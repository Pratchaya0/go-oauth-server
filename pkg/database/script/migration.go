package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Pratchaya0/go-oauth-server/config"
	"github.com/Pratchaya0/go-oauth-server/pkg/database"
	"github.com/Pratchaya0/go-oauth-server/pkg/database/migrations"
)

func main() {
	ctx := context.Background()

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Println("Run database migrations mode auth dev")
			return ("./env/dev/.env.auth")
		}

		log.Printf("Run database migrations mode %s %s", os.Args[1], os.Args[2])
		return (fmt.Sprintf("./env/%s/.env.%s", os.Args[2], os.Args[1]))
	}())

	db := database.DbConn(ctx, cfg)
	defer database.CloseDB(db)

	switch cfg.App.Name {
	case "RitsukoAuth":
		migrations.AuthMigerate(ctx, cfg, db)
	case "RitsukoUser":
		migrations.AuthMigerate(ctx, cfg, db)
	default:
		log.Fatal("Error: Invalid service!")
	}

	log.Printf("Migration successful!!!")

}
