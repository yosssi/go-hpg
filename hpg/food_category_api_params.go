package hpg

// FoodCategoryAPIParams は料理カテゴリマスタAPIの検索クエリパラメータを表す。
type FoodCategoryAPIParams struct {
	CommonParams
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *FoodCategoryAPIParams) Path() string {
	return "food_category"
}
