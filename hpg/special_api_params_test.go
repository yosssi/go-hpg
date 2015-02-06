package hpg

import "testing"

func TestSpecialAPIParams_Path(t *testing.T) {
	p := new(SpecialAPIParams)

	want := "special"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestSpecialAPIParams_String(t *testing.T) {
	p := new(SpecialAPIParams)

	want := "?key=&format=json"

	if s := p.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
