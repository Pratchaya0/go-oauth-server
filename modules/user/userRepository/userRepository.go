package userrepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Pratchaya0/go-oauth-server/modules/user"
	"gorm.io/gorm"
)

type (
	IUserRepository interface {
		IsUniqueUser(pctx context.Context, email, username string) bool
		InsertUser(pctx context.Context, request *user.User) (uint, error)
		FindOneUserByID(pctx context.Context, userID uint) (*user.User, error)
		UpdateOneUser(pctx context.Context, request *user.User) (*uint, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) IsUniqueUser(pctx context.Context, email, username string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	var existingUser user.User
	err := r.db.WithContext(ctx).
		Where("email = ? OR username = ?", email, username).
		First(&existingUser).Error

	if err == nil {
		log.Printf("IsUniqueUser: user with email or username already exists")
		return false
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	log.Printf("IsUniqueUser: failed to query user: %v", err)
	return false
}

func (r *userRepository) InsertUser(pctx context.Context, request *user.User) (uint, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	if err := r.db.WithContext(ctx).Create(request).Error; err != nil {
		log.Printf("CreateUser: failed to create user: %v", err)
		return 0, err
	}

	return request.ID, nil
}

func (r *userRepository) FindOneUserByID(pctx context.Context, userID uint) (*user.User, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	var user user.User

	if err := r.db.WithContext(ctx).Model(&user).Find(user.ID == userID).Error; err != nil {
		log.Printf("FindOneUserByID: user not found")
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateOneUser(pctx context.Context, request *user.User) (*uint, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	if err := r.db.WithContext(ctx).Model(&user.User{}).Where(&user.User{ID: request.ID}).Updates(&request).Error; err != nil {
		log.Printf("UpdateOneUser: cannot update user")
		return nil, err
	}

	return &request.ID, nil
}
