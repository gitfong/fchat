package main

import (
	_ "fmt"
	"log"
	"net"
	_"time"

	pb "fchat/pbgos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func init() {
}

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//time.Sleep(time.Second * 10)	//客户端会阻塞等待
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)

}
