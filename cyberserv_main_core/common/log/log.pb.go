// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: common/log/log.proto

package log

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Severity int32

const (
	Severity_Unknown Severity = 0
	Severity_Error   Severity = 1
	Severity_Warning Severity = 2
	Severity_Info    Severity = 3
	Severity_Debug   Severity = 4
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "Unknown",
		1: "Error",
		2: "Warning",
		3: "Info",
		4: "Debug",
	}
	Severity_value = map[string]int32{
		"Unknown": 0,
		"Error":   1,
		"Warning": 2,
		"Info":    3,
		"Debug":   4,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_common_log_log_proto_enumTypes[0].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_common_log_log_proto_enumTypes[0]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_common_log_log_proto_rawDescGZIP(), []int{0}
}

var File_common_log_log_proto protoreflect.FileDescriptor

var file_common_log_log_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6c, 0x6f, 0x67, 0x2a, 0x44, 0x0a,
	0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x10, 0x04, 0x42, 0x50, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6c, 0x6f, 0x67,
	0x50, 0x01, 0x5a, 0x19, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0xaa, 0x02, 0x15,
	0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_log_log_proto_rawDescOnce sync.Once
	file_common_log_log_proto_rawDescData = file_common_log_log_proto_rawDesc
)

func file_common_log_log_proto_rawDescGZIP() []byte {
	file_common_log_log_proto_rawDescOnce.Do(func() {
		file_common_log_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_log_log_proto_rawDescData)
	})
	return file_common_log_log_proto_rawDescData
}

var file_common_log_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_log_log_proto_goTypes = []interface{}{
	(Severity)(0), // 0: cyberservices.core.common.log.Severity
}
var file_common_log_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_log_log_proto_init() }
func file_common_log_log_proto_init() {
	if File_common_log_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_log_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_log_log_proto_goTypes,
		DependencyIndexes: file_common_log_log_proto_depIdxs,
		EnumInfos:         file_common_log_log_proto_enumTypes,
	}.Build()
	File_common_log_log_proto = out.File
	file_common_log_log_proto_rawDesc = nil
	file_common_log_log_proto_goTypes = nil
	file_common_log_log_proto_depIdxs = nil
}
