package hpg

import "testing"

func TestSpecialCategoryAPIParams_Path(t *testing.T) {
	p := new(SpecialCategoryAPIParams)

	want := "special_category"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestSpecialCategoryAPIParams_String(t *testing.T) {
	p := new(SpecialCategoryAPIParams)

	want := "?key=&format=json"

	if s := p.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
