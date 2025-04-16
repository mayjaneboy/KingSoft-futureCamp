package main

import "fmt"

type Notifier interface {
	Notify()
}

// 使用空结构体实现Notifier接口
type EmptyNotifier struct{}

func (e EmptyNotifier) Notify() {
	fmt.Println("Notification sent.")
}

func main() {
	var n Notifier = EmptyNotifier{}
	n.Notify()
}

//EmptyNotifier 结构体没有任何字段，但它实现了 Notifier 接口的 Notify 方法。
// 这种方式允许在不需要存储任何数据的情况下实现接口功能
