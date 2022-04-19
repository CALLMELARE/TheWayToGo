package main

import "fmt"

func main() {
	// if 基本写法
	var score = 65
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}
	// if 特殊写法
	if height := 176; height >= 180 {
		fmt.Println("High")
	} else if height >= 165 {
		fmt.Println("Medium")
	} else {
		fmt.Println("Low")
	}
	// for 基本写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// for 省略初始语句
	var j = 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}
	// for 省略初始语句和结束语句
	var k = 0
	for k < 10 {
		fmt.Println(k)
		k++
	}
	// for 无限循环
	for {
		fmt.Println("Infinite Loop")
	}
	// break 跳出循环
	for x := 0; x < 5; x++ {
		fmt.Println(x)
		if x == 3 {
			break
		}
	}
	// continue 继续下一次循环
	for y := 0; y < 5; y++ {
		if y == 3 {
			continue
		}
		fmt.Println(y)
	}
	// swicth single case
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效输入")
	}
	// switch multi case
	num := 5
	switch num {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}
	// switch case 表达式
	age := 30
	switch {
	case age >= 18:
		fmt.Println("成年")
	default:
		fmt.Println("未成年")
	}
}
