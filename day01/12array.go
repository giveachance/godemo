package main

import "fmt"

//数组初始化
func testArrayInit() {
	//方法一
	var a [5]int
	for _, value := range a {
		fmt.Println(value)
	}
	fmt.Println()
	//方法二

	var a1 = [5]int{1, 2, 3, 4, 5}
	for _, value := range a1 {
		fmt.Println(value)
	}
	var a2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(a2)
	//方法三

	a3 := [...]int{1: 1, 3: 5}
	fmt.Println(a3)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a3) //type of a:[4]int
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}
}

//二维数组
func test2Array() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}
