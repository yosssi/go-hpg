package hpg

// MiddleAreaAPIParams は中エリアマスタAPIの検索クエリパラメータを表す。
type MiddleAreaAPIParams struct {
	CommonParams
	MiddleArea []string
	LargeArea  []string
	Keyword    string
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *MiddleAreaAPIParams) Path() string {
	return "middle_area"
}

// String は、クエリパラメータ文字列を生成し、それを返却する。
func (p *MiddleAreaAPIParams) String() string {
	bf := p.queryBuffer()

	bf.appendStringSlice("middle_area", p.MiddleArea)
	bf.appendStringSlice("large_area", p.LargeArea)
	bf.appendString("keyword", p.Keyword)

	return bf.String()
}
