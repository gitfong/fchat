package main

import (
	"net"
	_ "time"

	rpcPb "fchat/protos2Go/rpc"

	"github.com/gitfong/fLog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var flog *fLog.FLogger

func init() {
	flog = fLog.New()
	if flog == nil {
		panic("fLog.New fail")
	}
}

const (
	accountSvrAddr = ":9091"
)

type accountServer struct{}

func (s *accountServer) Heartbeat(ctx context.Context, in *rpcPb.HeartbeatReq) (*rpcPb.HeartbeatRsp, error) {
	return &rpcPb.HeartbeatRsp{}, nil
}

func (s *accountServer) Register(ctx context.Context, in *rpcPb.RegisterReq) (*rpcPb.RegisterRsp, error) {
	//time.Sleep(time.Second * 10)	//客户端会阻塞等待
	flog.Debug("[accountSvr] on Register %s,%s", in.Account, in.Password)
	return &rpcPb.RegisterRsp{RetCode: 0, Uid: 1, Desc: "register succeed!"}, nil
}

func (s *accountServer) Login(ctx context.Context, in *rpcPb.LoginReq) (*rpcPb.LoginRsp, error) {
	flog.Debug("[accountSvr] on Login %s,%s", in.Account, in.Password)
	return &rpcPb.LoginRsp{RetCode: 0, Uid: 1, Desc: "login succeed!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", accountSvrAddr)
	if err != nil {
		flog.Fatal("failed to listen addr:%s, err:%v", accountSvrAddr, err)
	}

	rpcServer := grpc.NewServer()
	rpcPb.RegisterAccountServer(rpcServer, &accountServer{})
	rpcServer.Serve(lis)
}
