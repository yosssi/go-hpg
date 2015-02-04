package hpg

// ShopAPIResults は店名サーチAPIのレスポンスデータを表す。
type ShopAPIResults struct {
	Results struct {
		CommonResults
		Shop []Shop
	}
}
