// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enum/enumDay.proto

package enumPb

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DayofTheWeek int32

const (
	DayofTheWeek_MONDAY    DayofTheWeek = 0
	DayofTheWeek_TUESDAY   DayofTheWeek = 1
	DayofTheWeek_WEDNESDAY DayofTheWeek = 2
	DayofTheWeek_THURSDAY  DayofTheWeek = 3
	DayofTheWeek_FRIDAY    DayofTheWeek = 4
	DayofTheWeek_SATURDAY  DayofTheWeek = 5
	DayofTheWeek_SUNDAY    DayofTheWeek = 6
)

var DayofTheWeek_name = map[int32]string{
	0: "MONDAY",
	1: "TUESDAY",
	2: "WEDNESDAY",
	3: "THURSDAY",
	4: "FRIDAY",
	5: "SATURDAY",
	6: "SUNDAY",
}

var DayofTheWeek_value = map[string]int32{
	"MONDAY":    0,
	"TUESDAY":   1,
	"WEDNESDAY": 2,
	"THURSDAY":  3,
	"FRIDAY":    4,
	"SATURDAY":  5,
	"SUNDAY":    6,
}

func (x DayofTheWeek) String() string {
	return proto.EnumName(DayofTheWeek_name, int32(x))
}

func (DayofTheWeek) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce8da65676e1a44d, []int{0}
}

type EnumDay struct {
	Id                   int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DayOfTheWeek         DayofTheWeek `protobuf:"varint,2,opt,name=day_of_the_week,json=dayOfTheWeek,proto3,enum=enum.protobuf.DayofTheWeek" json:"day_of_the_week,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EnumDay) Reset()         { *m = EnumDay{} }
func (m *EnumDay) String() string { return proto.CompactTextString(m) }
func (*EnumDay) ProtoMessage()    {}
func (*EnumDay) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce8da65676e1a44d, []int{0}
}
func (m *EnumDay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumDay.Unmarshal(m, b)
}
func (m *EnumDay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumDay.Marshal(b, m, deterministic)
}
func (m *EnumDay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumDay.Merge(m, src)
}
func (m *EnumDay) XXX_Size() int {
	return xxx_messageInfo_EnumDay.Size(m)
}
func (m *EnumDay) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumDay.DiscardUnknown(m)
}

var xxx_messageInfo_EnumDay proto.InternalMessageInfo

func (m *EnumDay) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EnumDay) GetDayOfTheWeek() DayofTheWeek {
	if m != nil {
		return m.DayOfTheWeek
	}
	return DayofTheWeek_MONDAY
}

func init() {
	proto.RegisterEnum("enum.protobuf.DayofTheWeek", DayofTheWeek_name, DayofTheWeek_value)
	proto.RegisterType((*EnumDay)(nil), "enum.protobuf.EnumDay")
}

func init() { proto.RegisterFile("enum/enumDay.proto", fileDescriptor_ce8da65676e1a44d) }

var fileDescriptor_ce8da65676e1a44d = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xcd, 0x2b, 0xcd,
	0xd5, 0x07, 0x11, 0x2e, 0x89, 0x95, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x20, 0x2e,
	0x84, 0x9d, 0x54, 0x9a, 0xa6, 0x14, 0xcb, 0xc5, 0xee, 0x0a, 0x91, 0x17, 0xe2, 0xe3, 0x62, 0xca,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xca, 0x4c, 0x11, 0x72, 0xe2, 0xe2, 0x4f,
	0x49, 0xac, 0x8c, 0xcf, 0x4f, 0x8b, 0x2f, 0xc9, 0x48, 0x8d, 0x2f, 0x4f, 0x4d, 0xcd, 0x96, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x33, 0x92, 0xd6, 0x43, 0x31, 0x43, 0xcf, 0x25, 0xb1, 0x32, 0x3f, 0x2d,
	0x24, 0x23, 0x35, 0x3c, 0x35, 0x35, 0x3b, 0x88, 0x27, 0x25, 0xb1, 0xd2, 0x1f, 0xc6, 0xd3, 0xca,
	0xe2, 0xe2, 0x41, 0x96, 0x15, 0xe2, 0xe2, 0x62, 0xf3, 0xf5, 0xf7, 0x73, 0x71, 0x8c, 0x14, 0x60,
	0x10, 0xe2, 0xe6, 0x62, 0x0f, 0x09, 0x75, 0x0d, 0x06, 0x71, 0x18, 0x85, 0x78, 0xb9, 0x38, 0xc3,
	0x5d, 0x5d, 0xfc, 0x20, 0x5c, 0x26, 0x21, 0x1e, 0x2e, 0x8e, 0x10, 0x8f, 0xd0, 0x20, 0x30, 0x8f,
	0x19, 0xa4, 0xcb, 0x2d, 0xc8, 0x13, 0xc4, 0x66, 0x01, 0xc9, 0x04, 0x3b, 0x86, 0x84, 0x06, 0x81,
	0x78, 0xac, 0x20, 0x99, 0xe0, 0x50, 0xb0, 0x79, 0x6c, 0x4e, 0x1c, 0x51, 0x6c, 0x20, 0x77, 0x05,
	0x24, 0x25, 0xb1, 0x81, 0x9d, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x58, 0xc3, 0x7e, 0x8b,
	0x00, 0x01, 0x00, 0x00,
}
