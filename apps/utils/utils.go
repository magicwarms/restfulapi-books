package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Success bool         `json:"success"`
	Data    interface{}  `json:"data"`
	Error   *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// AppResponse generates a JSON response for the given request.
//
// It takes in the following parameters:
// - c: the echo.Context object representing the request and response.
// - code: an integer representing the HTTP status code of the response.
// - data: an interface{} representing the data to be included in the response.
//
// It returns an error.
func AppResponse(ctx echo.Context, code int, data interface{}) error {
	response := Response{
		Success: code <= 400,
		Data:    data,
		Error:   &ErrorDetail{},
	}
	if !response.Success {
		errMessage := fmt.Sprintf("%v", data)
		response.Data = nil

		if strings.Contains(errMessage, "not found") {
			code = http.StatusNotFound
		}

		response.Error = &ErrorDetail{
			Code:    code,
			Message: errMessage,
		}
	}

	if code == http.StatusBadRequest {
		response.Success = false
	}

	return ctx.JSONPretty(code, response, " ")
}
