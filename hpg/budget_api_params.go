package hpg

// BudgetAPIParams は検索用予算マスタAPIの検索クエリパラメータを表す。
type BudgetAPIParams struct {
	CommonParams
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *BudgetAPIParams) Path() string {
	return "budget"
}
