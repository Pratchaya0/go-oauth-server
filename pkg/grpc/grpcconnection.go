package pkggrpc

import (
	"errors"
	"log"
	"net"

	"github.com/Pratchaya0/go-oauth-server/config"
	authPb "github.com/Pratchaya0/go-oauth-server/modules/auth/authProtobuf"
	userPb "github.com/Pratchaya0/go-oauth-server/modules/user/userProtobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	IGrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		User() userPb.UserGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}

	grpcAuth struct{}
)

func (gcf *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(gcf.client)
}

func (gcf *grpcClientFactory) User() userPb.UserGrpcServiceClient {
	return userPb.NewUserGrpcServiceClient(gcf.client)
}

func NewGrpcClient(host string) (IGrpcClientFactoryHandler, error) {
	// need more certificate imprementation
	option := make([]grpc.DialOption, 0)

	option = append(option, grpc.WithTransportCredentials(insecure.NewCredentials()))

	cns, err := grpc.NewClient(host, option...)

	if err != nil {
		log.Fatalf("Error: Grpc client connection failed: %s", err.Error())
		return nil, errors.New("error: grpc client connection failed")
	}

	return &grpcClientFactory{client: cns}, nil
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	option := make([]grpc.ServerOption, 0)

	server := grpc.NewServer(option...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}

	return server, lis
}
