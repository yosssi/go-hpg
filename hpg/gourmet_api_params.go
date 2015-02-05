package hpg

// GourmetAPIParams はグルメサーチAPIの検索クエリパラメータを表す。
type GourmetAPIParams struct {
	CommonParams
	ID                []string
	Name              string
	NameKana          string
	NameAny           string
	TEL               string
	Address           string
	Special           []string
	SpecialOR         []string
	SpecialCategory   []string
	SpecialCategoryOR []string
	IsOpenTime        string
	IsOpenWeek        string
	LargeServiceArea  string
	ServiceArea       []string
	LargeArea         []string
	MiddleArea        []string
	SmallArea         []string
	Keyword           []string
	Lat               *float64
	Lng               *float64
	Range             string
	Datum             string
	KtaiCoupon        string
	Genre             []string
	Food              []string
	Budget            []string
	PartyCapacity     int
	Wifi              bool
	Wedding           bool
	Course            bool
	FreeDrink         bool
	FreeFood          bool
	PrivateRoom       bool
	Horigotatsu       bool
	Tatami            bool
	Cocktail          bool
	Shochu            bool
	Sake              bool
	Wine              bool
	Card              bool
	NonSmoking        bool
	Charter           bool
	Ktai              bool
	Parking           bool
	BarrierFree       bool
	Sommelier         bool
	NightView         bool
	OpenAir           bool
	Show              bool
	Equipment         bool
	Karaoke           bool
	Band              bool
	TV                bool
	Lunch             bool
	Midnight          bool
	MidnightMeal      bool
	English           bool
	Pet               bool
	Child             bool
	CreditCard        []string
	Type              string
	Order             string
}

// Path は、このパラメータに紐尽くAPIパスを返却する。
func (p *GourmetAPIParams) Path() string {
	return "gourmet"
}

// String は、クエリパラメータ文字列を生成し、それを返却する。
func (p *GourmetAPIParams) String() string {
	bf := p.queryBuffer()

	bf.appendStringSlice("id", p.ID)
	bf.appendString("name", p.Name)
	bf.appendString("name_kana", p.NameKana)
	bf.appendString("name_any", p.NameAny)
	bf.appendString("tel", p.TEL)
	bf.appendString("address", p.Address)
	bf.appendStringSlice("special", p.Special)
	bf.appendStringSlice("special_or", p.SpecialOR)
	bf.appendStringSlice("special_category", p.SpecialCategory)
	bf.appendStringSlice("special_category_or", p.SpecialCategory)
	bf.appendString("is_open_time", p.IsOpenTime)
	bf.appendString("is_open_week", p.IsOpenWeek)
	bf.appendString("large_service_area", p.LargeServiceArea)
	bf.appendStringSlice("service_area", p.ServiceArea)
	bf.appendStringSlice("large_area", p.LargeArea)
	bf.appendStringSlice("middle_area", p.MiddleArea)
	bf.appendStringSlice("small_area", p.SmallArea)
	bf.appendStringSlice("keyword", p.Keyword)
	bf.appendFloatPtr("lat", p.Lat)
	bf.appendFloatPtr("lng", p.Lng)
	bf.appendString("range", p.Range)
	bf.appendString("datum", p.Datum)
	bf.appendString("ktai_coupon", p.KtaiCoupon)
	bf.appendStringSlice("genre", p.Genre)
	bf.appendStringSlice("food", p.Food)
	bf.appendStringSlice("budget", p.Budget)
	bf.appendInt("party_capacity", p.PartyCapacity)
	bf.appendBool("wifi", p.Wifi)
	bf.appendBool("wedding", p.Wedding)
	bf.appendBool("course", p.Course)
	bf.appendBool("free_drink", p.FreeDrink)
	bf.appendBool("free_food", p.FreeFood)
	bf.appendBool("private_room", p.PrivateRoom)
	bf.appendBool("horigotatsu", p.Horigotatsu)
	bf.appendBool("tatami", p.Tatami)
	bf.appendBool("cocktail", p.Cocktail)
	bf.appendBool("shochu", p.Shochu)
	bf.appendBool("sake", p.Sake)
	bf.appendBool("wine", p.Wine)
	bf.appendBool("card", p.Card)
	bf.appendBool("non_smoking", p.NonSmoking)
	bf.appendBool("charter", p.Charter)
	bf.appendBool("ktai", p.Ktai)
	bf.appendBool("parking", p.Parking)
	bf.appendBool("barrier_free", p.BarrierFree)
	bf.appendBool("sommelier", p.Sommelier)
	bf.appendBool("night_view", p.NightView)
	bf.appendBool("open_air", p.OpenAir)
	bf.appendBool("show", p.Show)
	bf.appendBool("equipment", p.Equipment)
	bf.appendBool("karaoke", p.Karaoke)
	bf.appendBool("band", p.Band)
	bf.appendBool("tv", p.TV)
	bf.appendBool("lunch", p.Lunch)
	bf.appendBool("midnight", p.Midnight)
	bf.appendBool("midnight_meal", p.MidnightMeal)
	bf.appendBool("english", p.English)
	bf.appendBool("pet", p.Pet)
	bf.appendBool("child", p.Child)
	bf.appendStringSlice("credit_card", p.CreditCard)
	bf.appendString("type", p.Type)
	bf.appendString("order", p.Order)

	return bf.String()
}
