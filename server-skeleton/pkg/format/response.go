package format

import (
	"encoding/json"
)

type ResponseValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Other   any    `json:"other"`
}

type Response struct {
	Code   int             `json:"statusCode"`
	Errors []ResponseError `json:"errors"`
	Body   any             `json:"body"`
}

func ResponseFormat(code int, body any, errors []ResponseError) string {

	jsonData, err := json.Marshal(Response{
		Code:   code,
		Errors: errors,
		Body:   body,
	})

	if err != nil {
		panic(err)
	}

	return string(jsonData)
}
