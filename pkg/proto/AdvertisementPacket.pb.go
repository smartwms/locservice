// Code generated by protoc-gen-go. DO NOT EDIT.
// source: AdvertisementPacket.proto

package AdvertisementPacket

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

type AdvertisementPacket struct {
	Address    []byte  `protobuf:"bytes,10,req,name=address" json:"address,omitempty"`
	Rssi       *int32  `protobuf:"varint,11,req,name=rssi" json:"rssi,omitempty"`
	Channel    *uint32 `protobuf:"varint,12,req,name=channel" json:"channel,omitempty"`
	Timestamp  *uint32 `protobuf:"varint,13,req,name=timestamp" json:"timestamp,omitempty"`
	TimeMillis *uint32 `protobuf:"varint,14,opt,name=time_millis,json=timeMillis,def=0" json:"time_millis,omitempty"`
	TimeNanos  *uint32 `protobuf:"varint,15,opt,name=time_nanos,json=timeNanos,def=0" json:"time_nanos,omitempty"`
	// Type field is used to identify packet types.
	Type                         *ExtensionType `protobuf:"varint,16,opt,name=type,enum=ExtensionType,def=0" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}       `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *AdvertisementPacket) Reset()         { *m = AdvertisementPacket{} }
func (m *AdvertisementPacket) String() string { return proto.CompactTextString(m) }
func (*AdvertisementPacket) ProtoMessage()    {}
func (*AdvertisementPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8375302d2df4ca4, []int{0}
}

var extRange_AdvertisementPacket = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*AdvertisementPacket) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_AdvertisementPacket
}

func (m *AdvertisementPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdvertisementPacket.Unmarshal(m, b)
}
func (m *AdvertisementPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdvertisementPacket.Marshal(b, m, deterministic)
}
func (m *AdvertisementPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdvertisementPacket.Merge(m, src)
}
func (m *AdvertisementPacket) XXX_Size() int {
	return xxx_messageInfo_AdvertisementPacket.Size(m)
}
func (m *AdvertisementPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_AdvertisementPacket.DiscardUnknown(m)
}

var xxx_messageInfo_AdvertisementPacket proto.InternalMessageInfo

const Default_AdvertisementPacket_TimeMillis uint32 = 0
const Default_AdvertisementPacket_TimeNanos uint32 = 0
const Default_AdvertisementPacket_Type ExtensionType = ExtensionType_NONE

func (m *AdvertisementPacket) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AdvertisementPacket) GetRssi() int32 {
	if m != nil && m.Rssi != nil {
		return *m.Rssi
	}
	return 0
}

func (m *AdvertisementPacket) GetChannel() uint32 {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return 0
}

func (m *AdvertisementPacket) GetTimestamp() uint32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *AdvertisementPacket) GetTimeMillis() uint32 {
	if m != nil && m.TimeMillis != nil {
		return *m.TimeMillis
	}
	return Default_AdvertisementPacket_TimeMillis
}

func (m *AdvertisementPacket) GetTimeNanos() uint32 {
	if m != nil && m.TimeNanos != nil {
		return *m.TimeNanos
	}
	return Default_AdvertisementPacket_TimeNanos
}

func (m *AdvertisementPacket) GetType() ExtensionType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_AdvertisementPacket_Type
}

func init() {
	proto.RegisterType((*AdvertisementPacket)(nil), "AdvertisementPacket")
}

func init() { proto.RegisterFile("AdvertisementPacket.proto", fileDescriptor_c8375302d2df4ca4) }

var fileDescriptor_c8375302d2df4ca4 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x49, 0xa9, 0x58, 0x67, 0xb7, 0xb5, 0x44, 0x85, 0x28, 0x1e, 0xc2, 0x9e, 0x82, 0x87,
	0x45, 0x3c, 0xee, 0x4d, 0x61, 0x2f, 0x82, 0x55, 0x82, 0x27, 0x2f, 0x4b, 0xd8, 0x0e, 0x18, 0x6c,
	0xd3, 0xd2, 0x09, 0xe2, 0xde, 0xfa, 0xbf, 0xbd, 0x48, 0x22, 0x45, 0x84, 0xde, 0xe6, 0xfb, 0xe6,
	0xcd, 0xc0, 0x83, 0xcb, 0xfb, 0xfa, 0x13, 0x07, 0x6f, 0x09, 0x5b, 0x74, 0xfe, 0xc5, 0xec, 0x3f,
	0xd0, 0xaf, 0xfb, 0xa1, 0xf3, 0xdd, 0xd5, 0xf9, 0xf6, 0xcb, 0xa3, 0x23, 0xdb, 0xb9, 0xd7, 0x43,
	0x8f, 0xf4, 0x6b, 0x57, 0xdf, 0x0c, 0xce, 0x66, 0x6e, 0xb8, 0x80, 0x63, 0x53, 0xd7, 0x03, 0x12,
	0x09, 0x90, 0x89, 0x5a, 0xea, 0x09, 0x39, 0x87, 0x74, 0x20, 0xb2, 0x62, 0x21, 0x13, 0x75, 0xa4,
	0xe3, 0x1c, 0xd2, 0xfb, 0x77, 0xe3, 0x1c, 0x36, 0x62, 0x29, 0x13, 0x95, 0xeb, 0x09, 0xf9, 0x35,
	0x9c, 0x78, 0xdb, 0x22, 0x79, 0xd3, 0xf6, 0x22, 0x8f, 0xbb, 0x3f, 0xc1, 0x57, 0xb0, 0x08, 0xb0,
	0x6b, 0x6d, 0xd3, 0x58, 0x12, 0x85, 0x64, 0x2a, 0xdf, 0xb0, 0x5b, 0x0d, 0xc1, 0x3e, 0x45, 0xc9,
	0x25, 0x44, 0xda, 0x39, 0xe3, 0x3a, 0x12, 0xa7, 0x53, 0x24, 0x7e, 0xa9, 0x82, 0xe3, 0x0a, 0x52,
	0x7f, 0xe8, 0x51, 0x94, 0x92, 0xa9, 0xe2, 0xae, 0x58, 0xff, 0x2b, 0xba, 0x49, 0xab, 0xe7, 0x6a,
	0xab, 0x63, 0xe2, 0x26, 0xcb, 0xea, 0x72, 0x1c, 0xc7, 0x31, 0x79, 0x4c, 0x33, 0x56, 0xc2, 0xc3,
	0xc5, 0xdb, 0x5c, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x29, 0xaa, 0x83, 0x46, 0x01,
	0x00, 0x00,
}