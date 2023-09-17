package repository

import (
	"github.com/GorbachR/todo-app/api/data/dto"
	"github.com/GorbachR/todo-app/api/data/model"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func CreateMockTodoRepository() MockTodoRepository {
	return MockTodoRepository{}
}

func (m *MockTodoRepository) FindAll(LimitAndOffset dto.LimitAndOffset) ([]model.Todo, error) {
	mockArgs := m.Called(LimitAndOffset)
	return mockArgs.Get(0).([]model.Todo), mockArgs.Error(1)
}

func (m *MockTodoRepository) FindOne(todoId int) (model.Todo, error) {
	mockArgs := m.Called(todoId)
	return mockArgs.Get(0).(model.Todo), mockArgs.Error(1)
}

func (m *MockTodoRepository) Create(todo model.Todo) (int, error) {
	mockArgs := m.Called(todo)
	return mockArgs.Int(0), mockArgs.Error(1)
}

func (m *MockTodoRepository) Update(todoId int, todo model.Todo) error {
	mockArgs := m.Called(todoId, todo)
	return mockArgs.Error(0)
}

func (m *MockTodoRepository) Delete(todoId int) error {
	mockArgs := m.Called(todoId)
	return mockArgs.Error(0)
}
