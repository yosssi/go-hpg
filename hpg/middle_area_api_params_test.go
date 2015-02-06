package hpg

import "testing"

func TestMiddleAreaAPIParams_Path(t *testing.T) {
	p := new(MiddleAreaAPIParams)

	want := "middle_area"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestMiddleAreaAPIParams_String(t *testing.T) {
	p := new(MiddleAreaAPIParams)

	want := "?key=&format=json"

	if s := p.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
