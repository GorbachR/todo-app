package query

import (
	"database/sql"
	"github.com/GorbachR/todo-app/api/data"
	"github.com/GorbachR/todo-app/api/data/schema"
)

var selectTodosQuery = "select id, note, active from todo limit ? offset ?;"

func GetTodos(limit, offset int) (todos []schema.Todo, err error) {
	rows, err := data.Db.Query(selectTodosQuery, limit, offset)
	if err != nil {
		return
	}

	for rows.Next() {
		todo := schema.Todo{}
		if err = rows.Scan(&todo.Id, &todo.Note, &todo.Active); err != nil {
			return
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return
}

var selectOneTodoQuery = "select id, note, active from todo where id = ?;"

func GetTodo(id int) (todo schema.Todo, err error) {

	row := data.Db.QueryRow(selectOneTodoQuery, id)

	err = row.Scan(&todo.Id, &todo.Note, &todo.Active)

	return
}

var insertTodoQuery = "insert into todo (note, active) values (?, ?);"

func CreateTodo(todo schema.Todo) (res sql.Result, err error) {
	res, err = data.Db.Exec(insertTodoQuery, todo.Note, todo.Active)
	return
}

var updateTodoQuery = "update todo set note = ?, active = ?;"

func UpdateTodo(todo schema.Todo) (err error) {
	_, err = data.Db.Exec(updateTodoQuery, todo.Note, todo.Active)
	return
}

var deleteTodoQuery = "delete from todo where id = ?;"

func DeleteTodo(id int) (err error) {
	_, err = data.Db.Exec(deleteTodoQuery, id)
	return
}
