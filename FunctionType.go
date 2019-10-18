package main

import (
	"./models"
	"fmt"
)

/**
函数作为函数的入参
类似dotnet的"委托"和"匿名函数",只是概念上更轻量简洁。
可以占坑留空让功能消费者来实现
*/

// 定义函数类型
type delegate func(bag models.Bag)

// 函数的入参为函数类型
func exec(f delegate) {
	bag := models.Bag{Items: []int{1, 2, 3}, Name: "Mico Bag", Color: "Black"}
	// 占坑，逻辑由调用者编写
	f(bag)
}

func FunctionTypeMain() {
	// 实现函数类型
	var p = func(bag models.Bag) {
		fmt.Println("first", bag)
	}
	exec(p)

	// “匿名函数”作为参数直接传递给exec函数
	exec(func(bag models.Bag) {
		fmt.Println("second", bag)
	})
}
