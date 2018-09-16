package main

import (
	_ "bufio"
	"fmt"
	"net"
	"runtime"
	_ "strings"

	csMsg "fchat/protos2Go"
	"github.com/golang/protobuf/proto"

	"github.com/gitfong/fLog"
)

var flog *fLog.FLogger

func init() {
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)
	flog = fLog.New()
	if flog == nil {
		panic("fLog.New fail")
	}
}

func main() {
	flog.Info("launching server")

	listenAdd := ":9090"
	tcpAddr, err := net.ResolveTCPAddr("tcp", listenAdd)
	if err != nil {
		flog.Fatal("ResolveTCPAddr err:%v", listenAdd)
	}

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		flog.Fatal("listen on addr:", tcpAddr)
	}

	connMng := new(connsManager)
	connMng.init()
	go connMng.run()

	fmt.Println("server start finished")

	for {
		conn, err := ln.Accept()
		if err != nil {
			flog.Error("accept err:", err)
			continue
		}

		go handldConn(conn, connMng)

		flog.Debug("goroutine count:%d", runtime.NumGoroutine())
	}

}

func handldConn(c net.Conn, connMng *connsManager) {
	defer c.Close()

	addr := c.RemoteAddr().String()
	flog.Info("new conn from : '%s'", addr)

	connMng.addConn(c)

	buf := make([]byte, 4096, 4096)

	for {
		cnt, err := c.Read(buf)
		if err != nil {
			flog.Error("read err:%v", err)

			connMng.chDel <- addr
			return
		}

		msg := &csMsg.Msg{}
		pData := buf[:cnt]
		err = proto.Unmarshal(pData, msg)
		if err != nil {
			flog.Error("proto.Unmarshal err:%v", err)
			continue
		}

		flog.Debug("from '%s',msgLen:%d,%v", addr, cnt, msg)

		if MsgHandleFunc[msg.ID] != nil {
			//go MsgHandleFunc[msg.ID](addr, msg, connMng)
			MsgHandleFunc[msg.ID](addr, msg, connMng)
		} else {
			flog.Warn("MsgHandleFunc[%d] is nil", msg.ID)
		}

	}
}

type connInfo struct {
	conn net.Conn
	addr string //用户连接地址
	uid  int32  //用户唯一id
}

type connsManager struct {
	conns    map[string]*connInfo
	chAdd    chan *connInfo
	chDel    chan string
	chRsp    chan *rspData
	chNotify chan *rspData
}

func (cm *connsManager) init() {
	cm.conns = make(map[string]*connInfo)
	cm.chAdd = make(chan *connInfo)
	cm.chDel = make(chan string)
	cm.chRsp = make(chan *rspData)
	cm.chNotify = make(chan *rspData, 100)
}

func (cm *connsManager) addConn(c net.Conn) {
	cInfo := &connInfo{
		conn: c,
		addr: c.RemoteAddr().String(),
		uid:  0,
	}
	cm.chAdd <- cInfo
}

func (cm *connsManager) run() {
	for {
		select {
		case info := <-cm.chAdd:
			cm.conns[info.addr] = info
			flog.Debug("add '%v' to cm.", info.addr)
		case str := <-cm.chDel:
			delete(cm.conns, str)
			flog.Debug("delete '%v' from cm.", str)
		case rsp := <-cm.chRsp:
			pData, err := proto.Marshal(rsp.data)
			if err != nil {
				flog.Error("on connsManager run. mashal err:%v", err)
				continue
			}

			if info, ok := cm.conns[rsp.targetAddr]; ok {
				info.conn.Write(pData)
				flog.Debug("rsp data to '%v'", rsp.targetAddr)
			}
		case notifyData := <-cm.chNotify:
			pData, err := proto.Marshal(notifyData.data)
			if err != nil {
				flog.Error("on connsManager run. mashal err:%v", err)
				continue
			}

			for key, val := range cm.conns {
				val.conn.Write(pData)
				flog.Debug("notify to '%v'", key)
			}
		}
	}
}
