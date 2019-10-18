/**
类型、接收器Receiver、方法
接收器必须在包内用type定义，对包外的任何类型都不能做接收器，也不能有类型方法。
*/

// https://blog.csdn.net/kang___xi/article/details/86550705
// https://blog.csdn.net/qq_36874881/article/details/78065755
// https://www.jianshu.com/p/a537ee63d606

// 面向对象的语言为类提供了成员变量和成员方法，但是Go没有类。
// GO的方法(method)是用于特定类型的变量的函数，这种特定类型变量叫“接收器(Receiver)”,拥有方法的类型叫接收器。
package function

import (
	"fmt"
)
import "study/models"

func init() {
	fmt.Println("FunReceiver:1")
}

// 接收器(类型别名)必须在包内命名，并且只能在包内起作用
// bag可以做函数接收器，models.Bag无法做接收器
type bag models.Bag

/**
“指针类型”接收器
调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的。
*/
// 函数：面向过程编程
func insertPointer(b *models.Bag, itemid int) {
	b.Items = append(b.Items, itemid)
}

// 方法/行为：面向对象编程
// receiver：接收器
//           类型：接收器的类型可以是任何类型，任何类型都可以拥有方法，但必须是通过type定义的类型。
//           原型：func (接收器变量 接收器类型) 方法名(参数列表) (返回参数){ }
func (b *bag) InsertMethodPointer(itemid int) {
	b.Items = append(b.Items, itemid)
}

/**
非指针类型的接收器
    当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份。在非指针接收器的方法中可以获取接收器的成员值，但修改后无效。
*/
func insert(b models.Bag, itemid int) {
	b.Items = append(b.Items, itemid)
	fmt.Printf("%+v\n", b)
}
func (b bag) InsertMethod(itemid int) {
	b.Items = append(b.Items, itemid)
	fmt.Printf("%+v\n", b)
}

func FunReceiver() {
	// 声明定义变量b为Bag实例。
	// new:内置函数 func new(Type) *Type
	//     为入参“类型”分配空间，返回为类型新分配的空间的零值指针。

	// make:内置函数 func make(Type,size IntegerType) Type
	//      第一个参数类型，第二个参数为长度，返回值是一个类型(new返回指针)
	//      内建函数make分配并且初始化一个slice, 或map或chan对象。 并且只能是这三种对象。和new一样，第一个参数是类型，不是一个值。
	//      但是make的返回值就是这个类型（即使一个引用类型），而不是指针。 具体的返回值，依赖具体传入的类型。

	bagPointerModel := new(models.Bag)
	insertPointer(bagPointerModel, 100)
	insertPointer(bagPointerModel, 101)
	insertPointer(bagPointerModel, 102)
	fmt.Printf("%+v\n", bagPointerModel)
	bagPointer := new(bag)
	bagPointer.InsertMethodPointer(200)
	bagPointer.InsertMethodPointer(201)
	bagPointer.InsertMethodPointer(202)
	fmt.Printf("%+v\n", bagPointer)

	var bagModel models.Bag
	insert(bagModel, 500)
	insert(bagModel, 501)
	insert(bagModel, 502)
	fmt.Printf("%+v\n", bagModel)
	var bagObj bag
	bagObj.InsertMethod(600)
	bagObj.InsertMethod(601)
	bagObj.InsertMethod(602)
	fmt.Printf("%+v\n", bagObj)

}

/**
goland
文件名重构：shift+f6
变量重构  : shift+f6
*/
