package model

type SeguroVida struct {
	UsuarioID            string  `json:"usuarioID"`
	ProtecaoMorte        float64 `json:"protecaoMorte"`
	ProtecaoInvalidez    float64 `json:"protecaoInvalidez"`
	ProtecaoDoencas      float64 `json:"protecaoDoencas"`
	InternacaoHospitalar bool    `json:"internacaoHospitalar"`
	ValorPorMes          float64 `json:"valorPorMes"`
	FormaPagamento       string  `json:"formaPagamento"` // "cartao" ou "saldo"
}
