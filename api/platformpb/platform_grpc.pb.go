// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: platformpb/platform.proto

package platformpb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Platform_Connect_FullMethodName                        = "/platform.Platform/Connect"
	Platform_Disconnect_FullMethodName                     = "/platform.Platform/Disconnect"
	Platform_SearchOrganizationTickets_FullMethodName      = "/platform.Platform/SearchOrganizationTickets"
	Platform_SearchOrganizationEntitlements_FullMethodName = "/platform.Platform/SearchOrganizationEntitlements"
	Platform_GetContactInformation_FullMethodName          = "/platform.Platform/GetContactInformation"
	Platform_ServerInfo_FullMethodName                     = "/platform.Platform/ServerInfo"
	Platform_UserStatus_FullMethodName                     = "/platform.Platform/UserStatus"
)

// PlatformClient is the client API for Platform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Platform contains rpcs related to Percona Platform.
type PlatformClient interface {
	// Connect a PMM server to the organization created on Percona Portal. That allows the user to sign in to the PMM server with their Percona Account.
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	// Disconnect a PMM server from the organization created on Percona Portal.
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
	// SearchOrganizationTickets searches support tickets belonging to the Percona Portal Organization that the PMM server is connected to.
	SearchOrganizationTickets(ctx context.Context, in *SearchOrganizationTicketsRequest, opts ...grpc.CallOption) (*SearchOrganizationTicketsResponse, error)
	// SearchOrganizationEntitlements fetches details of the entitlement's available to the Portal organization that the PMM server is connected to.
	SearchOrganizationEntitlements(ctx context.Context, in *SearchOrganizationEntitlementsRequest, opts ...grpc.CallOption) (*SearchOrganizationEntitlementsResponse, error)
	// GetContactInformation fetches the contact details of the customer success employee handling the Percona customer account from Percona Platform.
	GetContactInformation(ctx context.Context, in *GetContactInformationRequest, opts ...grpc.CallOption) (*GetContactInformationResponse, error)
	// ServerInfo returns PMM server ID and name.
	ServerInfo(ctx context.Context, in *ServerInfoRequest, opts ...grpc.CallOption) (*ServerInfoResponse, error)
	// UserStatus returns a boolean indicating whether the current user is logged in with their Percona Account or not.
	UserStatus(ctx context.Context, in *UserStatusRequest, opts ...grpc.CallOption) (*UserStatusResponse, error)
}

type platformClient struct {
	cc grpc.ClientConnInterface
}

func NewPlatformClient(cc grpc.ClientConnInterface) PlatformClient {
	return &platformClient{cc}
}

func (c *platformClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, Platform_Connect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, Platform_Disconnect_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) SearchOrganizationTickets(ctx context.Context, in *SearchOrganizationTicketsRequest, opts ...grpc.CallOption) (*SearchOrganizationTicketsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchOrganizationTicketsResponse)
	err := c.cc.Invoke(ctx, Platform_SearchOrganizationTickets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) SearchOrganizationEntitlements(ctx context.Context, in *SearchOrganizationEntitlementsRequest, opts ...grpc.CallOption) (*SearchOrganizationEntitlementsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchOrganizationEntitlementsResponse)
	err := c.cc.Invoke(ctx, Platform_SearchOrganizationEntitlements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) GetContactInformation(ctx context.Context, in *GetContactInformationRequest, opts ...grpc.CallOption) (*GetContactInformationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContactInformationResponse)
	err := c.cc.Invoke(ctx, Platform_GetContactInformation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) ServerInfo(ctx context.Context, in *ServerInfoRequest, opts ...grpc.CallOption) (*ServerInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerInfoResponse)
	err := c.cc.Invoke(ctx, Platform_ServerInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformClient) UserStatus(ctx context.Context, in *UserStatusRequest, opts ...grpc.CallOption) (*UserStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserStatusResponse)
	err := c.cc.Invoke(ctx, Platform_UserStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServer is the server API for Platform service.
// All implementations must embed UnimplementedPlatformServer
// for forward compatibility.
//
// Platform contains rpcs related to Percona Platform.
type PlatformServer interface {
	// Connect a PMM server to the organization created on Percona Portal. That allows the user to sign in to the PMM server with their Percona Account.
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	// Disconnect a PMM server from the organization created on Percona Portal.
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
	// SearchOrganizationTickets searches support tickets belonging to the Percona Portal Organization that the PMM server is connected to.
	SearchOrganizationTickets(context.Context, *SearchOrganizationTicketsRequest) (*SearchOrganizationTicketsResponse, error)
	// SearchOrganizationEntitlements fetches details of the entitlement's available to the Portal organization that the PMM server is connected to.
	SearchOrganizationEntitlements(context.Context, *SearchOrganizationEntitlementsRequest) (*SearchOrganizationEntitlementsResponse, error)
	// GetContactInformation fetches the contact details of the customer success employee handling the Percona customer account from Percona Platform.
	GetContactInformation(context.Context, *GetContactInformationRequest) (*GetContactInformationResponse, error)
	// ServerInfo returns PMM server ID and name.
	ServerInfo(context.Context, *ServerInfoRequest) (*ServerInfoResponse, error)
	// UserStatus returns a boolean indicating whether the current user is logged in with their Percona Account or not.
	UserStatus(context.Context, *UserStatusRequest) (*UserStatusResponse, error)
	mustEmbedUnimplementedPlatformServer()
}

// UnimplementedPlatformServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPlatformServer struct{}

func (UnimplementedPlatformServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func (UnimplementedPlatformServer) Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}

func (UnimplementedPlatformServer) SearchOrganizationTickets(context.Context, *SearchOrganizationTicketsRequest) (*SearchOrganizationTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchOrganizationTickets not implemented")
}

func (UnimplementedPlatformServer) SearchOrganizationEntitlements(context.Context, *SearchOrganizationEntitlementsRequest) (*SearchOrganizationEntitlementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchOrganizationEntitlements not implemented")
}

func (UnimplementedPlatformServer) GetContactInformation(context.Context, *GetContactInformationRequest) (*GetContactInformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContactInformation not implemented")
}

func (UnimplementedPlatformServer) ServerInfo(context.Context, *ServerInfoRequest) (*ServerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerInfo not implemented")
}

func (UnimplementedPlatformServer) UserStatus(context.Context, *UserStatusRequest) (*UserStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserStatus not implemented")
}
func (UnimplementedPlatformServer) mustEmbedUnimplementedPlatformServer() {}
func (UnimplementedPlatformServer) testEmbeddedByValue()                  {}

// UnsafePlatformServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlatformServer will
// result in compilation errors.
type UnsafePlatformServer interface {
	mustEmbedUnimplementedPlatformServer()
}

func RegisterPlatformServer(s grpc.ServiceRegistrar, srv PlatformServer) {
	// If the following call pancis, it indicates UnimplementedPlatformServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Platform_ServiceDesc, srv)
}

func _Platform_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_Connect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_Disconnect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_SearchOrganizationTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchOrganizationTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).SearchOrganizationTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_SearchOrganizationTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).SearchOrganizationTickets(ctx, req.(*SearchOrganizationTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_SearchOrganizationEntitlements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchOrganizationEntitlementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).SearchOrganizationEntitlements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_SearchOrganizationEntitlements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).SearchOrganizationEntitlements(ctx, req.(*SearchOrganizationEntitlementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_GetContactInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).GetContactInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_GetContactInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).GetContactInformation(ctx, req.(*GetContactInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_ServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).ServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_ServerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).ServerInfo(ctx, req.(*ServerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Platform_UserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).UserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_UserStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).UserStatus(ctx, req.(*UserStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Platform_ServiceDesc is the grpc.ServiceDesc for Platform service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Platform_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "platform.Platform",
	HandlerType: (*PlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Platform_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _Platform_Disconnect_Handler,
		},
		{
			MethodName: "SearchOrganizationTickets",
			Handler:    _Platform_SearchOrganizationTickets_Handler,
		},
		{
			MethodName: "SearchOrganizationEntitlements",
			Handler:    _Platform_SearchOrganizationEntitlements_Handler,
		},
		{
			MethodName: "GetContactInformation",
			Handler:    _Platform_GetContactInformation_Handler,
		},
		{
			MethodName: "ServerInfo",
			Handler:    _Platform_ServerInfo_Handler,
		},
		{
			MethodName: "UserStatus",
			Handler:    _Platform_UserStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platformpb/platform.proto",
}
