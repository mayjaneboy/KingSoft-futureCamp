package main

import "fmt"

//定义一个只接收 指针 类型 的参数的函数
func mytest(ptr *int) {
	fmt.Println(*ptr)
}

func main() {
	astr := "hello"
	aint := 1
	abool := false
	arune := 'a'
	afloat := 1.2

	fmt.Printf("astr 指针类型是：%T\n", &astr)
	fmt.Printf("aint 指针类型是：%T\n", &aint)
	fmt.Printf("abool 指针类型是：%T\n", &abool)
	fmt.Printf("arune 指针类型是：%T\n", &arune)
	fmt.Printf("afloat 指针类型是：%T\n", &afloat)

	mytest(&aint)
}
