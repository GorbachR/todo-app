package repository_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/GorbachR/todo-app/api/config"
	"github.com/GorbachR/todo-app/api/data/dto"
	"github.com/GorbachR/todo-app/api/data/model"
	"github.com/GorbachR/todo-app/api/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var db *sql.DB
var migrateStr config.Migrations

func TestMain(m *testing.M) {

	db, err := config.InitDb("root:dev@/todo_app_test?multiStatements=true")
	if err != nil {
		log.Fatal(1)
	}
	migrateStr, err = config.CreateMigrations(db, "file://../data/migration")
	if err != nil {
		log.Fatal(1)
	}
	// err = migrateStr.Migrate.Drop()
	//
	// if err != nil {
	// 	log.Fatal(1)
	// }

	os.Exit(m.Run())
}

// Assumption: this function is 100% correct
func insertTodo(dbParam *sql.DB, todo model.Todo) (err error) {
	_, err = dbParam.Exec(repository.InsertTodoQuery, todo.Note, todo.Active)
	return
}

// Assumption: this function is 100% correct
func getTodo(dbParam *sql.DB, id int) (todo model.Todo, err error) {
	row := dbParam.QueryRow(repository.SelectOneTodoQuery, id)
	todo = model.Todo{}
	err = row.Scan(&todo)
	return
}

// Assumption: this function is 100% correct
func deleteTodos(db *sql.DB) (err error) {
	_, err = db.Exec("delete from todo;")
	return
}

func TestFindAll(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	require.NoError(t, err)

	repo := repository.CreateTodoRepository(db)

	for i := 1; i <= 100; i++ {
		todo := model.Todo{Note: fmt.Sprintf("%s %d", "todo", i), Active: i%2 == 0}
		insertTodo(db, todo)
	}

	limitAndOffset := dto.LimitAndOffset{Limit: 2, Offset: 0}
	expectedTodos := []model.Todo{
		{Id: 1, Note: "todo 1", Active: true},
		{Id: 2, Note: "todo 2", Active: false},
	}

	todos, err := repo.FindAll(limitAndOffset)

	assert.NoError(t, err)
	assert.Equal(t, expectedTodos, todos)
	assert.Len(t, todos, 2)

	limitAndOffset = dto.LimitAndOffset{Limit: 2, Offset: 14}
	expectedTodos = []model.Todo{
		{Id: 15, Note: "todo 15", Active: true},
		{Id: 16, Note: "todo 16", Active: false},
	}

	todos, err = repo.FindAll(limitAndOffset)

	assert.NoError(t, err)
	assert.Equal(t, expectedTodos, todos)
	assert.Len(t, todos, 2)

	err = migrateStr.MigrateDown(false)
	assert.NoError(t, err)
}

func TestFindOne(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	require.NoError(t, err)

	repo := repository.CreateTodoRepository(db)

	for i := 1; i < 10; i++ {
		todo := model.Todo{Note: fmt.Sprintf("%s %d", "todo", i), Active: i%2 == 0}
		insertTodo(db, todo)
	}

	expectedTodo := model.Todo{Note: "todo 5", Active: false}
	id := 5

	todo, err := repo.FindOne(id)

	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, todo)

	expectedTodo = model.Todo{}
	id = 999

	todo, err = repo.FindOne(id)

	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, todo)

	err = migrateStr.MigrateDown(false)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	require.NoError(t, err)

	repo := repository.CreateTodoRepository(db)

	todo := model.Todo{Note: "todo 1", Active: true}
	expectedId := 1

	id, err := repo.Create(todo)

	assert.NoError(t, err)
	assert.Equal(t, expectedId, id)

	insertedTodo, err := getTodo(db, expectedId)

	assert.NoError(t, err)
	assert.Equal(t, todo, insertedTodo)

	todo = model.Todo{Note: "todo 2", Active: false}
	expectedId = 1

	id, err = repo.Create(todo)

	assert.NoError(t, err)
	assert.Equal(t, expectedId, id)

	insertedTodo, err = getTodo(db, expectedId)

	assert.NoError(t, err)
	assert.Equal(t, todo, insertedTodo)

	err = migrateStr.MigrateDown(false)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	require.NoError(t, err)

	repo := repository.CreateTodoRepository(db)

	for i := 0; i < 100; i++ {
		todo := model.Todo{Note: fmt.Sprintf("%s %d", "todo", i), Active: i%2 == 0}
		insertTodo(db, todo)
	}

	expectedTodo := model.Todo{Note: "Hello world", Active: false}
	expectedId := 5

	err = repo.Update(expectedId, expectedTodo)

	assert.NoError(t, err)

	todo, err := getTodo(db, expectedId)

	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, todo)

	expectedTodo = model.Todo{}
	expectedId = 999

	err = repo.Update(expectedId, expectedTodo)

	assert.NoError(t, err)

	todo, err = getTodo(db, expectedId)

	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, todo)

	err = migrateStr.MigrateDown(false)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	require.NoError(t, err)

	repo := repository.CreateTodoRepository(db)

	for i := 0; i < 100; i++ {
		todo := model.Todo{Note: fmt.Sprintf("%s %d", "todo", i), Active: i%2 == 0}
		insertTodo(db, todo)
	}

	expectedId := 5
	expectedTodo := model.Todo{}

	err = repo.Delete(expectedId)

	assert.NoError(t, err)

	todo, err := getTodo(db, expectedId)

	assert.NoError(t, err)
	assert.Equal(t, expectedTodo, todo)

	err = migrateStr.MigrateDown(false)
	assert.NoError(t, err)
}
