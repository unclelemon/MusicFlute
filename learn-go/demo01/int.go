package main
import "fmt"

// 定义全局变量
var q1, q2, q3 = 100.200,300,400

var (
	q4 int
	q5 int
	q6 int
)

func main(){

	var test int = 100
	fmt.Println("i =", test)

	// 设置默认值
	var i int 
	fmt.Println("i =", i)

	// 根据值自动推导变量类型
	var num = 10.11
	fmt.Println("num =", num)

	// 省略var 使用:代替
	name := "name"
	fmt.Println("name =",name)

	// golang 一次性声明多个变量 1
	var n1, n2, n3 int
	fmt.Println("n1 =",n1, ",n2 =",n2,",n3 =",n3)

	// 其中23.0输出也为23 golang 一次性声明多个变量 1
	var n4, n5, n6 = 100, "tom", 23.0
	fmt.Println("n1 =",n4, ",n2 =",n5,",n3 =",n6)

	n7, n8, n9 := 100,"sdjk",888
	fmt.Println("n7 =",n7, ",n8 =",n8,",n9 =",n9)

	fmt.Println("n7 =",q1, ",n8 =",q2,",n9 =",q3)
 
	// 全局变量的赋值
	q4 = 10
	q5 = 20
	q6 = 30

	// uint表示正数
	// rune 与int32一样，表示一个Unicode码 -2^31 ~ 2^31-1
	// byte 等价uint8 0~255
	var zs rune 
	fmt.Println("zs =",zs)

	
}