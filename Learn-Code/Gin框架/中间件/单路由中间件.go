func main() {
	r := gin.Default()

	// 对单个路由应用中间件
	r.GET("/admin", AdminRequired(), getAdminPage)

	// 对单个路由应用多个中间件
	r.GET("/restricted", AuthRequired(), RoleCheck("admin"), getRestrictedPage)

	r.Run(":8080")
}

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := checkAdminStatus(c)
		if !isAdmin {
			c.AbortWithStatusJSON(403, gin.H{"error": "需要管理员权限"})
			return
		}
		c.Next()
	}
}

func RoleCheck(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := getUserRole(c)
		if userRole != role {
			c.AbortWithStatusJSON(403, gin.H{"error": "角色权限不足"})
			return
		}
		c.Next()
	}
}