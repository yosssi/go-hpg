package hpg

import (
	"errors"
	"io"
	"testing"
)

var errTest = errors.New("test error")

func TestCallAPI_GetErr(t *testing.T) {
	defer func(baseURLBack string) {
		baseURL = baseURLBack
	}(baseURL)

	baseURL = ""

	p := new(GourmetAPIParams)

	r := new(GourmetAPIResults)

	if err := CallAPI(V1, p, r); err == nil {
		notNilErrorExpected(t)
		return
	}
}

func TestCallAPI_invalidStatusCode(t *testing.T) {
	p := new(GourmetAPIParams)

	r := new(GourmetAPIResults)

	if err := CallAPI("invalidVersion", p, r); err == nil {
		notNilErrorExpected(t)
		return
	}
}

func TestCallAPI_readAllErr(t *testing.T) {
	defer func(readAllBak func(io.Reader) ([]byte, error)) {
		readAll = readAllBak
	}(readAll)

	readAll = func(_ io.Reader) ([]byte, error) {
		return nil, errTest
	}

	p := new(GourmetAPIParams)

	r := new(GourmetAPIResults)

	if err := CallAPI(V1, p, r); err != errTest {
		invalidError(t, err, errTest)
		return
	}
}

func TestCallAPI_jsonUnmarshalErr(t *testing.T) {
	defer func(readAllBak func(io.Reader) ([]byte, error)) {
		readAll = readAllBak
	}(readAll)

	readAll = func(_ io.Reader) ([]byte, error) {
		return []byte("test"), nil
	}

	p := new(GourmetAPIParams)

	r := new(GourmetAPIResults)

	if err := CallAPI(V1, p, r); err == nil {
		notNilErrorExpected(t)
		return
	}
}

func TestCallAPI(t *testing.T) {
	p := new(GourmetAPIParams)

	r := new(GourmetAPIResults)

	if err := CallAPI(V1, p, r); err != nil {
		nilErrorExpected(t, err)
		return
	}
}

func notNilErrorExpected(t *testing.T) {
	t.Error("err => nil, want => not nil")
}

func nilErrorExpected(t *testing.T, err error) {
	t.Errorf("err => %q, want => nil", err)
}

func invalidError(t *testing.T, err, want error) {
	if err == nil {
		t.Errorf("err => nil, want => %q", want)
	} else {
		t.Errorf("err => %q, want => %q", err, want)
	}
}
