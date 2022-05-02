package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protoValidate "github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_validate/proto"
)

type Server struct {
	protoValidate.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *protoValidate.Person) (*protoValidate.Person, error) {
	//return &protoInterceptor.HelloReply{
	//	Message: "Hello, demo of grpc interceptor",
	//}, nil
	return &protoValidate.Person{
		Id: request.Id,
	}, nil
}

type Validator interface {
	Validate() error
}

func main() {

	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新的请求")

		if r, ok := req.(Validator); ok {
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		resp, err = handler(ctx, req)
		fmt.Println("请求结束")
		return resp, err
	}

	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	protoValidate.RegisterGreeterServer(g, &Server{})

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("Failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("Failed to start grpc:" + err.Error())
	}
}
