// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.25.3
// source: internalapi/sensor/collector.proto

package sensor

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

type ExternalIpsEnabled int32

const (
	ExternalIpsEnabled_DISABLED ExternalIpsEnabled = 0
	ExternalIpsEnabled_ENABLED  ExternalIpsEnabled = 1
)

// Enum value maps for ExternalIpsEnabled.
var (
	ExternalIpsEnabled_name = map[int32]string{
		0: "DISABLED",
		1: "ENABLED",
	}
	ExternalIpsEnabled_value = map[string]int32{
		"DISABLED": 0,
		"ENABLED":  1,
	}
)

func (x ExternalIpsEnabled) Enum() *ExternalIpsEnabled {
	p := new(ExternalIpsEnabled)
	*p = x
	return p
}

func (x ExternalIpsEnabled) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExternalIpsEnabled) Descriptor() protoreflect.EnumDescriptor {
	return file_internalapi_sensor_collector_proto_enumTypes[0].Descriptor()
}

func (ExternalIpsEnabled) Type() protoreflect.EnumType {
	return &file_internalapi_sensor_collector_proto_enumTypes[0]
}

func (x ExternalIpsEnabled) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExternalIpsEnabled.Descriptor instead.
func (ExternalIpsEnabled) EnumDescriptor() ([]byte, []int) {
	return file_internalapi_sensor_collector_proto_rawDescGZIP(), []int{0}
}

// A request message sent by collector to register with Sensor. Typically the first message in any streams.
type CollectorRegisterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The hostname on which collector is running.
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// A unique identifier for an instance of collector.
	InstanceId    string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CollectorRegisterRequest) Reset() {
	*x = CollectorRegisterRequest{}
	mi := &file_internalapi_sensor_collector_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CollectorRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectorRegisterRequest) ProtoMessage() {}

func (x *CollectorRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_collector_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectorRegisterRequest.ProtoReflect.Descriptor instead.
func (*CollectorRegisterRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_collector_proto_rawDescGZIP(), []int{0}
}

func (x *CollectorRegisterRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *CollectorRegisterRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

// CollectorConfig controls which type of data is reported by collector
// and how it is processed by collector. These configurations are used
// to fine-tune collector's performance on large scale clusters.
// At this time it only controls if external IPs are aggregated at the
// cluster level and the maximum number of open connections reported
// for each container per minute.
type CollectorConfig struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Networking    *CollectorConfig_Networking `protobuf:"bytes,1,opt,name=networking,proto3" json:"networking,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CollectorConfig) Reset() {
	*x = CollectorConfig{}
	mi := &file_internalapi_sensor_collector_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CollectorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectorConfig) ProtoMessage() {}

func (x *CollectorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_collector_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectorConfig.ProtoReflect.Descriptor instead.
func (*CollectorConfig) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_collector_proto_rawDescGZIP(), []int{1}
}

func (x *CollectorConfig) GetNetworking() *CollectorConfig_Networking {
	if x != nil {
		return x.Networking
	}
	return nil
}

type CollectorConfig_ExternalIPs struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enabled       ExternalIpsEnabled     `protobuf:"varint,1,opt,name=enabled,proto3,enum=sensor.ExternalIpsEnabled" json:"enabled,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CollectorConfig_ExternalIPs) Reset() {
	*x = CollectorConfig_ExternalIPs{}
	mi := &file_internalapi_sensor_collector_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CollectorConfig_ExternalIPs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectorConfig_ExternalIPs) ProtoMessage() {}

func (x *CollectorConfig_ExternalIPs) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_collector_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectorConfig_ExternalIPs.ProtoReflect.Descriptor instead.
func (*CollectorConfig_ExternalIPs) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_collector_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CollectorConfig_ExternalIPs) GetEnabled() ExternalIpsEnabled {
	if x != nil {
		return x.Enabled
	}
	return ExternalIpsEnabled_DISABLED
}

type CollectorConfig_Networking struct {
	state                   protoimpl.MessageState       `protogen:"open.v1"`
	ExternalIps             *CollectorConfig_ExternalIPs `protobuf:"bytes,1,opt,name=external_ips,json=externalIps,proto3" json:"external_ips,omitempty"`
	MaxConnectionsPerMinute int64                        `protobuf:"varint,2,opt,name=max_connections_per_minute,json=maxConnectionsPerMinute,proto3" json:"max_connections_per_minute,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *CollectorConfig_Networking) Reset() {
	*x = CollectorConfig_Networking{}
	mi := &file_internalapi_sensor_collector_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CollectorConfig_Networking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectorConfig_Networking) ProtoMessage() {}

func (x *CollectorConfig_Networking) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_collector_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectorConfig_Networking.ProtoReflect.Descriptor instead.
func (*CollectorConfig_Networking) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_collector_proto_rawDescGZIP(), []int{1, 1}
}

func (x *CollectorConfig_Networking) GetExternalIps() *CollectorConfig_ExternalIPs {
	if x != nil {
		return x.ExternalIps
	}
	return nil
}

func (x *CollectorConfig_Networking) GetMaxConnectionsPerMinute() int64 {
	if x != nil {
		return x.MaxConnectionsPerMinute
	}
	return 0
}

var File_internalapi_sensor_collector_proto protoreflect.FileDescriptor

var file_internalapi_sensor_collector_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x22, 0x57, 0x0a, 0x18,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0xae, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x42, 0x0a, 0x0a, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x43, 0x0a,
	0x0b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x50, 0x73, 0x12, 0x34, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x70, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x1a, 0x91, 0x01, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x46, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x50, 0x73, 0x52, 0x0b, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x6d, 0x61, 0x78,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6d,
	0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x65, 0x72,
	0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x2a, 0x2f, 0x0a, 0x12, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x70, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x42, 0x20, 0x5a, 0x1b, 0x2e, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x3b,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internalapi_sensor_collector_proto_rawDescOnce sync.Once
	file_internalapi_sensor_collector_proto_rawDescData = file_internalapi_sensor_collector_proto_rawDesc
)

func file_internalapi_sensor_collector_proto_rawDescGZIP() []byte {
	file_internalapi_sensor_collector_proto_rawDescOnce.Do(func() {
		file_internalapi_sensor_collector_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_sensor_collector_proto_rawDescData)
	})
	return file_internalapi_sensor_collector_proto_rawDescData
}

var file_internalapi_sensor_collector_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internalapi_sensor_collector_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internalapi_sensor_collector_proto_goTypes = []any{
	(ExternalIpsEnabled)(0),             // 0: sensor.ExternalIpsEnabled
	(*CollectorRegisterRequest)(nil),    // 1: sensor.CollectorRegisterRequest
	(*CollectorConfig)(nil),             // 2: sensor.CollectorConfig
	(*CollectorConfig_ExternalIPs)(nil), // 3: sensor.CollectorConfig.ExternalIPs
	(*CollectorConfig_Networking)(nil),  // 4: sensor.CollectorConfig.Networking
}
var file_internalapi_sensor_collector_proto_depIdxs = []int32{
	4, // 0: sensor.CollectorConfig.networking:type_name -> sensor.CollectorConfig.Networking
	0, // 1: sensor.CollectorConfig.ExternalIPs.enabled:type_name -> sensor.ExternalIpsEnabled
	3, // 2: sensor.CollectorConfig.Networking.external_ips:type_name -> sensor.CollectorConfig.ExternalIPs
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internalapi_sensor_collector_proto_init() }
func file_internalapi_sensor_collector_proto_init() {
	if File_internalapi_sensor_collector_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_sensor_collector_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internalapi_sensor_collector_proto_goTypes,
		DependencyIndexes: file_internalapi_sensor_collector_proto_depIdxs,
		EnumInfos:         file_internalapi_sensor_collector_proto_enumTypes,
		MessageInfos:      file_internalapi_sensor_collector_proto_msgTypes,
	}.Build()
	File_internalapi_sensor_collector_proto = out.File
	file_internalapi_sensor_collector_proto_rawDesc = nil
	file_internalapi_sensor_collector_proto_goTypes = nil
	file_internalapi_sensor_collector_proto_depIdxs = nil
}
