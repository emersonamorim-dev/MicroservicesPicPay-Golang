package model

type TransacaoQRCode struct {
	ID         string  `bson:"_id,omitempty"`
	ChavePix   string  `json:"chave_pix"`
	Descricao  string  `json:"descricao"`
	Valor      float64 `json:"valor"`
	Emissor    string  `json:"emissor"`
	Data       string  `json:"data"`
	QRCodeData string  `json:"qrcode_data"`
}
