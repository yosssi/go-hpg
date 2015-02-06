package hpg

import "testing"

func TestFoodCategoryAPIParams_Path(t *testing.T) {
	p := new(FoodCategoryAPIParams)

	want := "food_category"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
