package main

import "fmt"

func main() {
	var a [3]int
	var b [4]int
	fmt.Println(a) // [0 0 0]
	fmt.Println(b) // [0 0 0 0]

	// 数组初始化
	// 1 定义时用初始值列表
	var cityArr = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArr) // [北京 上海 广州 深圳]
	// 2 编译器推导数组长度
	var boolArr = [...]bool{true, false, false, true, false}
	fmt.Println(boolArr) // [true false false true false]
	// 3 使用索引值初始化
	var langArr = [...]string{1: "Go", 3: "Python", 4: "Java"}
	fmt.Println(langArr)      // [ Go  Python Java]
	fmt.Printf("%T", langArr) // [5]string

	// 数组遍历
	// 1 for 循环遍历
	for i := 0; i < len(cityArr); i++ {
		fmt.Println(cityArr[i])
	}
	// 2 for range 遍历
	for _, val := range cityArr {
		fmt.Println(val)
	}

	// 二维数组
	cityMatrix := [3][2]string{
		{"北京", "西安"},
		{"上海", "重庆"},
		{"杭州", "成都"},
	}
	fmt.Println(cityMatrix)       // [[北京 西安] [上海 重庆] [杭州 成都]]
	fmt.Println(cityMatrix[2][0]) // 杭州
	for _, x := range cityMatrix {
		for _, y := range x {
			fmt.Printf("%s", y) // 北京西安上海重庆杭州成都
		}
	}
}
