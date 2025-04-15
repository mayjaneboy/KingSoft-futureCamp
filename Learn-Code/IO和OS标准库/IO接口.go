package main

// 导入需要使用的包
import (
	"fmt"     // 用于格式化输入输出
	"io"      // 提供I/O原语的基本接口
	"os"      // 提供操作系统功能
	"strings" // 提供字符串操作相关功能
)

func main() {
	// 创建一个字符串Reader
	// strings.NewReader 将字符串转换为一个实现了 io.Reader 接口的对象
	// 这样我们就可以像读取文件一样读取字符串
	r := strings.NewReader("Hello, Go I/O!")

	// 创建一个字节切片作为缓冲区
	// 大小为8字节，用于存储每次读取的数据
	// 缓冲区大小的选择会影响读取效率
	buf := make([]byte, 8)

	// 使用无限循环读取数据
	// 每次读取最多8个字节（缓冲区的大小）
	for {
		// r.Read 返回读取的字节数和可能的错误
		// n 表示实际读取的字节数
		n, err := r.Read(buf)

		// io.EOF（End Of File）表示已经读到数据末尾
		if err == io.EOF {
			break // 结束循环
		}

		// 处理其他可能的错误
		if err != nil {
			fmt.Println("读取错误:", err)
			break
		}

		// 打印读取的内容
		// buf[:n] 表示只取实际读取的字节数量
		fmt.Printf("读取了 %d 字节: %s\n", n, buf[:n])
	}

	// 获取标准输出的Writer
	// os.Stdout 是一个 *File 类型，实现了 io.Writer 接口
	w := os.Stdout

	// 直接写入字节切片到标准输出
	// Write 方法接收一个字节切片作为参数
	w.Write([]byte("这是写入到标准输出的内容\n"))

	// 演示使用 io.Copy 进行数据复制
	fmt.Println("\n使用io.Copy复制数据:")
	// 重新创建一个新的Reader
	r = strings.NewReader("这是使用io.Copy复制的数据!")
	// io.Copy 会自动处理数据复制，直到遇到EOF
	// 它内部会自动处理缓冲，比手动循环更方便
	io.Copy(w, r)
	fmt.Println() // 添加最后的换行
}
