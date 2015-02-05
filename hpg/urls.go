package hpg

// URLs は店舗URLを表す。
type URLs struct {
	PC     string `json:"pc"`
	Mobile string `json:"mobile"`
	QR     string `json:"qr"`
}
