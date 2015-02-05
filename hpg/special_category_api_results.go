package hpg

// SpecialCategoryAPIResults は特集カテゴリマスタAPIのレスポンスデータを表す。
type SpecialCategoryAPIResults struct {
	Results struct {
		CommonResults
		SpecialCategory []MasterData `json:"special_category"`
	}
}
