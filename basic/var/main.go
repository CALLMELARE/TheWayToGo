package main

func main() {
	// 标准声明格式
	var name string
	var age int
	var isOk bool
	// 批量声明格式
	var (
		str   string
		num   int
		flag  bool
		digit float32
	)
	// 声明并赋值
	var text1 string = "LARE"
	var size1 int = 18
	var text2, size2 = "LARE", 18
	// 类型推导，编译器根据初始值推导其类型
	var text3 = "LARE"
	var size3 = 18
	// 短变量声明
	m := 10
	// 匿名变量，_表示忽略值
	x, _ := foo()
}
