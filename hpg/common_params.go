package hpg

import "net/url"

// CommonParams は全APIに共通する検索クエリパラメータを表す。
type CommonParams struct {
	Key      string
	Start    int
	Count    int
	Callback string
}

// queryBuffer はqueryBufferを生成・初期化してそれを返却する。
func (p *CommonParams) queryBuffer() *queryBuffer {
	bf := new(queryBuffer)

	bf.WriteString("?key=" + url.QueryEscape(p.Key))

	bf.appendInt("start", p.Start)
	bf.appendInt("count", p.Count)

	if p.Callback == "" {
		bf.append("format", "json")
	} else {
		bf.append("format", "jsonp")
		bf.append("callback", p.Callback)
	}

	return bf
}

// String は、クエリパラメータ文字列を生成し、それを返却する。
func (p *CommonParams) String() string {
	return p.queryBuffer().String()
}
