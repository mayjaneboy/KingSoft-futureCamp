// 1. BindJSON
// 功能：专门用于将请求体中的JSON数据反序列化并绑定到指定的结构体上。它假定请求的 Content - Type 为 application/json。
// 错误处理：如果绑定失败，BindJSON 会直接调用 c.AbortWithError(400, err).SetType(ErrorTypeBind)，终止请求处理并返回HTTP 400 Bad Request错误，错误信息包含在响应体中。调用者无法在同一处理函数中对错误进行更细粒度的处理。
// 适用场景：适用于简单场景，当你希望在JSON绑定失败时直接返回默认的HTTP 400错误，而不需要在处理函数中显式处理错误时，可以使用 BindJSON。它能减少代码量，使处理函数逻辑更简洁，特别适用于对错误处理要求不高的内部API或快速原型开发。
// 2. ShouldBindJSON
// 功能：同样用于将请求体中的JSON数据反序列化并绑定到指定的结构体上，仅适用于 Content - Type 为 application/json 的请求。
// 错误处理：返回一个错误对象 error。调用者需要显式检查这个错误，并根据错误情况进行相应处理，例如返回不同的HTTP状态码和详细的错误信息给客户端。这种方式使错误处理更灵活。
// 适用场景：适用于复杂的API开发场景，需要对绑定错误进行自定义处理，比如根据不同的错误类型返回不同的HTTP状态码和详细错误描述，帮助客户端更好地理解问题所在。同时，在绑定失败时如需执行额外逻辑（如记录详细错误日志、进行清理操作等），ShouldBindJSON 提供了更大的灵活性。
// 3. ShouldBind
// 功能：具有通用性，它会根据请求的 Content - Type 自动选择合适的绑定器。如果 Content - Type 是 application/json，其行为类似于 ShouldBindJSON；如果是 application/x - www - form - urlencoded 或 multipart/form - data，它会将数据解析为相应的表单格式并绑定到结构体或 map 上。
// 错误处理：返回一个错误对象 error，调用者需显式检查错误并处理。
// 适用场景：适用于不确定请求数据格式，或者希望代码能够自动处理多种常见数据格式（JSON、表单格式等）的场景。例如开发一个既支持传统表单提交又支持JSON数据提交的接口时，ShouldBind 无需针对不同数据格式编写不同绑定逻辑，使用起来非常方便。
// 4. Bind
// 功能：根据请求的 Content - Type 自动选择合适的绑定器，功能与 ShouldBind 类似，能处理多种数据格式。
// 错误处理：如果绑定失败，Bind 会调用 c.AbortWithError(400, err).SetType(ErrorTypeBind)，终止请求处理并返回HTTP 400 Bad Request错误，与 BindJSON 的错误处理方式相同，调用者无法在同一处理函数中对错误进行更细粒度的处理。
// 适用场景：适用于简单场景，当希望在数据绑定失败时直接返回默认HTTP 400错误，且无需在处理函数中显式处理错误，同时又需要处理多种数据格式时，可使用 Bind。

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 定义用户结构体
type User struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}

func main() {
	r := gin.Default()

	r.POST("/bindjson", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.POST("/shouldbindjson", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.POST("/shouldbind", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.POST("/bind", func(c *gin.Context) {
		var user User
		c.Bind(&user)
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.Run(":8080")
}
