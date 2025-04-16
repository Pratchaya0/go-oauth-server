package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Pratchaya0/go-oauth-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DbConn establishes a connection to the PostgreSQL database
func DbConn(pctx context.Context, cfg *config.Config) *gorm.DB {

	// Build connection string from config
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Name,
		cfg.Db.SSLMode,
	)

	// Configure GORM with custom logger settings
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Attempt to connect to the database
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Set connection pool parameters
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	// Set connection pool limits
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Successfully connected to database")

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error retrieving sqlDB for closing: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}
