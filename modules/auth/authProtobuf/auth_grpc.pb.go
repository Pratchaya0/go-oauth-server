// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: modules/auth/authProtobuf/auth.proto

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
	AuthGrpcService_AccessTokenCheck_FullMethodName = "/AuthGrpcService/AccessTokenCheck"
	AuthGrpcService_RoleCount_FullMethodName        = "/AuthGrpcService/RoleCount"
)

// AuthGrpcServiceClient is the client API for AuthGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthGrpcServiceClient interface {
	AccessTokenCheck(ctx context.Context, in *AccessTokenCheckRequest, opts ...grpc.CallOption) (*AccessTokenCheckResponse, error)
	RoleCount(ctx context.Context, in *RoleCountRequest, opts ...grpc.CallOption) (*RoleCountResponse, error)
}

type authGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthGrpcServiceClient(cc grpc.ClientConnInterface) AuthGrpcServiceClient {
	return &authGrpcServiceClient{cc}
}

func (c *authGrpcServiceClient) AccessTokenCheck(ctx context.Context, in *AccessTokenCheckRequest, opts ...grpc.CallOption) (*AccessTokenCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AccessTokenCheckResponse)
	err := c.cc.Invoke(ctx, AuthGrpcService_AccessTokenCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authGrpcServiceClient) RoleCount(ctx context.Context, in *RoleCountRequest, opts ...grpc.CallOption) (*RoleCountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RoleCountResponse)
	err := c.cc.Invoke(ctx, AuthGrpcService_RoleCount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthGrpcServiceServer is the server API for AuthGrpcService service.
// All implementations must embed UnimplementedAuthGrpcServiceServer
// for forward compatibility.
type AuthGrpcServiceServer interface {
	AccessTokenCheck(context.Context, *AccessTokenCheckRequest) (*AccessTokenCheckResponse, error)
	RoleCount(context.Context, *RoleCountRequest) (*RoleCountResponse, error)
	mustEmbedUnimplementedAuthGrpcServiceServer()
}

// UnimplementedAuthGrpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthGrpcServiceServer struct{}

func (UnimplementedAuthGrpcServiceServer) AccessTokenCheck(context.Context, *AccessTokenCheckRequest) (*AccessTokenCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessTokenCheck not implemented")
}
func (UnimplementedAuthGrpcServiceServer) RoleCount(context.Context, *RoleCountRequest) (*RoleCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleCount not implemented")
}
func (UnimplementedAuthGrpcServiceServer) mustEmbedUnimplementedAuthGrpcServiceServer() {}
func (UnimplementedAuthGrpcServiceServer) testEmbeddedByValue()                         {}

// UnsafeAuthGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthGrpcServiceServer will
// result in compilation errors.
type UnsafeAuthGrpcServiceServer interface {
	mustEmbedUnimplementedAuthGrpcServiceServer()
}

func RegisterAuthGrpcServiceServer(s grpc.ServiceRegistrar, srv AuthGrpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthGrpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthGrpcService_ServiceDesc, srv)
}

func _AuthGrpcService_AccessTokenCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessTokenCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGrpcServiceServer).AccessTokenCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthGrpcService_AccessTokenCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGrpcServiceServer).AccessTokenCheck(ctx, req.(*AccessTokenCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthGrpcService_RoleCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGrpcServiceServer).RoleCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthGrpcService_RoleCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGrpcServiceServer).RoleCount(ctx, req.(*RoleCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthGrpcService_ServiceDesc is the grpc.ServiceDesc for AuthGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthGrpcService",
	HandlerType: (*AuthGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccessTokenCheck",
			Handler:    _AuthGrpcService_AccessTokenCheck_Handler,
		},
		{
			MethodName: "RoleCount",
			Handler:    _AuthGrpcService_RoleCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/auth/authProtobuf/auth.proto",
}
