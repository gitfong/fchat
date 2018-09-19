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
	return &rpcPb.SignInRsp{}, nil
}

func (s *statusSvr) UpdateStatus(ctx context.Context, in *rpcPb.UpdateStatusReq) (*rpcPb.UpdateStatusRsp, error) {
	return &rpcPb.UpdateStatusRsp{}, nil
}
