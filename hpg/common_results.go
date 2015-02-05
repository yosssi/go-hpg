package hpg

// CommonResults は全APIに共通するレスポンスデータを表す。
type CommonResults struct {
	APIVersion       string  `json:"api_version"`
	ResultsAvailable int     `json:"results_available"`
	ResultsReturned  string  `json:"results_returned"`
	ResultsStart     int     `json:"results_start"`
	Error            []Error `json:"error"`
}
