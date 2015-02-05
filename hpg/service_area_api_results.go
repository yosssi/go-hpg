package hpg

// ServiceAreaAPIResults はサービスエリアマスタAPIのレスポンスデータを表す。
type ServiceAreaAPIResults struct {
	Results struct {
		CommonResults
		ServiceArea []ServiceArea `json:"service_area"`
	}
}
