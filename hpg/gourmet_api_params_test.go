package hpg

import "testing"

func TestGourmetAPIParams_Path(t *testing.T) {
	p := new(GourmetAPIParams)

	want := "gourmet"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestGourmetAPIParams_String(t *testing.T) {
	p := new(GourmetAPIParams)

	want := "?key=&format=json"

	if s := p.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
