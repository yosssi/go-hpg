package hpg

import (
	"bytes"
	"strings"
)

// MiddleAreaAPIParams は大エリアマスタAPIの検索クエリパラメータを表す。
type MiddleAreaAPIParams struct {
	CommonParams
	MiddleArea []string
	LargeArea  []string
	Keyword    string
}

// Query は、クエリパラメータ文字列を生成し、それを返却する。
func (p *MiddleAreaAPIParams) Query() string {
	bf := new(bytes.Buffer)

	p.CommonParams.WriteStringTo(bf)

	if p.MiddleArea != nil {
		appendQuery(bf, "middle_area", strings.Join(p.MiddleArea, ","))
	}

	if p.LargeArea != nil {
		appendQuery(bf, "large_area", strings.Join(p.LargeArea, ","))
	}

	if p.Keyword != "" {
		appendQuery(bf, "keyword", p.Keyword)
	}

	return bf.String()
}
