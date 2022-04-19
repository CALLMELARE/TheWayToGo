package main

import "fmt"

func main() {
	// byte uint8的别名 ASCII
	// rune int32的别名
	var c1 byte = 'a'
	var c2 rune = 'a'
	fmt.Println(c1, c2)                 // 97 97
	fmt.Printf("c1=%T c2=%T\n", c1, c2) // c1=uint8 c2=int32

	s := "English中文"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
		// Englishä¸­æ
	}
	for _, r := range s {
		fmt.Printf("%c", r)
		// English中文
	}
}
