package main

import (
	"fmt"
	"sort"
)

func main() {
	// 基于数组得到切片
	a := [5]int{11, 22, 33, 44, 55}
	b := a[1:4]
	fmt.Println(a) // [11 22 33 44 55]
	fmt.Println(b) // [22 33 44]

	// 切片再次切片
	c := b[0 : len(b)-1] // [22 33]
	fmt.Println(c)

	// make函数构造切片
	d := make([]int, 5, 10)
	fmt.Println(d) // [0 0 0 0 0]

	// nil
	var e []int           // 声明
	fmt.Println(e == nil) // true
	var f = []int{}       // 声明并初始化
	fmt.Println(f == nil) // false

	// 切片的赋值拷贝
	g := make([]int, 3)
	h := g
	h[0] = 100
	fmt.Println(g) // [100 0 0]
	fmt.Println(h) // [100 0 0]

	// 切片的遍历
	i := []int{1, 2, 3, 4, 5}
	for x := 0; x < len(i); x++ {
		// 根据索引遍历
		fmt.Println(x, i[x])
	}
	for index, val := range i {
		// for-range遍历
		fmt.Println(index, val)
	}

	// 切片的扩容
	var j []int // 声明时未申请内存
	j = append(j, 10)
	fmt.Println(j) // [10] 容量1,扩容时容量翻倍(1,2,4,8,16...)
	j = append(j, 1, 2, 3, 4, 5)
	fmt.Println(j) // [10 1 2 3 4 5] 容量8
	k := []int{6, 7, 8}
	j = append(j, k...)
	fmt.Println(j) // [10 1 2 3 4 5 6 7 8]

	// 切片的拷贝
	l := []int{1, 2, 3, 4, 5}
	m := make([]int, 5, 5)
	copy(m, l)
	m[0] = 100
	fmt.Println(l) // [1 2 3 4 5]
	fmt.Println(m) // [100 2 3 4 5]

	// 切片的元素删除
	n := []string{"北京", "上海", "广州", "深圳"}
	n = append(n[0:2], n[3:]...)
	fmt.Println(n) // [北京 上海 深圳]

	// 切片的排序
	o := []int{4, 5, 7, 2, 3}
	sort.Ints(o[:])
	fmt.Println(o) // [2 3 4 5 7]
}
