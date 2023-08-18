package model

type DadosCartaoCredito struct {
	Numero           string  `json:"numero"`
	Validade         string  `json:"validade"`
	Limite           float64 `json:"limite"`
	LimiteDisponivel float64 `json:"limite_disponivel"`
}
