package hpg

import (
	"bytes"
	"net/url"
	"strconv"
)

// CommonParams は全APIに共通する検索クエリパラメータを表す。
type CommonParams struct {
	Key      string
	Start    int
	Count    int
	Callback string
}

// WriteStringTo は引数のBufferへパラメータを出力する。
func (p *CommonParams) WriteStringTo(bf *bytes.Buffer) {

	bf.WriteString("?key=" + url.QueryEscape(p.Key))

	if p.Start != 0 {
		appendQuery(bf, "start", strconv.Itoa(p.Start))
	}

	if p.Count != 0 {
		appendQuery(bf, "count", strconv.Itoa(p.Count))
	}

	if p.Callback == "" {
		appendQuery(bf, "format", "json")
	} else {
		appendQuery(bf, "format", "jsonp")
		appendQuery(bf, "callback", p.Callback)
	}
}

// AppendQuery は引数のBufferにkey-valueのクエリを追加する。
func appendQuery(bf *bytes.Buffer, key string, value string) {
	bf.WriteString("&" + key + "=" + url.QueryEscape(value))
}
