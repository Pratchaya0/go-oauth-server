package userhandler

import (
	"github.com/Pratchaya0/go-oauth-server/config"
	userusecase "github.com/Pratchaya0/go-oauth-server/modules/user/userUsecase"
)

type (
	IUserQueueHandler interface{}

	userQueueHandler struct {
		cfg         *config.Config
		userUsecase userusecase.IUserUsecase
	}
)

func NewUserQueueHandler(cfg *config.Config, userUsecase userusecase.IUserUsecase) IUserQueueHandler {
	return &userQueueHandler{cfg: cfg, userUsecase: userUsecase}
}
