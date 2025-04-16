package main

import "fmt"

func main() {
	// 按年龄分组用户
	users := []struct {
		Name string
		Age  int
	}{
		{"张三", 20},
		{"李四", 30},
		{"王五", 20},
		{"赵六", 25},
	}
	//users是一个struct数组

	groups := make(map[int][]string)
	//groups是一个键为int型，值为字符串数组的map

	for _, user := range users {
		//遍历users，拿到数组中的结构体元素
		//更新结构体元素对应的Age对应的string数组
		groups[user.Age] = append(groups[user.Age], user.Name)
	}

	// 打印分组结果
	for age, names := range groups {
		fmt.Printf("%d岁: %v\n", age, names)
	}
}
