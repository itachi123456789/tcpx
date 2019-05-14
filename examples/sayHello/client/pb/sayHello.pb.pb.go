// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sayHello.pb

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SayHelloRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloRequest) Reset()         { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()    {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3858ea14d1cc40e, []int{0}
}

func (m *SayHelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloRequest.Unmarshal(m, b)
}
func (m *SayHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloRequest.Marshal(b, m, deterministic)
}
func (m *SayHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloRequest.Merge(m, src)
}
func (m *SayHelloRequest) XXX_Size() int {
	return xxx_messageInfo_SayHelloRequest.Size(m)
}
func (m *SayHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloRequest proto.InternalMessageInfo

func (m *SayHelloRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type SayHelloReponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloReponse) Reset()         { *m = SayHelloReponse{} }
func (m *SayHelloReponse) String() string { return proto.CompactTextString(m) }
func (*SayHelloReponse) ProtoMessage()    {}
func (*SayHelloReponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3858ea14d1cc40e, []int{1}
}

func (m *SayHelloReponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloReponse.Unmarshal(m, b)
}
func (m *SayHelloReponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloReponse.Marshal(b, m, deterministic)
}
func (m *SayHelloReponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloReponse.Merge(m, src)
}
func (m *SayHelloReponse) XXX_Size() int {
	return xxx_messageInfo_SayHelloReponse.Size(m)
}
func (m *SayHelloReponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloReponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloReponse proto.InternalMessageInfo

func (m *SayHelloReponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloRequest)(nil), "pb.SayHelloRequest")
	proto.RegisterType((*SayHelloReponse)(nil), "pb.SayHelloReponse")
}

func init() { proto.RegisterFile("sayHello.pb", fileDescriptor_b3858ea14d1cc40e) }

var fileDescriptor_b3858ea14d1cc40e = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4e, 0xac, 0xf4,
	0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x2b, 0x48, 0x12, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe5, 0xe2, 0x0f,
	0x86, 0x0a, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x71, 0x71, 0x94, 0x16, 0xa7,
	0x16, 0xe5, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x4a, 0xda,
	0xc8, 0xca, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13,
	0xd3, 0x61, 0xaa, 0x61, 0xdc, 0x24, 0xb6, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf6, 0x56, 0x5f, 0x61, 0x75, 0x00, 0x00, 0x00,
}