package main

import (
	"io/ioutil"
	"net"
	_ "time"

	rpcPb "fchat/protos2Go/rpc"

	"encoding/json"

	"github.com/gitfong/fLog"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

type statusSvrCfg struct {
	StatusSvrAddr string `json:"statusSvrAddr"`
}

var flog *fLog.FLogger
var svrCfg *statusSvrCfg

func init() {
	flog = fLog.New()
	if flog == nil {
		panic("fLog.New fail")
	}

	data, err := ioutil.ReadFile("statusSvrCfg.json")
	if err != nil {
		flog.Fatal("read svr config err:%v", err)
	}

	svrCfg = new(statusSvrCfg)
	err = json.Unmarshal(data, svrCfg)
	if err != nil {
		flog.Fatal("unmarshal err:%v", err)
	}
}

func main() {
	lis, err := net.Listen("tcp", svrCfg.StatusSvrAddr)
	if err != nil {
		flog.Fatal("failed to listen addr:%s, err:%v", svrCfg.StatusSvrAddr, err)
	}

	rpcServer := grpc.NewServer()
	rpcPb.RegisterStatusServer(rpcServer, &statusSvr{})
	rpcServer.Serve(lis)
}
