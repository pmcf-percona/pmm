// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: management/v1/metrics.proto

package managementv1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MetricsMode defines desired metrics mode for agent,
// it can be pull, push or auto mode chosen by server.
type MetricsMode int32

const (
	// Auto
	MetricsMode_METRICS_MODE_UNSPECIFIED MetricsMode = 0
	MetricsMode_METRICS_MODE_PULL        MetricsMode = 1
	MetricsMode_METRICS_MODE_PUSH        MetricsMode = 2
)

// Enum value maps for MetricsMode.
var (
	MetricsMode_name = map[int32]string{
		0: "METRICS_MODE_UNSPECIFIED",
		1: "METRICS_MODE_PULL",
		2: "METRICS_MODE_PUSH",
	}
	MetricsMode_value = map[string]int32{
		"METRICS_MODE_UNSPECIFIED": 0,
		"METRICS_MODE_PULL":        1,
		"METRICS_MODE_PUSH":        2,
	}
)

func (x MetricsMode) Enum() *MetricsMode {
	p := new(MetricsMode)
	*p = x
	return p
}

func (x MetricsMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricsMode) Descriptor() protoreflect.EnumDescriptor {
	return file_management_v1_metrics_proto_enumTypes[0].Descriptor()
}

func (MetricsMode) Type() protoreflect.EnumType {
	return &file_management_v1_metrics_proto_enumTypes[0]
}

func (x MetricsMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricsMode.Descriptor instead.
func (MetricsMode) EnumDescriptor() ([]byte, []int) {
	return file_management_v1_metrics_proto_rawDescGZIP(), []int{0}
}

var File_management_v1_metrics_proto protoreflect.FileDescriptor

var file_management_v1_metrics_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2a, 0x59, 0x0a, 0x0b,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x4d,
	0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x55, 0x4c, 0x4c, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x50, 0x55, 0x53, 0x48, 0x10, 0x02, 0x42, 0xad, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_management_v1_metrics_proto_rawDescOnce sync.Once
	file_management_v1_metrics_proto_rawDescData = file_management_v1_metrics_proto_rawDesc
)

func file_management_v1_metrics_proto_rawDescGZIP() []byte {
	file_management_v1_metrics_proto_rawDescOnce.Do(func() {
		file_management_v1_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_management_v1_metrics_proto_rawDescData)
	})
	return file_management_v1_metrics_proto_rawDescData
}

var (
	file_management_v1_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_management_v1_metrics_proto_goTypes   = []any{
		(MetricsMode)(0), // 0: management.v1.MetricsMode
	}
)

var file_management_v1_metrics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_management_v1_metrics_proto_init() }
func file_management_v1_metrics_proto_init() {
	if File_management_v1_metrics_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_management_v1_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_management_v1_metrics_proto_goTypes,
		DependencyIndexes: file_management_v1_metrics_proto_depIdxs,
		EnumInfos:         file_management_v1_metrics_proto_enumTypes,
	}.Build()
	File_management_v1_metrics_proto = out.File
	file_management_v1_metrics_proto_rawDesc = nil
	file_management_v1_metrics_proto_goTypes = nil
	file_management_v1_metrics_proto_depIdxs = nil
}
