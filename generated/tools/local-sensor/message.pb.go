// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.3
// source: tools/local-sensor/message.proto

package local_sensor

import (
	storage "github.com/stackrox/rox/generated/storage"
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

type LocalSensorPolicies struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Policies      []*storage.Policy      `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LocalSensorPolicies) Reset() {
	*x = LocalSensorPolicies{}
	mi := &file_tools_local_sensor_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocalSensorPolicies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalSensorPolicies) ProtoMessage() {}

func (x *LocalSensorPolicies) ProtoReflect() protoreflect.Message {
	mi := &file_tools_local_sensor_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalSensorPolicies.ProtoReflect.Descriptor instead.
func (*LocalSensorPolicies) Descriptor() ([]byte, []int) {
	return file_tools_local_sensor_message_proto_rawDescGZIP(), []int{0}
}

func (x *LocalSensorPolicies) GetPolicies() []*storage.Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

var File_tools_local_sensor_message_proto protoreflect.FileDescriptor

var file_tools_local_sensor_message_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2d, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x1a,
	0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x13, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x08,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tools_local_sensor_message_proto_rawDescOnce sync.Once
	file_tools_local_sensor_message_proto_rawDescData = file_tools_local_sensor_message_proto_rawDesc
)

func file_tools_local_sensor_message_proto_rawDescGZIP() []byte {
	file_tools_local_sensor_message_proto_rawDescOnce.Do(func() {
		file_tools_local_sensor_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_tools_local_sensor_message_proto_rawDescData)
	})
	return file_tools_local_sensor_message_proto_rawDescData
}

var file_tools_local_sensor_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tools_local_sensor_message_proto_goTypes = []any{
	(*LocalSensorPolicies)(nil), // 0: localSensor.LocalSensorPolicies
	(*storage.Policy)(nil),      // 1: storage.Policy
}
var file_tools_local_sensor_message_proto_depIdxs = []int32{
	1, // 0: localSensor.LocalSensorPolicies.policies:type_name -> storage.Policy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tools_local_sensor_message_proto_init() }
func file_tools_local_sensor_message_proto_init() {
	if File_tools_local_sensor_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tools_local_sensor_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tools_local_sensor_message_proto_goTypes,
		DependencyIndexes: file_tools_local_sensor_message_proto_depIdxs,
		MessageInfos:      file_tools_local_sensor_message_proto_msgTypes,
	}.Build()
	File_tools_local_sensor_message_proto = out.File
	file_tools_local_sensor_message_proto_rawDesc = nil
	file_tools_local_sensor_message_proto_goTypes = nil
	file_tools_local_sensor_message_proto_depIdxs = nil
}
