// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: users/users.proto

package users

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

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Balance      int32  `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	FrozenAmount int32  `protobuf:"varint,4,opt,name=frozen_amount,json=frozenAmount,proto3" json:"frozen_amount,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{0}
}

func (x *Users) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Users) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Users) GetBalance() int32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Users) GetFrozenAmount() int32 {
	if x != nil {
		return x.FrozenAmount
	}
	return 0
}

type UsersId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UsersId) Reset() {
	*x = UsersId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersId) ProtoMessage() {}

func (x *UsersId) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersId.ProtoReflect.Descriptor instead.
func (*UsersId) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{1}
}

func (x *UsersId) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type BatchCreateUsersParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *BatchCreateUsersParams) Reset() {
	*x = BatchCreateUsersParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreateUsersParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateUsersParams) ProtoMessage() {}

func (x *BatchCreateUsersParams) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateUsersParams.ProtoReflect.Descriptor instead.
func (*BatchCreateUsersParams) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{2}
}

func (x *BatchCreateUsersParams) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type BatchCreateUsersCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *BatchCreateUsersCount) Reset() {
	*x = BatchCreateUsersCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreateUsersCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateUsersCount) ProtoMessage() {}

func (x *BatchCreateUsersCount) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateUsersCount.ProtoReflect.Descriptor instead.
func (*BatchCreateUsersCount) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{3}
}

func (x *BatchCreateUsersCount) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type RechargeParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RechargeMoney int32 `protobuf:"varint,2,opt,name=recharge_money,json=rechargeMoney,proto3" json:"recharge_money,omitempty"`
}

func (x *RechargeParams) Reset() {
	*x = RechargeParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RechargeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RechargeParams) ProtoMessage() {}

func (x *RechargeParams) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RechargeParams.ProtoReflect.Descriptor instead.
func (*RechargeParams) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{4}
}

func (x *RechargeParams) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RechargeParams) GetRechargeMoney() int32 {
	if x != nil {
		return x.RechargeMoney
	}
	return 0
}

type UsersSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UsersSuccess) Reset() {
	*x = UsersSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersSuccess) ProtoMessage() {}

func (x *UsersSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersSuccess.ProtoReflect.Descriptor instead.
func (*UsersSuccess) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{5}
}

func (x *UsersSuccess) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type FrozenTryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FrozenAmount int32 `protobuf:"varint,2,opt,name=frozen_amount,json=frozenAmount,proto3" json:"frozen_amount,omitempty"`
}

func (x *FrozenTryParams) Reset() {
	*x = FrozenTryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrozenTryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrozenTryParams) ProtoMessage() {}

func (x *FrozenTryParams) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrozenTryParams.ProtoReflect.Descriptor instead.
func (*FrozenTryParams) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{6}
}

func (x *FrozenTryParams) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FrozenTryParams) GetFrozenAmount() int32 {
	if x != nil {
		return x.FrozenAmount
	}
	return 0
}

type FrozenTryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersDetailId    int32 `protobuf:"varint,1,opt,name=users_detail_id,json=usersDetailId,proto3" json:"users_detail_id,omitempty"`
	UserBalance      int32 `protobuf:"varint,2,opt,name=user_balance,json=userBalance,proto3" json:"user_balance,omitempty"`
	UserFrozenAmount int32 `protobuf:"varint,3,opt,name=user_frozen_amount,json=userFrozenAmount,proto3" json:"user_frozen_amount,omitempty"`
}

func (x *FrozenTryResp) Reset() {
	*x = FrozenTryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrozenTryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrozenTryResp) ProtoMessage() {}

func (x *FrozenTryResp) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrozenTryResp.ProtoReflect.Descriptor instead.
func (*FrozenTryResp) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{7}
}

func (x *FrozenTryResp) GetUsersDetailId() int32 {
	if x != nil {
		return x.UsersDetailId
	}
	return 0
}

func (x *FrozenTryResp) GetUserBalance() int32 {
	if x != nil {
		return x.UserBalance
	}
	return 0
}

func (x *FrozenTryResp) GetUserFrozenAmount() int32 {
	if x != nil {
		return x.UserFrozenAmount
	}
	return 0
}

type FrozenCancelParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersDetailId int32 `protobuf:"varint,1,opt,name=users_detail_id,json=usersDetailId,proto3" json:"users_detail_id,omitempty"`
}

func (x *FrozenCancelParams) Reset() {
	*x = FrozenCancelParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrozenCancelParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrozenCancelParams) ProtoMessage() {}

func (x *FrozenCancelParams) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrozenCancelParams.ProtoReflect.Descriptor instead.
func (*FrozenCancelParams) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{8}
}

func (x *FrozenCancelParams) GetUsersDetailId() int32 {
	if x != nil {
		return x.UsersDetailId
	}
	return 0
}

type FrozenCancelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FrozenCancelResp) Reset() {
	*x = FrozenCancelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrozenCancelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrozenCancelResp) ProtoMessage() {}

func (x *FrozenCancelResp) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrozenCancelResp.ProtoReflect.Descriptor instead.
func (*FrozenCancelResp) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{9}
}

func (x *FrozenCancelResp) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type FrozenCommitParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersDetailId int32 `protobuf:"varint,1,opt,name=users_detail_id,json=usersDetailId,proto3" json:"users_detail_id,omitempty"`
}

func (x *FrozenCommitParams) Reset() {
	*x = FrozenCommitParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrozenCommitParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrozenCommitParams) ProtoMessage() {}

func (x *FrozenCommitParams) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrozenCommitParams.ProtoReflect.Descriptor instead.
func (*FrozenCommitParams) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{10}
}

func (x *FrozenCommitParams) GetUsersDetailId() int32 {
	if x != nil {
		return x.UsersDetailId
	}
	return 0
}

type FrozenCommitResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersDetailId int32 `protobuf:"varint,1,opt,name=users_detail_id,json=usersDetailId,proto3" json:"users_detail_id,omitempty"`
}

func (x *FrozenCommitResp) Reset() {
	*x = FrozenCommitResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_users_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrozenCommitResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrozenCommitResp) ProtoMessage() {}

func (x *FrozenCommitResp) ProtoReflect() protoreflect.Message {
	mi := &file_users_users_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrozenCommitResp.ProtoReflect.Descriptor instead.
func (*FrozenCommitResp) Descriptor() ([]byte, []int) {
	return file_users_users_proto_rawDescGZIP(), []int{11}
}

func (x *FrozenCommitResp) GetUsersDetailId() int32 {
	if x != nil {
		return x.UsersDetailId
	}
	return 0
}

var File_users_users_proto protoreflect.FileDescriptor

var file_users_users_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65,
	0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1f, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2e, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x15, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0x24, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x4f, 0x0a, 0x0f, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x54, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x88, 0x01, 0x0a, 0x0d, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x54, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x46,
	0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x12, 0x46,
	0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x46, 0x72, 0x6f,
	0x7a, 0x65, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x3c, 0x0a, 0x12, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49,
	0x64, 0x22, 0x3a, 0x0a, 0x10, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x32, 0xce, 0x02,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x06, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x1a, 0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x10,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x17, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x0f, 0x2e,
	0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0d,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a,
	0x0f, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x79,
	0x12, 0x10, 0x2e, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x54, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x0e, 0x2e, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x54, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x3c, 0x0a, 0x12, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x13, 0x2e, 0x46, 0x72, 0x6f, 0x7a, 0x65,
	0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x11, 0x2e,
	0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x3c, 0x0a, 0x12, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x13, 0x2e, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x11, 0x2e, 0x46, 0x72,
	0x6f, 0x7a, 0x65, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_users_users_proto_rawDescOnce sync.Once
	file_users_users_proto_rawDescData = file_users_users_proto_rawDesc
)

func file_users_users_proto_rawDescGZIP() []byte {
	file_users_users_proto_rawDescOnce.Do(func() {
		file_users_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_users_proto_rawDescData)
	})
	return file_users_users_proto_rawDescData
}

var file_users_users_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_users_users_proto_goTypes = []interface{}{
	(*Users)(nil),                  // 0: Users
	(*UsersId)(nil),                // 1: UsersId
	(*BatchCreateUsersParams)(nil), // 2: BatchCreateUsersParams
	(*BatchCreateUsersCount)(nil),  // 3: batchCreateUsersCount
	(*RechargeParams)(nil),         // 4: RechargeParams
	(*UsersSuccess)(nil),           // 5: UsersSuccess
	(*FrozenTryParams)(nil),        // 6: FrozenTryParams
	(*FrozenTryResp)(nil),          // 7: FrozenTryResp
	(*FrozenCancelParams)(nil),     // 8: FrozenCancelParams
	(*FrozenCancelResp)(nil),       // 9: FrozenCancelResp
	(*FrozenCommitParams)(nil),     // 10: FrozenCommitParams
	(*FrozenCommitResp)(nil),       // 11: FrozenCommitResp
}
var file_users_users_proto_depIdxs = []int32{
	0,  // 0: UsersInfo.createUsers:input_type -> Users
	2,  // 1: UsersInfo.batchCreateUsers:input_type -> BatchCreateUsersParams
	4,  // 2: UsersInfo.recharge:input_type -> RechargeParams
	6,  // 3: UsersInfo.frozenAmountTry:input_type -> FrozenTryParams
	8,  // 4: UsersInfo.frozenAmountCancel:input_type -> FrozenCancelParams
	10, // 5: UsersInfo.frozenAmountCommit:input_type -> FrozenCommitParams
	1,  // 6: UsersInfo.createUsers:output_type -> UsersId
	3,  // 7: UsersInfo.batchCreateUsers:output_type -> batchCreateUsersCount
	5,  // 8: UsersInfo.recharge:output_type -> UsersSuccess
	7,  // 9: UsersInfo.frozenAmountTry:output_type -> FrozenTryResp
	9,  // 10: UsersInfo.frozenAmountCancel:output_type -> FrozenCancelResp
	11, // 11: UsersInfo.frozenAmountCommit:output_type -> FrozenCommitResp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_users_users_proto_init() }
func file_users_users_proto_init() {
	if File_users_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_users_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersId); i {
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
		file_users_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchCreateUsersParams); i {
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
		file_users_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchCreateUsersCount); i {
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
		file_users_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RechargeParams); i {
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
		file_users_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersSuccess); i {
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
		file_users_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrozenTryParams); i {
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
		file_users_users_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrozenTryResp); i {
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
		file_users_users_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrozenCancelParams); i {
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
		file_users_users_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrozenCancelResp); i {
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
		file_users_users_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrozenCommitParams); i {
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
		file_users_users_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrozenCommitResp); i {
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
			RawDescriptor: file_users_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_users_proto_goTypes,
		DependencyIndexes: file_users_users_proto_depIdxs,
		MessageInfos:      file_users_users_proto_msgTypes,
	}.Build()
	File_users_users_proto = out.File
	file_users_users_proto_rawDesc = nil
	file_users_users_proto_goTypes = nil
	file_users_users_proto_depIdxs = nil
}