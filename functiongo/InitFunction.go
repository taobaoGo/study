package function

import "fmt"

/**
main包的函数在main函数调用之前被调用
其他包的init函数在被引用时调用
*/
func init() {
	fmt.Println("InitFunction:1")
}

func init() {
	fmt.Println("InitFunction:2")
}
