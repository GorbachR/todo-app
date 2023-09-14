package service

import (
	"github.com/GorbachR/todo-app/api/data/dto"
	"github.com/GorbachR/todo-app/api/data/model"
	"github.com/GorbachR/todo-app/api/repository"
)

type TodoService struct {
	TodoRepository repository.TodoRepository
}

func CreateTodoService(t repository.TodoRepository) TodoService {
	return TodoService{TodoRepository: t}
}

func (t *TodoService) FindAll(limitAndOffset dto.LimitAndOffset) (todos []model.Todo, err error) {
	todos, err = t.TodoRepository.FindAll(limitAndOffset)
	return
}

func (t *TodoService) FindOne(id int) (todo model.Todo, err error) {
	todo, err = t.TodoRepository.FindOne(id)
	return
}

func (t *TodoService) Create(newTodo model.Todo) (newId int, err error) {

	res, err := t.TodoRepository.Create(newTodo)

	if err != nil {
		return
	}

	lastInsert, err := res.LastInsertId()

	if err != nil {
		return
	}

	newId = int(lastInsert)
	return
}

func (t *TodoService) Update(id int, updatedTodo model.Todo) (err error) {
	err = t.TodoRepository.Update(id, updatedTodo)
	return
}

func (t *TodoService) Delete(id int) (err error) {
	err = t.TodoRepository.Delete(id)
	return
}
