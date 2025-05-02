package main

import (
	"log"
	"net/http"
)

func main() {
	// 创建一个处理器（Handler），专门用于提供静态文件内容（例如 HTML、JS、CSS、图片等）
	fileServer := http.FileServer(http.Dir("."))
	//http.Dir(".") 表示用当前程序所在目录作为“网站根目录”

	http.Handle("/", fileServer)
	//将 / 路径注册给 fileServer	所有以 / 开头的请求路径（即所有路径）都会交给 fileServer 处理

	log.Println("文件服务器启动在 :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	//启动 HTTP 服务器并监听 8080 端口
	//log.Fatal(...)：一旦 ListenAndServe 返回错误，会将错误信息打印并退出程序。
}

//浏览器访问 http://localhost:8080/index.html 流程：
//1、浏览器发起请求：	GET /index.html HTTP/1.1	Host: localhost:8080
//2、Go 服务监听了 8080 端口
//3、匹配路由 / 并由 fileServer 处理
//4、fileServer 根据请求路径找本地文件
//5、找到文件并返回内容
