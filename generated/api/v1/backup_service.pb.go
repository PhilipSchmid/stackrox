// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.3
// source: api/v1/backup_service.proto

package v1

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

type GetExternalBackupsResponse struct {
	state           protoimpl.MessageState    `protogen:"open.v1"`
	ExternalBackups []*storage.ExternalBackup `protobuf:"bytes,1,rep,name=external_backups,json=externalBackups,proto3" json:"external_backups,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetExternalBackupsResponse) Reset() {
	*x = GetExternalBackupsResponse{}
	mi := &file_api_v1_backup_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetExternalBackupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExternalBackupsResponse) ProtoMessage() {}

func (x *GetExternalBackupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_backup_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExternalBackupsResponse.ProtoReflect.Descriptor instead.
func (*GetExternalBackupsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_backup_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetExternalBackupsResponse) GetExternalBackups() []*storage.ExternalBackup {
	if x != nil {
		return x.ExternalBackups
	}
	return nil
}

type UpdateExternalBackupRequest struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	ExternalBackup *storage.ExternalBackup `protobuf:"bytes,1,opt,name=external_backup,json=externalBackup,proto3" json:"external_backup,omitempty"`
	// When false, use the stored credentials of an existing external backup configuration given its ID.
	UpdatePassword bool `protobuf:"varint,2,opt,name=update_password,json=updatePassword,proto3" json:"update_password,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateExternalBackupRequest) Reset() {
	*x = UpdateExternalBackupRequest{}
	mi := &file_api_v1_backup_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateExternalBackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExternalBackupRequest) ProtoMessage() {}

func (x *UpdateExternalBackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_backup_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExternalBackupRequest.ProtoReflect.Descriptor instead.
func (*UpdateExternalBackupRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_backup_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateExternalBackupRequest) GetExternalBackup() *storage.ExternalBackup {
	if x != nil {
		return x.ExternalBackup
	}
	return nil
}

func (x *UpdateExternalBackupRequest) GetUpdatePassword() bool {
	if x != nil {
		return x.UpdatePassword
	}
	return false
}

var File_api_v1_backup_service_proto protoreflect.FileDescriptor

var file_api_v1_backup_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x31, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x1b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x32, 0xb7, 0x07, 0x0a, 0x15, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x5c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x66,
	0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x1a, 0x17, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01,
	0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x6a, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x17, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x1a, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x5d, 0x0a, 0x12, 0x54, 0x65, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x12, 0x55, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x56, 0x0a, 0x15, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x79, 0x49, 0x44, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x85, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x32, 0x28,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2e, 0x69, 0x64, 0x7d, 0x12, 0x74, 0x0a, 0x19, 0x54, 0x65, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x27,
	0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1_backup_service_proto_rawDescOnce sync.Once
	file_api_v1_backup_service_proto_rawDescData = file_api_v1_backup_service_proto_rawDesc
)

func file_api_v1_backup_service_proto_rawDescGZIP() []byte {
	file_api_v1_backup_service_proto_rawDescOnce.Do(func() {
		file_api_v1_backup_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_backup_service_proto_rawDescData)
	})
	return file_api_v1_backup_service_proto_rawDescData
}

var file_api_v1_backup_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_backup_service_proto_goTypes = []any{
	(*GetExternalBackupsResponse)(nil),  // 0: v1.GetExternalBackupsResponse
	(*UpdateExternalBackupRequest)(nil), // 1: v1.UpdateExternalBackupRequest
	(*storage.ExternalBackup)(nil),      // 2: storage.ExternalBackup
	(*ResourceByID)(nil),                // 3: v1.ResourceByID
	(*Empty)(nil),                       // 4: v1.Empty
}
var file_api_v1_backup_service_proto_depIdxs = []int32{
	2,  // 0: v1.GetExternalBackupsResponse.external_backups:type_name -> storage.ExternalBackup
	2,  // 1: v1.UpdateExternalBackupRequest.external_backup:type_name -> storage.ExternalBackup
	3,  // 2: v1.ExternalBackupService.GetExternalBackup:input_type -> v1.ResourceByID
	4,  // 3: v1.ExternalBackupService.GetExternalBackups:input_type -> v1.Empty
	2,  // 4: v1.ExternalBackupService.PostExternalBackup:input_type -> storage.ExternalBackup
	2,  // 5: v1.ExternalBackupService.PutExternalBackup:input_type -> storage.ExternalBackup
	2,  // 6: v1.ExternalBackupService.TestExternalBackup:input_type -> storage.ExternalBackup
	3,  // 7: v1.ExternalBackupService.DeleteExternalBackup:input_type -> v1.ResourceByID
	3,  // 8: v1.ExternalBackupService.TriggerExternalBackup:input_type -> v1.ResourceByID
	1,  // 9: v1.ExternalBackupService.UpdateExternalBackup:input_type -> v1.UpdateExternalBackupRequest
	1,  // 10: v1.ExternalBackupService.TestUpdatedExternalBackup:input_type -> v1.UpdateExternalBackupRequest
	2,  // 11: v1.ExternalBackupService.GetExternalBackup:output_type -> storage.ExternalBackup
	0,  // 12: v1.ExternalBackupService.GetExternalBackups:output_type -> v1.GetExternalBackupsResponse
	2,  // 13: v1.ExternalBackupService.PostExternalBackup:output_type -> storage.ExternalBackup
	2,  // 14: v1.ExternalBackupService.PutExternalBackup:output_type -> storage.ExternalBackup
	4,  // 15: v1.ExternalBackupService.TestExternalBackup:output_type -> v1.Empty
	4,  // 16: v1.ExternalBackupService.DeleteExternalBackup:output_type -> v1.Empty
	4,  // 17: v1.ExternalBackupService.TriggerExternalBackup:output_type -> v1.Empty
	2,  // 18: v1.ExternalBackupService.UpdateExternalBackup:output_type -> storage.ExternalBackup
	4,  // 19: v1.ExternalBackupService.TestUpdatedExternalBackup:output_type -> v1.Empty
	11, // [11:20] is the sub-list for method output_type
	2,  // [2:11] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_backup_service_proto_init() }
func file_api_v1_backup_service_proto_init() {
	if File_api_v1_backup_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_backup_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_backup_service_proto_goTypes,
		DependencyIndexes: file_api_v1_backup_service_proto_depIdxs,
		MessageInfos:      file_api_v1_backup_service_proto_msgTypes,
	}.Build()
	File_api_v1_backup_service_proto = out.File
	file_api_v1_backup_service_proto_rawDesc = nil
	file_api_v1_backup_service_proto_goTypes = nil
	file_api_v1_backup_service_proto_depIdxs = nil
}
