package hpg

import "testing"

func TestServiceAreaAPIParams_Path(t *testing.T) {
	p := new(ServiceAreaAPIParams)

	want := "service_area"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
