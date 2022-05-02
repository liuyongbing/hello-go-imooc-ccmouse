package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_error_code/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := proto.NewGreeterClient(conn)
	rsp, err := cli.SayHello(context.Background(), &proto.HelloRequest{
		Name:   "Demo of grpc error codes",
		Url:    "http://keyi.art",
		Gender: proto.GenderEnum_MALE,
		Mp: map[string]string{
			"name":    "proto of map",
			"company": "imooc.com",
		},
		AddTime: timestamppb.New(time.Now()),
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic("Error解析失败")
		}
		fmt.Println("Error message: " + st.Message())
		fmt.Println("Error code: " + st.Code().String())
	} else {
		fmt.Println(rsp.Message)
	}
}
