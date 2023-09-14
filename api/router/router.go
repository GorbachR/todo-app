package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/todos", getTodos)
	r.POST("/todo", postTodo)
	r.GET("/todo/:id", getTodo)
	r.PUT("/todo/:id", putTodo)
	r.DELETE("/todo/:id", deleteTodo)
}
