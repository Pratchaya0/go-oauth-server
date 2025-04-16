package auth

import (
	"time"

	"gorm.io/datatypes"
)

type (
	User struct {
		ID           uint   `gorm:"primaryKey"`
		Username     string `gorm:"unique;not null"`
		PasswordHash string `gorm:"not null"`
		Sessions     []Session
		APIKeys      []APIKey
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

	OAuthClient struct {
		ID                uint   `gorm:"primaryKey"`
		ClientID          string `gorm:"unique;not null"`
		ClientSecret      string `gorm:"not null"`
		Name              string `gorm:"not null"`
		Description       string
		WebsiteURL        string
		RedirectURIs      []string `gorm:"type:text[]"`
		AllowedGrantTypes []string `gorm:"type:text[]"`
		IsConfidential    bool     `gorm:"default:true"`
		UserID            *uint
		User              *User `gorm:"constraint:OnDelete:SET NULL"`
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}

	OAuthAccessToken struct {
		ID          uint   `gorm:"primaryKey"`
		AccessToken string `gorm:"unique;not null"`
		ClientID    uint
		UserID      uint
		Scopes      []string `gorm:"type:text[]"`
		ExpiresAt   time.Time
		CreatedAt   time.Time
	}

	OAuthRefreshToken struct {
		ID            uint   `gorm:"primaryKey"`
		RefreshToken  string `gorm:"unique;not null"`
		AccessTokenID uint
		IsRevoked     bool `gorm:"default:false"`
		ExpiresAt     time.Time
		CreatedAt     time.Time
	}

	OAuthAuthorizationCode struct {
		ID                uint   `gorm:"primaryKey"`
		AuthorizationCode string `gorm:"unique;not null"`
		ClientID          uint
		UserID            uint
		RedirectURI       string   `gorm:"not null"`
		Scopes            []string `gorm:"type:text[]"`
		ExpiresAt         time.Time
		CreatedAt         time.Time
	}

	APIKey struct {
		ID         uint `gorm:"primaryKey"`
		UserID     uint
		APIKey     string   `gorm:"unique;not null"`
		APISecret  string   `gorm:"not null"`
		Name       string   `gorm:"not null"`
		IsActive   bool     `gorm:"default:true"`
		Scopes     []string `gorm:"type:text[]"`
		ExpiresAt  *time.Time
		LastUsedAt *time.Time
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}

	Service struct {
		ID          uint   `gorm:"primaryKey"`
		Name        string `gorm:"unique;not null"`
		Description string
		BaseURL     string
		IsActive    bool `gorm:"default:true"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	ServicePermission struct {
		ID           uint `gorm:"primaryKey"`
		ServiceID    uint
		PermissionID uint
		CreatedAt    time.Time
	}

	OAuthScope struct {
		ID          uint   `gorm:"primaryKey"`
		Name        string `gorm:"unique;not null"`
		Description string
		CreatedAt   time.Time
	}

	ScopePermission struct {
		ID           uint `gorm:"primaryKey"`
		ScopeID      uint
		PermissionID uint
	}

	Session struct {
		ID             uint `gorm:"primaryKey"`
		UserID         uint
		Token          string `gorm:"unique;not null"`
		IPAddress      string
		UserAgent      string
		IsValid        bool `gorm:"default:true"`
		ExpiresAt      time.Time
		CreatedAt      time.Time
		LastActivityAt time.Time
	}

	AuditLog struct {
		ID           uint `gorm:"primaryKey"`
		UserID       *uint
		ClientID     *uint
		Action       string `gorm:"not null"`
		ResourceType string
		ResourceID   string
		IPAddress    string
		UserAgent    string
		Details      datatypes.JSON
		CreatedAt    time.Time
	}
)
