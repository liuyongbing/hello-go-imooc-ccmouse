package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := proto.NewGreeterClient(conn)
	rsp, err := cli.SayHello(context.Background(), &proto.HelloRequest{Name: "Grpc service"})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Message)
}
