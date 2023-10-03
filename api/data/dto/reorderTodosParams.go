package dto

type ReorderPosTodoParams struct {
	AfterTodoId    int `json:"afterTodoId"`
	TodoToInsertId int `json:"todoToInsertId" binding:"required"`
}
