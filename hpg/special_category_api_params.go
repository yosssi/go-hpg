package hpg

// SpecialCategoryAPIParams は特集カテゴリマスタAPIの検索クエリパラメータを表す。
type SpecialCategoryAPIParams struct {
	CommonParams
	SpecialCategory []string
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *SpecialCategoryAPIParams) Path() string {
	return "special_category"
}

// String は、クエリパラメータ文字列を生成し、それを返却する。
func (p *SpecialCategoryAPIParams) String() string {
	bf := p.queryBuffer()

	bf.appendStringSlice("special_category", p.SpecialCategory)

	return bf.String()
}
