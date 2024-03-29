// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: server_streaming/proto/service.proto

package proto

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

// SomethingServiceClient is the client API for SomethingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SomethingServiceClient interface {
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (SomethingService_ExecuteClient, error)
}

type somethingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSomethingServiceClient(cc grpc.ClientConnInterface) SomethingServiceClient {
	return &somethingServiceClient{cc}
}

func (c *somethingServiceClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (SomethingService_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &SomethingService_ServiceDesc.Streams[0], "/SomethingService/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &somethingServiceExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SomethingService_ExecuteClient interface {
	Recv() (*ExecuteResponse, error)
	grpc.ClientStream
}

type somethingServiceExecuteClient struct {
	grpc.ClientStream
}

func (x *somethingServiceExecuteClient) Recv() (*ExecuteResponse, error) {
	m := new(ExecuteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SomethingServiceServer is the server API for SomethingService service.
// All implementations should embed UnimplementedSomethingServiceServer
// for forward compatibility
type SomethingServiceServer interface {
	Execute(*ExecuteRequest, SomethingService_ExecuteServer) error
}

// UnimplementedSomethingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSomethingServiceServer struct {
}

func (UnimplementedSomethingServiceServer) Execute(*ExecuteRequest, SomethingService_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

// UnsafeSomethingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SomethingServiceServer will
// result in compilation errors.
type UnsafeSomethingServiceServer interface {
	mustEmbedUnimplementedSomethingServiceServer()
}

func RegisterSomethingServiceServer(s grpc.ServiceRegistrar, srv SomethingServiceServer) {
	s.RegisterService(&SomethingService_ServiceDesc, srv)
}

func _SomethingService_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecuteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SomethingServiceServer).Execute(m, &somethingServiceExecuteServer{stream})
}

type SomethingService_ExecuteServer interface {
	Send(*ExecuteResponse) error
	grpc.ServerStream
}

type somethingServiceExecuteServer struct {
	grpc.ServerStream
}

func (x *somethingServiceExecuteServer) Send(m *ExecuteResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SomethingService_ServiceDesc is the grpc.ServiceDesc for SomethingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SomethingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SomethingService",
	HandlerType: (*SomethingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _SomethingService_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server_streaming/proto/service.proto",
}
