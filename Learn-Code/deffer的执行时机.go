package main

import (
	"fmt"
)

func example() {
	defer fmt.Println("这句话会被打印，即使有panic")

	fmt.Println("函数开始执行")
	panic("出现严重错误！")
	fmt.Println("这句话不会被打印")
}

func main() {
	defer fmt.Println("主函数结束")

	fmt.Println("调用example函数前")
	example()
	fmt.Println("这句话不会被打印") //无论这行代码前加不加defer，都不会执行这行代码
}
