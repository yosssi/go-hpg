package hpg

// FoodCategoryAPIResults は料理カテゴリマスタAPIのレスポンスデータを表す。
type FoodCategoryAPIResults struct {
	Results struct {
		CommonResults
		FoodCategory []MasterData `json:"food_category"`
	}
}
