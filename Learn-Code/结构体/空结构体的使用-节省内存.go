package main

import (
	"fmt"
)

func main() {
	// 使用map和空结构体来表示集合
	uniqueIDs := make(map[string]struct{})
	//map 的值类型为 struct{}，它不占用额外的内存空间，仅作为一种占位符表示键（即 ID）存在于集合中
	//相比使用 map[string]bool，使用 map[string]struct{} 可以节省内存，因为 bool 类型至少占用一个字节，而空结构体不占用任何内存
	id1 := "123"
	id2 := "456"

	// 将ID添加到集合中
	uniqueIDs[id1] = struct{}{}
	uniqueIDs[id2] = struct{}{}
	//第一个 {} 是用来定义空结构体类型 struct{}
	//第二个 {} 是用来创建该空结构体类型的实例

	// 检查ID是否存在
	if _, exists := uniqueIDs["123"]; exists {
		fmt.Println("ID 123 exists in the set.")
	}
}
