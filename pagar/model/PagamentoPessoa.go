package model

type PagamentoPessoa struct {
	ID            string  `bson:"_id,omitempty"`
	RecebedorID   string  `json:"recebedor_id"`
	PagadorID     string  `json:"pagador_id"`
	Valor         float64 `json:"valor"`
	DataPagamento string  `json:"data_pagamento"`
	Descricao     string  `json:"descricao"`
}
