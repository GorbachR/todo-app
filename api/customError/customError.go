package customError

type ErrResourceUnchanged struct {
	Resource string `json:"resource"`
}

func (r ErrResourceUnchanged) Error() string {
	return "Resource is unchanged"
}

type ErrResourceNotFound struct {
	Resource string `json:"resource"`
}

func (r ErrResourceNotFound) Error() string {
	return "Resource wasn't found"
}
