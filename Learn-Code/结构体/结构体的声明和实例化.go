package main

import "fmt"

// 定义一个名为Person的结构体
type Person struct {
	Name string
	Age  int
	City string
}

// 同样类型的字段也可以写在一行，
type person1 struct {
	name, city string
	age        int8
}

func main() {
	// 1. 使用var声明
	var p1 Person

	// 2. 使用字面量
	p2 := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}

	// 3. 使用new关键字
	p3 := new(Person)

	//键值对初始化
	p4 := Person{
		Name: "张三",
		Age:  18,
	}

	//值得列表初始化
	p5 := Person{
		"李四",
		20,
		"北京",
	}
	//这种方式必须初始化结构体的所有字段，且顺序必须与结构体定义时字段的顺序一致。

	// 打印结构体内容
	fmt.Printf("%+v\n", p1)
	// 打印结构体类型和内容
	fmt.Printf("%#v\n", p1)

	fmt.Printf("%+v\n", p2)
	// 打印结构体类型和内容
	fmt.Printf("%#v\n", p2)

	fmt.Printf("%+v\n", p3)
	// 打印结构体类型和内容
	fmt.Printf("%#v\n", p3)

	fmt.Printf("%+v\n", p4)
	// 打印结构体类型和内容
	fmt.Printf("%#v\n", p4)

	fmt.Printf("%+v\n", p5)
	// 打印结构体类型和内容
	fmt.Printf("%#v\n", p5)
}
