package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Definindo os serviços e suas portas
	services := map[string]string{
		"carteira":     "http://localhost:8081",
		"cobrar":       "http://localhost:8082",
		"crypto":       "http://localhost:8083",
		"emprestimos":  "http://localhost:8084",
		"finance":      "http://localhost:8085",
		"investir":     "http://localhost:8086",
		"notificações": "http://localhost:8087",
		"pagar":        "http://localhost:8088",
		"sacar":        "http://localhost:8089",
		"seguros":      "http://localhost:8090",
	}

	// Rota genérica para encaminhar as requisições para os microserviços
	r.Any("/*path", func(c *gin.Context) {
		service, exists := services[c.Param("path")]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Serviço não encontrado"})
			return
		}

		resp, err := http.Get(service)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao acessar o serviço"})
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler resposta do serviço"})
			return
		}

		// Enviar a resposta do serviço para o cliente
		for name, values := range resp.Header {
			for _, value := range values {
				c.Header(name, value)
			}
		}
		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
	})

	r.Run(":8094") // Inicie o servidor na porta 8094
}
