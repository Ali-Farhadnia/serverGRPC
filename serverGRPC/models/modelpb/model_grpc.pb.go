// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package modelpb

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

// CRUDClient is the client API for CRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDClient interface {
	InsertBook(ctx context.Context, in *Books, opts ...grpc.CallOption) (*Status, error)
	FindBookById(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*FindResponse, error)
	UpdateBook(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Status, error)
	DeleteBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*Status, error)
}

type cRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDClient(cc grpc.ClientConnInterface) CRUDClient {
	return &cRUDClient{cc}
}

func (c *cRUDClient) InsertBook(ctx context.Context, in *Books, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/model.CRUD/InsertBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) FindBookById(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/model.CRUD/FindBookById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) UpdateBook(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/model.CRUD/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) DeleteBook(ctx context.Context, in *BookID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/model.CRUD/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServer is the server API for CRUD service.
// All implementations must embed UnimplementedCRUDServer
// for forward compatibility
type CRUDServer interface {
	InsertBook(context.Context, *Books) (*Status, error)
	FindBookById(context.Context, *BookID) (*FindResponse, error)
	UpdateBook(context.Context, *UpdateRequest) (*Status, error)
	DeleteBook(context.Context, *BookID) (*Status, error)
	mustEmbedUnimplementedCRUDServer()
}

// UnimplementedCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedCRUDServer struct {
}

func (UnimplementedCRUDServer) InsertBook(context.Context, *Books) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertBook not implemented")
}
func (UnimplementedCRUDServer) FindBookById(context.Context, *BookID) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBookById not implemented")
}
func (UnimplementedCRUDServer) UpdateBook(context.Context, *UpdateRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedCRUDServer) DeleteBook(context.Context, *BookID) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedCRUDServer) mustEmbedUnimplementedCRUDServer() {}

// UnsafeCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CRUDServer will
// result in compilation errors.
type UnsafeCRUDServer interface {
	mustEmbedUnimplementedCRUDServer()
}

func RegisterCRUDServer(s grpc.ServiceRegistrar, srv CRUDServer) {
	s.RegisterService(&CRUD_ServiceDesc, srv)
}

func _CRUD_InsertBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Books)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).InsertBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.CRUD/InsertBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).InsertBook(ctx, req.(*Books))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_FindBookById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).FindBookById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.CRUD/FindBookById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).FindBookById(ctx, req.(*BookID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.CRUD/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).UpdateBook(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.CRUD/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).DeleteBook(ctx, req.(*BookID))
	}
	return interceptor(ctx, in, info, handler)
}

// CRUD_ServiceDesc is the grpc.ServiceDesc for CRUD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CRUD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.CRUD",
	HandlerType: (*CRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertBook",
			Handler:    _CRUD_InsertBook_Handler,
		},
		{
			MethodName: "FindBookById",
			Handler:    _CRUD_FindBookById_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _CRUD_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _CRUD_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "models/modelpb/model.proto",
}
