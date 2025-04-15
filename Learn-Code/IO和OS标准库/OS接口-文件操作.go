package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 创建文件
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("创建文件错误:", err)
		return
	}

	// 写入文件
	data := []byte("你好，这是一个示例文件！\n这是第二行内容。")
	n, err := file.Write(data)
	if err != nil {
		fmt.Println("写入文件错误:", err)
		file.Close()
		return
	}
	fmt.Printf("成功写入 %d 字节到文件\n", n)

	// 立即关闭写入的文件
	file.Close()

	// 打开文件进行读取
	readFile, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("打开文件错误:", err)
		return
	}

	// 读取文件内容
	content, err := ioutil.ReadAll(readFile)
	if err != nil {
		fmt.Println("读取文件错误:", err)
		readFile.Close()
		return
	}
	fmt.Println("文件内容:")
	fmt.Println(string(content))

	// 立即关闭读取的文件
	readFile.Close()

	// 获取文件信息
	fileInfo, err := os.Stat("example.txt")
	if err != nil {
		fmt.Println("获取文件信息错误:", err)
		return
	}
	fmt.Printf("文件名: %s\n", fileInfo.Name())
	fmt.Printf("文件大小: %d 字节\n", fileInfo.Size())
	fmt.Printf("文件权限: %v\n", fileInfo.Mode())
	fmt.Printf("最后修改时间: %v\n", fileInfo.ModTime())

	// 重命名文件
	err = os.Rename("example.txt", "renamed.txt")
	if err != nil {
		fmt.Println("重命名文件错误:", err)
		return
	}
	fmt.Println("文件已重命名为 renamed.txt")

	// 最后删除文件
	err = os.Remove("renamed.txt")
	if err != nil {
		fmt.Println("删除文件错误:", err)
		return
	}
	fmt.Println("文件已删除")
}
