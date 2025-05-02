package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//基本路由匹配方式
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})
	//参数化路由匹配
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})
	//路径中的:name是命名参数类型（有冒号：前缀定义），可以匹配任何值

	r.GET("/files/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath") // 将包含"/path/to/file.txt"
		c.String(http.StatusOK, "File Path: %s", filepath)
	})
	//*filepath是通配符参数，能匹配所有内容（包括/）

	//路由组
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/user", func(c *gin.Context) {
				c.String(http.StatusOK, "API v1 - User")
			})
			v1.GET("/product", func(c *gin.Context) {
				c.String(http.StatusOK, "API v1 - Product")
			})
		}
		v2 := api.Group("/v2")
		{
			v2.GET("/user", func(c *gin.Context) {
				c.String(http.StatusOK, "API v2 - User")
			})
			v2.GET("/product", func(c *gin.Context) {
				c.String(http.StatusOK, "API v2 - Product")
			})
		}
	}
	//首先创建了一个路由组 api，其前缀为 /api。
	// 然后在 api 组下又创建了两个子路由组 v1 和 v2，分别对应/api/v1和/api/v2前缀。
	// 每个子路由组下定义了不同的路由。
	// 这样可以方便地对不同版本的 API 或相关功能的路由进行组织和管理。
	// 同时，路由组还可以共享中间件，比如
	//api := r.Group("/api", authMiddleware)
	//这里的 authMiddleware 中间件会应用到 api 组下的所有路由
	r.Run(":8080")
}
