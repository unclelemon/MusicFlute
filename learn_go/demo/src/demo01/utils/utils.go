package main

// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)

func main(){
	zs := 100
	var isSingle bool 
	// 查看数据类型 和 占用字节大小
	// %v 原始输出 %T l类型输出 %d 整数
	fmt.Printf("zs 的类型%T，zs占用的字节数是%d\n",zs, unsafe.Sizeof(zs))
	fmt.Printf("isSingle 的类型%T，isSingle占用的字节数是%d, isSingle 的值是%v\n",isSingle, unsafe.Sizeof(isSingle),isSingle)

	// 
}
