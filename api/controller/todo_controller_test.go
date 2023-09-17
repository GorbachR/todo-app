package controller_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GorbachR/todo-app/api/config"
	"github.com/GorbachR/todo-app/api/controller"
	"github.com/GorbachR/todo-app/api/data/model"
	"github.com/GorbachR/todo-app/api/router"
	"github.com/GorbachR/todo-app/api/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine

type MockTodoService struct {
	service.TodoService
}

// func (m *MockTodoService) FindAll() (todos []model.Todo, err error) {
//
// }

func TestMain(m *testing.M) {

	db, _, err := sqlmock.New()

	if err != nil {
		os.Exit(1)
	}

	defer db.Close()

	r = gin.Default()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestGetTodos(t *testing.T) {

	w := httptest.NewRecorder()
	r := gin.Default()
	router.SetupRouter(db)

	req, _ := http.NewRequest("GET", "/todos", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
