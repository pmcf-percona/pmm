// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3-devel
// 	protoc        (unknown)
// source: management/v1/annotation.proto

package managementv1

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AddAnnotationRequest is a params to add new annotation.
type AddAnnotationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// An annotation description. Required.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Tags are used to filter annotations.
	Tags []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	// Used for annotating a node.
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Used for annotating services.
	ServiceNames  []string `protobuf:"bytes,4,rep,name=service_names,json=serviceNames,proto3" json:"service_names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddAnnotationRequest) Reset() {
	*x = AddAnnotationRequest{}
	mi := &file_management_v1_annotation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddAnnotationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAnnotationRequest) ProtoMessage() {}

func (x *AddAnnotationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_annotation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAnnotationRequest.ProtoReflect.Descriptor instead.
func (*AddAnnotationRequest) Descriptor() ([]byte, []int) {
	return file_management_v1_annotation_proto_rawDescGZIP(), []int{0}
}

func (x *AddAnnotationRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AddAnnotationRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AddAnnotationRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AddAnnotationRequest) GetServiceNames() []string {
	if x != nil {
		return x.ServiceNames
	}
	return nil
}

type AddAnnotationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddAnnotationResponse) Reset() {
	*x = AddAnnotationResponse{}
	mi := &file_management_v1_annotation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddAnnotationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAnnotationResponse) ProtoMessage() {}

func (x *AddAnnotationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_management_v1_annotation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAnnotationResponse.ProtoReflect.Descriptor instead.
func (*AddAnnotationResponse) Descriptor() ([]byte, []int) {
	return file_management_v1_annotation_proto_rawDescGZIP(), []int{1}
}

var File_management_v1_annotation_proto protoreflect.FileDescriptor

var file_management_v1_annotation_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x14, 0x41, 0x64, 0x64,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb0, 0x01,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x42, 0x0f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_management_v1_annotation_proto_rawDescOnce sync.Once
	file_management_v1_annotation_proto_rawDescData = file_management_v1_annotation_proto_rawDesc
)

func file_management_v1_annotation_proto_rawDescGZIP() []byte {
	file_management_v1_annotation_proto_rawDescOnce.Do(func() {
		file_management_v1_annotation_proto_rawDescData = string(protoimpl.X.CompressGZIP([]byte(file_management_v1_annotation_proto_rawDescData)))
	})
	return []byte(file_management_v1_annotation_proto_rawDescData)
}

var (
	file_management_v1_annotation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_management_v1_annotation_proto_goTypes  = []any{
		(*AddAnnotationRequest)(nil),  // 0: management.v1.AddAnnotationRequest
		(*AddAnnotationResponse)(nil), // 1: management.v1.AddAnnotationResponse
	}
)

var file_management_v1_annotation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_management_v1_annotation_proto_init() }
func file_management_v1_annotation_proto_init() {
	if File_management_v1_annotation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_management_v1_annotation_proto_rawDesc), len(file_management_v1_annotation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_management_v1_annotation_proto_goTypes,
		DependencyIndexes: file_management_v1_annotation_proto_depIdxs,
		MessageInfos:      file_management_v1_annotation_proto_msgTypes,
	}.Build()
	File_management_v1_annotation_proto = out.File
	file_management_v1_annotation_proto_goTypes = nil
	file_management_v1_annotation_proto_depIdxs = nil
}
