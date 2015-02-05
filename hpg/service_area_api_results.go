package hpg

// ServiceAreaAPIResults はサービスエリアマスタAPIのレスポンスデータを表す。
type ServiceAreaAPIResults struct {
	Results struct {
		CommonResults
		ServiceArea []MasterData `json:"service_area"`
	}
}
