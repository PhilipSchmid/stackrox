// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.25.3
// source: api/v1/service_identity_service.proto

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

type ServiceIdentityResponse struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Identities    []*storage.ServiceIdentity `protobuf:"bytes,1,rep,name=identities,proto3" json:"identities,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceIdentityResponse) Reset() {
	*x = ServiceIdentityResponse{}
	mi := &file_api_v1_service_identity_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceIdentityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceIdentityResponse) ProtoMessage() {}

func (x *ServiceIdentityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_service_identity_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceIdentityResponse.ProtoReflect.Descriptor instead.
func (*ServiceIdentityResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_service_identity_service_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceIdentityResponse) GetIdentities() []*storage.ServiceIdentity {
	if x != nil {
		return x.Identities
	}
	return nil
}

type CreateServiceIdentityRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          storage.ServiceType    `protobuf:"varint,2,opt,name=type,proto3,enum=storage.ServiceType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateServiceIdentityRequest) Reset() {
	*x = CreateServiceIdentityRequest{}
	mi := &file_api_v1_service_identity_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateServiceIdentityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceIdentityRequest) ProtoMessage() {}

func (x *CreateServiceIdentityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_service_identity_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceIdentityRequest.ProtoReflect.Descriptor instead.
func (*CreateServiceIdentityRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_service_identity_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateServiceIdentityRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateServiceIdentityRequest) GetType() storage.ServiceType {
	if x != nil {
		return x.Type
	}
	return storage.ServiceType(0)
}

type CreateServiceIdentityResponse struct {
	state          protoimpl.MessageState   `protogen:"open.v1"`
	Identity       *storage.ServiceIdentity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	CertificatePem []byte                   `protobuf:"bytes,2,opt,name=certificate_pem,json=certificatePem,proto3" json:"certificate_pem,omitempty"`
	PrivateKeyPem  []byte                   `protobuf:"bytes,3,opt,name=private_key_pem,json=privateKeyPem,proto3" json:"private_key_pem,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateServiceIdentityResponse) Reset() {
	*x = CreateServiceIdentityResponse{}
	mi := &file_api_v1_service_identity_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateServiceIdentityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceIdentityResponse) ProtoMessage() {}

func (x *CreateServiceIdentityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_service_identity_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceIdentityResponse.ProtoReflect.Descriptor instead.
func (*CreateServiceIdentityResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_service_identity_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateServiceIdentityResponse) GetIdentity() *storage.ServiceIdentity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *CreateServiceIdentityResponse) GetCertificatePem() []byte {
	if x != nil {
		return x.CertificatePem
	}
	return nil
}

func (x *CreateServiceIdentityResponse) GetPrivateKeyPem() []byte {
	if x != nil {
		return x.PrivateKeyPem
	}
	return nil
}

type Authority struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CertificatePem []byte                 `protobuf:"bytes,1,opt,name=certificate_pem,json=certificatePem,proto3" json:"certificate_pem,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Authority) Reset() {
	*x = Authority{}
	mi := &file_api_v1_service_identity_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Authority) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authority) ProtoMessage() {}

func (x *Authority) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_service_identity_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authority.ProtoReflect.Descriptor instead.
func (*Authority) Descriptor() ([]byte, []int) {
	return file_api_v1_service_identity_service_proto_rawDescGZIP(), []int{3}
}

func (x *Authority) GetCertificatePem() []byte {
	if x != nil {
		return x.CertificatePem
	}
	return nil
}

type Authorities struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Authorities   []*Authority           `protobuf:"bytes,1,rep,name=authorities,proto3" json:"authorities,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Authorities) Reset() {
	*x = Authorities{}
	mi := &file_api_v1_service_identity_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Authorities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authorities) ProtoMessage() {}

func (x *Authorities) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_service_identity_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authorities.ProtoReflect.Descriptor instead.
func (*Authorities) Descriptor() ([]byte, []int) {
	return file_api_v1_service_identity_service_proto_rawDescGZIP(), []int{4}
}

func (x *Authorities) GetAuthorities() []*Authority {
	if x != nil {
		return x.Authorities
	}
	return nil
}

var File_api_v1_service_identity_service_proto protoreflect.FileDescriptor

var file_api_v1_service_identity_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x12, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a,
	0x17, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x22, 0x58, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa6, 0x01, 0x0a,
	0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6d, 0x12, 0x26, 0x0a,
	0x0f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x65, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x50, 0x65, 0x6d, 0x22, 0x34, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6d, 0x22, 0x3e, 0x0a, 0x0b, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x0b, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0b,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x32, 0xbe, 0x02, 0x0a, 0x16,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x09,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x7e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x20,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x27, 0x0a, 0x18,
	0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_service_identity_service_proto_rawDescOnce sync.Once
	file_api_v1_service_identity_service_proto_rawDescData = file_api_v1_service_identity_service_proto_rawDesc
)

func file_api_v1_service_identity_service_proto_rawDescGZIP() []byte {
	file_api_v1_service_identity_service_proto_rawDescOnce.Do(func() {
		file_api_v1_service_identity_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_service_identity_service_proto_rawDescData)
	})
	return file_api_v1_service_identity_service_proto_rawDescData
}

var file_api_v1_service_identity_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v1_service_identity_service_proto_goTypes = []any{
	(*ServiceIdentityResponse)(nil),       // 0: v1.ServiceIdentityResponse
	(*CreateServiceIdentityRequest)(nil),  // 1: v1.CreateServiceIdentityRequest
	(*CreateServiceIdentityResponse)(nil), // 2: v1.CreateServiceIdentityResponse
	(*Authority)(nil),                     // 3: v1.Authority
	(*Authorities)(nil),                   // 4: v1.Authorities
	(*storage.ServiceIdentity)(nil),       // 5: storage.ServiceIdentity
	(storage.ServiceType)(0),              // 6: storage.ServiceType
	(*Empty)(nil),                         // 7: v1.Empty
}
var file_api_v1_service_identity_service_proto_depIdxs = []int32{
	5, // 0: v1.ServiceIdentityResponse.identities:type_name -> storage.ServiceIdentity
	6, // 1: v1.CreateServiceIdentityRequest.type:type_name -> storage.ServiceType
	5, // 2: v1.CreateServiceIdentityResponse.identity:type_name -> storage.ServiceIdentity
	3, // 3: v1.Authorities.authorities:type_name -> v1.Authority
	7, // 4: v1.ServiceIdentityService.GetServiceIdentities:input_type -> v1.Empty
	1, // 5: v1.ServiceIdentityService.CreateServiceIdentity:input_type -> v1.CreateServiceIdentityRequest
	7, // 6: v1.ServiceIdentityService.GetAuthorities:input_type -> v1.Empty
	0, // 7: v1.ServiceIdentityService.GetServiceIdentities:output_type -> v1.ServiceIdentityResponse
	2, // 8: v1.ServiceIdentityService.CreateServiceIdentity:output_type -> v1.CreateServiceIdentityResponse
	4, // 9: v1.ServiceIdentityService.GetAuthorities:output_type -> v1.Authorities
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_service_identity_service_proto_init() }
func file_api_v1_service_identity_service_proto_init() {
	if File_api_v1_service_identity_service_proto != nil {
		return
	}
	file_api_v1_empty_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_service_identity_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_service_identity_service_proto_goTypes,
		DependencyIndexes: file_api_v1_service_identity_service_proto_depIdxs,
		MessageInfos:      file_api_v1_service_identity_service_proto_msgTypes,
	}.Build()
	File_api_v1_service_identity_service_proto = out.File
	file_api_v1_service_identity_service_proto_rawDesc = nil
	file_api_v1_service_identity_service_proto_goTypes = nil
	file_api_v1_service_identity_service_proto_depIdxs = nil
}
