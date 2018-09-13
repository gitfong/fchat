package main

import (
	csMsg "fchat/protos2Go"
	"fmt"
	_ "time"
)

//
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

	rspMsg := &csMsg.Msg{
		ID: csMsg.MsgID_loginRsp,
		LoginRsp: &csMsg.MsgLoginRsp{
			RetCode: 0,
			Desc:    "登录成功",
		},
	}

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

//以下函数全是在gatewaySvr运行的，所以尽量不要有太复杂的逻辑，把复杂的逻辑都传给其他服去处理
func handleRegisterReq(addr string, msg *csMsg.Msg, cm *connsManager) {
	fmt.Printf("handleRegisterReq addr=%v, msg=%v\n", addr, msg)

	rspMsg := &csMsg.Msg{
		ID: csMsg.MsgID_registerRsp,
		RegisterRsp: &csMsg.MsgRegisterRsp{
			RetCode: 0,
			Desc:    "register succeed!",
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
