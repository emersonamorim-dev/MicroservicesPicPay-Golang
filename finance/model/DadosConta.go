package model

type DadosConta struct {
	Agencia   string  `json:"agencia"`
	Conta     string  `json:"conta"`
	TipoConta string  `json:"tipo_conta"`
	Saldo     float64 `json:"saldo"`
}
