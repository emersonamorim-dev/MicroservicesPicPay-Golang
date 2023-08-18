package cobrar

import (
	"MicroservicesPicPay/cobrar/model"
	"MicroservicesPicPay/cobrar/kafka"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
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
	db := client.Database("PicpayDB")

	// Endpoint: Cobrar Amigos
	r.POST("/cobrar-amigos", func(c *gin.Context) {
		var cobranca model.CobrancaAmigo
		if err := c.ShouldBindJSON(&cobranca); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		collection := db.Collection("cobrancas_amigos")
		_, err := collection.InsertOne(context.TODO(), cobranca)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir cobrança no banco de dados"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		message := fmt.Sprintf("Cobrança de %s para %s no valor de %.2f", cobranca.UsuarioID, cobranca.AmigoID, cobranca.Valor)
		err = kafka.SendMessage("cobrancas_amigos", message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Cobrança registrada com sucesso!"})
	})

	// Endpoint: Cobrar via Pix
	r.POST("/cobrar-pix", func(c *gin.Context) {
		var cobranca model.CobrancaPix
		if err := c.ShouldBindJSON(&cobranca); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		collection := db.Collection("cobrancas_pix")
		_, err := collection.InsertOne(context.TODO(), cobranca)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir cobrança no banco de dados"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		message := fmt.Sprintf("Cobrança via Pix de %s com chave %s no valor de %.2f", cobranca.UsuarioID, cobranca.ChavePix, cobranca.Valor)
		err = kafka.SendMessage("cobrancas_pix", message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Cobrança via Pix registrada com sucesso!"})
	})

	// Endpoint: Cobrar via QRCode
	r.POST("/cobrar-qrcode", func(c *gin.Context) {
		var cobranca model.CobrancaQRCode
		if err := c.ShouldBindJSON(&cobranca); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		collection := db.Collection("cobrancas_qrcode")
		_, err := collection.InsertOne(context.TODO(), cobranca)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir cobrança no banco de dados"})
			return
		}

		// Gerando QR Code conforme formato do Banco Central
		payload := fmt.Sprintf("[payload do Banco Central aqui]") // Substitua pelo formato correto do payload do Banco Central
		qrcode, err := qrcode.Encode(payload, qrcode.Medium, 256)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar QR Code"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		message := fmt.Sprintf("Cobrança via QR Code de %s com chave %s no valor de %.2f", cobranca.UsuarioID, cobranca.ChavePix, cobranca.Valor)
		err = kafka.SendMessage("cobrancas_qrcode", message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Cobrança via QR Code registrada com sucesso!", "qrcode": qrcode})
	})

	// Endpoint: Cobrar via Link Compartilhável
	r.POST("/cobrar-linkshare", func(c *gin.Context) {
		var cobranca model.CobrancaLinkShare
		if err := c.ShouldBindJSON(&cobranca); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gerando um link compartilhável único
		cobranca.Link = fmt.Sprintf("https://picpay.com/cobranca/%s", uuid.New().String())

		// Gravando no MongoDB
		collection := db.Collection("cobrancas_linkshare")
		_, err := collection.InsertOne(context.TODO(), cobranca)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir cobrança no banco de dados"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		message := fmt.Sprintf("Cobrança via Link Compartilhável de %s no valor de %.2f", cobranca.UsuarioID, cobranca.Valor)
		err = kafka.SendMessage("cobrancas_linkshare", message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Cobrança via Link Compartilhável registrada com sucesso!", "link": cobranca.Link})
	})

	r.Run(":8082")
}
