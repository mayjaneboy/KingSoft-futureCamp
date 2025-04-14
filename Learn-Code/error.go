package main

import (
	"errors"
	"fmt"
)

// 自定义错误类型
type MyError struct {
	Code    int
	Message string
}

// 使用自定义错误
func doSomething() error {
	return &MyError{
		Code:    404,
		Message: "资源未找到",
	}
}

// 实现 error 接口
func (e *MyError) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.Code, e.Message)
}
func main() {
	// 使用 errors.New() 创建错误
	err1 := errors.New("这是一个错误消息")
	if err1 != nil {
		fmt.Println(err1) // 输出: 这是一个错误消息
	}

	// 使用 fmt.Errorf() 创建格式化错误
	name := "文件"
	err2 := fmt.Errorf("%s不存在", name)
	if err2 != nil {
		fmt.Println(err2) // 输出: 文件不存在
	}

	//通过实现 error 接口来创建自定义错误类型
	if err3 := doSomething(); err3 != nil {
		fmt.Println(err3) // 输出: 错误码：404，错误信息：资源未找到

		// 类型断言获取详细错误信息
		if myErr, ok := err3.(*MyError); ok {
			//如果变量err3是MyError类型，则ok=true 否则false
			fmt.Printf("错误码：%d\n", myErr.Code)
		}
	}
}
