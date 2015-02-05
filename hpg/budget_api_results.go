package hpg

// BudgetAPIResults は検索用予算マスタAPIのレスポンスデータを表す。
type BudgetAPIResults struct {
	Results struct {
		CommonResults
		Budget []Budget `json:budget`
	}
}
