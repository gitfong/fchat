package main

import (
	"context"
	rpcPb "fchat/protos2Go/rpc"
	"time"

	"google.golang.org/grpc"
)

var (
	rpcClient = new(rpcClientProxy)
)

func init() {
	rpcClient.init()
	go rpcClient.heartbeat()
}

type rpcClientProxy struct {
	accountSvrAddr string
	accountClient  rpcPb.AccountClient
	statusSvrAddr  string
	statusClient   rpcPb.StatusClient
}

func (rc *rpcClientProxy) init() {
	rc.accountSvrAddr = "localhost:9091"
	rc.statusSvrAddr = "localhost:9092"

	//connect to account server
	conn, err := grpc.Dial(rc.accountSvrAddr, grpc.WithInsecure())
	if err != nil {
		flog.Fatal("[gatewaySvr] grpc Dial %s err: %v", rc.accountSvrAddr, err)
	}
	//defer conn.Close()
	rc.accountClient = rpcPb.NewAccountClient(conn)

	//connect to status server
	statusSvrConn, err := grpc.Dial(rc.statusSvrAddr, grpc.WithInsecure())
	if err != nil {
		flog.Fatal("[gatewaySvr] grpc Dial %s err: %v", rc.statusSvrAddr, err)
	}
	rc.statusClient = rpcPb.NewStatusClient(statusSvrConn)
}

//grpc心跳检测，重连
func (rc *rpcClientProxy) heartbeat() {
	for {
		time.Sleep(3 * time.Second)
		_, err := rc.getAccountCli().Heartbeat(context.Background(), &rpcPb.HeartbeatReq{})
		if err != nil {
			rc.reDialAccountSvr()
		}

		_, err = rc.getStatusCli().Heartbeat(context.Background(), &rpcPb.Heartbeat2StatusReq{})
		if err != nil {
			rc.reDialStatusSvr()
		}
	}
}

func (rc *rpcClientProxy) getAccountCli() rpcPb.AccountClient {
	return rc.accountClient
}

func (rc *rpcClientProxy) reDialAccountSvr() {
	flog.Debug("[gatewaySvr] grpc reDial %s", rc.accountSvrAddr)
	conn, err := grpc.Dial(rc.accountSvrAddr, grpc.WithInsecure())
	if err != nil {
		flog.Error("[gatewaySvr] grpc reDial %s err: %v", rc.accountSvrAddr, err)
		return
	}
	//defer conn.Close()
	rc.accountClient = rpcPb.NewAccountClient(conn)
}

func (rc *rpcClientProxy) getStatusCli() rpcPb.StatusClient {
	return rc.statusClient
}

func (rc *rpcClientProxy) reDialStatusSvr() {
	flog.Debug("[gatewaySvr] grpc reDial %s", rc.statusSvrAddr)
	conn, err := grpc.Dial(rc.statusSvrAddr, grpc.WithInsecure())
	if err != nil {
		flog.Error("[gatewaySvr] grpc reDial %s err: %v", rc.statusSvrAddr, err)
		return
	}
	//defer conn.Close()
	rc.statusClient = rpcPb.NewStatusClient(conn)
}
