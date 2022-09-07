// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package test

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NonSemantic struct {
	Semantic             *NonSemantic_Semantic `protobuf:"bytes,1,opt,name=semantic,proto3" json:"semantic,omitempty"`
	OverruledSemantic    *NonSemantic_Semantic `protobuf:"bytes,2,opt,name=overruled_semantic,json=overruledSemantic,proto3" json:"overruled_semantic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NonSemantic) Reset()         { *m = NonSemantic{} }
func (m *NonSemantic) String() string { return proto.CompactTextString(m) }
func (*NonSemantic) ProtoMessage()    {}
func (*NonSemantic) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *NonSemantic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonSemantic.Unmarshal(m, b)
}
func (m *NonSemantic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonSemantic.Marshal(b, m, deterministic)
}
func (m *NonSemantic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonSemantic.Merge(m, src)
}
func (m *NonSemantic) XXX_Size() int {
	return xxx_messageInfo_NonSemantic.Size(m)
}
func (m *NonSemantic) XXX_DiscardUnknown() {
	xxx_messageInfo_NonSemantic.DiscardUnknown(m)
}

var xxx_messageInfo_NonSemantic proto.InternalMessageInfo

func (m *NonSemantic) GetSemantic() *NonSemantic_Semantic {
	if m != nil {
		return m.Semantic
	}
	return nil
}

func (m *NonSemantic) GetOverruledSemantic() *NonSemantic_Semantic {
	if m != nil {
		return m.OverruledSemantic
	}
	return nil
}

type NonSemantic_Semantic struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonSemantic_Semantic) Reset()         { *m = NonSemantic_Semantic{} }
func (m *NonSemantic_Semantic) String() string { return proto.CompactTextString(m) }
func (*NonSemantic_Semantic) ProtoMessage()    {}
func (*NonSemantic_Semantic) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}
func (m *NonSemantic_Semantic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonSemantic_Semantic.Unmarshal(m, b)
}
func (m *NonSemantic_Semantic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonSemantic_Semantic.Marshal(b, m, deterministic)
}
func (m *NonSemantic_Semantic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonSemantic_Semantic.Merge(m, src)
}
func (m *NonSemantic_Semantic) XXX_Size() int {
	return xxx_messageInfo_NonSemantic_Semantic.Size(m)
}
func (m *NonSemantic_Semantic) XXX_DiscardUnknown() {
	xxx_messageInfo_NonSemantic_Semantic.DiscardUnknown(m)
}

var xxx_messageInfo_NonSemantic_Semantic proto.InternalMessageInfo

type OneOf struct {
	// Types that are valid to be assigned to Option:
	//
	//	*OneOf_Semantic_
	//	*OneOf_NonSemantic_
	Option               isOneOf_Option `protobuf_oneof:"option"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OneOf) Reset()         { *m = OneOf{} }
func (m *OneOf) String() string { return proto.CompactTextString(m) }
func (*OneOf) ProtoMessage()    {}
func (*OneOf) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}
func (m *OneOf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneOf.Unmarshal(m, b)
}
func (m *OneOf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneOf.Marshal(b, m, deterministic)
}
func (m *OneOf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneOf.Merge(m, src)
}
func (m *OneOf) XXX_Size() int {
	return xxx_messageInfo_OneOf.Size(m)
}
func (m *OneOf) XXX_DiscardUnknown() {
	xxx_messageInfo_OneOf.DiscardUnknown(m)
}

var xxx_messageInfo_OneOf proto.InternalMessageInfo

type isOneOf_Option interface {
	isOneOf_Option()
}

type OneOf_Semantic_ struct {
	Semantic *OneOf_Semantic `protobuf:"bytes,1,opt,name=semantic,proto3,oneof" json:"semantic,omitempty"`
}
type OneOf_NonSemantic_ struct {
	NonSemantic *OneOf_NonSemantic `protobuf:"bytes,2,opt,name=non_semantic,json=nonSemantic,proto3,oneof" json:"non_semantic,omitempty"`
}

func (*OneOf_Semantic_) isOneOf_Option()    {}
func (*OneOf_NonSemantic_) isOneOf_Option() {}

func (m *OneOf) GetOption() isOneOf_Option {
	if m != nil {
		return m.Option
	}
	return nil
}

func (m *OneOf) GetSemantic() *OneOf_Semantic {
	if x, ok := m.GetOption().(*OneOf_Semantic_); ok {
		return x.Semantic
	}
	return nil
}

func (m *OneOf) GetNonSemantic() *OneOf_NonSemantic {
	if x, ok := m.GetOption().(*OneOf_NonSemantic_); ok {
		return x.NonSemantic
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OneOf) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OneOf_Semantic_)(nil),
		(*OneOf_NonSemantic_)(nil),
	}
}

type OneOf_Semantic struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OneOf_Semantic) Reset()         { *m = OneOf_Semantic{} }
func (m *OneOf_Semantic) String() string { return proto.CompactTextString(m) }
func (*OneOf_Semantic) ProtoMessage()    {}
func (*OneOf_Semantic) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1, 0}
}
func (m *OneOf_Semantic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneOf_Semantic.Unmarshal(m, b)
}
func (m *OneOf_Semantic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneOf_Semantic.Marshal(b, m, deterministic)
}
func (m *OneOf_Semantic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneOf_Semantic.Merge(m, src)
}
func (m *OneOf_Semantic) XXX_Size() int {
	return xxx_messageInfo_OneOf_Semantic.Size(m)
}
func (m *OneOf_Semantic) XXX_DiscardUnknown() {
	xxx_messageInfo_OneOf_Semantic.DiscardUnknown(m)
}

var xxx_messageInfo_OneOf_Semantic proto.InternalMessageInfo

type OneOf_NonSemantic struct {
	Boolean              bool     `protobuf:"varint,1,opt,name=boolean,proto3" json:"boolean,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OneOf_NonSemantic) Reset()         { *m = OneOf_NonSemantic{} }
func (m *OneOf_NonSemantic) String() string { return proto.CompactTextString(m) }
func (*OneOf_NonSemantic) ProtoMessage()    {}
func (*OneOf_NonSemantic) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1, 1}
}
func (m *OneOf_NonSemantic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneOf_NonSemantic.Unmarshal(m, b)
}
func (m *OneOf_NonSemantic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneOf_NonSemantic.Marshal(b, m, deterministic)
}
func (m *OneOf_NonSemantic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneOf_NonSemantic.Merge(m, src)
}
func (m *OneOf_NonSemantic) XXX_Size() int {
	return xxx_messageInfo_OneOf_NonSemantic.Size(m)
}
func (m *OneOf_NonSemantic) XXX_DiscardUnknown() {
	xxx_messageInfo_OneOf_NonSemantic.DiscardUnknown(m)
}

var xxx_messageInfo_OneOf_NonSemantic proto.InternalMessageInfo

func (m *OneOf_NonSemantic) GetBoolean() bool {
	if m != nil {
		return m.Boolean
	}
	return false
}

func init() {
	proto.RegisterType((*NonSemantic)(nil), "thethings.flags.test.NonSemantic")
	proto.RegisterType((*NonSemantic_Semantic)(nil), "thethings.flags.test.NonSemantic.Semantic")
	proto.RegisterType((*OneOf)(nil), "thethings.flags.test.OneOf")
	proto.RegisterType((*OneOf_Semantic)(nil), "thethings.flags.test.OneOf.Semantic")
	proto.RegisterType((*OneOf_NonSemantic)(nil), "thethings.flags.test.OneOf.NonSemantic")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0x4e, 0xca, 0xef, 0x17, 0x97, 0xad, 0x82, 0x0d, 0x1e, 0xda, 0x80, 0x50, 0x8a, 0x50, 0x11,
	0xba, 0x11, 0x3d, 0x88, 0x1e, 0x73, 0xaa, 0x20, 0x2d, 0xc4, 0x9e, 0xbc, 0xc8, 0x36, 0x9d, 0x6e,
	0x02, 0xe9, 0x4c, 0xc9, 0x6e, 0x7d, 0xa8, 0x3e, 0x92, 0x67, 0x9f, 0xa0, 0x4f, 0x20, 0xdd, 0x9a,
	0x18, 0xb5, 0x08, 0xde, 0xbe, 0xd9, 0x99, 0xfd, 0xfe, 0xf0, 0xf1, 0xa3, 0x05, 0x68, 0x2d, 0x15,
	0x88, 0x65, 0x41, 0x86, 0xfc, 0x13, 0x93, 0x82, 0x49, 0x33, 0x54, 0x5a, 0xcc, 0x73, 0xa9, 0xb4,
	0x30, 0xa0, 0x4d, 0xd0, 0x92, 0x88, 0x64, 0xa4, 0xc9, 0x08, 0xf5, 0xee, 0x30, 0x38, 0xcd, 0xd0,
	0x40, 0x81, 0x32, 0x0f, 0x15, 0x29, 0xb2, 0x6f, 0x16, 0xed, 0xd6, 0xbd, 0x57, 0x97, 0x37, 0x47,
	0x84, 0x8f, 0xb0, 0x90, 0x68, 0xb2, 0xc4, 0x1f, 0x71, 0xa6, 0x3f, 0x70, 0xdb, 0xed, 0xba, 0xe7,
	0xcd, 0xab, 0x0b, 0xb1, 0x4f, 0x4a, 0xd4, 0x3e, 0x89, 0x12, 0x44, 0xde, 0x66, 0xdd, 0x69, 0x30,
	0x27, 0xae, 0x38, 0xfc, 0x84, 0xfb, 0xf4, 0x02, 0x45, 0xb1, 0xca, 0x61, 0xf6, 0x5c, 0x31, 0x37,
	0xfe, 0xcc, 0xcc, 0x36, 0xeb, 0xce, 0x3f, 0xe6, 0x5c, 0x3a, 0x71, 0xab, 0xe2, 0x2b, 0x97, 0x01,
	0xe7, 0xac, 0xc4, 0x77, 0x7c, 0xb3, 0xee, 0x78, 0xcc, 0x39, 0x76, 0xbb, 0x6e, 0xef, 0xcd, 0xe5,
	0xff, 0xc7, 0x08, 0xe3, 0xb9, 0x1f, 0xfd, 0x88, 0x75, 0xb6, 0x5f, 0xdc, 0x9e, 0x57, 0xb2, 0xc3,
	0x7a, 0x94, 0x07, 0x7e, 0x88, 0x84, 0xdf, 0x43, 0xf4, 0x7f, 0xe3, 0xa9, 0x45, 0x19, 0x3a, 0x71,
	0x13, 0x3f, 0xc7, 0xba, 0xe7, 0xa0, 0xff, 0xb5, 0x83, 0x36, 0x3f, 0x98, 0x12, 0xe5, 0x20, 0xd1,
	0x7a, 0x65, 0x71, 0x39, 0xd6, 0xc3, 0x45, 0x8c, 0x7b, 0xb4, 0xdc, 0x36, 0x1d, 0xdd, 0x3e, 0xdd,
	0xa8, 0xcc, 0xa4, 0xab, 0xa9, 0x48, 0x68, 0x11, 0x4e, 0x52, 0x98, 0x58, 0x3b, 0xf7, 0x38, 0x5b,
	0x69, 0x53, 0x64, 0xa0, 0x43, 0xdb, 0x75, 0x32, 0x50, 0x80, 0x03, 0x45, 0x03, 0x6b, 0x33, 0xdc,
	0xda, 0x9c, 0x7a, 0x76, 0x73, 0xfd, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x20, 0xfe, 0x6a, 0xd3, 0x5e,
	0x02, 0x00, 0x00,
}
