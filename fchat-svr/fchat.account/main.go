package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net"
	_ "time"

	rpcPb "fchat/protos2Go/rpc"

	"github.com/gitfong/fLog"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

const (
	accountSvrAddr = ":9091"
)

var flog *fLog.FLogger
var dbw = new(dbWorker)

func init() {
	flog = fLog.New()
	if flog == nil {
		panic("fLog.New fail")
	}
}

func main() {
	//init dbworker
	dbw.init()
	defer dbw.Db.Close()

	lis, err := net.Listen("tcp", accountSvrAddr)
	if err != nil {
		flog.Fatal("failed to listen addr:%s, err:%v", accountSvrAddr, err)
	}

	rpcServer := grpc.NewServer()
	rpcPb.RegisterAccountServer(rpcServer, &accountServer{})
	rpcServer.Serve(lis)
}

type dbWorker struct {
	DataSrcName string `json:"DataSrcName"`
	Db          *sql.DB
}

func (w *dbWorker) init() {
	data, err := ioutil.ReadFile("dbCfgAccount.json")
	if err != nil {
		flog.Fatal("read dbCfg.json err:%v", err)
	}

	err = json.Unmarshal(data, w)
	if err != nil {
		flog.Fatal("json.Unmarshal err:%v", err)
	}

	w.Db, err = sql.Open("mysql", w.DataSrcName)
	if err != nil {
		flog.Fatal("sql.Open err:", err)
	}
}
