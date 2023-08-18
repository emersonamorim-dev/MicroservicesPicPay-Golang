package model

type CobrancaLinkShare struct {
	ID        string  `bson:"_id,omitempty"`
	UsuarioID string  `json:"usuario_id"`
	Valor     float64 `json:"valor"`
	Descricao string  `json:"descricao"`
	Link      string  `bson:"link"`
}
