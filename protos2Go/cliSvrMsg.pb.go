// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cliSvrMsg.proto

package cliSvrMsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MsgID int32

const (
	MsgID_begin MsgID = 0
	// 账号状态管理相关协议id 1~1000
	MsgID_registerReq MsgID = 1
	MsgID_registerRsp MsgID = 2
	MsgID_loginReq    MsgID = 5
	MsgID_loginRsp    MsgID = 6
	MsgID_signInReq   MsgID = 7
	MsgID_signInRsp   MsgID = 8
)

var MsgID_name = map[int32]string{
	0: "begin",
	1: "registerReq",
	2: "registerRsp",
	5: "loginReq",
	6: "loginRsp",
	7: "signInReq",
	8: "signInRsp",
}
var MsgID_value = map[string]int32{
	"begin":       0,
	"registerReq": 1,
	"registerRsp": 2,
	"loginReq":    5,
	"loginRsp":    6,
	"signInReq":   7,
	"signInRsp":   8,
}

func (x MsgID) String() string {
	return proto.EnumName(MsgID_name, int32(x))
}
func (MsgID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cliSvrMsg_99da4b4836941f64, []int{0}
}

type Msg struct {
	ID                   MsgID           `protobuf:"varint,1,opt,name=ID,proto3,enum=cliSvrMsg.MsgID" json:"ID,omitempty"`
	RegisterReq          *MsgRegisterReq `protobuf:"bytes,2,opt,name=registerReq,proto3" json:"registerReq,omitempty"`
	RegisterRsp          *MsgRegisterRsp `protobuf:"bytes,3,opt,name=registerRsp,proto3" json:"registerRsp,omitempty"`
	LoginReq             *MsgLoginReq    `protobuf:"bytes,4,opt,name=loginReq,proto3" json:"loginReq,omitempty"`
	LoginRsp             *MsgLoginRsp    `protobuf:"bytes,5,opt,name=loginRsp,proto3" json:"loginRsp,omitempty"`
	SignInReq            *MsgSignInReq   `protobuf:"bytes,6,opt,name=signInReq,proto3" json:"signInReq,omitempty"`
	SignInRsp            *MsgSignInRsp   `protobuf:"bytes,7,opt,name=signInRsp,proto3" json:"signInRsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_cliSvrMsg_99da4b4836941f64, []int{0}
}
func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (dst *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(dst, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetID() MsgID {
	if m != nil {
		return m.ID
	}
	return MsgID_begin
}

func (m *Msg) GetRegisterReq() *MsgRegisterReq {
	if m != nil {
		return m.RegisterReq
	}
	return nil
}

func (m *Msg) GetRegisterRsp() *MsgRegisterRsp {
	if m != nil {
		return m.RegisterRsp
	}
	return nil
}

func (m *Msg) GetLoginReq() *MsgLoginReq {
	if m != nil {
		return m.LoginReq
	}
	return nil
}

func (m *Msg) GetLoginRsp() *MsgLoginRsp {
	if m != nil {
		return m.LoginRsp
	}
	return nil
}

func (m *Msg) GetSignInReq() *MsgSignInReq {
	if m != nil {
		return m.SignInReq
	}
	return nil
}

func (m *Msg) GetSignInRsp() *MsgSignInRsp {
	if m != nil {
		return m.SignInRsp
	}
	return nil
}

type MsgRegisterReq struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgRegisterReq) Reset()         { *m = MsgRegisterReq{} }
func (m *MsgRegisterReq) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterReq) ProtoMessage()    {}
func (*MsgRegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cliSvrMsg_99da4b4836941f64, []int{1}
}
func (m *MsgRegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRegisterReq.Unmarshal(m, b)
}
func (m *MsgRegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRegisterReq.Marshal(b, m, deterministic)
}
func (dst *MsgRegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterReq.Merge(dst, src)
}
func (m *MsgRegisterReq) XXX_Size() int {
	return xxx_messageInfo_MsgRegisterReq.Size(m)
}
func (m *MsgRegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterReq proto.InternalMessageInfo

func (m *MsgRegisterReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *MsgRegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type MsgRegisterRsp struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgRegisterRsp) Reset()         { *m = MsgRegisterRsp{} }
func (m *MsgRegisterRsp) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterRsp) ProtoMessage()    {}
func (*MsgRegisterRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cliSvrMsg_99da4b4836941f64, []int{2}
}
func (m *MsgRegisterRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRegisterRsp.Unmarshal(m, b)
}
func (m *MsgRegisterRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRegisterRsp.Marshal(b, m, deterministic)
}
func (dst *MsgRegisterRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterRsp.Merge(dst, src)
}
func (m *MsgRegisterRsp) XXX_Size() int {
	return xxx_messageInfo_MsgRegisterRsp.Size(m)
}
func (m *MsgRegisterRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterRsp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterRsp proto.InternalMessageInfo

func (m *MsgRegisterRsp) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *MsgRegisterRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type MsgLoginReq struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgLoginReq) Reset()         { *m = MsgLoginReq{} }
func (m *MsgLoginReq) String() string { return proto.CompactTextString(m) }
func (*MsgLoginReq) ProtoMessage()    {}
func (*MsgLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cliSvrMsg_99da4b4836941f64, []int{3}
}
func (m *MsgLoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgLoginReq.Unmarshal(m, b)
}
func (m *MsgLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgLoginReq.Marshal(b, m, deterministic)
}
func (dst *MsgLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLoginReq.Merge(dst, src)
}
func (m *MsgLoginReq) XXX_Size() int {
	return xxx_messageInfo_MsgLoginReq.Size(m)
}
func (m *MsgLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLoginReq proto.InternalMessageInfo

func (m *MsgLoginReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *MsgLoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type MsgLoginRsp struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Uid                  int32    `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgLoginRsp) Reset()         { *m = MsgLoginRsp{} }
func (m *MsgLoginRsp) String() string { return proto.CompactTextString(m) }
func (*MsgLoginRsp) ProtoMessage()    {}
func (*MsgLoginRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cliSvrMsg_99da4b4836941f64, []int{4}
}
func (m *MsgLoginRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgLoginRsp.Unmarshal(m, b)
}
func (m *MsgLoginRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgLoginRsp.Marshal(b, m, deterministic)
}
func (dst *MsgLoginRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLoginRsp.Merge(dst, src)
}
func (m *MsgLoginRsp) XXX_Size() int {
	return xxx_messageInfo_MsgLoginRsp.Size(m)
}
func (m *MsgLoginRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLoginRsp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLoginRsp proto.InternalMessageInfo

func (m *MsgLoginRsp) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *MsgLoginRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *MsgLoginRsp) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type MsgSignInReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgSignInReq) Reset()         { *m = MsgSignInReq{} }
func (m *MsgSignInReq) String() string { return proto.CompactTextString(m) }
func (*MsgSignInReq) ProtoMessage()    {}
func (*MsgSignInReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cliSvrMsg_99da4b4836941f64, []int{5}
}
func (m *MsgSignInReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSignInReq.Unmarshal(m, b)
}
func (m *MsgSignInReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSignInReq.Marshal(b, m, deterministic)
}
func (dst *MsgSignInReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignInReq.Merge(dst, src)
}
func (m *MsgSignInReq) XXX_Size() int {
	return xxx_messageInfo_MsgSignInReq.Size(m)
}
func (m *MsgSignInReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignInReq.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignInReq proto.InternalMessageInfo

type MsgSignInRsp struct {
	Retcode              int32    `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	SignInCount          int32    `protobuf:"varint,3,opt,name=signInCount,proto3" json:"signInCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgSignInRsp) Reset()         { *m = MsgSignInRsp{} }
func (m *MsgSignInRsp) String() string { return proto.CompactTextString(m) }
func (*MsgSignInRsp) ProtoMessage()    {}
func (*MsgSignInRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cliSvrMsg_99da4b4836941f64, []int{6}
}
func (m *MsgSignInRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSignInRsp.Unmarshal(m, b)
}
func (m *MsgSignInRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSignInRsp.Marshal(b, m, deterministic)
}
func (dst *MsgSignInRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignInRsp.Merge(dst, src)
}
func (m *MsgSignInRsp) XXX_Size() int {
	return xxx_messageInfo_MsgSignInRsp.Size(m)
}
func (m *MsgSignInRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignInRsp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignInRsp proto.InternalMessageInfo

func (m *MsgSignInRsp) GetRetcode() int32 {
	if m != nil {
		return m.Retcode
	}
	return 0
}

func (m *MsgSignInRsp) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *MsgSignInRsp) GetSignInCount() int32 {
	if m != nil {
		return m.SignInCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Msg)(nil), "cliSvrMsg.Msg")
	proto.RegisterType((*MsgRegisterReq)(nil), "cliSvrMsg.MsgRegisterReq")
	proto.RegisterType((*MsgRegisterRsp)(nil), "cliSvrMsg.MsgRegisterRsp")
	proto.RegisterType((*MsgLoginReq)(nil), "cliSvrMsg.MsgLoginReq")
	proto.RegisterType((*MsgLoginRsp)(nil), "cliSvrMsg.MsgLoginRsp")
	proto.RegisterType((*MsgSignInReq)(nil), "cliSvrMsg.MsgSignInReq")
	proto.RegisterType((*MsgSignInRsp)(nil), "cliSvrMsg.MsgSignInRsp")
	proto.RegisterEnum("cliSvrMsg.MsgID", MsgID_name, MsgID_value)
}

func init() { proto.RegisterFile("cliSvrMsg.proto", fileDescriptor_cliSvrMsg_99da4b4836941f64) }

var fileDescriptor_cliSvrMsg_99da4b4836941f64 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xb1, 0x6f, 0xf2, 0x30,
	0x10, 0xc5, 0xbf, 0x04, 0x02, 0xe4, 0xc2, 0x07, 0xd6, 0x0d, 0x6d, 0xda, 0x29, 0xca, 0x84, 0x3a,
	0x30, 0x50, 0x75, 0xaa, 0xd4, 0x05, 0x54, 0x29, 0x52, 0xb3, 0x98, 0xbd, 0x12, 0x24, 0x91, 0x15,
	0x09, 0x25, 0x6e, 0x2e, 0xb4, 0x7f, 0x7a, 0xd7, 0x0a, 0x03, 0xb1, 0x3d, 0x80, 0xd4, 0x6e, 0x7e,
	0xe7, 0xf7, 0x7b, 0x67, 0x8e, 0x0b, 0x4c, 0xb3, 0x5d, 0xb9, 0xfe, 0x6c, 0x52, 0x12, 0x73, 0xd9,
	0xd4, 0x6d, 0x8d, 0x7e, 0x57, 0x88, 0xbf, 0x5d, 0xe8, 0xa5, 0x24, 0x30, 0x02, 0x37, 0x59, 0x85,
	0x4e, 0xe4, 0xcc, 0x26, 0x0b, 0x36, 0xd7, 0x40, 0x4a, 0x22, 0x59, 0x71, 0x37, 0x59, 0xe1, 0x33,
	0x04, 0x4d, 0x21, 0x4a, 0x6a, 0x8b, 0x86, 0x17, 0x1f, 0xa1, 0x1b, 0x39, 0xb3, 0x60, 0x71, 0x67,
	0x5b, 0xb9, 0x36, 0x70, 0xd3, 0x6d, 0xc1, 0x24, 0xc3, 0xde, 0x55, 0x98, 0x24, 0x37, 0xdd, 0xb8,
	0x80, 0xd1, 0xae, 0x16, 0x65, 0x75, 0x68, 0xdb, 0x57, 0xe4, 0x8d, 0x4d, 0xbe, 0x9d, 0x6e, 0x79,
	0xe7, 0xd3, 0x0c, 0xc9, 0xd0, 0xbb, 0xcc, 0x90, 0xe4, 0x9d, 0x0f, 0x9f, 0xc0, 0xa7, 0x52, 0x54,
	0x89, 0x6a, 0x34, 0x50, 0xd0, 0xad, 0x0d, 0xad, 0xcf, 0xd7, 0x5c, 0x3b, 0x0d, 0x8c, 0x64, 0x38,
	0xbc, 0x82, 0x91, 0xe4, 0xda, 0x19, 0xbf, 0xc2, 0xc4, 0x9e, 0x18, 0x86, 0x30, 0xdc, 0x64, 0x59,
	0xbd, 0xaf, 0x5a, 0xf5, 0x47, 0xf8, 0xfc, 0x2c, 0xf1, 0x1e, 0x46, 0x72, 0x43, 0xf4, 0x55, 0x37,
	0xb9, 0x1a, 0xbc, 0xcf, 0x3b, 0x1d, 0xbf, 0xd8, 0x39, 0x24, 0x0f, 0x39, 0x4d, 0xd1, 0x2e, 0xeb,
	0xbc, 0x50, 0x39, 0x1e, 0x3f, 0x4b, 0x44, 0xe8, 0xe7, 0x05, 0x65, 0xa7, 0x0c, 0x75, 0x8e, 0x97,
	0x10, 0x18, 0x23, 0xfc, 0xe3, 0x23, 0x52, 0x23, 0xe4, 0xb7, 0x2f, 0x40, 0x06, 0xbd, 0x7d, 0x99,
	0xab, 0xa5, 0xf0, 0xf8, 0xe1, 0x18, 0x4f, 0x60, 0x6c, 0x4e, 0x3b, 0x7e, 0x37, 0x75, 0x97, 0x9f,
	0xd9, 0xf9, 0xd9, 0xa5, 0xfc, 0x08, 0x82, 0xe3, 0xd8, 0x97, 0xea, 0x67, 0x1d, 0xfb, 0x98, 0xa5,
	0x87, 0x0a, 0x3c, 0xb5, 0xe8, 0xe8, 0x83, 0xb7, 0x2d, 0x44, 0x59, 0xb1, 0x7f, 0x38, 0xb5, 0xf6,
	0x9d, 0x39, 0x56, 0x81, 0x24, 0x73, 0x71, 0xac, 0xf7, 0x92, 0x79, 0x5a, 0x91, 0x64, 0x03, 0xfc,
	0x6f, 0xec, 0x12, 0x1b, 0x1a, 0x92, 0x24, 0x1b, 0x6d, 0x07, 0xea, 0x3b, 0x7c, 0xfc, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x89, 0x27, 0x9d, 0x5b, 0x9a, 0x03, 0x00, 0x00,
}
