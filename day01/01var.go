package main

import "fmt"

var (
	name string
	age  int
	ok   bool
)

func test() {
	name = "hqh"
	age = 25
	ok = true
	fmt.Printf("name : %s", name)
	fmt.Printf("age: %d", age)
	fmt.Println("ok:", ok)
}

func testVar() {
	//标准声明
	var a1 string = "aa"
	//推断声明
	var a2 = "bb"
	//短变量声明
	a3 := "cc"
	//错误声明  ，必须加var或者使用短变量声明
	//a4 = "dd"
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	//fmt.Println(a4)

}

// := 用于声明局部变量，只能用在函数内部
// 声明局部变量与全局变量重名时，局部变量失效
func testDot() {
	age = 1
	fmt.Println(age)
	age := "111"
	fmt.Println(age)
	age = "15"
	fmt.Println(age)
	//声明变量的另一种方式
	var age1 string = "555"
	fmt.Println(age1)

}
