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

// 在结构体字段上使用了 json:"name" 等标签，用于指定结构体字段在 JSON 中的名称。
// 如果不指定标签，默认使用结构体字段的名称作为 JSON 中的键
func main() {
	// 创建一个Person结构体实例
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Hobbies: []string{"reading", "writing", "coding"},
	}

	// 将Person结构体转换为JSON格式的字节切片
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 输出JSON数据
	fmt.Println(string(jsonData))

	fmt.Println("JSON->结构体")
	// JSON数据
	jsonData2 := `{"name":"John Doe","age":30,"hobbies":["reading","writing","coding"]}`

	// 将JSON格式的字节切片转换为Person结构体
	var me Person
	err2 := json.Unmarshal([]byte(jsonData2), &me)
	//在使用 json.Unmarshal 函数时，需要传入一个指向结构体的指针作为第二个参数，用于接收转换后的结构体数据。
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	// 输出Person结构体的字段值
	fmt.Println("Name:", me.Name)
	fmt.Println("Age:", me.Age)
	fmt.Println("Hobbies:", me.Hobbies)
}
