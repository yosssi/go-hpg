package hpg

// GourmetAPIResults はグルメサーチAPIのレスポンスデータを表す。
type GourmetAPIResults struct {
	Results struct {
		CommonResults
		Shop []struct {
			ID               string     `json:"id"`
			Name             string     `json:"name"`
			LogoImage        string     `json:"logo_image"`
			NameKane         string     `json:"name_kana"`
			Address          string     `json:"address"`
			StationName      string     `json:"station_name"`
			KtaiCoupon       string     `json:"ktai_coupon"`
			LargeServiceArea MasterData `json:"large_service_area"`
			ServiceArea      MasterData `json:"service_area"`
			LargeArea        MasterData `json:"large_area"`
			MiddleArea       MasterData `json:"middle_area"`
			SmallArea        MasterData `json:"small_area"`
			Lat              string     `json:"lat"`
			Lng              string     `json:"lng"`
			Genre            struct {
				MasterData
				Catch string `json:"catch"`
			} `json:"genre"`
			SubGenre MasterData `json:"sub_genre"`
			Food     MasterData `json:"food"`
			SubFood  MasterData `json:"sub_food"`
			Budget   struct {
				MasterData
				Average string
			} `json:"budget"`
			BudgetMemo   string `json:"budget_memo"`
			Catch        string `json:"catch"`
			Capacity     string `json:"capacity"`
			Access       string `json:"access"`
			MobileAccess string `json:"mobile_access"`
			URLs         struct {
				PC     string `json:"pc"`
				Mobile string `json:"mobile"`
				QR     string `json:"qr"`
			} `json:"urls"`
			Photo struct {
				PC struct {
					L string `json:"l"`
					M string `json:"m"`
					S string `json:"s"`
				} `json:"pc"`
				Mobile struct {
					L string `json:"l"`
					S string `json:"s"`
				} `json:"mobile"`
			} `json:"photo"`
			Open           string `json:"open"`
			Close          string `json:"close"`
			PartyCapacity  string `json:"party_capacity"`
			Wifi           string `json:"wifi"`
			Wedding        string `json:"wedding"`
			Course         string `json:"course"`
			FreeDrink      string `json:"free_drink"`
			FreeFood       string `json:"free_food"`
			PrivateRoom    string `json:"private_room"`
			Horigotatsu    string `json:"horigotatsu"`
			Tatami         string `json:"tatami"`
			Card           string `json:"card"`
			NonSmoking     string `json:"non_smoking"`
			Charter        string `json:"charter"`
			Ktai           string `json:"ktai"`
			Parking        string `json:"parking"`
			BarrierFree    string `json:"barrier_free"`
			OtherMemo      string `json:"other_memo"`
			Sommelier      string `json:"sommelier"`
			OpenAir        string `json:"open_air"`
			Show           string `json:"show"`
			Equipment      string `json:"equipment"`
			Karaoke        string `json:"karaoke"`
			Band           string `json:"band"`
			TV             string `json:"tv"`
			English        string `json:"english"`
			Pet            string `json:"pet"`
			Child          string `json:"child"`
			Lunch          string `json:"lunch"`
			Midnight       string `json:"midnight"`
			ShopDetailMemo string `json:"shop_detail_memo"`
			CouponURLs     struct {
				PC     string `json:"pc"`
				Mobile string `json:"mobile"`
				QR     string `json:"qr"`
				SP     string `json:"sq"`
			} `json:"coupon_urls"`
			Special struct {
				MasterData
				SpecialCategory MasterData `json:"special_category"`
				Title           string     `json:"title"`
			} `json:"special"`
			CreditCard MasterData `json:"credit_card"`
		} `json:"shop"`
	}
}
