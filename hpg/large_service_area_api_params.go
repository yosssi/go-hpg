package hpg

// LargeServiceAreaAPIParams は大サービスエリアマスタAPIの検索クエリパラメータを表す。
type LargeServiceAreaAPIParams struct {
	CommonParams
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *LargeServiceAreaAPIParams) Path() string {
	return "large_service_area"
}
