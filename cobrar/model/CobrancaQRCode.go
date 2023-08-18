package model

type CobrancaQRCode struct {
	ID        string  `bson:"_id,omitempty"`
	UsuarioID string  `json:"usuario_id"`
	ChavePix  string  `json:"chave_pix"`
	Valor     float64 `json:"valor"`
	Descricao string  `json:"descricao"`
}
