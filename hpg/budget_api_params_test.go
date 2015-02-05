package hpg

import "testing"

func TestBudgetAPIParams_Query(t *testing.T) {
	p := &BudgetAPIParams{}

	want := "?key=&format=json"

	if s := p.Query(); s != want {
		t.Errorf("s => %q, want %q", s, want)
	}
}
