package dto

type GetTodosQueryParams struct {
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
	Q      string `form:"q"`
}
