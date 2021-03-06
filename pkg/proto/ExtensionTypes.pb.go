// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ExtensionTypes.proto

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

type ExtensionType int32

const (
	ExtensionType_NONE  ExtensionType = 0
	ExtensionType_RELAY ExtensionType = 1
	ExtensionType_TLM   ExtensionType = 2
	ExtensionType_TRACK ExtensionType = 5
	ExtensionType_BAND  ExtensionType = 6
)

var ExtensionType_name = map[int32]string{
	0: "NONE",
	1: "RELAY",
	2: "TLM",
	5: "TRACK",
	6: "BAND",
}

var ExtensionType_value = map[string]int32{
	"NONE":  0,
	"RELAY": 1,
	"TLM":   2,
	"TRACK": 5,
	"BAND":  6,
}

func (x ExtensionType) Enum() *ExtensionType {
	p := new(ExtensionType)
	*p = x
	return p
}

func (x ExtensionType) String() string {
	return proto.EnumName(ExtensionType_name, int32(x))
}

func (x *ExtensionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExtensionType_value, data, "ExtensionType")
	if err != nil {
		return err
	}
	*x = ExtensionType(value)
	return nil
}

func (ExtensionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3665e286bd22a0bf, []int{0}
}

func init() {
	proto.RegisterEnum("ExtensionType", ExtensionType_name, ExtensionType_value)
}

func init() { proto.RegisterFile("ExtensionTypes.proto", fileDescriptor_3665e286bd22a0bf) }

var fileDescriptor_3665e286bd22a0bf = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x71, 0xad, 0x28, 0x49,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x0b, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0xd7, 0x72, 0xe2, 0xe2, 0x45, 0x11, 0x17, 0xe2, 0xe0, 0x62, 0xf1, 0xf3, 0xf7, 0x73, 0x15, 0x60,
	0x10, 0xe2, 0xe4, 0x62, 0x0d, 0x72, 0xf5, 0x71, 0x8c, 0x14, 0x60, 0x14, 0x62, 0xe7, 0x62, 0x0e,
	0xf1, 0xf1, 0x15, 0x60, 0x02, 0x89, 0x85, 0x04, 0x39, 0x3a, 0x7b, 0x0b, 0xb0, 0x82, 0x14, 0x3a,
	0x39, 0xfa, 0xb9, 0x08, 0xb0, 0x39, 0x89, 0x46, 0x09, 0x3b, 0xa6, 0x94, 0xa5, 0x16, 0x95, 0x64,
	0x16, 0xa7, 0xe6, 0xa6, 0xe6, 0x95, 0x04, 0x24, 0x26, 0x67, 0xa7, 0x96, 0x00, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xed, 0x42, 0x75, 0xa2, 0x71, 0x00, 0x00, 0x00,
}
