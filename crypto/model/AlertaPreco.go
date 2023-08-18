package model

type AlertaPreco struct {
	UsuarioID string  `json:"usuario_id"`
	Moeda     string  `json:"moeda"`
	PrecoAlvo float64 `json:"preco_alvo"`
}
