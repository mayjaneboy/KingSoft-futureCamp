package main

import "fmt"

// 统计字符串中各字符出现的次数
func countCharacters(s string) map[rune]int {
	counter := make(map[rune]int)
	for _, char := range s {
		counter[char]++
	}
	return counter
}

func main() {
	str := "hello, 世界"
	counts := countCharacters(str)
	for char, count := range counts {
		fmt.Printf("字符 '%c' 出现 %d 次\n", char, count)
	}
}
