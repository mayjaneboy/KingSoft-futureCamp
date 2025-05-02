//有时，你可能不想将 Body 绑定到结构体，而是需要获取原始的、未经处理的字节数据。
// 例如，你需要将请求转发给另一个服务，或者你需要用自定义的方式解析 Body。
//调用 c.GetRawData(): 这个方法会读取请求的 Body 并返回一个 []byte 切片
//请求 Body 通常只能被读取一次。
// 一旦你调用了 GetRawData() 或任何 Bind/ShouldBind 方法，Body 流就会被消耗掉。
// 如果你后续还需要读取 Body（这很少见，但可能发生），需要使用 c.ShouldBindBodyWith 等更高级的方法，或者在读取后将数据存起来。

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// 注意：Go 1.16+ 推荐使用 io.ReadAll
	// import "io" // Go 1.16+
)

func main() {
	router := gin.Default()

	// 处理 /raw POST 请求
	router.POST("/raw", func(c *gin.Context) {
		// 1. 使用 GetRawData 获取原始 Body
		rawData, err := c.GetRawData()
		// 或者使用 Go 标准库的方式 (推荐，更通用)
		// rawData, err := io.ReadAll(c.Request.Body) // Go 1.16+
		// rawData, err := ioutil.ReadAll(c.Request.Body) // Go 1.15及更早版本

		if err != nil {
			log.Printf("读取原始 Body 失败: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "无法读取请求体"})
			return
		}

		// 2. 处理原始数据 (这里只是打印出来)
		log.Printf("接收到原始 Body 数据 (长度: %d):\n%s\n", len(rawData), string(rawData))

		// 3. 你可以根据需要对 rawData 进行操作，比如计算签名、记录日志等

		// 注意：此时 Body 已经被读取，不能再调用 c.ShouldBindJSON 等方法了
		// varjsonData map[string]interface{}
		// if err := c.ShouldBindJSON(&jsonData); err != nil {
		//     log.Printf("尝试再次绑定会失败: %v\n", err) // 这里会报错，因为 Body 已被读取
		// }

		c.JSON(http.StatusOK, gin.H{
			"message":         "成功接收到原始 Body 数据",
			"length":          len(rawData),
			"content_preview": string(rawData[:min(len(rawData), 50)]) + "...", // 显示前50个字符预览
		})
	})

	log.Println("服务器启动，监听端口 :8080")
	router.Run(":8080")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
