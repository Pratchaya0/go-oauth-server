package authhandler

import (
	"github.com/Pratchaya0/go-oauth-server/config"
	authusecase "github.com/Pratchaya0/go-oauth-server/modules/auth/authUsecase"
)

type (
	IAuthHttpHandler interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authusecase.IAuthUsecase
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authusecase.IAuthUsecase) IAuthHttpHandler {
	return &authHttpHandler{cfg, authUsecase}
}
