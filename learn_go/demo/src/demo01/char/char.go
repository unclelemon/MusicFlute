package char

// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)



func main() {
	c1 := 'a'
	var c2 byte = '0'
	fmt.Println("c1 = ", c1)
	fmt.Println("c2 = ", c2)

	// 输出格式化字符
	fmt.Printf("c1 = %c ", c1)
	fmt.Printf("c2 = %c ", c2)

	// golang 中字符都是使用 utf-8编码
	// 英文字母1个字节， 汉字3个字节
	// 如果使用的字符码值超过了255，使用int代替
	var c3 int = '林'
	fmt.Printf("c3 = %c，%d ", c3, c3)

	// 默认转成int32了
	var c4 = 10 + 'a' // 10 + 97 = 107
	fmt.Println("c4 = ", c4)
	fmt.Printf("c4 的类型%T，c4占用的字节数是%d，%c\n", c4, unsafe.Sizeof(c4), c4)
}
