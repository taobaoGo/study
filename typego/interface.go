package typego

import "fmt"
import "math"

/**
interface{}的动态特性很不能适应复杂的数据结构，
https://www.jianshu.com/p/16e8ebd289ce
 */

// 图形
type Shaper interface {
	// 面积
	Area() float32
	// 周长
	Circumference() float32
}

// 矩形
type Rect struct {
	Width float32
	Height float32
}
// 矩形实现接口面积方法
func (r Rect) Area() float32 {
	return r.Width * r.Height
}
// 矩形实现周长接口
func (r Rect) Circumference() float32 {
	return 2 * (r.Width + r.Height)
}

// 圆形
type Circle struct {
	Radius float32
}
// 圆形实现接口的面积方法
func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}
// 圆形实现接口的周长方法
func (c Circle) Circumference() float32 {
	return math.Pi * 2 * c.Radius
}

func ShaperPrint(shaper Shaper){

}

func InterfaceMain() {

	r := Rect{
		Height: 10,
		Width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())

	c := Circle{
		Radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area())
}

// 由于x.(T)只能是接口类型判断，所以传参时候，传入的是接口类型
// 为何test的类型可以是一个空接口？埋伏笔下文便知。
func CheckShaper(test interface{}) {
	if _, ok := test.(Shaper); ok {
		fmt.Printf("Student implements People")
	}
}


//func main() {
//	cbs := Student{Name:"咖啡色的羊驼"}
//	CheckShaper(cbs) // Student implements People
//}
