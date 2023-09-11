package schema

type Todo struct {
	Id     int    `json:"id"`
	Note   string `json:"note"`
	Active bool   `json:"active"`
}
