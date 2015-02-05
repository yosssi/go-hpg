package hpg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	shopAPIPath             = "shop"
	budgetAPIPath           = "budget"
	largeServiceAreaAPIPath = "large_service_area"
	serviceAreaAPIPath      = "service_area"
	largeAreaAPIPath        = "large_area"
	middleAreaAPIPath       = "middle_area"
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

// LargeServiceAreaAPI は大サービスエリアマスタAPIをコールし、その結果を返却する。
func LargeServiceAreaAPI(version string, params *LargeServiceAreaAPIParams) (*LargeServiceAreaAPIResults, error) {
	r := new(LargeServiceAreaAPIResults)

	if err := callAPI(largeServiceAreaAPIPath, version, params.Query(), r); err != nil {
		return nil, err
	}

	return r, nil
}

// ServiceAreaAPI はサービスエリアマスタAPIをコールし、その結果を返却する。
func ServiceAreaAPI(version string, params *ServiceAreaAPIParams) (*ServiceAreaAPIResults, error) {
	r := new(ServiceAreaAPIResults)

	if err := callAPI(serviceAreaAPIPath, version, params.Query(), r); err != nil {
		return nil, err
	}

	return r, nil
}

// LargeAreaAPI は大エリアマスタAPIをコールし、その結果を返却する。
func LargeAreaAPI(version string, params *LargeAreaAPIParams) (*LargeAreaAPIResults, error) {
	r := new(LargeAreaAPIResults)

	if err := callAPI(largeAreaAPIPath, version, params.Query(), r); err != nil {
		return nil, err
	}

	return r, nil
}

// MiddleAreaAPI は中エリアマスタAPIをコールし、その結果を返却する。
func MiddleAreaAPI(version string, params *MiddleAreaAPIParams) (*MiddleAreaAPIResults, error) {
	r := new(MiddleAreaAPIResults)

	if err := callAPI(middleAreaAPIPath, version, params.Query(), r); err != nil {
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
