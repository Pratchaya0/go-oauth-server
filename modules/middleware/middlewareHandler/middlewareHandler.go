package middlewarehandler

import (
	"github.com/Pratchaya0/go-oauth-server/config"
	middlewareusecase "github.com/Pratchaya0/go-oauth-server/modules/middleware/middlewareUsecase"
)

type (
	IMiddlewareHandler interface{}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase middlewareusecase.IMiddlewareUsecase
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareusecase.IMiddlewareUsecase) IMiddlewareHandler {
	return &middlewareHandler{cfg, middlewareUsecase}
}
