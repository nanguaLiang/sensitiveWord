// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wordfilter.proto

package wordfilter

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ContainSensitiveWordRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainSensitiveWordRequest) Reset()         { *m = ContainSensitiveWordRequest{} }
func (m *ContainSensitiveWordRequest) String() string { return proto.CompactTextString(m) }
func (*ContainSensitiveWordRequest) ProtoMessage()    {}
func (*ContainSensitiveWordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6007921b45e0112, []int{0}
}

func (m *ContainSensitiveWordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainSensitiveWordRequest.Unmarshal(m, b)
}
func (m *ContainSensitiveWordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainSensitiveWordRequest.Marshal(b, m, deterministic)
}
func (m *ContainSensitiveWordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainSensitiveWordRequest.Merge(m, src)
}
func (m *ContainSensitiveWordRequest) XXX_Size() int {
	return xxx_messageInfo_ContainSensitiveWordRequest.Size(m)
}
func (m *ContainSensitiveWordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainSensitiveWordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContainSensitiveWordRequest proto.InternalMessageInfo

func (m *ContainSensitiveWordRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type ContainSensitiveWordResponse struct {
	IsSensitiveWord      bool     `protobuf:"varint,1,opt,name=isSensitiveWord,proto3" json:"isSensitiveWord,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainSensitiveWordResponse) Reset()         { *m = ContainSensitiveWordResponse{} }
func (m *ContainSensitiveWordResponse) String() string { return proto.CompactTextString(m) }
func (*ContainSensitiveWordResponse) ProtoMessage()    {}
func (*ContainSensitiveWordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6007921b45e0112, []int{1}
}

func (m *ContainSensitiveWordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainSensitiveWordResponse.Unmarshal(m, b)
}
func (m *ContainSensitiveWordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainSensitiveWordResponse.Marshal(b, m, deterministic)
}
func (m *ContainSensitiveWordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainSensitiveWordResponse.Merge(m, src)
}
func (m *ContainSensitiveWordResponse) XXX_Size() int {
	return xxx_messageInfo_ContainSensitiveWordResponse.Size(m)
}
func (m *ContainSensitiveWordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainSensitiveWordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContainSensitiveWordResponse proto.InternalMessageInfo

func (m *ContainSensitiveWordResponse) GetIsSensitiveWord() bool {
	if m != nil {
		return m.IsSensitiveWord
	}
	return false
}

func init() {
	proto.RegisterType((*ContainSensitiveWordRequest)(nil), "wordfilter.ContainSensitiveWordRequest")
	proto.RegisterType((*ContainSensitiveWordResponse)(nil), "wordfilter.ContainSensitiveWordResponse")
}

func init() { proto.RegisterFile("wordfilter.proto", fileDescriptor_d6007921b45e0112) }

var fileDescriptor_d6007921b45e0112 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xcf, 0x2f, 0x4a,
	0x49, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x19, 0x72, 0x49, 0x3b, 0xe7, 0xe7, 0x95, 0x24, 0x66, 0xe6, 0x05, 0xa7, 0xe6, 0x15, 0x67,
	0x96, 0x64, 0x96, 0xa5, 0x86, 0xe7, 0x17, 0xa5, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08,
	0x09, 0x71, 0xb1, 0x94, 0xa4, 0x56, 0x94, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x4a, 0x1e, 0x5c, 0x32, 0xd8, 0xb5, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x69, 0x70, 0xf1,
	0x67, 0x16, 0xa3, 0x48, 0x81, 0xb5, 0x73, 0x04, 0xa1, 0x0b, 0x1b, 0x35, 0x32, 0x72, 0x09, 0x82,
	0x18, 0x6e, 0x60, 0xb7, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xe5, 0x70, 0x89, 0x42,
	0xcd, 0x47, 0x55, 0x2e, 0xa4, 0xae, 0x87, 0xe4, 0x15, 0x3c, 0xae, 0x96, 0xd2, 0x20, 0xac, 0x10,
	0xe2, 0x56, 0x25, 0x86, 0x24, 0x36, 0x70, 0x98, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf6,
	0x8b, 0x58, 0xf8, 0x27, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WordFilterServiceClient is the client API for WordFilterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WordFilterServiceClient interface {
	ContainsSensitiveWord(ctx context.Context, in *ContainSensitiveWordRequest, opts ...grpc.CallOption) (*ContainSensitiveWordResponse, error)
}

type wordFilterServiceClient struct {
	cc *grpc.ClientConn
}

func NewWordFilterServiceClient(cc *grpc.ClientConn) WordFilterServiceClient {
	return &wordFilterServiceClient{cc}
}

func (c *wordFilterServiceClient) ContainsSensitiveWord(ctx context.Context, in *ContainSensitiveWordRequest, opts ...grpc.CallOption) (*ContainSensitiveWordResponse, error) {
	out := new(ContainSensitiveWordResponse)
	err := c.cc.Invoke(ctx, "/wordfilter.WordFilterService/ContainsSensitiveWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordFilterServiceServer is the server API for WordFilterService service.
type WordFilterServiceServer interface {
	ContainsSensitiveWord(context.Context, *ContainSensitiveWordRequest) (*ContainSensitiveWordResponse, error)
}

func RegisterWordFilterServiceServer(s *grpc.Server, srv WordFilterServiceServer) {
	s.RegisterService(&_WordFilterService_serviceDesc, srv)
}

func _WordFilterService_ContainsSensitiveWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainSensitiveWordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordFilterServiceServer).ContainsSensitiveWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wordfilter.WordFilterService/ContainsSensitiveWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordFilterServiceServer).ContainsSensitiveWord(ctx, req.(*ContainSensitiveWordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WordFilterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wordfilter.WordFilterService",
	HandlerType: (*WordFilterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContainsSensitiveWord",
			Handler:    _WordFilterService_ContainsSensitiveWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wordfilter.proto",
}
