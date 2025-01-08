// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v4.25.3
// source: storage/report_configuration.proto

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

type ReportConfiguration_ReportType int32

const (
	ReportConfiguration_VULNERABILITY ReportConfiguration_ReportType = 0
)

// Enum value maps for ReportConfiguration_ReportType.
var (
	ReportConfiguration_ReportType_name = map[int32]string{
		0: "VULNERABILITY",
	}
	ReportConfiguration_ReportType_value = map[string]int32{
		"VULNERABILITY": 0,
	}
)

func (x ReportConfiguration_ReportType) Enum() *ReportConfiguration_ReportType {
	p := new(ReportConfiguration_ReportType)
	*p = x
	return p
}

func (x ReportConfiguration_ReportType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportConfiguration_ReportType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_configuration_proto_enumTypes[0].Descriptor()
}

func (ReportConfiguration_ReportType) Type() protoreflect.EnumType {
	return &file_storage_report_configuration_proto_enumTypes[0]
}

func (x ReportConfiguration_ReportType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportConfiguration_ReportType.Descriptor instead.
func (ReportConfiguration_ReportType) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{0, 0}
}

type ReportLastRunStatus_RunStatus int32

const (
	ReportLastRunStatus_SUCCESS ReportLastRunStatus_RunStatus = 0
	ReportLastRunStatus_FAILURE ReportLastRunStatus_RunStatus = 1
)

// Enum value maps for ReportLastRunStatus_RunStatus.
var (
	ReportLastRunStatus_RunStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
	}
	ReportLastRunStatus_RunStatus_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
	}
)

func (x ReportLastRunStatus_RunStatus) Enum() *ReportLastRunStatus_RunStatus {
	p := new(ReportLastRunStatus_RunStatus)
	*p = x
	return p
}

func (x ReportLastRunStatus_RunStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReportLastRunStatus_RunStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_configuration_proto_enumTypes[1].Descriptor()
}

func (ReportLastRunStatus_RunStatus) Type() protoreflect.EnumType {
	return &file_storage_report_configuration_proto_enumTypes[1]
}

func (x ReportLastRunStatus_RunStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReportLastRunStatus_RunStatus.Descriptor instead.
func (ReportLastRunStatus_RunStatus) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{1, 0}
}

type VulnerabilityReportFilters_Fixability int32

const (
	VulnerabilityReportFilters_BOTH        VulnerabilityReportFilters_Fixability = 0
	VulnerabilityReportFilters_FIXABLE     VulnerabilityReportFilters_Fixability = 1
	VulnerabilityReportFilters_NOT_FIXABLE VulnerabilityReportFilters_Fixability = 2
)

// Enum value maps for VulnerabilityReportFilters_Fixability.
var (
	VulnerabilityReportFilters_Fixability_name = map[int32]string{
		0: "BOTH",
		1: "FIXABLE",
		2: "NOT_FIXABLE",
	}
	VulnerabilityReportFilters_Fixability_value = map[string]int32{
		"BOTH":        0,
		"FIXABLE":     1,
		"NOT_FIXABLE": 2,
	}
)

func (x VulnerabilityReportFilters_Fixability) Enum() *VulnerabilityReportFilters_Fixability {
	p := new(VulnerabilityReportFilters_Fixability)
	*p = x
	return p
}

func (x VulnerabilityReportFilters_Fixability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VulnerabilityReportFilters_Fixability) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_configuration_proto_enumTypes[2].Descriptor()
}

func (VulnerabilityReportFilters_Fixability) Type() protoreflect.EnumType {
	return &file_storage_report_configuration_proto_enumTypes[2]
}

func (x VulnerabilityReportFilters_Fixability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VulnerabilityReportFilters_Fixability.Descriptor instead.
func (VulnerabilityReportFilters_Fixability) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{2, 0}
}

type VulnerabilityReportFilters_ImageType int32

const (
	VulnerabilityReportFilters_DEPLOYED VulnerabilityReportFilters_ImageType = 0
	VulnerabilityReportFilters_WATCHED  VulnerabilityReportFilters_ImageType = 1
)

// Enum value maps for VulnerabilityReportFilters_ImageType.
var (
	VulnerabilityReportFilters_ImageType_name = map[int32]string{
		0: "DEPLOYED",
		1: "WATCHED",
	}
	VulnerabilityReportFilters_ImageType_value = map[string]int32{
		"DEPLOYED": 0,
		"WATCHED":  1,
	}
)

func (x VulnerabilityReportFilters_ImageType) Enum() *VulnerabilityReportFilters_ImageType {
	p := new(VulnerabilityReportFilters_ImageType)
	*p = x
	return p
}

func (x VulnerabilityReportFilters_ImageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VulnerabilityReportFilters_ImageType) Descriptor() protoreflect.EnumDescriptor {
	return file_storage_report_configuration_proto_enumTypes[3].Descriptor()
}

func (VulnerabilityReportFilters_ImageType) Type() protoreflect.EnumType {
	return &file_storage_report_configuration_proto_enumTypes[3]
}

func (x VulnerabilityReportFilters_ImageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VulnerabilityReportFilters_ImageType.Descriptor instead.
func (VulnerabilityReportFilters_ImageType) EnumDescriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{2, 1}
}

type ReportConfiguration struct {
	state       protoimpl.MessageState         `protogen:"open.v1"`
	Id          string                         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" sql:"pk"`     // @gotags: sql:"pk"
	Name        string                         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" search:"Report Name"` // @gotags: search:"Report Name"
	Description string                         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        ReportConfiguration_ReportType `protobuf:"varint,4,opt,name=type,proto3,enum=storage.ReportConfiguration_ReportType" json:"type,omitempty" search:"Report Type"` // @gotags: search:"Report Type"
	// Types that are valid to be assigned to Filter:
	//
	//	*ReportConfiguration_VulnReportFilters
	Filter  isReportConfiguration_Filter `protobuf_oneof:"filter"`
	ScopeId string                       `protobuf:"bytes,6,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" search:"Embedded Collection ID"` // @gotags: search:"Embedded Collection ID"
	// Types that are valid to be assigned to NotifierConfig:
	//
	//	*ReportConfiguration_EmailConfig
	NotifierConfig        isReportConfiguration_NotifierConfig `protobuf_oneof:"notifier_config"`
	Schedule              *Schedule                            `protobuf:"bytes,8,opt,name=schedule,proto3" json:"schedule,omitempty"`
	LastRunStatus         *ReportLastRunStatus                 `protobuf:"bytes,9,opt,name=last_run_status,json=lastRunStatus,proto3" json:"last_run_status,omitempty"`
	LastSuccessfulRunTime *timestamppb.Timestamp               `protobuf:"bytes,10,opt,name=last_successful_run_time,json=lastSuccessfulRunTime,proto3" json:"last_successful_run_time,omitempty"`
	ResourceScope         *ResourceScope                       `protobuf:"bytes,11,opt,name=resource_scope,json=resourceScope,proto3" json:"resource_scope,omitempty"`
	Notifiers             []*NotifierConfiguration             `protobuf:"bytes,12,rep,name=notifiers,proto3" json:"notifiers,omitempty"`
	Creator               *SlimUser                            `protobuf:"bytes,13,opt,name=creator,proto3" json:"creator,omitempty" sql:"ignore_labels(User ID)"`  // @gotags: sql:"ignore_labels(User ID)"
	Version               int32                                `protobuf:"varint,14,opt,name=version,proto3" json:"version,omitempty"` // version=0 is unmigrated v1 config, version=1 is migrated v1 config and version=2 is v2 config
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ReportConfiguration) Reset() {
	*x = ReportConfiguration{}
	mi := &file_storage_report_configuration_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportConfiguration) ProtoMessage() {}

func (x *ReportConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_configuration_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportConfiguration.ProtoReflect.Descriptor instead.
func (*ReportConfiguration) Descriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *ReportConfiguration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReportConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReportConfiguration) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReportConfiguration) GetType() ReportConfiguration_ReportType {
	if x != nil {
		return x.Type
	}
	return ReportConfiguration_VULNERABILITY
}

func (x *ReportConfiguration) GetFilter() isReportConfiguration_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ReportConfiguration) GetVulnReportFilters() *VulnerabilityReportFilters {
	if x != nil {
		if x, ok := x.Filter.(*ReportConfiguration_VulnReportFilters); ok {
			return x.VulnReportFilters
		}
	}
	return nil
}

func (x *ReportConfiguration) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *ReportConfiguration) GetNotifierConfig() isReportConfiguration_NotifierConfig {
	if x != nil {
		return x.NotifierConfig
	}
	return nil
}

func (x *ReportConfiguration) GetEmailConfig() *EmailNotifierConfiguration {
	if x != nil {
		if x, ok := x.NotifierConfig.(*ReportConfiguration_EmailConfig); ok {
			return x.EmailConfig
		}
	}
	return nil
}

func (x *ReportConfiguration) GetSchedule() *Schedule {
	if x != nil {
		return x.Schedule
	}
	return nil
}

func (x *ReportConfiguration) GetLastRunStatus() *ReportLastRunStatus {
	if x != nil {
		return x.LastRunStatus
	}
	return nil
}

func (x *ReportConfiguration) GetLastSuccessfulRunTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSuccessfulRunTime
	}
	return nil
}

func (x *ReportConfiguration) GetResourceScope() *ResourceScope {
	if x != nil {
		return x.ResourceScope
	}
	return nil
}

func (x *ReportConfiguration) GetNotifiers() []*NotifierConfiguration {
	if x != nil {
		return x.Notifiers
	}
	return nil
}

func (x *ReportConfiguration) GetCreator() *SlimUser {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *ReportConfiguration) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type isReportConfiguration_Filter interface {
	isReportConfiguration_Filter()
}

type ReportConfiguration_VulnReportFilters struct {
	VulnReportFilters *VulnerabilityReportFilters `protobuf:"bytes,5,opt,name=vuln_report_filters,json=vulnReportFilters,proto3,oneof"`
}

func (*ReportConfiguration_VulnReportFilters) isReportConfiguration_Filter() {}

type isReportConfiguration_NotifierConfig interface {
	isReportConfiguration_NotifierConfig()
}

type ReportConfiguration_EmailConfig struct {
	EmailConfig *EmailNotifierConfiguration `protobuf:"bytes,7,opt,name=email_config,json=emailConfig,proto3,oneof"`
}

func (*ReportConfiguration_EmailConfig) isReportConfiguration_NotifierConfig() {}

type ReportLastRunStatus struct {
	state         protoimpl.MessageState        `protogen:"open.v1"`
	ReportStatus  ReportLastRunStatus_RunStatus `protobuf:"varint,1,opt,name=report_status,json=reportStatus,proto3,enum=storage.ReportLastRunStatus_RunStatus" json:"report_status,omitempty"`
	LastRunTime   *timestamppb.Timestamp        `protobuf:"bytes,2,opt,name=last_run_time,json=lastRunTime,proto3" json:"last_run_time,omitempty"`
	ErrorMsg      string                        `protobuf:"bytes,3,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReportLastRunStatus) Reset() {
	*x = ReportLastRunStatus{}
	mi := &file_storage_report_configuration_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportLastRunStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportLastRunStatus) ProtoMessage() {}

func (x *ReportLastRunStatus) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_configuration_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportLastRunStatus.ProtoReflect.Descriptor instead.
func (*ReportLastRunStatus) Descriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *ReportLastRunStatus) GetReportStatus() ReportLastRunStatus_RunStatus {
	if x != nil {
		return x.ReportStatus
	}
	return ReportLastRunStatus_SUCCESS
}

func (x *ReportLastRunStatus) GetLastRunTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastRunTime
	}
	return nil
}

func (x *ReportLastRunStatus) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

type VulnerabilityReportFilters struct {
	state           protoimpl.MessageState                 `protogen:"open.v1"`
	Fixability      VulnerabilityReportFilters_Fixability  `protobuf:"varint,1,opt,name=fixability,proto3,enum=storage.VulnerabilityReportFilters_Fixability" json:"fixability,omitempty"`
	SinceLastReport bool                                   `protobuf:"varint,2,opt,name=since_last_report,json=sinceLastReport,proto3" json:"since_last_report,omitempty"`
	Severities      []VulnerabilitySeverity                `protobuf:"varint,3,rep,packed,name=severities,proto3,enum=storage.VulnerabilitySeverity" json:"severities,omitempty"`
	ImageTypes      []VulnerabilityReportFilters_ImageType `protobuf:"varint,4,rep,packed,name=image_types,json=imageTypes,proto3,enum=storage.VulnerabilityReportFilters_ImageType" json:"image_types,omitempty"`
	// Types that are valid to be assigned to CvesSince:
	//
	//	*VulnerabilityReportFilters_AllVuln
	//	*VulnerabilityReportFilters_SinceLastSentScheduledReport
	//	*VulnerabilityReportFilters_SinceStartDate
	CvesSince        isVulnerabilityReportFilters_CvesSince `protobuf_oneof:"cves_since"`
	AccessScopeRules []*SimpleAccessScope_Rules             `protobuf:"bytes,8,rep,name=access_scope_rules,json=accessScopeRules,proto3" json:"access_scope_rules,omitempty"`
	IncludeNvdCvss   bool                                   `protobuf:"varint,9,opt,name=include_nvd_cvss,json=includeNvdCvss,proto3" json:"include_nvd_cvss,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *VulnerabilityReportFilters) Reset() {
	*x = VulnerabilityReportFilters{}
	mi := &file_storage_report_configuration_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VulnerabilityReportFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityReportFilters) ProtoMessage() {}

func (x *VulnerabilityReportFilters) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_configuration_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityReportFilters.ProtoReflect.Descriptor instead.
func (*VulnerabilityReportFilters) Descriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *VulnerabilityReportFilters) GetFixability() VulnerabilityReportFilters_Fixability {
	if x != nil {
		return x.Fixability
	}
	return VulnerabilityReportFilters_BOTH
}

func (x *VulnerabilityReportFilters) GetSinceLastReport() bool {
	if x != nil {
		return x.SinceLastReport
	}
	return false
}

func (x *VulnerabilityReportFilters) GetSeverities() []VulnerabilitySeverity {
	if x != nil {
		return x.Severities
	}
	return nil
}

func (x *VulnerabilityReportFilters) GetImageTypes() []VulnerabilityReportFilters_ImageType {
	if x != nil {
		return x.ImageTypes
	}
	return nil
}

func (x *VulnerabilityReportFilters) GetCvesSince() isVulnerabilityReportFilters_CvesSince {
	if x != nil {
		return x.CvesSince
	}
	return nil
}

func (x *VulnerabilityReportFilters) GetAllVuln() bool {
	if x != nil {
		if x, ok := x.CvesSince.(*VulnerabilityReportFilters_AllVuln); ok {
			return x.AllVuln
		}
	}
	return false
}

func (x *VulnerabilityReportFilters) GetSinceLastSentScheduledReport() bool {
	if x != nil {
		if x, ok := x.CvesSince.(*VulnerabilityReportFilters_SinceLastSentScheduledReport); ok {
			return x.SinceLastSentScheduledReport
		}
	}
	return false
}

func (x *VulnerabilityReportFilters) GetSinceStartDate() *timestamppb.Timestamp {
	if x != nil {
		if x, ok := x.CvesSince.(*VulnerabilityReportFilters_SinceStartDate); ok {
			return x.SinceStartDate
		}
	}
	return nil
}

func (x *VulnerabilityReportFilters) GetAccessScopeRules() []*SimpleAccessScope_Rules {
	if x != nil {
		return x.AccessScopeRules
	}
	return nil
}

func (x *VulnerabilityReportFilters) GetIncludeNvdCvss() bool {
	if x != nil {
		return x.IncludeNvdCvss
	}
	return false
}

type isVulnerabilityReportFilters_CvesSince interface {
	isVulnerabilityReportFilters_CvesSince()
}

type VulnerabilityReportFilters_AllVuln struct {
	AllVuln bool `protobuf:"varint,5,opt,name=all_vuln,json=allVuln,proto3,oneof"`
}

type VulnerabilityReportFilters_SinceLastSentScheduledReport struct {
	SinceLastSentScheduledReport bool `protobuf:"varint,6,opt,name=since_last_sent_scheduled_report,json=sinceLastSentScheduledReport,proto3,oneof"`
}

type VulnerabilityReportFilters_SinceStartDate struct {
	SinceStartDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=since_start_date,json=sinceStartDate,proto3,oneof"`
}

func (*VulnerabilityReportFilters_AllVuln) isVulnerabilityReportFilters_CvesSince() {}

func (*VulnerabilityReportFilters_SinceLastSentScheduledReport) isVulnerabilityReportFilters_CvesSince() {
}

func (*VulnerabilityReportFilters_SinceStartDate) isVulnerabilityReportFilters_CvesSince() {}

type ResourceScope struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to ScopeReference:
	//
	//	*ResourceScope_CollectionId
	ScopeReference isResourceScope_ScopeReference `protobuf_oneof:"scope_reference"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ResourceScope) Reset() {
	*x = ResourceScope{}
	mi := &file_storage_report_configuration_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceScope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceScope) ProtoMessage() {}

func (x *ResourceScope) ProtoReflect() protoreflect.Message {
	mi := &file_storage_report_configuration_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceScope.ProtoReflect.Descriptor instead.
func (*ResourceScope) Descriptor() ([]byte, []int) {
	return file_storage_report_configuration_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceScope) GetScopeReference() isResourceScope_ScopeReference {
	if x != nil {
		return x.ScopeReference
	}
	return nil
}

func (x *ResourceScope) GetCollectionId() string {
	if x != nil {
		if x, ok := x.ScopeReference.(*ResourceScope_CollectionId); ok {
			return x.CollectionId
		}
	}
	return ""
}

type isResourceScope_ScopeReference interface {
	isResourceScope_ScopeReference()
}

type ResourceScope_CollectionId struct {
	CollectionId string `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3,oneof" search:"Collection ID"` // @gotags: search:"Collection ID"
}

func (*ResourceScope_CollectionId) isResourceScope_ScopeReference() {}

var File_storage_report_configuration_proto protoreflect.FileDescriptor

var file_storage_report_configuration_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0,
	0x06, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x55, 0x0a, 0x13, 0x76, 0x75, 0x6c,
	0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x11, 0x76,
	0x75, 0x6c, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x0c, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x75, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x61,
	0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x53, 0x0a, 0x18, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x5f, 0x72, 0x75,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x3d, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x3c,
	0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x6c, 0x69, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x1f, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x55, 0x4c, 0x4e, 0x45, 0x52, 0x41, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x10, 0x00, 0x42, 0x08, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x11,
	0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0xe6, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x61, 0x73, 0x74,
	0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4b, 0x0a, 0x0d, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x26, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52,
	0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72,
	0x75, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52,
	0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x73, 0x67, 0x22, 0x25, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x22, 0xbd, 0x05, 0x0a, 0x1a, 0x56,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x66, 0x69, 0x78,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x46, 0x69, 0x78, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x66,
	0x69, 0x78, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x69, 0x6e,
	0x63, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x08, 0x61, 0x6c, 0x6c, 0x5f, 0x76, 0x75, 0x6c,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x61, 0x6c, 0x6c, 0x56, 0x75,
	0x6c, 0x6e, 0x12, 0x48, 0x0a, 0x20, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x1c,
	0x73, 0x69, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x46, 0x0a, 0x10,
	0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x4e, 0x0a, 0x12, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x52, 0x10, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x6e, 0x76, 0x64, 0x5f, 0x63, 0x76, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4e, 0x76, 0x64, 0x43, 0x76, 0x73, 0x73, 0x22, 0x34,
	0x0a, 0x0a, 0x46, 0x69, 0x78, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x08, 0x0a, 0x04,
	0x42, 0x4f, 0x54, 0x48, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x49, 0x58, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x49, 0x58, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x02, 0x22, 0x26, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x57, 0x41, 0x54, 0x43, 0x48, 0x45, 0x44, 0x10, 0x01, 0x42, 0x0c, 0x0a, 0x0a,
	0x63, 0x76, 0x65, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x49, 0x0a, 0x0d, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x42, 0x11, 0x0a, 0x0f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x2e, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5a, 0x11, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x3b, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_report_configuration_proto_rawDescOnce sync.Once
	file_storage_report_configuration_proto_rawDescData = file_storage_report_configuration_proto_rawDesc
)

func file_storage_report_configuration_proto_rawDescGZIP() []byte {
	file_storage_report_configuration_proto_rawDescOnce.Do(func() {
		file_storage_report_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_report_configuration_proto_rawDescData)
	})
	return file_storage_report_configuration_proto_rawDescData
}

var file_storage_report_configuration_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_storage_report_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_storage_report_configuration_proto_goTypes = []any{
	(ReportConfiguration_ReportType)(0),        // 0: storage.ReportConfiguration.ReportType
	(ReportLastRunStatus_RunStatus)(0),         // 1: storage.ReportLastRunStatus.RunStatus
	(VulnerabilityReportFilters_Fixability)(0), // 2: storage.VulnerabilityReportFilters.Fixability
	(VulnerabilityReportFilters_ImageType)(0),  // 3: storage.VulnerabilityReportFilters.ImageType
	(*ReportConfiguration)(nil),                // 4: storage.ReportConfiguration
	(*ReportLastRunStatus)(nil),                // 5: storage.ReportLastRunStatus
	(*VulnerabilityReportFilters)(nil),         // 6: storage.VulnerabilityReportFilters
	(*ResourceScope)(nil),                      // 7: storage.ResourceScope
	(*EmailNotifierConfiguration)(nil),         // 8: storage.EmailNotifierConfiguration
	(*Schedule)(nil),                           // 9: storage.Schedule
	(*timestamppb.Timestamp)(nil),              // 10: google.protobuf.Timestamp
	(*NotifierConfiguration)(nil),              // 11: storage.NotifierConfiguration
	(*SlimUser)(nil),                           // 12: storage.SlimUser
	(VulnerabilitySeverity)(0),                 // 13: storage.VulnerabilitySeverity
	(*SimpleAccessScope_Rules)(nil),            // 14: storage.SimpleAccessScope.Rules
}
var file_storage_report_configuration_proto_depIdxs = []int32{
	0,  // 0: storage.ReportConfiguration.type:type_name -> storage.ReportConfiguration.ReportType
	6,  // 1: storage.ReportConfiguration.vuln_report_filters:type_name -> storage.VulnerabilityReportFilters
	8,  // 2: storage.ReportConfiguration.email_config:type_name -> storage.EmailNotifierConfiguration
	9,  // 3: storage.ReportConfiguration.schedule:type_name -> storage.Schedule
	5,  // 4: storage.ReportConfiguration.last_run_status:type_name -> storage.ReportLastRunStatus
	10, // 5: storage.ReportConfiguration.last_successful_run_time:type_name -> google.protobuf.Timestamp
	7,  // 6: storage.ReportConfiguration.resource_scope:type_name -> storage.ResourceScope
	11, // 7: storage.ReportConfiguration.notifiers:type_name -> storage.NotifierConfiguration
	12, // 8: storage.ReportConfiguration.creator:type_name -> storage.SlimUser
	1,  // 9: storage.ReportLastRunStatus.report_status:type_name -> storage.ReportLastRunStatus.RunStatus
	10, // 10: storage.ReportLastRunStatus.last_run_time:type_name -> google.protobuf.Timestamp
	2,  // 11: storage.VulnerabilityReportFilters.fixability:type_name -> storage.VulnerabilityReportFilters.Fixability
	13, // 12: storage.VulnerabilityReportFilters.severities:type_name -> storage.VulnerabilitySeverity
	3,  // 13: storage.VulnerabilityReportFilters.image_types:type_name -> storage.VulnerabilityReportFilters.ImageType
	10, // 14: storage.VulnerabilityReportFilters.since_start_date:type_name -> google.protobuf.Timestamp
	14, // 15: storage.VulnerabilityReportFilters.access_scope_rules:type_name -> storage.SimpleAccessScope.Rules
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_storage_report_configuration_proto_init() }
func file_storage_report_configuration_proto_init() {
	if File_storage_report_configuration_proto != nil {
		return
	}
	file_storage_cve_proto_init()
	file_storage_report_notifier_configuration_proto_init()
	file_storage_role_proto_init()
	file_storage_schedule_proto_init()
	file_storage_user_proto_init()
	file_storage_report_configuration_proto_msgTypes[0].OneofWrappers = []any{
		(*ReportConfiguration_VulnReportFilters)(nil),
		(*ReportConfiguration_EmailConfig)(nil),
	}
	file_storage_report_configuration_proto_msgTypes[2].OneofWrappers = []any{
		(*VulnerabilityReportFilters_AllVuln)(nil),
		(*VulnerabilityReportFilters_SinceLastSentScheduledReport)(nil),
		(*VulnerabilityReportFilters_SinceStartDate)(nil),
	}
	file_storage_report_configuration_proto_msgTypes[3].OneofWrappers = []any{
		(*ResourceScope_CollectionId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_report_configuration_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_report_configuration_proto_goTypes,
		DependencyIndexes: file_storage_report_configuration_proto_depIdxs,
		EnumInfos:         file_storage_report_configuration_proto_enumTypes,
		MessageInfos:      file_storage_report_configuration_proto_msgTypes,
	}.Build()
	File_storage_report_configuration_proto = out.File
	file_storage_report_configuration_proto_rawDesc = nil
	file_storage_report_configuration_proto_goTypes = nil
	file_storage_report_configuration_proto_depIdxs = nil
}
