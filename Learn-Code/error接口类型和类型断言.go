package main

import (
	"errors"
	"fmt"
	"os"
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
	if err := doSomething(); err != nil {
		//doSomething 函数返回的(err)是MyError类型的指针，但err是一个error接口类型的变量
		//因为MyError实现了error接口，所以可以赋值给err，也就是接口变量 err 内部保存了 *MyError 这个具体类型的值和类型信息
		fmt.Println(err) // 输出: 错误码：404，错误信息：资源未找到
		//因为err实现了了 error的Error接口，索引直接打印err时，实际上时是调用了实现的这个接口Error方法，从而打印的是这个接口返回的字符串。

		// 类型断言获取详细错误信息
		if myErr, ok := err.(*MyError); ok {
			fmt.Printf("错误码：%d\n", myErr.Code)
		}

		//类型断言：x.(T)
		//x 是一个接口类型的变量，对应这里的err，它的类型是error，它存储了一个具体类型*MyError的值
		//T 是想要断言成的具体类型，对应这里的是 *MyError，是你想从接口里提取出来的具体类型。
		//类型断言的作用是：从接口变量 x 中提取它的具体类型 T 的值，如果接口底层的实际类型就是 T，则断言成功。

	}

	_, err := os.Open("non_existent_file.txt")
	// errors.Is判断错误是否为特定错误
	if errors.Is(err, os.ErrNotExist) {
		// 处理文件不存在的情况
	}

	// 将错误转换成特定类型的错误
	var pathError *os.PathError
	if errors.As(err, &pathError) {
		//errors.As 在错误链中查找第一个 能转换成指定类型 的错误，并把它赋值给目标变量
		fmt.Println("路径错误:", pathError.Path)
	}
}
