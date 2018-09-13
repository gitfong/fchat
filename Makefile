rpcproto:
	protoc -I ${GOPATH}/src/fchat/protos/rpcProtos --go_out=plugins=grpc:${GOPATH}/src/fchat/protos2Go/rpc ${GOPATH}/src/fchat/protos/rpcProtos/*.proto

proto:
	protoc -I ${GOPATH}/src/fchat/protos --go_out=${GOPATH}/src/fchat/protos2Go ${GOPATH}/src/fchat/protos/*.proto
