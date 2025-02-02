// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: grpc/process/process.proto

package process

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
	ProcessService_GetProcessInfo_FullMethodName = "/grpc_process.ProcessService/GetProcessInfo"
	ProcessService_GetProcessList_FullMethodName = "/grpc_process.ProcessService/GetProcessList"
)

// ProcessServiceClient is the client API for ProcessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessServiceClient interface {
	GetProcessInfo(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessReply, error)
	GetProcessList(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (ProcessService_GetProcessListClient, error)
}

type processServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessServiceClient(cc grpc.ClientConnInterface) ProcessServiceClient {
	return &processServiceClient{cc}
}

func (c *processServiceClient) GetProcessInfo(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessReply, error) {
	out := new(ProcessReply)
	err := c.cc.Invoke(ctx, ProcessService_GetProcessInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processServiceClient) GetProcessList(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (ProcessService_GetProcessListClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProcessService_ServiceDesc.Streams[0], ProcessService_GetProcessList_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &processServiceGetProcessListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessService_GetProcessListClient interface {
	Recv() (*ProcessReply, error)
	grpc.ClientStream
}

type processServiceGetProcessListClient struct {
	grpc.ClientStream
}

func (x *processServiceGetProcessListClient) Recv() (*ProcessReply, error) {
	m := new(ProcessReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessServiceServer is the server API for ProcessService service.
// All implementations must embed UnimplementedProcessServiceServer
// for forward compatibility
type ProcessServiceServer interface {
	GetProcessInfo(context.Context, *ProcessRequest) (*ProcessReply, error)
	GetProcessList(*EmptyRequest, ProcessService_GetProcessListServer) error
	mustEmbedUnimplementedProcessServiceServer()
}

// UnimplementedProcessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProcessServiceServer struct {
}

func (UnimplementedProcessServiceServer) GetProcessInfo(context.Context, *ProcessRequest) (*ProcessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProcessInfo not implemented")
}
func (UnimplementedProcessServiceServer) GetProcessList(*EmptyRequest, ProcessService_GetProcessListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProcessList not implemented")
}
func (UnimplementedProcessServiceServer) mustEmbedUnimplementedProcessServiceServer() {}

// UnsafeProcessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessServiceServer will
// result in compilation errors.
type UnsafeProcessServiceServer interface {
	mustEmbedUnimplementedProcessServiceServer()
}

func RegisterProcessServiceServer(s grpc.ServiceRegistrar, srv ProcessServiceServer) {
	s.RegisterService(&ProcessService_ServiceDesc, srv)
}

func _ProcessService_GetProcessInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServiceServer).GetProcessInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProcessService_GetProcessInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServiceServer).GetProcessInfo(ctx, req.(*ProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessService_GetProcessList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessServiceServer).GetProcessList(m, &processServiceGetProcessListServer{stream})
}

type ProcessService_GetProcessListServer interface {
	Send(*ProcessReply) error
	grpc.ServerStream
}

type processServiceGetProcessListServer struct {
	grpc.ServerStream
}

func (x *processServiceGetProcessListServer) Send(m *ProcessReply) error {
	return x.ServerStream.SendMsg(m)
}

// ProcessService_ServiceDesc is the grpc.ServiceDesc for ProcessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_process.ProcessService",
	HandlerType: (*ProcessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProcessInfo",
			Handler:    _ProcessService_GetProcessInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetProcessList",
			Handler:       _ProcessService_GetProcessList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/process/process.proto",
}
