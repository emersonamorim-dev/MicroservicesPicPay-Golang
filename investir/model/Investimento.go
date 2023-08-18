// model/investir.go
package model

type Investimento struct {
	ID      string  `json:"id"`
	Tipo    string  `json:"tipo"` // CDB, Crypto, Empresa, Pessoa, etc.
	Valor   float64 `json:"valor"`
	Retorno float64 `json:"retorno"` // Percentual de retorno esperado
}
