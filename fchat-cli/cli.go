package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"

	csMsg "fchat/protos2Go"
	"github.com/golang/protobuf/proto"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}

	go onKeyboardEvent(conn)

	buf := make([]byte,4096,4096)
	for {
		
		cnt,err := conn.Read(buf)
		if checkErr(err,"网络断开"){
			return
		}
		rcData := buf[:cnt]
		dataReceive := &csMsg.Msg{}
		err = proto.Unmarshal(rcData,dataReceive)
		if err != nil {
			fmt.Println("proto.Unmarshal err:",err)
		}
		
		fmt.Println("receive msg:", dataReceive)

		/*
			//listen for reply
			message, rErr := bufio.NewReader(conn).ReadString('\n')
			if rErr != nil {
				fmt.Println("rErr:",rErr)
				return
			}
			fmt.Println("message from server:" + message)
		*/
	}
}

func onKeyboardEvent(conn net.Conn){
	fmt.Println(`操作指令：
		输入1注册帐号；
		输入2登录帐号；
		输入3退出登录；
		输入4进入聊天室；
		按回车确定输入；`)

	reader := bufio.NewReader(os.Stdin)
	for{
		cmdStr, err := reader.ReadString('\n')
		if checkErr(err,"输入错误") {
			continue
		}
		cmdStr = strings.Replace(cmdStr, " ","", -1)
		cmdStr = strings.Replace(cmdStr, "\n","", -1)
		cmdStr = strings.Replace(cmdStr, "\r","", -1)

		cmdInt,err := strconv.Atoi(cmdStr)
		if checkErr(err,"请输入正确的指令数字。"){
			continue
		}
		
		switch cmdInt{
		case 1:
			fmt.Print("请输入要注册的帐号名：")
			accStr, err := reader.ReadString('\n')
			if checkErr(err,"输入帐号名失败"){
				continue
			}
			accStr = strings.Replace(accStr,"\r\n","",-1)

			fmt.Print("请输入注册密码：")
			pwStr,err := reader.ReadString('\n')
			if checkErr(err,"输入密码失败"){
				continue
			}
			pwStr = strings.Replace(pwStr,"\r\n","",-1)
			
			msg := &csMsg.Msg{
				ID:          csMsg.MsgID_registerReq,
				RegisterReq: &csMsg.MsgRegisterReq{
					Account:  accStr,
					Password: pwStr,
				},
			}
			fmt.Println("msg=", *msg)
			pData, err := proto.Marshal(msg)
			if checkErr(err,"网络错误"){
				continue
			}
			conn.Write(pData)
		case 2:
			fmt.Println("请输入要登录的帐号名：")
			accStr,err := reader.ReadString('\n')
			if checkErr(err,"输入帐号名失败"){
				continue
			}
			accStr = strings.Replace(accStr,"\r\n","",-1)

			fmt.Print("请输入登录密码：")
			pwStr,err := reader.ReadString('\n')
			if checkErr(err,"输入密码失败"){
				continue
			}
			pwStr = strings.Replace(pwStr,"\r\n","",-1)

			msg := &csMsg.Msg{
				ID:	csMsg.MsgID_loginReq,
				LoginReq: &csMsg.MsgLoginReq{
					Account: accStr,
					Password: pwStr,
				},
			}
			pData, err := proto.Marshal(msg)
			if checkErr(err,"网络错误"){
				continue
			}
			conn.Write(pData)
		default:
			fmt.Printf("unknow cmdID=%v\n",cmdInt)
		}
	}
}

func checkErr(e error, str string)bool{
	if e != nil {
		fmt.Println("error:",e)
		fmt.Println("[操作错误]:",str)
		return true
	}
	return false
}
