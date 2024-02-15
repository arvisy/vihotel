// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: hotel.proto

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
	HotelService_CreateHotel_FullMethodName = "/user.HotelService/CreateHotel"
	HotelService_UpdateHotel_FullMethodName = "/user.HotelService/UpdateHotel"
	HotelService_DeleteHotel_FullMethodName = "/user.HotelService/DeleteHotel"
	HotelService_GetHotel_FullMethodName    = "/user.HotelService/GetHotel"
	HotelService_ListHotels_FullMethodName  = "/user.HotelService/ListHotels"
	HotelService_CreateRoom_FullMethodName  = "/user.HotelService/CreateRoom"
	HotelService_UpdateRoom_FullMethodName  = "/user.HotelService/UpdateRoom"
	HotelService_DeleteRoom_FullMethodName  = "/user.HotelService/DeleteRoom"
	HotelService_GetRoom_FullMethodName     = "/user.HotelService/GetRoom"
	HotelService_ListRooms_FullMethodName   = "/user.HotelService/ListRooms"
)

// HotelServiceClient is the client API for HotelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HotelServiceClient interface {
	// hotel operations
	CreateHotel(ctx context.Context, in *CreateHotelRequest, opts ...grpc.CallOption) (*CreateHotelResponse, error)
	UpdateHotel(ctx context.Context, in *UpdateHotelRequest, opts ...grpc.CallOption) (*UpdateHotelResponse, error)
	DeleteHotel(ctx context.Context, in *DeleteHotelRequest, opts ...grpc.CallOption) (*DeleteHotelResponse, error)
	GetHotel(ctx context.Context, in *GetHotelRequest, opts ...grpc.CallOption) (*GetHotelResponse, error)
	ListHotels(ctx context.Context, in *ListHotelsRequest, opts ...grpc.CallOption) (*ListHotelsResponse, error)
	// room operations
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
	UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomResponse, error)
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error)
	GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error)
	ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...grpc.CallOption) (*ListRoomsResponse, error)
}

type hotelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHotelServiceClient(cc grpc.ClientConnInterface) HotelServiceClient {
	return &hotelServiceClient{cc}
}

func (c *hotelServiceClient) CreateHotel(ctx context.Context, in *CreateHotelRequest, opts ...grpc.CallOption) (*CreateHotelResponse, error) {
	out := new(CreateHotelResponse)
	err := c.cc.Invoke(ctx, HotelService_CreateHotel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) UpdateHotel(ctx context.Context, in *UpdateHotelRequest, opts ...grpc.CallOption) (*UpdateHotelResponse, error) {
	out := new(UpdateHotelResponse)
	err := c.cc.Invoke(ctx, HotelService_UpdateHotel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) DeleteHotel(ctx context.Context, in *DeleteHotelRequest, opts ...grpc.CallOption) (*DeleteHotelResponse, error) {
	out := new(DeleteHotelResponse)
	err := c.cc.Invoke(ctx, HotelService_DeleteHotel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetHotel(ctx context.Context, in *GetHotelRequest, opts ...grpc.CallOption) (*GetHotelResponse, error) {
	out := new(GetHotelResponse)
	err := c.cc.Invoke(ctx, HotelService_GetHotel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) ListHotels(ctx context.Context, in *ListHotelsRequest, opts ...grpc.CallOption) (*ListHotelsResponse, error) {
	out := new(ListHotelsResponse)
	err := c.cc.Invoke(ctx, HotelService_ListHotels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, HotelService_CreateRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomResponse, error) {
	out := new(UpdateRoomResponse)
	err := c.cc.Invoke(ctx, HotelService_UpdateRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error) {
	out := new(DeleteRoomResponse)
	err := c.cc.Invoke(ctx, HotelService_DeleteRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) GetRoom(ctx context.Context, in *GetRoomRequest, opts ...grpc.CallOption) (*GetRoomResponse, error) {
	out := new(GetRoomResponse)
	err := c.cc.Invoke(ctx, HotelService_GetRoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hotelServiceClient) ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...grpc.CallOption) (*ListRoomsResponse, error) {
	out := new(ListRoomsResponse)
	err := c.cc.Invoke(ctx, HotelService_ListRooms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelServiceServer is the server API for HotelService service.
// All implementations must embed UnimplementedHotelServiceServer
// for forward compatibility
type HotelServiceServer interface {
	// hotel operations
	CreateHotel(context.Context, *CreateHotelRequest) (*CreateHotelResponse, error)
	UpdateHotel(context.Context, *UpdateHotelRequest) (*UpdateHotelResponse, error)
	DeleteHotel(context.Context, *DeleteHotelRequest) (*DeleteHotelResponse, error)
	GetHotel(context.Context, *GetHotelRequest) (*GetHotelResponse, error)
	ListHotels(context.Context, *ListHotelsRequest) (*ListHotelsResponse, error)
	// room operations
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
	UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomResponse, error)
	DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error)
	GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error)
	ListRooms(context.Context, *ListRoomsRequest) (*ListRoomsResponse, error)
	mustEmbedUnimplementedHotelServiceServer()
}

// UnimplementedHotelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHotelServiceServer struct {
}

func (UnimplementedHotelServiceServer) CreateHotel(context.Context, *CreateHotelRequest) (*CreateHotelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHotel not implemented")
}
func (UnimplementedHotelServiceServer) UpdateHotel(context.Context, *UpdateHotelRequest) (*UpdateHotelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHotel not implemented")
}
func (UnimplementedHotelServiceServer) DeleteHotel(context.Context, *DeleteHotelRequest) (*DeleteHotelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHotel not implemented")
}
func (UnimplementedHotelServiceServer) GetHotel(context.Context, *GetHotelRequest) (*GetHotelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotel not implemented")
}
func (UnimplementedHotelServiceServer) ListHotels(context.Context, *ListHotelsRequest) (*ListHotelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHotels not implemented")
}
func (UnimplementedHotelServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedHotelServiceServer) UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}
func (UnimplementedHotelServiceServer) DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedHotelServiceServer) GetRoom(context.Context, *GetRoomRequest) (*GetRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedHotelServiceServer) ListRooms(context.Context, *ListRoomsRequest) (*ListRoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRooms not implemented")
}
func (UnimplementedHotelServiceServer) mustEmbedUnimplementedHotelServiceServer() {}

// UnsafeHotelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HotelServiceServer will
// result in compilation errors.
type UnsafeHotelServiceServer interface {
	mustEmbedUnimplementedHotelServiceServer()
}

func RegisterHotelServiceServer(s grpc.ServiceRegistrar, srv HotelServiceServer) {
	s.RegisterService(&HotelService_ServiceDesc, srv)
}

func _HotelService_CreateHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).CreateHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_CreateHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).CreateHotel(ctx, req.(*CreateHotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_UpdateHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).UpdateHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_UpdateHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).UpdateHotel(ctx, req.(*UpdateHotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_DeleteHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).DeleteHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_DeleteHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).DeleteHotel(ctx, req.(*DeleteHotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetHotel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetHotel(ctx, req.(*GetHotelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_ListHotels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHotelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).ListHotels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_ListHotels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).ListHotels(ctx, req.(*ListHotelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_CreateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_UpdateRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).UpdateRoom(ctx, req.(*UpdateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_DeleteRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).DeleteRoom(ctx, req.(*DeleteRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_GetRoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).GetRoom(ctx, req.(*GetRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HotelService_ListRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelServiceServer).ListRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HotelService_ListRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelServiceServer).ListRooms(ctx, req.(*ListRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HotelService_ServiceDesc is the grpc.ServiceDesc for HotelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HotelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.HotelService",
	HandlerType: (*HotelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHotel",
			Handler:    _HotelService_CreateHotel_Handler,
		},
		{
			MethodName: "UpdateHotel",
			Handler:    _HotelService_UpdateHotel_Handler,
		},
		{
			MethodName: "DeleteHotel",
			Handler:    _HotelService_DeleteHotel_Handler,
		},
		{
			MethodName: "GetHotel",
			Handler:    _HotelService_GetHotel_Handler,
		},
		{
			MethodName: "ListHotels",
			Handler:    _HotelService_ListHotels_Handler,
		},
		{
			MethodName: "CreateRoom",
			Handler:    _HotelService_CreateRoom_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _HotelService_UpdateRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _HotelService_DeleteRoom_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _HotelService_GetRoom_Handler,
		},
		{
			MethodName: "ListRooms",
			Handler:    _HotelService_ListRooms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hotel.proto",
}
