package main // 声明包名

import "fmt" // 导入格式化输出包

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
}
