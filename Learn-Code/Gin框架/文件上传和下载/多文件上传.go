func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm() //从请求中获取 multipart 表单对象
		files := form.File["files"]  //从表单中提取名为 "files" 的文件数组

		for _, file := range files {
			// 保存文件
			dst := "uploads/" + file.Filename
			c.SaveUploadedFile(file, dst)
		}

		c.JSON(200, gin.H{
			"message": "上传成功",
		})
	})
}