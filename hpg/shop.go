package hpg

// Shop は店舗を表す。
type Shop struct {
	ID       string
	Name     string
	NameKana string
	Address  string
	Genre    Genre
	URLs     URLs
	Desc     string
}
