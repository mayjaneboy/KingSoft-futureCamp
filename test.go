package main

import (
	"fmt"
	"sync"
)

// 模拟从文件读取数据的函数
func readFile(fileName string, resultChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	// 这里简单模拟读取文件，实际应用中应使用文件读取操作
	var data string
	switch fileName {
	case "file1.txt":
		data = "Content of file1"
	case "file2.txt":
		data = "Content of file2"
	case "file3.txt":
		data = "Content of file3"
	}
	resultChan <- fmt.Sprintf("Read data from %s: %s", fileName, data)
}

func main() {
	fileNames := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wg sync.WaitGroup
	resultChan := make(chan string)

	for _, fileName := range fileNames {
		wg.Add(1)
		go readFile(fileName, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Println(result)
	}

	fmt.Println("All files have been read.")
}
