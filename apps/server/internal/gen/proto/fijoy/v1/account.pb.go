// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: fijoy/v1/account.proto

package fijoyv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccountType int32

const (
	AccountType_CASH        AccountType = 0
	AccountType_DEBT        AccountType = 1
	AccountType_INVESTMENT  AccountType = 2
	AccountType_OTHER_ASSET AccountType = 3
)

// Enum value maps for AccountType.
var (
	AccountType_name = map[int32]string{
		0: "CASH",
		1: "DEBT",
		2: "INVESTMENT",
		3: "OTHER_ASSET",
	}
	AccountType_value = map[string]int32{
		"CASH":        0,
		"DEBT":        1,
		"INVESTMENT":  2,
		"OTHER_ASSET": 3,
	}
)

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}

func (x AccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_fijoy_v1_account_proto_enumTypes[0].Descriptor()
}

func (AccountType) Type() protoreflect.EnumType {
	return &file_fijoy_v1_account_proto_enumTypes[0]
}

func (x AccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountType.Descriptor instead.
func (AccountType) EnumDescriptor() ([]byte, []int) {
	return file_fijoy_v1_account_proto_rawDescGZIP(), []int{0}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WorkspaceId string                 `protobuf:"bytes,2,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AccountType AccountType            `protobuf:"varint,4,opt,name=account_type,json=accountType,proto3,enum=fijoy.v1.AccountType" json:"account_type,omitempty"`
	Balance     *Money                 `protobuf:"bytes,5,opt,name=balance,proto3" json:"balance,omitempty"`
	Currency    string                 `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Institution string                 `protobuf:"bytes,7,opt,name=institution,proto3" json:"institution,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fijoy_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_fijoy_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_fijoy_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetAccountType() AccountType {
	if x != nil {
		return x.AccountType
	}
	return AccountType_CASH
}

func (x *Account) GetBalance() *Money {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *Account) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Account) GetInstitution() string {
	if x != nil {
		return x.Institution
	}
	return ""
}

func (x *Account) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Account) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Accounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *Accounts) Reset() {
	*x = Accounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fijoy_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accounts) ProtoMessage() {}

func (x *Accounts) ProtoReflect() protoreflect.Message {
	mi := &file_fijoy_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accounts.ProtoReflect.Descriptor instead.
func (*Accounts) Descriptor() ([]byte, []int) {
	return file_fijoy_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *Accounts) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AccountType AccountType `protobuf:"varint,2,opt,name=account_type,json=accountType,proto3,enum=fijoy.v1.AccountType" json:"account_type,omitempty"`
	Balance     *Money      `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
	Currency    string      `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	Institution string      `protobuf:"bytes,5,opt,name=institution,proto3" json:"institution,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fijoy_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fijoy_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_fijoy_v1_account_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAccountRequest) GetAccountType() AccountType {
	if x != nil {
		return x.AccountType
	}
	return AccountType_CASH
}

func (x *CreateAccountRequest) GetBalance() *Money {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *CreateAccountRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CreateAccountRequest) GetInstitution() string {
	if x != nil {
		return x.Institution
	}
	return ""
}

type GetAccountByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAccountByIdRequest) Reset() {
	*x = GetAccountByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fijoy_v1_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountByIdRequest) ProtoMessage() {}

func (x *GetAccountByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fijoy_v1_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountByIdRequest.ProtoReflect.Descriptor instead.
func (*GetAccountByIdRequest) Descriptor() ([]byte, []int) {
	return file_fijoy_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAccountByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAccountByIdRequest) Reset() {
	*x = DeleteAccountByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fijoy_v1_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountByIdRequest) ProtoMessage() {}

func (x *DeleteAccountByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fijoy_v1_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccountByIdRequest) Descriptor() ([]byte, []int) {
	return file_fijoy_v1_account_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAccountByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_fijoy_v1_account_proto protoreflect.FileDescriptor

var file_fijoy_v1_account_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x96, 0x03, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x38, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66,
	0x69, 0x6a, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x29, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x39, 0x0a, 0x08,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x69, 0x6a,
	0x6f, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0xe8, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x29, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x42, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x53, 0x48,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x42, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x49, 0x4e, 0x56, 0x45, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b,
	0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x10, 0x03, 0x32, 0xad, 0x02,
	0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x66, 0x69,
	0x6a, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22,
	0x03, 0x90, 0x02, 0x01, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12,
	0x4c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x88, 0x01,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29,
	0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2f, 0x76,
	0x31, 0x3b, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa,
	0x02, 0x08, 0x46, 0x69, 0x6a, 0x6f, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x46, 0x69, 0x6a,
	0x6f, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x46, 0x69, 0x6a, 0x6f, 0x79, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x46,
	0x69, 0x6a, 0x6f, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fijoy_v1_account_proto_rawDescOnce sync.Once
	file_fijoy_v1_account_proto_rawDescData = file_fijoy_v1_account_proto_rawDesc
)

func file_fijoy_v1_account_proto_rawDescGZIP() []byte {
	file_fijoy_v1_account_proto_rawDescOnce.Do(func() {
		file_fijoy_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_fijoy_v1_account_proto_rawDescData)
	})
	return file_fijoy_v1_account_proto_rawDescData
}

var file_fijoy_v1_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fijoy_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_fijoy_v1_account_proto_goTypes = []interface{}{
	(AccountType)(0),                 // 0: fijoy.v1.AccountType
	(*Account)(nil),                  // 1: fijoy.v1.Account
	(*Accounts)(nil),                 // 2: fijoy.v1.Accounts
	(*CreateAccountRequest)(nil),     // 3: fijoy.v1.CreateAccountRequest
	(*GetAccountByIdRequest)(nil),    // 4: fijoy.v1.GetAccountByIdRequest
	(*DeleteAccountByIdRequest)(nil), // 5: fijoy.v1.DeleteAccountByIdRequest
	(*Money)(nil),                    // 6: fijoy.v1.Money
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_fijoy_v1_account_proto_depIdxs = []int32{
	0,  // 0: fijoy.v1.Account.account_type:type_name -> fijoy.v1.AccountType
	6,  // 1: fijoy.v1.Account.balance:type_name -> fijoy.v1.Money
	7,  // 2: fijoy.v1.Account.created_at:type_name -> google.protobuf.Timestamp
	7,  // 3: fijoy.v1.Account.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 4: fijoy.v1.Accounts.accounts:type_name -> fijoy.v1.Account
	0,  // 5: fijoy.v1.CreateAccountRequest.account_type:type_name -> fijoy.v1.AccountType
	6,  // 6: fijoy.v1.CreateAccountRequest.balance:type_name -> fijoy.v1.Money
	3,  // 7: fijoy.v1.AccountService.CreateAccount:input_type -> fijoy.v1.CreateAccountRequest
	8,  // 8: fijoy.v1.AccountService.GetAccounts:input_type -> google.protobuf.Empty
	4,  // 9: fijoy.v1.AccountService.GetAccountById:input_type -> fijoy.v1.GetAccountByIdRequest
	5,  // 10: fijoy.v1.AccountService.DeleteAccountById:input_type -> fijoy.v1.DeleteAccountByIdRequest
	1,  // 11: fijoy.v1.AccountService.CreateAccount:output_type -> fijoy.v1.Account
	2,  // 12: fijoy.v1.AccountService.GetAccounts:output_type -> fijoy.v1.Accounts
	1,  // 13: fijoy.v1.AccountService.GetAccountById:output_type -> fijoy.v1.Account
	1,  // 14: fijoy.v1.AccountService.DeleteAccountById:output_type -> fijoy.v1.Account
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_fijoy_v1_account_proto_init() }
func file_fijoy_v1_account_proto_init() {
	if File_fijoy_v1_account_proto != nil {
		return
	}
	file_fijoy_v1_money_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fijoy_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_fijoy_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accounts); i {
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
		file_fijoy_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_fijoy_v1_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountByIdRequest); i {
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
		file_fijoy_v1_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountByIdRequest); i {
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
			RawDescriptor: file_fijoy_v1_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fijoy_v1_account_proto_goTypes,
		DependencyIndexes: file_fijoy_v1_account_proto_depIdxs,
		EnumInfos:         file_fijoy_v1_account_proto_enumTypes,
		MessageInfos:      file_fijoy_v1_account_proto_msgTypes,
	}.Build()
	File_fijoy_v1_account_proto = out.File
	file_fijoy_v1_account_proto_rawDesc = nil
	file_fijoy_v1_account_proto_goTypes = nil
	file_fijoy_v1_account_proto_depIdxs = nil
}
