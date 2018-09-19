package main

import (
	rpcPb "fchat/protos2Go/rpc"
	"strconv"
	"time"

	"fchat/fchat-svr/fchat.status/redisCli"
	"github.com/garyburd/redigo/redis"
	"golang.org/x/net/context"
)

var beginSinginDate time.Time

func init() {
	beginSinginDate, _ = time.Parse("2006-01-02", "2018-09-09")
}

type statusSvr struct{}

func (s *statusSvr) Heartbeat(ctx context.Context, in *rpcPb.Heartbeat2StatusReq) (*rpcPb.Heartbeat2StatusRsp, error) {
	return &rpcPb.Heartbeat2StatusRsp{}, nil
}

func (s *statusSvr) SignIn(ctx context.Context, in *rpcPb.SignInReq) (*rpcPb.SignInRsp, error) {
	flog.Debug("signin uid:%d", in.Uid)

	offsetDay := int(time.Now().Sub(beginSinginDate).Hours() / 24)
	flog.Debug("offsetDay:%d", offsetDay)

	c := redisCli.Get()
	isSign, err := redis.Int(c.Do("getbit", getSignInKey(in.Uid), offsetDay))
	if err != nil {
		flog.Error("signIn. getbit key:%v, offset:%d, err:%v", getSignInKey(in.Uid), offsetDay, err)
		return &rpcPb.SignInRsp{RetCode: 1, Desc: "signin fail", SignInCount: 0}, nil
	}

	if isSign > 0 {
		return &rpcPb.SignInRsp{RetCode: 2, Desc: "you have signin today.", SignInCount: 0}, nil
	}

	reply, err := c.Do("setbit", getSignInKey(in.Uid), offsetDay, 1)
	flog.Debug("signin. setbit reply:%v", reply)
	if err != nil {
		flog.Error("signIn. setbit key:%v, offset:%d, err:%v", getSignInKey(in.Uid), offsetDay, err)
		return &rpcPb.SignInRsp{RetCode: 3, Desc: "signin fail", SignInCount: 0}, nil
	}

	count, _ := redis.Int(c.Do("bitcount", getSignInKey(in.Uid)))
	return &rpcPb.SignInRsp{RetCode: 0, Desc: "signin succeed", SignInCount: int32(count)}, nil
}

func (s *statusSvr) UpdateStatus(ctx context.Context, in *rpcPb.UpdateStatusReq) (*rpcPb.UpdateStatusRsp, error) {
	flog.Debug("updatestatus uid:%d,status:%d", in.Uid, in.Status)
	return &rpcPb.UpdateStatusRsp{RetCode: 0}, nil
}

func getSignInKey(uid int32) string {
	return "signIn_" + strconv.Itoa(int(uid))
}
