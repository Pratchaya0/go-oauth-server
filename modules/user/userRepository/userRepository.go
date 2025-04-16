package userrepository

import "gorm.io/gorm"

type (
	IUserRepository interface{}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}
