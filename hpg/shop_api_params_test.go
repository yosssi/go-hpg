package hpg

import "testing"

func TestShopAPIParams_Query(t *testing.T) {
	p := &ShopAPIParams{
		Keyword: []string{"testKeyword1", "testKeyword2"},
		Tel:     "0",
	}

	want := "?key=&format=json&keyword=testKeyword1%2CtestKeyword2&tel=0"

	if s := p.Query(); s != want {
		t.Errorf("s => %q, want %q", s, want)
	}
}
