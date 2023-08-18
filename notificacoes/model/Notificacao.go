package model

import (
	"time"
)

type Notificacao struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	UsuarioID string    `bson:"usuarioID" json:"usuarioID"`
	Titulo    string    `bson:"titulo" json:"titulo"`
	Mensagem  string    `bson:"mensagem" json:"mensagem"`
	Data      time.Time `bson:"data" json:"data"`
	Lida      bool      `bson:"lida" json:"lida"`
}
