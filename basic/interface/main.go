// Go语言中接口（interface）是一种类型，一种抽象的类型
package main

import "fmt"

// 接口的定义
type Singer interface {
	Sing()
}

type Duck struct{}

func (d Duck) Sing() {
	fmt.Println("Ga Ga Ga!")
}

func main() {
	Donald := Duck{}
	Donald.Sing() // Ga Ga Ga!
}
