package main

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("错误码: %d, 信息: %s", e.Code, e.Message)
}

func process() {
	panic(&CustomError{
		Code:    500,
		Message: "严重错误",
	})
}

func main() {
	process()
}

// 上面的示例中，我们定义了一个自定义的错误类型 CustomError，并实现了 Error 方法，
// 所以 CustomError 可以被当作 error 类型使用。
// 当 process 函数触发 panic 时，会创建一个 CustomError 实例，并将其作为参数传递给 panic 函数。
