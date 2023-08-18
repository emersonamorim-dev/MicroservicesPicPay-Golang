package notificacoes

import (
	"MicroservicesPicPay/notificacoes/kafka"
	"MicroservicesPicPay/notificacoes/model"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	Start()
}

func Start() {
	r := gin.Default()

	// Conexão com MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		panic(err)
	}
	db := client.Database("NotificacoesDB")

	// Endpoint: Notificações
	r.GET("/notificacoes/:usuarioID", func(c *gin.Context) {
		usuarioID := c.Param("usuarioID")
		var notificacoes []model.Notificacao

		cursor, err := db.Collection("notificacoes").Find(context.TODO(), bson.M{"usuarioID": usuarioID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar notificações"})
			return
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var notificacao model.Notificacao
			cursor.Decode(&notificacao)
			notificacoes = append(notificacoes, notificacao)
		}

		c.JSON(http.StatusOK, notificacoes)
	})

	// Endpoint: Compartilhar
	r.POST("/compartilhar", func(c *gin.Context) {
		var compartilhamento model.Compartilhamento
		if err := c.ShouldBindJSON(&compartilhamento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		_, err := db.Collection("compartilhamentos").InsertOne(context.TODO(), compartilhamento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir compartilhamento no banco de dados"})
			return
		}

		// Serializando compartilhamento para JSON
		compartilhamentoJSON, err := json.Marshal(compartilhamento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao serializar compartilhamento"})
			return
		}

		// Enviando mensagem para Kafka
		err = kafka.SendMessage("compartilhamentos", string(compartilhamentoJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Compartilhamento registrado com sucesso!"})
	})

	r.Run(":8087")
}
