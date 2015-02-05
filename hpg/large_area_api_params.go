package hpg

// LargeAreaAPIParams は大エリアマスタAPIの検索クエリパラメータを表す。
type LargeAreaAPIParams struct {
	CommonParams
	LargeArea []string
	Keyword   string
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *LargeAreaAPIParams) Path() string {
	return "large_area"
}

// String は、クエリパラメータ文字列を生成し、それを返却する。
func (p *LargeAreaAPIParams) String() string {
	bf := p.queryBuffer()

	bf.appendStringSlice("large_area", p.LargeArea)
	bf.appendString("keyword", p.Keyword)

	return bf.String()
}
