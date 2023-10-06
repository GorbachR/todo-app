package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/GorbachR/todo-app/api/controller"
	"github.com/GorbachR/todo-app/api/controller/response"
	"github.com/GorbachR/todo-app/api/customError"
	"github.com/GorbachR/todo-app/api/data/dto"
	"github.com/GorbachR/todo-app/api/data/model"
	"github.com/GorbachR/todo-app/api/router"
	mock_service "github.com/GorbachR/todo-app/api/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestGetTodos(t *testing.T) {

	r := gin.Default()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTodoService := mock_service.NewMockITodoService(ctrl)
	todoController := controller.CreateTodoController(mockTodoService)
	r.GET(router.GetTodosRoute, todoController.GetTodos)

	mockedTodos1 := []model.Todo{
		{Id: 2, Note: "fdsfsd", Completed: true, Position: 29},
		{Id: 4, Note: "fdsfsd", Completed: false, Position: 14},
	}

	type wrongParams struct {
		Q      string
		Limit  string
		Offset string
	}

	testsGet := []struct {
		params       any
		noMock       bool
		mockedTodos  []model.Todo
		mockedErr    error
		expectedCode int
		expectedRes  response.Response
	}{
		{
			params: dto.GetTodosQueryParams{
				Limit: 5,
			},
			mockedTodos:  mockedTodos1,
			mockedErr:    nil,
			expectedCode: 200,
			expectedRes: response.Response{
				Status: "success",
				Data:   mockedTodos1,
			},
		},
		{
			params: dto.GetTodosQueryParams{
				Limit:  2,
				Offset: 5,
				Q:      "test",
			},
			mockedTodos:  mockedTodos1,
			mockedErr:    mysql.ErrInvalidConn,
			expectedCode: 500,
			expectedRes: response.Response{
				Status:  "error",
				Message: "Internal Server Error",
			},
		},
		{
			params:       wrongParams{Q: "fdfd", Limit: "fdfd", Offset: "wrongDataType"},
			noMock:       true,
			expectedCode: 400,
			expectedRes: response.Response{
				Status: "fail",
				Data:   response.BindingFailureData{BindingFailure: "Binding the query payload to the corresponding datatype failed, check your inputs!"},
			},
		},
	}

	for _, test := range testsGet {
		w := httptest.NewRecorder()

		callUrl, err := url.Parse(router.GetTodosRoute)
		require.NoError(t, err)

		queries := callUrl.Query()
		switch v := test.params.(type) {
		case dto.GetTodosQueryParams:
			queries.Add("q", v.Q)
			queries.Add("limit", strconv.Itoa(v.Limit))
			queries.Add("offset", strconv.Itoa(v.Offset))
		case wrongParams:
			queries.Add("q", v.Q)
			queries.Add("limit", v.Limit)
			queries.Add("offset", v.Offset)
		}

		callUrl.RawQuery = queries.Encode()

		req := httptest.NewRequest(http.MethodGet, callUrl.String(), nil)

		if !test.noMock {
			mockTodoService.EXPECT().
				FindAll(gomock.Any(), gomock.Eq(test.params)).
				Return(test.mockedTodos, test.mockedErr)
		}
		r.ServeHTTP(w, req)

		expectedJsonBytes, err := json.Marshal(test.expectedRes)
		assert.NoError(t, err)
		expectedJson := string(expectedJsonBytes)
		assert.Equal(t, test.expectedCode, w.Code)
		assert.Equal(t, expectedJson, string(w.Body.Bytes()))
	}
}

func TestGetTodo(t *testing.T) {
	r := gin.Default()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTodoService := mock_service.NewMockITodoService(ctrl)
	todoController := controller.CreateTodoController(mockTodoService)
	r.GET(router.GetTodoRoute, todoController.GetTodo)

	mockedTodo1 := model.Todo{Id: 2, Note: "fdsfsd", Completed: true, Position: 29}

	testsGet := []struct {
		inputId      any
		noMock       bool
		mockedTodo   model.Todo
		mockedErr    error
		expectedCode int
		expectedRes  response.Response
	}{
		{
			inputId:      2,
			mockedTodo:   mockedTodo1,
			mockedErr:    nil,
			expectedCode: 200,
			expectedRes: response.Response{
				Status: "success",
				Data:   mockedTodo1,
			},
		},
		{
			inputId:      999,
			mockedTodo:   model.Todo{},
			mockedErr:    customError.ErrResourceNotFound{Resource: "999"},
			expectedCode: 404,
			expectedRes: response.Response{
				Status: "fail",
				Data:   response.ValidationFailureData{Key: "999", Details: "Resource 999 doesn't exist"},
			},
		},
		{
			inputId:      "notAnId",
			noMock:       true,
			expectedCode: 400,
			expectedRes: response.Response{
				Status: "fail",
				Data:   response.BindingFailureData{BindingFailure: "Uri params binding failed please stick to the appropriate format /todos/:id"},
			},
		},
	}

	for _, test := range testsGet {
		w := httptest.NewRecorder()

		var param string

		switch v := test.inputId.(type) {
		case int:
			param = strconv.Itoa(v)
		case string:
			param = v
		}
		req := httptest.NewRequest(http.MethodGet, strings.Replace(router.GetTodoRoute, ":id", param, 1), nil)

		if !test.noMock {
			mockTodoService.EXPECT().
				FindOne(gomock.Any(), gomock.Eq(test.inputId)).
				Return(test.mockedTodo, test.mockedErr)
		}
		r.ServeHTTP(w, req)

		expectedJsonBytes, err := json.Marshal(test.expectedRes)
		assert.NoError(t, err)
		expectedJson := string(expectedJsonBytes)
		assert.Equal(t, test.expectedCode, w.Code)
		assert.Equal(t, expectedJson, string(w.Body.Bytes()))
	}
}

func TestPostTodo(t *testing.T) {
	r := gin.Default()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTodoService := mock_service.NewMockITodoService(ctrl)
	todoController := controller.CreateTodoController(mockTodoService)
	r.POST(router.PostTodoRoute, todoController.PostTodo)

	mockedTodo1 := model.Todo{Id: 2, Note: "fdsfsd", Completed: true, Position: 29}
	mockedTodo2 := model.Todo{}

	testsPost := []struct {
		params       any
		noMock       bool
		mockedTodo   model.Todo
		mockedErr    error
		expectedCode int
		expectedRes  response.Response
	}{
		{
			params:       mockedTodo1,
			mockedTodo:   mockedTodo1,
			mockedErr:    nil,
			expectedCode: 201,
			expectedRes: response.Response{
				Status: "success",
				Data:   mockedTodo1,
			},
		},
		{
			params:       mockedTodo2,
			noMock:       true,
			expectedCode: 400,
			expectedRes: response.Response{
				Status: "fail",
				Data: []response.ValidationFailureData{
					{Key: "Note", Details: "The field Note is required"},
				},
			},
		},
	}

	for _, test := range testsPost {
		w := httptest.NewRecorder()

		input, err := json.Marshal(test.params)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, router.PostTodoRoute, bytes.NewReader(input))

		if !test.noMock {
			mockTodoService.EXPECT().
				Create(gomock.Any(), gomock.Eq(test.params)).
				Return(test.mockedTodo, test.mockedErr)
		}
		r.ServeHTTP(w, req)

		expectedJsonBytes, err := json.Marshal(test.expectedRes)
		assert.NoError(t, err)
		expectedJson := string(expectedJsonBytes)
		assert.Equal(t, test.expectedCode, w.Code)
		assert.Equal(t, expectedJson, string(w.Body.Bytes()))
	}
}

func TestPutTodo(t *testing.T) {
	r := gin.Default()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTodoService := mock_service.NewMockITodoService(ctrl)
	todoController := controller.CreateTodoController(mockTodoService)
	r.PUT(router.PutTodoRoute, todoController.PutTodo)

	inputTodo1 := model.Todo{Id: 2, Note: "fdsfsd", Completed: true, Position: 29}

	testsPut := []struct {
		inputId      any
		params       any
		noMock       bool
		mockedErr    error
		expectedCode int
		expectedRes  response.Response
	}{
		{
			inputId:      inputTodo1.Id,
			params:       inputTodo1,
			mockedErr:    nil,
			expectedCode: 200,
			expectedRes: response.Response{
				Status: "success",
			},
		},
		{
			inputId:      inputTodo1.Id,
			params:       inputTodo1,
			mockedErr:    customError.ErrResourceNotFound{Resource: "999"},
			expectedCode: 404,
			expectedRes: response.Response{
				Status: "fail",
				Data:   response.ValidationFailureData{Key: "999", Details: "Resource 999 doesn't exist"},
			},
		},
		{
			inputId:      "notAnId",
			noMock:       true,
			expectedCode: 400,
			expectedRes: response.Response{
				Status: "fail",
				Data:   response.BindingFailureData{BindingFailure: "Uri params binding failed please stick to the appropriate format /todos/:id"},
			},
		},
		{
			inputId:      inputTodo1.Id,
			params:       inputTodo1,
			mockedErr:    mysql.ErrInvalidConn,
			expectedCode: 500,
			expectedRes: response.Response{
				Status:  "error",
				Message: "Internal Server Error",
			},
		},
	}

	for _, test := range testsPut {
		w := httptest.NewRecorder()

		var inputId string

		switch v := test.inputId.(type) {
		case int:
			inputId = strconv.Itoa(v)
		case string:
			inputId = v
		}

		input, err := json.Marshal(test.params)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, strings.Replace(router.PutTodoRoute, ":id", inputId, 1), bytes.NewReader(input))

		if !test.noMock {
			mockTodoService.EXPECT().
				Update(gomock.Any(), gomock.Eq(test.inputId), gomock.Eq(test.params)).
				Return(test.mockedErr)
		}
		r.ServeHTTP(w, req)

		expectedJsonBytes, err := json.Marshal(test.expectedRes)
		assert.NoError(t, err)
		expectedJson := string(expectedJsonBytes)
		assert.Equal(t, test.expectedCode, w.Code)
		assert.Equal(t, expectedJson, string(w.Body.Bytes()))
	}
}

func TestDeleteTodo(t *testing.T) {
	r := gin.Default()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTodoService := mock_service.NewMockITodoService(ctrl)
	todoController := controller.CreateTodoController(mockTodoService)
	r.DELETE(router.PutTodoRoute, todoController.DeleteTodo)

	testsPut := []struct {
		inputId      any
		noMock       bool
		mockedErr    error
		expectedCode int
		expectedRes  response.Response
	}{
		{
			inputId:      2,
			mockedErr:    nil,
			expectedCode: 200,
			expectedRes: response.Response{
				Status: "success",
			},
		},
		{
			inputId:      999,
			mockedErr:    customError.ErrResourceNotFound{Resource: "999"},
			expectedCode: 404,
			expectedRes: response.Response{
				Status: "fail",
				Data:   response.ValidationFailureData{Key: "999", Details: "Resource 999 doesn't exist"},
			},
		},
		{
			inputId:      "notAnId",
			noMock:       true,
			expectedCode: 400,
			expectedRes: response.Response{
				Status: "fail",
				Data:   response.BindingFailureData{BindingFailure: "Uri params binding failed please stick to the appropriate format /todos/:id"},
			},
		},
		{
			inputId:      2,
			mockedErr:    mysql.ErrInvalidConn,
			expectedCode: 500,
			expectedRes: response.Response{
				Status:  "error",
				Message: "Internal Server Error",
			},
		},
	}

	for _, test := range testsPut {
		w := httptest.NewRecorder()

		var inputId string

		switch v := test.inputId.(type) {
		case int:
			inputId = strconv.Itoa(v)
		case string:
			inputId = v
		}

		req := httptest.NewRequest(http.MethodDelete, strings.Replace(router.DeleteTodoRoute, ":id", inputId, 1), nil)

		if !test.noMock {
			mockTodoService.EXPECT().
				Delete(gomock.Any(), gomock.Eq(test.inputId)).
				Return(test.mockedErr)
		}
		r.ServeHTTP(w, req)

		expectedJsonBytes, err := json.Marshal(test.expectedRes)
		assert.NoError(t, err)
		expectedJson := string(expectedJsonBytes)
		assert.Equal(t, test.expectedCode, w.Code)
		assert.Equal(t, expectedJson, string(w.Body.Bytes()))
	}
}

func TestReorderTodo(t *testing.T) {

	r := gin.Default()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTodoService := mock_service.NewMockITodoService(ctrl)
	todoController := controller.CreateTodoController(mockTodoService)
	r.PATCH(router.PatchReorderTodo, todoController.ReorderTodo)

	type wrongParams struct {
		AfterTodoId    string
		TodoToInsertId string
	}

	inputParams1 := dto.ReorderPosTodoParams{AfterTodoId: 2, TodoToInsertId: 5}
	inputParams2 := dto.ReorderPosTodoParams{AfterTodoId: 999, TodoToInsertId: 5}
	inputParams3 := wrongParams{AfterTodoId: "Invalid Id"}
	inputParams4 := dto.ReorderPosTodoParams{}

	testsPut := []struct {
		params       any
		noMock       bool
		mockedErr    error
		expectedCode int
		expectedRes  response.Response
	}{
		{
			params:       inputParams1,
			mockedErr:    nil,
			expectedCode: 200,
			expectedRes: response.Response{
				Status: "success",
			},
		},
		{
			params:       inputParams2,
			mockedErr:    customError.ErrResourceNotFound{Resource: "999"},
			expectedCode: 404,
			expectedRes: response.Response{
				Status: "fail",
				Data:   response.ValidationFailureData{Key: "999", Details: "Resource 999 doesn't exist"},
			},
		},
		{
			params:       inputParams3,
			noMock:       true,
			expectedCode: 400,
			expectedRes: response.Response{
				Status: "fail",
				Data:   response.BindingFailureData{BindingFailure: "Binding the json payload to the corresponding datatype failed, check your inputs!"},
			},
		},
		{
			params:       inputParams4,
			noMock:       true,
			expectedCode: 400,
			expectedRes: response.Response{
				Status: "fail",
				Data: []response.ValidationFailureData{
					{Key: "TodoToInsertId", Details: "The field TodoToInsertId is required"},
				},
			},
		},
		{
			params:       inputParams1,
			mockedErr:    mysql.ErrInvalidConn,
			expectedCode: 500,
			expectedRes: response.Response{
				Status:  "error",
				Message: "Internal Server Error",
			},
		},
	}

	for _, test := range testsPut {
		w := httptest.NewRecorder()

		input, err := json.Marshal(test.params)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPatch, router.PatchReorderTodo, bytes.NewReader(input))

		if !test.noMock {
			mockTodoService.EXPECT().
				ReorderInsert(gomock.Any(), gomock.Eq(test.params)).
				Return(test.mockedErr)
		}
		r.ServeHTTP(w, req)

		expectedJsonBytes, err := json.Marshal(test.expectedRes)
		assert.NoError(t, err)
		expectedJson := string(expectedJsonBytes)
		assert.Equal(t, test.expectedCode, w.Code)
		assert.Equal(t, expectedJson, string(w.Body.Bytes()))
	}
}
