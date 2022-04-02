package main

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/liuyongbing/hello-go-imooc-ccmouse/hello_grpc_stream/proto"
)

const PORT = ":50052"

type Server struct {
	proto.UnimplementedGreeterServer
}

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

func (s *Server) AllStream(allStr proto.Greeter_AllStreamServer) error {

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
