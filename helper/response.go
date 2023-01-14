package helper

import "strings"

type ResponseError struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Failed to process request"`
	Errors  string `json:"errors"`
}

type EmptyObj struct{}

func BuildErrorResponse(message string, err string) ResponseError {
	splitedError := strings.Split(err, "\n")
	res := ResponseError{
		Success: false,
		Message: message,
		Errors:  splitedError[0],
	}

	return res
}
