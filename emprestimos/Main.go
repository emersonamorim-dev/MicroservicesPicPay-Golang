package emprestimos

import (
	"MicroservicesPicPay/emprestimos/model"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func calcularEmprestimo(valor float64, parcelas int) float64 {
	taxaJuros := 0.02 // 2% taxa de juros mensal, por exemplo
	return valor * (1 + taxaJuros*float64(parcelas))
}

func calcularEmprestimoImovel(valor float64, parcelas int) (float64, float64) {
	taxaJuros := 0.02 // 2% taxa de juros mensal, por exemplo
	valorTotal := valor * (1 + taxaJuros*float64(parcelas))
	valorParcela := valorTotal / float64(parcelas)
	return valorParcela, valorTotal
}

func calcularEmprestimoFGTS(valor float64, parcelas int) (float64, float64) {
	taxaJuros := 0.015 // 1.5% taxa de juros mensal, por exemplo, para FGTS
	valorTotal := valor * (1 + taxaJuros*float64(parcelas))
	valorParcela := valorTotal / float64(parcelas)
	return valorParcela, valorTotal
}

func calcularParcelas(valorEmprestimo float64, parcelas int, taxaJuros float64) float64 {
	valorTotal := valorEmprestimo * (1 + taxaJuros*float64(parcelas))
	return valorTotal / float64(parcelas)
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

	// Endpoint: Empréstimo Garantia Veículo
	r.POST("/emprestimo-veiculo", func(c *gin.Context) {
		var emprestimo model.Emprestimo
		if err := c.ShouldBindJSON(&emprestimo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Calcula 90% do valor do veículo
		valorEmprestimo := emprestimo.Valor * 0.9
		emprestimo.ValorTotal = calcularEmprestimo(valorEmprestimo, emprestimo.Parcelas)

		// Gravando no MongoDB
		collection := db.Collection("emprestimos_veiculo")
		_, err := collection.InsertOne(context.TODO(), emprestimo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir empréstimo no banco de dados"})
			return
		}

		c.JSON(http.StatusOK, emprestimo)
	})

	// Endpoint: Empréstimo Garantia Imóvel
	r.POST("/emprestimo-imovel", func(c *gin.Context) {
		var emprestimo model.EmprestimoImovel
		if err := c.ShouldBindJSON(&emprestimo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Calcula até 90% do valor do imóvel
		valorEmprestimo := emprestimo.ValorImovel * 0.9
		emprestimo.ValorEmprestimo = valorEmprestimo
		emprestimo.ValorParcela, emprestimo.ValorTotal = calcularEmprestimoImovel(valorEmprestimo, emprestimo.Parcelas)

		// Gravando no MongoDB
		collection := db.Collection("emprestimos_imovel")
		_, err := collection.InsertOne(context.TODO(), emprestimo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir empréstimo no banco de dados"})
			return
		}

		c.JSON(http.StatusOK, emprestimo)
	})

	// Endpoint: Empréstimo Garantia FGTS
	r.POST("/emprestimo-fgts", func(c *gin.Context) {
		var emprestimo model.EmprestimoFGTS
		if err := c.ShouldBindJSON(&emprestimo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Calcula até 90% do valor do FGTS
		valorEmprestimo := emprestimo.ValorFGTS * 0.9
		emprestimo.ValorEmprestimo = valorEmprestimo
		emprestimo.ValorParcela, emprestimo.ValorTotal = calcularEmprestimoFGTS(valorEmprestimo, emprestimo.Parcelas)

		// Gravando no MongoDB
		collection := db.Collection("emprestimos_fgts")
		_, err := collection.InsertOne(context.TODO(), emprestimo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir empréstimo no banco de dados"})
			return
		}

		c.JSON(http.StatusOK, emprestimo)
	})

	// Endpoint: Cálculo Empréstimo
	r.GET("/calculo-emprestimo", func(c *gin.Context) {
		valorStr := c.DefaultQuery("valor", "0")
		parcelasStr := c.DefaultQuery("parcelas", "12")

		valor, err := strconv.ParseFloat(valorStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Valor inválido"})
			return
		}

		parcelas, err := strconv.Atoi(parcelasStr)
		if err != nil || (parcelas != 12 && parcelas != 18 && parcelas != 36 && parcelas != 60) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Número de parcelas inválido"})
			return
		}

		// Defina taxas de juros fictícias para cada tipo de empréstimo
		taxas := map[int]float64{
			12: 0.01,  // 1% ao mês
			18: 0.009, // 0.9% ao mês
			36: 0.008, // 0.8% ao mês
			60: 0.007, // 0.7% ao mês
		}

		valorParcela := calcularParcelas(valor, parcelas, taxas[parcelas])

		c.JSON(http.StatusOK, gin.H{
			"valor_emprestimo": valor,
			"parcelas":         parcelas,
			"valor_parcela":    valorParcela,
		})
	})

	r.Run(":8084")

}
