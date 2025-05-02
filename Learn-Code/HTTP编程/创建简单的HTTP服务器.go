package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 处理/hello路径的请求
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		//http.HandleFunc 是用于注册 某一路径（比如“/hello”） 对应的处理函数( func(){} ) 访问这个路径时会触发后面的函数
		//w 是 http.ResponseWriter 接口，用来向客户端写响应数据
		//r 是 *http.Request，表示客户端发来的请求对象，包含了请求方法、URL、头信息等
		fmt.Fprintf(w, "你好，Gopher！")
		//向响应写入“你好，Gopher！”，这段内容会作为 HTTP 响应体发送给客户端
	})

	// 启动服务器在8080端口
	fmt.Println("服务器启动在 :8080...")
	http.ListenAndServe(":8080", nil)
	//启动 HTTP 服务器并监听 8080 端口
	//第二个参数 nil 表示使用默认的 ServeMux 作为路由器（即用 http.HandleFunc 注册的路径都会被这个默认路由器处理）
}

//浏览器访问 http://localhost:8080/hello
