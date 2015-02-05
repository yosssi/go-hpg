package hpg

// ShopAPIResults は店名サーチAPIのレスポンスデータを表す。
type ShopAPIResults struct {
	Results struct {
		CommonResults
		Shop []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			NameKana string `json:"name_kana"`
			Address  string `json:"address"`
			Genre    struct {
				Name string `json:"name"`
			} `json:"genre"`
			URLs struct {
				PC     string `json:"pc"`
				Mobile string `json:"mobile"`
				QR     string `json:"qr"`
			} `json:"urls"`
			Desc string `json:"desc"`
		} `json:"shop"`
	}
}
