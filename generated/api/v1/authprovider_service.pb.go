// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.3
// source: api/v1/authprovider_service.proto

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

type GetAuthProviderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthProviderRequest) Reset() {
	*x = GetAuthProviderRequest{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthProviderRequest) ProtoMessage() {}

func (x *GetAuthProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthProviderRequest.ProtoReflect.Descriptor instead.
func (*GetAuthProviderRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAuthProviderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAuthProvidersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthProvidersRequest) Reset() {
	*x = GetAuthProvidersRequest{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthProvidersRequest) ProtoMessage() {}

func (x *GetAuthProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthProvidersRequest.ProtoReflect.Descriptor instead.
func (*GetAuthProvidersRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAuthProvidersRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAuthProvidersRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetLoginAuthProvidersResponse struct {
	state         protoimpl.MessageState                             `protogen:"open.v1"`
	AuthProviders []*GetLoginAuthProvidersResponse_LoginAuthProvider `protobuf:"bytes,1,rep,name=auth_providers,json=authProviders,proto3" json:"auth_providers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLoginAuthProvidersResponse) Reset() {
	*x = GetLoginAuthProvidersResponse{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLoginAuthProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginAuthProvidersResponse) ProtoMessage() {}

func (x *GetLoginAuthProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginAuthProvidersResponse.ProtoReflect.Descriptor instead.
func (*GetLoginAuthProvidersResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLoginAuthProvidersResponse) GetAuthProviders() []*GetLoginAuthProvidersResponse_LoginAuthProvider {
	if x != nil {
		return x.AuthProviders
	}
	return nil
}

type GetAuthProvidersResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	AuthProviders []*storage.AuthProvider `protobuf:"bytes,1,rep,name=auth_providers,json=authProviders,proto3" json:"auth_providers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthProvidersResponse) Reset() {
	*x = GetAuthProvidersResponse{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthProvidersResponse) ProtoMessage() {}

func (x *GetAuthProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthProvidersResponse.ProtoReflect.Descriptor instead.
func (*GetAuthProvidersResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuthProvidersResponse) GetAuthProviders() []*storage.AuthProvider {
	if x != nil {
		return x.AuthProviders
	}
	return nil
}

type PostAuthProviderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Provider      *storage.AuthProvider  `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostAuthProviderRequest) Reset() {
	*x = PostAuthProviderRequest{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostAuthProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostAuthProviderRequest) ProtoMessage() {}

func (x *PostAuthProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostAuthProviderRequest.ProtoReflect.Descriptor instead.
func (*PostAuthProviderRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{4}
}

func (x *PostAuthProviderRequest) GetProvider() *storage.AuthProvider {
	if x != nil {
		return x.Provider
	}
	return nil
}

type UpdateAuthProviderRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to NameOpt:
	//
	//	*UpdateAuthProviderRequest_Name
	NameOpt isUpdateAuthProviderRequest_NameOpt `protobuf_oneof:"name_opt"`
	// Types that are valid to be assigned to EnabledOpt:
	//
	//	*UpdateAuthProviderRequest_Enabled
	EnabledOpt    isUpdateAuthProviderRequest_EnabledOpt `protobuf_oneof:"enabled_opt"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAuthProviderRequest) Reset() {
	*x = UpdateAuthProviderRequest{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAuthProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthProviderRequest) ProtoMessage() {}

func (x *UpdateAuthProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthProviderRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuthProviderRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAuthProviderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAuthProviderRequest) GetNameOpt() isUpdateAuthProviderRequest_NameOpt {
	if x != nil {
		return x.NameOpt
	}
	return nil
}

func (x *UpdateAuthProviderRequest) GetName() string {
	if x != nil {
		if x, ok := x.NameOpt.(*UpdateAuthProviderRequest_Name); ok {
			return x.Name
		}
	}
	return ""
}

func (x *UpdateAuthProviderRequest) GetEnabledOpt() isUpdateAuthProviderRequest_EnabledOpt {
	if x != nil {
		return x.EnabledOpt
	}
	return nil
}

func (x *UpdateAuthProviderRequest) GetEnabled() bool {
	if x != nil {
		if x, ok := x.EnabledOpt.(*UpdateAuthProviderRequest_Enabled); ok {
			return x.Enabled
		}
	}
	return false
}

type isUpdateAuthProviderRequest_NameOpt interface {
	isUpdateAuthProviderRequest_NameOpt()
}

type UpdateAuthProviderRequest_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*UpdateAuthProviderRequest_Name) isUpdateAuthProviderRequest_NameOpt() {}

type isUpdateAuthProviderRequest_EnabledOpt interface {
	isUpdateAuthProviderRequest_EnabledOpt()
}

type UpdateAuthProviderRequest_Enabled struct {
	Enabled bool `protobuf:"varint,3,opt,name=enabled,proto3,oneof"`
}

func (*UpdateAuthProviderRequest_Enabled) isUpdateAuthProviderRequest_EnabledOpt() {}

type ExchangeTokenRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The external authentication token. The server will mask the value of this credential in responses and logs.
	ExternalToken string `protobuf:"bytes,1,opt,name=external_token,json=externalToken,proto3" json:"external_token,omitempty" scrub:"always"` // @gotags: scrub:"always"
	Type          string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	State         string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExchangeTokenRequest) Reset() {
	*x = ExchangeTokenRequest{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeTokenRequest) ProtoMessage() {}

func (x *ExchangeTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeTokenRequest.ProtoReflect.Descriptor instead.
func (*ExchangeTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{6}
}

func (x *ExchangeTokenRequest) GetExternalToken() string {
	if x != nil {
		return x.ExternalToken
	}
	return ""
}

func (x *ExchangeTokenRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ExchangeTokenRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type ExchangeTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ClientState   string                 `protobuf:"bytes,2,opt,name=client_state,json=clientState,proto3" json:"client_state,omitempty"`
	Test          bool                   `protobuf:"varint,3,opt,name=test,proto3" json:"test,omitempty"`
	User          *AuthStatus            `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExchangeTokenResponse) Reset() {
	*x = ExchangeTokenResponse{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeTokenResponse) ProtoMessage() {}

func (x *ExchangeTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeTokenResponse.ProtoReflect.Descriptor instead.
func (*ExchangeTokenResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{7}
}

func (x *ExchangeTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ExchangeTokenResponse) GetClientState() string {
	if x != nil {
		return x.ClientState
	}
	return ""
}

func (x *ExchangeTokenResponse) GetTest() bool {
	if x != nil {
		return x.Test
	}
	return false
}

func (x *ExchangeTokenResponse) GetUser() *AuthStatus {
	if x != nil {
		return x.User
	}
	return nil
}

type AvailableProviderTypesResponse struct {
	state             protoimpl.MessageState                             `protogen:"open.v1"`
	AuthProviderTypes []*AvailableProviderTypesResponse_AuthProviderType `protobuf:"bytes,1,rep,name=auth_provider_types,json=authProviderTypes,proto3" json:"auth_provider_types,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *AvailableProviderTypesResponse) Reset() {
	*x = AvailableProviderTypesResponse{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AvailableProviderTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableProviderTypesResponse) ProtoMessage() {}

func (x *AvailableProviderTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableProviderTypesResponse.ProtoReflect.Descriptor instead.
func (*AvailableProviderTypesResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{8}
}

func (x *AvailableProviderTypesResponse) GetAuthProviderTypes() []*AvailableProviderTypesResponse_AuthProviderType {
	if x != nil {
		return x.AuthProviderTypes
	}
	return nil
}

type GetLoginAuthProvidersResponse_LoginAuthProvider struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	LoginUrl      string                 `protobuf:"bytes,5,opt,name=login_url,json=loginUrl,proto3" json:"login_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLoginAuthProvidersResponse_LoginAuthProvider) Reset() {
	*x = GetLoginAuthProvidersResponse_LoginAuthProvider{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLoginAuthProvidersResponse_LoginAuthProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginAuthProvidersResponse_LoginAuthProvider) ProtoMessage() {}

func (x *GetLoginAuthProvidersResponse_LoginAuthProvider) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginAuthProvidersResponse_LoginAuthProvider.ProtoReflect.Descriptor instead.
func (*GetLoginAuthProvidersResponse_LoginAuthProvider) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetLoginAuthProvidersResponse_LoginAuthProvider) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLoginAuthProvidersResponse_LoginAuthProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetLoginAuthProvidersResponse_LoginAuthProvider) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetLoginAuthProvidersResponse_LoginAuthProvider) GetLoginUrl() string {
	if x != nil {
		return x.LoginUrl
	}
	return ""
}

type AvailableProviderTypesResponse_AuthProviderType struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Type                string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	SuggestedAttributes []string               `protobuf:"bytes,2,rep,name=suggested_attributes,json=suggestedAttributes,proto3" json:"suggested_attributes,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AvailableProviderTypesResponse_AuthProviderType) Reset() {
	*x = AvailableProviderTypesResponse_AuthProviderType{}
	mi := &file_api_v1_authprovider_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AvailableProviderTypesResponse_AuthProviderType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableProviderTypesResponse_AuthProviderType) ProtoMessage() {}

func (x *AvailableProviderTypesResponse_AuthProviderType) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_authprovider_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableProviderTypesResponse_AuthProviderType.ProtoReflect.Descriptor instead.
func (*AvailableProviderTypesResponse_AuthProviderType) Descriptor() ([]byte, []int) {
	return file_api_v1_authprovider_service_proto_rawDescGZIP(), []int{8, 0}
}

func (x *AvailableProviderTypesResponse_AuthProviderType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AvailableProviderTypesResponse_AuthProviderType) GetSuggestedAttributes() []string {
	if x != nil {
		return x.SuggestedAttributes
	}
	return nil
}

var File_api_v1_authprovider_service_proto protoreflect.FileDescriptor

var file_api_v1_authprovider_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x41, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x1a, 0x68, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0x58, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0x4c, 0x0a, 0x17, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x31, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x22, 0x78, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x42, 0x0d,
	0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x22, 0x67, 0x0a,
	0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x15, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0xe0, 0x01, 0x0a, 0x1e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x59, 0x0a, 0x10, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x31, 0x0a, 0x14, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x13, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x32, 0xca, 0x07, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x1a,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x64, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x09, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x68, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x6b, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22,
	0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x6d, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x32, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x62, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x1a, 0x16, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x58, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x57, 0x69, 0x74, 0x68, 0x46,
	0x6f, 0x72, 0x63, 0x65, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x70, 0x0a, 0x0d, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x18, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a,
	0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x03, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_authprovider_service_proto_rawDescOnce sync.Once
	file_api_v1_authprovider_service_proto_rawDescData = file_api_v1_authprovider_service_proto_rawDesc
)

func file_api_v1_authprovider_service_proto_rawDescGZIP() []byte {
	file_api_v1_authprovider_service_proto_rawDescOnce.Do(func() {
		file_api_v1_authprovider_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_authprovider_service_proto_rawDescData)
	})
	return file_api_v1_authprovider_service_proto_rawDescData
}

var file_api_v1_authprovider_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_v1_authprovider_service_proto_goTypes = []any{
	(*GetAuthProviderRequest)(nil),                          // 0: v1.GetAuthProviderRequest
	(*GetAuthProvidersRequest)(nil),                         // 1: v1.GetAuthProvidersRequest
	(*GetLoginAuthProvidersResponse)(nil),                   // 2: v1.GetLoginAuthProvidersResponse
	(*GetAuthProvidersResponse)(nil),                        // 3: v1.GetAuthProvidersResponse
	(*PostAuthProviderRequest)(nil),                         // 4: v1.PostAuthProviderRequest
	(*UpdateAuthProviderRequest)(nil),                       // 5: v1.UpdateAuthProviderRequest
	(*ExchangeTokenRequest)(nil),                            // 6: v1.ExchangeTokenRequest
	(*ExchangeTokenResponse)(nil),                           // 7: v1.ExchangeTokenResponse
	(*AvailableProviderTypesResponse)(nil),                  // 8: v1.AvailableProviderTypesResponse
	(*GetLoginAuthProvidersResponse_LoginAuthProvider)(nil), // 9: v1.GetLoginAuthProvidersResponse.LoginAuthProvider
	(*AvailableProviderTypesResponse_AuthProviderType)(nil), // 10: v1.AvailableProviderTypesResponse.AuthProviderType
	(*storage.AuthProvider)(nil),                            // 11: storage.AuthProvider
	(*AuthStatus)(nil),                                      // 12: v1.AuthStatus
	(*Empty)(nil),                                           // 13: v1.Empty
	(*DeleteByIDWithForce)(nil),                             // 14: v1.DeleteByIDWithForce
}
var file_api_v1_authprovider_service_proto_depIdxs = []int32{
	9,  // 0: v1.GetLoginAuthProvidersResponse.auth_providers:type_name -> v1.GetLoginAuthProvidersResponse.LoginAuthProvider
	11, // 1: v1.GetAuthProvidersResponse.auth_providers:type_name -> storage.AuthProvider
	11, // 2: v1.PostAuthProviderRequest.provider:type_name -> storage.AuthProvider
	12, // 3: v1.ExchangeTokenResponse.user:type_name -> v1.AuthStatus
	10, // 4: v1.AvailableProviderTypesResponse.auth_provider_types:type_name -> v1.AvailableProviderTypesResponse.AuthProviderType
	13, // 5: v1.AuthProviderService.ListAvailableProviderTypes:input_type -> v1.Empty
	0,  // 6: v1.AuthProviderService.GetAuthProvider:input_type -> v1.GetAuthProviderRequest
	13, // 7: v1.AuthProviderService.GetLoginAuthProviders:input_type -> v1.Empty
	1,  // 8: v1.AuthProviderService.GetAuthProviders:input_type -> v1.GetAuthProvidersRequest
	4,  // 9: v1.AuthProviderService.PostAuthProvider:input_type -> v1.PostAuthProviderRequest
	5,  // 10: v1.AuthProviderService.UpdateAuthProvider:input_type -> v1.UpdateAuthProviderRequest
	11, // 11: v1.AuthProviderService.PutAuthProvider:input_type -> storage.AuthProvider
	14, // 12: v1.AuthProviderService.DeleteAuthProvider:input_type -> v1.DeleteByIDWithForce
	6,  // 13: v1.AuthProviderService.ExchangeToken:input_type -> v1.ExchangeTokenRequest
	8,  // 14: v1.AuthProviderService.ListAvailableProviderTypes:output_type -> v1.AvailableProviderTypesResponse
	11, // 15: v1.AuthProviderService.GetAuthProvider:output_type -> storage.AuthProvider
	2,  // 16: v1.AuthProviderService.GetLoginAuthProviders:output_type -> v1.GetLoginAuthProvidersResponse
	3,  // 17: v1.AuthProviderService.GetAuthProviders:output_type -> v1.GetAuthProvidersResponse
	11, // 18: v1.AuthProviderService.PostAuthProvider:output_type -> storage.AuthProvider
	11, // 19: v1.AuthProviderService.UpdateAuthProvider:output_type -> storage.AuthProvider
	11, // 20: v1.AuthProviderService.PutAuthProvider:output_type -> storage.AuthProvider
	13, // 21: v1.AuthProviderService.DeleteAuthProvider:output_type -> v1.Empty
	7,  // 22: v1.AuthProviderService.ExchangeToken:output_type -> v1.ExchangeTokenResponse
	14, // [14:23] is the sub-list for method output_type
	5,  // [5:14] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_authprovider_service_proto_init() }
func file_api_v1_authprovider_service_proto_init() {
	if File_api_v1_authprovider_service_proto != nil {
		return
	}
	file_api_v1_auth_service_proto_init()
	file_api_v1_common_proto_init()
	file_api_v1_empty_proto_init()
	file_api_v1_authprovider_service_proto_msgTypes[5].OneofWrappers = []any{
		(*UpdateAuthProviderRequest_Name)(nil),
		(*UpdateAuthProviderRequest_Enabled)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_authprovider_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_authprovider_service_proto_goTypes,
		DependencyIndexes: file_api_v1_authprovider_service_proto_depIdxs,
		MessageInfos:      file_api_v1_authprovider_service_proto_msgTypes,
	}.Build()
	File_api_v1_authprovider_service_proto = out.File
	file_api_v1_authprovider_service_proto_rawDesc = nil
	file_api_v1_authprovider_service_proto_goTypes = nil
	file_api_v1_authprovider_service_proto_depIdxs = nil
}
