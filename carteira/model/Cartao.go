package model

type Cartao struct {
	ID     string `bson:"_id,omitempty"`
	Nome   string `bson:"nome"`
	Numero string `bson:"numero"`
}
