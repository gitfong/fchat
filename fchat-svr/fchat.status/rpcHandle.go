package main

import (
	rpcPb "fchat/protos2Go/rpc"

	"golang.org/x/net/context"
)

type statusSvr struct{}

func (s *statusSvr) Heartbeat(ctx context.Context, in *rpcPb.Heartbeat2StatusReq) (*rpcPb.Heartbeat2StatusRsp, error) {
	return &rpcPb.Heartbeat2StatusRsp{}, nil
}

func (s *statusSvr) SignIn(ctx context.Context, in *rpcPb.SignInReq) (*rpcPb.SignInRsp, error) {
	flog.Debug("signin uid:%d", in.Uid)
	return &rpcPb.SignInRsp{RetCode: 0, Desc: "签到成功"}, nil
}

func (s *statusSvr) UpdateStatus(ctx context.Context, in *rpcPb.UpdateStatusReq) (*rpcPb.UpdateStatusRsp, error) {
	flog.Debug("updatestatus uid:%d,status:%d", in.Uid, in.Status)
	return &rpcPb.UpdateStatusRsp{RetCode: 0}, nil
}
