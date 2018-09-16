package main

import (
	"context"
	rpcPb "fchat/protos2Go/rpc"
	"time"

	"google.golang.org/grpc"
)

func init() {
	rpcClient.init()
	go rpcClient.heartbeat()
}

var (
	rpcClient = new(rpcClientProxy)
)

type rpcClientProxy struct {
	accountClient rpcPb.AccountClient
	serverAddr    string
}

func (rc *rpcClientProxy) init() {
	rc.serverAddr = "localhost:9091"

	conn, err := grpc.Dial(rc.serverAddr, grpc.WithInsecure())
	if err != nil {
		flog.Fatal("[gatewaySvr] grpc Dial %s err: %v", rc.serverAddr, err)
	}
	//defer conn.Close()
	rc.accountClient = rpcPb.NewAccountClient(conn)
}

//grpc心跳检测，重连
func (rc *rpcClientProxy) heartbeat() {
	for {
		time.Sleep(3 * time.Second)
		_, err := rc.getCli().Heartbeat(context.Background(), &rpcPb.HeartbeatReq{})
		if err != nil {
			rc.reDial()
		}
	}
}

func (rc *rpcClientProxy) getCli() rpcPb.AccountClient {
	return rc.accountClient
}

func (rc *rpcClientProxy) reDial() {
	flog.Debug("[gatewaySvr] grpc reDial %s", rc.serverAddr)
	conn, err := grpc.Dial(rc.serverAddr, grpc.WithInsecure())
	if err != nil {
		flog.Error("[gatewaySvr] grpc reDial %s err: %v", rc.serverAddr, err)
		return
	}
	//defer conn.Close()
	rc.accountClient = rpcPb.NewAccountClient(conn)
}
