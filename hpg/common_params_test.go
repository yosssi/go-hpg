package hpg

import "testing"

func TestCommonParams_queryBuffer_CallbackEmpty(t *testing.T) {
	p := new(CommonParams)

	bf := p.queryBuffer()

	want := "?key=&format=json"

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestCommonParams_queryBuffer(t *testing.T) {
	p := &CommonParams{
		Callback: "testCallback",
	}

	bf := p.queryBuffer()

	want := "?key=&format=jsonp&callback=testCallback"

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}

func TestCommonParams_String(t *testing.T) {
	p := new(CommonParams)

	want := "?key=&format=json"

	if s := p.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
		return
	}
}
