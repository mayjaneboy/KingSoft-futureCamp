package main

import "fmt"

func main() {
	// 字符字面量使用单引号，默认类型为 rune
	ch := '中'                              // ch 的类型是 rune
	fmt.Printf("%T, %v, %c\n", ch, ch, ch) // 输出: int32, 20013, 中

	// 遍历字符串中的 Unicode 字符（而非字节）
	for _, r := range "Hello, 世界" {
		fmt.Printf("%c-", r) // 输出: H e l l o ,   世 界
	}
	fmt.Println()
	// 字符串与 rune 切片的转换
	s := "Go语言"
	runes := []rune(s)      // 字符串转换为 rune 切片
	fmt.Println(len(s))     // 输出: 8 (字节数)
	fmt.Println(len(runes)) // 输出: 4 (字符数)

	// rune 切片转回字符串
	s2 := string(runes) // rune 切片转换为字符串
	fmt.Println(s2)     // 输出: Go语言

	// 修改字符串中的字符
	runes[2] = '编' // 修改第三个字符
	s3 := string(runes)
	fmt.Println(s3) // 输出: Go编言
}
