package userhandler

import (
	"context"

	go_oauth_server "github.com/Pratchaya0/go-oauth-server/modules/user/userProtobuf"
	userusecase "github.com/Pratchaya0/go-oauth-server/modules/user/userUsecase"
)

type (
	userGrpcHandler struct {
		go_oauth_server.UnimplementedUserGrpcServiceServer
		userUsecase userusecase.IUserUsecase
	}
)

func NewUserGrpcHandler(userUsecase userusecase.IUserUsecase) *userGrpcHandler {
	return &userGrpcHandler{userUsecase: userUsecase}
}

func (gh *userGrpcHandler) CredetialSearch(ctx context.Context, request *go_oauth_server.CredentialSearchRequest) (*go_oauth_server.UserProfile, error) {
	return nil, nil
}

func (gh *userGrpcHandler) FindOneUserProfieToRefresh(ctx context.Context, request *go_oauth_server.FindOneUserProfieToRefreshRequest) (*go_oauth_server.UserProfile, error) {
	return nil, nil
}
