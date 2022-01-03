package main

import (
	"demo01/model"
	"fmt"
)
 
func main() {
	var i int = 10
	fmt.Printf("i 的地址是%d\n", &i)
	fmt.Println("i 的地址是", &i)

	// ptr 是一个指针变量
	// ptr 的类型是*int
	var ptr *int = &i
	fmt.Println("ptr的值是", ptr)
	fmt.Println("ptr的值是", *ptr)
	
	fmt.Println(model.Test)
	// fmt.Println(model.test)
}
