package main

import "fmt"

func main() {
	// 基本打印
	fmt.Println("Hello, Go!") // 输出: Hello, Go!

	// 格式化打印
	name := "张三"
	age := 25
	fmt.Printf("姓名: %s, 年龄: %d\n", name, age) // 输出: 姓名: 张三, 年龄: 25

	// 格式化字符串并返回
	s := fmt.Sprintf("姓名: %s, 年龄: %d", name, age)
	fmt.Println(s) // 输出: 姓名: 张三, 年龄: 25

	// 结构体打印
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "李四", Age: 30}
	fmt.Printf("%v\n", p)  // 输出: {李四 30}
	fmt.Printf("%+v\n", p) // 输出: {Name:李四 Age:30}
	fmt.Printf("%#v\n", p) // 输出: main.Person{Name:"李四", Age:30}
	fmt.Printf("%T\n", p)  // 输出: main.Person

	var str string
	fmt.Print("请输入str:")
	fmt.Scanf("%s", &str)
	fmt.Println("你输入的是" + str)
}
