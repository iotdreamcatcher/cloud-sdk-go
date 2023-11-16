// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: wechat.proto

package wechat

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
	WechatRpcService_Code2Token_FullMethodName                 = "/wechat.WechatRpcService/Code2Token"
	WechatRpcService_RefreshUserToken_FullMethodName           = "/wechat.WechatRpcService/RefreshUserToken"
	WechatRpcService_UserToken2UserInfo_FullMethodName         = "/wechat.WechatRpcService/UserToken2UserInfo"
	WechatRpcService_WebRedirectWechat_FullMethodName          = "/wechat.WechatRpcService/WebRedirectWechat"
	WechatRpcService_OfficialAccountAccessToken_FullMethodName = "/wechat.WechatRpcService/OfficialAccountAccessToken"
)

// WechatRpcServiceClient is the client API for WechatRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WechatRpcServiceClient interface {
	// note: common
	Code2Token(ctx context.Context, in *WebCodeReq, opts ...grpc.CallOption) (*WebCodeResp, error)
	RefreshUserToken(ctx context.Context, in *WebRefreshReq, opts ...grpc.CallOption) (*WebRefreshResp, error)
	UserToken2UserInfo(ctx context.Context, in *WebTokenReq, opts ...grpc.CallOption) (*WebUserInfoResp, error)
	// note: web
	WebRedirectWechat(ctx context.Context, in *WebRedirectReq, opts ...grpc.CallOption) (*WebRedirectResp, error)
	// note: base OfficialAccount 公众号基础能力
	OfficialAccountAccessToken(ctx context.Context, in *OaKeyReq, opts ...grpc.CallOption) (*OaAccessTokenResp, error)
}

type wechatRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWechatRpcServiceClient(cc grpc.ClientConnInterface) WechatRpcServiceClient {
	return &wechatRpcServiceClient{cc}
}

func (c *wechatRpcServiceClient) Code2Token(ctx context.Context, in *WebCodeReq, opts ...grpc.CallOption) (*WebCodeResp, error) {
	out := new(WebCodeResp)
	err := c.cc.Invoke(ctx, WechatRpcService_Code2Token_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatRpcServiceClient) RefreshUserToken(ctx context.Context, in *WebRefreshReq, opts ...grpc.CallOption) (*WebRefreshResp, error) {
	out := new(WebRefreshResp)
	err := c.cc.Invoke(ctx, WechatRpcService_RefreshUserToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatRpcServiceClient) UserToken2UserInfo(ctx context.Context, in *WebTokenReq, opts ...grpc.CallOption) (*WebUserInfoResp, error) {
	out := new(WebUserInfoResp)
	err := c.cc.Invoke(ctx, WechatRpcService_UserToken2UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatRpcServiceClient) WebRedirectWechat(ctx context.Context, in *WebRedirectReq, opts ...grpc.CallOption) (*WebRedirectResp, error) {
	out := new(WebRedirectResp)
	err := c.cc.Invoke(ctx, WechatRpcService_WebRedirectWechat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wechatRpcServiceClient) OfficialAccountAccessToken(ctx context.Context, in *OaKeyReq, opts ...grpc.CallOption) (*OaAccessTokenResp, error) {
	out := new(OaAccessTokenResp)
	err := c.cc.Invoke(ctx, WechatRpcService_OfficialAccountAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WechatRpcServiceServer is the server API for WechatRpcService service.
// All implementations must embed UnimplementedWechatRpcServiceServer
// for forward compatibility
type WechatRpcServiceServer interface {
	// note: common
	Code2Token(context.Context, *WebCodeReq) (*WebCodeResp, error)
	RefreshUserToken(context.Context, *WebRefreshReq) (*WebRefreshResp, error)
	UserToken2UserInfo(context.Context, *WebTokenReq) (*WebUserInfoResp, error)
	// note: web
	WebRedirectWechat(context.Context, *WebRedirectReq) (*WebRedirectResp, error)
	// note: base OfficialAccount 公众号基础能力
	OfficialAccountAccessToken(context.Context, *OaKeyReq) (*OaAccessTokenResp, error)
	mustEmbedUnimplementedWechatRpcServiceServer()
}

// UnimplementedWechatRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWechatRpcServiceServer struct {
}

func (UnimplementedWechatRpcServiceServer) Code2Token(context.Context, *WebCodeReq) (*WebCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Code2Token not implemented")
}
func (UnimplementedWechatRpcServiceServer) RefreshUserToken(context.Context, *WebRefreshReq) (*WebRefreshResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshUserToken not implemented")
}
func (UnimplementedWechatRpcServiceServer) UserToken2UserInfo(context.Context, *WebTokenReq) (*WebUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserToken2UserInfo not implemented")
}
func (UnimplementedWechatRpcServiceServer) WebRedirectWechat(context.Context, *WebRedirectReq) (*WebRedirectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebRedirectWechat not implemented")
}
func (UnimplementedWechatRpcServiceServer) OfficialAccountAccessToken(context.Context, *OaKeyReq) (*OaAccessTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OfficialAccountAccessToken not implemented")
}
func (UnimplementedWechatRpcServiceServer) mustEmbedUnimplementedWechatRpcServiceServer() {}

// UnsafeWechatRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WechatRpcServiceServer will
// result in compilation errors.
type UnsafeWechatRpcServiceServer interface {
	mustEmbedUnimplementedWechatRpcServiceServer()
}

func RegisterWechatRpcServiceServer(s grpc.ServiceRegistrar, srv WechatRpcServiceServer) {
	s.RegisterService(&WechatRpcService_ServiceDesc, srv)
}

func _WechatRpcService_Code2Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatRpcServiceServer).Code2Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WechatRpcService_Code2Token_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatRpcServiceServer).Code2Token(ctx, req.(*WebCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatRpcService_RefreshUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatRpcServiceServer).RefreshUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WechatRpcService_RefreshUserToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatRpcServiceServer).RefreshUserToken(ctx, req.(*WebRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatRpcService_UserToken2UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatRpcServiceServer).UserToken2UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WechatRpcService_UserToken2UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatRpcServiceServer).UserToken2UserInfo(ctx, req.(*WebTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatRpcService_WebRedirectWechat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebRedirectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatRpcServiceServer).WebRedirectWechat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WechatRpcService_WebRedirectWechat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatRpcServiceServer).WebRedirectWechat(ctx, req.(*WebRedirectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WechatRpcService_OfficialAccountAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OaKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WechatRpcServiceServer).OfficialAccountAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WechatRpcService_OfficialAccountAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WechatRpcServiceServer).OfficialAccountAccessToken(ctx, req.(*OaKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WechatRpcService_ServiceDesc is the grpc.ServiceDesc for WechatRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WechatRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wechat.WechatRpcService",
	HandlerType: (*WechatRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Code2Token",
			Handler:    _WechatRpcService_Code2Token_Handler,
		},
		{
			MethodName: "RefreshUserToken",
			Handler:    _WechatRpcService_RefreshUserToken_Handler,
		},
		{
			MethodName: "UserToken2UserInfo",
			Handler:    _WechatRpcService_UserToken2UserInfo_Handler,
		},
		{
			MethodName: "WebRedirectWechat",
			Handler:    _WechatRpcService_WebRedirectWechat_Handler,
		},
		{
			MethodName: "OfficialAccountAccessToken",
			Handler:    _WechatRpcService_OfficialAccountAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wechat.proto",
}
