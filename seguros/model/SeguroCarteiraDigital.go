package model

type SeguroCarteiraDigital struct {
	UsuarioID      string  `json:"usuarioID"`
	ValorCoberto   float64 `json:"valorCoberto"`
	DuracaoMeses   int     `json:"duracaoMeses"`
	ValorPorMes    float64 `json:"valorPorMes"`
	FormaPagamento string  `json:"formaPagamento"` // "cartao" ou "saldo"
}
