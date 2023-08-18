package model

type Notificacao struct {
	ID       string `bson:"_id,omitempty"`
	Titulo   string `bson:"titulo"`
	Mensagem string `bson:"mensagem"`
	Data     string `bson:"data"`
}
