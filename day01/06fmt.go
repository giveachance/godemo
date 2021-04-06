package main

import "fmt"

func testfmt() {
	var n = 100
	//十进制打印
	fmt.Printf("%d\n",n)
	//二进制打印
	fmt.Printf("%b\n",n)
	//八进制打印
	fmt.Printf("%o\n",n)
	//十六进制打印
	fmt.Printf("%x\n",n)
	//类型打印
	fmt.Printf("%T\n",n)
	//值打印
	fmt.Printf("%v\n",n)
	//类型和值打印
	fmt.Printf("%#v\n",n)
	//浮点数打印
	fmt.Printf("%f\n",float32(n))
	fmt.Println()
	var m = "hello"
	fmt.Printf("%s\n",m)
	fmt.Printf("%v\n",m)
	fmt.Printf("%#v\n",m)

}
