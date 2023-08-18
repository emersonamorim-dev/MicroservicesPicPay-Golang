package model

type Transacao struct {
	UsuarioID  string  `json:"usuario_id"`
	Moeda      string  `json:"moeda"`
	Quantidade float64 `json:"quantidade"`
	Preco      float64 `json:"preco"`
	Data       string  `json:"data"`
}
