package main

import "fmt"

func main() {
	// 算术运算符
	a := 10
	b := 2
	fmt.Println(a + b) // 12
	fmt.Println(a - b) // 8
	fmt.Println(a * b) // 20
	fmt.Println(a / b) // 5
	fmt.Println(a % b) // 0
	// 关系运算符
	fmt.Println(10 > 2)  // true
	fmt.Println(10 != 2) // true
	fmt.Println(10 <= 2) // false
	// 逻辑运算符
	fmt.Println(10 > 5 && 2 == 2) // true
	fmt.Println(!(10 > 5))        // false
	fmt.Println(1 > 5 || 2 == 2)  // true
	// 位运算符
	c := 1              // 001
	d := 5              // 101
	fmt.Println(c & d)  // 001=>1
	fmt.Println(c | d)  // 101=>5
	fmt.Println(c ^ d)  // 100=>4
	fmt.Println(1 << 2) // 100=>4
	fmt.Println(4 >> 2) // 001=>1
	// 赋值运算符
	var e int
	e = 5
	e += 1 // 6
	e -= 2 // 4
	e *= 2 // 8
	e /= 4 // 2
	e %= 2 // 0
	fmt.Println(e)
}
