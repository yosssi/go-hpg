package hpg

// CreditCardAPIParams はクレジットカードマスタAPIの検索クエリパラメータを表す。
type CreditCardAPIParams struct {
	CommonParams
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *CreditCardAPIParams) Path() string {
	return "credit_card"
}
