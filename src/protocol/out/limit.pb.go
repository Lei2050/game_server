// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: limit.proto

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

type SC_LimitList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsedItem    []*UsedItem    `protobuf:"bytes,1,rep,name=used_item,json=usedItem" json:"used_item,omitempty"`
	ShopBuy     []*ShopBuy     `protobuf:"bytes,2,rep,name=shop_buy,json=shopBuy" json:"shop_buy,omitempty"`
	ShopRefresh []*ShopRefresh `protobuf:"bytes,3,rep,name=shop_refresh,json=shopRefresh" json:"shop_refresh,omitempty"`
}

func (x *SC_LimitList) Reset() {
	*x = SC_LimitList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_LimitList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_LimitList) ProtoMessage() {}

func (x *SC_LimitList) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_LimitList.ProtoReflect.Descriptor instead.
func (*SC_LimitList) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{0}
}

func (x *SC_LimitList) GetUsedItem() []*UsedItem {
	if x != nil {
		return x.UsedItem
	}
	return nil
}

func (x *SC_LimitList) GetShopBuy() []*ShopBuy {
	if x != nil {
		return x.ShopBuy
	}
	return nil
}

func (x *SC_LimitList) GetShopRefresh() []*ShopRefresh {
	if x != nil {
		return x.ShopRefresh
	}
	return nil
}

type UsedItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId *uint64 `protobuf:"varint,1,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	Num    *uint32 `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
}

func (x *UsedItem) Reset() {
	*x = UsedItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsedItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsedItem) ProtoMessage() {}

func (x *UsedItem) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsedItem.ProtoReflect.Descriptor instead.
func (*UsedItem) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{1}
}

func (x *UsedItem) GetItemId() uint64 {
	if x != nil && x.ItemId != nil {
		return *x.ItemId
	}
	return 0
}

func (x *UsedItem) GetNum() uint32 {
	if x != nil && x.Num != nil {
		return *x.Num
	}
	return 0
}

type ShopBuy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId *uint64        `protobuf:"varint,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
	Goods  []*ShopBuyGood `protobuf:"bytes,2,rep,name=goods" json:"goods,omitempty"`
}

func (x *ShopBuy) Reset() {
	*x = ShopBuy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopBuy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopBuy) ProtoMessage() {}

func (x *ShopBuy) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopBuy.ProtoReflect.Descriptor instead.
func (*ShopBuy) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{2}
}

func (x *ShopBuy) GetShopId() uint64 {
	if x != nil && x.ShopId != nil {
		return *x.ShopId
	}
	return 0
}

func (x *ShopBuy) GetGoods() []*ShopBuyGood {
	if x != nil {
		return x.Goods
	}
	return nil
}

type ShopBuyGood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodId *uint64 `protobuf:"varint,1,opt,name=good_id,json=goodId" json:"good_id,omitempty"`
	Num    *uint32 `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
}

func (x *ShopBuyGood) Reset() {
	*x = ShopBuyGood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopBuyGood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopBuyGood) ProtoMessage() {}

func (x *ShopBuyGood) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopBuyGood.ProtoReflect.Descriptor instead.
func (*ShopBuyGood) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{3}
}

func (x *ShopBuyGood) GetGoodId() uint64 {
	if x != nil && x.GoodId != nil {
		return *x.GoodId
	}
	return 0
}

func (x *ShopBuyGood) GetNum() uint32 {
	if x != nil && x.Num != nil {
		return *x.Num
	}
	return 0
}

type ShopRefresh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId *uint64 `protobuf:"varint,1,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
	Num    *uint32 `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
}

func (x *ShopRefresh) Reset() {
	*x = ShopRefresh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopRefresh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopRefresh) ProtoMessage() {}

func (x *ShopRefresh) ProtoReflect() protoreflect.Message {
	mi := &file_limit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopRefresh.ProtoReflect.Descriptor instead.
func (*ShopRefresh) Descriptor() ([]byte, []int) {
	return file_limit_proto_rawDescGZIP(), []int{4}
}

func (x *ShopRefresh) GetShopId() uint64 {
	if x != nil && x.ShopId != nil {
		return *x.ShopId
	}
	return 0
}

func (x *ShopRefresh) GetNum() uint32 {
	if x != nil && x.Num != nil {
		return *x.Num
	}
	return 0
}

var File_limit_proto protoreflect.FileDescriptor

var file_limit_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0xb0, 0x01, 0x0a, 0x0c, 0x53,
	0x43, 0x5f, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x73, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x75, 0x73, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x2f, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x62, 0x75, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x53, 0x68, 0x6f, 0x70, 0x42, 0x75, 0x79, 0x52, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x42, 0x75, 0x79,
	0x12, 0x3b, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x52, 0x0b, 0x73, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x22, 0x35, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6e, 0x75, 0x6d, 0x22, 0x52, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x75, 0x79, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f,
	0x64, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x38, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x70,
	0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x6f, 0x6f, 0x64, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x22, 0x38, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x42, 0x11, 0x5a, 0x0f,
	0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6f, 0x75, 0x74,
}

var (
	file_limit_proto_rawDescOnce sync.Once
	file_limit_proto_rawDescData = file_limit_proto_rawDesc
)

func file_limit_proto_rawDescGZIP() []byte {
	file_limit_proto_rawDescOnce.Do(func() {
		file_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_limit_proto_rawDescData)
	})
	return file_limit_proto_rawDescData
}

var file_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_limit_proto_goTypes = []interface{}{
	(*SC_LimitList)(nil), // 0: am.protocol.SC_LimitList
	(*UsedItem)(nil),     // 1: am.protocol.UsedItem
	(*ShopBuy)(nil),      // 2: am.protocol.ShopBuy
	(*ShopBuyGood)(nil),  // 3: am.protocol.ShopBuyGood
	(*ShopRefresh)(nil),  // 4: am.protocol.ShopRefresh
}
var file_limit_proto_depIdxs = []int32{
	1, // 0: am.protocol.SC_LimitList.used_item:type_name -> am.protocol.UsedItem
	2, // 1: am.protocol.SC_LimitList.shop_buy:type_name -> am.protocol.ShopBuy
	4, // 2: am.protocol.SC_LimitList.shop_refresh:type_name -> am.protocol.ShopRefresh
	3, // 3: am.protocol.ShopBuy.goods:type_name -> am.protocol.ShopBuyGood
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_limit_proto_init() }
func file_limit_proto_init() {
	if File_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_LimitList); i {
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
		file_limit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsedItem); i {
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
		file_limit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopBuy); i {
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
		file_limit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopBuyGood); i {
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
		file_limit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShopRefresh); i {
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
			RawDescriptor: file_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_limit_proto_goTypes,
		DependencyIndexes: file_limit_proto_depIdxs,
		MessageInfos:      file_limit_proto_msgTypes,
	}.Build()
	File_limit_proto = out.File
	file_limit_proto_rawDesc = nil
	file_limit_proto_goTypes = nil
	file_limit_proto_depIdxs = nil
}