// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.25.3
// source: storage/blob.proto

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

// Next Tag: 7
type Blob struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" sql:"pk" search:"Blob Name"` // @gotags: sql:"pk" search:"Blob Name"
	Oid           uint32                 `protobuf:"varint,2,opt,name=oid,proto3" json:"oid,omitempty"`
	Checksum      string                 `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Length        int64                  `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty" search:"Blob Length"` // @gotags: search:"Blob Length"
	LastUpdated   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	ModifiedTime  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=modified_time,json=modifiedTime,proto3" json:"modified_time,omitempty" search:"Blob Modified On"` // @gotags: search:"Blob Modified On"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Blob) Reset() {
	*x = Blob{}
	mi := &file_storage_blob_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Blob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blob) ProtoMessage() {}

func (x *Blob) ProtoReflect() protoreflect.Message {
	mi := &file_storage_blob_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blob.ProtoReflect.Descriptor instead.
func (*Blob) Descriptor() ([]byte, []int) {
	return file_storage_blob_proto_rawDescGZIP(), []int{0}
}

func (x *Blob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Blob) GetOid() uint32 {
	if x != nil {
		return x.Oid
	}
	return 0
}

func (x *Blob) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

func (x *Blob) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Blob) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *Blob) GetModifiedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedTime
	}
	return nil
}

var File_storage_blob_proto protoreflect.FileDescriptor

var file_storage_blob_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0,
	0x01, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6f, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x3f, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5a, 0x11,
	0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_blob_proto_rawDescOnce sync.Once
	file_storage_blob_proto_rawDescData = file_storage_blob_proto_rawDesc
)

func file_storage_blob_proto_rawDescGZIP() []byte {
	file_storage_blob_proto_rawDescOnce.Do(func() {
		file_storage_blob_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_blob_proto_rawDescData)
	})
	return file_storage_blob_proto_rawDescData
}

var file_storage_blob_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_blob_proto_goTypes = []any{
	(*Blob)(nil),                  // 0: storage.Blob
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_storage_blob_proto_depIdxs = []int32{
	1, // 0: storage.Blob.last_updated:type_name -> google.protobuf.Timestamp
	1, // 1: storage.Blob.modified_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_storage_blob_proto_init() }
func file_storage_blob_proto_init() {
	if File_storage_blob_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_blob_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_blob_proto_goTypes,
		DependencyIndexes: file_storage_blob_proto_depIdxs,
		MessageInfos:      file_storage_blob_proto_msgTypes,
	}.Build()
	File_storage_blob_proto = out.File
	file_storage_blob_proto_rawDesc = nil
	file_storage_blob_proto_goTypes = nil
	file_storage_blob_proto_depIdxs = nil
}
