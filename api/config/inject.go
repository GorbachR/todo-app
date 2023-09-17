package config

import (
	"database/sql"
	"github.com/GorbachR/todo-app/api/controller"
	"github.com/GorbachR/todo-app/api/repository"
	"github.com/GorbachR/todo-app/api/service"
)

func CreateTodoController(db *sql.DB) controller.TodoController {
	repo := repository.CreateTodoRepository(db)
	serv := service.CreateTodoService(repo)
	return controller.CreateTodoController(serv)
}
