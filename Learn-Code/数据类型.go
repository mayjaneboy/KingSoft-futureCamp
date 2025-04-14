package main // 声明包名

import (
	"fmt"
	"unicode"
) // 导入格式化输出包

// main 函数是程序执行的入口
func main() {
	// 打印 Hello, Go! 到控制台
	fmt.Println("Hello, Go!")

	var age int = 20
	var name string = "张三"
	var isStudent bool = true

	fmt.Println("年龄:", age)
	fmt.Println("姓名:", name)
	fmt.Println("是否是学生:", isStudent)

	isLive := true
	fmt.Println("是否活着:", isLive)

	var x, y, z = 10, "hello", true
	fmt.Println(x, y, z)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("i 的值为 %v\n", i)
	fmt.Printf("f 的值为 %v\n", f)
	fmt.Printf("b 的值为 %v\n", b)
	fmt.Printf("s 的值为 %v\n", s)

	const pi = 3.14159
	fmt.Println("圆周率:", pi)

	const (
		Monday = iota
		Tuesday
		Wednesday
	)
	fmt.Println(Monday, Tuesday, Wednesday)

	// 判断字符是否为字母
	fmt.Println(unicode.IsLetter('A')) // true
	fmt.Println(unicode.IsLetter('中')) // true
	fmt.Println(unicode.IsLetter('1')) // false

	// 判断字符是否为数字
	fmt.Println(unicode.IsDigit('5')) // true

	// 判断字符是否为空白符
	fmt.Println(unicode.IsSpace(' '))  // true
	fmt.Println(unicode.IsSpace('\t')) // true

	// 大小写转换
	fmt.Println(string(unicode.ToUpper('a'))) // A
	fmt.Println(string(unicode.ToLower('A'))) // a

	//在 Go 中，布尔值不能与整数进行转换，也不能用于算术运算
}
