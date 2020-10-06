// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dominion.proto

package grpcd

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// GetServicesRequest contains the type of the service the client is searching for
type GetServicesRequest struct {
	// Name is the name of the service which is requested
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServicesRequest) Reset()         { *m = GetServicesRequest{} }
func (m *GetServicesRequest) String() string { return proto.CompactTextString(m) }
func (*GetServicesRequest) ProtoMessage()    {}
func (*GetServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb6c7947614a23d, []int{0}
}

func (m *GetServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServicesRequest.Unmarshal(m, b)
}
func (m *GetServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServicesRequest.Marshal(b, m, deterministic)
}
func (m *GetServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServicesRequest.Merge(m, src)
}
func (m *GetServicesRequest) XXX_Size() int {
	return xxx_messageInfo_GetServicesRequest.Size(m)
}
func (m *GetServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServicesRequest proto.InternalMessageInfo

func (m *GetServicesRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// GetServicesReply contains the list of known addresses hosting the requested service
type GetServicesReply struct {
	// Services is the list of services of the requested type
	Services             []*ServiceIdentity `protobuf:"bytes,1,rep,name=Services,proto3" json:"Services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetServicesReply) Reset()         { *m = GetServicesReply{} }
func (m *GetServicesReply) String() string { return proto.CompactTextString(m) }
func (*GetServicesReply) ProtoMessage()    {}
func (*GetServicesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb6c7947614a23d, []int{1}
}

func (m *GetServicesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServicesReply.Unmarshal(m, b)
}
func (m *GetServicesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServicesReply.Marshal(b, m, deterministic)
}
func (m *GetServicesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServicesReply.Merge(m, src)
}
func (m *GetServicesReply) XXX_Size() int {
	return xxx_messageInfo_GetServicesReply.Size(m)
}
func (m *GetServicesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServicesReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetServicesReply proto.InternalMessageInfo

func (m *GetServicesReply) GetServices() []*ServiceIdentity {
	if m != nil {
		return m.Services
	}
	return nil
}

// GetDomainsReply contains all domains and their services
type GetDomainsReply struct {
	// Domains is the list of all domains
	Domains              []*DomainIdentity `protobuf:"bytes,1,rep,name=Domains,proto3" json:"Domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetDomainsReply) Reset()         { *m = GetDomainsReply{} }
func (m *GetDomainsReply) String() string { return proto.CompactTextString(m) }
func (*GetDomainsReply) ProtoMessage()    {}
func (*GetDomainsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb6c7947614a23d, []int{2}
}

func (m *GetDomainsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDomainsReply.Unmarshal(m, b)
}
func (m *GetDomainsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDomainsReply.Marshal(b, m, deterministic)
}
func (m *GetDomainsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDomainsReply.Merge(m, src)
}
func (m *GetDomainsReply) XXX_Size() int {
	return xxx_messageInfo_GetDomainsReply.Size(m)
}
func (m *GetDomainsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDomainsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetDomainsReply proto.InternalMessageInfo

func (m *GetDomainsReply) GetDomains() []*DomainIdentity {
	if m != nil {
		return m.Domains
	}
	return nil
}

// Empty is empty, Duh.
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb6c7947614a23d, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetServicesRequest)(nil), "grpcd.GetServicesRequest")
	proto.RegisterType((*GetServicesReply)(nil), "grpcd.GetServicesReply")
	proto.RegisterType((*GetDomainsReply)(nil), "grpcd.GetDomainsReply")
	proto.RegisterType((*Empty)(nil), "grpcd.Empty")
}

func init() { proto.RegisterFile("dominion.proto", fileDescriptor_3fb6c7947614a23d) }

var fileDescriptor_3fb6c7947614a23d = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xc9, 0xcf, 0xcd,
	0xcc, 0xcb, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2f, 0x2a, 0x48,
	0x4e, 0x91, 0xe2, 0xcb, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x84, 0x08, 0x2b, 0x69, 0x70,
	0x09, 0xb9, 0xa7, 0x96, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x16, 0x07, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0x84, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0xd9, 0x4a, 0x6e, 0x5c, 0x02, 0x28, 0x2a, 0x0b, 0x72, 0x2a, 0x85, 0x8c, 0xb8,
	0x38, 0x60, 0x02, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x62, 0x7a, 0x60, 0x7b, 0xf4, 0xa0,
	0xc2, 0x9e, 0x50, 0xdb, 0x82, 0xe0, 0xea, 0x94, 0x9c, 0xb8, 0xf8, 0xdd, 0x53, 0x4b, 0x5c, 0xf2,
	0x73, 0x13, 0x33, 0xf3, 0xa0, 0xc6, 0xe8, 0x73, 0xb1, 0x43, 0xf9, 0x50, 0x53, 0x44, 0xa1, 0xa6,
	0x40, 0x44, 0xe1, 0x86, 0xc0, 0x54, 0x29, 0xb1, 0x73, 0xb1, 0xba, 0xe6, 0x16, 0x94, 0x54, 0x1a,
	0xb5, 0x32, 0x72, 0x71, 0xb8, 0x40, 0x3d, 0x2a, 0xe4, 0xcc, 0xc5, 0x8d, 0xe4, 0x42, 0x21, 0x49,
	0xa8, 0x21, 0x98, 0xfe, 0x93, 0x12, 0xc7, 0x26, 0x55, 0x90, 0x53, 0xa9, 0xc4, 0x20, 0x64, 0xc2,
	0xc5, 0x85, 0x70, 0x9e, 0x10, 0x0f, 0x54, 0x21, 0xd8, 0x36, 0x29, 0x31, 0x84, 0x36, 0x64, 0xf7,
	0x2b, 0x31, 0x24, 0xb1, 0x81, 0x43, 0xd3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x22, 0x29, 0x8d,
	0xdb, 0x76, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DominionClient is the client API for Dominion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DominionClient interface {
	// GetServices returns the availible services and their locations
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesReply, error)
	// GetDomains returns all domains and their services
	GetDomains(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetDomainsReply, error)
}

type dominionClient struct {
	cc *grpc.ClientConn
}

func NewDominionClient(cc *grpc.ClientConn) DominionClient {
	return &dominionClient{cc}
}

func (c *dominionClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesReply, error) {
	out := new(GetServicesReply)
	err := c.cc.Invoke(ctx, "/grpcd.Dominion/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dominionClient) GetDomains(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetDomainsReply, error) {
	out := new(GetDomainsReply)
	err := c.cc.Invoke(ctx, "/grpcd.Dominion/GetDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DominionServer is the server API for Dominion service.
type DominionServer interface {
	// GetServices returns the availible services and their locations
	GetServices(context.Context, *GetServicesRequest) (*GetServicesReply, error)
	// GetDomains returns all domains and their services
	GetDomains(context.Context, *Empty) (*GetDomainsReply, error)
}

// UnimplementedDominionServer can be embedded to have forward compatible implementations.
type UnimplementedDominionServer struct {
}

func (*UnimplementedDominionServer) GetServices(ctx context.Context, req *GetServicesRequest) (*GetServicesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (*UnimplementedDominionServer) GetDomains(ctx context.Context, req *Empty) (*GetDomainsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomains not implemented")
}

func RegisterDominionServer(s *grpc.Server, srv DominionServer) {
	s.RegisterService(&_Dominion_serviceDesc, srv)
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
		FullMethod: "/grpcd.Dominion/GetServices",
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
		FullMethod: "/grpcd.Dominion/GetDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DominionServer).GetDomains(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dominion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcd.Dominion",
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
