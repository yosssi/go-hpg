package hpg

import (
	"bytes"
	"testing"
)

func TestCommonParams_Query(t *testing.T) {
	p := &CommonParams{}

	want := "?key=&format=json"

	if s := p.Query(); s != want {
		t.Errorf("s => %q, want %q", s, want)
	}
}

func TestCommonParams_WriteStringTo(t *testing.T) {
	testCases := []struct {
		in  *CommonParams
		out string
	}{
		{
			in: &CommonParams{
				Key:   "testKey",
				Start: 1,
				Count: 2,
			},
			out: "?key=testKey&start=1&count=2&format=json",
		},
		{
			in: &CommonParams{
				Key:      "testKey",
				Start:    1,
				Count:    2,
				Callback: "testCallback",
			},
			out: "?key=testKey&start=1&count=2&format=jsonp&callback=testCallback",
		},
	}

	for _, testCase := range testCases {
		bf := new(bytes.Buffer)

		testCase.in.WriteStringTo(bf)

		if s := bf.String(); s != testCase.out {
			t.Errorf("s => %q, want %q", s, testCase.out)
		}
	}
}

func Test_appendQuery(t *testing.T) {
	bf := new(bytes.Buffer)

	appendQuery(bf, "testKey", "test1,test2")

	want := "&testKey=test1%2Ctest2"

	if s := bf.String(); s != want {
		t.Errorf("s => %q, want %q", s, want)
	}
}
