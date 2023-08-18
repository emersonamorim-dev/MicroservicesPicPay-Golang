package model

type SeguroCelular struct {
	UsuarioID      string  `json:"usuarioID"`
	ModeloCelular  string  `json:"modeloCelular"`
	ValorCobertura float64 `json:"valorCobertura"`
	ValorMensal    float64 `json:"valorMensal"`
	FormaPagamento string  `json:"formaPagamento"` // "cartao" ou "saldo"
}
