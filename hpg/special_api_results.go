package hpg

// SpecialAPIResults は特集マスタAPIのレスポンスデータを表す。
type SpecialAPIResults struct {
	Results struct {
		CommonResults
		Special []struct {
			MasterData
			SpecialCategory MasterData `json:"special_category"`
		} `json:"special"`
	}
}
