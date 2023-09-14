package routes

import (
	"github.com/GorbachR/todo-app/api/data"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMain(m *testing.M) {

}

func TestGetTodos(t *testing.T) {
	r := gin.New()

	r.GET("/todos", getTodos)
}
