package hpg

import "testing"

func TestLargeAreaAPIParams_Path(t *testing.T) {
	p := new(LargeAreaAPIParams)

	want := "large_area"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestLargeAreaAPIParams_String(t *testing.T) {
	p := new(LargeAreaAPIParams)

	want := "?key=&format=json"

	if s := p.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
