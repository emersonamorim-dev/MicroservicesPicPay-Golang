package finance

import (
	"MicroservicesPicPay/finance/model"
	"context"
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

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		panic(err)
	}
	db := client.Database("OpenFinanceDB")

	// Endpoint: Trazer Dados
	r.GET("/dados/:cpf", func(c *gin.Context) {
		cpf := c.Param("cpf")
		var compartilhamento model.Compartilhamento
		err := db.Collection("dados_compartilhados").FindOne(context.TODO(), gin.H{"dados_cadastrais.cpf": cpf}).Decode(&compartilhamento)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar dados"})
			return
		}
		c.JSON(http.StatusOK, compartilhamento)
	})

	// Endpoint: Buscar Instituição
	r.GET("/instituicao/:id", func(c *gin.Context) {
		id := c.Param("id")
		var instituicao model.Instituicao
		err := db.Collection("instituicoes").FindOne(context.TODO(), gin.H{"id": id}).Decode(&instituicao)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar instituição"})
			return
		}
		c.JSON(http.StatusOK, instituicao)
	})

	// Endpoint: Mais Procurados (por simplicidade, retornaremos todas as instituições)
	r.GET("/mais-procurados", func(c *gin.Context) {
		cursor, err := db.Collection("instituicoes").Find(context.TODO(), gin.H{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar instituições"})
			return
		}
		var instituicoes []model.Instituicao
		if err = cursor.All(context.TODO(), &instituicoes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar instituições"})
			return
		}
		c.JSON(http.StatusOK, instituicoes)
	})

	r.Run(":8085")
}
