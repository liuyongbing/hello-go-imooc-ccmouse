package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"

	protoMetadata "github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_metadata/proto"
)

type Server struct {
	protoMetadata.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *protoMetadata.HelloRequest) (*protoMetadata.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get grpc metadata error")
	}
	//for key, val := range md {
	//	fmt.Println(key, val)
	//}
	nameSlice, ok := md["name"]
	if !ok {
		fmt.Println("No md[name]")
	}
	fmt.Println("Server incoming metadata:")
	fmt.Println(nameSlice)
	for i, e := range nameSlice {
		fmt.Println(i, e)
	}

	return &protoMetadata.HelloReply{
		Message: "Hello, demo of grpc metadata",
	}, nil
}

func main() {
	g := grpc.NewServer()
	protoMetadata.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("Failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("Failed to start grpc:" + err.Error())
	}
}
