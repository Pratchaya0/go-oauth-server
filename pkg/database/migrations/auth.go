package migrations

import (
	"context"
	"log"

	"github.com/Pratchaya0/go-oauth-server/config"
	"github.com/Pratchaya0/go-oauth-server/modules/auth"
	"gorm.io/gorm"
)

func AuthMigerate(pctx context.Context, cfg *config.Config, db *gorm.DB) {

	// Auto-migrate database models
	err := db.WithContext(pctx).AutoMigrate(
		// Add your models here
		&auth.User{},
		&auth.Role{},
		&auth.Permission{},
		&auth.OAuthClient{},
		&auth.OAuthAccessToken{},
		&auth.OAuthRefreshToken{},
		&auth.OAuthAuthorizationCode{},
		&auth.APIKey{},
		&auth.Service{},
		&auth.ServicePermission{},
		&auth.OAuthScope{},
		&auth.ScopePermission{},
		&auth.Session{},
		&auth.AuditLog{},
	)

	if err != nil {
		log.Fatalf("Auto-migration failed: %v", err)
	}

}
