package main

import "fmt"

//for 初始语句;条件表达式;结束语句{
//    循环体语句
//}

func testFor() {
	//完整写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//省略初始语句，;不可省略
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	//省略初始语句和结束语句，在循环内部控制结束
	i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	a := "hello Go语言"
	//for range
	for i, v := range a {
		fmt.Printf("%d:%c\n", i, v)
	}
}
