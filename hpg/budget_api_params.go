package hpg

import "bytes"

// BudgetAPIParams は検索用予算マスタAPIの検索クエリパラメータを表す。
type BudgetAPIParams struct {
	CommonParams
}

// Query は、クエリパラメータ文字列を生成し、それを返却する。
func (p *BudgetAPIParams) Query() string {
	bf := new(bytes.Buffer)

	p.CommonParams.WriteStringTo(bf)

	return bf.String()
}
