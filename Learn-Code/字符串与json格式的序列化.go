package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	myMap := map[int]string{1: "one", 2: "two"}
	fmt.Println(myMap)

	test, err := json.Marshal(myMap)
	//json格式只支持字符串类型的键，所以这里会自动把int类型的键转换为字符串类型的键
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(test)) // 转换为字符串输出

	var anthorMap map[int]string
	err1 := json.Unmarshal(test, &anthorMap)
	// 反序列化时，json格式的键是字符串类型的，所以这里可以自动把字符串类型的键转换为int类型的键
	if err1 != nil {
		fmt.Println("Error unmarshaling JSON:", err1)
	}
	fmt.Println(anthorMap) // 输出反序列化后的结果
	delete(anthorMap, 1)
	fmt.Println(anthorMap) // 输出反序列化后的结果

	m := map[string]int{}
	fmt.Println(m) // 输出反序列化后的结果

}
