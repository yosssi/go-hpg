package hpg

import (
	"bytes"
	"strings"
)

// LargeAreaAPIParams は大エリアマスタAPIの検索クエリパラメータを表す。
type LargeAreaAPIParams struct {
	CommonParams
	LargeArea []string
	Keyword   string
}

// Query は、クエリパラメータ文字列を生成し、それを返却する。
func (p *LargeAreaAPIParams) Query() string {
	bf := new(bytes.Buffer)

	p.CommonParams.WriteStringTo(bf)

	if p.LargeArea != nil {
		appendQuery(bf, "large_area", strings.Join(p.LargeArea, ","))
	}

	if p.Keyword != "" {
		appendQuery(bf, "keyword", p.Keyword)
	}

	return bf.String()
}
