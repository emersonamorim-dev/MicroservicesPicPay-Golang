package model

type CobrancaAmigo struct {
	ID        string  `bson:"_id,omitempty"`
	UsuarioID string  `json:"usuario_id"`
	AmigoID   string  `json:"amigo_id"`
	Valor     float64 `json:"valor"`
	Descricao string  `json:"descricao"`
}
