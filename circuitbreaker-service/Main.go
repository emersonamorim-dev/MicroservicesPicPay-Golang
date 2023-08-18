package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

var cbSettings = gobreaker.Settings{
	Interval:    60 * time.Second,
	Timeout:     2 * time.Second,
	MaxRequests: 1,
	ReadyToTrip: func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	},
}

var cb = gobreaker.NewCircuitBreaker(cbSettings)

func redirectToService(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	r := gin.Default()

	// Carteira
	r.GET("/carteira/:action", func(c *gin.Context) {
		action := c.Param("action")
		url := "http://localhost:8081/" + action
		_, err := cb.Execute(func() (interface{}, error) {
			return redirectToService(url)
		})
		handleCircuitBreakerResponse(c, err, url)
	})

	// Cobrar
	r.GET("/cobrar/:action", func(c *gin.Context) {
		action := c.Param("action")
		url := "http://localhost:8082/" + action
		_, err := cb.Execute(func() (interface{}, error) {
			return redirectToService(url)
		})
		handleCircuitBreakerResponse(c, err, url)
	})

	// Crypto
	r.GET("/crypto/:action", func(c *gin.Context) {
		action := c.Param("action")
		url := "http://localhost:8083/" + action
		_, err := cb.Execute(func() (interface{}, error) {
			return redirectToService(url)
		})
		handleCircuitBreakerResponse(c, err, url)
	})

	// Empréstimo
	r.GET("/emprestimo/:action", func(c *gin.Context) {
		action := c.Param("action")
		url := "http://localhost:8084/" + action
		_, err := cb.Execute(func() (interface{}, error) {
			return redirectToService(url)
		})
		handleCircuitBreakerResponse(c, err, url)
	})

	// Finance
	r.GET("/finance/:action", func(c *gin.Context) {
		action := c.Param("action")
		url := "http://localhost:8085/" + action
		_, err := cb.Execute(func() (interface{}, error) {
			return redirectToService(url)
		})
		handleCircuitBreakerResponse(c, err, url)
	})

	// Investir
	r.GET("/investir/:action", func(c *gin.Context) {
		action := c.Param("action")
		url := "http://localhost:8086/" + action
		_, err := cb.Execute(func() (interface{}, error) {
			return redirectToService(url)
		})
		handleCircuitBreakerResponse(c, err, url)
	})

	// Notificações
	r.GET("/notificacoes/:action", func(c *gin.Context) {
		action := c.Param("action")
		url := "http://localhost:8087/" + action
		_, err := cb.Execute(func() (interface{}, error) {
			return redirectToService(url)
		})
		handleCircuitBreakerResponse(c, err, url)
	})

	// Pagar
	r.GET("/pagar/:action", func(c *gin.Context) {
		action := c.Param("action")
		url := "http://localhost:8088/" + action
		_, err := cb.Execute(func() (interface{}, error) {
			return redirectToService(url)
		})
		handleCircuitBreakerResponse(c, err, url)
	})

	// Sacar
	r.GET("/sacar/:action", func(c *gin.Context) {
		action := c.Param("action")
		url := "http://localhost:8089/" + action
		_, err := cb.Execute(func() (interface{}, error) {
			return redirectToService(url)
		})
		handleCircuitBreakerResponse(c, err, url)
	})

	// Seguros
	r.GET("/seguros/:action", func(c *gin.Context) {
		action := c.Param("action")
		url := "http://localhost:8090/" + action
		_, err := cb.Execute(func() (interface{}, error) {
			return redirectToService(url)
		})
		handleCircuitBreakerResponse(c, err, url)
	})

	r.Run(":8093") // Porta do circuitbreaker-service
}

func handleCircuitBreakerResponse(c *gin.Context, err error, url string) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service unavailable"})
		return
	}
	http.Redirect(c.Writer, c.Request, url, http.StatusSeeOther)
}
