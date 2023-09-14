package repository

import (
	"database/sql"
	"github.com/GorbachR/todo-app/api/data/dto"
	"github.com/GorbachR/todo-app/api/data/model"
)

type TodoRepository struct {
	Db *sql.DB
}

func CreateTodoRepository(db *sql.DB) TodoRepository {
	return TodoRepository{Db: db}
}

var selectTodosQuery = "select id, note, active from todo limit ? offset ?;"

func (t *TodoRepository) FindAll(limitAndOffset dto.LimitAndOffset) (todos []model.Todo, err error) {
	rows, err := t.Db.Query(selectTodosQuery, limitAndOffset.Limit, limitAndOffset.Offset)
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

func (t *TodoRepository) FindOne(id int) (todo model.Todo, err error) {

	row := t.Db.QueryRow(selectOneTodoQuery, id)

	err = row.Scan(&todo.Id, &todo.Note, &todo.Active)

	return
}

var insertTodoQuery = "insert into todo (note, active) values (?, ?);"

func (t *TodoRepository) Create(todo model.Todo) (res sql.Result, err error) {
	res, err = t.Db.Exec(insertTodoQuery, todo.Note, todo.Active)
	return
}

var updateTodoQuery = "update todo set note = ?, active = ? where id = ?;"

func (t *TodoRepository) Update(id int, todo model.Todo) (err error) {
	_, err = t.Db.Exec(updateTodoQuery, todo.Note, todo.Active, id)
	return
}

var deleteTodoQuery = "delete from todo where id = ?;"

func (t *TodoRepository) Delete(id int) (err error) {
	_, err = t.Db.Exec(deleteTodoQuery, id)
	return
}
