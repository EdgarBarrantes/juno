// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: state.proto

package state

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

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code [][]byte `protobuf:"bytes,1,rep,name=code,proto3" json:"code,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{0}
}

func (x *Code) GetCode() [][]byte {
	if x != nil {
		return x.Code
	}
	return nil
}

type Storage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage map[string]string `protobuf:"bytes,1,rep,name=storage,proto3" json:"storage,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Storage) Reset() {
	*x = Storage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Storage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Storage) ProtoMessage() {}

func (x *Storage) ProtoReflect() protoreflect.Message {
	mi := &file_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Storage.ProtoReflect.Descriptor instead.
func (*Storage) Descriptor() ([]byte, []int) {
	return file_state_proto_rawDescGZIP(), []int{1}
}

func (x *Storage) GetStorage() map[string]string {
	if x != nil {
		return x.Storage
	}
	return nil
}

var File_state_proto protoreflect.FileDescriptor

var file_state_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x76, 0x0a, 0x07, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x64, 0x45, 0x74, 0x68, 0x2f, 0x6a, 0x75,
	0x6e, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_state_proto_rawDescOnce sync.Once
	file_state_proto_rawDescData = file_state_proto_rawDesc
)

func file_state_proto_rawDescGZIP() []byte {
	file_state_proto_rawDescOnce.Do(func() {
		file_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_proto_rawDescData)
	})
	return file_state_proto_rawDescData
}

var file_state_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_state_proto_goTypes = []interface{}{
	(*Code)(nil),    // 0: Code
	(*Storage)(nil), // 1: Storage
	nil,             // 2: Storage.StorageEntry
}
var file_state_proto_depIdxs = []int32{
	2, // 0: Storage.storage:type_name -> Storage.StorageEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_state_proto_init() }
func file_state_proto_init() {
	if File_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Storage); i {
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
			RawDescriptor: file_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_state_proto_goTypes,
		DependencyIndexes: file_state_proto_depIdxs,
		MessageInfos:      file_state_proto_msgTypes,
	}.Build()
	File_state_proto = out.File
	file_state_proto_rawDesc = nil
	file_state_proto_goTypes = nil
	file_state_proto_depIdxs = nil
}
