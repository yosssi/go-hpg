package hpg

// Shop は店舗を表す。
type Shop struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	NameKana string `json:"name_kana"`
	Address  string `json:"address"`
	Genre    Genre  `json:"genre"`
	URLs     URLs   `json:"urls"`
	Desc     string `json:"desc"`
}
