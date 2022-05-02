package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	protoInterceptor "github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_interceptor/proto"
)

type Server struct {
	protoInterceptor.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *protoInterceptor.HelloRequest) (*protoInterceptor.HelloReply, error) {
	return &protoInterceptor.HelloReply{
		Message: "Hello, demo of grpc interceptor",
	}, nil
}

func main() {

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新的请求")
		resp, err = handler(ctx, req)
		fmt.Println("请求结束")
		return resp, err
	}

	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	protoInterceptor.RegisterGreeterServer(g, &Server{})

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("Failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("Failed to start grpc:" + err.Error())
	}
}
