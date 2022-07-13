// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: PiCalculus.proto

package PiCalculus

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

type CalculusData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Steps  uint64 `protobuf:"varint,2,opt,name=steps,proto3" json:"steps,omitempty"`
	Result int64  `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalculusData) Reset() {
	*x = CalculusData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PiCalculus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculusData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculusData) ProtoMessage() {}

func (x *CalculusData) ProtoReflect() protoreflect.Message {
	mi := &file_PiCalculus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculusData.ProtoReflect.Descriptor instead.
func (*CalculusData) Descriptor() ([]byte, []int) {
	return file_PiCalculus_proto_rawDescGZIP(), []int{0}
}

func (x *CalculusData) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CalculusData) GetSteps() uint64 {
	if x != nil {
		return x.Steps
	}
	return 0
}

func (x *CalculusData) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_PiCalculus_proto protoreflect.FileDescriptor

var file_PiCalculus_proto_rawDesc = []byte{
	0x0a, 0x10, 0x50, 0x69, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x75, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x65, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x79, 0x0a, 0x0a, 0x50, 0x69, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x50, 0x69, 0x4c, 0x65, 0x69, 0x62, 0x6e, 0x69, 0x7a, 0x12, 0x0d, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x75, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x13,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x69, 0x6c, 0x61, 0x6b, 0x61, 0x6e,
	0x74, 0x68, 0x61, 0x12, 0x0d, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x75, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x75, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PiCalculus_proto_rawDescOnce sync.Once
	file_PiCalculus_proto_rawDescData = file_PiCalculus_proto_rawDesc
)

func file_PiCalculus_proto_rawDescGZIP() []byte {
	file_PiCalculus_proto_rawDescOnce.Do(func() {
		file_PiCalculus_proto_rawDescData = protoimpl.X.CompressGZIP(file_PiCalculus_proto_rawDescData)
	})
	return file_PiCalculus_proto_rawDescData
}

var file_PiCalculus_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PiCalculus_proto_goTypes = []interface{}{
	(*CalculusData)(nil), // 0: CalculusData
}
var file_PiCalculus_proto_depIdxs = []int32{
	0, // 0: PiCalculus.CalculatePiLeibniz:input_type -> CalculusData
	0, // 1: PiCalculus.CalculateNilakantha:input_type -> CalculusData
	0, // 2: PiCalculus.CalculatePiLeibniz:output_type -> CalculusData
	0, // 3: PiCalculus.CalculateNilakantha:output_type -> CalculusData
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PiCalculus_proto_init() }
func file_PiCalculus_proto_init() {
	if File_PiCalculus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PiCalculus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculusData); i {
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
			RawDescriptor: file_PiCalculus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_PiCalculus_proto_goTypes,
		DependencyIndexes: file_PiCalculus_proto_depIdxs,
		MessageInfos:      file_PiCalculus_proto_msgTypes,
	}.Build()
	File_PiCalculus_proto = out.File
	file_PiCalculus_proto_rawDesc = nil
	file_PiCalculus_proto_goTypes = nil
	file_PiCalculus_proto_depIdxs = nil
}
