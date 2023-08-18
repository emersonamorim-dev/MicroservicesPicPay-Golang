package model

type Cobranca struct {
	ID           string  `bson:"_id,omitempty"`
	UsuarioID    string  `json:"usuario_id"`
	Valor        float64 `json:"valor"`
	Descricao    string  `json:"descricao"`
	TipoCobranca string  `json:"tipo_cobranca"` // Amigos, Pix, QRCode, Link Compartilhado
}
