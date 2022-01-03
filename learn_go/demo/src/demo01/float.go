package main


import (
	"fmt"
	"unsafe"
)

func main(){
	
	// 默认float64 占8个字节
	zs := 100.0
	// 查看数据类型 和 占用字节大小
	fmt.Printf("zs 的类型%T，zs占用的字节数是%d\n",zs, unsafe.Sizeof(zs))

	// float32 占4个字节
	var f1 float32 = 100.0
	// 查看数据类型 和 占用字节大小
	fmt.Printf("f1 的类型%T，zs占用的字节数是%d\n",f1, unsafe.Sizeof(f1))

	// float64 占8个字节
	var f2 float64 = 100.0
	// 查看数据类型 和 占用字节大小
	fmt.Printf("f2 的类型%T，zs占用的字节数是%d\n",f2, unsafe.Sizeof(f2))

	// 有符号
	var f3 float32 = -100.98098934893845945
	var f4 float64 = -100.98098934893845945
	fmt.Println("f3 = ",f3)
	fmt.Println("f4 = ",f4)

	// 科学计数法
	f5 := 5.12e2
	fmt.Println("f5 = ",f5)
	// 科学计数法
	f6 := 5.12e-2
	fmt.Println("f6 = ",f6)
}
