package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func HandleErrorValidator(err error) error {
	var errMessage string
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				errMessage = fmt.Sprintf("%s is required", err.Field())
			case "email":
				errMessage = fmt.Sprintf("%s is not valid", err.Field())
			case "min":
				errMessage = fmt.Sprintf("%s min 8 character", err.Field())
			case "max":
				errMessage = fmt.Sprintf("%s max 12 character", err.Field())
			case "numeric":
				errMessage = fmt.Sprintf("%s character must is numeric", err.Field())
			case "url":
				errMessage = fmt.Sprintf("%s is not valid", err.Field())
			case "eq=active|eq=draft":
				errMessage = fmt.Sprintf("%s is not valid", err.Field())
			}

		}
	}

	return errors.New(errMessage)
}

// handler error message and status code
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) StatusCode() int {
	return e.Code
}

func GetStatusCode(err error) int {
	if err == nil {
		return 200
	}

	if e, ok := err.(*Error); ok {
		return e.StatusCode()
	}

	return 500
}

func HandleError(code int, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
