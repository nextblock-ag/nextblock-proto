// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: api.proto

package api

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
	Api_PostSubmitV2_FullMethodName      = "/api.Api/PostSubmitV2"
	Api_PostSubmitBatchV2_FullMethodName = "/api.Api/PostSubmitBatchV2"
)

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	PostSubmitV2(ctx context.Context, in *PostSubmitRequest, opts ...grpc.CallOption) (*PostSubmitResponse, error)
	PostSubmitBatchV2(ctx context.Context, in *PostSubmitBatchRequest, opts ...grpc.CallOption) (*PostSubmitBatchResponse, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) PostSubmitV2(ctx context.Context, in *PostSubmitRequest, opts ...grpc.CallOption) (*PostSubmitResponse, error) {
	out := new(PostSubmitResponse)
	err := c.cc.Invoke(ctx, Api_PostSubmitV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PostSubmitBatchV2(ctx context.Context, in *PostSubmitBatchRequest, opts ...grpc.CallOption) (*PostSubmitBatchResponse, error) {
	out := new(PostSubmitBatchResponse)
	err := c.cc.Invoke(ctx, Api_PostSubmitBatchV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	PostSubmitV2(context.Context, *PostSubmitRequest) (*PostSubmitResponse, error)
	PostSubmitBatchV2(context.Context, *PostSubmitBatchRequest) (*PostSubmitBatchResponse, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) PostSubmitV2(context.Context, *PostSubmitRequest) (*PostSubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSubmitV2 not implemented")
}
func (UnimplementedApiServer) PostSubmitBatchV2(context.Context, *PostSubmitBatchRequest) (*PostSubmitBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSubmitBatchV2 not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_PostSubmitV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostSubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PostSubmitV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_PostSubmitV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PostSubmitV2(ctx, req.(*PostSubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PostSubmitBatchV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostSubmitBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PostSubmitBatchV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_PostSubmitBatchV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PostSubmitBatchV2(ctx, req.(*PostSubmitBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostSubmitV2",
			Handler:    _Api_PostSubmitV2_Handler,
		},
		{
			MethodName: "PostSubmitBatchV2",
			Handler:    _Api_PostSubmitBatchV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
