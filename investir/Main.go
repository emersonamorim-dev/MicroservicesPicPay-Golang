package investir

import (
	"MicroservicesPicPay/investir/model"
	"context"
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

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		panic(err)
	}
	db := client.Database("InvestirDB")

	// Endpoint: Investir em CDB PicPay Bank
	r.POST("/investir-cdb", func(c *gin.Context) {
		var investimento model.Investimento
		if err := c.ShouldBindJSON(&investimento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		investimento.Tipo = "CDB PicPay Bank"
		_, err := db.Collection("investimentos").InsertOne(context.TODO(), investimento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir investimento"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Investimento em CDB PicPay Bank realizado com sucesso!"})
	})

	// Endpoint: Investir em Crypto
	r.POST("/investir-crypto", func(c *gin.Context) {
		var investimento model.Investimento
		if err := c.ShouldBindJSON(&investimento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		investimento.Tipo = "Crypto"
		_, err := db.Collection("investimentos").InsertOne(context.TODO(), investimento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir investimento"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Investimento em Crypto realizado com sucesso!"})
	})

	// Endpoint: Investir em Empresas
	r.POST("/investir-empresas", func(c *gin.Context) {
		var investimento model.Investimento
		if err := c.ShouldBindJSON(&investimento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		investimento.Tipo = "Empresa"
		_, err := db.Collection("investimentos").InsertOne(context.TODO(), investimento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir investimento"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Investimento em Empresa realizado com sucesso!"})
	})

	// Endpoint: Investir em Pessoas
	r.POST("/investir-pessoas", func(c *gin.Context) {
		var investimento model.Investimento
		if err := c.ShouldBindJSON(&investimento); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		investimento.Tipo = "Pessoa"
		_, err := db.Collection("investimentos").InsertOne(context.TODO(), investimento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir investimento"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Investimento em Pessoa realizado com sucesso!"})
	})

	// Endpoint: Perfil do Investidor
	r.GET("/perfil-investidor/:cpf", func(c *gin.Context) {
		cpf := c.Param("cpf")
		var perfil model.PerfilInvestidor

		err := db.Collection("perfis_investidores").FindOne(context.TODO(), bson.M{"cpf": cpf}).Decode(&perfil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar perfil do investidor"})
			return
		}

		c.JSON(http.StatusOK, perfil)
	})

	// Endpoint: Total Investido
	r.GET("/total-investido/:cpf", func(c *gin.Context) {
		cpf := c.Param("cpf")
		var investimentos []model.Investimento
		total := 0.0

		cursor, err := db.Collection("investimentos").Find(context.TODO(), bson.M{"id": cpf})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar investimentos"})
			return
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var investimento model.Investimento
			cursor.Decode(&investimento)
			investimentos = append(investimentos, investimento)
			total += investimento.Valor
		}

		c.JSON(http.StatusOK, gin.H{"total_investido": total})
	})

	r.Run(":8086")
}
