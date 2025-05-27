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
	fmt.Println("这句话不会被打印")
}

//panic一但发生，程序会沿调用栈向上传播，直到被recover捕获或程序终止。
//defer语句会在函数返回前执行，即使函数中发生了panic，也会执行。
