package main

import (
	rpcPb "fchat/protos2Go/rpc"

	"golang.org/x/net/context"
)

type statusSvr struct{}

func (s *statusSvr) SignIn(ctx context.Context, in *rpcPb.HeartbeatReq) (*rpcPb.HeartbeatRsp, error) {
	return &rpcPb.HeartbeatRsp{}, nil
}
