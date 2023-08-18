package model

type DadosOperacoesCredito struct {
	Tipo         string  `json:"tipo"`
	Valor        float64 `json:"valor"`
	DataOperacao string  `json:"data_operacao"`
}
