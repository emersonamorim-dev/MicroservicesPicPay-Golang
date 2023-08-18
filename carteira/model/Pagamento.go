package model

type Pagamento struct {
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}
