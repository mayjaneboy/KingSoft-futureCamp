package main

import (
	"bytes"         //用于将 JSON 数据转换为可以传输的字节流
	"encoding/json" //提供 JSON 编码和解码功能
	"fmt"
	"net/http"
)

func main() {
	// 准备要发送的数据	创建一个字符串映射 map[string]string
	data := map[string]string{
		"name": "小明",
		"age":  "18",
	}

	jsonData, err := json.Marshal(data) //将 map 编码成 JSON 格式的 {"age":"18","name":"小明"}
	if err != nil {
		fmt.Println("JSON编码失败:", err)
		return
	}

	// 发送POST请求
	resp, err := http.Post("http://example.com/api",
		"application/json",
		bytes.NewBuffer(jsonData))
	//application/json是Content-Type：表明发送的是什么格式的数据（json）
	//bytes.NewBuffer(jsonData)是请求体：把 JSON 字节数据封装为一个读写流，作为 POST 请求的 Body
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("状态码: %d\n", resp.StatusCode)
}
