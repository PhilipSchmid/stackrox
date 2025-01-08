// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.25.3
// source: storage/declarative_config_health.proto

package storage

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

type DeclarativeConfigHealth_Status int32

const (
	DeclarativeConfigHealth_UNHEALTHY DeclarativeConfigHealth_Status = 0
	DeclarativeConfigHealth_HEALTHY   DeclarativeConfigHealth_Status = 1
)

// Enum value maps for DeclarativeConfigHealth_Status.
var (
	DeclarativeConfigHealth_Status_name = map[int32]string{
		0: "UNHEALTHY",
		1: "HEALTHY",
	}
	DeclarativeConfigHealth_Status_value = map[string]int32{
		"UNHEALTHY": 0,
		"HEALTHY":   1,
	}
)

func (x DeclarativeConfigHealth_Status) Enum() *DeclarativeConfigHealth_Status {
	p := new(DeclarativeConfigHealth_Status)
	*p = x
	return p
}

func (x DeclarativeConfigHealth_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeclarativeConfigHealth_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_declarative_config_health_proto_enumTypes[0].Descriptor()
}

func (DeclarativeConfigHealth_Status) Type() protoreflect.EnumType {
	return &file_storage_declarative_config_health_proto_enumTypes[0]
}

func (x DeclarativeConfigHealth_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeclarativeConfigHealth_Status.Descriptor instead.
func (DeclarativeConfigHealth_Status) EnumDescriptor() ([]byte, []int) {
	return file_storage_declarative_config_health_proto_rawDescGZIP(), []int{0, 0}
}

type DeclarativeConfigHealth_ResourceType int32

const (
	DeclarativeConfigHealth_CONFIG_MAP     DeclarativeConfigHealth_ResourceType = 0
	DeclarativeConfigHealth_ACCESS_SCOPE   DeclarativeConfigHealth_ResourceType = 1
	DeclarativeConfigHealth_PERMISSION_SET DeclarativeConfigHealth_ResourceType = 2
	DeclarativeConfigHealth_ROLE           DeclarativeConfigHealth_ResourceType = 3
	DeclarativeConfigHealth_AUTH_PROVIDER  DeclarativeConfigHealth_ResourceType = 4
	DeclarativeConfigHealth_GROUP          DeclarativeConfigHealth_ResourceType = 5
	DeclarativeConfigHealth_NOTIFIER       DeclarativeConfigHealth_ResourceType = 6
)

// Enum value maps for DeclarativeConfigHealth_ResourceType.
var (
	DeclarativeConfigHealth_ResourceType_name = map[int32]string{
		0: "CONFIG_MAP",
		1: "ACCESS_SCOPE",
		2: "PERMISSION_SET",
		3: "ROLE",
		4: "AUTH_PROVIDER",
		5: "GROUP",
		6: "NOTIFIER",
	}
	DeclarativeConfigHealth_ResourceType_value = map[string]int32{
		"CONFIG_MAP":     0,
		"ACCESS_SCOPE":   1,
		"PERMISSION_SET": 2,
		"ROLE":           3,
		"AUTH_PROVIDER":  4,
		"GROUP":          5,
		"NOTIFIER":       6,
	}
)

func (x DeclarativeConfigHealth_ResourceType) Enum() *DeclarativeConfigHealth_ResourceType {
	p := new(DeclarativeConfigHealth_ResourceType)
	*p = x
	return p
}

func (x DeclarativeConfigHealth_ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeclarativeConfigHealth_ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_declarative_config_health_proto_enumTypes[1].Descriptor()
}

func (DeclarativeConfigHealth_ResourceType) Type() protoreflect.EnumType {
	return &file_storage_declarative_config_health_proto_enumTypes[1]
}

func (x DeclarativeConfigHealth_ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeclarativeConfigHealth_ResourceType.Descriptor instead.
func (DeclarativeConfigHealth_ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_storage_declarative_config_health_proto_rawDescGZIP(), []int{0, 1}
}

type DeclarativeConfigHealth struct {
	state        protoimpl.MessageState               `protogen:"open.v1"`
	Id           string                               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk,type(uuid)"` // @gotags: sql:"pk,type(uuid)"
	Name         string                               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status       DeclarativeConfigHealth_Status       `protobuf:"varint,4,opt,name=status,proto3,enum=storage.DeclarativeConfigHealth_Status" json:"status,omitempty"`
	ErrorMessage string                               `protobuf:"bytes,5,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	ResourceName string                               `protobuf:"bytes,6,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	ResourceType DeclarativeConfigHealth_ResourceType `protobuf:"varint,7,opt,name=resource_type,json=resourceType,proto3,enum=storage.DeclarativeConfigHealth_ResourceType" json:"resource_type,omitempty"`
	// Timestamp when the current status was set.
	LastTimestamp *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=last_timestamp,json=lastTimestamp,proto3" json:"last_timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeclarativeConfigHealth) Reset() {
	*x = DeclarativeConfigHealth{}
	mi := &file_storage_declarative_config_health_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeclarativeConfigHealth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclarativeConfigHealth) ProtoMessage() {}

func (x *DeclarativeConfigHealth) ProtoReflect() protoreflect.Message {
	mi := &file_storage_declarative_config_health_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclarativeConfigHealth.ProtoReflect.Descriptor instead.
func (*DeclarativeConfigHealth) Descriptor() ([]byte, []int) {
	return file_storage_declarative_config_health_proto_rawDescGZIP(), []int{0}
}

func (x *DeclarativeConfigHealth) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeclarativeConfigHealth) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeclarativeConfigHealth) GetStatus() DeclarativeConfigHealth_Status {
	if x != nil {
		return x.Status
	}
	return DeclarativeConfigHealth_UNHEALTHY
}

func (x *DeclarativeConfigHealth) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *DeclarativeConfigHealth) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *DeclarativeConfigHealth) GetResourceType() DeclarativeConfigHealth_ResourceType {
	if x != nil {
		return x.ResourceType
	}
	return DeclarativeConfigHealth_CONFIG_MAP
}

func (x *DeclarativeConfigHealth) GetLastTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastTimestamp
	}
	return nil
}

var File_storage_declarative_config_health_proto protoreflect.FileDescriptor

var file_storage_declarative_config_health_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x81, 0x04, 0x0a, 0x17, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65,
	0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x52,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x24, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x01, 0x22, 0x7a, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x43,
	0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x4d, 0x41, 0x50, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a,
	0x0e, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x54, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x41,
	0x55, 0x54, 0x48, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x10, 0x04, 0x12, 0x09,
	0x0a, 0x05, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54,
	0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x06, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_declarative_config_health_proto_rawDescOnce sync.Once
	file_storage_declarative_config_health_proto_rawDescData = file_storage_declarative_config_health_proto_rawDesc
)

func file_storage_declarative_config_health_proto_rawDescGZIP() []byte {
	file_storage_declarative_config_health_proto_rawDescOnce.Do(func() {
		file_storage_declarative_config_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_declarative_config_health_proto_rawDescData)
	})
	return file_storage_declarative_config_health_proto_rawDescData
}

var file_storage_declarative_config_health_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_storage_declarative_config_health_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_declarative_config_health_proto_goTypes = []any{
	(DeclarativeConfigHealth_Status)(0),       // 0: storage.DeclarativeConfigHealth.Status
	(DeclarativeConfigHealth_ResourceType)(0), // 1: storage.DeclarativeConfigHealth.ResourceType
	(*DeclarativeConfigHealth)(nil),           // 2: storage.DeclarativeConfigHealth
	(*timestamppb.Timestamp)(nil),             // 3: google.protobuf.Timestamp
}
var file_storage_declarative_config_health_proto_depIdxs = []int32{
	0, // 0: storage.DeclarativeConfigHealth.status:type_name -> storage.DeclarativeConfigHealth.Status
	1, // 1: storage.DeclarativeConfigHealth.resource_type:type_name -> storage.DeclarativeConfigHealth.ResourceType
	3, // 2: storage.DeclarativeConfigHealth.last_timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storage_declarative_config_health_proto_init() }
func file_storage_declarative_config_health_proto_init() {
	if File_storage_declarative_config_health_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_declarative_config_health_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_declarative_config_health_proto_goTypes,
		DependencyIndexes: file_storage_declarative_config_health_proto_depIdxs,
		EnumInfos:         file_storage_declarative_config_health_proto_enumTypes,
		MessageInfos:      file_storage_declarative_config_health_proto_msgTypes,
	}.Build()
	File_storage_declarative_config_health_proto = out.File
	file_storage_declarative_config_health_proto_rawDesc = nil
	file_storage_declarative_config_health_proto_goTypes = nil
	file_storage_declarative_config_health_proto_depIdxs = nil
}
