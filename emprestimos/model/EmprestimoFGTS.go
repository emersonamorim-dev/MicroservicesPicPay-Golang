package model

type EmprestimoFGTS struct {
	UsuarioID       string  `json:"usuario_id"`
	ValorFGTS       float64 `json:"valor_fgts"`
	ValorEmprestimo float64 `json:"valor_emprestimo"`
	Parcelas        int     `json:"parcelas"`
	ValorParcela    float64 `json:"valor_parcela"`
	ValorTotal      float64 `json:"valor_total"`
}
