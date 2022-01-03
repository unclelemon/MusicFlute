package main

import (
	"fmt"
)

func main() {
	var str string = "test"
	slice1 := []byte(str)

	fmt.Println("str:"+str)
	fmt.Println("slice1:",slice1)
	slice1[0] = 'a'

	fmt.Println("str:"+str)
	fmt.Println("slice1:",slice1)

	var array = [...]int{1,2,3}
	slice2 := array[0:2]
	slice2[0] = 5
	fmt.Println(array)
	fmt.Println(slice2)

	slice3 := make([]int, 3)
	fmt.Println(slice3)
	slice3[2] = 4
	fmt.Println(slice3)
	slice3 = append(slice3,2,3,4)
	fmt.Println(slice3)
}