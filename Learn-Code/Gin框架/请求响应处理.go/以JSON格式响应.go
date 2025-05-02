func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "这是 JSON 响应",
		})
	})
}