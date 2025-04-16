package authusecase

import authrepository "github.com/Pratchaya0/go-oauth-server/modules/auth/authRepository"

type (
	IAuthUsecase interface{}

	authUsecase struct {
		authRepository authrepository.IAuthRepository
	}
)

func NewAuthUsecase(authRepository authrepository.IAuthRepository) IAuthUsecase {
	return &authUsecase{authRepository: authRepository}
}
