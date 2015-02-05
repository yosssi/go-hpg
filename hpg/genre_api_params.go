package hpg

// GenreAPIParams はジャンルマスタAPIの検索クエリパラメータを表す。
type GenreAPIParams struct {
	CommonParams
	Code    []string
	Keyword string
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *GenreAPIParams) Path() string {
	return "genre"
}

// String は、クエリパラメータ文字列を生成し、それを返却する。
func (p *GenreAPIParams) String() string {
	bf := p.queryBuffer()

	bf.appendStringSlice("code", p.Code)
	bf.appendString("keyword", p.Keyword)

	return bf.String()
}
