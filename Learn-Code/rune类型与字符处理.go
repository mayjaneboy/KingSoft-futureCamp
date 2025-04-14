package main

import "fmt"

//在 Go 语言中，rune 类型是 int32 的别名，专门用来表示 Unicode 码点
//为什么需要 rune 类型？
//在 Go 中，字符串是以 UTF-8 编码存储的。UTF-8 是一种变长编码：
//ASCII 字符（如英文字母、数字）占用 1 个字节
//中文等字符通常占用 3 个字节
//某些特殊符号可能占用 4 个字节
//因此，当我们遍历字符串时，如果按字节遍历，会将一个多字节字符拆分成多个无意义的字节。而 rune 类型正是为了解决这个问题而设计的。
func main() {
	// 字符字面量使用单引号，默认类型为 rune
	ch := '中'                              // ch 的类型是 rune
	fmt.Printf("%T, %v, %c\n", ch, ch, ch) // 输出: int32, 20013, 中

	//%v表示按值的默认格式进行输出，可以用于任何数据类型，会自动选择适合该类型的输出格式
	// 不同类型使用 %v 的例子
	// num := 42
	// fmt.Printf("%v\n", num) // 输出: 42

	// str := "hello"
	// fmt.Printf("%v\n", str) // 输出: hello

	// arr := []int{1, 2, 3}
	// fmt.Printf("%v\n", arr) // 输出: [1 2 3]

	// type Person struct {
	// 	Name string
	// 	Age  int
	// }
	// p := Person{"张三", 25}
	// fmt.Printf("%v\n", p) // 输出: {张三 25}
	//而对于字符，则是输出字符的Unicode 码点值

	// 遍历字符串中的 Unicode 字符（而非字节）
	for _, r := range "Hello, 世界" {
		fmt.Printf("%c-", r) // 输出: H-e-l-l-o-,- -世-界- 主页字符串中的空格
	}
	fmt.Println()

	// 字符串与 rune 切片的转换
	s := "Go语言"
	runes := []rune(s)      // 字符串转换为 rune 切片
	fmt.Println(len(s))     // 输出: 8 (字节数)
	fmt.Println(len(runes)) // 输出: 4 (字符数)

	// rune 切片转回字符串
	s2 := string(runes) // rune 切片转换为字符串
	fmt.Println(s2)     // 输出: Go语言

	// 修改字符串中的字符
	runes[2] = '编' // 修改第三个字符
	s3 := string(runes)
	fmt.Println(s3) // 输出: Go编言

}

//所以处理中文字符串时，应当使用 []rune 而非 []byte，这样可以正确处理每个 Unicode 字符，避免出现乱码
