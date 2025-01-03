// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.3
// source: internalapi/sensor/network_connection_iservice.proto

package sensor

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

type NetworkConnectionInfoMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Msg:
	//
	//	*NetworkConnectionInfoMessage_Register
	//	*NetworkConnectionInfoMessage_Info
	Msg           isNetworkConnectionInfoMessage_Msg `protobuf_oneof:"msg"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkConnectionInfoMessage) Reset() {
	*x = NetworkConnectionInfoMessage{}
	mi := &file_internalapi_sensor_network_connection_iservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkConnectionInfoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkConnectionInfoMessage) ProtoMessage() {}

func (x *NetworkConnectionInfoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_network_connection_iservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkConnectionInfoMessage.ProtoReflect.Descriptor instead.
func (*NetworkConnectionInfoMessage) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_network_connection_iservice_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkConnectionInfoMessage) GetMsg() isNetworkConnectionInfoMessage_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *NetworkConnectionInfoMessage) GetRegister() *CollectorRegisterRequest {
	if x != nil {
		if x, ok := x.Msg.(*NetworkConnectionInfoMessage_Register); ok {
			return x.Register
		}
	}
	return nil
}

func (x *NetworkConnectionInfoMessage) GetInfo() *NetworkConnectionInfo {
	if x != nil {
		if x, ok := x.Msg.(*NetworkConnectionInfoMessage_Info); ok {
			return x.Info
		}
	}
	return nil
}

type isNetworkConnectionInfoMessage_Msg interface {
	isNetworkConnectionInfoMessage_Msg()
}

type NetworkConnectionInfoMessage_Register struct {
	Register *CollectorRegisterRequest `protobuf:"bytes,1,opt,name=register,proto3,oneof"`
}

type NetworkConnectionInfoMessage_Info struct {
	Info *NetworkConnectionInfo `protobuf:"bytes,2,opt,name=info,proto3,oneof"`
}

func (*NetworkConnectionInfoMessage_Register) isNetworkConnectionInfoMessage_Msg() {}

func (*NetworkConnectionInfoMessage_Info) isNetworkConnectionInfoMessage_Msg() {}

type NetworkFlowsControlMessage struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	PublicIpAddresses *IPAddressList         `protobuf:"bytes,1,opt,name=public_ip_addresses,json=publicIpAddresses,proto3" json:"public_ip_addresses,omitempty"`
	IpNetworks        *IPNetworkList         `protobuf:"bytes,2,opt,name=ip_networks,json=ipNetworks,proto3" json:"ip_networks,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *NetworkFlowsControlMessage) Reset() {
	*x = NetworkFlowsControlMessage{}
	mi := &file_internalapi_sensor_network_connection_iservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkFlowsControlMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkFlowsControlMessage) ProtoMessage() {}

func (x *NetworkFlowsControlMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_network_connection_iservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkFlowsControlMessage.ProtoReflect.Descriptor instead.
func (*NetworkFlowsControlMessage) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_network_connection_iservice_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkFlowsControlMessage) GetPublicIpAddresses() *IPAddressList {
	if x != nil {
		return x.PublicIpAddresses
	}
	return nil
}

func (x *NetworkFlowsControlMessage) GetIpNetworks() *IPNetworkList {
	if x != nil {
		return x.IpNetworks
	}
	return nil
}

type IPAddressList struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A list of IPv4 addresses, encoded in correct byte order (i.e., 127.0.0.1 is encoded as 0x7f000001). On a
	// little-endian system, you must convert to big endian to interpret the value as a binary address.
	Ipv4Addresses []uint32 `protobuf:"fixed32,1,rep,packed,name=ipv4_addresses,json=ipv4Addresses,proto3" json:"ipv4_addresses,omitempty"`
	// A list of IPv6 addresses, as uint64 pairs. Each pair is ordered in network order (big endian, i.e., first high,
	// then low); each uint64 is encoded in the correct byte order and may need to be converted to big endian on little
	// endian System. E.g., the IPv6 address ::1 (local loopback) is encoded as a `0` uint64 followed by a `1` uint64.
	// This field must always have an even number of values; otherwise it should be discarded.
	Ipv6Addresses []uint64 `protobuf:"fixed64,2,rep,packed,name=ipv6_addresses,json=ipv6Addresses,proto3" json:"ipv6_addresses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IPAddressList) Reset() {
	*x = IPAddressList{}
	mi := &file_internalapi_sensor_network_connection_iservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPAddressList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPAddressList) ProtoMessage() {}

func (x *IPAddressList) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_network_connection_iservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPAddressList.ProtoReflect.Descriptor instead.
func (*IPAddressList) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_network_connection_iservice_proto_rawDescGZIP(), []int{2}
}

func (x *IPAddressList) GetIpv4Addresses() []uint32 {
	if x != nil {
		return x.Ipv4Addresses
	}
	return nil
}

func (x *IPAddressList) GetIpv6Addresses() []uint64 {
	if x != nil {
		return x.Ipv6Addresses
	}
	return nil
}

type IPNetworkList struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// A list of IPv4 networks, as a sequence of bytes. First four bytes represent IP followed by the byte representing
	// the network prefix length. Each sequence encoded in correct byte order (i.e., 127.0.0.1/8 is encoded as 0x7f0000018).
	// On a little-endian system, you must convert to big endian to interpret the value as a binary address.
	// This field must always have an 5x number of values; otherwise it should be discarded.
	Ipv4Networks []byte `protobuf:"bytes,1,opt,name=ipv4_networks,json=ipv4Networks,proto3" json:"ipv4_networks,omitempty"`
	// A list of IPv6 networks, as a sequence of bytes. First 16 bytes representing IP followed by the byte representing
	// the network prefix length. Each IP sequence is ordered in network order (big endian, i.e., first high, then low).
	// Each order is encoded in the correct byte order and may need to be converted to big endian on little endian System.
	// This field must always have an 17x number of values; otherwise it should be discarded.
	Ipv6Networks  []byte `protobuf:"bytes,2,opt,name=ipv6_networks,json=ipv6Networks,proto3" json:"ipv6_networks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IPNetworkList) Reset() {
	*x = IPNetworkList{}
	mi := &file_internalapi_sensor_network_connection_iservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPNetworkList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPNetworkList) ProtoMessage() {}

func (x *IPNetworkList) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_sensor_network_connection_iservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPNetworkList.ProtoReflect.Descriptor instead.
func (*IPNetworkList) Descriptor() ([]byte, []int) {
	return file_internalapi_sensor_network_connection_iservice_proto_rawDescGZIP(), []int{3}
}

func (x *IPNetworkList) GetIpv4Networks() []byte {
	if x != nil {
		return x.Ipv4Networks
	}
	return nil
}

func (x *IPNetworkList) GetIpv6Networks() []byte {
	if x != nil {
		return x.Ipv6Networks
	}
	return nil
}

var File_internalapi_sensor_network_connection_iservice_proto protoreflect.FileDescriptor

var file_internalapi_sensor_network_connection_iservice_proto_rawDesc = []byte{
	0x0a, 0x34, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x1a, 0x22,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x30, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x1c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x9b, 0x01, 0x0a, 0x1a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x6c, 0x6f,
	0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x45, 0x0a, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x69, 0x70, 0x5f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x49, 0x50, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x0a, 0x69, 0x70, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x22,
	0x5d, 0x0a, 0x0d, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x07, 0x52, 0x0d, 0x69, 0x70, 0x76, 0x34, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x70, 0x76, 0x36, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x06, 0x52,
	0x0d, 0x69, 0x70, 0x76, 0x36, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x59,
	0x0a, 0x0d, 0x49, 0x50, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x69, 0x70, 0x76, 0x34, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x70, 0x76, 0x36, 0x5f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x69, 0x70, 0x76,
	0x36, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x32, 0x89, 0x01, 0x0a, 0x1c, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x19, 0x50, 0x75,
	0x73, 0x68, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x22, 0x2e,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x6c,
	0x6f, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x20, 0x5a, 0x1b, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x3b, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internalapi_sensor_network_connection_iservice_proto_rawDescOnce sync.Once
	file_internalapi_sensor_network_connection_iservice_proto_rawDescData = file_internalapi_sensor_network_connection_iservice_proto_rawDesc
)

func file_internalapi_sensor_network_connection_iservice_proto_rawDescGZIP() []byte {
	file_internalapi_sensor_network_connection_iservice_proto_rawDescOnce.Do(func() {
		file_internalapi_sensor_network_connection_iservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_sensor_network_connection_iservice_proto_rawDescData)
	})
	return file_internalapi_sensor_network_connection_iservice_proto_rawDescData
}

var file_internalapi_sensor_network_connection_iservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internalapi_sensor_network_connection_iservice_proto_goTypes = []any{
	(*NetworkConnectionInfoMessage)(nil), // 0: sensor.NetworkConnectionInfoMessage
	(*NetworkFlowsControlMessage)(nil),   // 1: sensor.NetworkFlowsControlMessage
	(*IPAddressList)(nil),                // 2: sensor.IPAddressList
	(*IPNetworkList)(nil),                // 3: sensor.IPNetworkList
	(*CollectorRegisterRequest)(nil),     // 4: sensor.CollectorRegisterRequest
	(*NetworkConnectionInfo)(nil),        // 5: sensor.NetworkConnectionInfo
}
var file_internalapi_sensor_network_connection_iservice_proto_depIdxs = []int32{
	4, // 0: sensor.NetworkConnectionInfoMessage.register:type_name -> sensor.CollectorRegisterRequest
	5, // 1: sensor.NetworkConnectionInfoMessage.info:type_name -> sensor.NetworkConnectionInfo
	2, // 2: sensor.NetworkFlowsControlMessage.public_ip_addresses:type_name -> sensor.IPAddressList
	3, // 3: sensor.NetworkFlowsControlMessage.ip_networks:type_name -> sensor.IPNetworkList
	0, // 4: sensor.NetworkConnectionInfoService.PushNetworkConnectionInfo:input_type -> sensor.NetworkConnectionInfoMessage
	1, // 5: sensor.NetworkConnectionInfoService.PushNetworkConnectionInfo:output_type -> sensor.NetworkFlowsControlMessage
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internalapi_sensor_network_connection_iservice_proto_init() }
func file_internalapi_sensor_network_connection_iservice_proto_init() {
	if File_internalapi_sensor_network_connection_iservice_proto != nil {
		return
	}
	file_internalapi_sensor_collector_proto_init()
	file_internalapi_sensor_network_connection_info_proto_init()
	file_internalapi_sensor_network_connection_iservice_proto_msgTypes[0].OneofWrappers = []any{
		(*NetworkConnectionInfoMessage_Register)(nil),
		(*NetworkConnectionInfoMessage_Info)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internalapi_sensor_network_connection_iservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_sensor_network_connection_iservice_proto_goTypes,
		DependencyIndexes: file_internalapi_sensor_network_connection_iservice_proto_depIdxs,
		MessageInfos:      file_internalapi_sensor_network_connection_iservice_proto_msgTypes,
	}.Build()
	File_internalapi_sensor_network_connection_iservice_proto = out.File
	file_internalapi_sensor_network_connection_iservice_proto_rawDesc = nil
	file_internalapi_sensor_network_connection_iservice_proto_goTypes = nil
	file_internalapi_sensor_network_connection_iservice_proto_depIdxs = nil
}
