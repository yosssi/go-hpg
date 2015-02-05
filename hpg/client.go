package hpg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	shopAPIPath   = "shop"
	budgetAPIPath = "budget"
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
	r := new(ShopAPIResults)

	if err := callAPI(shopAPIPath, version, params.Query(), r); err != nil {
		return nil, err
	}

	return r, nil
}

// BudgetAPI は検索用予算マスタAPIをコールし、その結果を返却する。
func BudgetAPI(version string, params *BudgetAPIParams) (*BudgetAPIResults, error) {
	r := new(BudgetAPIResults)

	if err := callAPI(budgetAPIPath, version, params.Query(), r); err != nil {
		return nil, err
	}

	return r, nil
}

// callAPI はホットペッパー WebサービスのAPIをコールする。
func callAPI(apiPath, version, query string, r interface{}) error {
	url := baseURL + "/" + apiPath + "/" + version + "/" + query

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTPステータスコードが不正です。(期待：%d、実際：%d)", http.StatusOK, res.StatusCode)
	}

	b, err := readAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, r); err != nil {
		return err
	}

	return nil
}
