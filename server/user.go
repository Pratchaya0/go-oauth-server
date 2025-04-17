package server

import (
	"log"

	userhandler "github.com/Pratchaya0/go-oauth-server/modules/user/userHandler"
	go_oauth_server "github.com/Pratchaya0/go-oauth-server/modules/user/userProtobuf"
	userrepository "github.com/Pratchaya0/go-oauth-server/modules/user/userRepository"
	userusecase "github.com/Pratchaya0/go-oauth-server/modules/user/userUsecase"
	pkggrpc "github.com/Pratchaya0/go-oauth-server/pkg/grpc"
)

func (s *server) userService() {
	repo := userrepository.NewUserRepository(s.db)
	usecase := userusecase.NewUserUsecase(repo)
	httpHandler := userhandler.NewUserHandler(s.cfg, usecase)
	grpcHandler := userhandler.NewUserGrpcHandler(usecase)
	queueHandler := userhandler.NewUserQueueHandler(s.cfg, usecase)

	// grpc
	go func() {
		grpcServer, lis := pkggrpc.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.UserUrl)

		go_oauth_server.RegisterUserGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("%s gRPC server listening on %s", s.cfg.App.Name, s.cfg.Grpc.UserUrl)

		grpcServer.Serve(lis)
	}()

	_ = queueHandler

	user := s.app.Group("api/v1/user")

	user.GET("/health", s.healthCheckService)
	user.POST("/create", httpHandler.CreateOneUser)
	user.POST("/update/:userId", httpHandler.UpdateOneUserDetails)

}
