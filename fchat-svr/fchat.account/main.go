package main

import (
	"fmt"
	"log"
	"net"
	_ "time"

	rpcPb "fchat/protos2Go/rpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func init() {
}

const (
	accountSvrddr = ":9091"
)

type accountServer struct{}

func (s *accountServer) Register(ctx context.Context, in *rpcPb.RegisterReq) (*rpcPb.RegisterRsp, error) {
	//time.Sleep(time.Second * 10)	//客户端会阻塞等待
	fmt.Println("[accountSvr] on Register ", in.Account, in.Password)
	return &rpcPb.RegisterRsp{RetCode: 0, Uid: 1, Desc: "注册成功"}, nil
}

func (s *accountServer) Login(ctx context.Context, in *rpcPb.LoginReq) (*rpcPb.LoginRsp, error) {
	fmt.Println("[accountSvr] on Login ", in.Account, in.Password)
	return &rpcPb.LoginRsp{RetCode: 0, Uid: 1, Desc: "注册成功"}, nil
}

func main() {
	lis, err := net.Listen("tcp", accountSvrddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	rpcPb.RegisterAccountServer(rpcServer, &accountServer{})
	rpcServer.Serve(lis)

}
