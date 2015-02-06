package hpg

import "testing"

func TestLargeServiceAreaAPIParams_Path(t *testing.T) {
	p := new(LargeServiceAreaAPIParams)

	want := "large_service_area"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
