// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto/sample-manager.proto

package sample_manager

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

// SampleServiceClient is the client API for SampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleServiceClient interface {
	GetSampleItemID(ctx context.Context, in *GetSampleItemIDRequest, opts ...grpc.CallOption) (*GetSampleItemIDResponse, error)
	CreateSampleItem(ctx context.Context, in *CreateSampleItemRequest, opts ...grpc.CallOption) (*CreateSampleItemResponse, error)
}

type sampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleServiceClient(cc grpc.ClientConnInterface) SampleServiceClient {
	return &sampleServiceClient{cc}
}

func (c *sampleServiceClient) GetSampleItemID(ctx context.Context, in *GetSampleItemIDRequest, opts ...grpc.CallOption) (*GetSampleItemIDResponse, error) {
	out := new(GetSampleItemIDResponse)
	err := c.cc.Invoke(ctx, "/SampleService/GetSampleItemID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleServiceClient) CreateSampleItem(ctx context.Context, in *CreateSampleItemRequest, opts ...grpc.CallOption) (*CreateSampleItemResponse, error) {
	out := new(CreateSampleItemResponse)
	err := c.cc.Invoke(ctx, "/SampleService/CreateSampleItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServiceServer is the server API for SampleService service.
// All implementations must embed UnimplementedSampleServiceServer
// for forward compatibility
type SampleServiceServer interface {
	GetSampleItemID(context.Context, *GetSampleItemIDRequest) (*GetSampleItemIDResponse, error)
	CreateSampleItem(context.Context, *CreateSampleItemRequest) (*CreateSampleItemResponse, error)
	mustEmbedUnimplementedSampleServiceServer()
}

// UnimplementedSampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSampleServiceServer struct {
}

func (UnimplementedSampleServiceServer) GetSampleItemID(context.Context, *GetSampleItemIDRequest) (*GetSampleItemIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSampleItemID not implemented")
}
func (UnimplementedSampleServiceServer) CreateSampleItem(context.Context, *CreateSampleItemRequest) (*CreateSampleItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSampleItem not implemented")
}
func (UnimplementedSampleServiceServer) mustEmbedUnimplementedSampleServiceServer() {}

// UnsafeSampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleServiceServer will
// result in compilation errors.
type UnsafeSampleServiceServer interface {
	mustEmbedUnimplementedSampleServiceServer()
}

func RegisterSampleServiceServer(s grpc.ServiceRegistrar, srv SampleServiceServer) {
	s.RegisterService(&SampleService_ServiceDesc, srv)
}

func _SampleService_GetSampleItemID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSampleItemIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServiceServer).GetSampleItemID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SampleService/GetSampleItemID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServiceServer).GetSampleItemID(ctx, req.(*GetSampleItemIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleService_CreateSampleItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSampleItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServiceServer).CreateSampleItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SampleService/CreateSampleItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServiceServer).CreateSampleItem(ctx, req.(*CreateSampleItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SampleService_ServiceDesc is the grpc.ServiceDesc for SampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SampleService",
	HandlerType: (*SampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSampleItemID",
			Handler:    _SampleService_GetSampleItemID_Handler,
		},
		{
			MethodName: "CreateSampleItem",
			Handler:    _SampleService_CreateSampleItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sample-manager.proto",
}
