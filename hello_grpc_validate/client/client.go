package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	protoValidate "github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_validate/proto"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("耗时:%s\n", time.Since(start))

		return err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	conn, err := grpc.Dial("127.0.0.1:8080", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := protoValidate.NewGreeterClient(conn)
	rsp, err := cli.SayHello(context.Background(), &protoValidate.Person{
		Id:     1000,
		Email:  "mike.doon.lau@gmail.com",
		Mobile: "18516553344",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id)
}
