package main

import "fmt"

func main() {
	var s1 string = "Hello, 世界"
	s2 := "Go 语言"
	// 字符串连接
	s3 := s1 + " " + s2
	fmt.Println(s3)

	sub := s1[0:5]
	fmt.Println(sub)

}
