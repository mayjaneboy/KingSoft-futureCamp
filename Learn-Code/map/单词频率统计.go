package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go语言是一个开源的编程语言 它能让构造简单 可靠且高效的软件变得容易 Go语言是从2007年末由Robert Griesemer Rob Pike及Ken Thompson开始设计 Go是一个快速 静态类型 编译型语言 但是它又像动态类型的解释型语言"

	// 将文本分割成单词
	words := strings.Fields(text)

	// 创建一个map来存储每个单词出现的次数
	frequency := make(map[string]int)

	// 统计每个单词出现的次数
	for _, word := range words {
		frequency[word]++
	}

	// 打印结果
	fmt.Println("单词频率统计:")
	for word, count := range frequency {
		fmt.Printf("%s: %d次\n", word, count)
	}

	// 找出出现次数最多的单词
	maxWord := ""
	maxCount := 0
	for word, count := range frequency {
		if count > maxCount {
			maxCount = count
			maxWord = word
		}
	}

	fmt.Printf("\n出现次数最多的单词是 \"%s\"，共出现了 %d 次\n", maxWord, maxCount)
}
