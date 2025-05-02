// 当请求进来时，中间件按照注册的顺序依次执行，直到某个中间件调用了c.Next()或路由处理函数被执行。
// 执行路由处理函数后，控制流会返回到最后一个调用c.Next()的中间件，然后逆序执行剩余的中间件代码。
package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(MiddlewareA())
	r.Use(MiddlewareB())
	r.Use(MiddlewareC())

	r.GET("/example", func(c *gin.Context) {
		log.Println("执行路由处理函数")
		c.JSON(200, gin.H{"message": "成功"})
	})

	r.Run(":8080")
}

func MiddlewareA() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("MiddlewareA: 进入")
		c.Next()
		log.Println("MiddlewareA: 退出")
	}
}

func MiddlewareB() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("MiddlewareB: 进入")
		c.Next()
		log.Println("MiddlewareB: 退出")
	}
}

func MiddlewareC() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("MiddlewareC: 进入")
		c.Next()
		log.Println("MiddlewareC: 退出")
	}
}
