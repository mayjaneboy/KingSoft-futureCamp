package main

import (
	"fmt"
	"time"
)

func worker(done chan struct{}) {
	// 模拟工作
	time.Sleep(2 * time.Second)
	fmt.Println("Worker finished")
	// 发送完成信号
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	go worker(done)

	// 等待工作完成信号
	<-done
	fmt.Println("Main function received done signal")
}

//worker goroutine 完成工作后，通过向 done 通道发送一个空结构体值来通知主 goroutine 工作已完成。
// 主 goroutine 通过从 done 通道接收数据来等待这个信号，这里空结构体仅作为信号标识，不需要携带任何数据。
