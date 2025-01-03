// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.3
// source: api/v2/compliance_integration_service.proto

package v2

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

// Represents the status of compliance operator
type COStatus int32

const (
	COStatus_HEALTHY   COStatus = 0
	COStatus_UNHEALTHY COStatus = 1
)

// Enum value maps for COStatus.
var (
	COStatus_name = map[int32]string{
		0: "HEALTHY",
		1: "UNHEALTHY",
	}
	COStatus_value = map[string]int32{
		"HEALTHY":   0,
		"UNHEALTHY": 1,
	}
)

func (x COStatus) Enum() *COStatus {
	p := new(COStatus)
	*p = x
	return p
}

func (x COStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (COStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v2_compliance_integration_service_proto_enumTypes[0].Descriptor()
}

func (COStatus) Type() protoreflect.EnumType {
	return &file_api_v2_compliance_integration_service_proto_enumTypes[0]
}

func (x COStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use COStatus.Descriptor instead.
func (COStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_v2_compliance_integration_service_proto_rawDescGZIP(), []int{0}
}

type ClusterProviderType int32

const (
	ClusterProviderType_UNSPECIFIED ClusterProviderType = 0
	ClusterProviderType_AKS         ClusterProviderType = 1
	ClusterProviderType_ARO         ClusterProviderType = 2
	ClusterProviderType_EKS         ClusterProviderType = 3
	ClusterProviderType_GKE         ClusterProviderType = 4
	ClusterProviderType_OCP         ClusterProviderType = 5
	ClusterProviderType_OSD         ClusterProviderType = 6
	ClusterProviderType_ROSA        ClusterProviderType = 7
)

// Enum value maps for ClusterProviderType.
var (
	ClusterProviderType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "AKS",
		2: "ARO",
		3: "EKS",
		4: "GKE",
		5: "OCP",
		6: "OSD",
		7: "ROSA",
	}
	ClusterProviderType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"AKS":         1,
		"ARO":         2,
		"EKS":         3,
		"GKE":         4,
		"OCP":         5,
		"OSD":         6,
		"ROSA":        7,
	}
)

func (x ClusterProviderType) Enum() *ClusterProviderType {
	p := new(ClusterProviderType)
	*p = x
	return p
}

func (x ClusterProviderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterProviderType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v2_compliance_integration_service_proto_enumTypes[1].Descriptor()
}

func (ClusterProviderType) Type() protoreflect.EnumType {
	return &file_api_v2_compliance_integration_service_proto_enumTypes[1]
}

func (x ClusterProviderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterProviderType.Descriptor instead.
func (ClusterProviderType) EnumDescriptor() ([]byte, []int) {
	return file_api_v2_compliance_integration_service_proto_rawDescGZIP(), []int{1}
}

type ClusterPlatformType int32

const (
	ClusterPlatformType_GENERIC_CLUSTER    ClusterPlatformType = 0
	ClusterPlatformType_KUBERNETES_CLUSTER ClusterPlatformType = 1
	ClusterPlatformType_OPENSHIFT_CLUSTER  ClusterPlatformType = 2
	ClusterPlatformType_OPENSHIFT4_CLUSTER ClusterPlatformType = 5
)

// Enum value maps for ClusterPlatformType.
var (
	ClusterPlatformType_name = map[int32]string{
		0: "GENERIC_CLUSTER",
		1: "KUBERNETES_CLUSTER",
		2: "OPENSHIFT_CLUSTER",
		5: "OPENSHIFT4_CLUSTER",
	}
	ClusterPlatformType_value = map[string]int32{
		"GENERIC_CLUSTER":    0,
		"KUBERNETES_CLUSTER": 1,
		"OPENSHIFT_CLUSTER":  2,
		"OPENSHIFT4_CLUSTER": 5,
	}
)

func (x ClusterPlatformType) Enum() *ClusterPlatformType {
	p := new(ClusterPlatformType)
	*p = x
	return p
}

func (x ClusterPlatformType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterPlatformType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v2_compliance_integration_service_proto_enumTypes[2].Descriptor()
}

func (ClusterPlatformType) Type() protoreflect.EnumType {
	return &file_api_v2_compliance_integration_service_proto_enumTypes[2]
}

func (x ClusterPlatformType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterPlatformType.Descriptor instead.
func (ClusterPlatformType) EnumDescriptor() ([]byte, []int) {
	return file_api_v2_compliance_integration_service_proto_rawDescGZIP(), []int{2}
}

// Next Tag: 11
type ComplianceIntegration struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version     string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	ClusterId   string                 `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ClusterName string                 `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Namespace   string                 `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Collection of errors that occurred while trying to obtain compliance operator health info.
	StatusErrors        []string            `protobuf:"bytes,6,rep,name=status_errors,json=statusErrors,proto3" json:"status_errors,omitempty"`
	OperatorInstalled   bool                `protobuf:"varint,7,opt,name=operator_installed,json=operatorInstalled,proto3" json:"operator_installed,omitempty"`
	Status              COStatus            `protobuf:"varint,8,opt,name=status,proto3,enum=v2.COStatus" json:"status,omitempty"`
	ClusterPlatformType ClusterPlatformType `protobuf:"varint,9,opt,name=cluster_platform_type,json=clusterPlatformType,proto3,enum=v2.ClusterPlatformType" json:"cluster_platform_type,omitempty"`
	ClusterProviderType ClusterProviderType `protobuf:"varint,10,opt,name=cluster_provider_type,json=clusterProviderType,proto3,enum=v2.ClusterProviderType" json:"cluster_provider_type,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ComplianceIntegration) Reset() {
	*x = ComplianceIntegration{}
	mi := &file_api_v2_compliance_integration_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplianceIntegration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceIntegration) ProtoMessage() {}

func (x *ComplianceIntegration) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_integration_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceIntegration.ProtoReflect.Descriptor instead.
func (*ComplianceIntegration) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_integration_service_proto_rawDescGZIP(), []int{0}
}

func (x *ComplianceIntegration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ComplianceIntegration) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ComplianceIntegration) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ComplianceIntegration) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ComplianceIntegration) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ComplianceIntegration) GetStatusErrors() []string {
	if x != nil {
		return x.StatusErrors
	}
	return nil
}

func (x *ComplianceIntegration) GetOperatorInstalled() bool {
	if x != nil {
		return x.OperatorInstalled
	}
	return false
}

func (x *ComplianceIntegration) GetStatus() COStatus {
	if x != nil {
		return x.Status
	}
	return COStatus_HEALTHY
}

func (x *ComplianceIntegration) GetClusterPlatformType() ClusterPlatformType {
	if x != nil {
		return x.ClusterPlatformType
	}
	return ClusterPlatformType_GENERIC_CLUSTER
}

func (x *ComplianceIntegration) GetClusterProviderType() ClusterProviderType {
	if x != nil {
		return x.ClusterProviderType
	}
	return ClusterProviderType_UNSPECIFIED
}

type ComplianceIntegrationStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClusterId     string                 `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ComplianceIntegrationStatusRequest) Reset() {
	*x = ComplianceIntegrationStatusRequest{}
	mi := &file_api_v2_compliance_integration_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComplianceIntegrationStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplianceIntegrationStatusRequest) ProtoMessage() {}

func (x *ComplianceIntegrationStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_integration_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplianceIntegrationStatusRequest.ProtoReflect.Descriptor instead.
func (*ComplianceIntegrationStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_integration_service_proto_rawDescGZIP(), []int{1}
}

func (x *ComplianceIntegrationStatusRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

type ListComplianceIntegrationsResponse struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Integrations  []*ComplianceIntegration `protobuf:"bytes,1,rep,name=integrations,proto3" json:"integrations,omitempty"`
	TotalCount    int32                    `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListComplianceIntegrationsResponse) Reset() {
	*x = ListComplianceIntegrationsResponse{}
	mi := &file_api_v2_compliance_integration_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListComplianceIntegrationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComplianceIntegrationsResponse) ProtoMessage() {}

func (x *ListComplianceIntegrationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_compliance_integration_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComplianceIntegrationsResponse.ProtoReflect.Descriptor instead.
func (*ListComplianceIntegrationsResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_compliance_integration_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListComplianceIntegrationsResponse) GetIntegrations() []*ComplianceIntegration {
	if x != nil {
		return x.Integrations
	}
	return nil
}

func (x *ListComplianceIntegrationsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_api_v2_compliance_integration_service_proto protoreflect.FileDescriptor

var file_api_v2_compliance_integration_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x32, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x03, 0x0a, 0x15, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x65, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x4f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4b, 0x0a, 0x15, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x13, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x15, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x43, 0x0a, 0x22, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x22, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x26,
	0x0a, 0x08, 0x43, 0x4f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45,
	0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x48, 0x45, 0x41,
	0x4c, 0x54, 0x48, 0x59, 0x10, 0x01, 0x2a, 0x66, 0x0a, 0x13, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x4b, 0x53, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x52, 0x4f, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x45, 0x4b, 0x53, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x4b, 0x45,
	0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x43, 0x50, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x4f,
	0x53, 0x44, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x53, 0x41, 0x10, 0x07, 0x2a, 0x7d,
	0x0a, 0x13, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43,
	0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4b, 0x55,
	0x42, 0x45, 0x52, 0x4e, 0x45, 0x54, 0x45, 0x53, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x50, 0x45, 0x4e, 0x53, 0x48, 0x49, 0x46, 0x54, 0x5f,
	0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x50, 0x45,
	0x4e, 0x53, 0x48, 0x49, 0x46, 0x54, 0x34, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10,
	0x05, 0x22, 0x04, 0x08, 0x03, 0x10, 0x03, 0x22, 0x04, 0x08, 0x04, 0x10, 0x04, 0x32, 0x97, 0x01,
	0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77,
	0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x2e, 0x76,
	0x32, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x26, 0x2e, 0x76, 0x32, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x32, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x27, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32,
	0x58, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v2_compliance_integration_service_proto_rawDescOnce sync.Once
	file_api_v2_compliance_integration_service_proto_rawDescData = file_api_v2_compliance_integration_service_proto_rawDesc
)

func file_api_v2_compliance_integration_service_proto_rawDescGZIP() []byte {
	file_api_v2_compliance_integration_service_proto_rawDescOnce.Do(func() {
		file_api_v2_compliance_integration_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_compliance_integration_service_proto_rawDescData)
	})
	return file_api_v2_compliance_integration_service_proto_rawDescData
}

var file_api_v2_compliance_integration_service_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_v2_compliance_integration_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v2_compliance_integration_service_proto_goTypes = []any{
	(COStatus)(0),                              // 0: v2.COStatus
	(ClusterProviderType)(0),                   // 1: v2.ClusterProviderType
	(ClusterPlatformType)(0),                   // 2: v2.ClusterPlatformType
	(*ComplianceIntegration)(nil),              // 3: v2.ComplianceIntegration
	(*ComplianceIntegrationStatusRequest)(nil), // 4: v2.ComplianceIntegrationStatusRequest
	(*ListComplianceIntegrationsResponse)(nil), // 5: v2.ListComplianceIntegrationsResponse
	(*RawQuery)(nil),                           // 6: v2.RawQuery
}
var file_api_v2_compliance_integration_service_proto_depIdxs = []int32{
	0, // 0: v2.ComplianceIntegration.status:type_name -> v2.COStatus
	2, // 1: v2.ComplianceIntegration.cluster_platform_type:type_name -> v2.ClusterPlatformType
	1, // 2: v2.ComplianceIntegration.cluster_provider_type:type_name -> v2.ClusterProviderType
	3, // 3: v2.ListComplianceIntegrationsResponse.integrations:type_name -> v2.ComplianceIntegration
	6, // 4: v2.ComplianceIntegrationService.ListComplianceIntegrations:input_type -> v2.RawQuery
	5, // 5: v2.ComplianceIntegrationService.ListComplianceIntegrations:output_type -> v2.ListComplianceIntegrationsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v2_compliance_integration_service_proto_init() }
func file_api_v2_compliance_integration_service_proto_init() {
	if File_api_v2_compliance_integration_service_proto != nil {
		return
	}
	file_api_v2_search_query_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v2_compliance_integration_service_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_compliance_integration_service_proto_goTypes,
		DependencyIndexes: file_api_v2_compliance_integration_service_proto_depIdxs,
		EnumInfos:         file_api_v2_compliance_integration_service_proto_enumTypes,
		MessageInfos:      file_api_v2_compliance_integration_service_proto_msgTypes,
	}.Build()
	File_api_v2_compliance_integration_service_proto = out.File
	file_api_v2_compliance_integration_service_proto_rawDesc = nil
	file_api_v2_compliance_integration_service_proto_goTypes = nil
	file_api_v2_compliance_integration_service_proto_depIdxs = nil
}
