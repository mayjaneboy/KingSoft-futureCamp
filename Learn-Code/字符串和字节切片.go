package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	ch := '中'                              // ch 的类型是 rune
	fmt.Printf("%T, %v, %c\n", ch, ch, ch) // 输出: int32, 20013, 中
	num := 20013                           // num 的类型是 int
	fmt.Printf("%c", num)                  // 输出: 中

	s := "你好abc"

	for i := 0; i < len(s); i++ {
		fmt.Printf("字节 %d: %x (%q)\n", i, s[i], s[i])
	}
	b := []byte(s)

	for i, by := range b {
		fmt.Printf("字节 %d: %x (%q)\n", i, by, by)
	}

	// 字符串与 rune 切片的转换
	s1 := "Go语言"
	runes := []rune(s1)     // 字符串转换为 rune 切片
	fmt.Println(len(s1))    // 输出: 8 (字节数)
	fmt.Println(len(runes)) // 输出: 4 (字符数)

	// rune 切片转回字符串
	s2 := string(runes)  // rune 切片转换为字符串
	fmt.Println(s2)      // 输出: Go语言
	fmt.Println(len(s2)) // 输出: 8 (字节数)

	// 修改字符串中的字符
	runes[2] = '编' // 修改第三个字符，字符串不能直接修改，但可以转换为 rune 字符切片修改
	s3 := string(runes)
	fmt.Println(s3) // 输出: Go编言

	var s4 string = "hello"
	s5 := "世界"

	s6 := s4 + " " + s5
	length := len(s6) //获取字符串的字节数
	println(length)

	count := utf8.RuneCountInString(s2)
	println(count) //获取字符串的字符数
}
