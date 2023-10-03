package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/GorbachR/todo-app/api/customError"
	"github.com/GorbachR/todo-app/api/data/dto"
	"github.com/GorbachR/todo-app/api/data/model"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type TodoRepository struct {
	Db DBTX
}

type ITodoRepository interface {
	FindAll(context.Context, dto.GetTodosQueryParams) ([]model.Todo, error)
	FindOne(context.Context, int) (model.Todo, error)
	Create(context.Context, model.Todo) (model.Todo, error)
	Update(context.Context, int, model.Todo) error
	Delete(context.Context, int) error
	ReorderInsert(context.Context, dto.ReorderPosTodoParams) error
}

type DBTX interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

func CreateTodoRepository(db DBTX) *TodoRepository {
	return &TodoRepository{Db: db}
}

var SelectTodosQuery = "select id, note, completed, position, lastChanged from todo where note like ? order by position desc limit ? offset ?;"

func (t TodoRepository) FindAll(ctx context.Context, getTodosQueryParams dto.GetTodosQueryParams) (todos []model.Todo, err error) {
	todos = []model.Todo{}
	if getTodosQueryParams.Q == "" {
		getTodosQueryParams.Q = "%"
	} else {
		getTodosQueryParams.Q = fmt.Sprintf("%s%s%s", "%", getTodosQueryParams.Q, "%")
	}
	rows, err := t.Db.QueryContext(ctx, SelectTodosQuery, getTodosQueryParams.Q, getTodosQueryParams.Limit, getTodosQueryParams.Offset)
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		todo := model.Todo{}
		if err = rows.Scan(&todo.Id, &todo.Note, &todo.Completed, &todo.Position, &todo.LastChanged); err != nil {
			return
		}
		todos = append(todos, todo)
	}
	return
}

var SelectOneTodoQuery = "select id, note, completed, position, lastChanged from todo where id = ?;"

func (t TodoRepository) FindOne(ctx context.Context, id int) (todo model.Todo, err error) {

	row := t.Db.QueryRowContext(ctx, SelectOneTodoQuery, id)

	err = row.Scan(&todo.Id, &todo.Note, &todo.Completed, &todo.Position, &todo.LastChanged)
	if errors.Is(err, sql.ErrNoRows) {
		err = customError.ErrResourceNotFound{Resource: strconv.Itoa(id)}
	}
	return
}

var InsertTodoQuery = "insert into todo (note, completed, lastChanged) values (?, ?, ?);"

func (t TodoRepository) Create(ctx context.Context, inputTodo model.Todo) (todo model.Todo, err error) {
	timeNow := time.Now().UTC().Round(0)
	res, err := t.Db.ExecContext(ctx, InsertTodoQuery, inputTodo.Note, inputTodo.Completed, timeNow)

	if err != nil {
		return
	}

	id, err := res.LastInsertId()

	if err != nil {
		return
	}

	todo, err = t.FindOne(ctx, int(id))
	if errors.Is(err, sql.ErrNoRows) {
		err = customError.ErrResourceNotFound{}
	}

	return
}

var UpdateTodoQuery = "update todo set note = ?, completed = ?, lastChanged = ? where id = ?;"

func (t TodoRepository) Update(ctx context.Context, id int, todo model.Todo) (err error) {
	timeNow := time.Now().UTC().Round(0)
	res, err := t.Db.ExecContext(ctx, UpdateTodoQuery, todo.Note, todo.Completed, timeNow, id)

	if err != nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 {
		err = customError.ErrResourceNotFound{}
	}
	return
}

var DeleteTodoQuery = "delete from todo where id = ?;"

func (t TodoRepository) Delete(ctx context.Context, id int) (err error) {
	res, err := t.Db.ExecContext(ctx, DeleteTodoQuery, id)
	if err != nil {
		return
	}
	rowsAffected64, err := res.RowsAffected()
	if rowsAffected64 == 0 {
		err = customError.ErrResourceNotFound{Resource: strconv.Itoa(id)}
	}
	return
}

var UpdateTodoOrder = "update todo set position = ? where id = ?;"
var ReorderTodoOrder = "update todo set position = position + 1 where position > ?;"

func (t TodoRepository) ReorderInsert(ctx context.Context, params dto.ReorderPosTodoParams) (err error) {

	var todo model.Todo
	if params.AfterTodoId != 0 {
		todo, err = t.FindOne(ctx, params.AfterTodoId)
	}

	if errors.Is(err, sql.ErrNoRows) {
		err = customError.ErrResourceNotFound{Resource: strconv.Itoa(params.AfterTodoId)}
		return
	} else if err != nil {
		return
	}

	_, err = t.Db.ExecContext(ctx, ReorderTodoOrder, todo.Position)

	if err != nil {
		return
	}

	res, err := t.Db.ExecContext(ctx, UpdateTodoOrder, todo.Position+1, params.TodoToInsertId)

	if err != nil {
		return
	}
	rowsAffected64, err := res.RowsAffected()

	if err != nil {
		return
	}

	if rowsAffected64 == 0 {
		err = customError.ErrResourceUnchanged{Resource: strconv.Itoa(params.TodoToInsertId)}
	}
	return
}
