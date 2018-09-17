package main

import (
	"net"
	_ "time"

	rpcPb "fchat/protos2Go/rpc"

	"github.com/gitfong/fLog"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

const (
	statusSvrAddr = ":9092"
)

var flog *fLog.FLogger

func init() {
	flog = fLog.New()
	if flog == nil {
		panic("fLog.New fail")
	}
}

func main() {
	lis, err := net.Listen("tcp", statusSvrAddr)
	if err != nil {
		flog.Fatal("failed to listen addr:%s, err:%v", statusSvrAddr, err)
	}

	rpcServer := grpc.NewServer()
	rpcPb.RegisterAccountServer(rpcServer, &statusSvr{})
	rpcServer.Serve(lis)
}
