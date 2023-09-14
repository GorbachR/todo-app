package dao

type Todo struct {
	Id     int    `json:"id"`
	Note   string `json:"note" binding:"required"`
	Active bool   `json:"active" binding:"required"`
}
