package hpg

// ServiceArea はサービスエリアマスタデータを表す。
type ServiceArea struct {
	Code             string           `json:"code"`
	Name             string           `json:"name"`
	LargeServiceArea LargeServiceArea `json:"large_service_area"`
}
