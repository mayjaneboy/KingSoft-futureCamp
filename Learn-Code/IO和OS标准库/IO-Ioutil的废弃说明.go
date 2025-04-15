package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 读取整个文件
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}
	fmt.Println("文件内容:", string(content))

	// 写入整个文件
	err = os.WriteFile("output.txt", []byte("Hello, Go I/O!"), 0644)
	if err != nil {
		fmt.Println("写入文件错误:", err)
		return
	}

	// 创建临时文件
	tempFile, err := os.CreateTemp("", "example")
	if err != nil {
		fmt.Println("创建临时文件错误:", err)
		return
	}
	defer os.Remove(tempFile.Name()) // 确保删除临时文件

	fmt.Println("创建的临时文件:", tempFile.Name())

	// 从文件中读取所有内容
	f, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("打开文件错误:", err)
		return
	}
	defer f.Close()

	// 使用io.ReadAll替代ioutil.ReadAll
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}
	fmt.Println("文件内容:", string(data))

	//os.ReadFile()：直接从文件路径读取整个文件内容，其内部封装了打开文件、读取内容和关闭文件的操作
	//io.ReadAll()：从任何 io.Reader 接口读取所有数据，其只负责读取数据，文件的打开和关闭需要调用者自己处理
}
