package models

// 跨包小写属性无法访问
type Bag struct {
	Items []int
	Name string
	Color string
}
