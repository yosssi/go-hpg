package hpg

import "testing"

func TestMiddleAreaAPIParams_Query(t *testing.T) {
	p := &MiddleAreaAPIParams{
		MiddleArea: []string{"Y005", "Y006"},
		LargeArea:  []string{"Z011", "Z012"},
		Keyword:    "testKeyword",
	}

	want := "?key=&format=json&middle_area=Y005%2CY006&large_area=Z011%2CZ012&keyword=testKeyword"

	if s := p.Query(); s != want {
		t.Errorf("s => %q, want %q", s, want)
	}
}
