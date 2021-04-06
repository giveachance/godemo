package main

import "fmt"
//float 默认为float64
func testFloat() {
	a1 := 1.11
	fmt.Printf("%f\n",a1)
	//默认float64
	fmt.Printf("%T\n",a1)
	a2:= float32(1.22)
	fmt.Printf("%f\n",a2)
	fmt.Printf("%T\n",a2)
	a1 = float64(a2)
	a2 = float32(a1)
	fmt.Printf("%f\n",a1)
	//默认float64
	fmt.Printf("%T\n",a1)
	fmt.Printf("%f\n",a2)
	fmt.Printf("%T\n",a2)
}