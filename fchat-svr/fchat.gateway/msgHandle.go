package main

import (
	"context"
	csMsg "fchat/protos2Go"
	rpcPb "fchat/protos2Go/rpc"
	_ "time"
)

//以下函数全是在gatewaySvr运行的，所以尽量不要有太复杂的逻辑，把复杂的逻辑都传给其他服去处理
var MsgHandleFunc = map[csMsg.MsgID]func(string, *csMsg.Msg, *connsManager){
	csMsg.MsgID_registerReq: handleRegisterReq,
	csMsg.MsgID_registerRsp: handleRegisterRsp,
	csMsg.MsgID_loginReq:    handleLoginReq,
	csMsg.MsgID_loginRsp:    handleLoginRsp,
	csMsg.MsgID_signInReq:   handleSignInReq,
}

func handleSignInReq(addr string, msg *csMsg.Msg, cm *connsManager) {
	flog.Debug("handleSignInReq msg=%v", msg)
}

func handleLoginRsp(addr string, msg *csMsg.Msg, cm *connsManager) {
	flog.Debug("handleLoginRsp msg=%v", *msg)
}

func handleLoginReq(addr string, msg *csMsg.Msg, cm *connsManager) {
	flog.Debug("handleLoginReq msg=%v", *msg)

	rsp, err := rpcClient.getCli().Login(context.Background(), &rpcPb.LoginReq{Account: msg.GetLoginReq().Account, Password: msg.GetLoginReq().Password})
	if err != nil {
		flog.Error("[gatewaySvr] Login account:%v, password:%v err:%v", msg.GetLoginReq().Account, msg.GetLoginReq().Password, err)

		rData := new(rspData)
		rData.setData(addr, addr,
			&csMsg.Msg{
				ID: csMsg.MsgID_loginRsp,
				LoginRsp: &csMsg.MsgLoginRsp{
					RetCode: 1,
					Desc:    "Network exception, please try again later.",
				},
			},
		)
		cm.chRsp <- rData
		return
	}

	rspMsg := new(csMsg.Msg)
	rspMsg.ID = csMsg.MsgID_loginRsp
	rspMsg.LoginRsp = new(csMsg.MsgLoginRsp)
	rspMsg.GetLoginRsp().RetCode = rsp.RetCode
	rspMsg.GetLoginRsp().Desc = rsp.Desc

	rData := new(rspData)
	rData.setData(addr, addr, rspMsg)
	cm.chRsp <- rData
}

func handleRegisterRsp(addr string, msg *csMsg.Msg, cm *connsManager) {
	flog.Debug("handleRegisterRsp msg=%v", *msg)
}

func handleRegisterReq(addr string, msg *csMsg.Msg, cm *connsManager) {
	flog.Debug("handleRegisterReq addr=%v, msg=%v", addr, msg)

	rsp, err := rpcClient.getCli().Register(context.Background(), &rpcPb.RegisterReq{Account: msg.GetRegisterReq().Account, Password: msg.GetRegisterReq().Password})
	if err != nil {
		flog.Error("[gatewaySvr] Register account:%v, password:%v err:%v", msg.GetRegisterReq().Account, msg.GetRegisterReq().Password, err)

		rData := new(rspData)
		rData.setData(addr, addr,
			&csMsg.Msg{
				ID: csMsg.MsgID_registerRsp,
				RegisterRsp: &csMsg.MsgRegisterRsp{
					RetCode: 1,
					Desc:    "Network exception, please try again later.",
				},
			},
		)
		cm.chRsp <- rData
		return
	}

	rspMsg := new(csMsg.Msg)
	rspMsg.ID = csMsg.MsgID_registerRsp
	rspMsg.RegisterRsp = new(csMsg.MsgRegisterRsp)
	rspMsg.GetRegisterRsp().RetCode = rsp.RetCode
	rspMsg.GetRegisterRsp().Desc = rsp.Desc

	rData := new(rspData)
	rData.setData(addr, addr, rspMsg)
	cm.chRsp <- rData
}

type rspData struct {
	targetAddr string
	fromAddr   string
	data       *csMsg.Msg
}

func (rd *rspData) setData(fromAddr string, targetAddr string, data *csMsg.Msg) {
	rd.targetAddr = targetAddr
	rd.fromAddr = fromAddr
	rd.data = data
}
