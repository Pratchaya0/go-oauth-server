package migrations

import (
	"context"
	"log"

	"github.com/Pratchaya0/go-oauth-server/config"
	"github.com/Pratchaya0/go-oauth-server/modules/user"
	"gorm.io/gorm"
)

func UserMigerate(pctx context.Context, cfg *config.Config, db *gorm.DB) {

	// Auto-migrate database models
	err := db.WithContext(pctx).AutoMigrate(
		// Add your models here
		&user.User{},
		&user.Role{},
		&user.Permission{},
		&user.UserRole{},
		&user.RolePermission{},
	)

	if err != nil {
		log.Fatalf("Auto-migration failed: %v", err)
	}

}
