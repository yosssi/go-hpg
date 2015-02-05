package hpg

// LargeAreaAPIResults は大エリアマスタAPIのレスポンスデータを表す。
type LargeAreaAPIResults struct {
	Results struct {
		CommonResults
		LargeArea []LargeArea `json:"large_area"`
	}
}
