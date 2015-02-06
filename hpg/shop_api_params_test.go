package hpg

import "testing"

func TestShopAPIParams_Path(t *testing.T) {
	p := new(ShopAPIParams)

	want := "shop"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestShopAPIParams(t *testing.T) {
	p := new(ShopAPIParams)

	want := "?key=&format=json"

	if s := p.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
