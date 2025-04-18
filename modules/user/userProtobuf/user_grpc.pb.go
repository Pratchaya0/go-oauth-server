// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: modules/user/userProtobuf/user.proto

package go_oauth_server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserGrpcService_CredetialSearch_FullMethodName            = "/UserGrpcService/CredetialSearch"
	UserGrpcService_FindOneUserProfieToRefresh_FullMethodName = "/UserGrpcService/FindOneUserProfieToRefresh"
)

// UserGrpcServiceClient is the client API for UserGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGrpcServiceClient interface {
	CredetialSearch(ctx context.Context, in *CredentialSearchRequest, opts ...grpc.CallOption) (*UserProfile, error)
	FindOneUserProfieToRefresh(ctx context.Context, in *FindOneUserProfieToRefreshRequest, opts ...grpc.CallOption) (*UserProfile, error)
}

type userGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGrpcServiceClient(cc grpc.ClientConnInterface) UserGrpcServiceClient {
	return &userGrpcServiceClient{cc}
}

func (c *userGrpcServiceClient) CredetialSearch(ctx context.Context, in *CredentialSearchRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, UserGrpcService_CredetialSearch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGrpcServiceClient) FindOneUserProfieToRefresh(ctx context.Context, in *FindOneUserProfieToRefreshRequest, opts ...grpc.CallOption) (*UserProfile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserProfile)
	err := c.cc.Invoke(ctx, UserGrpcService_FindOneUserProfieToRefresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGrpcServiceServer is the server API for UserGrpcService service.
// All implementations must embed UnimplementedUserGrpcServiceServer
// for forward compatibility.
type UserGrpcServiceServer interface {
	CredetialSearch(context.Context, *CredentialSearchRequest) (*UserProfile, error)
	FindOneUserProfieToRefresh(context.Context, *FindOneUserProfieToRefreshRequest) (*UserProfile, error)
	mustEmbedUnimplementedUserGrpcServiceServer()
}

// UnimplementedUserGrpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserGrpcServiceServer struct{}

func (UnimplementedUserGrpcServiceServer) CredetialSearch(context.Context, *CredentialSearchRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CredetialSearch not implemented")
}
func (UnimplementedUserGrpcServiceServer) FindOneUserProfieToRefresh(context.Context, *FindOneUserProfieToRefreshRequest) (*UserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneUserProfieToRefresh not implemented")
}
func (UnimplementedUserGrpcServiceServer) mustEmbedUnimplementedUserGrpcServiceServer() {}
func (UnimplementedUserGrpcServiceServer) testEmbeddedByValue()                         {}

// UnsafeUserGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGrpcServiceServer will
// result in compilation errors.
type UnsafeUserGrpcServiceServer interface {
	mustEmbedUnimplementedUserGrpcServiceServer()
}

func RegisterUserGrpcServiceServer(s grpc.ServiceRegistrar, srv UserGrpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserGrpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserGrpcService_ServiceDesc, srv)
}

func _UserGrpcService_CredetialSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).CredetialSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGrpcService_CredetialSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).CredetialSearch(ctx, req.(*CredentialSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGrpcService_FindOneUserProfieToRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneUserProfieToRefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGrpcServiceServer).FindOneUserProfieToRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserGrpcService_FindOneUserProfieToRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGrpcServiceServer).FindOneUserProfieToRefresh(ctx, req.(*FindOneUserProfieToRefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGrpcService_ServiceDesc is the grpc.ServiceDesc for UserGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserGrpcService",
	HandlerType: (*UserGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CredetialSearch",
			Handler:    _UserGrpcService_CredetialSearch_Handler,
		},
		{
			MethodName: "FindOneUserProfieToRefresh",
			Handler:    _UserGrpcService_FindOneUserProfieToRefresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/user/userProtobuf/user.proto",
}
