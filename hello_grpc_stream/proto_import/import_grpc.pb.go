// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: import.proto

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProtoImportClient is the client API for ProtoImport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProtoImportClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	//  使用 google.protobuf.Empty 代替
	//  rpc Ping(Empty) returns (Pong);
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Pong, error)
}

type protoImportClient struct {
	cc grpc.ClientConnInterface
}

func NewProtoImportClient(cc grpc.ClientConnInterface) ProtoImportClient {
	return &protoImportClient{cc}
}

func (c *protoImportClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/ProtoImport/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protoImportClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/ProtoImport/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtoImportServer is the server API for ProtoImport service.
// All implementations must embed UnimplementedProtoImportServer
// for forward compatibility
type ProtoImportServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	//  使用 google.protobuf.Empty 代替
	//  rpc Ping(Empty) returns (Pong);
	Ping(context.Context, *empty.Empty) (*Pong, error)
	mustEmbedUnimplementedProtoImportServer()
}

// UnimplementedProtoImportServer must be embedded to have forward compatible implementations.
type UnimplementedProtoImportServer struct {
}

func (UnimplementedProtoImportServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedProtoImportServer) Ping(context.Context, *empty.Empty) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedProtoImportServer) mustEmbedUnimplementedProtoImportServer() {}

// UnsafeProtoImportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProtoImportServer will
// result in compilation errors.
type UnsafeProtoImportServer interface {
	mustEmbedUnimplementedProtoImportServer()
}

func RegisterProtoImportServer(s grpc.ServiceRegistrar, srv ProtoImportServer) {
	s.RegisterService(&ProtoImport_ServiceDesc, srv)
}

func _ProtoImport_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoImportServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProtoImport/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoImportServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtoImport_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoImportServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProtoImport/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoImportServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ProtoImport_ServiceDesc is the grpc.ServiceDesc for ProtoImport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProtoImport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProtoImport",
	HandlerType: (*ProtoImportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ProtoImport_SayHello_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ProtoImport_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "import.proto",
}