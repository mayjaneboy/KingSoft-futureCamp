package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 发送GET请求
	resp, err := http.Get("http://localhost:8080/index.html")
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应内容: %s\n", body)
}
