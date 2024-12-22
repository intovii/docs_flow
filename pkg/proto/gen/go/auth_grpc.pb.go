// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: auth.proto

package auth_protov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Auth_GetUserToken_FullMethodName             = "/auth_proto.v1.auth/GetUserToken"
	Auth_RegisterUser_FullMethodName             = "/auth_proto.v1.auth/RegisterUser"
	Auth_CreateCourse_FullMethodName             = "/auth_proto.v1.auth/CreateCourse"
	Auth_UpdateCourse_FullMethodName             = "/auth_proto.v1.auth/UpdateCourse"
	Auth_GetCourse_FullMethodName                = "/auth_proto.v1.auth/GetCourse"
	Auth_GetAllCoursesCatalog_FullMethodName     = "/auth_proto.v1.auth/GetAllCoursesCatalog"
	Auth_DeleteCourse_FullMethodName             = "/auth_proto.v1.auth/DeleteCourse"
	Auth_CreateUserCoursePayment_FullMethodName  = "/auth_proto.v1.auth/CreateUserCoursePayment"
	Auth_CreateCourseContentMulti_FullMethodName = "/auth_proto.v1.auth/CreateCourseContentMulti"
	Auth_GetCourseContent_FullMethodName         = "/auth_proto.v1.auth/GetCourseContent"
	Auth_UpdateCourseContent_FullMethodName      = "/auth_proto.v1.auth/UpdateCourseContent"
	Auth_DeleteCourseContent_FullMethodName      = "/auth_proto.v1.auth/DeleteCourseContent"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	GetUserToken(ctx context.Context, in *GetUserTokenRequest, opts ...grpc.CallOption) (*GetUserTokenResponse, error)
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CreateCourseResponse, error)
	UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*UpdateCourseResponse, error)
	GetCourse(ctx context.Context, in *GetCourseRequest, opts ...grpc.CallOption) (*GetCourseResponse, error)
	GetAllCoursesCatalog(ctx context.Context, in *GetAllCoursesCatalogRequest, opts ...grpc.CallOption) (*GetAllCoursesCatalogResponse, error)
	DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error)
	CreateUserCoursePayment(ctx context.Context, in *CreateUserCoursePaymentRequest, opts ...grpc.CallOption) (*CreateUserCoursePaymentResponse, error)
	CreateCourseContentMulti(ctx context.Context, in *CreateCourseContentMultiRequest, opts ...grpc.CallOption) (*CreateCourseContentMultiResponse, error)
	GetCourseContent(ctx context.Context, in *GetCourseContentRequest, opts ...grpc.CallOption) (*GetCourseContentResponse, error)
	UpdateCourseContent(ctx context.Context, in *UpdateCourseContentRequest, opts ...grpc.CallOption) (*UpdateCourseContentResponse, error)
	DeleteCourseContent(ctx context.Context, in *DeleteCourseContentRequest, opts ...grpc.CallOption) (*DeleteCourseContentResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetUserToken(ctx context.Context, in *GetUserTokenRequest, opts ...grpc.CallOption) (*GetUserTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserTokenResponse)
	err := c.cc.Invoke(ctx, Auth_GetUserToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, Auth_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CreateCourseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCourseResponse)
	err := c.cc.Invoke(ctx, Auth_CreateCourse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateCourse(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*UpdateCourseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCourseResponse)
	err := c.cc.Invoke(ctx, Auth_UpdateCourse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetCourse(ctx context.Context, in *GetCourseRequest, opts ...grpc.CallOption) (*GetCourseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCourseResponse)
	err := c.cc.Invoke(ctx, Auth_GetCourse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAllCoursesCatalog(ctx context.Context, in *GetAllCoursesCatalogRequest, opts ...grpc.CallOption) (*GetAllCoursesCatalogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllCoursesCatalogResponse)
	err := c.cc.Invoke(ctx, Auth_GetAllCoursesCatalog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteCourse(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCourseResponse)
	err := c.cc.Invoke(ctx, Auth_DeleteCourse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateUserCoursePayment(ctx context.Context, in *CreateUserCoursePaymentRequest, opts ...grpc.CallOption) (*CreateUserCoursePaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserCoursePaymentResponse)
	err := c.cc.Invoke(ctx, Auth_CreateUserCoursePayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateCourseContentMulti(ctx context.Context, in *CreateCourseContentMultiRequest, opts ...grpc.CallOption) (*CreateCourseContentMultiResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCourseContentMultiResponse)
	err := c.cc.Invoke(ctx, Auth_CreateCourseContentMulti_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetCourseContent(ctx context.Context, in *GetCourseContentRequest, opts ...grpc.CallOption) (*GetCourseContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCourseContentResponse)
	err := c.cc.Invoke(ctx, Auth_GetCourseContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateCourseContent(ctx context.Context, in *UpdateCourseContentRequest, opts ...grpc.CallOption) (*UpdateCourseContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCourseContentResponse)
	err := c.cc.Invoke(ctx, Auth_UpdateCourseContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteCourseContent(ctx context.Context, in *DeleteCourseContentRequest, opts ...grpc.CallOption) (*DeleteCourseContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCourseContentResponse)
	err := c.cc.Invoke(ctx, Auth_DeleteCourseContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility.
type AuthServer interface {
	GetUserToken(context.Context, *GetUserTokenRequest) (*GetUserTokenResponse, error)
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	CreateCourse(context.Context, *CreateCourseRequest) (*CreateCourseResponse, error)
	UpdateCourse(context.Context, *UpdateCourseRequest) (*UpdateCourseResponse, error)
	GetCourse(context.Context, *GetCourseRequest) (*GetCourseResponse, error)
	GetAllCoursesCatalog(context.Context, *GetAllCoursesCatalogRequest) (*GetAllCoursesCatalogResponse, error)
	DeleteCourse(context.Context, *DeleteCourseRequest) (*DeleteCourseResponse, error)
	CreateUserCoursePayment(context.Context, *CreateUserCoursePaymentRequest) (*CreateUserCoursePaymentResponse, error)
	CreateCourseContentMulti(context.Context, *CreateCourseContentMultiRequest) (*CreateCourseContentMultiResponse, error)
	GetCourseContent(context.Context, *GetCourseContentRequest) (*GetCourseContentResponse, error)
	UpdateCourseContent(context.Context, *UpdateCourseContentRequest) (*UpdateCourseContentResponse, error)
	DeleteCourseContent(context.Context, *DeleteCourseContentRequest) (*DeleteCourseContentResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServer struct{}

func (UnimplementedAuthServer) GetUserToken(context.Context, *GetUserTokenRequest) (*GetUserTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserToken not implemented")
}
func (UnimplementedAuthServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAuthServer) CreateCourse(context.Context, *CreateCourseRequest) (*CreateCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedAuthServer) UpdateCourse(context.Context, *UpdateCourseRequest) (*UpdateCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourse not implemented")
}
func (UnimplementedAuthServer) GetCourse(context.Context, *GetCourseRequest) (*GetCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedAuthServer) GetAllCoursesCatalog(context.Context, *GetAllCoursesCatalogRequest) (*GetAllCoursesCatalogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCoursesCatalog not implemented")
}
func (UnimplementedAuthServer) DeleteCourse(context.Context, *DeleteCourseRequest) (*DeleteCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedAuthServer) CreateUserCoursePayment(context.Context, *CreateUserCoursePaymentRequest) (*CreateUserCoursePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserCoursePayment not implemented")
}
func (UnimplementedAuthServer) CreateCourseContentMulti(context.Context, *CreateCourseContentMultiRequest) (*CreateCourseContentMultiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourseContentMulti not implemented")
}
func (UnimplementedAuthServer) GetCourseContent(context.Context, *GetCourseContentRequest) (*GetCourseContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourseContent not implemented")
}
func (UnimplementedAuthServer) UpdateCourseContent(context.Context, *UpdateCourseContentRequest) (*UpdateCourseContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourseContent not implemented")
}
func (UnimplementedAuthServer) DeleteCourseContent(context.Context, *DeleteCourseContentRequest) (*DeleteCourseContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourseContent not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}
func (UnimplementedAuthServer) testEmbeddedByValue()              {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	// If the following call pancis, it indicates UnimplementedAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_GetUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetUserToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserToken(ctx, req.(*GetUserTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_CreateCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateCourse(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UpdateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UpdateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_UpdateCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UpdateCourse(ctx, req.(*UpdateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetCourse(ctx, req.(*GetCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetAllCoursesCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCoursesCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetAllCoursesCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetAllCoursesCatalog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetAllCoursesCatalog(ctx, req.(*GetAllCoursesCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DeleteCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteCourse(ctx, req.(*DeleteCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateUserCoursePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserCoursePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateUserCoursePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_CreateUserCoursePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateUserCoursePayment(ctx, req.(*CreateUserCoursePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateCourseContentMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseContentMultiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateCourseContentMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_CreateCourseContentMulti_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateCourseContentMulti(ctx, req.(*CreateCourseContentMultiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetCourseContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetCourseContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetCourseContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetCourseContent(ctx, req.(*GetCourseContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UpdateCourseContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourseContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UpdateCourseContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_UpdateCourseContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UpdateCourseContent(ctx, req.(*UpdateCourseContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteCourseContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourseContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteCourseContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DeleteCourseContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteCourseContent(ctx, req.(*DeleteCourseContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_proto.v1.auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserToken",
			Handler:    _Auth_GetUserToken_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _Auth_RegisterUser_Handler,
		},
		{
			MethodName: "CreateCourse",
			Handler:    _Auth_CreateCourse_Handler,
		},
		{
			MethodName: "UpdateCourse",
			Handler:    _Auth_UpdateCourse_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _Auth_GetCourse_Handler,
		},
		{
			MethodName: "GetAllCoursesCatalog",
			Handler:    _Auth_GetAllCoursesCatalog_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _Auth_DeleteCourse_Handler,
		},
		{
			MethodName: "CreateUserCoursePayment",
			Handler:    _Auth_CreateUserCoursePayment_Handler,
		},
		{
			MethodName: "CreateCourseContentMulti",
			Handler:    _Auth_CreateCourseContentMulti_Handler,
		},
		{
			MethodName: "GetCourseContent",
			Handler:    _Auth_GetCourseContent_Handler,
		},
		{
			MethodName: "UpdateCourseContent",
			Handler:    _Auth_UpdateCourseContent_Handler,
		},
		{
			MethodName: "DeleteCourseContent",
			Handler:    _Auth_DeleteCourseContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
