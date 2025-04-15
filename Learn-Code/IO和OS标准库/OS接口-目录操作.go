package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 创建目录
	err := os.Mkdir("testdir", 0755)
	if err != nil {
		fmt.Println("创建目录错误:", err)
		return
	}
	fmt.Println("目录已创建")

	// 创建嵌套目录
	err = os.MkdirAll("testdir/subdir/subsubdir", 0755)
	if err != nil {
		fmt.Println("创建嵌套目录错误:", err)
		return
	}
	fmt.Println("嵌套目录已创建")

	// 创建文件到子目录
	file, err := os.Create("testdir/subdir/test.txt")
	if err != nil {
		fmt.Println("创建文件错误:", err)
		return
	}
	file.WriteString("测试内容")
	file.Close()

	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录错误:", err)
		return
	}
	fmt.Println("当前工作目录:", currentDir)

	// 遍历目录
	fmt.Println("遍历目录:")
	err = filepath.Walk("testdir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("%-40s %-10v %8d 字节\n", path, info.IsDir(), info.Size())
		return nil
	})
	if err != nil {
		fmt.Println("遍历目录错误:", err)
	}

	// 读取目录内容
	fmt.Println("\n读取顶层目录内容:")
	entries, err := os.ReadDir("testdir")
	if err != nil {
		fmt.Println("读取目录错误:", err)
		return
	}

	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Printf("%-20s %-10v %8d 字节\n", entry.Name(), entry.IsDir(), info.Size())
	}

	// 清理：删除目录及其内容
	err = os.RemoveAll("testdir")
	if err != nil {
		fmt.Println("删除目录错误:", err)
		return
	}
	fmt.Println("\n目录及其内容已删除")
}
