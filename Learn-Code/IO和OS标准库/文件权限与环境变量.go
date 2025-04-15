package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件并设置权限
	file, err := os.OpenFile("permission.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("创建文件错误:", err)
		return
	}
	file.Close()

	// 修改文件权限
	err = os.Chmod("permission.txt", 0755)
	if err != nil {
		fmt.Println("修改权限错误:", err)
		return
	}
	fmt.Println("文件权限已修改为 0755")

	// 获取环境变量
	home := os.Getenv("HOME")
	fmt.Println("HOME 环境变量:", home)

	// 设置环境变量
	err = os.Setenv("MY_VARIABLE", "hello")
	if err != nil {
		fmt.Println("设置环境变量错误:", err)
		return
	}

	// 读取环境变量
	value := os.Getenv("MY_VARIABLE")
	fmt.Println("MY_VARIABLE 环境变量:", value)

	// 获取所有环境变量
	fmt.Println("\n部分环境变量:")
	count := 0
	for _, env := range os.Environ() {
		if count < 5 { // 只显示前5个，避免输出过多
			fmt.Println(env)
			count++
		} else {
			break
		}
	}

	// 清理
	os.Remove("permission.txt")
}
