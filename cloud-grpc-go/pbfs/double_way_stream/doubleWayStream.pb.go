// Code generated by protoc-gen-go. DO NOT EDIT.
// source: doubleWayStream.proto

package double_way_stream

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

type RequestMessage struct {
	ReqMsg               string   `protobuf:"bytes,1,opt,name=req_msg,json=reqMsg,proto3" json:"req_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMessage) Reset()         { *m = RequestMessage{} }
func (m *RequestMessage) String() string { return proto.CompactTextString(m) }
func (*RequestMessage) ProtoMessage()    {}
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ce2a594dcac19e3, []int{0}
}

func (m *RequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMessage.Unmarshal(m, b)
}
func (m *RequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMessage.Marshal(b, m, deterministic)
}
func (m *RequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMessage.Merge(m, src)
}
func (m *RequestMessage) XXX_Size() int {
	return xxx_messageInfo_RequestMessage.Size(m)
}
func (m *RequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMessage proto.InternalMessageInfo

func (m *RequestMessage) GetReqMsg() string {
	if m != nil {
		return m.ReqMsg
	}
	return ""
}

type ResponseMessage struct {
	RspMsg               string   `protobuf:"bytes,1,opt,name=rsp_msg,json=rspMsg,proto3" json:"rsp_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMessage) Reset()         { *m = ResponseMessage{} }
func (m *ResponseMessage) String() string { return proto.CompactTextString(m) }
func (*ResponseMessage) ProtoMessage()    {}
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ce2a594dcac19e3, []int{1}
}

func (m *ResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMessage.Unmarshal(m, b)
}
func (m *ResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMessage.Marshal(b, m, deterministic)
}
func (m *ResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMessage.Merge(m, src)
}
func (m *ResponseMessage) XXX_Size() int {
	return xxx_messageInfo_ResponseMessage.Size(m)
}
func (m *ResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMessage proto.InternalMessageInfo

func (m *ResponseMessage) GetRspMsg() string {
	if m != nil {
		return m.RspMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestMessage)(nil), "double_way_stream.RequestMessage")
	proto.RegisterType((*ResponseMessage)(nil), "double_way_stream.ResponseMessage")
}

func init() { proto.RegisterFile("doubleWayStream.proto", fileDescriptor_6ce2a594dcac19e3) }

var fileDescriptor_6ce2a594dcac19e3 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xc9, 0x2f, 0x4d,
	0xca, 0x49, 0x0d, 0x4f, 0xac, 0x0c, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x84, 0x08, 0xc7, 0x97, 0x27, 0x56, 0xc6, 0x17, 0x83, 0x25, 0x94, 0x34, 0xb9,
	0xf8, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85,
	0xc4, 0xb9, 0xd8, 0x8b, 0x52, 0x0b, 0xe3, 0x73, 0x8b, 0xd3, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xd8, 0x8a, 0x52, 0x0b, 0x7d, 0x8b, 0xd3, 0x95, 0xb4, 0xb8, 0xf8, 0x83, 0x52, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x91, 0xd5, 0x16, 0x17, 0xa0, 0xa8, 0x2d, 0x2e, 0xf0, 0x2d, 0x4e, 0x37,
	0xaa, 0xe6, 0x12, 0x73, 0x41, 0x75, 0x42, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x50, 0x22,
	0x97, 0x10, 0x9a, 0x8c, 0x5b, 0x69, 0x9e, 0x90, 0xa2, 0x1e, 0x86, 0xd3, 0xf4, 0x50, 0xdd, 0x25,
	0xa5, 0x84, 0x55, 0x09, 0x8a, 0x7b, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x9d, 0xc2, 0xb8, 0xe4,
	0x92, 0xf3, 0x73, 0xf5, 0x92, 0x73, 0xf2, 0x4b, 0x53, 0xf4, 0xd2, 0x8b, 0x0a, 0x92, 0xf5, 0xd0,
	0x82, 0xc3, 0x49, 0x04, 0xcd, 0x09, 0x01, 0xa0, 0xe0, 0x09, 0x60, 0x8c, 0x12, 0x2b, 0x48, 0x4a,
	0x2b, 0xd6, 0xc7, 0xb0, 0x67, 0x11, 0x13, 0xb3, 0x4b, 0x78, 0x70, 0x12, 0x1b, 0x38, 0x14, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xba, 0x72, 0x50, 0x5e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DoubleWayStreamServiceClient is the client API for DoubleWayStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DoubleWayStreamServiceClient interface {
	DoubleWayStreamFun(ctx context.Context, opts ...grpc.CallOption) (DoubleWayStreamService_DoubleWayStreamFunClient, error)
}

type doubleWayStreamServiceClient struct {
	cc *grpc.ClientConn
}

func NewDoubleWayStreamServiceClient(cc *grpc.ClientConn) DoubleWayStreamServiceClient {
	return &doubleWayStreamServiceClient{cc}
}

func (c *doubleWayStreamServiceClient) DoubleWayStreamFun(ctx context.Context, opts ...grpc.CallOption) (DoubleWayStreamService_DoubleWayStreamFunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DoubleWayStreamService_serviceDesc.Streams[0], "/double_way_stream.DoubleWayStreamService/DoubleWayStreamFun", opts...)
	if err != nil {
		return nil, err
	}
	x := &doubleWayStreamServiceDoubleWayStreamFunClient{stream}
	return x, nil
}

type DoubleWayStreamService_DoubleWayStreamFunClient interface {
	Send(*RequestMessage) error
	Recv() (*ResponseMessage, error)
	grpc.ClientStream
}

type doubleWayStreamServiceDoubleWayStreamFunClient struct {
	grpc.ClientStream
}

func (x *doubleWayStreamServiceDoubleWayStreamFunClient) Send(m *RequestMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *doubleWayStreamServiceDoubleWayStreamFunClient) Recv() (*ResponseMessage, error) {
	m := new(ResponseMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DoubleWayStreamServiceServer is the server API for DoubleWayStreamService service.
type DoubleWayStreamServiceServer interface {
	DoubleWayStreamFun(DoubleWayStreamService_DoubleWayStreamFunServer) error
}

// UnimplementedDoubleWayStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDoubleWayStreamServiceServer struct {
}

func (*UnimplementedDoubleWayStreamServiceServer) DoubleWayStreamFun(srv DoubleWayStreamService_DoubleWayStreamFunServer) error {
	return status.Errorf(codes.Unimplemented, "method DoubleWayStreamFun not implemented")
}

func RegisterDoubleWayStreamServiceServer(s *grpc.Server, srv DoubleWayStreamServiceServer) {
	s.RegisterService(&_DoubleWayStreamService_serviceDesc, srv)
}

func _DoubleWayStreamService_DoubleWayStreamFun_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DoubleWayStreamServiceServer).DoubleWayStreamFun(&doubleWayStreamServiceDoubleWayStreamFunServer{stream})
}

type DoubleWayStreamService_DoubleWayStreamFunServer interface {
	Send(*ResponseMessage) error
	Recv() (*RequestMessage, error)
	grpc.ServerStream
}

type doubleWayStreamServiceDoubleWayStreamFunServer struct {
	grpc.ServerStream
}

func (x *doubleWayStreamServiceDoubleWayStreamFunServer) Send(m *ResponseMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *doubleWayStreamServiceDoubleWayStreamFunServer) Recv() (*RequestMessage, error) {
	m := new(RequestMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DoubleWayStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "double_way_stream.DoubleWayStreamService",
	HandlerType: (*DoubleWayStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoubleWayStreamFun",
			Handler:       _DoubleWayStreamService_DoubleWayStreamFun_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "doubleWayStream.proto",
}
