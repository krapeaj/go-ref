// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	gateway.proto

It has these top-level messages:
	WebSocketResponseHeader
	WebSocketRealtimeResponseBody
	WebSocketRealtimeResponse
*/
package grpc

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

type WebSocketResponseHeader_WebSocketEventType int32

const (
	WebSocketResponseHeader_UPDATE_TOTAL_ACCESS_COUNT        WebSocketResponseHeader_WebSocketEventType = 0
	WebSocketResponseHeader_UPDATE_CURRENT_CONNECTION_COUNT  WebSocketResponseHeader_WebSocketEventType = 1
	WebSocketResponseHeader_UPDATE_CURRENT_NODE_COUNT        WebSocketResponseHeader_WebSocketEventType = 2
	WebSocketResponseHeader_UPDATE_CURRENT_MASTER_IDENTIFIER WebSocketResponseHeader_WebSocketEventType = 3
)

var WebSocketResponseHeader_WebSocketEventType_name = map[int32]string{
	0: "UPDATE_TOTAL_ACCESS_COUNT",
	1: "UPDATE_CURRENT_CONNECTION_COUNT",
	2: "UPDATE_CURRENT_NODE_COUNT",
	3: "UPDATE_CURRENT_MASTER_IDENTIFIER",
}
var WebSocketResponseHeader_WebSocketEventType_value = map[string]int32{
	"UPDATE_TOTAL_ACCESS_COUNT":        0,
	"UPDATE_CURRENT_CONNECTION_COUNT":  1,
	"UPDATE_CURRENT_NODE_COUNT":        2,
	"UPDATE_CURRENT_MASTER_IDENTIFIER": 3,
}

func (x WebSocketResponseHeader_WebSocketEventType) String() string {
	return proto.EnumName(WebSocketResponseHeader_WebSocketEventType_name, int32(x))
}
func (WebSocketResponseHeader_WebSocketEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type WebSocketResponseHeader struct {
	EventType WebSocketResponseHeader_WebSocketEventType `protobuf:"varint,1,opt,name=eventType,enum=grpc.WebSocketResponseHeader_WebSocketEventType" json:"eventType,omitempty"`
}

func (m *WebSocketResponseHeader) Reset()                    { *m = WebSocketResponseHeader{} }
func (m *WebSocketResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*WebSocketResponseHeader) ProtoMessage()               {}
func (*WebSocketResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WebSocketResponseHeader) GetEventType() WebSocketResponseHeader_WebSocketEventType {
	if m != nil {
		return m.EventType
	}
	return WebSocketResponseHeader_UPDATE_TOTAL_ACCESS_COUNT
}

type WebSocketRealtimeResponseBody struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *WebSocketRealtimeResponseBody) Reset()                    { *m = WebSocketRealtimeResponseBody{} }
func (m *WebSocketRealtimeResponseBody) String() string            { return proto.CompactTextString(m) }
func (*WebSocketRealtimeResponseBody) ProtoMessage()               {}
func (*WebSocketRealtimeResponseBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WebSocketRealtimeResponseBody) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type WebSocketRealtimeResponse struct {
	Header *WebSocketResponseHeader       `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Body   *WebSocketRealtimeResponseBody `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *WebSocketRealtimeResponse) Reset()                    { *m = WebSocketRealtimeResponse{} }
func (m *WebSocketRealtimeResponse) String() string            { return proto.CompactTextString(m) }
func (*WebSocketRealtimeResponse) ProtoMessage()               {}
func (*WebSocketRealtimeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WebSocketRealtimeResponse) GetHeader() *WebSocketResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *WebSocketRealtimeResponse) GetBody() *WebSocketRealtimeResponseBody {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*WebSocketResponseHeader)(nil), "grpc.WebSocketResponseHeader")
	proto.RegisterType((*WebSocketRealtimeResponseBody)(nil), "grpc.WebSocketRealtimeResponseBody")
	proto.RegisterType((*WebSocketRealtimeResponse)(nil), "grpc.WebSocketRealtimeResponse")
	proto.RegisterEnum("grpc.WebSocketResponseHeader_WebSocketEventType", WebSocketResponseHeader_WebSocketEventType_name, WebSocketResponseHeader_WebSocketEventType_value)
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0xed, 0x9c, 0x83, 0x1d, 0x51, 0x46, 0x10, 0xdc, 0x2e, 0x8a, 0xa3, 0xf3, 0xc2, 0xab,
	0x22, 0x93, 0xe1, 0x75, 0x4d, 0x23, 0x16, 0x34, 0x95, 0x34, 0xc5, 0xcb, 0xd2, 0xae, 0x87, 0x29,
	0xce, 0xa5, 0x74, 0x75, 0xd2, 0x47, 0x10, 0x9f, 0xc1, 0x77, 0x15, 0xb3, 0x96, 0x0d, 0xcb, 0xbc,
	0x4c, 0xfe, 0xef, 0x3b, 0xe4, 0xfc, 0x81, 0xa3, 0x59, 0x5c, 0xe0, 0x47, 0x5c, 0xda, 0x59, 0xae,
	0x0a, 0x45, 0xda, 0xb3, 0x3c, 0x9b, 0x5a, 0x9f, 0x2d, 0x38, 0x7d, 0xc2, 0x24, 0x50, 0xd3, 0x57,
	0x2c, 0x04, 0x2e, 0x33, 0xb5, 0x58, 0xe2, 0x1d, 0xc6, 0x29, 0xe6, 0x84, 0x43, 0x17, 0x57, 0xb8,
	0x28, 0x64, 0x99, 0x61, 0xdf, 0x18, 0x1a, 0x17, 0xc7, 0xe3, 0x4b, 0xfb, 0xd7, 0xb2, 0x77, 0x18,
	0x9b, 0x7b, 0x56, 0x7b, 0x62, 0x33, 0xc2, 0xfa, 0x36, 0x80, 0x34, 0x09, 0x62, 0xc2, 0x20, 0x7c,
	0x74, 0x1d, 0xc9, 0x22, 0xe9, 0x4b, 0xe7, 0x3e, 0x72, 0x28, 0x65, 0x41, 0x10, 0x51, 0x3f, 0xe4,
	0xb2, 0xb7, 0x47, 0x46, 0x70, 0x56, 0xc5, 0x34, 0x14, 0x82, 0x71, 0x19, 0x51, 0x9f, 0x73, 0x46,
	0xa5, 0xe7, 0xf3, 0x0a, 0x32, 0xb6, 0x66, 0xd4, 0x10, 0xf7, 0x5d, 0x56, 0xc5, 0x2d, 0x72, 0x0e,
	0xc3, 0x3f, 0xf1, 0x83, 0x13, 0x48, 0x26, 0x22, 0xcf, 0x65, 0x5c, 0x7a, 0xb7, 0x1e, 0x13, 0xbd,
	0x7d, 0x6b, 0x02, 0xe6, 0xd6, 0x62, 0xf1, 0xbc, 0x78, 0x79, 0xc3, 0x7a, 0xc1, 0x1b, 0x95, 0x96,
	0xe4, 0x04, 0x0e, 0x56, 0xf1, 0xfc, 0x7d, 0x5d, 0x46, 0x57, 0xac, 0x0f, 0xd6, 0x97, 0x01, 0x83,
	0x9d, 0x1e, 0x99, 0x40, 0xe7, 0x59, 0x97, 0xa3, 0xa5, 0xc3, 0xb1, 0xf9, 0x6f, 0x83, 0xa2, 0x82,
	0xc9, 0x35, 0xb4, 0x13, 0x95, 0x96, 0xfd, 0x96, 0x96, 0x46, 0x0d, 0xa9, 0xf9, 0x3a, 0xa1, 0x85,
	0xa4, 0xa3, 0x7f, 0xf7, 0xea, 0x27, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xd5, 0x5e, 0x33, 0xee, 0x01,
	0x00, 0x00,
}
