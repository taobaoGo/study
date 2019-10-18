package typego

import "fmt"
import "../models"

func Struct_fun() {
	bag := models.Bag{Color: "red", Name: "IBM Bag", Items: []int{1, 2}}
	fmt.Printf("%+v\n", bag)
}
