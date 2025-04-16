package authrepository

import (
	"gorm.io/gorm"
)

type (
	IAuthRepository interface{}

	authRepository struct {
		db *gorm.DB
	}
)

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &authRepository{db: db}
}
