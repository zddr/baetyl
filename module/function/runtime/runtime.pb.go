// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runtime.proto

package runtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// The request message.
type Message struct {
	QOS                  uint32   `protobuf:"varint,1,opt,name=QOS,proto3" json:"QOS,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=Payload,proto3" json:"Payload,omitempty"`
	FunctionName         string   `protobuf:"bytes,11,opt,name=FunctionName,proto3" json:"FunctionName,omitempty"`
	FunctionInvokeID     string   `protobuf:"bytes,12,opt,name=FunctionInvokeID,proto3" json:"FunctionInvokeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2dd377c869464, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetQOS() uint32 {
	if m != nil {
		return m.QOS
	}
	return 0
}

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *Message) GetFunctionInvokeID() string {
	if m != nil {
		return m.FunctionInvokeID
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "runtime.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RuntimeClient is the client API for Runtime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RuntimeClient interface {
	// Handle handles request
	Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type runtimeClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeClient(cc *grpc.ClientConn) RuntimeClient {
	return &runtimeClient{cc}
}

func (c *runtimeClient) Handle(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/runtime.Runtime/Handle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeServer is the server API for Runtime service.
type RuntimeServer interface {
	// Handle handles request
	Handle(context.Context, *Message) (*Message, error)
}

func RegisterRuntimeServer(s *grpc.Server, srv RuntimeServer) {
	s.RegisterService(&_Runtime_serviceDesc, srv)
}

func _Runtime_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/runtime.Runtime/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServer).Handle(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Runtime_serviceDesc = grpc.ServiceDesc{
	ServiceName: "runtime.Runtime",
	HandlerType: (*RuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _Runtime_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runtime.proto",
}

func init() { proto.RegisterFile("runtime.proto", fileDescriptor_86e2dd377c869464) }

var fileDescriptor_86e2dd377c869464 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2a, 0xcd, 0x2b,
	0xc9, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x66, 0x33,
	0x72, 0xb1, 0xfb, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0x09, 0x70, 0x31, 0x07, 0xfa, 0x07,
	0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x21, 0xf9, 0x05,
	0x99, 0xc9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x04, 0x17, 0x7b, 0x40,
	0x62, 0x65, 0x4e, 0x7e, 0x62, 0x8a, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x8c, 0x2b, 0xa4,
	0xc4, 0xc5, 0xe3, 0x56, 0x9a, 0x97, 0x5c, 0x92, 0x99, 0x9f, 0xe7, 0x97, 0x98, 0x9b, 0x2a, 0xc1,
	0x0d, 0xd6, 0x86, 0x22, 0x26, 0xa4, 0xc5, 0x25, 0x00, 0xe3, 0x7b, 0xe6, 0x95, 0xe5, 0x67, 0xa7,
	0x7a, 0xba, 0x48, 0xf0, 0x80, 0xd5, 0x61, 0x88, 0x1b, 0x59, 0x72, 0xb1, 0x07, 0x41, 0x1c, 0x2a,
	0xa4, 0xc7, 0xc5, 0xe6, 0x91, 0x98, 0x97, 0x92, 0x93, 0x2a, 0x24, 0xa0, 0x07, 0xf3, 0x0b, 0xd4,
	0xe1, 0x52, 0x18, 0x22, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x8f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xd0, 0x55, 0x89, 0x55, 0xf9, 0x00, 0x00, 0x00,
}
