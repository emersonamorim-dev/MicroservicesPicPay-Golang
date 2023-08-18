package model

type TransacaoPix struct {
	ID           string  `bson:"_id,omitempty"`
	ChavePix     string  `json:"chave_pix"`
	Descricao    string  `json:"descricao"`
	Valor        float64 `json:"valor"`
	Destinatario string  `json:"destinatario"`
	Data         string  `json:"data"`
}
