package carteira

import (
	"MicroservicesPicPay/carteira/kafka"
	"MicroservicesPicPay/carteira/model"
	"context"
	"fmt"
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
	db := client.Database("PicpayDB")

	// Endpoint: Saldo Disponível
	r.GET("/saldo", func(c *gin.Context) {
		// Lógica para obter saldo
		saldo := "7.800.00"

		// Gravando saldo no MongoDB
		collection := db.Collection("saldos")
		_, err := collection.InsertOne(context.TODO(), gin.H{"saldo": saldo})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir saldo no banco de dados"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		err = kafka.SendMessage("saldos", "Saldo consultado: "+saldo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(200, gin.H{"saldo": saldo})
	})

	// Endpoint: Meus Cartões
	r.GET("/cartoes", func(c *gin.Context) {
		collection := db.Collection("cartoes")
		cursor, err := collection.Find(context.TODO(), gin.H{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar cartões"})
			return
		}
		defer cursor.Close(context.TODO())

		var cartoes []model.Cartao
		for cursor.Next(context.TODO()) {
			var cartao model.Cartao
			if err = cursor.Decode(&cartao); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao decodificar cartão"})
				return
			}
			cartoes = append(cartoes, cartao)
		}

		c.JSON(200, gin.H{"cartoes": cartoes})
	})

	// Endpoint: Adicionar Cartão
	r.POST("/cartoes", func(c *gin.Context) {
		var cartao model.Cartao // Usando o modelo Cartao
		if err := c.ShouldBindJSON(&cartao); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		collection := db.Collection("cartoes")
		_, err := collection.InsertOne(context.TODO(), cartao)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir cartão no banco de dados"})
			return
		}
		c.JSON(201, gin.H{"message": "Cartão adicionado com sucesso!"})
	})

	// Endpoint: Minhas Chaves Pix
	r.GET("/chaves-pix", func(c *gin.Context) {
		// Lógica para listar chaves Pix
		chaves := []string{"Chave 1", "Chave 2"}

		// Gravando chaves no MongoDB
		collection := db.Collection("chaves_pix")
		_, err := collection.InsertOne(context.TODO(), gin.H{"chaves": chaves})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir chaves no banco de dados"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		err = kafka.SendMessage("chaves_pix", "Chaves Pix consultadas: "+chaves[0]+", "+chaves[1])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(200, gin.H{"chaves": chaves})
	})

	// Endpoint: Pagamentos
	r.POST("/pagamento", func(c *gin.Context) {
		var pagamento model.Pagamento // Declare uma variável do tipo model.Pagamento

		if err := c.ShouldBindJSON(&pagamento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		movimentacao := "Pagamento: " + pagamento.Descricao + " - Valor: " + fmt.Sprintf("%.2f", pagamento.Valor)
		collection := db.Collection("movimentacoes")
		_, err := collection.InsertOne(context.TODO(), gin.H{"movimentacao": movimentacao})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir movimentação no banco de dados"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		err = kafka.SendMessage("movimentacoes", movimentacao)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(200, gin.H{"message": "Pagamento registrado com sucesso!"})
	})

	// Endpoint: Transferências
	r.POST("/transferencia", func(c *gin.Context) {
		var transferencia model.Transferencia // Declare uma variável do tipo model.Transferencia

		if err := c.ShouldBindJSON(&transferencia); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		movimentacao := "Transferência: " + transferencia.Descricao + " - Valor: " + fmt.Sprintf("%.2f", transferencia.Valor)
		collection := db.Collection("movimentacoes")
		_, err := collection.InsertOne(context.TODO(), gin.H{"movimentacao": movimentacao})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir movimentação no banco de dados"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		err = kafka.SendMessage("movimentacoes", movimentacao)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(200, gin.H{"message": "Transferência registrada com sucesso!"})
	})

	// Endpoint: Últimas Movimentações
	r.GET("/movimentacoes", func(c *gin.Context) {
		collection := db.Collection("movimentacoes")
		cursor, err := collection.Find(context.TODO(), gin.H{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar movimentações"})
			return
		}
		defer cursor.Close(context.TODO())

		var movimentacoes []string
		for cursor.Next(context.TODO()) {
			var mov gin.H
			if err = cursor.Decode(&mov); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao decodificar movimentação"})
				return
			}
			movimentacoes = append(movimentacoes, mov["movimentacao"].(string))
		}

		c.JSON(200, gin.H{"movimentacoes": movimentacoes})
	})

	r.Run(":8081")
}
