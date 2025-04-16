package main

import (
	"fmt"
)

// make：仅适用于创建 slice、map 和 channel 这三种引用类型。返回的是类型本身，而不是指向类型的指针
// new：用于为任何类型分配零值内存空间，返回指向该类型零值的指针。
// new适用于所有类型，包括结构体、基本类型（如 int、float 等）以及指针类型等
func main() {
	// make 示例
	// 创建并初始化一个切片
	slice := make([]int, 3)
	fmt.Printf("Slice: %v\n", slice)

	// 创建并初始化一个 map
	m := make(map[string]int)
	fmt.Printf("Map: %v\n", m)
	m["key"] = 1
	fmt.Printf("Map: %v\n", m)

	// 创建并初始化一个通道
	ch := make(chan int)
	fmt.Printf("Channel: %v\n", ch)

	// new 示例
	numPtr := new(int)
	fmt.Printf("Pointer to int: %v, Value: %d\n", numPtr, *numPtr)

	type MyStruct struct {
		a int
		b string
	}
	structPtr := new(MyStruct)
	fmt.Printf("Pointer to struct: %v, a: %d, b: %s\n", structPtr, structPtr.a, structPtr.b)
}
