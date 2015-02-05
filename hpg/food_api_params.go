package hpg

// FoodAPIParams は料理名マスタAPIの検索クエリパラメータを表す。
type FoodAPIParams struct {
	CommonParams
	Code         []string
	FoodCategory string
	Keyword      string
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *FoodAPIParams) Path() string {
	return "food"
}

// String は、クエリパラメータ文字列を生成し、それを返却する。
func (p *FoodAPIParams) String() string {
	bf := p.queryBuffer()

	bf.appendStringSlice("code", p.Code)
	bf.appendString("food_category", p.FoodCategory)
	bf.appendString("keyword", p.Keyword)

	return bf.String()
}
