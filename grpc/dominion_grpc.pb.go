// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dominion

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

// DominionClient is the client API for Dominion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DominionClient interface {
	// GetServices returns the available services and their locations
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesReply, error)
	// GetDomains returns all domains and their services
	GetDomains(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetDomainsReply, error)
}

type dominionClient struct {
	cc grpc.ClientConnInterface
}

func NewDominionClient(cc grpc.ClientConnInterface) DominionClient {
	return &dominionClient{cc}
}

func (c *dominionClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesReply, error) {
	out := new(GetServicesReply)
	err := c.cc.Invoke(ctx, "/grpc.Dominion/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dominionClient) GetDomains(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetDomainsReply, error) {
	out := new(GetDomainsReply)
	err := c.cc.Invoke(ctx, "/grpc.Dominion/GetDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DominionServer is the server API for Dominion service.
// All implementations must embed UnimplementedDominionServer
// for forward compatibility
type DominionServer interface {
	// GetServices returns the available services and their locations
	GetServices(context.Context, *GetServicesRequest) (*GetServicesReply, error)
	// GetDomains returns all domains and their services
	GetDomains(context.Context, *Empty) (*GetDomainsReply, error)
	mustEmbedUnimplementedDominionServer()
}

// UnimplementedDominionServer must be embedded to have forward compatible implementations.
type UnimplementedDominionServer struct {
}

func (UnimplementedDominionServer) GetServices(context.Context, *GetServicesRequest) (*GetServicesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedDominionServer) GetDomains(context.Context, *Empty) (*GetDomainsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomains not implemented")
}
func (UnimplementedDominionServer) mustEmbedUnimplementedDominionServer() {}

// UnsafeDominionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DominionServer will
// result in compilation errors.
type UnsafeDominionServer interface {
	mustEmbedUnimplementedDominionServer()
}

func RegisterDominionServer(s grpc.ServiceRegistrar, srv DominionServer) {
	s.RegisterService(&Dominion_ServiceDesc, srv)
}

func _Dominion_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DominionServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Dominion/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DominionServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dominion_GetDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DominionServer).GetDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Dominion/GetDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DominionServer).GetDomains(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Dominion_ServiceDesc is the grpc.ServiceDesc for Dominion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dominion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Dominion",
	HandlerType: (*DominionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServices",
			Handler:    _Dominion_GetServices_Handler,
		},
		{
			MethodName: "GetDomains",
			Handler:    _Dominion_GetDomains_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dominion.proto",
}
