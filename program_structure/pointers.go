package program_structure

import "fmt"

func Pointers() {

	// 变量是包含值的存储区，用变量名称标识这个存储区,通过变量名可以访问被标识的存储区。
	var x int
	println(x) // 0 ，变量没有被初始化"zero"
	x = 123    // 改变x标识的存储区的值
	fmt.Printf("&x value %v type %T \n", &x,&x)

	var y int
	y = 456
	x = y // x标识区的值更新为y存储区的值
	fmt.Printf("after x=y,&x value %v type %T \n", &x,&x) // &x没有改变

	// 表达式&x获取一个指向x的
	var p=&x
	fmt.Printf("p value %v type %T \n", p,p)
	fmt.Printf("*p value %v type %T \n", *p,*p)

}
