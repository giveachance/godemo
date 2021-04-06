package main

import "fmt"

func testIf() {

	a := 1
	if a<0{
		fmt.Println("hh")
	}else if a > 0 &&a<1{
		fmt.Println("hehe")
	}else {
		fmt.Println("kk")
	}
}