package userusecase

import userrepository "github.com/Pratchaya0/go-oauth-server/modules/user/userRepository"

type (
	IUserUsecase interface{}

	userUsercase struct {
		userRepository userrepository.IUserRepository
	}
)

func NewUserUsecase(userRepository userrepository.IUserRepository) IUserUsecase {
	return &userUsercase{userRepository}
}
