package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

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

	buf := make([]byte, 4096, 4096)
	for {

		cnt, err := conn.Read(buf)
		if checkErr(err, "网络断开") {
			return
		}
		rcData := buf[:cnt]
		dataReceive := &csMsg.Msg{}
		err = proto.Unmarshal(rcData, dataReceive)
		if err != nil {
			fmt.Println("proto.Unmarshal err:", err)
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

var (
	cmdRegister = 1
	cmdLogin    = 2
	cmdLogout   = 3
	cmdSignIn   = 4
	cmdChat     = 5
)

func onKeyboardEvent(conn net.Conn) {
	fmt.Printf(`【操作提示】：
	输入%d注册帐号；
	输入%d登录帐号；
	输入%d退出登录；
	输入%d签到；
	输入%d进入聊天室；
	按回车确定输入：
`, cmdRegister, cmdLogin, cmdLogout, cmdSignIn, cmdChat)

	reader := bufio.NewReader(os.Stdin)
	for {
		cmdStr, err := reader.ReadString('\n')
		if checkErr(err, "输入错误") {
			continue
		}
		cmdStr = getRealStr(cmdStr)

		cmdInt, err := strconv.Atoi(cmdStr)
		if checkErr(err, "请输入正确的指令数字。") {
			continue
		}

		switch cmdInt {
		case cmdRegister:
			fmt.Print("请输入要注册的帐号名：")
			accStr, err := reader.ReadString('\n')
			if checkErr(err, "输入帐号名失败") {
				continue
			}
			accStr = getRealStr(accStr)

			fmt.Print("请输入注册密码：")
			pwStr, err := reader.ReadString('\n')
			if checkErr(err, "输入密码失败") {
				continue
			}
			pwStr = getRealStr(pwStr)

			msg := &csMsg.Msg{
				ID: csMsg.MsgID_registerReq,
				RegisterReq: &csMsg.MsgRegisterReq{
					Account:  accStr,
					Password: pwStr,
				},
			}
			fmt.Println("msg=", *msg)
			pData, err := proto.Marshal(msg)
			if checkErr(err, "网络错误") {
				continue
			}
			conn.Write(pData)
		case cmdLogin:
			fmt.Println("请输入要登录的帐号名：")
			accStr, err := reader.ReadString('\n')
			if checkErr(err, "输入帐号名失败") {
				continue
			}
			accStr = getRealStr(accStr)

			fmt.Print("请输入登录密码：")
			pwStr, err := reader.ReadString('\n')
			if checkErr(err, "输入密码失败") {
				continue
			}
			pwStr = getRealStr(pwStr)

			msg := &csMsg.Msg{
				ID: csMsg.MsgID_loginReq,
				LoginReq: &csMsg.MsgLoginReq{
					Account:  accStr,
					Password: pwStr,
				},
			}
			pData, err := proto.Marshal(msg)
			if checkErr(err, "网络错误") {
				continue
			}
			conn.Write(pData)
		case cmdSignIn:
			msg := &csMsg.Msg{
				ID:        csMsg.MsgID_signInReq,
				SignInReq: &csMsg.MsgSignInReq{},
			}
			pData, err := proto.Marshal(msg)
			if checkErr(err, "网络错误") {
				continue
			}
			conn.Write(pData)
		default:
			fmt.Printf("unknow cmdID=%v\n", cmdInt)
		}
	}
}

func checkErr(e error, str string) bool {
	if e != nil {
		fmt.Println("error:", e)
		fmt.Println("[操作失败]:", str)
		return true
	}
	return false
}

func getRealStr(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\r\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}
