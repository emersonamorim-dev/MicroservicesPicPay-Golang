package model

type Compartilhamento struct {
	ID             string `bson:"_id,omitempty" json:"id"`
	UsuarioID      string `bson:"usuarioID" json:"usuarioID"`
	DestinatarioID string `bson:"destinatarioID" json:"destinatarioID"`
	Conteudo       string `bson:"conteudo" json:"conteudo"`
}
