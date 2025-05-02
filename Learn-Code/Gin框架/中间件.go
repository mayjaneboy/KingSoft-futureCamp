package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 模拟用户信息结构体
type User struct {
	ID   int
	Name string
}

// 模拟考试信息结构体
type Exam struct {
	ID    int
	Title string
}

// 模拟数据库存储用户信息
var users = map[int]User{
	1: {ID: 1, Name: "Alice"},
	2: {ID: 2, Name: "Bob"},
}

// 模拟数据库存储考试信息，这里假设每个用户有固定的考试列表
var exams = map[int][]Exam{
	1: {
		{ID: 101, Title: "Math Exam"},
		{ID: 102, Title: "English Exam"},
	},
	2: {
		{ID: 201, Title: "Science Exam"},
		{ID: 202, Title: "History Exam"},
	},
}

// 获取用户信息中间件
func getUserInfo() gin.HandlerFunc {
	//gin.HandlerFunc是 Gin 框架中用于定义 中间件函数或处理函数 的函数类型。本质上是func(c *gin.Context)
	//也就是说，getUserInfo()返回一个这样类型的中间件函数，这个中间件可以被 Gin 路由注册使用。
	return func(c *gin.Context) {
		userID := 1 // 简单模拟获取用户ID，实际应用中可能从认证等方式获取
		user, exists := users[userID]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			//http.StatusNotFound是404
			c.Abort()
			//用 c.Abort() 方法终止请求的处理流程，并返回相应的错误响应
			//后续的中间件（包括主处理函数）将 不会被执行
			return
		}
		c.Set("user", user)
		//c.Set("user", userInfo) 用于在上下文中设置数据
		//可以用value, exists := c.Get("user") 从上下文中获取相应数据
		c.Next()
		//在当前中间件中执行后续的处理函数
	}
}

// 获取用户考试列表中间件
func getUserExams() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user") //从上下文中获取相应数据
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not set in context"})
			//http.StatusBadRequest表示HTTP400错误：客户端请求无效，服务器无法理解。
			c.Abort()
			return
		}
		userObj, ok := user.(User)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data in context"})
			c.Abort()
			return
		}
		userExams, exists := exams[userObj.ID]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Exams not found for this user"})
			c.Abort()
			return
		}
		// 将考试列表存储到上下文中
		c.Set("exams", userExams)
		c.Next()
	}
}

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()

	r.GET("/user/exams", getUserInfo(), getUserExams(), func(c *gin.Context) {
		// 从上下文中获取考试列表
		exams, _ := c.Get("exams")
		c.JSON(http.StatusOK, gin.H{"exams": exams})
	})

	r.Run(":8080")
}
