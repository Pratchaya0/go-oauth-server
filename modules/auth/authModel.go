package auth

import (
	"time"

	"github.com/Pratchaya0/go-oauth-server/modules/user"
)

type (
	LoginRequestDTO struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
	}

	RefreshTokenRequestDTO struct {
		RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required,max=500"`
	}

	InsertUserRole struct {
		UserId   string `json:"user_id" validate:"required"`
		RoleCode []int  `json:"role" validate:"required"`
	}

	ProfileIntercepterRequestDTO struct {
		// Fix this letter
		*user.UserProfileResponseDTO
		Credentail *CredentialResponseDTO `json:"credential"`
	}

	CredentialResponseDTO struct {
		ID          uint      `json:"credential_id"`
		UserID      uint      `json:"user_id"`
		RoleID      uint      `json:"role_id"`
		AccessToken string    `json:"access_token"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updaeted_at"`
	}
)
