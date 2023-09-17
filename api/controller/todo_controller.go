package controller

import (
	"errors"
	"github.com/GorbachR/todo-app/api/controller/validation"
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
}

func CreateTodoController(s service.ITodoService) TodoController {
	return TodoController{TodoService: s}
}

func (t TodoController) GetTodos(c *gin.Context) {
	var limitAndOffset dto.LimitAndOffset

	if err := c.ShouldBindQuery(&limitAndOffset); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &err)
		return
	}

	if limitAndOffset.Limit == 0 {
		limitAndOffset.Limit = 5
	}

	todos, err := t.TodoService.FindAll(limitAndOffset)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	if len(todos) == 0 {
		c.JSON(http.StatusOK, []model.Todo{})
		return
	}

	c.JSON(http.StatusOK, &todos)

}

func (t TodoController) GetTodo(c *gin.Context) {
	var todoId dto.IdRequest

	if err := c.ShouldBindUri(&todoId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := t.TodoService.FindOne(todoId.Id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, &todo)
}

func (t TodoController) PostTodo(c *gin.Context) {
	var newTodo model.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]validation.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = validation.ErrorMsg{Field: fe.Field(), Message: validation.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	newId, err := t.TodoService.Create(newTodo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
	}

	c.JSON(http.StatusCreated, gin.H{"id": int(newId)})
}

func (t TodoController) PutTodo(c *gin.Context) {
	var changedTodo model.Todo
	var todoId dto.IdRequest

	if err := c.ShouldBindUri(&todoId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&changedTodo); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]validation.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = validation.ErrorMsg{Field: fe.Field(), Message: validation.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	err := t.TodoService.Update(todoId.Id, changedTodo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.Status(http.StatusNoContent)

}

func (t TodoController) DeleteTodo(c *gin.Context) {
	var todoId dto.IdRequest

	if err := c.ShouldBindUri(&todoId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := t.TodoService.Delete(todoId.Id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.Status(http.StatusOK)
}
