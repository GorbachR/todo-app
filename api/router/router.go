package router

import (
	"database/sql"

	"github.com/GorbachR/todo-app/api/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {

	r := gin.Default()

	todoController := config.CreateTodoController(db)

	r.GET("/todos", todoController.GetTodos)
	r.GET("/todos/:id", todoController.GetTodo)
	r.POST("/todos", todoController.PostTodo)
	r.PUT("/todos/:id", todoController.PutTodo)
	r.DELETE("/todos/:id", todoController.DeleteTodo)

	return r
}
