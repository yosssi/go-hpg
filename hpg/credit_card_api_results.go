package hpg

// CreditCardAPIResults はクレジットカードマスタAPIのレスポンスデータを表す。
type CreditCardAPIResults struct {
	Results struct {
		CommonResults
		CreditCard []MasterData `json:"credit_card"`
	}
}
