// Code generated by protoc-gen-go. DO NOT EDIT.
// source: time.proto

/*
Package mypackage is a generated protocol buffer package.

It is generated from these files:
	time.proto

It has these top-level messages:
	Result
	Request
*/
package mypackage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Result struct {
	Now           string `protobuf:"bytes,1,opt,name=Now,json=now" json:"Now,omitempty"`
	ProcesingTime int64  `protobuf:"varint,2,opt,name=ProcesingTime,json=procesingTime" json:"ProcesingTime,omitempty"`
	Error         string `protobuf:"bytes,3,opt,name=Error,json=error" json:"Error,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Result) GetNow() string {
	if m != nil {
		return m.Now
	}
	return ""
}

func (m *Result) GetProcesingTime() int64 {
	if m != nil {
		return m.ProcesingTime
	}
	return 0
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Request struct {
	TimeZone string `protobuf:"bytes,1,opt,name=TimeZone,json=timeZone" json:"TimeZone,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func init() {
	proto.RegisterType((*Result)(nil), "time.Result")
	proto.RegisterType((*Request)(nil), "time.Request")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Timer service

type TimerClient interface {
	ReturnTimeNow(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type timerClient struct {
	cc *grpc.ClientConn
}

func NewTimerClient(cc *grpc.ClientConn) TimerClient {
	return &timerClient{cc}
}

func (c *timerClient) ReturnTimeNow(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/time.Timer/ReturnTimeNow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Timer service

type TimerServer interface {
	ReturnTimeNow(context.Context, *Request) (*Result, error)
}

func RegisterTimerServer(s *grpc.Server, srv TimerServer) {
	s.RegisterService(&_Timer_serviceDesc, srv)
}

func _Timer_ReturnTimeNow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimerServer).ReturnTimeNow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/time.Timer/ReturnTimeNow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimerServer).ReturnTimeNow(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Timer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "time.Timer",
	HandlerType: (*TimerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReturnTimeNow",
			Handler:    _Timer_ReturnTimeNow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time.proto",
}

func init() { proto.RegisterFile("time.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcf, 0x0a, 0x82, 0x40,
	0x10, 0x87, 0xb3, 0x4d, 0xb3, 0xa1, 0x85, 0x18, 0x3a, 0x88, 0x27, 0x91, 0x02, 0x4f, 0x1e, 0xea,
	0xd0, 0x13, 0x74, 0x95, 0x58, 0xa2, 0x43, 0xc7, 0x62, 0x08, 0x21, 0x77, 0x6c, 0x5c, 0xf1, 0xf5,
	0x63, 0xc5, 0xa0, 0xdb, 0x7c, 0xf3, 0xe7, 0xc7, 0x37, 0x00, 0xae, 0x6e, 0xa8, 0x6c, 0x85, 0x1d,
	0xe3, 0xc2, 0xd7, 0xf9, 0x0d, 0x22, 0x43, 0x5d, 0xff, 0x76, 0xb8, 0x01, 0x55, 0xf1, 0x90, 0x04,
	0x59, 0x50, 0xac, 0x8c, 0xb2, 0x3c, 0xe0, 0x0e, 0xf4, 0x45, 0xf8, 0x49, 0x5d, 0x6d, 0x5f, 0xd7,
	0xba, 0xa1, 0x64, 0x9e, 0x05, 0x85, 0x32, 0xba, 0xfd, 0x6f, 0xe2, 0x16, 0xc2, 0xb3, 0x08, 0x4b,
	0xa2, 0xc6, 0xcb, 0x90, 0x3c, 0xe4, 0x7b, 0x58, 0x1a, 0xfa, 0xf4, 0xd4, 0x39, 0x4c, 0x21, 0xf6,
	0x8b, 0x77, 0xb6, 0x34, 0xa5, 0xc7, 0x6e, 0xe2, 0xc3, 0x09, 0x42, 0x3f, 0x13, 0x2c, 0x41, 0x1b,
	0x72, 0xbd, 0x58, 0x8f, 0x15, 0x0f, 0xa8, 0xcb, 0xd1, 0x75, 0x0a, 0x49, 0xd7, 0x3f, 0xf4, 0xae,
	0xf9, 0xec, 0x11, 0x8d, 0x4f, 0x1c, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x25, 0x6a, 0xc6, 0x95,
	0xd2, 0x00, 0x00, 0x00,
}
