package seguros

import (
	"MicroservicesPicPay/seguros/kafka"
	"MicroservicesPicPay/seguros/model"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
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
	db := client.Database("SeguroDB")

	// Endpoint: Seguro de Vida
	r.POST("/seguro-vida", func(c *gin.Context) {
		var seguro model.SeguroVida
		if err := c.ShouldBindJSON(&seguro); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		_, err := db.Collection("seguros_vida").InsertOne(context.TODO(), seguro)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar seguro de vida no banco de dados"})
			return
		}

		// Convertendo o objeto seguro em uma string JSON
		seguroJSON, err := json.Marshal(seguro)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter seguro para JSON"})
			return
		}

		// Enviando mensagem para Kafka
		err = kafka.SendMessage("seguros_vida", string(seguroJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Seguro de vida registrado com sucesso!"})
	})

	// Endpoint: Seguro Carteira Digital
	r.POST("/seguro-carteira-digital", func(c *gin.Context) {
		var seguro model.SeguroCarteiraDigital
		if err := c.ShouldBindJSON(&seguro); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		_, err := db.Collection("seguros_carteira_digital").InsertOne(context.TODO(), seguro)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar seguro de carteira digital no banco de dados"})
			return
		}

		// Convertendo o objeto seguro em uma string JSON
		seguroJSON, err := json.Marshal(seguro)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter seguro para JSON"})
			return
		}

		// Enviando mensagem para Kafka
		err = kafka.SendMessage("seguros_carteira_digital", string(seguroJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Seguro de carteira digital registrado com sucesso!"})
	})

	// Endpoint: Seguro Celular
	r.POST("/seguro-celular", func(c *gin.Context) {
		var seguro model.SeguroCelular
		if err := c.ShouldBindJSON(&seguro); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		_, err := db.Collection("seguros_celular").InsertOne(context.TODO(), seguro)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar seguro de celular no banco de dados"})
			return
		}

		// Convertendo o objeto seguro em uma string JSON
		seguroJSON, err := json.Marshal(seguro)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter seguro para JSON"})
			return
		}

		// Enviando mensagem para Kafka
		err = kafka.SendMessage("seguros_celular", string(seguroJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Seguro de celular registrado com sucesso!"})
	})

	r.Run(":8090")
}
