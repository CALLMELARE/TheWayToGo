package main

import "fmt"

// 结构体的定义
type person struct {
	name, city string
	age        int8
}

func main() {
	// 自定义类型
	type myInt int
	// 类型别名
	type myIntAlias = myInt

	// 基本实例化
	var jack person
	jack.name = "杰克"
	jack.city = "北京"
	jack.age = 20
	fmt.Println(jack) // {杰克 北京 20}

	// 匿名结构体
	var user struct {
		name string
		age  int
	}
	user.name = "Anonymous"
	user.age = 18
	fmt.Println(user) // {Anonymous 18}

	// 创建指针类型结构体
	var tom = new(person)
	fmt.Printf("%#v\n", tom) // &main.person{name:"", city:"", age:0}

	// 取结构体的地址实例化
	jerry := &person{}
	fmt.Printf("%#v\n", jerry) // &main.person{name:"", city:"", age:0}
	jerry.name = "杰瑞"
	jerry.city = "北京"
	jerry.age = 8
	fmt.Printf("%#v\n", jerry) // &main.person{name:"杰瑞", city:"北京", age:8}

	// 使用键值对初始化结构体
	bella := person{
		name: "贝拉",
		city: "杭州",
		age:  19,
	}
	fmt.Printf("%#v\n", bella) // main.person{name:"贝拉", city:"杭州", age:19}

	// 使用值的列表初始化
	diana := person{
		"戴安娜",
		"杭州",
		19,
	}
	fmt.Printf("%#v\n", diana) // main.person{name:"戴安娜", city:"杭州", age:19}
}
