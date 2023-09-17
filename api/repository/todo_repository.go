package repository

import (
	"database/sql"
	"github.com/GorbachR/todo-app/api/data/dto"
	"github.com/GorbachR/todo-app/api/data/model"
	_ "github.com/go-sql-driver/mysql"
)

type TodoRepository struct {
	Db DBTX
}

type ITodoRepository interface {
	FindAll(dto.LimitAndOffset) ([]model.Todo, error)
	FindOne(int) (model.Todo, error)
	Create(model.Todo) (int, error)
	Update(int, model.Todo) error
	Delete(int) error
}

type DBTX interface {
	Exec(string, ...any) (sql.Result, error)
	Query(string, ...any) (*sql.Rows, error)
	QueryRow(string, ...any) *sql.Row
}

func CreateTodoRepository(db DBTX) TodoRepository {
	return TodoRepository{Db: db}
}

var SelectTodosQuery = "select id, note, active from todo order by id asc limit ? offset ?;"

func (t TodoRepository) FindAll(limitAndOffset dto.LimitAndOffset) (todos []model.Todo, err error) {
	rows, err := t.Db.Query(SelectTodosQuery, limitAndOffset.Limit, limitAndOffset.Offset)
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

var SelectOneTodoQuery = "select id, note, active from todo where id = ?;"

func (t TodoRepository) FindOne(id int) (todo model.Todo, err error) {

	row := t.Db.QueryRow(SelectOneTodoQuery, id)

	err = row.Scan(&todo.Id, &todo.Note, &todo.Active)

	return
}

var InsertTodoQuery = "insert into todo (note, active) values (?, ?);"

func (t TodoRepository) Create(todo model.Todo) (newId int, err error) {
	res, err := t.Db.Exec(InsertTodoQuery, todo.Note, todo.Active)

	if err != nil {
		return
	}

	id, err := res.LastInsertId()

	newId = int(id)

	return
}

var UpdateTodoQuery = "update todo set note = ?, active = ? where id = ?;"

func (t TodoRepository) Update(id int, todo model.Todo) (err error) {
	_, err = t.Db.Exec(UpdateTodoQuery, todo.Note, todo.Active, id)
	return
}

var DeleteTodoQuery = "delete from todo where id = ?;"

func (t TodoRepository) Delete(id int) (err error) {
	_, err = t.Db.Exec(DeleteTodoQuery, id)
	return
}
