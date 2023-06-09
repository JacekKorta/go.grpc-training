// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job-data.proto

package jobs

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

type StandardRequest struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	ExpLevel             string   `protobuf:"bytes,2,opt,name=exp_level,json=expLevel,proto3" json:"exp_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StandardRequest) Reset()         { *m = StandardRequest{} }
func (m *StandardRequest) String() string { return proto.CompactTextString(m) }
func (*StandardRequest) ProtoMessage()    {}
func (*StandardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_50c0c0edafd2e172, []int{0}
}

func (m *StandardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StandardRequest.Unmarshal(m, b)
}
func (m *StandardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StandardRequest.Marshal(b, m, deterministic)
}
func (m *StandardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StandardRequest.Merge(m, src)
}
func (m *StandardRequest) XXX_Size() int {
	return xxx_messageInfo_StandardRequest.Size(m)
}
func (m *StandardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StandardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StandardRequest proto.InternalMessageInfo

func (m *StandardRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *StandardRequest) GetExpLevel() string {
	if m != nil {
		return m.ExpLevel
	}
	return ""
}

type CountResponse struct {
	TotalOfferCount      int32    `protobuf:"varint,1,opt,name=total_offer_count,json=totalOfferCount,proto3" json:"total_offer_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountResponse) Reset()         { *m = CountResponse{} }
func (m *CountResponse) String() string { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()    {}
func (*CountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50c0c0edafd2e172, []int{1}
}

func (m *CountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountResponse.Unmarshal(m, b)
}
func (m *CountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountResponse.Marshal(b, m, deterministic)
}
func (m *CountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountResponse.Merge(m, src)
}
func (m *CountResponse) XXX_Size() int {
	return xxx_messageInfo_CountResponse.Size(m)
}
func (m *CountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountResponse proto.InternalMessageInfo

func (m *CountResponse) GetTotalOfferCount() int32 {
	if m != nil {
		return m.TotalOfferCount
	}
	return 0
}

type SalaryResponse struct {
	Min                  int64    `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max                  int64    `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	Avg                  float32  `protobuf:"fixed32,3,opt,name=avg,proto3" json:"avg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SalaryResponse) Reset()         { *m = SalaryResponse{} }
func (m *SalaryResponse) String() string { return proto.CompactTextString(m) }
func (*SalaryResponse) ProtoMessage()    {}
func (*SalaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50c0c0edafd2e172, []int{2}
}

func (m *SalaryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SalaryResponse.Unmarshal(m, b)
}
func (m *SalaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SalaryResponse.Marshal(b, m, deterministic)
}
func (m *SalaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SalaryResponse.Merge(m, src)
}
func (m *SalaryResponse) XXX_Size() int {
	return xxx_messageInfo_SalaryResponse.Size(m)
}
func (m *SalaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SalaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SalaryResponse proto.InternalMessageInfo

func (m *SalaryResponse) GetMin() int64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *SalaryResponse) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *SalaryResponse) GetAvg() float32 {
	if m != nil {
		return m.Avg
	}
	return 0
}

func init() {
	proto.RegisterType((*StandardRequest)(nil), "StandardRequest")
	proto.RegisterType((*CountResponse)(nil), "CountResponse")
	proto.RegisterType((*SalaryResponse)(nil), "SalaryResponse")
}

func init() { proto.RegisterFile("job-data.proto", fileDescriptor_50c0c0edafd2e172) }

var fileDescriptor_50c0c0edafd2e172 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xfb, 0x87, 0x96, 0x76, 0xc0, 0xa4, 0xee, 0x29, 0xd4, 0x4b, 0xc9, 0xa9, 0x28, 0x46,
	0xd0, 0xa3, 0x07, 0xa1, 0x8a, 0x42, 0x11, 0x84, 0xcd, 0xcd, 0x4b, 0x98, 0x34, 0xd3, 0x92, 0x12,
	0x33, 0x71, 0x33, 0x0d, 0x89, 0x9f, 0x5e, 0x76, 0x0b, 0x85, 0x8a, 0xb7, 0x99, 0xdf, 0xdb, 0xf7,
	0xf6, 0xed, 0x82, 0xb7, 0xe7, 0xf4, 0x36, 0x43, 0xc1, 0xa8, 0x32, 0x2c, 0x1c, 0xae, 0xc1, 0x8f,
	0x05, 0xcb, 0x0c, 0x4d, 0xa6, 0xe9, 0xfb, 0x40, 0xb5, 0xa8, 0x39, 0x4c, 0x36, 0x28, 0xb4, 0x63,
	0xd3, 0x05, 0xfd, 0x45, 0x7f, 0x39, 0xd5, 0xa7, 0x5d, 0x5d, 0xc1, 0x94, 0xda, 0x2a, 0x29, 0xa8,
	0xa1, 0x22, 0x18, 0x1c, 0x45, 0x6a, 0xab, 0x77, 0xbb, 0x87, 0x8f, 0x70, 0xf1, 0xcc, 0x87, 0x52,
	0x34, 0xd5, 0x15, 0x97, 0x35, 0xa9, 0x6b, 0xb8, 0x14, 0x16, 0x2c, 0x12, 0xde, 0x6e, 0xc9, 0x24,
	0x1b, 0x2b, 0xba, 0xc8, 0x91, 0xf6, 0x9d, 0xf0, 0x61, 0xb9, 0xf3, 0x84, 0xaf, 0xe0, 0xc5, 0x58,
	0xa0, 0xe9, 0x4e, 0xee, 0x19, 0x0c, 0xbf, 0xf2, 0xd2, 0x9d, 0x1f, 0x6a, 0x3b, 0x3a, 0x82, 0xad,
	0xbb, 0xd7, 0x12, 0x6c, 0x2d, 0xc1, 0x66, 0x17, 0x0c, 0x17, 0xfd, 0xe5, 0x40, 0xdb, 0xf1, 0xfe,
	0x07, 0x60, 0xcd, 0xe9, 0x0b, 0x0a, 0xc6, 0xa6, 0x51, 0x37, 0x30, 0x72, 0xf1, 0x6a, 0x16, 0xfd,
	0x79, 0xe6, 0xdc, 0x8b, 0xce, 0xca, 0x86, 0x3d, 0xf5, 0x04, 0xc1, 0x1b, 0xc9, 0xb1, 0x85, 0x0d,
	0x58, 0x75, 0x31, 0x95, 0x39, 0x9b, 0x5c, 0xba, 0x7f, 0xfc, 0x7e, 0x74, 0xde, 0x37, 0xec, 0xad,
	0x26, 0x9f, 0xe3, 0xe8, 0x6e, 0xcf, 0x69, 0x9d, 0x8e, 0xdd, 0xef, 0x3e, 0xfc, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x99, 0x6d, 0x90, 0xe1, 0x6f, 0x01, 0x00, 0x00,
}
