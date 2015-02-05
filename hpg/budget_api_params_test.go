package hpg

import "testing"

func TestBudgetAPIParams_Path(t *testing.T) {
	p := new(BudgetAPIParams)

	want := "budget"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
