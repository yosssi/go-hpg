package hpg

import "testing"

func TestFoodAPIParams_Path(t *testing.T) {
	p := new(FoodAPIParams)

	want := "food"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestFoodAPIParams_String(t *testing.T) {
	p := new(FoodAPIParams)

	want := "?key=&format=json"

	if s := p.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
