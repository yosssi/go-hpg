package hpg

// SpecialAPIParams は特集マスタAPIの検索クエリパラメータを表す。
type SpecialAPIParams struct {
	CommonParams
	Special         []string
	SpecialCategory []string
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *SpecialAPIParams) Path() string {
	return "special"
}

// String は、クエリパラメータ文字列を生成し、それを返却する。
func (p *SpecialAPIParams) String() string {
	bf := p.queryBuffer()

	bf.appendStringSlice("special", p.Special)
	bf.appendStringSlice("special_category", p.SpecialCategory)

	return bf.String()
}
