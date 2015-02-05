package hpg

import (
	"errors"
	"io"
	"os"
	"testing"
)

var key = os.Getenv("HPG_KEY")

var errTest = errors.New("テストエラー")

func TestShopAPI_GetErr(t *testing.T) {
	defer func(baseURLBacked string) {
		baseURL = baseURLBacked
	}(baseURL)

	baseURL = ""

	p := &ShopAPIParams{
		CommonParams: CommonParams{
			Key: key,
		},
		Keyword: []string{"テスト"},
	}

	if _, err := ShopAPI(V1, p); err == nil {
		notNilErrorExpected(t)
		return
	}
}

func TestShopAPI_invalidStatusCode(t *testing.T) {
	p := &ShopAPIParams{
		CommonParams: CommonParams{
			Key: key,
		},
		Keyword: []string{"テスト"},
	}

	if _, err := ShopAPI("invalidVersion", p); err == nil {
		notNilErrorExpected(t)
		return
	}
}

func TestShopAPI_readAllErr(t *testing.T) {
	defer func(readAllBacked func(io.Reader) ([]byte, error)) {
		readAll = readAllBacked
	}(readAll)

	p := &ShopAPIParams{
		CommonParams: CommonParams{
			Key: key,
		},
		Keyword: []string{"テスト"},
	}

	readAll = func(_ io.Reader) ([]byte, error) {
		return nil, errTest
	}

	if _, err := ShopAPI(V1, p); err == nil {
		notNilErrorExpected(t)
		return
	}
}

func TestShopAPI_UnmarshalErr(t *testing.T) {
	defer func(readAllBacked func(io.Reader) ([]byte, error)) {
		readAll = readAllBacked
	}(readAll)

	p := &ShopAPIParams{
		CommonParams: CommonParams{
			Key: key,
		},
		Keyword: []string{"テスト"},
	}

	readAll = func(_ io.Reader) ([]byte, error) {
		return []byte("test"), nil
	}

	if _, err := ShopAPI(V1, p); err == nil {
		notNilErrorExpected(t)
		return
	}
}

func TestShopAPI(t *testing.T) {
	p := &ShopAPIParams{
		CommonParams: CommonParams{
			Key: key,
		},
		Keyword: []string{"テスト"},
	}

	if _, err := ShopAPI(V1, p); err != nil {
		nilErrorExpected(t, err)
		return
	}
}

func notNilErrorExpected(t *testing.T) {
	t.Error("err => nil, want not nil")
}

func nilErrorExpected(t *testing.T, err error) {
	t.Errorf("err => %q, want nil", err)
}

func invalidError(t *testing.T, err, want error) {
	if err == nil {
		t.Errorf("err => nil, want %q", want)

	} else {
		t.Errorf("err => %q, want %q", err, want)

	}

}
