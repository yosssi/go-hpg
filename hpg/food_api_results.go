package hpg

// FoodAPIResults は料理名マスタAPIのレスポンスデータを表す。
type FoodAPIResults struct {
	Results struct {
		CommonResults
		Food []struct {
			MasterData
			FoodCategory MasterData `json:"food_category"`
		} `json:"food"`
	}
}
