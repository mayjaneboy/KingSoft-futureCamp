package main

import "fmt"

func main() {
	i := 0
	//defer 的参数在 程序执行流程经过defer语句 时确定值，而不是在defer语句实际执行时确定。
	defer fmt.Println("deferred value:", i) // 会打印0，而不是1
	i++
	fmt.Println("current value:", i)
}
