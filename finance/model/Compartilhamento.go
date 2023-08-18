package model

type Compartilhamento struct {
	DadosCadastrais       DadosCadastrais         `json:"dados_cadastrais"`
	DadosConta            DadosConta              `json:"dados_conta"`
	DadosCartaoCredito    DadosCartaoCredito      `json:"dados_cartao_credito"`
	DadosOperacoesCredito []DadosOperacoesCredito `json:"dados_operacoes_credito"`
}
