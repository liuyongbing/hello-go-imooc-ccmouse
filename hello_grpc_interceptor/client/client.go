package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	protoInterceptor "github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_interceptor/proto"
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

	cli := protoInterceptor.NewGreeterClient(conn)
	rsp, err := cli.SayHello(context.Background(), &protoInterceptor.HelloRequest{
		Name: "Demo of grpc interceptor from client",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Message)
}
