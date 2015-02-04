package hpg

import (
	"bytes"
	"net/url"
	"strings"
)

// ShopAPIParams は店名サーチAPIの検索クエリパラメータを表す。
type ShopAPIParams struct {
	CommonParams
	KeyWord []string
	Tel     string
}

// Query は、クエリパラメータ文字列を生成し、それを返却する。
func (p *ShopAPIParams) Query() string {
	bf := new(bytes.Buffer)

	p.CommonParams.WriteStringTo(bf)

	if p.KeyWord != nil {
		bf.WriteString("&keyword=" + strings.Join(p.KeyWord, ","))
	}

	if p.Tel != "" {
		bf.WriteString("&tel=" + p.Tel)
	}

	return url.QueryEscape(bf.String())
}
