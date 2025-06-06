package format

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

const ValidationErrorMessage = "Error validation"

var validate = validator.New()

func ValidateFormat(s interface{}) []ResponseError {
	var messages []ResponseError
	if err := validate.Struct(s); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, ve := range validationErrors {
			responseValidationError := ResponseValidationError{
				Field:   ve.Field(),
				Message: ve.ActualTag(),
			}

			responseError := ResponseError{
				Code:    http.StatusBadRequest,
				Message: ValidationErrorMessage,
				Other:   responseValidationError,
			}
			messages = append(messages, responseError)
		}
	}
	return messages
}
