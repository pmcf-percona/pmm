// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: managementpb/dbaas/logs.proto

package dbaasv1beta1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Logs contain logs for certain pod's container. If container is an empty
// string, logs contain pod's events.
type Logs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Pod name.
	Pod string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	// Container name.
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// Content of container's log or pod's events.
	Logs []string `protobuf:"bytes,3,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *Logs) Reset() {
	*x = Logs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logs) ProtoMessage() {}

func (x *Logs) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logs.ProtoReflect.Descriptor instead.
func (*Logs) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_logs_proto_rawDescGZIP(), []int{0}
}

func (x *Logs) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *Logs) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

func (x *Logs) GetLogs() []string {
	if x != nil {
		return x.Logs
	}
	return nil
}

type GetLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kubernetes cluster name.
	KubernetesClusterName string `protobuf:"bytes,1,opt,name=kubernetes_cluster_name,json=kubernetesClusterName,proto3" json:"kubernetes_cluster_name,omitempty"`
	// DB cluster name.
	ClusterName string `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (x *GetLogsRequest) Reset() {
	*x = GetLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_logs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsRequest) ProtoMessage() {}

func (x *GetLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_logs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsRequest.ProtoReflect.Descriptor instead.
func (*GetLogsRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_logs_proto_rawDescGZIP(), []int{1}
}

func (x *GetLogsRequest) GetKubernetesClusterName() string {
	if x != nil {
		return x.KubernetesClusterName
	}
	return ""
}

func (x *GetLogsRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type GetLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Log represents list of logs. Each entry contains either container's logs or,
	// when container field is empty, pod's events.
	Logs []*Logs `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *GetLogsResponse) Reset() {
	*x = GetLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_dbaas_logs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsResponse) ProtoMessage() {}

func (x *GetLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_dbaas_logs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsResponse.ProtoReflect.Descriptor instead.
func (*GetLogsResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_dbaas_logs_proto_rawDescGZIP(), []int{2}
}

func (x *GetLogsResponse) GetLogs() []*Logs {
	if x != nil {
		return x.Logs
	}
	return nil
}

var File_managementpb_dbaas_logs_proto protoreflect.FileDescriptor

var file_managementpb_dbaas_logs_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x64,
	0x62, 0x61, 0x61, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x22, 0x7d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x17, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x15, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x32, 0x7c, 0x0a, 0x07,
	0x4c, 0x6f, 0x67, 0x73, 0x41, 0x50, 0x49, 0x12, 0x71, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x12, 0x1d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x44, 0x42, 0x61,
	0x61, 0x53, 0x2f, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0xaf, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x42, 0x09, 0x4c, 0x6f, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x64, 0x62, 0x61, 0x61, 0x73, 0x3b, 0x64, 0x62, 0x61,
	0x61, 0x73, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58, 0xaa,
	0x02, 0x0d, 0x44, 0x62, 0x61, 0x61, 0x73, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca,
	0x02, 0x0d, 0x44, 0x62, 0x61, 0x61, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2,
	0x02, 0x19, 0x44, 0x62, 0x61, 0x61, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x44, 0x62,
	0x61, 0x61, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_dbaas_logs_proto_rawDescOnce sync.Once
	file_managementpb_dbaas_logs_proto_rawDescData = file_managementpb_dbaas_logs_proto_rawDesc
)

func file_managementpb_dbaas_logs_proto_rawDescGZIP() []byte {
	file_managementpb_dbaas_logs_proto_rawDescOnce.Do(func() {
		file_managementpb_dbaas_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_dbaas_logs_proto_rawDescData)
	})
	return file_managementpb_dbaas_logs_proto_rawDescData
}

var (
	file_managementpb_dbaas_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_managementpb_dbaas_logs_proto_goTypes  = []interface{}{
		(*Logs)(nil),            // 0: dbaas.v1beta1.Logs
		(*GetLogsRequest)(nil),  // 1: dbaas.v1beta1.GetLogsRequest
		(*GetLogsResponse)(nil), // 2: dbaas.v1beta1.GetLogsResponse
	}
)

var file_managementpb_dbaas_logs_proto_depIdxs = []int32{
	0, // 0: dbaas.v1beta1.GetLogsResponse.logs:type_name -> dbaas.v1beta1.Logs
	1, // 1: dbaas.v1beta1.LogsAPI.GetLogs:input_type -> dbaas.v1beta1.GetLogsRequest
	2, // 2: dbaas.v1beta1.LogsAPI.GetLogs:output_type -> dbaas.v1beta1.GetLogsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_managementpb_dbaas_logs_proto_init() }
func file_managementpb_dbaas_logs_proto_init() {
	if File_managementpb_dbaas_logs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managementpb_dbaas_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logs); i {
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
		file_managementpb_dbaas_logs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsRequest); i {
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
		file_managementpb_dbaas_logs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsResponse); i {
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
			RawDescriptor: file_managementpb_dbaas_logs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_dbaas_logs_proto_goTypes,
		DependencyIndexes: file_managementpb_dbaas_logs_proto_depIdxs,
		MessageInfos:      file_managementpb_dbaas_logs_proto_msgTypes,
	}.Build()
	File_managementpb_dbaas_logs_proto = out.File
	file_managementpb_dbaas_logs_proto_rawDesc = nil
	file_managementpb_dbaas_logs_proto_goTypes = nil
	file_managementpb_dbaas_logs_proto_depIdxs = nil
}
