package model

type EmprestimoImovel struct {
	UsuarioID       string  `json:"usuario_id"`
	ValorImovel     float64 `json:"valor_imovel"`
	ValorEmprestimo float64 `json:"valor_emprestimo"`
	Parcelas        int     `json:"parcelas"`
	ValorParcela    float64 `json:"valor_parcela"`
	ValorTotal      float64 `json:"valor_total"`
}
