package hpg

import (
	"bytes"
	"net/url"
	"strconv"
	"strings"
)

// queryBuffer はURLクエリパラメータのバイトバッファを表す。
type queryBuffer struct {
	bytes.Buffer
}

// append はkey-valueのクエリを追加する。
func (bf *queryBuffer) append(k string, v string) {
	bf.WriteString("&" + k + "=" + url.QueryEscape(v))
}

// appendString はkey-valueのクエリを追加する。
func (bf *queryBuffer) appendString(k string, v string) {
	if v == "" {
		return
	}

	bf.append(k, v)
}

// appendStringSlice はkey-valueのクエリを追加する。
func (bf *queryBuffer) appendStringSlice(k string, v []string) {
	if v == nil {
		return
	}

	bf.append(k, strings.Join(v, ","))
}

// appendInt はkey-valueのクエリを追加する。
func (bf *queryBuffer) appendInt(k string, v int) {
	if v == 0 {
		return
	}

	bf.append(k, strconv.Itoa(v))
}

// appendFloatPtr はkey-valueのクエリを追加する。
func (bf *queryBuffer) appendFloatPtr(k string, v *float64) {
	if v == nil {
		return
	}

	bf.append(k, strconv.FormatFloat(*v, 'f', 10, 64))
}

// appendBool はkey-valueのクエリを追加する。
func (bf *queryBuffer) appendBool(k string, v bool) {
	if !v {
		return
	}

	bf.append(k, "1")
}
