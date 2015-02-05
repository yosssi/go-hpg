package hpg

// SmallAreaAPIParams は小エリアマスタAPIの検索クエリパラメータを表す。
type SmallAreaAPIParams struct {
	CommonParams
	SmallArea  []string
	MiddleArea []string
	Keyword    string
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *SmallAreaAPIParams) Path() string {
	return "small_area"
}

// String は、クエリパラメータ文字列を生成し、それを返却する。
func (p *SmallAreaAPIParams) String() string {
	bf := p.queryBuffer()

	bf.appendStringSlice("small_area", p.SmallArea)
	bf.appendStringSlice("middle_area", p.MiddleArea)
	bf.appendString("keyword", p.Keyword)

	return bf.String()
}
