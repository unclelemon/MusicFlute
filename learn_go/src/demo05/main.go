package main

import "fmt"
import "math"

func main()  {
	
	// 可以在if中赋值一个新的变量，但是不可以用var声明一个新的变量
	// 不要在if后面加（）,除非有判断
	// {} 必须有 else 必须在{ 同一行
	if age := 17; age > 29 {
		fmt.Println("你已经30了")
	} else if age < 18 {
		fmt.Println("你还没到18")
	}
	fmt.Println(math.Sqrt(4))
}