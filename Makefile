all: proto rpcproto gateway account status

allsvr: gateway account status

rpcproto:
	protoc -I ${GOPATH}/src/fchat/protos/rpcProtos --go_out=plugins=grpc:${GOPATH}/src/fchat/protos2Go/rpc ${GOPATH}/src/fchat/protos/rpcProtos/*.proto

proto:
	protoc -I ${GOPATH}/src/fchat/protos --go_out=${GOPATH}/src/fchat/protos2Go ${GOPATH}/src/fchat/protos/*.proto

gateway:
	go install ./fchat-svr/fchat.gateway/

account:
	go install ./fchat-svr/fchat.account/

status:
	go install ./fchat-svr/fchat.status/
