package model

type PerfilInvestidor struct {
	CPF    string `json:"cpf"`
	Perfil string `json:"perfil"` // Conservador, Moderado, Agressivo, etc.
}
