package hpg

// CommonResults は全APIに共通するレスポンスデータを表す。
type CommonResults struct {
	APIVersion       float64
	ResultsAvailable int
	ResultsReturned  int
	ResultsStart     int
	Error            Error
}
