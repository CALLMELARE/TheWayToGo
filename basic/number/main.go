package main

import (
	"fmt"
	"math"
)

func main() {
	// 十进制
	a := 10
	fmt.Printf("%d\n", a) // 10
	fmt.Printf("%b\n", a) // 1010
	// 八进制 以0开头
	b := 077
	fmt.Printf("%o\n", b) // 77
	// 十六进制 以0x开头
	c := 0xff
	fmt.Printf("%x\n", c) // ff
	fmt.Printf("%X\n", c) // FF
	// uint8 (0~255)
	var age uint8
	age = 255
	fmt.Println(age) // 255
	// 浮点数
	fmt.Println(math.MaxInt32) // 2147483647
	fmt.Println(math.MaxInt64) // 9223372036854775807
	// 复数
	var com1 complex64
	com1 = 1 + 2i
	fmt.Println(com1) // (1+2i)
	var com2 complex128
	com2 = 2 + 4i // (2+4i)
	fmt.Println(com2)
	// 变量的内存地址
	fmt.Printf("%p\n", &a) // 0xc000014088
}
