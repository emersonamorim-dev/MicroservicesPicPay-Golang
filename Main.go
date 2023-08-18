package main

import (
	"sync"

	"MicroservicesPicPay/carteira"
	"MicroservicesPicPay/cobrar"
	"MicroservicesPicPay/crypto"
	"MicroservicesPicPay/emprestimos"
	"MicroservicesPicPay/finance"
	"MicroservicesPicPay/investir"
	"MicroservicesPicPay/notificacoes"
	"MicroservicesPicPay/pagar"
	"MicroservicesPicPay/sacar"
	"MicroservicesPicPay/seguros"
)

func main() {
	var wg sync.WaitGroup

	// Inicie o serviço Carteira
	wg.Add(1)
	go func() {
		defer wg.Done()
		carteira.Start()
	}()

	// Inicie o serviço Cobrar
	wg.Add(1)
	go func() {
		defer wg.Done()
		cobrar.Start()
	}()

	// Inicie o serviço Crypto
	wg.Add(1)
	go func() {
		defer wg.Done()
		crypto.Start()
	}()

	// Inicie o serviço Empréstimos
	wg.Add(1)
	go func() {
		defer wg.Done()
		emprestimos.Start()
	}()

	// Inicie o serviço Finance
	wg.Add(1)
	go func() {
		defer wg.Done()
		finance.Start()
	}()

	// Inicie o serviço Investir
	wg.Add(1)
	go func() {
		defer wg.Done()
		investir.Start()
	}()

	// Inicie o serviço Notificações
	wg.Add(1)
	go func() {
		defer wg.Done()
		notificacoes.Start()
	}()

	// Inicie o serviço Pagar
	wg.Add(1)
	go func() {
		defer wg.Done()
		pagar.Start()
	}()

	// Inicie o serviço Sacar
	wg.Add(1)
	go func() {
		defer wg.Done()
		sacar.Start()
	}()

	// Inicie o serviço Seguros
	wg.Add(1)
	go func() {
		defer wg.Done()
		seguros.Start()
	}()

	wg.Wait() // Espere todos os serviços terminarem
}
