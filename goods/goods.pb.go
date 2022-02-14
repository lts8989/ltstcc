// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: goods/goods.proto

package goods

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

type GoodsId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GoodsId) Reset() {
	*x = GoodsId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsId) ProtoMessage() {}

func (x *GoodsId) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsId.ProtoReflect.Descriptor instead.
func (*GoodsId) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsId) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Goods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GoodsName   string `protobuf:"bytes,2,opt,name=goods_name,json=goodsName,proto3" json:"goods_name,omitempty"`
	Price       int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	FrozenCount int32  `protobuf:"varint,5,opt,name=frozen_count,json=frozenCount,proto3" json:"frozen_count,omitempty"`
}

func (x *Goods) Reset() {
	*x = Goods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Goods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Goods) ProtoMessage() {}

func (x *Goods) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Goods.ProtoReflect.Descriptor instead.
func (*Goods) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{1}
}

func (x *Goods) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Goods) GetGoodsName() string {
	if x != nil {
		return x.GoodsName
	}
	return ""
}

func (x *Goods) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Goods) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Goods) GetFrozenCount() int32 {
	if x != nil {
		return x.FrozenCount
	}
	return 0
}

type GoodsFrozenParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsId int32 `protobuf:"varint,1,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	Count   int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GoodsFrozenParams) Reset() {
	*x = GoodsFrozenParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsFrozenParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsFrozenParams) ProtoMessage() {}

func (x *GoodsFrozenParams) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsFrozenParams.ProtoReflect.Descriptor instead.
func (*GoodsFrozenParams) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{2}
}

func (x *GoodsFrozenParams) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *GoodsFrozenParams) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GoodsFrozenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GoodsFrozenResp) Reset() {
	*x = GoodsFrozenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsFrozenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsFrozenResp) ProtoMessage() {}

func (x *GoodsFrozenResp) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsFrozenResp.ProtoReflect.Descriptor instead.
func (*GoodsFrozenResp) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{3}
}

func (x *GoodsFrozenResp) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type GoodsFrozenCancelParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsDetailId int32 `protobuf:"varint,1,opt,name=goods_detail_id,json=goodsDetailId,proto3" json:"goods_detail_id,omitempty"`
}

func (x *GoodsFrozenCancelParams) Reset() {
	*x = GoodsFrozenCancelParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsFrozenCancelParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsFrozenCancelParams) ProtoMessage() {}

func (x *GoodsFrozenCancelParams) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsFrozenCancelParams.ProtoReflect.Descriptor instead.
func (*GoodsFrozenCancelParams) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{4}
}

func (x *GoodsFrozenCancelParams) GetGoodsDetailId() int32 {
	if x != nil {
		return x.GoodsDetailId
	}
	return 0
}

type GoodsSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GoodsSuccess) Reset() {
	*x = GoodsSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsSuccess) ProtoMessage() {}

func (x *GoodsSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsSuccess.ProtoReflect.Descriptor instead.
func (*GoodsSuccess) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{5}
}

func (x *GoodsSuccess) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_goods_goods_proto protoreflect.FileDescriptor

var file_goods_goods_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x72, 0x6f,
	0x7a, 0x65, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27,
	0x0a, 0x0f, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x41, 0x0a, 0x17, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x32, 0xa0, 0x01, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20,
	0x0a, 0x0c, 0x67, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x08,
	0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x1a, 0x06, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x12, 0x33, 0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x12,
	0x12, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x46, 0x72, 0x6f, 0x7a, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x11, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x46, 0x72,
	0x6f, 0x7a, 0x65, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x18, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0d, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_goods_proto_rawDescOnce sync.Once
	file_goods_goods_proto_rawDescData = file_goods_goods_proto_rawDesc
)

func file_goods_goods_proto_rawDescGZIP() []byte {
	file_goods_goods_proto_rawDescOnce.Do(func() {
		file_goods_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_goods_proto_rawDescData)
	})
	return file_goods_goods_proto_rawDescData
}

var file_goods_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_goods_goods_proto_goTypes = []interface{}{
	(*GoodsId)(nil),                 // 0: GoodsId
	(*Goods)(nil),                   // 1: Goods
	(*GoodsFrozenParams)(nil),       // 2: GoodsFrozenParams
	(*GoodsFrozenResp)(nil),         // 3: GoodsFrozenResp
	(*GoodsFrozenCancelParams)(nil), // 4: GoodsFrozenCancelParams
	(*GoodsSuccess)(nil),            // 5: GoodsSuccess
}
var file_goods_goods_proto_depIdxs = []int32{
	0, // 0: GoodsInfo.getGoodsInfo:input_type -> GoodsId
	2, // 1: GoodsInfo.goodsFrozen:input_type -> GoodsFrozenParams
	4, // 2: GoodsInfo.goodsFrozenCancel:input_type -> GoodsFrozenCancelParams
	1, // 3: GoodsInfo.getGoodsInfo:output_type -> Goods
	3, // 4: GoodsInfo.goodsFrozen:output_type -> GoodsFrozenResp
	5, // 5: GoodsInfo.goodsFrozenCancel:output_type -> GoodsSuccess
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goods_goods_proto_init() }
func file_goods_goods_proto_init() {
	if File_goods_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsId); i {
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
		file_goods_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Goods); i {
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
		file_goods_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsFrozenParams); i {
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
		file_goods_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsFrozenResp); i {
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
		file_goods_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsFrozenCancelParams); i {
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
		file_goods_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsSuccess); i {
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
			RawDescriptor: file_goods_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goods_goods_proto_goTypes,
		DependencyIndexes: file_goods_goods_proto_depIdxs,
		MessageInfos:      file_goods_goods_proto_msgTypes,
	}.Build()
	File_goods_goods_proto = out.File
	file_goods_goods_proto_rawDesc = nil
	file_goods_goods_proto_goTypes = nil
	file_goods_goods_proto_depIdxs = nil
}