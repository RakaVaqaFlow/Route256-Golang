// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/service.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	// Очистка всех данных в базе
	//
	// Безвозвратное удаление всей пользовательской информации из базы данных.
	Clear(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Получение инфомарции о базе данных
	//
	// Получение инфомарции о базе данных.
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) Clear(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/github.moguchev.BD_Forum.api.Admin/Clear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/github.moguchev.BD_Forum.api.Admin/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	// Очистка всех данных в базе
	//
	// Безвозвратное удаление всей пользовательской информации из базы данных.
	Clear(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// Получение инфомарции о базе данных
	//
	// Получение инфомарции о базе данных.
	Status(context.Context, *emptypb.Empty) (*StatusResponse, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) Clear(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedAdminServer) Status(context.Context, *emptypb.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.BD_Forum.api.Admin/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Clear(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.moguchev.BD_Forum.api.Admin/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.moguchev.BD_Forum.api.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Clear",
			Handler:    _Admin_Clear_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Admin_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}
