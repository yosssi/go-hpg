package hpg

// LargeServiceAreaAPIResults は大サービスエリアマスタAPIのレスポンスデータを表す。
type LargeServiceAreaAPIResults struct {
	Results struct {
		CommonResults
		LargeServiceArea []LargeServiceArea `json:"large_service_area"`
	}
}
