package model

type Saque struct {
	UsuarioID             string  `json:"usuarioID"`
	InstituicaoFinanceira string  `json:"instituicaoFinanceira"`
	ValorSacado           float64 `json:"valorSacado"`
	DataHora              string  `json:"dataHora"`
}
