package config

import (
	"database/sql"
	"github.com/GorbachR/todo-app/api/controller"
	"github.com/GorbachR/todo-app/api/repository"
	"github.com/GorbachR/todo-app/api/service"
	"github.com/google/wire"
)

func initTodoController(db *sql.DB) controller.TodoController {
	wire.Build(repository.CreateTodoRepository, service.CreateTodoService, controller.CreateTodoController)
	return controller.TodoController{}
}
