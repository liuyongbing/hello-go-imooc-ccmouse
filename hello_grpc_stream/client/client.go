package main

import (
	"context"
	"fmt"
	"time"

	"github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_stream/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 流模式: 服务端(发送) > 客户端(接收)
	//CliGetStream(conn)

	// 流模式: 服务端(发送) > 客户端(接收)
	CliPutStream(conn)

}

// 流模式: 客户端(发送) > 服务端(接收)
func CliPutStream(conn *grpc.ClientConn) {
	cli := proto.NewGreeterClient(conn)
	putS, _ := cli.PutStream(context.Background())

	i := 0
	for {
		i++

		_ = putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("Client put stream:%d", i),
		})
		time.Sleep(time.Second)

		if i >= 5 {
			break
		}
	}
}

// 流模式: 服务端(发送) > 客户端(接收)
func CliGetStream(conn *grpc.ClientConn) {
	cli := proto.NewGreeterClient(conn)
	rsp, _ := cli.GetStream(context.Background(), &proto.StreamReqData{Data: "stream of client"})
	i := 0
	for {
		i++

		data, err := rsp.Recv() // 底层为 socket send()/recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(i, ":", data)
	}
}
