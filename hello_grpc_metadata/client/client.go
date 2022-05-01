package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"

	protoMetadata "github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_metadata/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	md := metadata.New(map[string]string{
		"name":     "bobby",
		"password": "imooc",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	cli := protoMetadata.NewGreeterClient(conn)
	rsp, err := cli.SayHello(ctx, &protoMetadata.HelloRequest{
		Name: "Grpc metadata",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Message)
}
