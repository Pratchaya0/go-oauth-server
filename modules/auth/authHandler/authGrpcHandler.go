package authhandler

import (
	"context"

	go_oauth_server "github.com/Pratchaya0/go-oauth-server/modules/auth/authProtobuf"
	authusecase "github.com/Pratchaya0/go-oauth-server/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		go_oauth_server.UnimplementedAuthGrpcServiceServer
		authUsecase authusecase.IAuthUsecase
	}
)

func NewAuthGrpcHandler(authUsecase authusecase.IAuthUsecase) *authGrpcHandler {
	return &authGrpcHandler{authUsecase: authUsecase}
}

func (gh *authGrpcHandler) AccessTokenCheck(ctx context.Context, request *go_oauth_server.AccessTokenCheckRequest) (*go_oauth_server.AccessTokenCheckResponse, error) {
	return nil, nil
}

func (gh *authGrpcHandler) RoleCount(ctx context.Context, request *go_oauth_server.RoleCountRequest) (*go_oauth_server.RoleCountResponse, error) {
	return nil, nil
}
