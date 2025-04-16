package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	//构建一个读取file的reader
	reader := bufio.NewReader(file)

	//按字节读取
	// buf := make([]byte, 5)
	// n, err := reader.Read(buf)
	// if err != nil {
	// 	fmt.Println("读取时出错:", err)
	// 	return
	// }
	// fmt.Printf("读取了 %d 个字节: %s\n", n, string(buf[:n]))

	//按行读取
	for {
		line, err := reader.ReadString('\n')
		fmt.Print(line)

		if err == io.EOF {
			fmt.Println()
			break
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
