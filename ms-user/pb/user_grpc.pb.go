// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_Register_FullMethodName              = "/user.UserService/Register"
	UserService_Login_FullMethodName                 = "/user.UserService/Login"
	UserService_GetCustomer_FullMethodName           = "/user.UserService/GetCustomer"
	UserService_UpdateCustomer_FullMethodName        = "/user.UserService/UpdateCustomer"
	UserService_DeleteCustomer_FullMethodName        = "/user.UserService/DeleteCustomer"
	UserService_AddAddressCustomer_FullMethodName    = "/user.UserService/AddAddressCustomer"
	UserService_UpdateAddressCustomer_FullMethodName = "/user.UserService/UpdateAddressCustomer"
	UserService_GetCustomerAddress_FullMethodName    = "/user.UserService/GetCustomerAddress"
	UserService_GetCustomerAdmin_FullMethodName      = "/user.UserService/GetCustomerAdmin"
	UserService_GetAllCustomerAdmin_FullMethodName   = "/user.UserService/GetAllCustomerAdmin"
	UserService_UpdateCustomerAdmin_FullMethodName   = "/user.UserService/UpdateCustomerAdmin"
	UserService_DeleteCustomerAdmin_FullMethodName   = "/user.UserService/DeleteCustomerAdmin"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// user
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// customer
	GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*GetCustomerResponse, error)
	UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*UpdateCustomerResponse, error)
	DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*Empty, error)
	AddAddressCustomer(ctx context.Context, in *AddAddressCustomerRequest, opts ...grpc.CallOption) (*AddAddressCustomerResponse, error)
	UpdateAddressCustomer(ctx context.Context, in *UpdateAddressCustomerRequest, opts ...grpc.CallOption) (*UpdateAddressCustomerResponse, error)
	GetCustomerAddress(ctx context.Context, in *GetCustomerAddressRequest, opts ...grpc.CallOption) (*GetCustomerAddressResponse, error)
	// admin-customer
	GetCustomerAdmin(ctx context.Context, in *GetCustomerAdminRequest, opts ...grpc.CallOption) (*GetCustomerAdminResponse, error)
	GetAllCustomerAdmin(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllCustomerAdminResponse, error)
	UpdateCustomerAdmin(ctx context.Context, in *UpdateCustomerAdminRequest, opts ...grpc.CallOption) (*UpdateCustomerAdminResponse, error)
	DeleteCustomerAdmin(ctx context.Context, in *DeleteCustomerAdminRequest, opts ...grpc.CallOption) (*Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, UserService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*GetCustomerResponse, error) {
	out := new(GetCustomerResponse)
	err := c.cc.Invoke(ctx, UserService_GetCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*UpdateCustomerResponse, error) {
	out := new(UpdateCustomerResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_DeleteCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddAddressCustomer(ctx context.Context, in *AddAddressCustomerRequest, opts ...grpc.CallOption) (*AddAddressCustomerResponse, error) {
	out := new(AddAddressCustomerResponse)
	err := c.cc.Invoke(ctx, UserService_AddAddressCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateAddressCustomer(ctx context.Context, in *UpdateAddressCustomerRequest, opts ...grpc.CallOption) (*UpdateAddressCustomerResponse, error) {
	out := new(UpdateAddressCustomerResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateAddressCustomer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCustomerAddress(ctx context.Context, in *GetCustomerAddressRequest, opts ...grpc.CallOption) (*GetCustomerAddressResponse, error) {
	out := new(GetCustomerAddressResponse)
	err := c.cc.Invoke(ctx, UserService_GetCustomerAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCustomerAdmin(ctx context.Context, in *GetCustomerAdminRequest, opts ...grpc.CallOption) (*GetCustomerAdminResponse, error) {
	out := new(GetCustomerAdminResponse)
	err := c.cc.Invoke(ctx, UserService_GetCustomerAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllCustomerAdmin(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllCustomerAdminResponse, error) {
	out := new(GetAllCustomerAdminResponse)
	err := c.cc.Invoke(ctx, UserService_GetAllCustomerAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateCustomerAdmin(ctx context.Context, in *UpdateCustomerAdminRequest, opts ...grpc.CallOption) (*UpdateCustomerAdminResponse, error) {
	out := new(UpdateCustomerAdminResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateCustomerAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteCustomerAdmin(ctx context.Context, in *DeleteCustomerAdminRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_DeleteCustomerAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// user
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// customer
	GetCustomer(context.Context, *GetCustomerRequest) (*GetCustomerResponse, error)
	UpdateCustomer(context.Context, *UpdateCustomerRequest) (*UpdateCustomerResponse, error)
	DeleteCustomer(context.Context, *DeleteCustomerRequest) (*Empty, error)
	AddAddressCustomer(context.Context, *AddAddressCustomerRequest) (*AddAddressCustomerResponse, error)
	UpdateAddressCustomer(context.Context, *UpdateAddressCustomerRequest) (*UpdateAddressCustomerResponse, error)
	GetCustomerAddress(context.Context, *GetCustomerAddressRequest) (*GetCustomerAddressResponse, error)
	// admin-customer
	GetCustomerAdmin(context.Context, *GetCustomerAdminRequest) (*GetCustomerAdminResponse, error)
	GetAllCustomerAdmin(context.Context, *Empty) (*GetAllCustomerAdminResponse, error)
	UpdateCustomerAdmin(context.Context, *UpdateCustomerAdminRequest) (*UpdateCustomerAdminResponse, error)
	DeleteCustomerAdmin(context.Context, *DeleteCustomerAdminRequest) (*Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) GetCustomer(context.Context, *GetCustomerRequest) (*GetCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (UnimplementedUserServiceServer) UpdateCustomer(context.Context, *UpdateCustomerRequest) (*UpdateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedUserServiceServer) DeleteCustomer(context.Context, *DeleteCustomerRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
func (UnimplementedUserServiceServer) AddAddressCustomer(context.Context, *AddAddressCustomerRequest) (*AddAddressCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAddressCustomer not implemented")
}
func (UnimplementedUserServiceServer) UpdateAddressCustomer(context.Context, *UpdateAddressCustomerRequest) (*UpdateAddressCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddressCustomer not implemented")
}
func (UnimplementedUserServiceServer) GetCustomerAddress(context.Context, *GetCustomerAddressRequest) (*GetCustomerAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerAddress not implemented")
}
func (UnimplementedUserServiceServer) GetCustomerAdmin(context.Context, *GetCustomerAdminRequest) (*GetCustomerAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerAdmin not implemented")
}
func (UnimplementedUserServiceServer) GetAllCustomerAdmin(context.Context, *Empty) (*GetAllCustomerAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCustomerAdmin not implemented")
}
func (UnimplementedUserServiceServer) UpdateCustomerAdmin(context.Context, *UpdateCustomerAdminRequest) (*UpdateCustomerAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerAdmin not implemented")
}
func (UnimplementedUserServiceServer) DeleteCustomerAdmin(context.Context, *DeleteCustomerAdminRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomerAdmin not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCustomer(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateCustomer(ctx, req.(*UpdateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteCustomer(ctx, req.(*DeleteCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddAddressCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAddressCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddAddressCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddAddressCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddAddressCustomer(ctx, req.(*AddAddressCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateAddressCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddressCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateAddressCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateAddressCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateAddressCustomer(ctx, req.(*UpdateAddressCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCustomerAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCustomerAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetCustomerAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCustomerAddress(ctx, req.(*GetCustomerAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCustomerAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCustomerAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetCustomerAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCustomerAdmin(ctx, req.(*GetCustomerAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllCustomerAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllCustomerAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAllCustomerAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllCustomerAdmin(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateCustomerAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateCustomerAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateCustomerAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateCustomerAdmin(ctx, req.(*UpdateCustomerAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteCustomerAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteCustomerAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteCustomerAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteCustomerAdmin(ctx, req.(*DeleteCustomerAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "GetCustomer",
			Handler:    _UserService_GetCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _UserService_UpdateCustomer_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _UserService_DeleteCustomer_Handler,
		},
		{
			MethodName: "AddAddressCustomer",
			Handler:    _UserService_AddAddressCustomer_Handler,
		},
		{
			MethodName: "UpdateAddressCustomer",
			Handler:    _UserService_UpdateAddressCustomer_Handler,
		},
		{
			MethodName: "GetCustomerAddress",
			Handler:    _UserService_GetCustomerAddress_Handler,
		},
		{
			MethodName: "GetCustomerAdmin",
			Handler:    _UserService_GetCustomerAdmin_Handler,
		},
		{
			MethodName: "GetAllCustomerAdmin",
			Handler:    _UserService_GetAllCustomerAdmin_Handler,
		},
		{
			MethodName: "UpdateCustomerAdmin",
			Handler:    _UserService_UpdateCustomerAdmin_Handler,
		},
		{
			MethodName: "DeleteCustomerAdmin",
			Handler:    _UserService_DeleteCustomerAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
