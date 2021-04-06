package main

import "fmt"

//在对普通变量使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用*操作，也就是指针取值
func testPoint() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

func testNew() {

	var a *int //nil
	//*a = 100 //空指针错误
	fmt.Println(*a)

	a1 := new(int)

	*a1 = 100
	fmt.Println(*a1)
}

//二者都是用来做内存分配的。
//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
func testMake() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
