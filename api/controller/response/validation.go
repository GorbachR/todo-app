package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const (
	BindingJSON  = "json"
	BindingQuery = "query"
)

type ValidationFailureData struct {
	Key     string `json:"key"`
	Details string `json:"details"`
}

type BindingFailureData struct {
	BindingFailure string `json:"bindingFailure"`
}

func GetErrorMessage(code int) string {
	return errorMessageMap[code]
}

func TagToDetails(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s %s %s", "The field", fe.Field(), "is required")
	}
	return ""
}

func ConstructValidationFailResponse(ve validator.ValidationErrors) (re Response) {
	re.Status = Fail
	valErrData := make([]ValidationFailureData, len(ve))
	for i, fe := range ve {
		valErrData[i] = ValidationFailureData{Key: fe.Field(), Details: TagToDetails(fe)}
	}
	re.Data = valErrData
	return
}

func ConstructBindingFailResponse(binding string) (re Response) {
	re.Status = Fail
	re.Data = BindingFailureData{
		BindingFailure: fmt.Sprintf("%s %s %s", "Binding the", binding, "payload to the corresponding datatype failed, check your inputs!")}
	return
}

func ConstructUriBindingFailResponse(url string) (re Response) {
	re.Status = Fail
	re.Data = BindingFailureData{BindingFailure: "Uri params binding failed please stick to the appropriate format " + url}
	return
}
