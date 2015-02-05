package hpg

import "testing"

func TestCreditCardAPIParams_Path(t *testing.T) {
	p := new(CreditCardAPIParams)

	want := "credit_card"

	if s := p.Path(); s != want {
		t.Errorf("s => %q, want %q", s, want)
	}
}
