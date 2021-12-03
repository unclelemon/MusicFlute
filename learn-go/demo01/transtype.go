package main

// import "fmt"
// import "unsafe"

import (
	"fmt"
)

func main(){
	// golang 数据只能显示转换 
	var i int = 100
	var j float32 = float32(i)

	fmt.Printf("i = %v, j = %v ",i,j)

	// 被转换的是变量存储的数值，变量本身并没有发生改变
	// 将转换后的值赋值给新变量
	var n1 int32 = 12
	var n2 int8
	// var n3 int8
	n2 = int8(n1) + 127  // [编译通过，但是结果不正确，按溢出处理]
	// n3 = int8(n1) + 128  // [编译不通过，128 已经超过了int8]
	fmt.Println(n2)

	var num1 int = 99
	var num2 float64 = 98.88998
	var str string 
	str = fmt.Sprintf("%d", num1)
	fmt.Println(str)
	str = fmt.Sprintf("%f%f", num2,num2)
	fmt.Println(str)
	

}
