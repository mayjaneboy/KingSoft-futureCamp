package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string  `json:"username" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	Age      *int    `json:"age"`      // 使用指针，可以为null
	Nickname *string `json:"nickname"` // 可以为null
	VIP      *bool   `json:"vip"`      // 可以为null
}

func main() {
	router := gin.Default()

	router.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 检查Age是否为空
		if user.Age == nil {
			log.Println("年龄字段为空")
		} else {
			log.Printf("年龄: %d\n", *user.Age)
		}

		// 检查Nickname是否为空
		if user.Nickname == nil {
			log.Println("昵称字段为空")
		} else {
			log.Printf("昵称: %s\n", *user.Nickname)
		}

		c.JSON(http.StatusOK, gin.H{"message": "用户数据接收成功", "user": user})
	})

	router.Run(":8080")
}
