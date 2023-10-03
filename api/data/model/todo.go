package model

import (
	"time"
)

type Todo struct {
	Id          int       `json:"id"`
	Note        string    `json:"note" binding:"required"`
	Completed   bool      `json:"completed"`
	Position    int       `json:"position"`
	LastChanged time.Time `json:"lastChanged"`
}
