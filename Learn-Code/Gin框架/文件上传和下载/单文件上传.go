func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")                      //从请求的表单中提取 name="file" 的上传文件
		c.SaveUploadedFile(file, "uploads/"+file.Filename) //将上传的文件保存到服务器 uploads/ 文件夹下，文件名保持不变。

		c.JSON(200, gin.H{
			"message": "上传成功",
		})
	})
}