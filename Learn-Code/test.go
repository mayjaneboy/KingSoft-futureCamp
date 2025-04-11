package main

import (
	"fmt"
	"time"
)

// 无缓冲channel的咖啡店示例
func coffeeShopUnbuffered() {
	// 创建一个无缓冲的订单通道
	orders := make(chan string)

	// 咖啡师goroutine
	go func() {
		for order := range orders {
			fmt.Printf("正在制作: %s\n", order)
			time.Sleep(2 * time.Second) // 模拟制作时间
			fmt.Printf("%s 制作完成\n", order)
		}
		fmt.Println("所有订单已完成")
	}()

	// 顾客点单（必须等待咖啡师接单）
	go func() {
		orders <- "拿铁"
		fmt.Println("拿铁订单已接收")
		//打印时间
		fmt.Println("拿铁订单接收时间", time.Now().Format("2006-01-02 15:04:05"))
		orders <- "美式"
		fmt.Println("美式订单已接收")
		//打印时间
		fmt.Println("美式订单接收时间", time.Now().Format("2006-01-02 15:04:05"))
		orders <- "卡布奇诺"
		fmt.Println("卡布奇诺订单已接收")
		//打印时间
		fmt.Println("卡布奇诺订单接收时间", time.Now().Format("2006-01-02 15:04:05"))
		close(orders)
	}()

	// 给咖啡师时间完成所有订单
	time.Sleep(4 * time.Second)
}

func main() {
	coffeeShopUnbuffered()
}
