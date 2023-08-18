package crypto

import (
	"MicroservicesPicPay/crypto/model"
	"MicroservicesPicPay/crypto/kafka"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
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

	// Endpoint: Comprar Crypto
	r.POST("/comprar-crypto", func(c *gin.Context) {
		var compra model.CompraCrypto
		if err := c.ShouldBindJSON(&compra); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		collection := db.Collection("compras_crypto")
		_, err := collection.InsertOne(context.TODO(), compra)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir compra no banco de dados"})
			return
		}

		// Enviando mensagem para Kafka
		message := fmt.Sprintf("Compra de %s por %s de %s %s ao preço de %.2f", compra.Moeda, compra.UsuarioID, fmt.Sprintf("%.2f", compra.Quantidade), compra.Moeda, compra.Preco)
		err = kafka.SendMessage("compras_crypto", message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Compra de criptomoeda registrada com sucesso!"})
	})

	// Endpoint: Minhas Transações
	r.GET("/minhas-transacoes/:usuario_id", func(c *gin.Context) {
		usuarioID := c.Param("usuario_id")

		// Consultando MongoDB
		collection := db.Collection("compras_crypto")
		cursor, err := collection.Find(context.TODO(), gin.H{"usuario_id": usuarioID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar transações"})
			return
		}

		var transacoes []model.Transacao
		if err = cursor.All(context.TODO(), &transacoes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler transações"})
			return
		}

		c.JSON(http.StatusOK, transacoes)
	})

	// Endpoint: Alerta de Preços
	r.POST("/alerta-precos", func(c *gin.Context) {
		var alerta model.AlertaPreco
		if err := c.ShouldBindJSON(&alerta); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		collection := db.Collection("alertas_preco")
		_, err := collection.InsertOne(context.TODO(), alerta)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir alerta no banco de dados"})
			return
		}

		// Enviando mensagem para Kafka
		message := fmt.Sprintf("Alerta de preço para %s definido por %s ao preço de %.2f", alerta.Moeda, alerta.UsuarioID, alerta.PrecoAlvo)
		err = kafka.SendMessage("alertas_preco", message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Alerta de preço definido com sucesso!"})
	})

	// Endpoint: Moeda e Preço
	r.GET("/moeda-preco/:moeda", func(c *gin.Context) {
		moeda := c.Param("moeda")

		client := resty.New()
		resp, err := client.R().
			SetHeader("Accepts", "application/json").
			SetHeader("X-CMC_PRO_API_KEY", "YOUR_COINMARKETCAP_API_KEY").
			Get("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar CoinMarketcap"})
			return
		}

		var result map[string]interface{}
		if err := json.Unmarshal(resp.Body(), &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar resposta do CoinMarketcap"})
			return
		}

		data := result["data"].([]interface{})
		var preco float64
		for _, item := range data {
			crypto := item.(map[string]interface{})
			if crypto["symbol"].(string) == moeda {
				quote := crypto["quote"].(map[string]interface{})
				usd := quote["USD"].(map[string]interface{})
				preco = usd["price"].(float64)
				break
			}
		}

		if preco == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Moeda não encontrada"})
			return
		}

		// Gravando no MongoDB
		collection := db.Collection("precos_moedas")
		_, err = collection.InsertOne(context.TODO(), gin.H{"moeda": moeda, "preco": preco})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir preço no banco de dados"})
			return
		}

		// Enviando mensagem para Kafka
		message := fmt.Sprintf("Preço atual de %s é %.2f", moeda, preco)
		err = kafka.SendMessage("precos_moedas", message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"moeda": moeda, "preco": preco})
	})

	r.Run(":8083")
}
