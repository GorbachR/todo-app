package service

import (
	"context"
	"github.com/GorbachR/todo-app/api/data/dto"
	"github.com/GorbachR/todo-app/api/data/model"
	"github.com/GorbachR/todo-app/api/repository"
)

type TodoService struct {
	TodoRepository repository.ITodoRepository
}

type ITodoService interface {
	FindAll(context.Context, dto.GetTodosQueryParams) ([]model.Todo, error)
	FindOne(context.Context, int) (model.Todo, error)
	Create(context.Context, model.Todo) (todo model.Todo, err error)
	Update(context.Context, int, model.Todo) error
	Delete(context.Context, int) error
	ReorderInsert(context.Context, dto.ReorderPosTodoParams) error
}

func CreateTodoService(t repository.ITodoRepository) *TodoService {
	return &TodoService{TodoRepository: t}
}

func (t TodoService) FindAll(ctx context.Context, getTodosQueryParams dto.GetTodosQueryParams) (todos []model.Todo, err error) {
	todos, err = t.TodoRepository.FindAll(ctx, getTodosQueryParams)
	return
}

func (t TodoService) FindOne(ctx context.Context, id int) (todo model.Todo, err error) {
	todo, err = t.TodoRepository.FindOne(ctx, id)
	return
}

func (t TodoService) Create(ctx context.Context, newTodo model.Todo) (todo model.Todo, err error) {
	todo, err = t.TodoRepository.Create(ctx, newTodo)
	return
}

func (t TodoService) Update(ctx context.Context, id int, updatedTodo model.Todo) (err error) {
	err = t.TodoRepository.Update(ctx, id, updatedTodo)
	return
}

func (t TodoService) Delete(ctx context.Context, id int) (err error) {
	err = t.TodoRepository.Delete(ctx, id)
	return
}

func (t TodoService) ReorderInsert(ctx context.Context, params dto.ReorderPosTodoParams) (err error) {
	err = t.TodoRepository.ReorderInsert(ctx, params)
	return
}
