// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bridge

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

// BlockTranserClient is the client API for BlockTranser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockTranserClient interface {
	// 由发送方调用函数，接收方实现函数
	TransBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*DhtStatus, error)
	LoadConfig(ctx context.Context, in *DhtStatus, opts ...grpc.CallOption) (*Block, error)
}

type blockTranserClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockTranserClient(cc grpc.ClientConnInterface) BlockTranserClient {
	return &blockTranserClient{cc}
}

func (c *blockTranserClient) TransBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*DhtStatus, error) {
	out := new(DhtStatus)
	err := c.cc.Invoke(ctx, "/bridge.BlockTranser/TransBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockTranserClient) LoadConfig(ctx context.Context, in *DhtStatus, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := c.cc.Invoke(ctx, "/bridge.BlockTranser/LoadConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockTranserServer is the server API for BlockTranser service.
// All implementations must embed UnimplementedBlockTranserServer
// for forward compatibility
type BlockTranserServer interface {
	// 由发送方调用函数，接收方实现函数
	TransBlock(context.Context, *Block) (*DhtStatus, error)
	LoadConfig(context.Context, *DhtStatus) (*Block, error)
	mustEmbedUnimplementedBlockTranserServer()
}

// UnimplementedBlockTranserServer must be embedded to have forward compatible implementations.
type UnimplementedBlockTranserServer struct {
}

func (UnimplementedBlockTranserServer) TransBlock(context.Context, *Block) (*DhtStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransBlock not implemented")
}
func (UnimplementedBlockTranserServer) LoadConfig(context.Context, *DhtStatus) (*Block, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadConfig not implemented")
}
func (UnimplementedBlockTranserServer) mustEmbedUnimplementedBlockTranserServer() {}

// UnsafeBlockTranserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockTranserServer will
// result in compilation errors.
type UnsafeBlockTranserServer interface {
	mustEmbedUnimplementedBlockTranserServer()
}

func RegisterBlockTranserServer(s grpc.ServiceRegistrar, srv BlockTranserServer) {
	s.RegisterService(&BlockTranser_ServiceDesc, srv)
}

func _BlockTranser_TransBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockTranserServer).TransBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bridge.BlockTranser/TransBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockTranserServer).TransBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockTranser_LoadConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DhtStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockTranserServer).LoadConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bridge.BlockTranser/LoadConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockTranserServer).LoadConfig(ctx, req.(*DhtStatus))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockTranser_ServiceDesc is the grpc.ServiceDesc for BlockTranser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockTranser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bridge.BlockTranser",
	HandlerType: (*BlockTranserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransBlock",
			Handler:    _BlockTranser_TransBlock_Handler,
		},
		{
			MethodName: "LoadConfig",
			Handler:    _BlockTranser_LoadConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bridge.proto",
}

// MsgTranserClient is the client API for MsgTranser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgTranserClient interface {
	TransMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*DhtStatus, error)
}

type msgTranserClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgTranserClient(cc grpc.ClientConnInterface) MsgTranserClient {
	return &msgTranserClient{cc}
}

func (c *msgTranserClient) TransMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*DhtStatus, error) {
	out := new(DhtStatus)
	err := c.cc.Invoke(ctx, "/bridge.MsgTranser/TransMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgTranserServer is the server API for MsgTranser service.
// All implementations must embed UnimplementedMsgTranserServer
// for forward compatibility
type MsgTranserServer interface {
	TransMsg(context.Context, *Msg) (*DhtStatus, error)
	mustEmbedUnimplementedMsgTranserServer()
}

// UnimplementedMsgTranserServer must be embedded to have forward compatible implementations.
type UnimplementedMsgTranserServer struct {
}

func (UnimplementedMsgTranserServer) TransMsg(context.Context, *Msg) (*DhtStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransMsg not implemented")
}
func (UnimplementedMsgTranserServer) mustEmbedUnimplementedMsgTranserServer() {}

// UnsafeMsgTranserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgTranserServer will
// result in compilation errors.
type UnsafeMsgTranserServer interface {
	mustEmbedUnimplementedMsgTranserServer()
}

func RegisterMsgTranserServer(s grpc.ServiceRegistrar, srv MsgTranserServer) {
	s.RegisterService(&MsgTranser_ServiceDesc, srv)
}

func _MsgTranser_TransMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgTranserServer).TransMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bridge.MsgTranser/TransMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgTranserServer).TransMsg(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

// MsgTranser_ServiceDesc is the grpc.ServiceDesc for MsgTranser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MsgTranser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bridge.MsgTranser",
	HandlerType: (*MsgTranserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransMsg",
			Handler:    _MsgTranser_TransMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bridge.proto",
}
