package hpg

// ServiceAreaAPIParams はサービスエリアマスタAPIの検索クエリパラメータを表す。
type ServiceAreaAPIParams struct {
	CommonParams
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *ServiceAreaAPIParams) Path() string {
	return "service_area"
}
