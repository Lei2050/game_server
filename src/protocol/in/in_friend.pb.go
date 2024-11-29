// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: in_friend.proto

package in

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

//好友查找
type FriendSearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (x *FriendSearchReq) Reset() {
	*x = FriendSearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendSearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendSearchReq) ProtoMessage() {}

func (x *FriendSearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_in_friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendSearchReq.ProtoReflect.Descriptor instead.
func (*FriendSearchReq) Descriptor() ([]byte, []int) {
	return file_in_friend_proto_rawDescGZIP(), []int{0}
}

func (x *FriendSearchReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type FriendSearchRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId *uint64 `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
}

func (x *FriendSearchRes) Reset() {
	*x = FriendSearchRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_friend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendSearchRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendSearchRes) ProtoMessage() {}

func (x *FriendSearchRes) ProtoReflect() protoreflect.Message {
	mi := &file_in_friend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendSearchRes.ProtoReflect.Descriptor instead.
func (*FriendSearchRes) Descriptor() ([]byte, []int) {
	return file_in_friend_proto_rawDescGZIP(), []int{1}
}

func (x *FriendSearchRes) GetPlayerId() uint64 {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return 0
}

//好友检查
type FriendCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId *uint64 `protobuf:"varint,1,opt,name=player_id,json=playerId" json:"player_id,omitempty"` //添加的人
	AddUid   *uint64 `protobuf:"varint,2,opt,name=add_uid,json=addUid" json:"add_uid,omitempty"`       //被添加的人
	AddType  *uint32 `protobuf:"varint,3,opt,name=add_type,json=addType" json:"add_type,omitempty"`    //添加类型 好友请求 或者 好友
}

func (x *FriendCheck) Reset() {
	*x = FriendCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_friend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendCheck) ProtoMessage() {}

func (x *FriendCheck) ProtoReflect() protoreflect.Message {
	mi := &file_in_friend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendCheck.ProtoReflect.Descriptor instead.
func (*FriendCheck) Descriptor() ([]byte, []int) {
	return file_in_friend_proto_rawDescGZIP(), []int{2}
}

func (x *FriendCheck) GetPlayerId() uint64 {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return 0
}

func (x *FriendCheck) GetAddUid() uint64 {
	if x != nil && x.AddUid != nil {
		return *x.AddUid
	}
	return 0
}

func (x *FriendCheck) GetAddType() uint32 {
	if x != nil && x.AddType != nil {
		return *x.AddType
	}
	return 0
}

type FriendCheckRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddUid   *uint64 `protobuf:"varint,1,opt,name=add_uid,json=addUid" json:"add_uid,omitempty"`       //被添加的人
	PlayerId *uint64 `protobuf:"varint,2,opt,name=player_id,json=playerId" json:"player_id,omitempty"` //添加的人
	Ret      *uint32 `protobuf:"varint,3,opt,name=ret" json:"ret,omitempty"`                           //检查结果 为 0表示检查通过
	AddType  *uint32 `protobuf:"varint,4,opt,name=add_type,json=addType" json:"add_type,omitempty"`    //添加类型
}

func (x *FriendCheckRes) Reset() {
	*x = FriendCheckRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_friend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendCheckRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendCheckRes) ProtoMessage() {}

func (x *FriendCheckRes) ProtoReflect() protoreflect.Message {
	mi := &file_in_friend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendCheckRes.ProtoReflect.Descriptor instead.
func (*FriendCheckRes) Descriptor() ([]byte, []int) {
	return file_in_friend_proto_rawDescGZIP(), []int{3}
}

func (x *FriendCheckRes) GetAddUid() uint64 {
	if x != nil && x.AddUid != nil {
		return *x.AddUid
	}
	return 0
}

func (x *FriendCheckRes) GetPlayerId() uint64 {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return 0
}

func (x *FriendCheckRes) GetRet() uint32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *FriendCheckRes) GetAddType() uint32 {
	if x != nil && x.AddType != nil {
		return *x.AddType
	}
	return 0
}

//好友推荐
type G2R_FriendRecommend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      *uint64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`          //获取好友推荐的玩家
	Language *string `protobuf:"bytes,2,opt,name=language" json:"language,omitempty"` //设备语言
	Arena    *string `protobuf:"bytes,3,opt,name=arena" json:"arena,omitempty"`       //下载地区相同
	Level    *uint32 `protobuf:"varint,4,opt,name=level" json:"level,omitempty"`      //玩家等级
}

func (x *G2R_FriendRecommend) Reset() {
	*x = G2R_FriendRecommend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_friend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *G2R_FriendRecommend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*G2R_FriendRecommend) ProtoMessage() {}

func (x *G2R_FriendRecommend) ProtoReflect() protoreflect.Message {
	mi := &file_in_friend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use G2R_FriendRecommend.ProtoReflect.Descriptor instead.
func (*G2R_FriendRecommend) Descriptor() ([]byte, []int) {
	return file_in_friend_proto_rawDescGZIP(), []int{4}
}

func (x *G2R_FriendRecommend) GetUid() uint64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *G2R_FriendRecommend) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

func (x *G2R_FriendRecommend) GetArena() string {
	if x != nil && x.Arena != nil {
		return *x.Arena
	}
	return ""
}

func (x *G2R_FriendRecommend) GetLevel() uint32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

type R2G_FriendRecommend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     *uint64                `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"` //获取好友推荐的玩家
	Friends []*FriendRecommendInfo `protobuf:"bytes,2,rep,name=friends" json:"friends,omitempty"`
}

func (x *R2G_FriendRecommend) Reset() {
	*x = R2G_FriendRecommend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_friend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *R2G_FriendRecommend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*R2G_FriendRecommend) ProtoMessage() {}

func (x *R2G_FriendRecommend) ProtoReflect() protoreflect.Message {
	mi := &file_in_friend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use R2G_FriendRecommend.ProtoReflect.Descriptor instead.
func (*R2G_FriendRecommend) Descriptor() ([]byte, []int) {
	return file_in_friend_proto_rawDescGZIP(), []int{5}
}

func (x *R2G_FriendRecommend) GetUid() uint64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *R2G_FriendRecommend) GetFriends() []*FriendRecommendInfo {
	if x != nil {
		return x.Friends
	}
	return nil
}

type FriendRecommendInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid           *uint64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Online        *uint32 `protobuf:"varint,2,opt,name=online" json:"online,omitempty"`                                    //是否在线 1 在线 0 不在线
	Language      *uint32 `protobuf:"varint,3,opt,name=language" json:"language,omitempty"`                                //语言是否相同 1 相同 0 不同
	Arena         *uint32 `protobuf:"varint,4,opt,name=arena" json:"arena,omitempty"`                                      //下载地区是否相同
	LevelInterval *uint32 `protobuf:"varint,5,opt,name=level_interval,json=levelInterval" json:"level_interval,omitempty"` //等级差
	OfflineTime   *int64  `protobuf:"varint,6,opt,name=offline_time,json=offlineTime" json:"offline_time,omitempty"`       //离线时间
}

func (x *FriendRecommendInfo) Reset() {
	*x = FriendRecommendInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_friend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRecommendInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRecommendInfo) ProtoMessage() {}

func (x *FriendRecommendInfo) ProtoReflect() protoreflect.Message {
	mi := &file_in_friend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRecommendInfo.ProtoReflect.Descriptor instead.
func (*FriendRecommendInfo) Descriptor() ([]byte, []int) {
	return file_in_friend_proto_rawDescGZIP(), []int{6}
}

func (x *FriendRecommendInfo) GetUid() uint64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *FriendRecommendInfo) GetOnline() uint32 {
	if x != nil && x.Online != nil {
		return *x.Online
	}
	return 0
}

func (x *FriendRecommendInfo) GetLanguage() uint32 {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return 0
}

func (x *FriendRecommendInfo) GetArena() uint32 {
	if x != nil && x.Arena != nil {
		return *x.Arena
	}
	return 0
}

func (x *FriendRecommendInfo) GetLevelInterval() uint32 {
	if x != nil && x.LevelInterval != nil {
		return *x.LevelInterval
	}
	return 0
}

func (x *FriendRecommendInfo) GetOfflineTime() int64 {
	if x != nil && x.OfflineTime != nil {
		return *x.OfflineTime
	}
	return 0
}

var File_in_friend_proto protoreflect.FileDescriptor

var file_in_friend_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x25,
	0x0a, 0x0f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x0f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x0b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x61, 0x64, 0x64, 0x55, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x73, 0x0a, 0x0e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x5f, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x64, 0x64, 0x55, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6f, 0x0a, 0x13, 0x47, 0x32,
	0x52, 0x5f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x72, 0x65, 0x6e, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x63, 0x0a, 0x13, 0x52,
	0x32, 0x47, 0x5f, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x22, 0xbb, 0x01, 0x0a, 0x13, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61,
	0x72, 0x65, 0x6e, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x10,
	0x5a, 0x0e, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x69, 0x6e,
}

var (
	file_in_friend_proto_rawDescOnce sync.Once
	file_in_friend_proto_rawDescData = file_in_friend_proto_rawDesc
)

func file_in_friend_proto_rawDescGZIP() []byte {
	file_in_friend_proto_rawDescOnce.Do(func() {
		file_in_friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_in_friend_proto_rawDescData)
	})
	return file_in_friend_proto_rawDescData
}

var file_in_friend_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_in_friend_proto_goTypes = []interface{}{
	(*FriendSearchReq)(nil),     // 0: am.protocol.FriendSearchReq
	(*FriendSearchRes)(nil),     // 1: am.protocol.FriendSearchRes
	(*FriendCheck)(nil),         // 2: am.protocol.FriendCheck
	(*FriendCheckRes)(nil),      // 3: am.protocol.FriendCheckRes
	(*G2R_FriendRecommend)(nil), // 4: am.protocol.G2R_FriendRecommend
	(*R2G_FriendRecommend)(nil), // 5: am.protocol.R2G_FriendRecommend
	(*FriendRecommendInfo)(nil), // 6: am.protocol.FriendRecommendInfo
}
var file_in_friend_proto_depIdxs = []int32{
	6, // 0: am.protocol.R2G_FriendRecommend.friends:type_name -> am.protocol.FriendRecommendInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_in_friend_proto_init() }
func file_in_friend_proto_init() {
	if File_in_friend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_in_friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendSearchReq); i {
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
		file_in_friend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendSearchRes); i {
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
		file_in_friend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendCheck); i {
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
		file_in_friend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendCheckRes); i {
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
		file_in_friend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*G2R_FriendRecommend); i {
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
		file_in_friend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*R2G_FriendRecommend); i {
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
		file_in_friend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRecommendInfo); i {
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
			RawDescriptor: file_in_friend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_in_friend_proto_goTypes,
		DependencyIndexes: file_in_friend_proto_depIdxs,
		MessageInfos:      file_in_friend_proto_msgTypes,
	}.Build()
	File_in_friend_proto = out.File
	file_in_friend_proto_rawDesc = nil
	file_in_friend_proto_goTypes = nil
	file_in_friend_proto_depIdxs = nil
}