package main

import "fmt"

//切片本质上是对数组的引用
func testSlice() {

	var a []string //声明一个字符串切片

	a1 := [3]string{"1", "2", "3"}

	a = a1[:]
	fmt.Println(a)
	fmt.Println(a1)
	a1[2] = "2222"
	fmt.Println(a)
	fmt.Println(a1)

	var a2 []string = []string{"aa"}
	a3 := []string{"aa"}
	fmt.Println(a2)
	fmt.Println(a3)

}

func testMake() {

	a:=make([]int,1,2)
}
