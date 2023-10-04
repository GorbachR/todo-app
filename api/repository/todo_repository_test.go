package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/GorbachR/todo-app/api/config"
	"github.com/GorbachR/todo-app/api/customError"
	"github.com/GorbachR/todo-app/api/data/dto"
	"github.com/GorbachR/todo-app/api/data/model"
	"github.com/GorbachR/todo-app/api/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
	"time"
)

var db *sql.DB
var migrateStr config.Migrations
var repo *repository.TodoRepository

func TestMain(m *testing.M) {

	var err error
	db, err = config.InitDb("root:dev@/todo_app_test?multiStatements=true&parseTime=true&loc=UTC")
	if err != nil {
		log.Fatal(1)
	}
	migrateStr, err = config.CreateMigrations(db, "file://../data/migration")
	if err != nil {
		log.Fatal(1)
	}

	err = migrateStr.MigrateDown(true)

	if err != nil {
		log.Fatal(1)
	}

	repo = repository.CreateTodoRepository(db)

	os.Exit(m.Run())
}

// Assumption: this function is 100% correct
func insertTodo(dbParam *sql.DB, todo model.Todo) (err error) {
	if todo.LastChanged.IsZero() {
		todo.LastChanged = time.Now()
	}
	_, err = dbParam.ExecContext(context.Background(), repository.InsertTodoQuery, todo.Note, todo.Completed, todo.LastChanged)
	return
}

// Assumption: this function is 100% correct
func getTodo(dbParam *sql.DB, id int) (todo model.Todo, err error) {
	row := dbParam.QueryRowContext(context.Background(), repository.SelectOneTodoQuery, id)
	todo = model.Todo{}
	err = row.Scan(&todo.Id, &todo.Note, &todo.Completed, &todo.Position, &todo.LastChanged)
	return
}

func zeroNanoseconds(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}

func TestFindAll(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	defer migrateStr.MigrateDown(true)
	require.NoError(t, err)

	equalTime := zeroNanoseconds(time.Now().UTC().Round(0))
	for i := 1; i <= 10; i++ {
		todo := model.Todo{Note: fmt.Sprintf("%s %d", "todo", i), Completed: i%2 == 0, LastChanged: equalTime}
		err := insertTodo(db, todo)
		require.NoError(t, err)
	}

	findAllTests := []struct {
		params         dto.GetTodosQueryParams
		expectedRes    []model.Todo
		expectedLength int
	}{
		{
			params: dto.GetTodosQueryParams{
				Limit:  2,
				Offset: 0,
			},
			expectedRes: []model.Todo{
				{Id: 10, Note: "todo 10", Completed: true, Position: 10, LastChanged: equalTime},
				{Id: 9, Note: "todo 9", Completed: false, Position: 9, LastChanged: equalTime},
			},
			expectedLength: 2,
		},
		{
			params: dto.GetTodosQueryParams{
				Limit:  2,
				Offset: 5,
			},
			expectedRes: []model.Todo{
				{Id: 5, Note: "todo 5", Completed: false, Position: 5, LastChanged: equalTime},
				{Id: 4, Note: "todo 4", Completed: true, Position: 4, LastChanged: equalTime},
			},
			expectedLength: 2,
		},
		{
			params: dto.GetTodosQueryParams{
				Limit:  3,
				Offset: 0,
				Q:      "6",
			},
			expectedRes: []model.Todo{
				{Id: 6, Note: "todo 6", Completed: true, Position: 6, LastChanged: equalTime},
			},
			expectedLength: 1,
		},
	}

	for _, test := range findAllTests {
		res, err := repo.FindAll(context.Background(), test.params)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedRes, res)
		assert.Len(t, res, test.expectedLength)
	}
}

func TestFindOne(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	defer migrateStr.MigrateDown(true)
	require.NoError(t, err)

	equalTime := zeroNanoseconds(time.Now().UTC().Round(0))
	for i := 1; i < 10; i++ {
		todo := model.Todo{Note: fmt.Sprintf("%s %d", "todo", i), Completed: i%2 == 0, LastChanged: equalTime}
		err := insertTodo(db, todo)
		require.NoError(t, err)
	}

	findOneTests := []struct {
		params      int
		expectedErr error
		expectedRes model.Todo
	}{
		{
			params:      5,
			expectedErr: nil,
			expectedRes: model.Todo{Id: 5, Note: "todo 5", Completed: false, Position: 5, LastChanged: equalTime},
		},
		{
			params:      999,
			expectedErr: customError.ErrResourceNotFound{Resource: "999"},
			expectedRes: model.Todo{},
		},
	}

	for _, test := range findOneTests {
		res, err := repo.FindOne(context.Background(), test.params)
		if test.expectedErr == nil {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.expectedErr.Error())
		}
		assert.Equal(t, test.expectedRes, res)
	}
}

func TestCreate(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	defer migrateStr.MigrateDown(true)
	require.NoError(t, err)

	createTests := []struct {
		params      model.Todo
		expectedErr error
		expectedRes model.Todo
	}{
		{
			params:      model.Todo{Id: 1, Note: "todo 1", Completed: true, Position: 1},
			expectedErr: nil,
			expectedRes: model.Todo{Id: 1, Note: "todo 1", Completed: true, Position: 1},
		},
		{
			params:      model.Todo{Id: 2, Note: "todo 2", Completed: false, Position: 2},
			expectedErr: nil,
			expectedRes: model.Todo{Id: 2, Note: "todo 2", Completed: false, Position: 2},
		},
	}

	for _, test := range createTests {
		res, err := repo.Create(context.Background(), test.params)
		if test.expectedErr == nil {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.expectedErr.Error())
		}
		assert.Equal(t, test.expectedRes.Id, res.Id)
		assert.Equal(t, test.expectedRes.Note, res.Note)
		assert.Equal(t, test.expectedRes.Position, res.Position)
		assert.Equal(t, test.expectedRes.Completed, res.Completed)
	}
}

func TestUpdate(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	defer migrateStr.MigrateDown(true)
	require.NoError(t, err)

	for i := 0; i < 10; i++ {
		todo := model.Todo{Note: fmt.Sprintf("%s %d", "todo", i), Completed: i%2 == 0}
		err := insertTodo(db, todo)
		require.NoError(t, err)
	}

	updateTests := []struct {
		params1     int
		params2     model.Todo
		expectedErr error
	}{
		{
			params1:     5,
			params2:     model.Todo{Note: "Hello world", Completed: false, Position: 5},
			expectedErr: nil,
		},
		{
			params1:     999,
			params2:     model.Todo{},
			expectedErr: customError.ErrResourceNotFound{Resource: "999"},
		},
	}

	for _, test := range updateTests {
		err := repo.Update(context.Background(), test.params1, test.params2)
		if test.expectedErr == nil {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.expectedErr.Error())
		}
	}
}

func TestDelete(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	defer migrateStr.MigrateDown(true)
	require.NoError(t, err)

	for i := 0; i < 10; i++ {
		todo := model.Todo{Note: fmt.Sprintf("%s %d", "todo", i), Completed: i%2 == 0}
		err := insertTodo(db, todo)
		require.NoError(t, err)
	}

	deleteTests := []struct {
		params      int
		expectedErr error
	}{
		{
			params:      5,
			expectedErr: nil,
		},
		{
			params:      999,
			expectedErr: customError.ErrResourceNotFound{Resource: "999"},
		},
	}

	for _, test := range deleteTests {
		err := repo.Delete(context.Background(), test.params)
		if test.expectedErr == nil {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.expectedErr.Error())
		}
	}
}

func TestReorderInsert(t *testing.T) {

	err := migrateStr.MigrateUp(false)
	defer migrateStr.MigrateDown(true)
	require.NoError(t, err)

	for i := 0; i < 10; i++ {
		todo := model.Todo{Note: fmt.Sprintf("%s %d", "todo", i), Completed: i%2 == 0}
		err := insertTodo(db, todo)
		require.NoError(t, err)
	}

	reorderInsertTests := []struct {
		params      dto.ReorderPosTodoParams
		expectedErr error
		expectedPos int
	}{
		{
			params:      dto.ReorderPosTodoParams{TodoToInsertId: 10, AfterTodoId: 5},
			expectedPos: 6,
			expectedErr: nil,
		},
		{
			params:      dto.ReorderPosTodoParams{TodoToInsertId: 999, AfterTodoId: 5},
			expectedErr: customError.ErrResourceNotFound{Resource: "999"},
		},
		{
			params:      dto.ReorderPosTodoParams{TodoToInsertId: 4, AfterTodoId: 997},
			expectedErr: customError.ErrResourceNotFound{Resource: "997"},
		},
	}

	for _, test := range reorderInsertTests {
		err := repo.ReorderInsert(context.Background(), test.params)
		if test.expectedErr == nil {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.expectedErr.Error())
		}
		todo, err := getTodo(db, test.params.TodoToInsertId)
		if test.expectedPos != 0 {
			assert.NoError(t, err)
			assert.Equal(t, todo.Position, test.expectedPos)
		}
	}
}
