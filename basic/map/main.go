package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// 仅声明map类型，不进行初始化
	var a map[string]int
	fmt.Println(a == nil) // true

	// map 的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil) // false

	// map 添加键值对
	a["id"] = 1
	a["score"] = 97
	fmt.Printf("%#v\n", a) // map[string]int{"id":1, "score":97}

	// 声明的同时完成初始化
	b := map[int]bool{
		1: true, 2: false,
	}
	fmt.Printf("%#v\n", b) // map[int]bool{1:true, 2:false}

	// 判断某个键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["张三"] = 100
	val, ok := scoreMap["张三"]
	fmt.Println(val, ok) // 100 true

	// 遍历map
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 只遍历value
	for _, v := range scoreMap {
		fmt.Println(v)
	}

	// 删除键值对
	delete(scoreMap, "张三")
	fmt.Println(scoreMap) // map[]

	// 按照特定顺序遍历map
	var newMap = make(map[string]int, 100)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("uid%02d", i)
		value := rand.Intn(100)
		newMap[key] = value
	}
	keys := make([]string, 0, 100)
	for k := range newMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, newMap[key])
	}

	// 元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8)
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["张三"] = 100
	fmt.Println(mapSlice) // [map[张三:100] map[] map[] map[] map[] map[] map[] map[]]

	// 值为切片的map
	var sliceMap = make(map[string][]int, 8)
	v, ok := sliceMap["李四"]
	fmt.Println(v, ok) // [] false
	if !ok {
		sliceMap["李四"] = make([]int, 8)
		sliceMap["李四"][0] = 10
		sliceMap["李四"][2] = 20
		sliceMap["李四"][4] = 40
	}

	// 遍历sliceMap
	for k, v := range sliceMap {
		fmt.Println(k, v) // 李四 [10 0 20 0 40 0 0 0]
	}

	// 统计一个字符串中每个单词出现的次数
	var wordCount = make(map[string]int, 10)
	var str = "how to go to the way to go"
	words := strings.Split(str, " ")
	for _, w := range words {
		v, ok := wordCount[w]
		if ok {
			wordCount[w] = v + 1
		} else {
			wordCount[w] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Print(k, v) // way1how1to3go2the1
	}

}
