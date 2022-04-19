package main

import (
	"fmt"
)

// Go语言中的指针不能进行偏移和运算，是安全指针

func modify(x int) {
	x = 100
}

func remodify(y *int) {
	*y = 100
}

func main() {
	// 指针取值
	a := 10
	b := &a
	fmt.Printf("%T\n", b) // *int
	c := *b
	fmt.Printf("%T\n", c) // int
	fmt.Printf("%v\n", c) // 10

	// 指针传值
	modify(a)
	fmt.Println(a) // 10
	remodify(&a)
	fmt.Println(a) // 100

	// new
	var p *int
	p = new(int)
	*p = 10
	fmt.Println(*p) // 10

	// make
	var q map[string]int
	q = make(map[string]int, 10)
	q["张三"] = 100
	fmt.Println(q) // map[张三:100]

	// 二者都是用来做内存分配的。
	// make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	// 而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
}
