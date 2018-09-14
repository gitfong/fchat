package main

import (
	"context"
	csMsg "fchat/protos2Go"
	rpcPb "fchat/protos2Go/rpc"
	"fmt"
	"log"
	_ "time"

	//"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	accountClient rpcPb.AccountClient
)

func init() {
	accountSvrAddr := "localhost:9091"
	conn, err := grpc.Dial(accountSvrAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not ocnnect:%v", err)
	}

	//defer conn.Close()
	accountClient = rpcPb.NewAccountClient(conn)
}

//以下函数全是在gatewaySvr运行的，所以尽量不要有太复杂的逻辑，把复杂的逻辑都传给其他服去处理
var MsgHandleFunc = map[csMsg.MsgID]func(string, *csMsg.Msg, *connsManager){
	csMsg.MsgID_registerReq: handleRegisterReq,
	csMsg.MsgID_registerRsp: handleRegisterRsp,
	csMsg.MsgID_loginReq:    handleLoginReq,
	csMsg.MsgID_loginRsp:    handleLoginRsp,
	csMsg.MsgID_signInReq:   handleSignInReq,
}

func handleSignInReq(addr string, msg *csMsg.Msg, cm *connsManager) {
	fmt.Printf("handleSignInReq msg=%v\n", msg)
}

func handleLoginRsp(addr string, msg *csMsg.Msg, cm *connsManager) {
	fmt.Printf("handleLoginRsp msg=%v\n", *msg)
}

func handleLoginReq(addr string, msg *csMsg.Msg, cm *connsManager) {
	fmt.Printf("handleLoginReq msg=%v\n", *msg)

	rsp, err := accountClient.Login(context.Background(), &rpcPb.LoginReq{Account: msg.GetLoginReq().Account, Password: msg.GetLoginReq().Password})
	if err != nil {
		fmt.Printf("[gatewaySvr] Login account:%v, password:%v err:%v\n", msg.GetLoginReq().Account, msg.GetLoginReq().Password, err)
	}

	rspMsg := &csMsg.Msg{
		ID: csMsg.MsgID_loginRsp,
		LoginRsp: &csMsg.MsgLoginRsp{
			RetCode: rsp.RetCode,
			Desc:    rsp.Desc,
		},
	}
	//rspMsg.GetRegisterRsp().RetCode = 0
	//rspMsg.GetRegisterRsp().Desc = "register succeed!"

	rData := &rspData{
		targetAddr: addr,
		fromAddr:   addr,
		data:       rspMsg,
	}
	cm.chRsp <- rData
}

func handleRegisterRsp(addr string, msg *csMsg.Msg, cm *connsManager) {
	fmt.Printf("handleRegisterRsp msg=%v\n", *msg)
}

func handleRegisterReq(addr string, msg *csMsg.Msg, cm *connsManager) {
	fmt.Printf("handleRegisterReq addr=%v, msg=%v\n", addr, msg)

	rsp, err := accountClient.Register(context.Background(), &rpcPb.RegisterReq{Account: msg.GetRegisterReq().Account, Password: msg.GetRegisterReq().Password})
	if err != nil {
		fmt.Printf("[gatewaySvr] Register account:%v, password:%v err:%v\n", msg.GetRegisterReq().Account, msg.GetRegisterReq().Password, err)
	}

	rspMsg := &csMsg.Msg{
		ID: csMsg.MsgID_registerRsp,
		RegisterRsp: &csMsg.MsgRegisterRsp{
			RetCode: rsp.RetCode,
			Desc:    rsp.Desc,
		},
	}
	//rspMsg.GetRegisterRsp().RetCode = 0
	//rspMsg.GetRegisterRsp().Desc = "register succeed!"

	rData := &rspData{
		targetAddr: addr,
		fromAddr:   addr,
		data:       rspMsg,
	}
	cm.chRsp <- rData
}

type rspData struct {
	targetAddr string
	fromAddr   string
	data       *csMsg.Msg
}
