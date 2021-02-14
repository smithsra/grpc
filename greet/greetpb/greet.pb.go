// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greetpb/greet.proto

package greetpb

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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

// Request message
type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

// Response message
type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetManytTimesRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetManytTimesRequest) Reset()         { *m = GreetManytTimesRequest{} }
func (m *GreetManytTimesRequest) String() string { return proto.CompactTextString(m) }
func (*GreetManytTimesRequest) ProtoMessage()    {}
func (*GreetManytTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{3}
}

func (m *GreetManytTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManytTimesRequest.Unmarshal(m, b)
}
func (m *GreetManytTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManytTimesRequest.Marshal(b, m, deterministic)
}
func (m *GreetManytTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManytTimesRequest.Merge(m, src)
}
func (m *GreetManytTimesRequest) XXX_Size() int {
	return xxx_messageInfo_GreetManytTimesRequest.Size(m)
}
func (m *GreetManytTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManytTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManytTimesRequest proto.InternalMessageInfo

func (m *GreetManytTimesRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetManytTimesResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetManytTimesResponse) Reset()         { *m = GreetManytTimesResponse{} }
func (m *GreetManytTimesResponse) String() string { return proto.CompactTextString(m) }
func (*GreetManytTimesResponse) ProtoMessage()    {}
func (*GreetManytTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe6f881da19a2871, []int{4}
}

func (m *GreetManytTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManytTimesResponse.Unmarshal(m, b)
}
func (m *GreetManytTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManytTimesResponse.Marshal(b, m, deterministic)
}
func (m *GreetManytTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManytTimesResponse.Merge(m, src)
}
func (m *GreetManytTimesResponse) XXX_Size() int {
	return xxx_messageInfo_GreetManytTimesResponse.Size(m)
}
func (m *GreetManytTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManytTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManytTimesResponse proto.InternalMessageInfo

func (m *GreetManytTimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*GreetManytTimesRequest)(nil), "greet.GreetManytTimesRequest")
	proto.RegisterType((*GreetManytTimesResponse)(nil), "greet.GreetManytTimesResponse")
}

func init() {
	proto.RegisterFile("greet/greetpb/greet.proto", fileDescriptor_fe6f881da19a2871)
}

var fileDescriptor_fe6f881da19a2871 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x07, 0x93, 0x05, 0x49, 0x10, 0x5a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x15,
	0xcc, 0x51, 0x72, 0xe3, 0xe2, 0x70, 0x07, 0x31, 0x32, 0xf3, 0xd2, 0x85, 0x64, 0xb9, 0xb8, 0xd2,
	0x32, 0x8b, 0x8a, 0x4b, 0xe2, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x38, 0xc1, 0x22, 0x7e, 0x40, 0x01, 0x21, 0x69, 0x2e, 0xce, 0x9c, 0x44, 0x98, 0x2c, 0x13, 0x58,
	0x96, 0x03, 0x24, 0x00, 0x92, 0x54, 0xb2, 0xe6, 0xe2, 0x01, 0x9b, 0x13, 0x94, 0x5a, 0x58, 0x9a,
	0x5a, 0x5c, 0x22, 0xa4, 0xcd, 0xc5, 0x91, 0x0e, 0x35, 0x17, 0x6c, 0x12, 0xb7, 0x11, 0xbf, 0x1e,
	0xc4, 0x7a, 0x98, 0x75, 0x41, 0x70, 0x05, 0x4a, 0xea, 0x5c, 0xbc, 0x50, 0xcd, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x50, 0x57, 0x40,
	0x79, 0x4a, 0xae, 0x5c, 0x62, 0x60, 0x85, 0xbe, 0x89, 0x79, 0x95, 0x25, 0x21, 0x99, 0xb9, 0xa9,
	0xc5, 0x64, 0xd9, 0x67, 0xc8, 0x25, 0x8e, 0x61, 0x0c, 0x7e, 0x9b, 0x8d, 0x66, 0x30, 0x42, 0x3d,
	0x18, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xc2, 0xc5, 0x0a, 0xe6, 0x0b, 0x09, 0x23,
	0xdb, 0x03, 0x75, 0x8e, 0x94, 0x08, 0xaa, 0x20, 0xc4, 0x70, 0x25, 0x06, 0xa1, 0x20, 0x2e, 0x7e,
	0x34, 0x9b, 0x85, 0x64, 0x91, 0x95, 0x62, 0x78, 0x4c, 0x4a, 0x0e, 0x97, 0x34, 0xc4, 0x4c, 0x03,
	0x46, 0x27, 0xce, 0x28, 0x76, 0x68, 0x04, 0x27, 0xb1, 0x81, 0xe3, 0xd6, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xf9, 0x0f, 0x8d, 0xca, 0xf8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// Unary API
	// Create an rpc that takes a GreetRequest and returns a GreetResponse
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	//Server streaming
	GreetManytTimes(ctx context.Context, in *GreetManytTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManytTimesClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetManytTimes(ctx context.Context, in *GreetManytTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManytTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetManytTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManytTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManytTimesClient interface {
	Recv() (*GreetManytTimesResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManytTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManytTimesClient) Recv() (*GreetManytTimesResponse, error) {
	m := new(GreetManytTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Unary API
	// Create an rpc that takes a GreetRequest and returns a GreetResponse
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	//Server streaming
	GreetManytTimes(*GreetManytTimesRequest, GreetService_GreetManytTimesServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreetServiceServer) GreetManytTimes(req *GreetManytTimesRequest, srv GreetService_GreetManytTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManytTimes not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetManytTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManytTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManytTimes(m, &greetServiceGreetManytTimesServer{stream})
}

type GreetService_GreetManytTimesServer interface {
	Send(*GreetManytTimesResponse) error
	grpc.ServerStream
}

type greetServiceGreetManytTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManytTimesServer) Send(m *GreetManytTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManytTimes",
			Handler:       _GreetService_GreetManytTimes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greet/greetpb/greet.proto",
}
