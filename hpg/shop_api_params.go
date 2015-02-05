package hpg

// ShopAPIParams は店名サーチAPIの検索クエリパラメータを表す。
type ShopAPIParams struct {
	CommonParams
	Keyword []string
	TEL     string
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *ShopAPIParams) Path() string {
	return "shop"
}

// String は、クエリパラメータ文字列を生成し、それを返却する。
func (p *ShopAPIParams) String() string {
	bf := p.queryBuffer()

	bf.appendStringSlice("keyword", p.Keyword)
	bf.appendString("tel", p.TEL)

	return bf.String()
}
