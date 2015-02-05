package hpg

// MiddleArea は中エリアマスタデータを表す。
type MiddleArea struct {
	Code             string           `json:"code"`
	Name             string           `json:"name"`
	LargeArea        LargeArea        `json:"large_area"`
	ServiceArea      ServiceArea      `json:"service_area"`
	LargeServiceArea LargeServiceArea `json:"large_service_area"`
}
