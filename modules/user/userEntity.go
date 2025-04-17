package user

import "time"

type (
	User struct {
		ID                          uint   `gorm:"primaryKey"`
		Username                    string `gorm:"unique;not null"`
		Email                       string `gorm:"unique;not null"`
		PasswordHash                []byte `gorm:"not null"`
		FirstName                   string
		LastName                    string
		IsActive                    bool `gorm:"default:true"`
		IsVerified                  bool `gorm:"default:false"`
		VerificationToken           string
		VerificationTokenExpiresAt  *time.Time
		ResetPasswordToken          string
		ResetPasswordTokenExpiresAt *time.Time
		CreatedAt                   time.Time
		UpdatedAt                   time.Time
		Roles                       []Role `gorm:"many2many:user_roles"`
	}

	Role struct {
		ID           uint   `gorm:"primaryKey"`
		Name         string `gorm:"unique;not null"`
		Description  string
		IsSystemRole bool `gorm:"default:false"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
		Permissions  []Permission `gorm:"many2many:role_permissions"`
	}

	Permission struct {
		ID          uint   `gorm:"primaryKey"`
		Name        string `gorm:"unique;not null"`
		Description string
		Resource    string `gorm:"not null"`
		Action      string `gorm:"not null"`
		CreatedAt   time.Time
	}

	UserRole struct {
		ID        uint `gorm:"primaryKey"`
		UserID    uint
		RoleID    uint
		CreatedAt time.Time
	}

	RolePermission struct {
		ID           uint `gorm:"primaryKey"`
		RoleID       uint
		PermissionID uint
		CreatedAt    time.Time
	}
)
