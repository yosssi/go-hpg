package hpg

// BudgetAPIResults は検索用予算マスタAPIのレスポンスデータを表す。
type BudgetAPIResults struct {
	Results struct {
		CommonResults
		Budget []MasterData `json:"budget"`
	}
}
