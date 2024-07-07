// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.11.4
// source: api/kvstore/v1/kv.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Kv_Set_FullMethodName = "/api.kvstore.v1.Kv/Set"
	Kv_Get_FullMethodName = "/api.kvstore.v1.Kv/Get"
	Kv_Del_FullMethodName = "/api.kvstore.v1.Kv/Del"
)

// KvClient is the client API for Kv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KvClient interface {
	Set(ctx context.Context, in *KvRequest, opts ...grpc.CallOption) (*KvReply, error)
	Get(ctx context.Context, in *KvRequest, opts ...grpc.CallOption) (*KvReply, error)
	Del(ctx context.Context, in *KvRequest, opts ...grpc.CallOption) (*KvReply, error)
}

type kvClient struct {
	cc grpc.ClientConnInterface
}

func NewKvClient(cc grpc.ClientConnInterface) KvClient {
	return &kvClient{cc}
}

func (c *kvClient) Set(ctx context.Context, in *KvRequest, opts ...grpc.CallOption) (*KvReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KvReply)
	err := c.cc.Invoke(ctx, Kv_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvClient) Get(ctx context.Context, in *KvRequest, opts ...grpc.CallOption) (*KvReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KvReply)
	err := c.cc.Invoke(ctx, Kv_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvClient) Del(ctx context.Context, in *KvRequest, opts ...grpc.CallOption) (*KvReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KvReply)
	err := c.cc.Invoke(ctx, Kv_Del_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KvServer is the server API for Kv service.
// All implementations must embed UnimplementedKvServer
// for forward compatibility
type KvServer interface {
	Set(context.Context, *KvRequest) (*KvReply, error)
	Get(context.Context, *KvRequest) (*KvReply, error)
	Del(context.Context, *KvRequest) (*KvReply, error)
	mustEmbedUnimplementedKvServer()
}

// UnimplementedKvServer must be embedded to have forward compatible implementations.
type UnimplementedKvServer struct {
}

func (UnimplementedKvServer) Set(context.Context, *KvRequest) (*KvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedKvServer) Get(context.Context, *KvRequest) (*KvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKvServer) Del(context.Context, *KvRequest) (*KvReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedKvServer) mustEmbedUnimplementedKvServer() {}

// UnsafeKvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KvServer will
// result in compilation errors.
type UnsafeKvServer interface {
	mustEmbedUnimplementedKvServer()
}

func RegisterKvServer(s grpc.ServiceRegistrar, srv KvServer) {
	s.RegisterService(&Kv_ServiceDesc, srv)
}

func _Kv_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kv_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvServer).Set(ctx, req.(*KvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kv_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kv_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvServer).Get(ctx, req.(*KvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kv_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kv_Del_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvServer).Del(ctx, req.(*KvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Kv_ServiceDesc is the grpc.ServiceDesc for Kv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.kvstore.v1.Kv",
	HandlerType: (*KvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _Kv_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Kv_Get_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Kv_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/kvstore/v1/kv.proto",
}
