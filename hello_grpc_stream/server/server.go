package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_stream/proto"
)

const PORT = ":50052"

type Server struct {
	proto.UnimplementedGreeterServer
}

// 流模式: 服务端(发送) > 客户端(接收)
func (s *Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++

		res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)

		if i > 5 {
			break
		}
	}

	return nil
}

// 流模式: 客户端(发送) > 服务端(接收)
func (s *Server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		req, err := cliStr.Recv()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(req.Data)
	}

	return nil
}

// 流模式: 双向流
func (s *Server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	// 协程: 开启
	group := sync.WaitGroup{}
	group.Add(2)

	// 接收
	go func() {
		defer group.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("双向流(服务端接收): ", data.Data)
		}
	}()

	// 发送
	go func() {
		defer group.Done()
		for {
			_ = allStr.Send(&proto.StreamResData{
				Data: fmt.Sprintf("双向流(服务端发送): %v", time.Now().Unix()),
			})
			time.Sleep(time.Second)
		}
	}()

	group.Wait()

	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
