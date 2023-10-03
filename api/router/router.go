package router

import (
	"database/sql"

	"github.com/GorbachR/todo-app/api/config"
	"github.com/gin-gonic/gin"
)

const (
	GetTodosRoute    = "/todos"
	GetTodoRoute     = "/todos/:id"
	PostTodoRoute    = "/todos"
	PutTodoRoute     = "/todos/:id"
	DeleteTodoRoute  = "/todos/:id"
	PatchReorderTodo = "/todos/reorder"
)

func SetupRouter(db *sql.DB) *gin.Engine {

	r := gin.Default()

	todoController := config.CreateTodoController(db)

	r.GET(GetTodosRoute, todoController.GetTodos)
	r.GET(GetTodoRoute, todoController.GetTodo)
	r.POST(PostTodoRoute, todoController.PostTodo)
	r.PUT(PutTodoRoute, todoController.PutTodo)
	r.DELETE(DeleteTodoRoute, todoController.DeleteTodo)
	r.PATCH(PatchReorderTodo, todoController.ReorderTodo)

	return r
}
