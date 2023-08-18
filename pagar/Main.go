package pagar

import (
	"context"
	"fmt"
	"net/http"

	"MicroservicesPicPay/pagar/kafka"
	"MicroservicesPicPay/pagar/model"

	"encoding/json"

	qrcode "github.com/skip2/go-qrcode"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func generateQRCode(transacao model.TransacaoQRCode) (string, error) {
	// Formate os dados conforme as especificações do Banco Central
	data := fmt.Sprintf("ChavePix: %s, Valor: %f", transacao.ChavePix, transacao.Valor)

	// Gere o QRCode
	err := qrcode.WriteFile(data, qrcode.Medium, 256, "qrcode.png")
	if err != nil {
		return "", err
	}

	return data, nil
}

func validateBoleto(codigoBoleto string) (bool, error) {
	// Aqui, estamos apenas verificando se o código do boleto tem 44 caracteres, que é o padrão para boletos bancários.
	if len(codigoBoleto) == 44 {
		return true, nil
	}
	return false, nil
}

func validatePaymentData(pagamento model.PagamentoPessoa) (bool, error) {
	if pagamento.PagadorID == "" || pagamento.RecebedorID == "" || pagamento.Valor <= 0 {
		return false, nil
	}
	return true, nil
}

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

	// Endpoint: Pagar
	r.POST("/pagar", func(c *gin.Context) {
		var pagamento model.Pagamento
		if err := c.ShouldBindJSON(&pagamento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		collection := db.Collection("pagamentos")
		_, err := collection.InsertOne(context.TODO(), pagamento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir pagamento no banco de dados"})
			return
		}

		// Serializando pagamento para JSON
		pagamentoJSON, err := json.Marshal(pagamento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao serializar pagamento"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		err = kafka.SendMessage("pagamento", string(pagamentoJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Pagamento realizado!"})
	})

	// Endpoint: Notificações
	r.GET("/notificacoes", func(c *gin.Context) {
		// Consultando o MongoDB para obter notificações
		collection := db.Collection("notificacoes")
		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar notificações no banco de dados"})
			return
		}
		defer cursor.Close(context.TODO())

		var notificacoes []model.Notificacao
		if err = cursor.All(context.TODO(), &notificacoes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao decodificar notificações"})
			return
		}

		// Enviando mensagem para Kafka informando que as notificações foram consultadas
		msg := "Notificações consultadas"
		err = kafka.SendMessage("notificacoes", msg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"notificacoes": notificacoes})
	})

	// Endpoint: Pix
	r.POST("/pix", func(c *gin.Context) {
		var transacaoPix model.TransacaoPix // Declare uma variável do tipo TransacaoPix

		if err := c.ShouldBindJSON(&transacaoPix); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gravando no MongoDB
		collection := db.Collection("transacoes_pix")
		_, err := collection.InsertOne(context.TODO(), transacaoPix)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir transação Pix no banco de dados"})
			return
		}

		// Serializando transacaoPix para JSON
		transacaoJSON, err := json.Marshal(transacaoPix)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao serializar transação Pix"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		err = kafka.SendMessage("transacoes_pix", string(transacaoJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Transação Pix realizada com sucesso!"})
	})

	// Endpoint: Pagar com QRCODE
	r.POST("/pagar-qrcode", func(c *gin.Context) {
		var transacaoQR model.TransacaoQRCode

		if err := c.ShouldBindJSON(&transacaoQR); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gere o QRCode no formato do Banco Central do Brasil
		qrcodeData, err := generateQRCode(transacaoQR)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar QRCode"})
			return
		}
		transacaoQR.QRCodeData = qrcodeData

		// Gravando no MongoDB
		collection := db.Collection("transacoes_qrcode")
		_, err = collection.InsertOne(context.TODO(), transacaoQR)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir transação via QRCode no banco de dados"})
			return
		}

		// Serializando transacaoQR para JSON
		transacaoJSON, err := json.Marshal(transacaoQR)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao serializar transação via QRCode"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		err = kafka.SendMessage("transacoes_qrcode", string(transacaoJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Pagamento via QRCode realizado com sucesso!"})
	})

	// Endpoint: Pagar Boleto
	r.POST("/pagar-boleto", func(c *gin.Context) {
		var transacaoBoleto model.TransacaoBoleto

		if err := c.ShouldBindJSON(&transacaoBoleto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Valide o boleto por exemplo, verifique se o código do boleto é válido
		isValid, err := validateBoleto(transacaoBoleto.CodigoBoleto)
		if err != nil || !isValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Boleto inválido"})
			return
		}

		// Gravando no MongoDB
		collection := db.Collection("transacoes_boleto")
		_, err = collection.InsertOne(context.TODO(), transacaoBoleto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir transação via boleto no banco de dados"})
			return
		}

		// Serializando transacaoBoleto para JSON
		transacaoJSON, err := json.Marshal(transacaoBoleto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao serializar transação via boleto"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		err = kafka.SendMessage("transacoes_boleto", string(transacaoJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Boleto pago com sucesso!"})
	})

	// Endpoint: Pagar Pessoas
	r.POST("/pagar-pessoas", func(c *gin.Context) {
		var pagamento model.PagamentoPessoa

		if err := c.ShouldBindJSON(&pagamento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Valide os dados por exemplo, verifique se o ID do pagador e do recebedor são válidos
		isValid, err := validatePaymentData(pagamento)
		if err != nil || !isValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de pagamento inválidos"})
			return
		}

		// Gravando no MongoDB
		collection := db.Collection("pagamentos_pessoas")
		_, err = collection.InsertOne(context.TODO(), pagamento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir pagamento no banco de dados"})
			return
		}

		// Serializando pagamento para JSON
		pagamentoJSON, err := json.Marshal(pagamento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao serializar pagamento"})
			return
		}

		// Enviando mensagem para Kafka usando Confluent Kafka Go
		err = kafka.SendMessage("pagamentos_pessoas", string(pagamentoJSON))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem para Kafka"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Pagamento para pessoa realizado com sucesso!"})
	})

	r.Run(":8088")
}
