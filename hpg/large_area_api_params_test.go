package hpg

import "testing"

func TestLargeAreaAPIParams_Query(t *testing.T) {
	p := &LargeAreaAPIParams{
		LargeArea: []string{"Z011", "Z012"},
		Keyword:   "testKeyword",
	}

	want := "?key=&format=json&large_area=Z011%2CZ012&keyword=testKeyword"

	if s := p.Query(); s != want {
		t.Errorf("s => %q, want %q", s, want)
	}
}
