package main

// import "fmt"
// import "unsafe"

import (
	"fmt"
)

func main(){
	zs := "beijing"
	fmt.Println(zs)

	// 双引号会识别转义字符
	var s1 string = "shanghai"
	fmt.Println(s1)

	// 反引号  不会识别转义字符
	var s2 string = `、df\df\ndfjkjk`
	fmt.Println(s2)

	// + 必须保留在上一行
	var s3 string ;
	s3 = "hello" + " world" +
	" ni"+" hao"
	fmt.Println(s3)
}
 