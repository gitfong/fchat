// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway2Account.proto

package rpcProtos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterReq struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway2Account_7a1adb2517d797d4, []int{0}
}
func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (dst *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(dst, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *RegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterRsp struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	Uid                  int32    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRsp) Reset()         { *m = RegisterRsp{} }
func (m *RegisterRsp) String() string { return proto.CompactTextString(m) }
func (*RegisterRsp) ProtoMessage()    {}
func (*RegisterRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway2Account_7a1adb2517d797d4, []int{1}
}
func (m *RegisterRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRsp.Unmarshal(m, b)
}
func (m *RegisterRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRsp.Marshal(b, m, deterministic)
}
func (dst *RegisterRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRsp.Merge(dst, src)
}
func (m *RegisterRsp) XXX_Size() int {
	return xxx_messageInfo_RegisterRsp.Size(m)
}
func (m *RegisterRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRsp proto.InternalMessageInfo

func (m *RegisterRsp) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RegisterRsp) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *RegisterRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type LoginReq struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway2Account_7a1adb2517d797d4, []int{2}
}
func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (dst *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(dst, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginRsp struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	Uid                  int32    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRsp) Reset()         { *m = LoginRsp{} }
func (m *LoginRsp) String() string { return proto.CompactTextString(m) }
func (*LoginRsp) ProtoMessage()    {}
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_gateway2Account_7a1adb2517d797d4, []int{3}
}
func (m *LoginRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRsp.Unmarshal(m, b)
}
func (m *LoginRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRsp.Marshal(b, m, deterministic)
}
func (dst *LoginRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRsp.Merge(dst, src)
}
func (m *LoginRsp) XXX_Size() int {
	return xxx_messageInfo_LoginRsp.Size(m)
}
func (m *LoginRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRsp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRsp proto.InternalMessageInfo

func (m *LoginRsp) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *LoginRsp) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *LoginRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "rpcProtos.RegisterReq")
	proto.RegisterType((*RegisterRsp)(nil), "rpcProtos.RegisterRsp")
	proto.RegisterType((*LoginReq)(nil), "rpcProtos.LoginReq")
	proto.RegisterType((*LoginRsp)(nil), "rpcProtos.LoginRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRsp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRsp, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRsp, error) {
	out := new(RegisterRsp)
	err := c.cc.Invoke(ctx, "/rpcProtos.Account/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRsp, error) {
	out := new(LoginRsp)
	err := c.cc.Invoke(ctx, "/rpcProtos.Account/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	Register(context.Context, *RegisterReq) (*RegisterRsp, error)
	Login(context.Context, *LoginReq) (*LoginRsp, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcProtos.Account/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcProtos.Account/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcProtos.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Account_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway2Account.proto",
}

func init() {
	proto.RegisterFile("gateway2Account.proto", fileDescriptor_gateway2Account_7a1adb2517d797d4)
}

var fileDescriptor_gateway2Account_7a1adb2517d797d4 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0x34, 0x72, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x2c, 0x2a, 0x48, 0x0e, 0x00, 0xb1, 0x8a, 0x95, 0x9c, 0xb9, 0xb8, 0x83, 0x52,
	0xd3, 0x33, 0x8b, 0x4b, 0x52, 0x8b, 0x82, 0x52, 0x0b, 0x85, 0x24, 0xb8, 0xd8, 0x13, 0x21, 0x4a,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x29, 0x2e, 0x8e, 0x82, 0xc4, 0xe2,
	0xe2, 0xf2, 0xfc, 0xa2, 0x14, 0x09, 0x26, 0xb0, 0x14, 0x9c, 0xaf, 0xe4, 0x8b, 0x64, 0x48, 0x71,
	0x01, 0xc8, 0x90, 0xa2, 0xd4, 0x12, 0xe7, 0xfc, 0x94, 0x54, 0xb0, 0x21, 0xac, 0x41, 0x30, 0xae,
	0x90, 0x00, 0x17, 0x73, 0x69, 0x26, 0x44, 0x3f, 0x6b, 0x10, 0x88, 0x29, 0x24, 0xc4, 0xc5, 0x92,
	0x92, 0x5a, 0x9c, 0x2c, 0xc1, 0x0c, 0x36, 0x12, 0xcc, 0x56, 0x72, 0xe0, 0xe2, 0xf0, 0xc9, 0x4f,
	0xcf, 0xcc, 0x23, 0xdf, 0x41, 0x5e, 0x30, 0x13, 0x28, 0x77, 0x8d, 0x51, 0x0d, 0x17, 0x3b, 0x34,
	0xf4, 0x84, 0x6c, 0xb8, 0x38, 0x60, 0xfe, 0x14, 0x12, 0xd3, 0x83, 0x07, 0xa2, 0x1e, 0x52, 0x08,
	0x4a, 0x61, 0x15, 0x2f, 0x2e, 0x50, 0x62, 0x10, 0x32, 0xe6, 0x62, 0x05, 0x3b, 0x4a, 0x48, 0x18,
	0x49, 0x09, 0xcc, 0xa3, 0x52, 0x98, 0x82, 0x20, 0x4d, 0x49, 0x6c, 0xe0, 0x18, 0x33, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x97, 0xd9, 0x2b, 0x71, 0xca, 0x01, 0x00, 0x00,
}
