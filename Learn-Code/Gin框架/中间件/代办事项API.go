package main

import (
	"uuid"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

var todos = make(map[string]Todo)

func main() {
	r := gin.Default()

	// 获取所有待办事项
	r.GET("/todos", func(c *gin.Context) {
		var todoList []Todo
		for _, todo := range todos {
			todoList = append(todoList, todo)
		}
		c.JSON(200, todoList)
	})

	// 创建待办事项
	r.POST("/todos", func(c *gin.Context) {
		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		todo.ID = uuid.New().String()
		todos[todo.ID] = todo

		c.JSON(201, todo)
	})

	r.Run(":8080")
}
