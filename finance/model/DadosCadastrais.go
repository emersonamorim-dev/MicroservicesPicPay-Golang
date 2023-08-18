package model

type DadosCadastrais struct {
	Nome     string `json:"nome"`
	CPF      string `json:"cpf"`
	Endereco string `json:"endereco"`
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}
