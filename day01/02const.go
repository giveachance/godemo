package main

//const
const (
	n1 = 100
	n2
	n3
)

//const
const n4 = 200

//iota 在每个常量声明中值重置为0
//每新增一行常量声明，iota将+1
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

//iota
const (
	b1 = iota //0
	b2 = 100  //100
	b3 = iota //2
)

//iota 单行声明多个常量
const (
	c1, c2 = iota, iota         //c1:0,c2:0
	c3, c4 = iota + 1, iota + 2 //c3:2,c4:3
	_,c5 = iota,iota   //c5:2
)

func testConst() {

	println(c1)
	println(c2)
	println(c3)
	println(c4)
	println(c5)
}
