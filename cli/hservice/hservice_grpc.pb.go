// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/hservice.proto

package hservice

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

// HookServiceClient is the client API for HookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HookServiceClient interface {
	// Sends a greeting
	Send(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error)
}

type hookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHookServiceClient(cc grpc.ClientConnInterface) HookServiceClient {
	return &hookServiceClient{cc}
}

func (c *hookServiceClient) Send(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageReply, error) {
	out := new(MessageReply)
	err := c.cc.Invoke(ctx, "/hservice.HookService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HookServiceServer is the server API for HookService service.
// All implementations must embed UnimplementedHookServiceServer
// for forward compatibility
type HookServiceServer interface {
	// Sends a greeting
	Send(context.Context, *MessageRequest) (*MessageReply, error)
	mustEmbedUnimplementedHookServiceServer()
}

// UnimplementedHookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHookServiceServer struct {
}

func (UnimplementedHookServiceServer) Send(context.Context, *MessageRequest) (*MessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedHookServiceServer) mustEmbedUnimplementedHookServiceServer() {}

// UnsafeHookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HookServiceServer will
// result in compilation errors.
type UnsafeHookServiceServer interface {
	mustEmbedUnimplementedHookServiceServer()
}

func RegisterHookServiceServer(s grpc.ServiceRegistrar, srv HookServiceServer) {
	s.RegisterService(&HookService_ServiceDesc, srv)
}

func _HookService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HookServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hservice.HookService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HookServiceServer).Send(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HookService_ServiceDesc is the grpc.ServiceDesc for HookService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HookService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hservice.HookService",
	HandlerType: (*HookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _HookService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hservice.proto",
}
