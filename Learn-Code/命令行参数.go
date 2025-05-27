package main

import (
	"flag"
	"fmt"
)

func main() {
	//定义命令行参数
	name := flag.String("name", "世界", "姓名")
	age := flag.Int("age", 0, "年龄")
	married := flag.Bool("married", false, "是否已婚")
	// flag.String(name string, defaultValue string, usage string) *string	获取命令行参数中的一个名为name的字符串类型的参数
	//第一个参数是命令行参数的名字，比如--name
	//第二个参数是默认值，当用户为传入该参数时使用
	//第三个参数是参数说明，在-h或-help时显示
	//返回的是 *string，要用 *name 才能拿到真实的值
	//类似的还要flag.Float64等

	//解析命令行参数
	flag.Parse()
	//使用参数
	fmt.Println("你好，", *name)

	if *age > 0 {
		fmt.Printf("你的年龄：%d\n", *age)
	}
	if *married {
		fmt.Println("你已婚")
	} else {
		fmt.Println("你未婚")
	}
	// height := flag.Float64("height", 170, "身高")
	// hobby := flag.String("hobby", "dance", "爱好")
	fmt.Println("其他参数：", flag.Args())
	//flag.Args()用于获取命令行中的非选项参数，即没有“-”开头的
	//go run test.go -name=bravo -age=25 -married=true 175.5 sing
	//175.5 sing就时非选项参数
}
