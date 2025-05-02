package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 创建一个需要认证的API路由组
	authorized := r.Group("/api")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/users", listUsers)
		authorized.POST("/users", createUser)
	}

	// 创建一个不需要认证的公开路由组
	public := r.Group("/public")
	{
		public.GET("/login", loginEndpoint)
		public.GET("/register", registerEndpoint)
	}

	r.Run(":8080")
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "未授权访问"})
			return
		}

		// 这里可以添加更复杂的令牌验证逻辑
		if token != "valid-token" {
			c.AbortWithStatusJSON(401, gin.H{"error": "无效的令牌"})
			return
		}

		c.Next()
	}
}
