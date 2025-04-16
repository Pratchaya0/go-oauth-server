package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	messageResponse[T any] struct {
		IsSuccess bool   `json:"isSuccess"`
		Message   string `json:"message"`
		Data      *T     `json:"data"`
	}
)

func ErrorResponse[T any](ec echo.Context, statusCode int, message string) error {
	return ec.JSON(statusCode, &messageResponse[T]{
		IsSuccess: false,
		Message:   message,
		Data:      nil,
	})
}

func SuccessReponse[T any](ec echo.Context, data T) error {
	return ec.JSON(http.StatusOK, &messageResponse[T]{
		IsSuccess: true,
		Message:   "Success",
		Data:      &data,
	})
}
