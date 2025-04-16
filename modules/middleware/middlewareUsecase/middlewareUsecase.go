package middlewareusecase

import middlewarerepository "github.com/Pratchaya0/go-oauth-server/modules/middleware/middlewareRepository"

type (
	IMiddlewareUsecase interface{}

	middlewareUsecase struct {
		middlewareRepository middlewarerepository.IMiddlewareRepository
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewarerepository.IMiddlewareRepository) IMiddlewareUsecase {
	return &middlewareUsecase{middlewareRepository: middlewareRepository}
}
