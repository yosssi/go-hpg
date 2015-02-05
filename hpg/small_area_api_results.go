package hpg

// SmallAreaAPIResults は小エリアマスタAPIのレスポンスデータを表す。
type SmallAreaAPIResults struct {
	Results struct {
		CommonResults
		SmallArea []MasterData `json:"small_area"`
	}
}
