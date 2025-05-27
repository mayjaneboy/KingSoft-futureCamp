package main

import (
	"fmt"
)

// Calculator 是一个 接受两个整数参数并返回一个整数 的 函数类型
type Calculator func(int, int) int

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func calculate(a, b int, operation Calculator) int {
	return operation(a, b)
}

func main() {
	fmt.Println(calculate(10, 5, add))      // 输出: 15
	fmt.Println(calculate(10, 5, subtract)) // 输出: 5
	//在这里add和subtract是函数，作为了参数传递，对应的形参是operation，是Calculator类型
}
