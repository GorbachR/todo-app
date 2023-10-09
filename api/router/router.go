package router

import (
	"database/sql"
	"net/http"
    "github.com/gin-contrib/cors"
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

// 	r.Use(cors.New(cors.Config{
//        AllowOrigins: []string{"http://localhost:5173/"},
//        AllowMethods: []string{http.MethodOptions,http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch},
//        AllowHeaders: []string{"Origin"},
//        ExposeHeaders: []string{"Content-Length"},
//        AllowCredentials: true,
//        MaxAge: 12 * time.Hour,
// 	}))

    corsConfig := cors.DefaultConfig()

    corsConfig.AllowOrigins = []string{"*"}
    corsConfig.AllowCredentials = true
    corsConfig.AddAllowMethods(http.MethodOptions)
    r.Use(cors.New(corsConfig))

	todoController := config.CreateTodoController(db)

	r.GET(GetTodosRoute, todoController.GetTodos)
	r.GET(GetTodoRoute, todoController.GetTodo)
	r.POST(PostTodoRoute, todoController.PostTodo)
	r.PUT(PutTodoRoute, todoController.PutTodo)
	r.DELETE(DeleteTodoRoute, todoController.DeleteTodo)
	r.PATCH(PatchReorderTodo, todoController.ReorderTodo)

	return r
}
