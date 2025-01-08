// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.25.3
// source: storage/version.proto

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

type Version struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This is a strictly incrementing integer describing the DB version.
	SeqNum int32 `protobuf:"varint,1,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty" search:"Current Sequence Number"` // @gotags: search:"Current Sequence Number"
	// Associated version metadata. (For example, the corresponding product version.)
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" search:"Version"` // @gotags: search:"Version"
	// Last time version was updated
	LastPersisted *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_persisted,json=lastPersisted,proto3" json:"last_persisted,omitempty" search:"Last Persisted"` // @gotags: search:"Last Persisted"
	// The minimum sequence number supported by this iteration of the database.  Rollbacks
	// to versions prior to this sequence number are not supported.
	MinSeqNum     int32 `protobuf:"varint,4,opt,name=min_seq_num,json=minSeqNum,proto3" json:"min_seq_num,omitempty" search:"Minimum Sequence Number"` // @gotags: search:"Minimum Sequence Number"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Version) Reset() {
	*x = Version{}
	mi := &file_storage_version_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_storage_version_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_storage_version_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetSeqNum() int32 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

func (x *Version) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Version) GetLastPersisted() *timestamppb.Timestamp {
	if x != nil {
		return x.LastPersisted
	}
	return nil
}

func (x *Version) GetMinSeqNum() int32 {
	if x != nil {
		return x.MinSeqNum
	}
	return 0
}

var File_storage_version_proto protoreflect.FileDescriptor

var file_storage_version_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x41, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x71, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x71,
	0x4e, 0x75, 0x6d, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72,
	0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_version_proto_rawDescOnce sync.Once
	file_storage_version_proto_rawDescData = file_storage_version_proto_rawDesc
)

func file_storage_version_proto_rawDescGZIP() []byte {
	file_storage_version_proto_rawDescOnce.Do(func() {
		file_storage_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_version_proto_rawDescData)
	})
	return file_storage_version_proto_rawDescData
}

var file_storage_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_version_proto_goTypes = []any{
	(*Version)(nil),               // 0: storage.Version
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_storage_version_proto_depIdxs = []int32{
	1, // 0: storage.Version.last_persisted:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_storage_version_proto_init() }
func file_storage_version_proto_init() {
	if File_storage_version_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_version_proto_goTypes,
		DependencyIndexes: file_storage_version_proto_depIdxs,
		MessageInfos:      file_storage_version_proto_msgTypes,
	}.Build()
	File_storage_version_proto = out.File
	file_storage_version_proto_rawDesc = nil
	file_storage_version_proto_goTypes = nil
	file_storage_version_proto_depIdxs = nil
}
