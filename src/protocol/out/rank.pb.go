// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: rank.proto

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

type CS_GetRankList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId *uint32 `protobuf:"varint,1,opt,name=rank_id,json=rankId" json:"rank_id,omitempty"`
}

func (x *CS_GetRankList) Reset() {
	*x = CS_GetRankList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_GetRankList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_GetRankList) ProtoMessage() {}

func (x *CS_GetRankList) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_GetRankList.ProtoReflect.Descriptor instead.
func (*CS_GetRankList) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{0}
}

func (x *CS_GetRankList) GetRankId() uint32 {
	if x != nil && x.RankId != nil {
		return *x.RankId
	}
	return 0
}

type SC_GetRankList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId  *uint32     `protobuf:"varint,1,opt,name=rank_id,json=rankId" json:"rank_id,omitempty"`
	List    []*RankInfo `protobuf:"bytes,2,rep,name=list" json:"list,omitempty"`
	Ranking *uint32     `protobuf:"varint,3,opt,name=ranking" json:"ranking,omitempty"` //排名  为0，表示未上榜
	Score   *uint32     `protobuf:"varint,4,opt,name=score" json:"score,omitempty"`     //排行榜积分
}

func (x *SC_GetRankList) Reset() {
	*x = SC_GetRankList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_GetRankList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_GetRankList) ProtoMessage() {}

func (x *SC_GetRankList) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_GetRankList.ProtoReflect.Descriptor instead.
func (*SC_GetRankList) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{1}
}

func (x *SC_GetRankList) GetRankId() uint32 {
	if x != nil && x.RankId != nil {
		return *x.RankId
	}
	return 0
}

func (x *SC_GetRankList) GetList() []*RankInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *SC_GetRankList) GetRanking() uint32 {
	if x != nil && x.Ranking != nil {
		return *x.Ranking
	}
	return 0
}

func (x *SC_GetRankList) GetScore() uint32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

type RankInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   *uint64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Score *uint32 `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
}

func (x *RankInfo) Reset() {
	*x = RankInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankInfo) ProtoMessage() {}

func (x *RankInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankInfo.ProtoReflect.Descriptor instead.
func (*RankInfo) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{2}
}

func (x *RankInfo) GetUid() uint64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *RankInfo) GetScore() uint32 {
	if x != nil && x.Score != nil {
		return *x.Score
	}
	return 0
}

//排行榜点赞
type CS_RankLike struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId *uint32 `protobuf:"varint,1,opt,name=rank_id,json=rankId" json:"rank_id,omitempty"`
	Rank   *uint32 `protobuf:"varint,2,opt,name=rank" json:"rank,omitempty"` //点赞的排行 （前3选一）
}

func (x *CS_RankLike) Reset() {
	*x = CS_RankLike{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_RankLike) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_RankLike) ProtoMessage() {}

func (x *CS_RankLike) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_RankLike.ProtoReflect.Descriptor instead.
func (*CS_RankLike) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{3}
}

func (x *CS_RankLike) GetRankId() uint32 {
	if x != nil && x.RankId != nil {
		return *x.RankId
	}
	return 0
}

func (x *CS_RankLike) GetRank() uint32 {
	if x != nil && x.Rank != nil {
		return *x.Rank
	}
	return 0
}

type RankUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  *uint64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`   // uid
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`  // 玩家名字
	Head *uint32 `protobuf:"varint,3,opt,name=head" json:"head,omitempty"` // 头像
}

func (x *RankUserInfo) Reset() {
	*x = RankUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankUserInfo) ProtoMessage() {}

func (x *RankUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankUserInfo.ProtoReflect.Descriptor instead.
func (*RankUserInfo) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{4}
}

func (x *RankUserInfo) GetUid() uint64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *RankUserInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *RankUserInfo) GetHead() uint32 {
	if x != nil && x.Head != nil {
		return *x.Head
	}
	return 0
}

type TowerRankInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info  *RankUserInfo `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Layer *float32      `protobuf:"fixed32,2,opt,name=layer" json:"layer,omitempty"`
	Rank  *int32        `protobuf:"varint,3,opt,name=rank" json:"rank,omitempty"`
	Like  *uint32       `protobuf:"varint,4,opt,name=like" json:"like,omitempty"`
}

func (x *TowerRankInfo) Reset() {
	*x = TowerRankInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerRankInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerRankInfo) ProtoMessage() {}

func (x *TowerRankInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerRankInfo.ProtoReflect.Descriptor instead.
func (*TowerRankInfo) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{5}
}

func (x *TowerRankInfo) GetInfo() *RankUserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *TowerRankInfo) GetLayer() float32 {
	if x != nil && x.Layer != nil {
		return *x.Layer
	}
	return 0
}

func (x *TowerRankInfo) GetRank() int32 {
	if x != nil && x.Rank != nil {
		return *x.Rank
	}
	return 0
}

func (x *TowerRankInfo) GetLike() uint32 {
	if x != nil && x.Like != nil {
		return *x.Like
	}
	return 0
}

var File_rank_proto protoreflect.FileDescriptor

var file_rank_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x29, 0x0a, 0x0e, 0x43, 0x53, 0x5f,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61,
	0x6e, 0x6b, 0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x6b, 0x49, 0x64,
	0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x61, 0x6e,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x61,
	0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x32, 0x0a, 0x08, 0x52,
	0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22,
	0x3a, 0x0a, 0x0b, 0x43, 0x53, 0x5f, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x72, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x48, 0x0a, 0x0c, 0x52,
	0x61, 0x6e, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x22, 0x7c, 0x0a, 0x0d, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x61,
	0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6c,
	0x69, 0x6b, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x6f, 0x75, 0x74,
}

var (
	file_rank_proto_rawDescOnce sync.Once
	file_rank_proto_rawDescData = file_rank_proto_rawDesc
)

func file_rank_proto_rawDescGZIP() []byte {
	file_rank_proto_rawDescOnce.Do(func() {
		file_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_rank_proto_rawDescData)
	})
	return file_rank_proto_rawDescData
}

var file_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rank_proto_goTypes = []interface{}{
	(*CS_GetRankList)(nil), // 0: am.protocol.CS_GetRankList
	(*SC_GetRankList)(nil), // 1: am.protocol.SC_GetRankList
	(*RankInfo)(nil),       // 2: am.protocol.RankInfo
	(*CS_RankLike)(nil),    // 3: am.protocol.CS_RankLike
	(*RankUserInfo)(nil),   // 4: am.protocol.RankUserInfo
	(*TowerRankInfo)(nil),  // 5: am.protocol.TowerRankInfo
}
var file_rank_proto_depIdxs = []int32{
	2, // 0: am.protocol.SC_GetRankList.list:type_name -> am.protocol.RankInfo
	4, // 1: am.protocol.TowerRankInfo.info:type_name -> am.protocol.RankUserInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rank_proto_init() }
func file_rank_proto_init() {
	if File_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_GetRankList); i {
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
		file_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_GetRankList); i {
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
		file_rank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankInfo); i {
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
		file_rank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_RankLike); i {
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
		file_rank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankUserInfo); i {
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
		file_rank_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerRankInfo); i {
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
			RawDescriptor: file_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rank_proto_goTypes,
		DependencyIndexes: file_rank_proto_depIdxs,
		MessageInfos:      file_rank_proto_msgTypes,
	}.Build()
	File_rank_proto = out.File
	file_rank_proto_rawDesc = nil
	file_rank_proto_goTypes = nil
	file_rank_proto_depIdxs = nil
}