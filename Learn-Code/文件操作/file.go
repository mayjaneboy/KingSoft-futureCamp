package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("读取到的行:", line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时出错:", err)
	}
}
