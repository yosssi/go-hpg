package hpg

// MiddleAreaAPIResults は中エリアマスタAPIのレスポンスデータを表す。
type MiddleAreaAPIResults struct {
	Results struct {
		CommonResults
		MiddleArea []MasterData `json:"middle_area"`
	}
}
