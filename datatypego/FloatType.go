package datatypego

import (
	"fmt"
	"math"
)
/**
浮点数不是一种精确地表达方式，所以像整型那样直接用 == 来判断两个浮点数是否相等是不可行的
 */

const MIN = 0.000001
// MIN 为用户自定义的比较精度
func IsEqual(f1, f2 float64) bool {
	return math.Dim(f1, f2) < MIN
}

func FloatMain() {
	a := 0.0000123
	b := 0.000012234
	if IsEqual(a, b) {
		fmt.Println("a < b")
	}
}