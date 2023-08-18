package model

type Pagamento struct {
	ID           string  `bson:"_id,omitempty"`
	Descricao    string  `bson:"descricao"`
	Valor        float64 `bson:"valor"`
	Destinatario string  `bson:"destinatario"`
}
