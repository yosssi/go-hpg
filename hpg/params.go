package hpg

import "fmt"

// Params はパラメータを表すインターフェースである。
type Params interface {
	Path() string
	fmt.Stringer
}
