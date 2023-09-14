package service

import (
	"errors"
	"net/http"

	"github.com/GorbachR/todo-app/api/domain/dao"
	"github.com/GorbachR/todo-app/api/domain/dto"
	"github.com/GorbachR/todo-app/api/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func getTodos(c *gin.Context) {

	var limitAndOffset dto.LimitAndOffset

	if err := c.ShouldBindQuery(&limitAndOffset); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &err)
		return
	}

	if limitAndOffset.Limit == 0 {
		limitAndOffset.Limit = 5
	}

	todos, err := data.QueryDb.GetTodos(limitAndOffset)

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

func postTodo(c *gin.Context) {
	var newTodo dao.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]pkg.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = pkg.ErrorMsg{fe.Field(), pkg.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	res, err := data.QueryDb.CreateTodo(newTodo)
	lastInsert, insertErr := res.LastInsertId()

	if err != nil || insertErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": int(lastInsert)})
}

func putTodo(c *gin.Context) {
	var changedTodo dao.Todo
	var todoId dto.IdRequest

	if err := c.ShouldBindUri(&todoId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&changedTodo); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]pkg.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = pkg.ErrorMsg{fe.Field(), pkg.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	err := data.QueryDb.UpdateTodo(todoId.Id, changedTodo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.Status(http.StatusNoContent)
}

func deleteTodo(c *gin.Context) {

	var todoId dto.IdRequest

	if err := c.ShouldBindUri(&todoId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.QueryDb.DeleteTodo(todoId.Id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.Status(http.StatusOK)
}

func getTodo(c *gin.Context) {

	var todoId dto.IdRequest

	if err := c.ShouldBindUri(&todoId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := data.QueryDb.GetTodo(todoId.Id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, &todo)
}
