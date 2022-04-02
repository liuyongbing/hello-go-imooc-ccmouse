package main

import (
	"context"
	"fmt"
	"sync"
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

	// 流模式: 客户端(发送) > 服务端(接收)
	//CliPutStream(conn)

	// 流模式: 双向流
	CliAllStream(conn)

}

func CliAllStream(conn *grpc.ClientConn) {
	cli := proto.NewGreeterClient(conn)
	allStr, _ := cli.AllStream(context.Background())

	// 协程: 开启
	group := sync.WaitGroup{}
	group.Add(2)

	// 接收
	go func() {
		defer group.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("双向流(客户端接收): ", data.Data)
		}
	}()

	// 发送
	go func() {
		defer group.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{
				Data: fmt.Sprintf("双向流(客户端发送): %v", time.Now().Unix()),
			})
			time.Sleep(time.Second)
		}
	}()

	group.Wait()
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
