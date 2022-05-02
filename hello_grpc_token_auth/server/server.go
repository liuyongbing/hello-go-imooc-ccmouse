package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"

	protoTokenAuth "github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_token_auth/proto"
)

type Server struct {
	protoTokenAuth.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *protoTokenAuth.HelloRequest) (*protoTokenAuth.HelloReply, error) {
	return &protoTokenAuth.HelloReply{
		Message: "Hello, demo of grpc token authed",
	}, nil
}

func main() {

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新的请求")

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "Token authed failed.")
		}
		var (
			appid  string
			appkey string
		)
		if appSlice, ok := md["appid"]; ok {
			appid = appSlice[0]
		}
		if appSlice, ok := md["appkey"]; ok {
			appkey = appSlice[0]
		}
		if appid != "10001" || appkey != "com.imooc" {
			return resp, status.Error(codes.Unauthenticated, "Appid or appkey error")
		}

		resp, err = handler(ctx, req)
		fmt.Println("请求结束")

		return resp, err
	}

	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	protoTokenAuth.RegisterGreeterServer(g, &Server{})

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("Failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("Failed to start grpc:" + err.Error())
	}
}
