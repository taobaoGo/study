package datatypego

import "fmt"
/**
## 整型
- 有符号整数类型：int8、int16、int32和int64
- 无符号整数类型：uint8、uint16、uint32和uint64
- 跟CPU平台相关的数据类型：有符号整数int和无符号整数uint
- 特殊的无符号整数类型：uintptr（同指针），在32位平台下4字节，在64位平台下8字节
 */
func IntegerFunction() {
	var apples int32 = 1
	var oranges int16 = 2
	// 不同的整型类型需要转换成相同的才可以进行“算数运算”和“比较运算”
	var compote= int(apples) + int(oranges)
	if (apples == int32(oranges)) {
		fmt.Println("apples and oranges are equal.")
	}

	fmt.Print(compote)
}