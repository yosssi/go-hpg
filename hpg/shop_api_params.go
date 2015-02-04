package hpg

import (
	"bytes"
	"strings"
)

// ShopAPIParams は店名サーチAPIの検索クエリパラメータを表す。
type ShopAPIParams struct {
	CommonParams
	Keyword []string
	Tel     string
}

// Query は、クエリパラメータ文字列を生成し、それを返却する。
func (p *ShopAPIParams) Query() string {
	bf := new(bytes.Buffer)

	p.CommonParams.WriteStringTo(bf)

	if p.Keyword != nil {
		appendQuery(bf, "keyword", strings.Join(p.Keyword, ","))
	}

	if p.Tel != "" {
		appendQuery(bf, "tel", p.Tel)
	}

	return bf.String()
}
