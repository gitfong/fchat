package main

import (
	_ "bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	_ "strings"

	csMsg "fchat/protos2Go"
	"github.com/golang/protobuf/proto"
)

func init() {
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)

}

func main() {
	fmt.Println("Launching server")

	listenAdd := ":9090"
	tcpAddr, err := net.ResolveTCPAddr("tcp", listenAdd)
	if err != nil {
		fmt.Println("resolveTCPAddr err:", err)
		os.Exit(0)
	}

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("listen on addr:", tcpAddr)
	}

	connMng := new(connsManager)
	connMng.init()
	go connMng.run()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}

		go handldConn(conn, connMng)

		fmt.Println("goroutine count:", runtime.NumGoroutine())
	}

}

func handldConn(c net.Conn, connMng *connsManager) {
	defer c.Close()

	addr := c.RemoteAddr().String()
	fmt.Printf("new conn: '%s'\n", addr)

	connMng.addConn(c)
	
	buf := make([]byte, 4096, 4096)

	for {
		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)

			connMng.chDel <- addr
			return
		}

		msg := &csMsg.Msg{}
		pData := buf[:cnt]
		err = proto.Unmarshal(pData, msg)
		if err != nil {
			fmt.Println("proto.Unmarshal err:", err)
			continue
		}

		fmt.Printf("from '%s',cnt:%d,%v\n", addr, cnt, msg)

		if MsgHandleFunc[msg.ID] != nil {
			//go MsgHandleFunc[msg.ID](addr, msg, connMng)
			MsgHandleFunc[msg.ID](addr, msg, connMng)
		} else {
			fmt.Printf("MsgHandleFunc[%d] is nil\n", msg.ID)
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

func (cm *connsManager) addConn(c net.Conn){
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
			fmt.Printf("add '%v' to cm.\n", info.addr)
		case str := <-cm.chDel:
			delete(cm.conns, str)
			fmt.Printf("delete '%v' from cm.\n", str)
		case rsp := <-cm.chRsp:
			pData, err := proto.Marshal(rsp.data)
			if err != nil {
				fmt.Println("on connsManager run. mashal err:", err)
				continue
			}

			if info, ok := cm.conns[rsp.targetAddr]; ok {
				info.conn.Write(pData)
				fmt.Printf("rsp data to '%v'\n", rsp.targetAddr)

			} 
		case notifyData := <-cm.chNotify:
			pData, err := proto.Marshal(notifyData.data)
			if err != nil {
				fmt.Println("on connsManager run. mashal err:", err)
				continue
			}

			for key, val := range cm.conns {
				val.conn.Write(pData)
				fmt.Println("notify to ", key)
			}
		}
	}
}
