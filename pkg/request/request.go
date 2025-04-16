package request

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	IContextWrapper interface {
		Bind(data any) error
	}

	contextWrapper struct {
		context   echo.Context
		validator *validator.Validate
	}
)

func ContextWrapper(ctx echo.Context) IContextWrapper {
	return &contextWrapper{
		context:   ctx,
		validator: validator.New(),
	}
}

func (c *contextWrapper) Bind(data any) error {
	if err := c.context.Bind(data); err != nil {
		log.Printf("Error: Bind data failed: %s", err.Error())
	}

	if err := c.validator.Struct(data); err != nil {
		log.Printf("Error: Validate data failed: %s", err.Error())
	}

	return nil
}
