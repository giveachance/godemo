package main

import "fmt"
//进制转换
func testInt() {
	var a1 int = 100
	//十进制
	fmt.Printf("%d\n",a1)
	//二进制
	fmt.Printf("%b\n",a1)
	//八进制
	fmt.Printf("%o\n",a1)
	//
	fmt.Printf("%x\n",a1)

}

//符号转换
func testUint(){
	var a1  = 255
	fmt.Printf("%d\n",a1)
	fmt.Printf("%T\n",a1)
	a2 := int8(127)
	fmt.Println()
	fmt.Printf("%d\n",a2)
	fmt.Printf("%T\n",a2)
}