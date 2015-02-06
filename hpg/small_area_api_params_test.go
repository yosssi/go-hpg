package hpg

import "testing"

func TestSmallAreaAPIParams_Path(t *testing.T) {
	p := new(SmallAreaAPIParams)

	want := "small_area"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestSmallAreaAPIParams_String(t *testing.T) {
	p := new(SmallAreaAPIParams)

	want := "?key=&format=json"

	if s := p.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
