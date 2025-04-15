package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义命令行参数
	name := flag.String("name", "世界", "姓名")
	age := flag.Int("age", 0, "年龄")
	married := flag.Bool("married", false, "是否已婚")

	// 解析命令行参数
	flag.Parse()

	// 使用参数
	fmt.Printf("你好, %s!\n", *name)

	if *age > 0 {
		fmt.Printf("你的年龄是: %d\n", *age)
	}

	if *married {
		fmt.Println("你已婚")
	} else {
		fmt.Println("你未婚")
	}

	// 获取非选项参数
	fmt.Println("其他参数:", flag.Args())
}

//go run flag标准库.go -name=张三 -age=25 -married [其他参数1] [其他参数2] ...
//go run flag标准库.go -name=张三 -age=25 -married true
//go run flag标准库.go -name=张三 -age=25 -married
//go run flag标准库.go -name=张三 -age=25
