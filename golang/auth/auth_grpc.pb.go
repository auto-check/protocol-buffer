// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	LoginAuth(ctx context.Context, in *LoginAuthRequest, opts ...grpc.CallOption) (*LoginAuthResponse, error)
	GetStudentWithId(ctx context.Context, in *GetStudentWithIdRequest, opts ...grpc.CallOption) (*GetStudentWithIdResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) LoginAuth(ctx context.Context, in *LoginAuthRequest, opts ...grpc.CallOption) (*LoginAuthResponse, error) {
	out := new(LoginAuthResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/LoginAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetStudentWithId(ctx context.Context, in *GetStudentWithIdRequest, opts ...grpc.CallOption) (*GetStudentWithIdResponse, error) {
	out := new(GetStudentWithIdResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetStudentWithId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	LoginAuth(context.Context, *LoginAuthRequest) (*LoginAuthResponse, error)
	GetStudentWithId(context.Context, *GetStudentWithIdRequest) (*GetStudentWithIdResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) LoginAuth(context.Context, *LoginAuthRequest) (*LoginAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAuth not implemented")
}
func (UnimplementedAuthServer) GetStudentWithId(context.Context, *GetStudentWithIdRequest) (*GetStudentWithIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentWithId not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_LoginAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/LoginAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginAuth(ctx, req.(*LoginAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetStudentWithId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentWithIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetStudentWithId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetStudentWithId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetStudentWithId(ctx, req.(*GetStudentWithIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginAuth",
			Handler:    _Auth_LoginAuth_Handler,
		},
		{
			MethodName: "GetStudentWithId",
			Handler:    _Auth_GetStudentWithId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
