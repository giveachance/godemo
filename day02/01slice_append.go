package main

import "fmt"

func appendTest() {
	//创建一个切片
	s1 := []string{"北京", "上海", "广州"}
	//必须接收
	s1 = append(s1, "深圳")
	fmt.Println(s1)
	s2 := []string{"北京1", "上海1", "广州1"}

	s1 = append(s1, s2...) //...表示拆开
}

//make([]T,a,b)，创建T类型的切片，长度为a(前面的数据为0或者nil或空串）,初始容量为b
//copy 将 原来的元素覆盖，且最多覆盖len(target)个元素，深拷贝
//append，在原来的长度上追加元素
func copyTest() {
	//创建一个切片
	//切片是引用类型
	s1 := []string{"北京", "上海", "广州"}
	//必须接收
	var s2 = make([]string, 4, 6)
	fmt.Println(s2)
	fmt.Println(len(s2), " ", cap(s2))
	copy(s2, s1)
	//s2 = append(s2,s1...)
	fmt.Println(s2)
	fmt.Println(len(s2), " ", cap(s2))

}

//切片不保存具体的值
//切片底层对应一个数组
//底层数组是一块连续的内存
//当append元素时，若原来的数组放不下时，go底层会把底层数组换一个
func appendPractice() {
	x1 := [...]byte{1, 3, 5, 7}     //创建数组
	s1 := x1[:]                     //创建切片，引用
	fmt.Println(s1)                 // [1 3 5 7]
	fmt.Printf("%p\n", s1)          //0xc0000101e0
	fmt.Printf("%p\n", &x1)         //0xc0000101e0
	s1 = append(s1[1:3], s1[3:]...) //长度超过了x1数组，更换底层数组，x1和s1二者没有了关联；当长度未超过时，二者仍有 关联
	fmt.Printf("%p\n", s1)          //0xc00000c300 ，地址改变了
	fmt.Printf("%p\n", &x1)         //0xc0000101e0
	fmt.Println(s1)                 //[3 5 5 7]
	fmt.Println(x1)                 //[1 3 5 7]
	s1[3] = 100
	fmt.Println(s1) //[3 5 5 100]
	fmt.Println(x1) //[1 3 5 7]
}

func makePractice() {
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a) //[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
}
