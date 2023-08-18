package model

type Emprestimo struct {
	Tipo       string  `json:"tipo"`
	Valor      float64 `json:"valor"`
	Parcelas   int     `json:"parcelas"`
	ValorTotal float64 `json:"valor_total"`
}
