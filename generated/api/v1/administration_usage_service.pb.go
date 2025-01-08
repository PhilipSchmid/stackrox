// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.25.3
// source: api/v1/administration_usage_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// TimeRange allows for requesting data by a time range.
type TimeRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	mi := &file_api_v1_administration_usage_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_administration_usage_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_api_v1_administration_usage_service_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRange) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TimeRange) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

// SecuredUnitsUsageResponse holds the values of the currently observable
// administration usage metrics.
type SecuredUnitsUsageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NumNodes      int64                  `protobuf:"varint,1,opt,name=num_nodes,json=numNodes,proto3" json:"num_nodes,omitempty"`
	NumCpuUnits   int64                  `protobuf:"varint,2,opt,name=num_cpu_units,json=numCpuUnits,proto3" json:"num_cpu_units,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SecuredUnitsUsageResponse) Reset() {
	*x = SecuredUnitsUsageResponse{}
	mi := &file_api_v1_administration_usage_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SecuredUnitsUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecuredUnitsUsageResponse) ProtoMessage() {}

func (x *SecuredUnitsUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_administration_usage_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecuredUnitsUsageResponse.ProtoReflect.Descriptor instead.
func (*SecuredUnitsUsageResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_administration_usage_service_proto_rawDescGZIP(), []int{1}
}

func (x *SecuredUnitsUsageResponse) GetNumNodes() int64 {
	if x != nil {
		return x.NumNodes
	}
	return 0
}

func (x *SecuredUnitsUsageResponse) GetNumCpuUnits() int64 {
	if x != nil {
		return x.NumCpuUnits
	}
	return 0
}

// MaxSecuredUnitsUsageResponse holds the maximum values of the secured nodes
// and CPU Units (as reported by Kubernetes) with the time at which these
// values were aggregated, with the aggregation period accuracy (1h).
type MaxSecuredUnitsUsageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MaxNodesAt    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=max_nodes_at,json=maxNodesAt,proto3" json:"max_nodes_at,omitempty"`
	MaxNodes      int64                  `protobuf:"varint,2,opt,name=max_nodes,json=maxNodes,proto3" json:"max_nodes,omitempty"`
	MaxCpuUnitsAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=max_cpu_units_at,json=maxCpuUnitsAt,proto3" json:"max_cpu_units_at,omitempty"`
	MaxCpuUnits   int64                  `protobuf:"varint,4,opt,name=max_cpu_units,json=maxCpuUnits,proto3" json:"max_cpu_units,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MaxSecuredUnitsUsageResponse) Reset() {
	*x = MaxSecuredUnitsUsageResponse{}
	mi := &file_api_v1_administration_usage_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MaxSecuredUnitsUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxSecuredUnitsUsageResponse) ProtoMessage() {}

func (x *MaxSecuredUnitsUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_administration_usage_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxSecuredUnitsUsageResponse.ProtoReflect.Descriptor instead.
func (*MaxSecuredUnitsUsageResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_administration_usage_service_proto_rawDescGZIP(), []int{2}
}

func (x *MaxSecuredUnitsUsageResponse) GetMaxNodesAt() *timestamppb.Timestamp {
	if x != nil {
		return x.MaxNodesAt
	}
	return nil
}

func (x *MaxSecuredUnitsUsageResponse) GetMaxNodes() int64 {
	if x != nil {
		return x.MaxNodes
	}
	return 0
}

func (x *MaxSecuredUnitsUsageResponse) GetMaxCpuUnitsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.MaxCpuUnitsAt
	}
	return nil
}

func (x *MaxSecuredUnitsUsageResponse) GetMaxCpuUnits() int64 {
	if x != nil {
		return x.MaxCpuUnits
	}
	return 0
}

var File_api_v1_administration_usage_service_proto protoreflect.FileDescriptor

var file_api_v1_administration_usage_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a,
	0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x67, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x19, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x75, 0x6d,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x70, 0x75,
	0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x75,
	0x6d, 0x43, 0x70, 0x75, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x22, 0xe2, 0x01, 0x0a, 0x1c, 0x4d, 0x61,
	0x78, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x6d, 0x61,
	0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d, 0x61,
	0x78, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x70, 0x75,
	0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6d, 0x61, 0x78,
	0x43, 0x70, 0x75, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61,
	0x78, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x43, 0x70, 0x75, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x32, 0x9d,
	0x02, 0x0a, 0x1a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7f, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x75, 0x72,
	0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x09, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64,
	0x2d, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x7e,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x78, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2c, 0x12, 0x2a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x65, 0x64, 0x2d, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2f, 0x6d, 0x61, 0x78, 0x42, 0x27,
	0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1_administration_usage_service_proto_rawDescOnce sync.Once
	file_api_v1_administration_usage_service_proto_rawDescData = file_api_v1_administration_usage_service_proto_rawDesc
)

func file_api_v1_administration_usage_service_proto_rawDescGZIP() []byte {
	file_api_v1_administration_usage_service_proto_rawDescOnce.Do(func() {
		file_api_v1_administration_usage_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_administration_usage_service_proto_rawDescData)
	})
	return file_api_v1_administration_usage_service_proto_rawDescData
}

var file_api_v1_administration_usage_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_administration_usage_service_proto_goTypes = []any{
	(*TimeRange)(nil),                    // 0: v1.TimeRange
	(*SecuredUnitsUsageResponse)(nil),    // 1: v1.SecuredUnitsUsageResponse
	(*MaxSecuredUnitsUsageResponse)(nil), // 2: v1.MaxSecuredUnitsUsageResponse
	(*timestamppb.Timestamp)(nil),        // 3: google.protobuf.Timestamp
	(*Empty)(nil),                        // 4: v1.Empty
}
var file_api_v1_administration_usage_service_proto_depIdxs = []int32{
	3, // 0: v1.TimeRange.from:type_name -> google.protobuf.Timestamp
	3, // 1: v1.TimeRange.to:type_name -> google.protobuf.Timestamp
	3, // 2: v1.MaxSecuredUnitsUsageResponse.max_nodes_at:type_name -> google.protobuf.Timestamp
	3, // 3: v1.MaxSecuredUnitsUsageResponse.max_cpu_units_at:type_name -> google.protobuf.Timestamp
	4, // 4: v1.AdministrationUsageService.GetCurrentSecuredUnitsUsage:input_type -> v1.Empty
	0, // 5: v1.AdministrationUsageService.GetMaxSecuredUnitsUsage:input_type -> v1.TimeRange
	1, // 6: v1.AdministrationUsageService.GetCurrentSecuredUnitsUsage:output_type -> v1.SecuredUnitsUsageResponse
	2, // 7: v1.AdministrationUsageService.GetMaxSecuredUnitsUsage:output_type -> v1.MaxSecuredUnitsUsageResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_administration_usage_service_proto_init() }
func file_api_v1_administration_usage_service_proto_init() {
	if File_api_v1_administration_usage_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_administration_usage_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_administration_usage_service_proto_goTypes,
		DependencyIndexes: file_api_v1_administration_usage_service_proto_depIdxs,
		MessageInfos:      file_api_v1_administration_usage_service_proto_msgTypes,
	}.Build()
	File_api_v1_administration_usage_service_proto = out.File
	file_api_v1_administration_usage_service_proto_rawDesc = nil
	file_api_v1_administration_usage_service_proto_goTypes = nil
	file_api_v1_administration_usage_service_proto_depIdxs = nil
}
