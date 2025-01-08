// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.25.3
// source: storage/group.proto

package storage

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

// Group is a GroupProperties : Role mapping.
type Group struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// GroupProperties define the properties of a group, applying to users when their properties match.
	// They also uniquely identify the group with the props.id field.
	Props *GroupProperties `protobuf:"bytes,1,opt,name=props,proto3" json:"props,omitempty"`
	// This is the name of the role that will apply to users in this group.
	RoleName      string `protobuf:"bytes,3,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty" search:"Role,hidden" sql:"index=name:groups_unique_indicator;category:unique"` // @gotags: search:"Role,hidden" sql:"index=name:groups_unique_indicator;category:unique"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_storage_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_storage_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_storage_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetProps() *GroupProperties {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *Group) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

// GroupProperties defines the properties of a group. Groups apply to users when
// their properties match. For instance:
//   - If GroupProperties has only an auth_provider_id, then that group applies
//     to all users logged in with that auth provider.
//   - If GroupProperties in addition has a claim key, then it applies to all
//     users with that auth provider and the claim key, etc.
//
// Note: Changes to GroupProperties may require changes to v1.DeleteGroupRequest.
type GroupProperties struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique identifier for group properties and respectively the group.
	Id             string  `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty" sql:"pk"` // @gotags: sql:"pk"
	Traits         *Traits `protobuf:"bytes,5,opt,name=traits,proto3" json:"traits,omitempty"`
	AuthProviderId string  `protobuf:"bytes,1,opt,name=auth_provider_id,json=authProviderId,proto3" json:"auth_provider_id,omitempty" search:"Group Auth Provider,hidden" sql:"index=category:unique;name:groups_unique_indicator"` // @gotags: search:"Group Auth Provider,hidden" sql:"index=category:unique;name:groups_unique_indicator"
	Key            string  `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty" search:"Group Key,hidden" sql:"index=category:unique;name:groups_unique_indicator"`                                               // @gotags: search:"Group Key,hidden" sql:"index=category:unique;name:groups_unique_indicator"
	Value          string  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty" search:"Group Value,hidden" sql:"index=category:unique;name:groups_unique_indicator"`                                           // @gotags: search:"Group Value,hidden" sql:"index=category:unique;name:groups_unique_indicator"
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GroupProperties) Reset() {
	*x = GroupProperties{}
	mi := &file_storage_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupProperties) ProtoMessage() {}

func (x *GroupProperties) ProtoReflect() protoreflect.Message {
	mi := &file_storage_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupProperties.ProtoReflect.Descriptor instead.
func (*GroupProperties) Descriptor() ([]byte, []int) {
	return file_storage_group_proto_rawDescGZIP(), []int{1}
}

func (x *GroupProperties) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GroupProperties) GetTraits() *Traits {
	if x != nil {
		return x.Traits
	}
	return nil
}

func (x *GroupProperties) GetAuthProviderId() string {
	if x != nil {
		return x.AuthProviderId
	}
	return ""
}

func (x *GroupProperties) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GroupProperties) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_storage_group_proto protoreflect.FileDescriptor

var file_storage_group_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x14,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2e, 0x0a,
	0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x0f, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27,
	0x0a, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x52,
	0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_storage_group_proto_rawDescOnce sync.Once
	file_storage_group_proto_rawDescData = file_storage_group_proto_rawDesc
)

func file_storage_group_proto_rawDescGZIP() []byte {
	file_storage_group_proto_rawDescOnce.Do(func() {
		file_storage_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_group_proto_rawDescData)
	})
	return file_storage_group_proto_rawDescData
}

var file_storage_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_storage_group_proto_goTypes = []any{
	(*Group)(nil),           // 0: storage.Group
	(*GroupProperties)(nil), // 1: storage.GroupProperties
	(*Traits)(nil),          // 2: storage.Traits
}
var file_storage_group_proto_depIdxs = []int32{
	1, // 0: storage.Group.props:type_name -> storage.GroupProperties
	2, // 1: storage.GroupProperties.traits:type_name -> storage.Traits
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_storage_group_proto_init() }
func file_storage_group_proto_init() {
	if File_storage_group_proto != nil {
		return
	}
	file_storage_traits_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_group_proto_goTypes,
		DependencyIndexes: file_storage_group_proto_depIdxs,
		MessageInfos:      file_storage_group_proto_msgTypes,
	}.Build()
	File_storage_group_proto = out.File
	file_storage_group_proto_rawDesc = nil
	file_storage_group_proto_goTypes = nil
	file_storage_group_proto_depIdxs = nil
}
