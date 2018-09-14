package main

import (
	"context"
	rpcPb "fchat/protos2Go/rpc"
	"fmt"
	"log"
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
		log.Fatalf("[gatewaySvr] grpc Dial %s err: %v", rc.serverAddr, err)
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
		} else {
			fmt.Printf("[gatewaySvr] heartbeat normal. serverAddr:%s\n", rc.serverAddr)
		}
	}
}

func (rc *rpcClientProxy) getCli() rpcPb.AccountClient {
	return rc.accountClient
}

func (rc *rpcClientProxy) reDial() {
	fmt.Printf("[gatewaySvr] grpc reDial %s\n", rc.serverAddr)
	conn, err := grpc.Dial(rc.serverAddr, grpc.WithInsecure())
	if err != nil {
		//log.Fatalf("[gatewaySvr] grpc reDial %s err: %v", rc.serverAddr, err)
		fmt.Printf("[gatewaySvr] grpc reDial %s err: %v\n", rc.serverAddr, err)
		return
	}
	//defer conn.Close()
	rc.accountClient = rpcPb.NewAccountClient(conn)
}
