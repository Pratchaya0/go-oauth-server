package user

import "time"

type (
	UserProfileResponseDTO struct {
		UserId     string    `json:"user_id"`
		Email      string    `json:"email"`
		Username   string    `json:"username"`
		CreatedAt  time.Time `json:"created_at"`
		UppdatedAt time.Time `json:"updated_at"`
	}

	UserClaimsReponseDTO struct {
		ID     uint `json:"user_id"`
		RoleID uint `json:"role_id"`
	}

	CreateUserRequestDTO struct {
		UserName string `json:"username" form:"username" validate:"required,max=64"`
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
	}
)
