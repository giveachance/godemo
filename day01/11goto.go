package main

import "fmt"

//定义label，当满足条件时
//continue label ，跳出到label循环后continue
//break label ， 跳出到label循环后break
func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
