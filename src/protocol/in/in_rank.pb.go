// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: in_rank.proto

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

// 排行榜插入
type RankInsert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId  *uint32 `protobuf:"varint,1,opt,name=rank_id,json=rankId" json:"rank_id,omitempty"`
	RankKey *string `protobuf:"bytes,2,opt,name=rank_key,json=rankKey" json:"rank_key,omitempty"` // 排行榜key 由排行榜id和zone_id组成
	Element *string `protobuf:"bytes,3,opt,name=element" json:"element,omitempty"`                // 插入的元素
}

func (x *RankInsert) Reset() {
	*x = RankInsert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankInsert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankInsert) ProtoMessage() {}

func (x *RankInsert) ProtoReflect() protoreflect.Message {
	mi := &file_in_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankInsert.ProtoReflect.Descriptor instead.
func (*RankInsert) Descriptor() ([]byte, []int) {
	return file_in_rank_proto_rawDescGZIP(), []int{0}
}

func (x *RankInsert) GetRankId() uint32 {
	if x != nil && x.RankId != nil {
		return *x.RankId
	}
	return 0
}

func (x *RankInsert) GetRankKey() string {
	if x != nil && x.RankKey != nil {
		return *x.RankKey
	}
	return ""
}

func (x *RankInsert) GetElement() string {
	if x != nil && x.Element != nil {
		return *x.Element
	}
	return ""
}

//排行榜同步
type RankSync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranks []*RankSyncInfo `protobuf:"bytes,1,rep,name=ranks" json:"ranks,omitempty"`
	Load  *int32          `protobuf:"varint,2,opt,name=load" json:"load,omitempty"` //是否为load
}

func (x *RankSync) Reset() {
	*x = RankSync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankSync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankSync) ProtoMessage() {}

func (x *RankSync) ProtoReflect() protoreflect.Message {
	mi := &file_in_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankSync.ProtoReflect.Descriptor instead.
func (*RankSync) Descriptor() ([]byte, []int) {
	return file_in_rank_proto_rawDescGZIP(), []int{1}
}

func (x *RankSync) GetRanks() []*RankSyncInfo {
	if x != nil {
		return x.Ranks
	}
	return nil
}

func (x *RankSync) GetLoad() int32 {
	if x != nil && x.Load != nil {
		return *x.Load
	}
	return 0
}

type RankSyncInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId   *uint32  `protobuf:"varint,1,opt,name=rank_id,json=rankId" json:"rank_id,omitempty"`
	RankKey  *string  `protobuf:"bytes,2,opt,name=rank_key,json=rankKey" json:"rank_key,omitempty"`
	RankList []string `protobuf:"bytes,3,rep,name=rank_list,json=rankList" json:"rank_list,omitempty"`
	Done     *bool    `protobuf:"varint,4,opt,name=done" json:"done,omitempty"` //是否全部同步完毕
}

func (x *RankSyncInfo) Reset() {
	*x = RankSyncInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_in_rank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankSyncInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankSyncInfo) ProtoMessage() {}

func (x *RankSyncInfo) ProtoReflect() protoreflect.Message {
	mi := &file_in_rank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankSyncInfo.ProtoReflect.Descriptor instead.
func (*RankSyncInfo) Descriptor() ([]byte, []int) {
	return file_in_rank_proto_rawDescGZIP(), []int{2}
}

func (x *RankSyncInfo) GetRankId() uint32 {
	if x != nil && x.RankId != nil {
		return *x.RankId
	}
	return 0
}

func (x *RankSyncInfo) GetRankKey() string {
	if x != nil && x.RankKey != nil {
		return *x.RankKey
	}
	return ""
}

func (x *RankSyncInfo) GetRankList() []string {
	if x != nil {
		return x.RankList
	}
	return nil
}

func (x *RankSyncInfo) GetDone() bool {
	if x != nil && x.Done != nil {
		return *x.Done
	}
	return false
}

var File_in_rank_proto protoreflect.FileDescriptor

var file_in_rank_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x5a, 0x0a, 0x0a,
	0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61,
	0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x6e,
	0x6b, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x4b, 0x65, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4f, 0x0a, 0x08, 0x52, 0x61, 0x6e, 0x6b,
	0x53, 0x79, 0x6e, 0x63, 0x12, 0x2f, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x72, 0x61, 0x6e, 0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x73, 0x0a, 0x0c, 0x52, 0x61, 0x6e,
	0x6b, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x6e,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x6b,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x6e, 0x6b, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f,
	0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x42, 0x10,
	0x5a, 0x0e, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x69, 0x6e,
}

var (
	file_in_rank_proto_rawDescOnce sync.Once
	file_in_rank_proto_rawDescData = file_in_rank_proto_rawDesc
)

func file_in_rank_proto_rawDescGZIP() []byte {
	file_in_rank_proto_rawDescOnce.Do(func() {
		file_in_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_in_rank_proto_rawDescData)
	})
	return file_in_rank_proto_rawDescData
}

var file_in_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_in_rank_proto_goTypes = []interface{}{
	(*RankInsert)(nil),   // 0: am.protocol.RankInsert
	(*RankSync)(nil),     // 1: am.protocol.RankSync
	(*RankSyncInfo)(nil), // 2: am.protocol.RankSyncInfo
}
var file_in_rank_proto_depIdxs = []int32{
	2, // 0: am.protocol.RankSync.ranks:type_name -> am.protocol.RankSyncInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_in_rank_proto_init() }
func file_in_rank_proto_init() {
	if File_in_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_in_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankInsert); i {
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
		file_in_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankSync); i {
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
		file_in_rank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankSyncInfo); i {
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
			RawDescriptor: file_in_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_in_rank_proto_goTypes,
		DependencyIndexes: file_in_rank_proto_depIdxs,
		MessageInfos:      file_in_rank_proto_msgTypes,
	}.Build()
	File_in_rank_proto = out.File
	file_in_rank_proto_rawDesc = nil
	file_in_rank_proto_goTypes = nil
	file_in_rank_proto_depIdxs = nil
}
