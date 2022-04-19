package main

import (
	"errors"
	"fmt"
)

// 函数的定义
func twoSum(x int, y int) int {
	return x + y
}

// 函数的定义（无参数，无返回值）
func helloWorld() {
	fmt.Println("Hello world")
}

// 函数参数（类型简写）
func twoTimes(x, y int) int {
	return x * y
}

// 函数参数（可变参数）
func multiSum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 多返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名
func namedCalc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 变量作用域（全局变量）
var globalVar int64 = 10

// 变量作用域（局部变量）
func local() {
	var localVar int64 = 100
	fmt.Println(localVar)
}

// 高阶函数

// 函数作为参数(1/2)
func add(x, y int) int {
	return x + y
}
func opadd(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

// 闭包=函数+引用环境
func closure() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	// 函数的调用
	helloWorld() // Hello world
	sum := twoSum(2, 3)
	fmt.Println(sum) // 5
	mSum := multiSum(1, 2, 3, 4, 5)
	fmt.Println(mSum) // 15

	// 函数作为参数(2/2)
	opans := opadd(2, 3, add)
	fmt.Println(opans) // 5

	// 匿名函数
	anonyFunc := func(x int) {
		fmt.Println(x)
	}
	anonyFunc(1) // 1

	// 自执行函数
	func(x int) {
		fmt.Println(x) // 2
	}(2)

	// defer语句：在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
	fmt.Println("start")
	defer fmt.Println("defer")
	fmt.Println("end")
}
