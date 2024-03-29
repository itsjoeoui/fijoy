// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: fijoy/v1/category.proto

package fijoyv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WorkspaceId  string          `protobuf:"bytes,2,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Name         string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CategoryType TransactionType `protobuf:"varint,4,opt,name=category_type,json=categoryType,proto3,enum=fijoy.v1.TransactionType" json:"category_type,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fijoy_v1_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_fijoy_v1_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_fijoy_v1_category_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetCategoryType() TransactionType {
	if x != nil {
		return x.CategoryType
	}
	return TransactionType_TRANSACTION_TYPE_UNSPECIFIED
}

var File_fijoy_v1_category_proto protoreflect.FileDescriptor

var file_fijoy_v1_category_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x69, 0x6a, 0x6f, 0x79,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x08, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2a, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x07, 0xba, 0x48, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x89, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x69, 0x6a, 0x6f, 0x79,
	0x2e, 0x76, 0x31, 0x42, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66,
	0x69, 0x6a, 0x6f, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x69, 0x6a, 0x6f, 0x79, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x46, 0x69, 0x6a, 0x6f, 0x79, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x08, 0x46, 0x69, 0x6a, 0x6f, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x46, 0x69,
	0x6a, 0x6f, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x09, 0x46, 0x69, 0x6a, 0x6f, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fijoy_v1_category_proto_rawDescOnce sync.Once
	file_fijoy_v1_category_proto_rawDescData = file_fijoy_v1_category_proto_rawDesc
)

func file_fijoy_v1_category_proto_rawDescGZIP() []byte {
	file_fijoy_v1_category_proto_rawDescOnce.Do(func() {
		file_fijoy_v1_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_fijoy_v1_category_proto_rawDescData)
	})
	return file_fijoy_v1_category_proto_rawDescData
}

var file_fijoy_v1_category_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fijoy_v1_category_proto_goTypes = []interface{}{
	(*Category)(nil),     // 0: fijoy.v1.Category
	(TransactionType)(0), // 1: fijoy.v1.TransactionType
}
var file_fijoy_v1_category_proto_depIdxs = []int32{
	1, // 0: fijoy.v1.Category.category_type:type_name -> fijoy.v1.TransactionType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fijoy_v1_category_proto_init() }
func file_fijoy_v1_category_proto_init() {
	if File_fijoy_v1_category_proto != nil {
		return
	}
	file_fijoy_v1_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fijoy_v1_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
			RawDescriptor: file_fijoy_v1_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fijoy_v1_category_proto_goTypes,
		DependencyIndexes: file_fijoy_v1_category_proto_depIdxs,
		MessageInfos:      file_fijoy_v1_category_proto_msgTypes,
	}.Build()
	File_fijoy_v1_category_proto = out.File
	file_fijoy_v1_category_proto_rawDesc = nil
	file_fijoy_v1_category_proto_goTypes = nil
	file_fijoy_v1_category_proto_depIdxs = nil
}
