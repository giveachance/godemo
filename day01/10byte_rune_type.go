package main

import "fmt"

//字符串为不可变对象
//当字符存储为asii码时，以byte类型存储(uint8)
//当字符存储为非asii码时，以rune类型存储（int32)
func testByte() {

	str := "红色的萝卜"
	str1 := []rune(str)
	str1[0] = '白'
	fmt.Printf("%v\n", str1)
	fmt.Printf("%v\n", string(str1))

	c := '白'
	fmt.Printf("%T\n", c)
	d := rune(c)
	fmt.Printf("%T\n", d) //int32
	fmt.Printf("%c\n", d) //白
	fmt.Printf("%v\n", d) //3033
	fmt.Println()

	e := 'a'
	f := byte(e) //uint8
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}

//类型转换 T()
//T为想要转换的类型
func testType() {
	a := 11
	b := string(11)
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%s\n", b)
}
