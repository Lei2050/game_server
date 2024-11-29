// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: statistics.proto

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

type SC_StatisticsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CareerDrawNum          *uint32 `protobuf:"varint,1,opt,name=career_draw_num,json=careerDrawNum" json:"career_draw_num,omitempty"`                            // 生涯抽卡次数
	DrawOrangeNum          *uint32 `protobuf:"varint,2,opt,name=draw_orange_num,json=drawOrangeNum" json:"draw_orange_num,omitempty"`                            // 获得橙卡次数
	DrawGuaranteedSchedule *uint32 `protobuf:"varint,3,opt,name=draw_guaranteed_schedule,json=drawGuaranteedSchedule" json:"draw_guaranteed_schedule,omitempty"` //抽卡保底进度
}

func (x *SC_StatisticsList) Reset() {
	*x = SC_StatisticsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_StatisticsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_StatisticsList) ProtoMessage() {}

func (x *SC_StatisticsList) ProtoReflect() protoreflect.Message {
	mi := &file_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_StatisticsList.ProtoReflect.Descriptor instead.
func (*SC_StatisticsList) Descriptor() ([]byte, []int) {
	return file_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *SC_StatisticsList) GetCareerDrawNum() uint32 {
	if x != nil && x.CareerDrawNum != nil {
		return *x.CareerDrawNum
	}
	return 0
}

func (x *SC_StatisticsList) GetDrawOrangeNum() uint32 {
	if x != nil && x.DrawOrangeNum != nil {
		return *x.DrawOrangeNum
	}
	return 0
}

func (x *SC_StatisticsList) GetDrawGuaranteedSchedule() uint32 {
	if x != nil && x.DrawGuaranteedSchedule != nil {
		return *x.DrawGuaranteedSchedule
	}
	return 0
}

var File_statistics_proto protoreflect.FileDescriptor

var file_statistics_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22,
	0x9d, 0x01, 0x0a, 0x11, 0x53, 0x43, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x5f,
	0x64, 0x72, 0x61, 0x77, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x44, 0x72, 0x61, 0x77, 0x4e, 0x75, 0x6d, 0x12, 0x26, 0x0a,
	0x0f, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x6f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x64, 0x72, 0x61, 0x77, 0x4f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x72, 0x61, 0x77, 0x5f, 0x67, 0x75,
	0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x64, 0x72, 0x61, 0x77, 0x47, 0x75, 0x61,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x42,
	0x11, 0x5a, 0x0f, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6f,
	0x75, 0x74,
}

var (
	file_statistics_proto_rawDescOnce sync.Once
	file_statistics_proto_rawDescData = file_statistics_proto_rawDesc
)

func file_statistics_proto_rawDescGZIP() []byte {
	file_statistics_proto_rawDescOnce.Do(func() {
		file_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_statistics_proto_rawDescData)
	})
	return file_statistics_proto_rawDescData
}

var file_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_statistics_proto_goTypes = []interface{}{
	(*SC_StatisticsList)(nil), // 0: am.protocol.SC_StatisticsList
}
var file_statistics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_statistics_proto_init() }
func file_statistics_proto_init() {
	if File_statistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_StatisticsList); i {
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
			RawDescriptor: file_statistics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_statistics_proto_goTypes,
		DependencyIndexes: file_statistics_proto_depIdxs,
		MessageInfos:      file_statistics_proto_msgTypes,
	}.Build()
	File_statistics_proto = out.File
	file_statistics_proto_rawDesc = nil
	file_statistics_proto_goTypes = nil
	file_statistics_proto_depIdxs = nil
}