package userhandler

import (
	"github.com/Pratchaya0/go-oauth-server/config"
	userusecase "github.com/Pratchaya0/go-oauth-server/modules/user/userUsecase"
)

type (
	IUserHttpHandler interface{}

	userHttpHandler struct {
		cfg         *config.Config
		userUsecase userusecase.IUserUsecase
	}
)

func NewUserHandler(cfg *config.Config, userUsscase userusecase.IUserUsecase) IUserHttpHandler {
	return &userHttpHandler{cfg, userUsscase}
}
