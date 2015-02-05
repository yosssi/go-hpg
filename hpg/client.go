package hpg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	shopAPIPath = "shop"
)

// APIバージョン
const (
	V1 = "v1"
)

// テストコードにて書き換えられるようにするため、定数ではなく変数として定義する。
var baseURL = "http://webservice.recruit.co.jp/hotpepper"

// テストコードで書き換えられるようにするため、変数にて定義する。
var readAll = ioutil.ReadAll

// ShopAPI は店名サーチAPIをコールし、その結果を返却する。
func ShopAPI(version string, params *ShopAPIParams) (*ShopAPIResults, error) {
	url := baseURL + "/" + shopAPIPath + "/" + version + "/" + params.Query()

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTPステータスコードが不正です。(期待：%d、実際：%d)", http.StatusOK, res.StatusCode)
	}

	b, err := readAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	r := new(ShopAPIResults)

	if err := json.Unmarshal(b, r); err != nil {
		return nil, err
	}

	return r, nil
}
