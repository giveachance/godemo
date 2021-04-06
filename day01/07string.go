package main

import (
	"fmt"
	"strings"
)

func testString() {

	str := "E:\\GoProject\\src"
	fmt.Println(str)
	str1 := `E:\GoProject\src`
	fmt.Println(str1)
	str2 := `E:\GoProject\src
	E:\GoProject\src\github.com
	`
	fmt.Println(str2)

	fmt.Println(strings.Split(str1, "\\"))
	str3:= "abcdeb"
	fmt.Println(strings.HasPrefix(str3,"a"))
	fmt.Println(strings.Index(str3,"b"))
	fmt.Println(strings.LastIndex(str3,"b"))
	fmt.Println(strings.IndexAny(str3,"b"))
	fmt.Println(strings.Contains(str3,"b"))
	
}
