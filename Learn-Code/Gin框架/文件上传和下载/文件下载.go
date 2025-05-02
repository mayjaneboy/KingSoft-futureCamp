func main() {
	r := gin.Default()

	r.GET("/download", func(c *gin.Context) {
		// 指定下载文件
		filename := "test.txt"
		c.Header("Content-Type", "application/octet-stream")
		//设置响应的 MIME 类型为 application/octet-stream，表示 二进制流，强制浏览器下载文件，而不是解析它。
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		//告诉浏览器：这是一个附件（attachment），下载时的默认文件名为 filename。
		c.File("./files/" + filename)
		//告诉 Gin：从服务器的 ./files/ 目录中读取该文件并发送到客户端。
	})
}