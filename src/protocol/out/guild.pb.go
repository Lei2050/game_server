// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: guild.proto

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

//获取公会信息
type CS_GetGuildList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Placeholder *string `protobuf:"bytes,1,opt,name=placeholder" json:"placeholder,omitempty"`
}

func (x *CS_GetGuildList) Reset() {
	*x = CS_GetGuildList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guild_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_GetGuildList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_GetGuildList) ProtoMessage() {}

func (x *CS_GetGuildList) ProtoReflect() protoreflect.Message {
	mi := &file_guild_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_GetGuildList.ProtoReflect.Descriptor instead.
func (*CS_GetGuildList) Descriptor() ([]byte, []int) {
	return file_guild_proto_rawDescGZIP(), []int{0}
}

func (x *CS_GetGuildList) GetPlaceholder() string {
	if x != nil && x.Placeholder != nil {
		return *x.Placeholder
	}
	return ""
}

type SC_GetGuildList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guilds  []*GuildInfo `protobuf:"bytes,1,rep,name=guilds" json:"guilds,omitempty"`
	Applier []uint64     `protobuf:"varint,2,rep,name=applier" json:"applier,omitempty"` //申请人（leader 可看）
}

func (x *SC_GetGuildList) Reset() {
	*x = SC_GetGuildList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guild_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_GetGuildList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_GetGuildList) ProtoMessage() {}

func (x *SC_GetGuildList) ProtoReflect() protoreflect.Message {
	mi := &file_guild_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_GetGuildList.ProtoReflect.Descriptor instead.
func (*SC_GetGuildList) Descriptor() ([]byte, []int) {
	return file_guild_proto_rawDescGZIP(), []int{1}
}

func (x *SC_GetGuildList) GetGuilds() []*GuildInfo {
	if x != nil {
		return x.Guilds
	}
	return nil
}

func (x *SC_GetGuildList) GetApplier() []uint64 {
	if x != nil {
		return x.Applier
	}
	return nil
}

type GuildInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Members *uint32 `protobuf:"varint,3,opt,name=members" json:"members,omitempty"` //成员数量
	Message *string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`  //公会公告
}

func (x *GuildInfo) Reset() {
	*x = GuildInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guild_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuildInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuildInfo) ProtoMessage() {}

func (x *GuildInfo) ProtoReflect() protoreflect.Message {
	mi := &file_guild_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuildInfo.ProtoReflect.Descriptor instead.
func (*GuildInfo) Descriptor() ([]byte, []int) {
	return file_guild_proto_rawDescGZIP(), []int{2}
}

func (x *GuildInfo) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *GuildInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GuildInfo) GetMembers() uint32 {
	if x != nil && x.Members != nil {
		return *x.Members
	}
	return 0
}

func (x *GuildInfo) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

//创建公会
type CS_GuildCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"` //公会名字
}

func (x *CS_GuildCreate) Reset() {
	*x = CS_GuildCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guild_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_GuildCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_GuildCreate) ProtoMessage() {}

func (x *CS_GuildCreate) ProtoReflect() protoreflect.Message {
	mi := &file_guild_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_GuildCreate.ProtoReflect.Descriptor instead.
func (*CS_GuildCreate) Descriptor() ([]byte, []int) {
	return file_guild_proto_rawDescGZIP(), []int{3}
}

func (x *CS_GuildCreate) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

//解散公会
type CS_GuildDismiss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Placeholder *int32 `protobuf:"varint,1,opt,name=placeholder" json:"placeholder,omitempty"`
}

func (x *CS_GuildDismiss) Reset() {
	*x = CS_GuildDismiss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guild_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_GuildDismiss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_GuildDismiss) ProtoMessage() {}

func (x *CS_GuildDismiss) ProtoReflect() protoreflect.Message {
	mi := &file_guild_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_GuildDismiss.ProtoReflect.Descriptor instead.
func (*CS_GuildDismiss) Descriptor() ([]byte, []int) {
	return file_guild_proto_rawDescGZIP(), []int{4}
}

func (x *CS_GuildDismiss) GetPlaceholder() int32 {
	if x != nil && x.Placeholder != nil {
		return *x.Placeholder
	}
	return 0
}

//申请加入公会
type CS_GuildApplyJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuildId *uint64 `protobuf:"varint,1,opt,name=guild_id,json=guildId" json:"guild_id,omitempty"`
}

func (x *CS_GuildApplyJoin) Reset() {
	*x = CS_GuildApplyJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guild_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_GuildApplyJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_GuildApplyJoin) ProtoMessage() {}

func (x *CS_GuildApplyJoin) ProtoReflect() protoreflect.Message {
	mi := &file_guild_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_GuildApplyJoin.ProtoReflect.Descriptor instead.
func (*CS_GuildApplyJoin) Descriptor() ([]byte, []int) {
	return file_guild_proto_rawDescGZIP(), []int{5}
}

func (x *CS_GuildApplyJoin) GetGuildId() uint64 {
	if x != nil && x.GuildId != nil {
		return *x.GuildId
	}
	return 0
}

//申请审核
type CS_GuildApplyCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId *uint64 `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	Accept   *uint32 `protobuf:"varint,2,opt,name=accept" json:"accept,omitempty"` // 1 同意 0 拒绝
}

func (x *CS_GuildApplyCheck) Reset() {
	*x = CS_GuildApplyCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guild_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_GuildApplyCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_GuildApplyCheck) ProtoMessage() {}

func (x *CS_GuildApplyCheck) ProtoReflect() protoreflect.Message {
	mi := &file_guild_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_GuildApplyCheck.ProtoReflect.Descriptor instead.
func (*CS_GuildApplyCheck) Descriptor() ([]byte, []int) {
	return file_guild_proto_rawDescGZIP(), []int{6}
}

func (x *CS_GuildApplyCheck) GetPlayerId() uint64 {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return 0
}

func (x *CS_GuildApplyCheck) GetAccept() uint32 {
	if x != nil && x.Accept != nil {
		return *x.Accept
	}
	return 0
}

//退出公会
type CS_GuildLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Placeholder *int32 `protobuf:"varint,1,opt,name=placeholder" json:"placeholder,omitempty"`
}

func (x *CS_GuildLeave) Reset() {
	*x = CS_GuildLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guild_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_GuildLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_GuildLeave) ProtoMessage() {}

func (x *CS_GuildLeave) ProtoReflect() protoreflect.Message {
	mi := &file_guild_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_GuildLeave.ProtoReflect.Descriptor instead.
func (*CS_GuildLeave) Descriptor() ([]byte, []int) {
	return file_guild_proto_rawDescGZIP(), []int{7}
}

func (x *CS_GuildLeave) GetPlaceholder() int32 {
	if x != nil && x.Placeholder != nil {
		return *x.Placeholder
	}
	return 0
}

//踢出公会
type CS_GuildKick struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId *uint64 `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
}

func (x *CS_GuildKick) Reset() {
	*x = CS_GuildKick{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guild_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_GuildKick) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_GuildKick) ProtoMessage() {}

func (x *CS_GuildKick) ProtoReflect() protoreflect.Message {
	mi := &file_guild_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_GuildKick.ProtoReflect.Descriptor instead.
func (*CS_GuildKick) Descriptor() ([]byte, []int) {
	return file_guild_proto_rawDescGZIP(), []int{8}
}

func (x *CS_GuildKick) GetPlayerId() uint64 {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return 0
}

//修改公告
type CS_GuildMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (x *CS_GuildMessage) Reset() {
	*x = CS_GuildMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guild_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_GuildMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_GuildMessage) ProtoMessage() {}

func (x *CS_GuildMessage) ProtoReflect() protoreflect.Message {
	mi := &file_guild_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_GuildMessage.ProtoReflect.Descriptor instead.
func (*CS_GuildMessage) Descriptor() ([]byte, []int) {
	return file_guild_proto_rawDescGZIP(), []int{9}
}

func (x *CS_GuildMessage) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_guild_proto protoreflect.FileDescriptor

var file_guild_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x33, 0x0a, 0x0f, 0x43, 0x53,
	0x5f, 0x47, 0x65, 0x74, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22,
	0x5b, 0x0a, 0x0f, 0x53, 0x43, 0x5f, 0x47, 0x65, 0x74, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x67, 0x75, 0x69, 0x6c,
	0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x22, 0x63, 0x0a, 0x09,
	0x47, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x53, 0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x0f, 0x43, 0x53, 0x5f, 0x47, 0x75,
	0x69, 0x6c, 0x64, 0x44, 0x69, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x11,
	0x43, 0x53, 0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4a, 0x6f, 0x69,
	0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x12,
	0x43, 0x53, 0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x22, 0x31, 0x0a, 0x0d, 0x43, 0x53, 0x5f, 0x47, 0x75,
	0x69, 0x6c, 0x64, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0c, 0x43, 0x53,
	0x5f, 0x47, 0x75, 0x69, 0x6c, 0x64, 0x4b, 0x69, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x43, 0x53, 0x5f, 0x47, 0x75,
	0x69, 0x6c, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x6f, 0x75, 0x74,
}

var (
	file_guild_proto_rawDescOnce sync.Once
	file_guild_proto_rawDescData = file_guild_proto_rawDesc
)

func file_guild_proto_rawDescGZIP() []byte {
	file_guild_proto_rawDescOnce.Do(func() {
		file_guild_proto_rawDescData = protoimpl.X.CompressGZIP(file_guild_proto_rawDescData)
	})
	return file_guild_proto_rawDescData
}

var file_guild_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_guild_proto_goTypes = []interface{}{
	(*CS_GetGuildList)(nil),    // 0: am.protocol.CS_GetGuildList
	(*SC_GetGuildList)(nil),    // 1: am.protocol.SC_GetGuildList
	(*GuildInfo)(nil),          // 2: am.protocol.GuildInfo
	(*CS_GuildCreate)(nil),     // 3: am.protocol.CS_GuildCreate
	(*CS_GuildDismiss)(nil),    // 4: am.protocol.CS_GuildDismiss
	(*CS_GuildApplyJoin)(nil),  // 5: am.protocol.CS_GuildApplyJoin
	(*CS_GuildApplyCheck)(nil), // 6: am.protocol.CS_GuildApplyCheck
	(*CS_GuildLeave)(nil),      // 7: am.protocol.CS_GuildLeave
	(*CS_GuildKick)(nil),       // 8: am.protocol.CS_GuildKick
	(*CS_GuildMessage)(nil),    // 9: am.protocol.CS_GuildMessage
}
var file_guild_proto_depIdxs = []int32{
	2, // 0: am.protocol.SC_GetGuildList.guilds:type_name -> am.protocol.GuildInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_guild_proto_init() }
func file_guild_proto_init() {
	if File_guild_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_guild_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_GetGuildList); i {
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
		file_guild_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_GetGuildList); i {
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
		file_guild_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuildInfo); i {
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
		file_guild_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_GuildCreate); i {
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
		file_guild_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_GuildDismiss); i {
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
		file_guild_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_GuildApplyJoin); i {
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
		file_guild_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_GuildApplyCheck); i {
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
		file_guild_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_GuildLeave); i {
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
		file_guild_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_GuildKick); i {
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
		file_guild_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_GuildMessage); i {
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
			RawDescriptor: file_guild_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_guild_proto_goTypes,
		DependencyIndexes: file_guild_proto_depIdxs,
		MessageInfos:      file_guild_proto_msgTypes,
	}.Build()
	File_guild_proto = out.File
	file_guild_proto_rawDesc = nil
	file_guild_proto_goTypes = nil
	file_guild_proto_depIdxs = nil
}
