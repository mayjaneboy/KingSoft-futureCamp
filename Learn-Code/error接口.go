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

// 实现 error 接口
func (e *MyError) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.Code, e.Message)
}

// 使用自定义错误
func doSomething() error {
	return &MyError{
		Code:    404,
		Message: "资源未找到",
	}
}

func main() {
	err := errors.New("这是一个错误")
	fmt.Println(err)

	err2 := fmt.Errorf("错误信息: %w", err)
	fmt.Println(err2)

	if err1 := doSomething(); err1 != nil {
		fmt.Println(err1)
	}
}
