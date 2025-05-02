// 这段代码是一个使用 Gin 框架 + gin-contrib/sessions 中间件实现 Session 会话管理 的示例程序。
// 它实现了一个 /session 路由，用于统计用户访问次数。
// Session（会话）是 Web 开发中用于 在多个请求之间保持用户状态 的一种机制。
// 由于 HTTP 是无状态协议，每次请求服务器都不会自动记住用户的信息，而 Session 就是用来“记住你是谁”的。
package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie" //使用 Cookie 存储 Session 的实现
	"github.com/gin-gonic/gin"               //Gin 的 Session 支持中间件
)

func main() {
	r := gin.Default()

	// 创建一个使用 Cookie 存储 Session 的中间件 Store
	//[]byte("secret"): 是加密 Session 数据用的密钥
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	//将 Session 中间件应用到整个路由。
	//"mysession" 是 Cookie 的名字，也就是用户浏览器里看到的键名

	r.GET("/session", func(c *gin.Context) {
		session := sessions.Default(c) //获取当前请求的 Session 对象

		// 获取访问次数
		count := session.Get("count")
		if count == nil {
			count = 0
		}

		// 更新访问次数
		session.Set("count", count.(int)+1) //将访问次数加 1，然后写回 Session 中。
		session.Save()                      //Save() 表示保存更改的 Session 数据到客户端 Cookie 中。

		c.JSON(200, gin.H{
			"访问次数": count,
		})
	})
}
