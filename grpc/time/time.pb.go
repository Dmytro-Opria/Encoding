// Code generated by protoc-gen-go. DO NOT EDIT.
// source: time.proto

/*
Package time is a generated protocol buffer package.

It is generated from these files:
	time.proto

It has these top-level messages:
	Result
	Request
*/
package time

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

func init() { proto.RegisterFile("time.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x1c, 0xb8, 0xd8, 0x82, 0x52,
	0x8b, 0x4b, 0x73, 0x4a, 0x84, 0x04, 0xb8, 0x98, 0xfd, 0xf2, 0xcb, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0x98, 0xf3, 0xf2, 0xcb, 0x85, 0x54, 0xb8, 0x78, 0x03, 0x8a, 0xf2, 0x93, 0x53, 0x8b,
	0x33, 0xf3, 0xd2, 0x43, 0x32, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x98, 0x83, 0x78, 0x0b,
	0x90, 0x05, 0x95, 0x54, 0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8,
	0x38, 0x40, 0x42, 0x51, 0xf9, 0x79, 0xa9, 0x50, 0x73, 0x38, 0x4a, 0xa0, 0xfc, 0x24, 0x36, 0xb0,
	0xad, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x9f, 0x69, 0xc1, 0x83, 0x00, 0x00, 0x00,
}