package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello Goroutine!")
}
func main() {
	go hello()
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)
}

//注意：main 函数也是一个 goroutine，当 main 函数执行完毕，程序就结束了。
//如果不加time.Sleep(time.Second)，就不会输出Hello Goroutine!
//因为main也是一个goroutine，首先创建main的goroutin，main执行时，创建hello的goroutine，
// 当main的执行完了后，hello的goroutine也不一定创建成功
