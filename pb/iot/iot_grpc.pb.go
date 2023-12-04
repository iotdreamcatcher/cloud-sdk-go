// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: iot.proto

package iot

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
	IotHomeService_Create_FullMethodName    = "/iot.IotHomeService/Create"
	IotHomeService_Update_FullMethodName    = "/iot.IotHomeService/Update"
	IotHomeService_Delete_FullMethodName    = "/iot.IotHomeService/Delete"
	IotHomeService_DeleteIds_FullMethodName = "/iot.IotHomeService/DeleteIds"
	IotHomeService_Query_FullMethodName     = "/iot.IotHomeService/Query"
	IotHomeService_QueryList_FullMethodName = "/iot.IotHomeService/QueryList"
)

// IotHomeServiceClient is the client API for IotHomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IotHomeServiceClient interface {
	Create(ctx context.Context, in *CreateHomeReq, opts ...grpc.CallOption) (*CreateResp, error)
	Update(ctx context.Context, in *UpdateHomeReq, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Response, error)
	DeleteIds(ctx context.Context, in *DeleteIdsReq, opts ...grpc.CallOption) (*Response, error)
	Query(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryResp, error)
	QueryList(ctx context.Context, in *QueryListReq, opts ...grpc.CallOption) (*QueryListResp, error)
}

type iotHomeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIotHomeServiceClient(cc grpc.ClientConnInterface) IotHomeServiceClient {
	return &iotHomeServiceClient{cc}
}

func (c *iotHomeServiceClient) Create(ctx context.Context, in *CreateHomeReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, IotHomeService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotHomeServiceClient) Update(ctx context.Context, in *UpdateHomeReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, IotHomeService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotHomeServiceClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, IotHomeService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotHomeServiceClient) DeleteIds(ctx context.Context, in *DeleteIdsReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, IotHomeService_DeleteIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotHomeServiceClient) Query(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryResp, error) {
	out := new(QueryResp)
	err := c.cc.Invoke(ctx, IotHomeService_Query_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotHomeServiceClient) QueryList(ctx context.Context, in *QueryListReq, opts ...grpc.CallOption) (*QueryListResp, error) {
	out := new(QueryListResp)
	err := c.cc.Invoke(ctx, IotHomeService_QueryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IotHomeServiceServer is the server API for IotHomeService service.
// All implementations must embed UnimplementedIotHomeServiceServer
// for forward compatibility
type IotHomeServiceServer interface {
	Create(context.Context, *CreateHomeReq) (*CreateResp, error)
	Update(context.Context, *UpdateHomeReq) (*Response, error)
	Delete(context.Context, *DeleteReq) (*Response, error)
	DeleteIds(context.Context, *DeleteIdsReq) (*Response, error)
	Query(context.Context, *QueryReq) (*QueryResp, error)
	QueryList(context.Context, *QueryListReq) (*QueryListResp, error)
	mustEmbedUnimplementedIotHomeServiceServer()
}

// UnimplementedIotHomeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIotHomeServiceServer struct {
}

func (UnimplementedIotHomeServiceServer) Create(context.Context, *CreateHomeReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIotHomeServiceServer) Update(context.Context, *UpdateHomeReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedIotHomeServiceServer) Delete(context.Context, *DeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedIotHomeServiceServer) DeleteIds(context.Context, *DeleteIdsReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIds not implemented")
}
func (UnimplementedIotHomeServiceServer) Query(context.Context, *QueryReq) (*QueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedIotHomeServiceServer) QueryList(context.Context, *QueryListReq) (*QueryListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryList not implemented")
}
func (UnimplementedIotHomeServiceServer) mustEmbedUnimplementedIotHomeServiceServer() {}

// UnsafeIotHomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IotHomeServiceServer will
// result in compilation errors.
type UnsafeIotHomeServiceServer interface {
	mustEmbedUnimplementedIotHomeServiceServer()
}

func RegisterIotHomeServiceServer(s grpc.ServiceRegistrar, srv IotHomeServiceServer) {
	s.RegisterService(&IotHomeService_ServiceDesc, srv)
}

func _IotHomeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHomeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IotHomeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IotHomeService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IotHomeServiceServer).Create(ctx, req.(*CreateHomeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IotHomeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHomeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IotHomeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IotHomeService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IotHomeServiceServer).Update(ctx, req.(*UpdateHomeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IotHomeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IotHomeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IotHomeService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IotHomeServiceServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IotHomeService_DeleteIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IotHomeServiceServer).DeleteIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IotHomeService_DeleteIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IotHomeServiceServer).DeleteIds(ctx, req.(*DeleteIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IotHomeService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IotHomeServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IotHomeService_Query_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IotHomeServiceServer).Query(ctx, req.(*QueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IotHomeService_QueryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IotHomeServiceServer).QueryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IotHomeService_QueryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IotHomeServiceServer).QueryList(ctx, req.(*QueryListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IotHomeService_ServiceDesc is the grpc.ServiceDesc for IotHomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IotHomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iot.IotHomeService",
	HandlerType: (*IotHomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _IotHomeService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _IotHomeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _IotHomeService_Delete_Handler,
		},
		{
			MethodName: "DeleteIds",
			Handler:    _IotHomeService_DeleteIds_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _IotHomeService_Query_Handler,
		},
		{
			MethodName: "QueryList",
			Handler:    _IotHomeService_QueryList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iot.proto",
}