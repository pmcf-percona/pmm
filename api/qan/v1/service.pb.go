// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3-devel
// 	protoc        (unknown)
// source: qan/v1/service.proto

package qanv1

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MetricsNamesRequest is empty.
type GetMetricsNamesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMetricsNamesRequest) Reset() {
	*x = GetMetricsNamesRequest{}
	mi := &file_qan_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetricsNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsNamesRequest) ProtoMessage() {}

func (x *GetMetricsNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qan_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsNamesRequest.ProtoReflect.Descriptor instead.
func (*GetMetricsNamesRequest) Descriptor() ([]byte, []int) {
	return file_qan_v1_service_proto_rawDescGZIP(), []int{0}
}

// MetricsNamesReply is map of stored metrics:
// key is root of metric name in db (Ex:. [m_]query_time[_sum]);
// value - Human readable name of metrics.
type GetMetricsNamesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          map[string]string      `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMetricsNamesResponse) Reset() {
	*x = GetMetricsNamesResponse{}
	mi := &file_qan_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetricsNamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsNamesResponse) ProtoMessage() {}

func (x *GetMetricsNamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qan_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsNamesResponse.ProtoReflect.Descriptor instead.
func (*GetMetricsNamesResponse) Descriptor() ([]byte, []int) {
	return file_qan_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetMetricsNamesResponse) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_qan_v1_service_proto protoreflect.FileDescriptor

var file_qan_v1_service_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x71, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x71, 0x61,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x71, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x71, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x91, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x71, 0x61, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x32, 0x94, 0x10, 0x0a, 0x0a, 0x51, 0x41, 0x4e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xb8, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x18, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x71, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x76, 0x92, 0x41, 0x4f, 0x12, 0x0a, 0x47, 0x65, 0x74, 0x20,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x41, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20,
	0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x69, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x20, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a,
	0x01, 0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x3a, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0xcc, 0x01,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x71, 0x61, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x92, 0x41, 0x38, 0x12,
	0x0b, 0x47, 0x65, 0x74, 0x20, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x29, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x61, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x20, 0x6d, 0x61, 0x70, 0x20, 0x6f, 0x66, 0x20, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a,
	0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x3a, 0x67, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0xb3, 0x01, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x1e, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x5f, 0x92, 0x41, 0x39, 0x12, 0x11, 0x47, 0x65, 0x74, 0x20, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x24, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x73, 0x20, 0x61, 0x20, 0x6d, 0x61, 0x70, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x6c, 0x6c,
	0x20, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x61,
	0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x3a, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0xa5, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x19, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x71,
	0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x92, 0x41, 0x40, 0x12, 0x0b, 0x47,
	0x65, 0x74, 0x20, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x31, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x73, 0x20, 0x61, 0x20, 0x6d, 0x61, 0x70, 0x20, 0x6f, 0x66, 0x20, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x3a,
	0x67, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x9c, 0x01, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x18, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x92,
	0x41, 0x3b, 0x12, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x2d,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x3a,
	0x67, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0xae, 0x01, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x71, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x63, 0x92, 0x41, 0x41, 0x12, 0x0d, 0x47, 0x65, 0x74, 0x20,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x1a, 0x30, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x73, 0x20, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x20, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x3a, 0x67, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0xee, 0x01, 0x0a, 0x1b, 0x45,
	0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x42, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x12, 0x2a, 0x2e, 0x71, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x46, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x42, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x76, 0x92, 0x41, 0x4e, 0x12, 0x17, 0x47, 0x65, 0x74, 0x20, 0x45, 0x78,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x20, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x1a, 0x33, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x65,
	0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x20, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x20, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x20, 0x49, 0x44, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22,
	0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x3a, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e,
	0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0xbd, 0x01, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1b, 0x2e, 0x71,
	0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x71, 0x61, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x72, 0x92, 0x41, 0x4b, 0x12, 0x0e, 0x47, 0x65,
	0x74, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x20, 0x50, 0x6c, 0x61, 0x6e, 0x1a, 0x39, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x61, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x70,
	0x6c, 0x61, 0x6e, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x70, 0x6c, 0x61, 0x6e, 0x20, 0x69, 0x64, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f,
	0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x7b, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x12, 0xa8, 0x01, 0x0a, 0x0b,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x71, 0x61,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x92, 0x41, 0x3e, 0x12, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x20, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x1a, 0x25, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x20, 0x69, 0x66, 0x20, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x63, 0x6c, 0x69, 0x63,
	0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a,
	0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0xbd, 0x01, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x42, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x71, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x79, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x71, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x79, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x92, 0x41, 0x44,
	0x12, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x36, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x20, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x49, 0x44, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x44, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f,
	0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x67, 0x65, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0xb1, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x71, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x71, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x92, 0x41, 0x37,
	0x12, 0x11, 0x47, 0x65, 0x74, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x20, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x1a, 0x22, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x73, 0x20, 0x61, 0x20,
	0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a,
	0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x61, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a,
	0x67, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x7c, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x2e, 0x71, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x61, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x71, 0x61, 0x6e, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x51, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x51, 0x61, 0x6e, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x06, 0x51, 0x61, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x51, 0x61, 0x6e, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x07, 0x51, 0x61, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_qan_v1_service_proto_rawDescOnce sync.Once
	file_qan_v1_service_proto_rawDescData = file_qan_v1_service_proto_rawDesc
)

func file_qan_v1_service_proto_rawDescGZIP() []byte {
	file_qan_v1_service_proto_rawDescOnce.Do(func() {
		file_qan_v1_service_proto_rawDescData = string(protoimpl.X.CompressGZIP([]byte(file_qan_v1_service_proto_rawDescData)))
	})
	return []byte(file_qan_v1_service_proto_rawDescData)
}

var (
	file_qan_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_qan_v1_service_proto_goTypes  = []any{
		(*GetMetricsNamesRequest)(nil),              // 0: qan.v1.GetMetricsNamesRequest
		(*GetMetricsNamesResponse)(nil),             // 1: qan.v1.GetMetricsNamesResponse
		nil,                                         // 2: qan.v1.GetMetricsNamesResponse.DataEntry
		(*GetReportRequest)(nil),                    // 3: qan.v1.GetReportRequest
		(*GetFilteredMetricsNamesRequest)(nil),      // 4: qan.v1.GetFilteredMetricsNamesRequest
		(*GetMetricsRequest)(nil),                   // 5: qan.v1.GetMetricsRequest
		(*GetLabelsRequest)(nil),                    // 6: qan.v1.GetLabelsRequest
		(*GetHistogramRequest)(nil),                 // 7: qan.v1.GetHistogramRequest
		(*ExplainFingerprintByQueryIDRequest)(nil),  // 8: qan.v1.ExplainFingerprintByQueryIDRequest
		(*GetQueryPlanRequest)(nil),                 // 9: qan.v1.GetQueryPlanRequest
		(*QueryExistsRequest)(nil),                  // 10: qan.v1.QueryExistsRequest
		(*SchemaByQueryIDRequest)(nil),              // 11: qan.v1.SchemaByQueryIDRequest
		(*GetQueryExampleRequest)(nil),              // 12: qan.v1.GetQueryExampleRequest
		(*GetReportResponse)(nil),                   // 13: qan.v1.GetReportResponse
		(*GetFilteredMetricsNamesResponse)(nil),     // 14: qan.v1.GetFilteredMetricsNamesResponse
		(*GetMetricsResponse)(nil),                  // 15: qan.v1.GetMetricsResponse
		(*GetLabelsResponse)(nil),                   // 16: qan.v1.GetLabelsResponse
		(*GetHistogramResponse)(nil),                // 17: qan.v1.GetHistogramResponse
		(*ExplainFingerprintByQueryIDResponse)(nil), // 18: qan.v1.ExplainFingerprintByQueryIDResponse
		(*GetQueryPlanResponse)(nil),                // 19: qan.v1.GetQueryPlanResponse
		(*QueryExistsResponse)(nil),                 // 20: qan.v1.QueryExistsResponse
		(*SchemaByQueryIDResponse)(nil),             // 21: qan.v1.SchemaByQueryIDResponse
		(*GetQueryExampleResponse)(nil),             // 22: qan.v1.GetQueryExampleResponse
	}
)

var file_qan_v1_service_proto_depIdxs = []int32{
	2,  // 0: qan.v1.GetMetricsNamesResponse.data:type_name -> qan.v1.GetMetricsNamesResponse.DataEntry
	3,  // 1: qan.v1.QANService.GetReport:input_type -> qan.v1.GetReportRequest
	4,  // 2: qan.v1.QANService.GetFilteredMetricsNames:input_type -> qan.v1.GetFilteredMetricsNamesRequest
	0,  // 3: qan.v1.QANService.GetMetricsNames:input_type -> qan.v1.GetMetricsNamesRequest
	5,  // 4: qan.v1.QANService.GetMetrics:input_type -> qan.v1.GetMetricsRequest
	6,  // 5: qan.v1.QANService.GetLabels:input_type -> qan.v1.GetLabelsRequest
	7,  // 6: qan.v1.QANService.GetHistogram:input_type -> qan.v1.GetHistogramRequest
	8,  // 7: qan.v1.QANService.ExplainFingerprintByQueryID:input_type -> qan.v1.ExplainFingerprintByQueryIDRequest
	9,  // 8: qan.v1.QANService.GetQueryPlan:input_type -> qan.v1.GetQueryPlanRequest
	10, // 9: qan.v1.QANService.QueryExists:input_type -> qan.v1.QueryExistsRequest
	11, // 10: qan.v1.QANService.SchemaByQueryID:input_type -> qan.v1.SchemaByQueryIDRequest
	12, // 11: qan.v1.QANService.GetQueryExample:input_type -> qan.v1.GetQueryExampleRequest
	13, // 12: qan.v1.QANService.GetReport:output_type -> qan.v1.GetReportResponse
	14, // 13: qan.v1.QANService.GetFilteredMetricsNames:output_type -> qan.v1.GetFilteredMetricsNamesResponse
	1,  // 14: qan.v1.QANService.GetMetricsNames:output_type -> qan.v1.GetMetricsNamesResponse
	15, // 15: qan.v1.QANService.GetMetrics:output_type -> qan.v1.GetMetricsResponse
	16, // 16: qan.v1.QANService.GetLabels:output_type -> qan.v1.GetLabelsResponse
	17, // 17: qan.v1.QANService.GetHistogram:output_type -> qan.v1.GetHistogramResponse
	18, // 18: qan.v1.QANService.ExplainFingerprintByQueryID:output_type -> qan.v1.ExplainFingerprintByQueryIDResponse
	19, // 19: qan.v1.QANService.GetQueryPlan:output_type -> qan.v1.GetQueryPlanResponse
	20, // 20: qan.v1.QANService.QueryExists:output_type -> qan.v1.QueryExistsResponse
	21, // 21: qan.v1.QANService.SchemaByQueryID:output_type -> qan.v1.SchemaByQueryIDResponse
	22, // 22: qan.v1.QANService.GetQueryExample:output_type -> qan.v1.GetQueryExampleResponse
	12, // [12:23] is the sub-list for method output_type
	1,  // [1:12] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_qan_v1_service_proto_init() }
func file_qan_v1_service_proto_init() {
	if File_qan_v1_service_proto != nil {
		return
	}
	file_qan_v1_filters_proto_init()
	file_qan_v1_object_details_proto_init()
	file_qan_v1_profile_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qan_v1_service_proto_rawDesc), len(file_qan_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qan_v1_service_proto_goTypes,
		DependencyIndexes: file_qan_v1_service_proto_depIdxs,
		MessageInfos:      file_qan_v1_service_proto_msgTypes,
	}.Build()
	File_qan_v1_service_proto = out.File
	file_qan_v1_service_proto_goTypes = nil
	file_qan_v1_service_proto_depIdxs = nil
}
