package response

import (
	"fmt"
)

const (
	Fail    = "fail"
	Success = "success"
	Error   = "error"
)

const (
	UnknownError      = 1000
	InvalidQueryParam = 1001
)

var errorMessageMap = map[int]string{
	InvalidQueryParam: "The supplied query param is invalid",
	UnknownError:      "Something went wrong",
}

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ResourceData struct {
	Resource interface{}
}

// TODO an update to a record that already exists with the same data will result in this customError

func ConstructResourceNotFound(id string) (re Response) {
	re.Status = Fail
	re.Data = ValidationFailureData{Key: id, Details: fmt.Sprintf("%s %s %s", "Resource", id, "doesn't exist")}
	return
}

func ConstructInternalServerError(message string, data ...any) (re Response) {
	re.Status = Error
	if message == "" {
		re.Message = "Internal Server Error"
	} else {
		re.Message = message
	}
	if len(data) != 0 {
		re.Data = data
	}
	return
}

func ConstructSuccess(data any) (re Response) {
	re.Status = Success
	re.Data = data
	return
}
