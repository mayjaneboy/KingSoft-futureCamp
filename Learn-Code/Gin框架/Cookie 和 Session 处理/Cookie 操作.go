package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 设置 Cookie
	r.GET("/cookie", func(c *gin.Context) {
		// 设置 Cookie
		c.SetCookie("user", "张三", 3600, "/", "localhost", false, true)
		//参数的含义分别是：Cookie 名称、Cookie 的值、有效时间（秒）、	Cookie 所属于的路径、绑定的域名、
		// 是否仅在HTTPS请求中发送该Cookie（false 表示 HTTP 也可发送）、是否只能通过 HTTP 协议访问该 Cookie（JavaScript 无法访问，提高安全性）
		// 读取 Cookie
		user, err := c.Cookie("user")
		if err != nil {
			user = "游客"
		}

		c.String(200, "当前用户：%s", user)
	})
}
