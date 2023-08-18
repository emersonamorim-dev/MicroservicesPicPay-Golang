package sacar

import (
	"MicroservicesPicPay/sacar/kafka"
	"MicroservicesPicPay/sacar/model"
	"context"
	"encoding/json"
	"net/http"
	"time"

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
	db := client.Database("SaqueDB")

	// Endpoint: Saque 24 horas
	r.POST("/saque", func(c *gin.Context) {
		var saque model.Saque
		if err := c.ShouldBindJSON(&saque); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Adicionando data e hora ao saque
		saque.DataHora = time.Now().Format(time.RFC3339)

		// Gravando no MongoDB
		_, err := db.Collection("saques").InsertOne(context.TODO(), saque)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar saque no banco de dados"})
			return
		}

		// Convertendo o objeto saque em uma string JSON
		saqueJSON, err := json.Marshal(saque)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter saque para JSON"})
			return
		}

		// Enviando mensagem para Kafka
		err = kafka.SendMessage("saques", string(saqueJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Saque registrado com sucesso!"})
	})

	// Endpoint: Saldo
	r.GET("/saldo/:usuarioID", func(c *gin.Context) {
		usuarioID := c.Param("usuarioID")

		// Criando um filtro para buscar o saldo pelo ID do usuário
		filter := bson.M{"usuarioID": usuarioID}

		// Consultando o MongoDB
		var saldo model.Saldo
		err := db.Collection("saldos").FindOne(context.TODO(), filter).Decode(&saldo)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{"error": "Saldo não encontrado para o usuário"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar o saldo no banco de dados"})
			return
		}

		c.JSON(http.StatusOK, saldo)
	})

	r.Run(":8089")
}
