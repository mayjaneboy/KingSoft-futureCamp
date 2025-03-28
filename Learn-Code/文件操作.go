package main

import (
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0644)
	// if err != nil {
	// 	fmt.Println("open file failed")
	// 	return
	// } else {
	// 	fmt.Println("open file success")
	// 	data := []byte("这是要写入的数据")
	// 	n, err := file.Write(data)
	// 	if err != nil {
	// 		fmt.Println("write file failed")
	// 		return
	// 	}
	// 	fmt.Println("write file success")
	// 	fmt.Println("write file bytes:", n)
	// }
	// defer file.Close() //延迟关闭文件

	// if err != nil {
	// 	fmt.Println("无法打开或创建文件:", err)
	// 	return
	// }
	// defer file.Close()

	// str := "使用 WriteString 写入的字符串"
	// n, err := file.WriteString(str)
	// if err != nil {
	// 	fmt.Println("写入文件时出错:", err)
	// 	return
	// }
	// fmt.Printf("成功写入 %d 个字节\n", n)

}
