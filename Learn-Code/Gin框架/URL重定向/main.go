func main() {
	r := gin.Default()

	// HTTP 重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	// 路由重定向
	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test3"
		r.HandleContext(c) //强制重新执行上下文 c 对应的路由处理逻辑（包括中间件和最终的处理函数）
	})

	r.GET("/test3", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
}