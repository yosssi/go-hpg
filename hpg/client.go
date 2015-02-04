package hpg

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL = "http://webservice.recruit.co.jp/hotpepper"

const (
	shopAPIPath = "shop"
)

// ShopAPI は店名サーチAPIをコールし、その結果を返却する。
func ShopAPI(version string, params *ShopAPIParams) (*ShopAPIResults, error) {
	url := baseURL + "/" + shopAPIPath + "/" + version + "/" + params.Query()

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))
	return nil, nil
}
