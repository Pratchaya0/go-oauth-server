package userusecase

import (
	"context"
	"errors"

	"github.com/Pratchaya0/go-oauth-server/modules/user"
	userrepository "github.com/Pratchaya0/go-oauth-server/modules/user/userRepository"
	"github.com/Pratchaya0/go-oauth-server/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	IUserUsecase interface {
		CreateOneUser(pctx context.Context, request *user.CreateUserRequestDTO) (*uint, error)
		UpdateOneUserDetails(pctx context.Context, userID uint, request *user.UpdateUserRequestDTO) (*uint, error)
	}

	userUsercase struct {
		userRepository userrepository.IUserRepository
	}
)

func NewUserUsecase(userRepository userrepository.IUserRepository) IUserUsecase {
	return &userUsercase{userRepository}
}

func (u *userUsercase) CreateOneUser(pctx context.Context, request *user.CreateUserRequestDTO) (*uint, error) {

	if !u.userRepository.IsUniqueUser(pctx, request.Email, request.UserName) {
		return nil, errors.New("error: email or username already exist")
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error: failed to hash password")
	}

	userID, err := u.userRepository.InsertUser(pctx, &user.User{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Username:     request.UserName,
		Email:        request.Email,
		PasswordHash: hashed,
		CreatedAt:    utils.LocalTime(),
		UpdatedAt:    utils.LocalTime(),
	})

	if err != nil {
		return nil, errors.New("error: failed to inser user")
	}

	return &userID, nil
}

func (u *userUsercase) UpdateOneUserDetails(pctx context.Context, userID uint, request *user.UpdateUserRequestDTO) (*uint, error) {

	user, err := u.userRepository.FindOneUserByID(pctx, userID)
	if err != nil {
		return nil, errors.New("error: user not found")
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName

	result, err := u.userRepository.UpdateOneUser(pctx, user)
	if err != nil {
		return nil, errors.New("error: failed to update user")
	}

	return result, nil

}
