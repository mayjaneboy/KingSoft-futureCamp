package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s1 string = "Hello, 世界"
	s2 := "Go 语言"
	// 字符串连接
	s3 := s1 + " " + s2
	fmt.Println(s3)

	// 获取字符串长度（字节数）
	length := len(s1)
	fmt.Println(length)

	// 获取字符串中的字符（Unicode 码点）数量
	count := utf8.RuneCountInString(s1)
	fmt.Println(count) //9

	//字符串切片
	sub := s1[0:5]
	fmt.Println(sub)

	//Go 语言不支持隐式数据类型转换，必须使用显式转换
}
