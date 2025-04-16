package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio 包提供了带缓冲的 I/O 操作
func main() {
	file, err := os.OpenFile("bufwritefile.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("无法打开或创建文件:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	//bufio的Writer 提供了带缓冲的写入功能。
	// 它会在缓冲区满或调用 Flush 方法时，将数据写入底层的 io.Writer（如 os.File）
	// 获取缓冲区大小
	bufferSize := writer.Size()
	fmt.Printf("缓冲区大小: %d 字节\n", bufferSize)
	// 手动创建一个缓冲区为 8KB 的 Writer
	writer2 := bufio.NewWriterSize(file, 8192)
	bufferSize2 := writer2.Size()
	fmt.Printf("缓冲区大小: %d 字节\n", bufferSize2)

	data := []byte("使用 bufio.Writer 写入的数据")
	_, err = writer.Write(data)
	if err != nil {
		fmt.Println("写入时出错:", err)
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区时出错:", err)
		return
	}
	fmt.Println("数据已成功写入")
}
