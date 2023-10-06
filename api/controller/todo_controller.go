package controller

import (
	"errors"
	"github.com/GorbachR/todo-app/api/controller/response"
	"github.com/GorbachR/todo-app/api/customError"
	"github.com/GorbachR/todo-app/api/data/dto"
	"github.com/GorbachR/todo-app/api/data/model"
	"github.com/GorbachR/todo-app/api/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type TodoController struct {
	TodoService service.ITodoService
}

type ITodoController interface {
	GetTodos(*gin.Context)
	GetTodo(*gin.Context)
	PostTodo(*gin.Context)
	PutTodo(*gin.Context)
	DeleteTodo(*gin.Context)
	ReorderTodo(*gin.Context)
}

func CreateTodoController(s service.ITodoService) *TodoController {
	return &TodoController{TodoService: s}
}

func (t TodoController) GetTodos(c *gin.Context) {
	var getTodoQueryParams dto.GetTodosQueryParams

	if err := c.ShouldBindQuery(&getTodoQueryParams); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ConstructBindingFailResponse(response.BindingQuery))
		return
	}

	if getTodoQueryParams.Limit == 0 {
		getTodoQueryParams.Limit = 5
	}

	todos, err := t.TodoService.FindAll(c.Request.Context(), getTodoQueryParams)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ConstructInternalServerError(""))
		return
	}

	resp := response.ConstructSuccess(todos)
	c.JSON(http.StatusOK, &resp)
}

func (t TodoController) GetTodo(c *gin.Context) {
	var todoId dto.IdRequest

	if err := c.ShouldBindUri(&todoId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ConstructUriBindingFailResponse(c.FullPath()))
		return
	}

	todo, err := t.TodoService.FindOne(c.Request.Context(), todoId.Id)

	if err != nil {
		var errResNF customError.ErrResourceNotFound
		if errors.As(err, &errResNF) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.ConstructResourceNotFound(errResNF.Resource))
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.ConstructInternalServerError(""))
		}
		return
	}

	c.JSON(http.StatusOK, response.ConstructSuccess(todo))
}

func (t TodoController) PostTodo(c *gin.Context) {
	var newTodo model.Todo

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		var valFail validator.ValidationErrors
		if errors.As(err, &valFail) {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ConstructValidationFailResponse(valFail))
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ConstructBindingFailResponse(response.BindingJSON))
		}
		return
	}

	todo, err := t.TodoService.Create(c.Request.Context(), newTodo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ConstructInternalServerError(""))
		return
	}

	//TODO add location header

	c.JSON(http.StatusCreated, response.ConstructSuccess(todo))
}

func (t TodoController) PutTodo(c *gin.Context) {
	var changedTodo model.Todo
	var todoId dto.IdRequest

	if err := c.ShouldBindUri(&todoId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ConstructUriBindingFailResponse(c.FullPath()))
		return
	}

	if err := c.ShouldBindJSON(&changedTodo); err != nil {
		var valFail validator.ValidationErrors
		if errors.As(err, &valFail) {
			resp := response.ConstructValidationFailResponse(valFail)
			c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ConstructBindingFailResponse(response.BindingJSON))
		}
		return
	}

	err := t.TodoService.Update(c.Request.Context(), todoId.Id, changedTodo)

	if err != nil {
		var errResNF customError.ErrResourceNotFound
		if errors.As(err, &errResNF) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.ConstructResourceNotFound(errResNF.Resource))
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.ConstructInternalServerError(""))
		}
		return
	}

	c.JSON(http.StatusOK, response.ConstructSuccess(nil))
}

func (t TodoController) DeleteTodo(c *gin.Context) {
	var todoId dto.IdRequest

	if err := c.ShouldBindUri(&todoId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ConstructUriBindingFailResponse(c.FullPath()))
		return
	}

	err := t.TodoService.Delete(c.Request.Context(), todoId.Id)

	if err != nil {
		var errResNF customError.ErrResourceNotFound
		if errors.As(err, &errResNF) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.ConstructResourceNotFound(errResNF.Resource))
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.ConstructInternalServerError(""))
		}
		return
	}

	c.JSON(http.StatusOK, response.ConstructSuccess(nil))
}

func (t TodoController) ReorderTodo(c *gin.Context) {
	var reorderTodoParams dto.ReorderPosTodoParams

	if err := c.ShouldBindJSON(&reorderTodoParams); err != nil {
		var valFail validator.ValidationErrors
		if errors.As(err, &valFail) {
			resp := response.ConstructValidationFailResponse(valFail)
			c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ConstructBindingFailResponse(response.BindingJSON))
		}
		return
	}

	err := t.TodoService.ReorderInsert(c.Request.Context(), reorderTodoParams)

	if err != nil {
		var errResNF customError.ErrResourceNotFound
		if errors.As(err, &errResNF) {
			c.AbortWithStatusJSON(http.StatusNotFound, response.ConstructResourceNotFound(errResNF.Resource))
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.ConstructInternalServerError(""))
		}
		return
	}

	c.JSON(http.StatusOK, response.ConstructSuccess(nil))
}
