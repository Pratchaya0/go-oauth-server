package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Pratchaya0/go-oauth-server/config"
	middlewarehandler "github.com/Pratchaya0/go-oauth-server/modules/middleware/middlewareHandler"
	middlewarerepository "github.com/Pratchaya0/go-oauth-server/modules/middleware/middlewareRepository"
	middlewareusecase "github.com/Pratchaya0/go-oauth-server/modules/middleware/middlewareUsecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	server struct {
		app        *echo.Echo
		db         *gorm.DB
		cfg        *config.Config
		middleware middlewarehandler.IMiddlewareHandler
	}
)

func newMiddleware(cfg *config.Config) middlewarehandler.IMiddlewareHandler {
	repo := middlewarerepository.NewMiddlewareRepository()
	usecase := middlewareusecase.NewMiddlewareUsecase(repo)
	return middlewarehandler.NewMiddlewareHandler(cfg, usecase)
}

func (s *server) gracefulShutDown(pctx context.Context, quit <-chan os.Signal) {
	log.Printf("Start service: %s", s.cfg.App.Name)
	<-quit
	log.Println("Shutting down server ...")

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func (s *server) httpListening() {
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}
}

func Start(ptcx context.Context, cfg *config.Config, db *gorm.DB) {
	s := &server{
		app:        echo.New(),
		db:         db,
		cfg:        cfg,
		middleware: newMiddleware(cfg),
	}

	s.middlewareService()

	switch s.cfg.App.Name {
	case "RitsukoAuth":
		s.authService()
	case "RitsukoUser":
		s.userService()
	default:
		log.Fatal("Error: Invalid service!")
	}

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go s.gracefulShutDown(ptcx, quit)

	s.httpListening()

}
