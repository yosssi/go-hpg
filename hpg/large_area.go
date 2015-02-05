package hpg

// LargeArea は大エリアマスタデータを表す。
type LargeArea struct {
	Code             string           `json:"code"`
	Name             string           `json:"name"`
	ServiceArea      ServiceArea      `json:"service_area"`
	LargeServiceArea LargeServiceArea `json:"large_service_area"`
}
