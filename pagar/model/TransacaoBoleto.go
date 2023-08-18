package model

type TransacaoBoleto struct {
	ID            string  `bson:"_id,omitempty"`
	CodigoBoleto  string  `json:"codigo_boleto"`
	Valor         float64 `json:"valor"`
	DataPagamento string  `json:"data_pagamento"`
	Emissor       string  `json:"emissor"`
}
