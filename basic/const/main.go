package main

import "fmt"

func main() {
	// 标准声明
	const pi = 3.1415926
	const E = 2.7182818
	// 批量声明，省略值则表示与上一行相等
	const (
		c1 = 1
		c2 = 10
		c3
	)
	fmt.Println(c3) // 10
	// iota，_可跳过某些值
	const (
		n1 = iota // 0
		_
		n3 = iota //2
		n4 = iota //3
		n5 = 100  //100
		n6 = iota //5
	)
	// 利用iota产生数量级
	const (
		_  = iota             //0
		KB = 1 << (10 * iota) //1024
		MB = 1 << (10 * iota) //1024*1024
		GB = 1 << (10 * iota) //1024*1024*1024
		TB = 1 << (10 * iota) //1024*1024*1024*1024
		PB = 1 << (10 * iota) //1024*1024*1024*1024*1024
	)
	// 多个iota定义在一行
	const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
}
