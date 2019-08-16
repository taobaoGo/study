package function

import "fmt"

func init()  {
	fmt.Println("FuncVariable:1")
}

/** *******
函数变量
无论是“普通函数”还是“结构体的方法”，只要它们的签名一致，与它们签名一致的函数变量就可以保存“普通函数”或“结构体方法”。
var delegate func(int)
*/
type class struct {
}
func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}
func funcDo(v int) {
	fmt.Println("call function do:", v)
}
func FuncVar()  {
	var delegate func(int)
	c := new(class)
	// 保存结构体方法
	delegate = c.Do
	delegate(100)
	// 保存普通函数
	delegate = funcDo
	delegate(100)
}

