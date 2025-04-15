package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println("当前时间:", now)

	// 创建指定时间
	t := time.Date(2023, time.October, 1, 10, 30, 0, 0, time.Local)
	fmt.Println("指定时间:", t)

	// 时间格式化
	fmt.Println("格式化时间(yyyy-MM-dd):", now.Format("2006-01-02"))
	fmt.Println("格式化时间(yyyy-MM-dd HH:mm:ss):", now.Format("2006-01-02 15:04:05"))

	// 时间解析
	timeStr := "2023-10-01 10:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err == nil {
		fmt.Println("解析的时间:", parsedTime)
	}

	// 时间计算
	tomorrow := now.Add(24 * time.Hour)
	fmt.Println("明天:", tomorrow.Format("2006-01-02"))

	// 时间间隔
	duration := tomorrow.Sub(now)
	fmt.Printf("距离明天还有: %.2f 小时\n", duration.Hours())

	// 休眠示例
	fmt.Println("休眠前:", time.Now().Format("15:04:05"))
	time.Sleep(2 * time.Second) // 休眠2秒
	fmt.Println("休眠后:", time.Now().Format("15:04:05"))
}
