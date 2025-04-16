package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	// 创建一个Person结构体实例
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Hobbies: []string{"reading", "writing", "coding"},
	}

	// 将Person结构体转换为格式化的JSON格式的字节切片
	// jsonData, err := json.MarshalIndent(person, "前缀", "--")
	jsonData, err := json.MarshalIndent(person, "", "  ")
	//第一个参数是要转换的结构体，第二个参数是缩进的前缀，第三个参数是缩进的宽度。
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 输出格式化的JSON数据
	fmt.Println(string(jsonData))
}
