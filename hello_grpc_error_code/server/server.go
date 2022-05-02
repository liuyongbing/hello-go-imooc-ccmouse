package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_error_code/proto"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	//return &proto.HelloReply{
	//	Message: "Hello, demo of grpc. Name:" + request.Name + " Url:" + request.Url + " Gender:" + request.Gender.String(),
	//}, nil
	return nil, status.Errorf(codes.NotFound, "记录未找到: %s", request.Name)
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("Failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("Failed to start grpc:" + err.Error())
	}
}
