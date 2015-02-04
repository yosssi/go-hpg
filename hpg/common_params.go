package hpg

import (
	"bytes"
	"strconv"
)

// CommonParams は全APIに共通する検索クエリパラメータを表す。
type CommonParams struct {
	Key      string
	Start    int
	Count    int
	Callback string
}

// WriteStringTo は引数のwへパラメータを出力する。
func (p *CommonParams) WriteStringTo(bf *bytes.Buffer) {

	bf.WriteString("?key=" + p.Key)

	if p.Start != 0 {
		bf.WriteString("&start=" + strconv.Itoa(p.Start))
	}

	if p.Count != 0 {
		bf.WriteString("&count=" + strconv.Itoa(p.Count))
	}

	if p.Callback == "" {
		bf.WriteString("&format=json")
	} else {
		bf.WriteString("&format=jsonp&callback=" + p.Callback)
	}
}
