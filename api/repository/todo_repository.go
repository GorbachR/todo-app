package repository

import (
	"database/sql"
	"github.com/GorbachR/todo-app/api/data/model"
)

type TodosInterface interface {
	GetTodos() ([]model.Todo, error)
	GetTodo() (model.Todo, error)
	CreateTodo() (sql.Result, error)
	UpdateTodo() error
	DeleteTodo() error
}

type TodoRepositoryImpl struct {
	db *sql.DB
}

var selectTodosQuery = "select id, note, active from todo limit ? offset ?;"

func (q *Queries) GetTodos(limitAndOffset LimitAndOffset) (todos []model.Todo, err error) {
	rows, err := q.Db.Query(selectTodosQuery, limitAndOffset.Limit, limitAndOffset.Offset)
	if err != nil {
		return
	}

	for rows.Next() {
		todo := model.Todo{}
		if err = rows.Scan(&todo.Id, &todo.Note, &todo.Active); err != nil {
			return
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return
}

var selectOneTodoQuery = "select id, note, active from todo where id = ?;"

func (q *Queries) GetTodo(id int) (todo model.Todo, err error) {

	row := q.Db.QueryRow(selectOneTodoQuery, id)

	err = row.Scan(&todo.Id, &todo.Note, &todo.Active)

	return
}

var insertTodoQuery = "insert into todo (note, active) values (?, ?);"

func (q *Queries) CreateTodo(todo model.Todo) (res sql.Result, err error) {
	res, err = q.Db.Exec(insertTodoQuery, todo.Note, todo.Active)
	return
}

var updateTodoQuery = "update todo set note = ?, active = ? where id = ?;"

func (q *Queries) UpdateTodo(id int, todo model.Todo) (err error) {
	_, err = q.Db.Exec(updateTodoQuery, todo.Note, todo.Active, id)
	return
}

var deleteTodoQuery = "delete from todo where id = ?;"

func (q *Queries) DeleteTodo(id int) (err error) {
	_, err = q.Db.Exec(deleteTodoQuery, id)
	return
}
