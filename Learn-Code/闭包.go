package main

import (
	"fmt"
)

// makeCounter返回的是一个返回整数的匿名函数，因为这个匿名函数使用了外部变量即count，所以它是一个闭包。
// 闭包可以记住它被创建时的环境，在这里就是使用这个闭包函数时，对count的使用会被记住。
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := makeCounter()
	fmt.Println(counter()) // 输出: 1
	fmt.Println(counter()) // 输出: 2
	fmt.Println(counter()) // 输出: 3

	// 创建新计数器
	counter2 := makeCounter()
	fmt.Println(counter2()) // 输出: 1
}
