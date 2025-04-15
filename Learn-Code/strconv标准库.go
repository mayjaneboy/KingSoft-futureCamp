package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整数与字符串转换
	i := 42
	s := strconv.Itoa(i) // 整数转字符串
	fmt.Printf("整数 %d 转字符串: %s\n", i, s)

	s = "42"
	i, err := strconv.Atoi(s) // 字符串转整数
	if err == nil {
		fmt.Printf("字符串 %s 转整数: %d\n", s, i)
	}

	// 解析布尔值
	b, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Printf("解析的布尔值: %t\n", b)
	}

	// 解析浮点数
	f, err := strconv.ParseFloat("3.14159", 64)
	if err == nil {
		fmt.Printf("解析的浮点数: %f\n", f)
	}

	// 解析整数(带基数)
	// i, err = strconv.ParseInt("1010", 2, 64) // 二进制解析
	// if err == nil {
	// 	fmt.Printf("二进制 1010 解析为: %d\n", i)
	// }

	// i, err = strconv.ParseInt("FF", 16, 64) // 十六进制解析
	// if err == nil {
	// 	fmt.Printf("十六进制 FF 解析为: %d\n", i)
	// }

	// 格式化布尔值
	s = strconv.FormatBool(true)
	fmt.Printf("格式化布尔值 true: %s\n", s)

	// 格式化浮点数
	s = strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Printf("格式化浮点数 3.14159 (保留2位小数): %s\n", s)

	// 格式化整数
	s = strconv.FormatInt(255, 16)
	fmt.Printf("格式化整数 255 为十六进制: %s\n", s)
}
