package hpg

// GenreAPIResults はジャンルマスタAPIのレスポンスデータを表す。
type GenreAPIResults struct {
	Results struct {
		CommonResults
		Genre []MasterData `json:"genre"`
	}
}
