package routes

import (
	"database/sql"
	"github.com/GorbachR/todo-app/api/config"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {

	todoController := config.InitTodoController(db)

	r.GET("/todos", todoController.GetTodos)
	r.GET("/todos/:id", todoController.GetTodo)
	r.POST("/todos", todoController.PostTodo)
	r.PUT("/todos/:id", todoController.PutTodo)
	r.DELETE("/todos/:id", todoController.DeleteTodo)
}
