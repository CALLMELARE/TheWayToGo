package main

import (
	"fmt"
	"strings"
)

// UTF-8编码

func main() {
	// English
	str1 := "hello world"
	fmt.Println(str1) // hello world
	// 中文
	str2 := "你好世界"
	fmt.Println(str2) // 你好世界
	// 多行字符串
	multistr := `
	多行字符串
	会原样输出
	`
	fmt.Println(multistr)
	// 多行字符串
	// 会原样输出

	// 拼接字符串
	fmt.Println(str1 + str2) // hello world你好世界
	combine := fmt.Sprintf("%s - %s", str1, str2)
	fmt.Println(combine) // hello world - 你好世界
	// Split方法:字符串分割
	str := "aa bb cc dd"
	fmt.Println(strings.Split(str, " "))        // [aa bb cc dd]
	fmt.Printf("%T\n", strings.Split(str, " ")) // []string
	// Contains方法:判断是否包含子串
	fmt.Println(strings.Contains(str, "cc")) // true
	// HasPrefix方法:判断前缀
	fmt.Println(strings.HasPrefix(str, "aa")) // true
	// HasSuffix方法:判断后缀
	fmt.Println(strings.HasSuffix(str, "dd")) // true
	// Index方法:判断子串最先出现的位置
	fmt.Println(strings.Index(str, "bb")) // 3
	// LastIndex方法:判断子串最后出现的位置
	fmt.Println(strings.LastIndex(str, "bb")) // 3
	// Join方法:连接字符串
	joinstr := []string{"aa", "bb", "cc", "dd"}
	fmt.Println(joinstr)                    // [aa bb cc dd]
	fmt.Println(strings.Join(joinstr, "&")) // aa&bb&cc&dd
}
