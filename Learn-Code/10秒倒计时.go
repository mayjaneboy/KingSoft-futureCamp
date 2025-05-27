package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("10秒倒计时开始")
	var i int32
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		fmt.Println()
		time.Sleep(time.Second)
	}
	fmt.Println(i)
}
