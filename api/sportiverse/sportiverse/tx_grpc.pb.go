// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sportiverse/sportiverse/tx.proto

package sportiverse

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
	Msg_UpdateParams_FullMethodName       = "/sportiverse.sportiverse.Msg/UpdateParams"
	Msg_CreatePost_FullMethodName         = "/sportiverse.sportiverse.Msg/CreatePost"
	Msg_UpdatePost_FullMethodName         = "/sportiverse.sportiverse.Msg/UpdatePost"
	Msg_DeletePost_FullMethodName         = "/sportiverse.sportiverse.Msg/DeletePost"
	Msg_CreateComment_FullMethodName      = "/sportiverse.sportiverse.Msg/CreateComment"
	Msg_UpdateComment_FullMethodName      = "/sportiverse.sportiverse.Msg/UpdateComment"
	Msg_DeleteComment_FullMethodName      = "/sportiverse.sportiverse.Msg/DeleteComment"
	Msg_CreateSubscription_FullMethodName = "/sportiverse.sportiverse.Msg/CreateSubscription"
	Msg_DeleteSubscription_FullMethodName = "/sportiverse.sportiverse.Msg/DeleteSubscription"
	Msg_CreateLike_FullMethodName         = "/sportiverse.sportiverse.Msg/CreateLike"
	Msg_DeleteLike_FullMethodName         = "/sportiverse.sportiverse.Msg/DeleteLike"
	Msg_CreateAccount_FullMethodName      = "/sportiverse.sportiverse.Msg/CreateAccount"
	Msg_DeleteAccount_FullMethodName      = "/sportiverse.sportiverse.Msg/DeleteAccount"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreatePost(ctx context.Context, in *MsgCreatePost, opts ...grpc.CallOption) (*MsgCreatePostResponse, error)
	UpdatePost(ctx context.Context, in *MsgUpdatePost, opts ...grpc.CallOption) (*MsgUpdatePostResponse, error)
	DeletePost(ctx context.Context, in *MsgDeletePost, opts ...grpc.CallOption) (*MsgDeletePostResponse, error)
	CreateComment(ctx context.Context, in *MsgCreateComment, opts ...grpc.CallOption) (*MsgCreateCommentResponse, error)
	UpdateComment(ctx context.Context, in *MsgUpdateComment, opts ...grpc.CallOption) (*MsgUpdateCommentResponse, error)
	DeleteComment(ctx context.Context, in *MsgDeleteComment, opts ...grpc.CallOption) (*MsgDeleteCommentResponse, error)
	CreateSubscription(ctx context.Context, in *MsgCreateSubscription, opts ...grpc.CallOption) (*MsgCreateSubscriptionResponse, error)
	DeleteSubscription(ctx context.Context, in *MsgDeleteSubscription, opts ...grpc.CallOption) (*MsgDeleteSubscriptionResponse, error)
	CreateLike(ctx context.Context, in *MsgCreateLike, opts ...grpc.CallOption) (*MsgCreateLikeResponse, error)
	DeleteLike(ctx context.Context, in *MsgDeleteLike, opts ...grpc.CallOption) (*MsgDeleteLikeResponse, error)
	CreateAccount(ctx context.Context, in *MsgCreateAccount, opts ...grpc.CallOption) (*MsgCreateAccountResponse, error)
	DeleteAccount(ctx context.Context, in *MsgDeleteAccount, opts ...grpc.CallOption) (*MsgDeleteAccountResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePost(ctx context.Context, in *MsgCreatePost, opts ...grpc.CallOption) (*MsgCreatePostResponse, error) {
	out := new(MsgCreatePostResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePost(ctx context.Context, in *MsgUpdatePost, opts ...grpc.CallOption) (*MsgUpdatePostResponse, error) {
	out := new(MsgUpdatePostResponse)
	err := c.cc.Invoke(ctx, Msg_UpdatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeletePost(ctx context.Context, in *MsgDeletePost, opts ...grpc.CallOption) (*MsgDeletePostResponse, error) {
	out := new(MsgDeletePostResponse)
	err := c.cc.Invoke(ctx, Msg_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateComment(ctx context.Context, in *MsgCreateComment, opts ...grpc.CallOption) (*MsgCreateCommentResponse, error) {
	out := new(MsgCreateCommentResponse)
	err := c.cc.Invoke(ctx, Msg_CreateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateComment(ctx context.Context, in *MsgUpdateComment, opts ...grpc.CallOption) (*MsgUpdateCommentResponse, error) {
	out := new(MsgUpdateCommentResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteComment(ctx context.Context, in *MsgDeleteComment, opts ...grpc.CallOption) (*MsgDeleteCommentResponse, error) {
	out := new(MsgDeleteCommentResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateSubscription(ctx context.Context, in *MsgCreateSubscription, opts ...grpc.CallOption) (*MsgCreateSubscriptionResponse, error) {
	out := new(MsgCreateSubscriptionResponse)
	err := c.cc.Invoke(ctx, Msg_CreateSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteSubscription(ctx context.Context, in *MsgDeleteSubscription, opts ...grpc.CallOption) (*MsgDeleteSubscriptionResponse, error) {
	out := new(MsgDeleteSubscriptionResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateLike(ctx context.Context, in *MsgCreateLike, opts ...grpc.CallOption) (*MsgCreateLikeResponse, error) {
	out := new(MsgCreateLikeResponse)
	err := c.cc.Invoke(ctx, Msg_CreateLike_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteLike(ctx context.Context, in *MsgDeleteLike, opts ...grpc.CallOption) (*MsgDeleteLikeResponse, error) {
	out := new(MsgDeleteLikeResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteLike_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateAccount(ctx context.Context, in *MsgCreateAccount, opts ...grpc.CallOption) (*MsgCreateAccountResponse, error) {
	out := new(MsgCreateAccountResponse)
	err := c.cc.Invoke(ctx, Msg_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteAccount(ctx context.Context, in *MsgDeleteAccount, opts ...grpc.CallOption) (*MsgDeleteAccountResponse, error) {
	out := new(MsgDeleteAccountResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreatePost(context.Context, *MsgCreatePost) (*MsgCreatePostResponse, error)
	UpdatePost(context.Context, *MsgUpdatePost) (*MsgUpdatePostResponse, error)
	DeletePost(context.Context, *MsgDeletePost) (*MsgDeletePostResponse, error)
	CreateComment(context.Context, *MsgCreateComment) (*MsgCreateCommentResponse, error)
	UpdateComment(context.Context, *MsgUpdateComment) (*MsgUpdateCommentResponse, error)
	DeleteComment(context.Context, *MsgDeleteComment) (*MsgDeleteCommentResponse, error)
	CreateSubscription(context.Context, *MsgCreateSubscription) (*MsgCreateSubscriptionResponse, error)
	DeleteSubscription(context.Context, *MsgDeleteSubscription) (*MsgDeleteSubscriptionResponse, error)
	CreateLike(context.Context, *MsgCreateLike) (*MsgCreateLikeResponse, error)
	DeleteLike(context.Context, *MsgDeleteLike) (*MsgDeleteLikeResponse, error)
	CreateAccount(context.Context, *MsgCreateAccount) (*MsgCreateAccountResponse, error)
	DeleteAccount(context.Context, *MsgDeleteAccount) (*MsgDeleteAccountResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreatePost(context.Context, *MsgCreatePost) (*MsgCreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedMsgServer) UpdatePost(context.Context, *MsgUpdatePost) (*MsgUpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedMsgServer) DeletePost(context.Context, *MsgDeletePost) (*MsgDeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedMsgServer) CreateComment(context.Context, *MsgCreateComment) (*MsgCreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedMsgServer) UpdateComment(context.Context, *MsgUpdateComment) (*MsgUpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedMsgServer) DeleteComment(context.Context, *MsgDeleteComment) (*MsgDeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedMsgServer) CreateSubscription(context.Context, *MsgCreateSubscription) (*MsgCreateSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscription not implemented")
}
func (UnimplementedMsgServer) DeleteSubscription(context.Context, *MsgDeleteSubscription) (*MsgDeleteSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}
func (UnimplementedMsgServer) CreateLike(context.Context, *MsgCreateLike) (*MsgCreateLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLike not implemented")
}
func (UnimplementedMsgServer) DeleteLike(context.Context, *MsgDeleteLike) (*MsgDeleteLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLike not implemented")
}
func (UnimplementedMsgServer) CreateAccount(context.Context, *MsgCreateAccount) (*MsgCreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedMsgServer) DeleteAccount(context.Context, *MsgDeleteAccount) (*MsgDeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePost(ctx, req.(*MsgCreatePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePost(ctx, req.(*MsgUpdatePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeletePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeletePost(ctx, req.(*MsgDeletePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateComment(ctx, req.(*MsgCreateComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateComment(ctx, req.(*MsgUpdateComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteComment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteComment(ctx, req.(*MsgDeleteComment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateSubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateSubscription(ctx, req.(*MsgCreateSubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteSubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteSubscription(ctx, req.(*MsgDeleteSubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateLike)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateLike(ctx, req.(*MsgCreateLike))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteLike)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteLike(ctx, req.(*MsgDeleteLike))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAccount(ctx, req.(*MsgCreateAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteAccount(ctx, req.(*MsgDeleteAccount))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sportiverse.sportiverse.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _Msg_CreatePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Msg_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Msg_DeletePost_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _Msg_CreateComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _Msg_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Msg_DeleteComment_Handler,
		},
		{
			MethodName: "CreateSubscription",
			Handler:    _Msg_CreateSubscription_Handler,
		},
		{
			MethodName: "DeleteSubscription",
			Handler:    _Msg_DeleteSubscription_Handler,
		},
		{
			MethodName: "CreateLike",
			Handler:    _Msg_CreateLike_Handler,
		},
		{
			MethodName: "DeleteLike",
			Handler:    _Msg_DeleteLike_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Msg_CreateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Msg_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sportiverse/sportiverse/tx.proto",
}
