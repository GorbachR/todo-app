package query

import (
	"github.com/GorbachR/todo-app/api/data"
	"github.com/GorbachR/todo-app/api/data/schema"
)

var selectQuery = "select id, note, active from todo limit ? offset ?;"

func Todos(offset int) (todos []schema.Todo, err error) {
	rows, err := data.Db.Query(selectQuery, 5, offset)
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
