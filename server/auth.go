package server

import (
	"log"

	authhandler "github.com/Pratchaya0/go-oauth-server/modules/auth/authHandler"
	go_oauth_server "github.com/Pratchaya0/go-oauth-server/modules/auth/authProtobuf"
	authrepository "github.com/Pratchaya0/go-oauth-server/modules/auth/authRepository"
	authusecase "github.com/Pratchaya0/go-oauth-server/modules/auth/authUsecase"
	pkggrpc "github.com/Pratchaya0/go-oauth-server/pkg/grpc"
)

func (s *server) authService() {
	repo := authrepository.NewAuthRepository(s.db)
	usecase := authusecase.NewAuthUsecase(repo)
	httpHandler := authhandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authhandler.NewAuthGrpcHandler(usecase)

	// grpc
	go func() {
		grpcServer, lis := pkggrpc.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		go_oauth_server.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("%s gRPC server listening on %s", s.cfg.App.Name, s.cfg.Grpc.AuthUrl)

		grpcServer.Serve(lis)
	}()

	_ = httpHandler

	auth := s.app.Group("api/v1/auth")

	// Health Check
	auth.GET("/health", s.healthCheckService)

}
