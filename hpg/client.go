package hpg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIバージョン
const (
	V1 = "v1"
)

// テストコードにて書き換えられるようにするため、定数ではなく変数として定義する。
var baseURL = "http://webservice.recruit.co.jp/hotpepper"

// テストコードにて書き換えられるようにするため、変数にて定義する。
var readAll = ioutil.ReadAll

// CallAPI はホットペッパー WebサービスのAPIをコールする。
func CallAPI(version string, p Params, r Results) error {
	url := baseURL + "/" + p.Path() + "/" + version + "/" + p.String()

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
