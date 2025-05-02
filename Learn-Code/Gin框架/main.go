package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建默认的路由引擎
	r := gin.Default()

	// 定义一个GET路由	当访问这个路径时，执行后面的匿名函数。
	// c *gin.Context 是上下文对象，封装了这次请求和对应的响应。
	// c.Request即原生*http.Request 对象
	r.GET("/hello", func(c *gin.Context) {
		// 获取请求方法
		method := c.Request.Method
		//以json格式返回
		c.JSON(http.StatusOK, gin.H{
			//http.StatusOK 是状态码 200
			"method":  method,
			"message": "Hello from Gin",
		})
		//gin.H 是 Gin 框架定义的一个快捷方式，用于创建 map[string]interface{} 类型的对象
	})

	// 启动服务器
	r.Run(":8080")
}
