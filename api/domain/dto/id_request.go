package dto

type IdRequest struct {
	Id int `uri:"id" binding:"required"`
}
