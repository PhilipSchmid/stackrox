// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/v2/compliance_integration_service.proto

package v2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ComplianceIntegrationService_ListComplianceIntegrations_FullMethodName     = "/v2.ComplianceIntegrationService/ListComplianceIntegrations"
	ComplianceIntegrationService_GetComplianceIntegrationsCount_FullMethodName = "/v2.ComplianceIntegrationService/GetComplianceIntegrationsCount"
)

// ComplianceIntegrationServiceClient is the client API for ComplianceIntegrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComplianceIntegrationServiceClient interface {
	// ListComplianceIntegrations lists all the compliance operator metadata for the secured clusters
	ListComplianceIntegrations(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListComplianceIntegrationsResponse, error)
	// GetComplianceIntegrationsCount returns the number of compliance operator integrations
	// matching the given query
	GetComplianceIntegrationsCount(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*CountComplianceIntegrationsResponse, error)
}

type complianceIntegrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComplianceIntegrationServiceClient(cc grpc.ClientConnInterface) ComplianceIntegrationServiceClient {
	return &complianceIntegrationServiceClient{cc}
}

func (c *complianceIntegrationServiceClient) ListComplianceIntegrations(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*ListComplianceIntegrationsResponse, error) {
	out := new(ListComplianceIntegrationsResponse)
	err := c.cc.Invoke(ctx, ComplianceIntegrationService_ListComplianceIntegrations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceIntegrationServiceClient) GetComplianceIntegrationsCount(ctx context.Context, in *RawQuery, opts ...grpc.CallOption) (*CountComplianceIntegrationsResponse, error) {
	out := new(CountComplianceIntegrationsResponse)
	err := c.cc.Invoke(ctx, ComplianceIntegrationService_GetComplianceIntegrationsCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplianceIntegrationServiceServer is the server API for ComplianceIntegrationService service.
// All implementations should embed UnimplementedComplianceIntegrationServiceServer
// for forward compatibility
type ComplianceIntegrationServiceServer interface {
	// ListComplianceIntegrations lists all the compliance operator metadata for the secured clusters
	ListComplianceIntegrations(context.Context, *RawQuery) (*ListComplianceIntegrationsResponse, error)
	// GetComplianceIntegrationsCount returns the number of compliance operator integrations
	// matching the given query
	GetComplianceIntegrationsCount(context.Context, *RawQuery) (*CountComplianceIntegrationsResponse, error)
}

// UnimplementedComplianceIntegrationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedComplianceIntegrationServiceServer struct {
}

func (UnimplementedComplianceIntegrationServiceServer) ListComplianceIntegrations(context.Context, *RawQuery) (*ListComplianceIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComplianceIntegrations not implemented")
}
func (UnimplementedComplianceIntegrationServiceServer) GetComplianceIntegrationsCount(context.Context, *RawQuery) (*CountComplianceIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComplianceIntegrationsCount not implemented")
}

// UnsafeComplianceIntegrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComplianceIntegrationServiceServer will
// result in compilation errors.
type UnsafeComplianceIntegrationServiceServer interface {
	mustEmbedUnimplementedComplianceIntegrationServiceServer()
}

func RegisterComplianceIntegrationServiceServer(s grpc.ServiceRegistrar, srv ComplianceIntegrationServiceServer) {
	s.RegisterService(&ComplianceIntegrationService_ServiceDesc, srv)
}

func _ComplianceIntegrationService_ListComplianceIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceIntegrationServiceServer).ListComplianceIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceIntegrationService_ListComplianceIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceIntegrationServiceServer).ListComplianceIntegrations(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceIntegrationService_GetComplianceIntegrationsCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceIntegrationServiceServer).GetComplianceIntegrationsCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComplianceIntegrationService_GetComplianceIntegrationsCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceIntegrationServiceServer).GetComplianceIntegrationsCount(ctx, req.(*RawQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// ComplianceIntegrationService_ServiceDesc is the grpc.ServiceDesc for ComplianceIntegrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComplianceIntegrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v2.ComplianceIntegrationService",
	HandlerType: (*ComplianceIntegrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListComplianceIntegrations",
			Handler:    _ComplianceIntegrationService_ListComplianceIntegrations_Handler,
		},
		{
			MethodName: "GetComplianceIntegrationsCount",
			Handler:    _ComplianceIntegrationService_GetComplianceIntegrationsCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v2/compliance_integration_service.proto",
}