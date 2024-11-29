// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: login.proto

package out

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginResult int32

const (
	LoginResult_LOGIN_SUCCESS LoginResult = 0
	//* token过期
	LoginResult_TOKEN_EXPIRED LoginResult = 1
	//* 服务器没有开服
	LoginResult_WORLD_CLOSE LoginResult = 2
	//* 时间戳非法
	LoginResult_TIMESTAMP_VALID LoginResult = 3
	//* 参数错误
	LoginResult_PARAM_INVALID LoginResult = 4
	//* 服务器繁忙
	LoginResult_SERVER_IS_BUSY LoginResult = 5
	//* 注册已达上限
	LoginResult_REGISTER_LIMIT LoginResult = 6
	//* 服务器已满
	LoginResult_SERVER_IS_FULL LoginResult = 7
	//* 账号被封
	LoginResult_ACCOUNT_BAN LoginResult = 8
	//* 还未创角
	LoginResult_NO_ROLE LoginResult = 9
)

// Enum value maps for LoginResult.
var (
	LoginResult_name = map[int32]string{
		0: "LOGIN_SUCCESS",
		1: "TOKEN_EXPIRED",
		2: "WORLD_CLOSE",
		3: "TIMESTAMP_VALID",
		4: "PARAM_INVALID",
		5: "SERVER_IS_BUSY",
		6: "REGISTER_LIMIT",
		7: "SERVER_IS_FULL",
		8: "ACCOUNT_BAN",
		9: "NO_ROLE",
	}
	LoginResult_value = map[string]int32{
		"LOGIN_SUCCESS":   0,
		"TOKEN_EXPIRED":   1,
		"WORLD_CLOSE":     2,
		"TIMESTAMP_VALID": 3,
		"PARAM_INVALID":   4,
		"SERVER_IS_BUSY":  5,
		"REGISTER_LIMIT":  6,
		"SERVER_IS_FULL":  7,
		"ACCOUNT_BAN":     8,
		"NO_ROLE":         9,
	}
)

func (x LoginResult) Enum() *LoginResult {
	p := new(LoginResult)
	*p = x
	return p
}

func (x LoginResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginResult) Descriptor() protoreflect.EnumDescriptor {
	return file_login_proto_enumTypes[0].Descriptor()
}

func (LoginResult) Type() protoreflect.EnumType {
	return &file_login_proto_enumTypes[0]
}

func (x LoginResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LoginResult) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LoginResult(num)
	return nil
}

// Deprecated: Use LoginResult.Descriptor instead.
func (LoginResult) EnumDescriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

type CreateRet int32

const (
	CreateRet_OK            CreateRet = 0
	CreateRet_REPEATED_NAME CreateRet = 1 //重名
	CreateRet_OCCUPYING     CreateRet = 2 //该名字正在被占用
	CreateRet_ILLEGAL_NAME  CreateRet = 3 //非法名字
)

// Enum value maps for CreateRet.
var (
	CreateRet_name = map[int32]string{
		0: "OK",
		1: "REPEATED_NAME",
		2: "OCCUPYING",
		3: "ILLEGAL_NAME",
	}
	CreateRet_value = map[string]int32{
		"OK":            0,
		"REPEATED_NAME": 1,
		"OCCUPYING":     2,
		"ILLEGAL_NAME":  3,
	}
)

func (x CreateRet) Enum() *CreateRet {
	p := new(CreateRet)
	*p = x
	return p
}

func (x CreateRet) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateRet) Descriptor() protoreflect.EnumDescriptor {
	return file_login_proto_enumTypes[1].Descriptor()
}

func (CreateRet) Type() protoreflect.EnumType {
	return &file_login_proto_enumTypes[1]
}

func (x CreateRet) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CreateRet) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CreateRet(num)
	return nil
}

// Deprecated: Use CreateRet.Descriptor instead.
func (CreateRet) EnumDescriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

type KickoffReason int32

const (
	KickoffReason_ERROR   KickoffReason = 0 //异常
	KickoffReason_RELOGIN KickoffReason = 1 //顶号
	KickoffReason_GM      KickoffReason = 2 //系统踢下线
)

// Enum value maps for KickoffReason.
var (
	KickoffReason_name = map[int32]string{
		0: "ERROR",
		1: "RELOGIN",
		2: "GM",
	}
	KickoffReason_value = map[string]int32{
		"ERROR":   0,
		"RELOGIN": 1,
		"GM":      2,
	}
)

func (x KickoffReason) Enum() *KickoffReason {
	p := new(KickoffReason)
	*p = x
	return p
}

func (x KickoffReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KickoffReason) Descriptor() protoreflect.EnumDescriptor {
	return file_login_proto_enumTypes[2].Descriptor()
}

func (KickoffReason) Type() protoreflect.EnumType {
	return &file_login_proto_enumTypes[2]
}

func (x KickoffReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *KickoffReason) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = KickoffReason(num)
	return nil
}

// Deprecated: Use KickoffReason.Descriptor instead.
func (KickoffReason) EnumDescriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

type CS_Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  *string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`                   // 账号
	Serverid *uint32 `protobuf:"varint,2,opt,name=serverid" json:"serverid,omitempty"`                // 服务器id
	Device   *string `protobuf:"bytes,3,opt,name=device" json:"device,omitempty"`                     //设备类型、设备号等
	ClientIp *string `protobuf:"bytes,4,opt,name=client_ip,json=clientIp" json:"client_ip,omitempty"` //登录ip
	Location *string `protobuf:"bytes,5,opt,name=location" json:"location,omitempty"`                 //地理位置
}

func (x *CS_Login) Reset() {
	*x = CS_Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_Login) ProtoMessage() {}

func (x *CS_Login) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_Login.ProtoReflect.Descriptor instead.
func (*CS_Login) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *CS_Login) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

func (x *CS_Login) GetServerid() uint32 {
	if x != nil && x.Serverid != nil {
		return *x.Serverid
	}
	return 0
}

func (x *CS_Login) GetDevice() string {
	if x != nil && x.Device != nil {
		return *x.Device
	}
	return ""
}

func (x *CS_Login) GetClientIp() string {
	if x != nil && x.ClientIp != nil {
		return *x.ClientIp
	}
	return ""
}

func (x *CS_Login) GetLocation() string {
	if x != nil && x.Location != nil {
		return *x.Location
	}
	return ""
}

type SC_Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     *LoginResult `protobuf:"varint,1,opt,name=result,enum=am.protocol.LoginResult" json:"result,omitempty"` // 登陆结果
	TimeZone   *int32       `protobuf:"varint,2,opt,name=time_zone,json=timeZone" json:"time_zone,omitempty"`          // 服务器时区
	ServerTime *int64       `protobuf:"varint,3,opt,name=server_time,json=serverTime" json:"server_time,omitempty"`    //服务器时间戳
}

func (x *SC_Login) Reset() {
	*x = SC_Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_Login) ProtoMessage() {}

func (x *SC_Login) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_Login.ProtoReflect.Descriptor instead.
func (*SC_Login) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *SC_Login) GetResult() LoginResult {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return LoginResult_LOGIN_SUCCESS
}

func (x *SC_Login) GetTimeZone() int32 {
	if x != nil && x.TimeZone != nil {
		return *x.TimeZone
	}
	return 0
}

func (x *SC_Login) GetServerTime() int64 {
	if x != nil && x.ServerTime != nil {
		return *x.ServerTime
	}
	return 0
}

type CS_Create struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type *int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
}

func (x *CS_Create) Reset() {
	*x = CS_Create{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_Create) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_Create) ProtoMessage() {}

func (x *CS_Create) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_Create.ProtoReflect.Descriptor instead.
func (*CS_Create) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

func (x *CS_Create) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CS_Create) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

type SC_Create struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret *CreateRet `protobuf:"varint,1,opt,name=ret,enum=am.protocol.CreateRet" json:"ret,omitempty"`
}

func (x *SC_Create) Reset() {
	*x = SC_Create{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_Create) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_Create) ProtoMessage() {}

func (x *SC_Create) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_Create.ProtoReflect.Descriptor instead.
func (*SC_Create) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{3}
}

func (x *SC_Create) GetRet() CreateRet {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return CreateRet_OK
}

type CS_Ready struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Placeholder *int32 `protobuf:"varint,1,opt,name=placeholder" json:"placeholder,omitempty"` //占坑
}

func (x *CS_Ready) Reset() {
	*x = CS_Ready{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_Ready) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_Ready) ProtoMessage() {}

func (x *CS_Ready) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_Ready.ProtoReflect.Descriptor instead.
func (*CS_Ready) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{4}
}

func (x *CS_Ready) GetPlaceholder() int32 {
	if x != nil && x.Placeholder != nil {
		return *x.Placeholder
	}
	return 0
}

type SC_Ready struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Placeholder *int32 `protobuf:"varint,1,opt,name=placeholder" json:"placeholder,omitempty"` //占坑
}

func (x *SC_Ready) Reset() {
	*x = SC_Ready{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_Ready) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_Ready) ProtoMessage() {}

func (x *SC_Ready) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_Ready.ProtoReflect.Descriptor instead.
func (*SC_Ready) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{5}
}

func (x *SC_Ready) GetPlaceholder() int32 {
	if x != nil && x.Placeholder != nil {
		return *x.Placeholder
	}
	return 0
}

type SC_Kickoff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason *KickoffReason `protobuf:"varint,1,opt,name=reason,enum=am.protocol.KickoffReason" json:"reason,omitempty"`
}

func (x *SC_Kickoff) Reset() {
	*x = SC_Kickoff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_Kickoff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_Kickoff) ProtoMessage() {}

func (x *SC_Kickoff) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_Kickoff.ProtoReflect.Descriptor instead.
func (*SC_Kickoff) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{6}
}

func (x *SC_Kickoff) GetReason() KickoffReason {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return KickoffReason_ERROR
}

type CS_Gm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmdLine *string `protobuf:"bytes,1,opt,name=cmd_line,json=cmdLine" json:"cmd_line,omitempty"`
}

func (x *CS_Gm) Reset() {
	*x = CS_Gm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_Gm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_Gm) ProtoMessage() {}

func (x *CS_Gm) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_Gm.ProtoReflect.Descriptor instead.
func (*CS_Gm) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{7}
}

func (x *CS_Gm) GetCmdLine() string {
	if x != nil && x.CmdLine != nil {
		return *x.CmdLine
	}
	return ""
}

type SC_Gm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content *string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
}

func (x *SC_Gm) Reset() {
	*x = SC_Gm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_Gm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_Gm) ProtoMessage() {}

func (x *SC_Gm) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_Gm.ProtoReflect.Descriptor instead.
func (*SC_Gm) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{8}
}

func (x *SC_Gm) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x91, 0x01, 0x0a, 0x08, 0x43,
	0x53, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7a,
	0x0a, 0x08, 0x53, 0x43, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x09, 0x43, 0x53,
	0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x35, 0x0a, 0x09, 0x53, 0x43, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x03,
	0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x74, 0x52, 0x03, 0x72, 0x65, 0x74, 0x22, 0x2c, 0x0a, 0x08, 0x43, 0x53, 0x5f, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x08, 0x53, 0x43, 0x5f, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x22, 0x40, 0x0a, 0x0a, 0x53, 0x43, 0x5f, 0x4b, 0x69, 0x63, 0x6b, 0x6f, 0x66, 0x66,
	0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4b,
	0x69, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x05, 0x43, 0x53, 0x5f, 0x47, 0x6d, 0x12, 0x19, 0x0a,
	0x08, 0x63, 0x6d, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6d, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x22, 0x21, 0x0a, 0x05, 0x53, 0x43, 0x5f, 0x47,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0xc6, 0x01, 0x0a, 0x0b,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x11, 0x0a, 0x0d, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x5f, 0x49, 0x53, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x05, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54,
	0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x49, 0x53, 0x5f,
	0x46, 0x55, 0x4c, 0x4c, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x5f, 0x42, 0x41, 0x4e, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x5f, 0x52, 0x4f,
	0x4c, 0x45, 0x10, 0x09, 0x2a, 0x47, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x50,
	0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x4f, 0x43, 0x43, 0x55, 0x50, 0x59, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x49,
	0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x03, 0x2a, 0x2f, 0x0a,
	0x0d, 0x4b, 0x69, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x4d, 0x10, 0x02, 0x42, 0x11,
	0x5a, 0x0f, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6f, 0x75,
	0x74,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_login_proto_goTypes = []interface{}{
	(LoginResult)(0),   // 0: am.protocol.LoginResult
	(CreateRet)(0),     // 1: am.protocol.CreateRet
	(KickoffReason)(0), // 2: am.protocol.KickoffReason
	(*CS_Login)(nil),   // 3: am.protocol.CS_Login
	(*SC_Login)(nil),   // 4: am.protocol.SC_Login
	(*CS_Create)(nil),  // 5: am.protocol.CS_Create
	(*SC_Create)(nil),  // 6: am.protocol.SC_Create
	(*CS_Ready)(nil),   // 7: am.protocol.CS_Ready
	(*SC_Ready)(nil),   // 8: am.protocol.SC_Ready
	(*SC_Kickoff)(nil), // 9: am.protocol.SC_Kickoff
	(*CS_Gm)(nil),      // 10: am.protocol.CS_Gm
	(*SC_Gm)(nil),      // 11: am.protocol.SC_Gm
}
var file_login_proto_depIdxs = []int32{
	0, // 0: am.protocol.SC_Login.result:type_name -> am.protocol.LoginResult
	1, // 1: am.protocol.SC_Create.ret:type_name -> am.protocol.CreateRet
	2, // 2: am.protocol.SC_Kickoff.reason:type_name -> am.protocol.KickoffReason
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_Login); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_Login); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_Create); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_Create); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_Ready); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_Ready); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_Kickoff); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_Gm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_login_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_Gm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		EnumInfos:         file_login_proto_enumTypes,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
