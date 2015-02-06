package hpg

import "testing"

func TestGenreAPIParams_Path(t *testing.T) {
	p := new(GenreAPIParams)

	want := "genre"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestGenreAPIParams_String(t *testing.T) {
	p := new(GenreAPIParams)

	want := "?key=&format=json"

	if s := p.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
