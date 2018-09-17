package main

import (
	rpcPb "fchat/protos2Go/rpc"
	"time"

	"golang.org/x/net/context"
)

type accountServer struct{}

func (s *accountServer) Heartbeat(ctx context.Context, in *rpcPb.HeartbeatReq) (*rpcPb.HeartbeatRsp, error) {
	return &rpcPb.HeartbeatRsp{}, nil
}

func (s *accountServer) Register(ctx context.Context, in *rpcPb.RegisterReq) (*rpcPb.RegisterRsp, error) {
	flog.Debug("[accountSvr] on Register %s,%s", in.Account, in.Password)

	stmt, err := dbw.Db.Prepare(`INSERT INTO accounts(account,password,registerTime,lastLoginTime)VALUES(?,?,?,?)`)
	defer stmt.Close()
	if err != nil {
		flog.Error("Register db.Prepare err:%v", err)
		return &rpcPb.RegisterRsp{RetCode: 1, Uid: 0, Desc: "register fail."}, err
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	ret, err := stmt.Exec(in.Account, in.Password, nowStr, nowStr)
	if err != nil {
		flog.Error("Register account:%s,password:%s,err:%v", in.Account, in.Password, err)
		return &rpcPb.RegisterRsp{RetCode: 2, Uid: 0, Desc: "register fail."}, err
	}

	var Uid int32 = 0
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		Uid = int32(LastInsertId)
		flog.Info("Register account:%s,lastInsertID:%d", in.Account, LastInsertId)
	}

	if RowsAffected, err := ret.RowsAffected(); nil == err {
		flog.Debug("Register account:%s,RowsAffected:%d", in.Account, RowsAffected)
	}
	return &rpcPb.RegisterRsp{RetCode: 0, Uid: Uid, Desc: "register succeed!"}, nil
}

func (s *accountServer) Login(ctx context.Context, in *rpcPb.LoginReq) (*rpcPb.LoginRsp, error) {
	flog.Debug("[accountSvr] on Login %s,%s", in.Account, in.Password)

	//check account and password
	var Uid int32 = 0
	err := dbw.Db.QueryRow(`SELECT id FROM accounts WHERE account=? AND passWord=?`).Scan(&Uid)
	if err != nil {
		flog.Error("Login account:%s,password:%s,err:%v", in.Account, in.Password, err)
		return &rpcPb.LoginRsp{RetCode: 1, Uid: 0, Desc: "login fail."}, err
	}
	if Uid <= 0 {
		flog.Info("Login account:%s not exist. Uid=%d", in.Account, Uid)
		return &rpcPb.LoginRsp{RetCode: 2, Uid: 0, Desc: "login fail.Account not exist."}, err
	}

	stmt, err := dbw.Db.Prepare(`UPDATE accounts SET lastLoginTime=? WHERE account=?`)
	defer stmt.Close()
	if err != nil {
		flog.Error("Login db.Prepare err:%v", err)
		return &rpcPb.LoginRsp{RetCode: 2, Uid: 0, Desc: "login fail."}, nil
	}
	ret, err := stmt.Exec(time.Now().Format("2006-01-02 15:04:05"), in.Account)
	if err != nil {
		flog.Error("Login check succeed. account:%s. but update lastLoginTime err:%v", in.Account, err)
		return &rpcPb.LoginRsp{RetCode: 3, Uid: 0, Desc: "login fail."}, nil
	}
	if rowAf, er := ret.RowsAffected(); nil != er || rowAf <= 0 {
		flog.Error("Login check succeed. account:%s. but update lastLoginTime err:%v, rowAf:%d", in.Account, err, rowAf)
		return &rpcPb.LoginRsp{RetCode: 4, Uid: 0, Desc: "login fail."}, nil
	}

	return &rpcPb.LoginRsp{RetCode: 0, Uid: Uid, Desc: "login succeed!"}, nil
}
