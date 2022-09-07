// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enums.proto

package test

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
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

type RegularEnum int32

const (
	REGULAR_UNKNOWN RegularEnum = 0
	REGULAR_A       RegularEnum = 1
	REGULAR_B       RegularEnum = 2
)

var RegularEnum_name = map[int32]string{
	0: "REGULAR_UNKNOWN",
	1: "REGULAR_A",
	2: "REGULAR_B",
}

var RegularEnum_value = map[string]int32{
	"REGULAR_UNKNOWN": 0,
	"REGULAR_A":       1,
	"REGULAR_B":       2,
}

func (RegularEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{0}
}

type CustomEnum int32

const (
	CustomEnum_CUSTOM_UNKNOWN CustomEnum = 0
	CustomEnum_CUSTOM_V1_0    CustomEnum = 1
	CustomEnum_CUSTOM_V1_0_1  CustomEnum = 2
)

var CustomEnum_name = map[int32]string{
	0: "CUSTOM_UNKNOWN",
	1: "CUSTOM_V1_0",
	2: "CUSTOM_V1_0_1",
}

var CustomEnum_value = map[string]int32{
	"CUSTOM_UNKNOWN": 0,
	"CUSTOM_V1_0":    1,
	"CUSTOM_V1_0_1":  2,
}

func (CustomEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{1}
}

type MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum int32

const (
	AT_MOST_ONCE  MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum = 0
	AT_LEAST_ONCE MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum = 1
	EXACTLY_ONCE  MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum = 2
)

var MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum_name = map[int32]string{
	0: "AT_MOST_ONCE",
	1: "AT_LEAST_ONCE",
	2: "EXACTLY_ONCE",
}

var MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum_value = map[string]int32{
	"AT_MOST_ONCE":  0,
	"AT_LEAST_ONCE": 1,
	"EXACTLY_ONCE":  2,
}

func (MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{3, 0, 0}
}

type CustomEnumValue struct {
	Value                CustomEnum `protobuf:"varint,1,opt,name=value,proto3,enum=thethings.flags.test.CustomEnum" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CustomEnumValue) Reset()         { *m = CustomEnumValue{} }
func (m *CustomEnumValue) String() string { return proto.CompactTextString(m) }
func (*CustomEnumValue) ProtoMessage()    {}
func (*CustomEnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{0}
}
func (m *CustomEnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomEnumValue.Unmarshal(m, b)
}
func (m *CustomEnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomEnumValue.Marshal(b, m, deterministic)
}
func (m *CustomEnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomEnumValue.Merge(m, src)
}
func (m *CustomEnumValue) XXX_Size() int {
	return xxx_messageInfo_CustomEnumValue.Size(m)
}
func (m *CustomEnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomEnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_CustomEnumValue proto.InternalMessageInfo

func (m *CustomEnumValue) GetValue() CustomEnum {
	if m != nil {
		return m.Value
	}
	return CustomEnum_CUSTOM_UNKNOWN
}

type MessageWithEnums struct {
	Regular              RegularEnum        `protobuf:"varint,1,opt,name=regular,proto3,enum=thethings.flags.test.RegularEnum" json:"regular,omitempty"`
	Regulars             []RegularEnum      `protobuf:"varint,2,rep,packed,name=regulars,proto3,enum=thethings.flags.test.RegularEnum" json:"regulars,omitempty"`
	Custom               CustomEnum         `protobuf:"varint,3,opt,name=custom,proto3,enum=thethings.flags.test.CustomEnum" json:"custom,omitempty"`
	Customs              []CustomEnum       `protobuf:"varint,4,rep,packed,name=customs,proto3,enum=thethings.flags.test.CustomEnum" json:"customs,omitempty"`
	WrappedCustom        *CustomEnumValue   `protobuf:"bytes,5,opt,name=wrapped_custom,json=wrappedCustom,proto3" json:"wrapped_custom,omitempty"`
	WrappedCustoms       []*CustomEnumValue `protobuf:"bytes,6,rep,name=wrapped_customs,json=wrappedCustoms,proto3" json:"wrapped_customs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MessageWithEnums) Reset()         { *m = MessageWithEnums{} }
func (m *MessageWithEnums) String() string { return proto.CompactTextString(m) }
func (*MessageWithEnums) ProtoMessage()    {}
func (*MessageWithEnums) Descriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{1}
}
func (m *MessageWithEnums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithEnums.Unmarshal(m, b)
}
func (m *MessageWithEnums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithEnums.Marshal(b, m, deterministic)
}
func (m *MessageWithEnums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithEnums.Merge(m, src)
}
func (m *MessageWithEnums) XXX_Size() int {
	return xxx_messageInfo_MessageWithEnums.Size(m)
}
func (m *MessageWithEnums) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithEnums.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithEnums proto.InternalMessageInfo

func (m *MessageWithEnums) GetRegular() RegularEnum {
	if m != nil {
		return m.Regular
	}
	return REGULAR_UNKNOWN
}

func (m *MessageWithEnums) GetRegulars() []RegularEnum {
	if m != nil {
		return m.Regulars
	}
	return nil
}

func (m *MessageWithEnums) GetCustom() CustomEnum {
	if m != nil {
		return m.Custom
	}
	return CustomEnum_CUSTOM_UNKNOWN
}

func (m *MessageWithEnums) GetCustoms() []CustomEnum {
	if m != nil {
		return m.Customs
	}
	return nil
}

func (m *MessageWithEnums) GetWrappedCustom() *CustomEnumValue {
	if m != nil {
		return m.WrappedCustom
	}
	return nil
}

func (m *MessageWithEnums) GetWrappedCustoms() []*CustomEnumValue {
	if m != nil {
		return m.WrappedCustoms
	}
	return nil
}

type MessageWithOneofEnums struct {
	// Types that are valid to be assigned to Value:
	//
	//	*MessageWithOneofEnums_Regular
	//	*MessageWithOneofEnums_Custom
	//	*MessageWithOneofEnums_WrappedCustom
	Value                isMessageWithOneofEnums_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MessageWithOneofEnums) Reset()         { *m = MessageWithOneofEnums{} }
func (m *MessageWithOneofEnums) String() string { return proto.CompactTextString(m) }
func (*MessageWithOneofEnums) ProtoMessage()    {}
func (*MessageWithOneofEnums) Descriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{2}
}
func (m *MessageWithOneofEnums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithOneofEnums.Unmarshal(m, b)
}
func (m *MessageWithOneofEnums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithOneofEnums.Marshal(b, m, deterministic)
}
func (m *MessageWithOneofEnums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithOneofEnums.Merge(m, src)
}
func (m *MessageWithOneofEnums) XXX_Size() int {
	return xxx_messageInfo_MessageWithOneofEnums.Size(m)
}
func (m *MessageWithOneofEnums) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithOneofEnums.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithOneofEnums proto.InternalMessageInfo

type isMessageWithOneofEnums_Value interface {
	isMessageWithOneofEnums_Value()
}

type MessageWithOneofEnums_Regular struct {
	Regular RegularEnum `protobuf:"varint,1,opt,name=regular,proto3,enum=thethings.flags.test.RegularEnum,oneof" json:"regular,omitempty"`
}
type MessageWithOneofEnums_Custom struct {
	Custom CustomEnum `protobuf:"varint,2,opt,name=custom,proto3,enum=thethings.flags.test.CustomEnum,oneof" json:"custom,omitempty"`
}
type MessageWithOneofEnums_WrappedCustom struct {
	WrappedCustom *CustomEnumValue `protobuf:"bytes,3,opt,name=wrapped_custom,json=wrappedCustom,proto3,oneof" json:"wrapped_custom,omitempty"`
}

func (*MessageWithOneofEnums_Regular) isMessageWithOneofEnums_Value()       {}
func (*MessageWithOneofEnums_Custom) isMessageWithOneofEnums_Value()        {}
func (*MessageWithOneofEnums_WrappedCustom) isMessageWithOneofEnums_Value() {}

func (m *MessageWithOneofEnums) GetValue() isMessageWithOneofEnums_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MessageWithOneofEnums) GetRegular() RegularEnum {
	if x, ok := m.GetValue().(*MessageWithOneofEnums_Regular); ok {
		return x.Regular
	}
	return REGULAR_UNKNOWN
}

func (m *MessageWithOneofEnums) GetCustom() CustomEnum {
	if x, ok := m.GetValue().(*MessageWithOneofEnums_Custom); ok {
		return x.Custom
	}
	return CustomEnum_CUSTOM_UNKNOWN
}

func (m *MessageWithOneofEnums) GetWrappedCustom() *CustomEnumValue {
	if x, ok := m.GetValue().(*MessageWithOneofEnums_WrappedCustom); ok {
		return x.WrappedCustom
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MessageWithOneofEnums) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MessageWithOneofEnums_Regular)(nil),
		(*MessageWithOneofEnums_Custom)(nil),
		(*MessageWithOneofEnums_WrappedCustom)(nil),
	}
}

type MessageWithEmbeddedMessageDefinitionWithEnums struct {
	TestMessageField     *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum `protobuf:"bytes,1,opt,name=test_message_field,json=testMessageField,proto3" json:"test_message_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                         `json:"-"`
	XXX_unrecognized     []byte                                                                           `json:"-"`
	XXX_sizecache        int32                                                                            `json:"-"`
}

func (m *MessageWithEmbeddedMessageDefinitionWithEnums) Reset() {
	*m = MessageWithEmbeddedMessageDefinitionWithEnums{}
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums) String() string {
	return proto.CompactTextString(m)
}
func (*MessageWithEmbeddedMessageDefinitionWithEnums) ProtoMessage() {}
func (*MessageWithEmbeddedMessageDefinitionWithEnums) Descriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{3}
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums.Unmarshal(m, b)
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums.Marshal(b, m, deterministic)
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums.Merge(m, src)
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums) XXX_Size() int {
	return xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums.Size(m)
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums proto.InternalMessageInfo

func (m *MessageWithEmbeddedMessageDefinitionWithEnums) GetTestMessageField() *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum {
	if m != nil {
		return m.TestMessageField
	}
	return nil
}

type MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum struct {
	StringValue          string                                                                                       `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	TestEnumValue        MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum `protobuf:"varint,2,opt,name=test_enum_value,json=testEnumValue,proto3,enum=thethings.flags.test.MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum" json:"test_enum_value,omitempty"`
	TestAnotherEnumValue MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum `protobuf:"varint,3,opt,name=test_another_enum_value,json=testAnotherEnumValue,proto3,enum=thethings.flags.test.MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum" json:"test_another_enum_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                     `json:"-"`
	XXX_unrecognized     []byte                                                                                       `json:"-"`
	XXX_sizecache        int32                                                                                        `json:"-"`
}

func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) Reset() {
	*m = MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum{}
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) String() string {
	return proto.CompactTextString(m)
}
func (*MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) ProtoMessage() {
}
func (*MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_888b6bd9597961ff, []int{3, 0}
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum.Unmarshal(m, b)
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum.Marshal(b, m, deterministic)
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum.Merge(m, src)
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) XXX_Size() int {
	return xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum.Size(m)
}
func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum proto.InternalMessageInfo

func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) GetTestEnumValue() MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum {
	if m != nil {
		return m.TestEnumValue
	}
	return AT_MOST_ONCE
}

func (m *MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum) GetTestAnotherEnumValue() MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum {
	if m != nil {
		return m.TestAnotherEnumValue
	}
	return AT_MOST_ONCE
}

func init() {
	proto.RegisterEnum("thethings.flags.test.RegularEnum", RegularEnum_name, RegularEnum_value)
	proto.RegisterEnum("thethings.flags.test.CustomEnum", CustomEnum_name, CustomEnum_value)
	proto.RegisterEnum("thethings.flags.test.MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum", MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum_name, MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum_EmbeddedEnum_value)
	proto.RegisterType((*CustomEnumValue)(nil), "thethings.flags.test.CustomEnumValue")
	proto.RegisterType((*MessageWithEnums)(nil), "thethings.flags.test.MessageWithEnums")
	proto.RegisterType((*MessageWithOneofEnums)(nil), "thethings.flags.test.MessageWithOneofEnums")
	proto.RegisterType((*MessageWithEmbeddedMessageDefinitionWithEnums)(nil), "thethings.flags.test.MessageWithEmbeddedMessageDefinitionWithEnums")
	proto.RegisterType((*MessageWithEmbeddedMessageDefinitionWithEnums_EmbeddedMessageDefinitionWithEnum)(nil), "thethings.flags.test.MessageWithEmbeddedMessageDefinitionWithEnums.EmbeddedMessageDefinitionWithEnum")
}

func init() { proto.RegisterFile("enums.proto", fileDescriptor_888b6bd9597961ff) }

var fileDescriptor_888b6bd9597961ff = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xf6, 0x26, 0x6d, 0x9a, 0x4e, 0x9a, 0xc4, 0xdd, 0x5f, 0x7f, 0x22, 0x8d, 0x44, 0x94, 0x46,
	0x42, 0xaa, 0x2a, 0xc5, 0xa1, 0x45, 0x02, 0x14, 0xd4, 0x83, 0x1b, 0x02, 0x05, 0xd2, 0x44, 0x72,
	0x9d, 0x96, 0x72, 0xb1, 0xdc, 0x66, 0xe3, 0x58, 0x4a, 0xd6, 0x95, 0xd7, 0x86, 0x57, 0xe0, 0xce,
	0x8d, 0x0b, 0x07, 0xb8, 0x50, 0x71, 0xe0, 0x01, 0x78, 0x9a, 0x9e, 0x10, 0x27, 0x54, 0x89, 0x3b,
	0xda, 0xb5, 0x13, 0xbb, 0x50, 0x29, 0x29, 0x07, 0xb8, 0x4d, 0x66, 0xbf, 0xf9, 0xe6, 0xcf, 0x37,
	0xe3, 0x40, 0x86, 0x50, 0x7f, 0xc4, 0x94, 0x53, 0xd7, 0xf1, 0x1c, 0xbc, 0xe2, 0x0d, 0x88, 0x37,
	0xb0, 0xa9, 0xc5, 0x94, 0xfe, 0xd0, 0xb4, 0x98, 0xe2, 0x11, 0xe6, 0x15, 0x97, 0x4d, 0x4a, 0x1d,
	0xcf, 0xf4, 0x6c, 0x87, 0x86, 0xc0, 0xe2, 0x4d, 0x9b, 0x7a, 0xc4, 0xa5, 0xe6, 0xb0, 0x66, 0x39,
	0x96, 0x23, 0x7c, 0xc2, 0x0a, 0x9e, 0x2b, 0x5d, 0xc8, 0x37, 0x7c, 0xe6, 0x39, 0xa3, 0x26, 0xf5,
	0x47, 0x07, 0xe6, 0xd0, 0x27, 0xf8, 0x2e, 0xcc, 0xbf, 0xe4, 0x46, 0x01, 0x95, 0xd1, 0x7a, 0x6e,
	0xab, 0xac, 0x5c, 0x95, 0x4a, 0x89, 0xa2, 0xb4, 0x00, 0x5e, 0x87, 0x8b, 0xb3, 0xd5, 0x54, 0x1a,
	0xc9, 0xa8, 0x80, 0x2a, 0x9f, 0x92, 0x20, 0xef, 0x11, 0xc6, 0x4c, 0x8b, 0x1c, 0xda, 0xde, 0x80,
	0xc3, 0x18, 0x7e, 0x00, 0x0b, 0x2e, 0xb1, 0xfc, 0xa1, 0xe9, 0x86, 0xd4, 0x6b, 0x57, 0x53, 0x6b,
	0x01, 0x48, 0x70, 0x8f, 0x23, 0xf0, 0x36, 0xa4, 0x43, 0x93, 0x15, 0x12, 0xe5, 0xe4, 0x6c, 0xd1,
	0x93, 0x10, 0x7c, 0x1f, 0x52, 0x27, 0xa2, 0xe2, 0x42, 0x72, 0xc6, 0xae, 0x42, 0x3c, 0xae, 0xc3,
	0x42, 0x60, 0xb1, 0xc2, 0x9c, 0xc8, 0x3b, 0x3d, 0x74, 0x1c, 0x80, 0x5b, 0x90, 0x7b, 0xe5, 0x9a,
	0xa7, 0xa7, 0xa4, 0x67, 0x84, 0xd9, 0xe7, 0xcb, 0x68, 0x3d, 0xb3, 0x75, 0x6b, 0x1a, 0x85, 0x50,
	0x42, 0xcb, 0x86, 0xc1, 0x81, 0x1f, 0xb7, 0x21, 0x7f, 0x99, 0x8d, 0x15, 0x52, 0xe5, 0xe4, 0xec,
	0x74, 0xb9, 0x4b, 0x74, 0xac, 0x9e, 0xbe, 0x38, 0x5b, 0x9d, 0xe3, 0x82, 0x55, 0x7e, 0x20, 0xf8,
	0x3f, 0x26, 0x57, 0x87, 0x12, 0xa7, 0x1f, 0x68, 0xb6, 0x7d, 0x7d, 0xcd, 0x76, 0xa5, 0x48, 0xb5,
	0xfa, 0x64, 0xec, 0x89, 0xd9, 0xc6, 0xbe, 0x2b, 0x4d, 0x06, 0xdf, 0xfe, 0x6d, 0x78, 0xc9, 0x6b,
	0x0c, 0x6f, 0x57, 0xfa, 0x65, 0x7c, 0x51, 0xbb, 0x3b, 0x0b, 0xe1, 0x86, 0x57, 0xbe, 0xcc, 0x43,
	0x35, 0xbe, 0xa6, 0xa3, 0x63, 0xd2, 0xeb, 0x91, 0x5e, 0xe8, 0x7a, 0x48, 0xfa, 0x36, 0xb5, 0xf9,
	0x31, 0x45, 0x3b, 0xfc, 0x06, 0x01, 0xe6, 0xe9, 0x8c, 0x51, 0x80, 0x31, 0xfa, 0x36, 0x19, 0xf6,
	0xc4, 0x6c, 0x32, 0x5b, 0xe4, 0xea, 0xca, 0xae, 0x95, 0x41, 0x99, 0x0a, 0xd1, 0x64, 0xce, 0x1a,
	0x3e, 0x3f, 0xe2, 0xe9, 0x8b, 0xdf, 0x92, 0xb0, 0x36, 0x35, 0x0e, 0xaf, 0xc1, 0x12, 0xf3, 0x5c,
	0x9b, 0x5a, 0x46, 0x74, 0xdf, 0x8b, 0x5a, 0x26, 0xf0, 0x05, 0xb7, 0xff, 0x16, 0x41, 0x5e, 0xb4,
	0xc7, 0xbf, 0x35, 0x21, 0x2c, 0x50, 0xce, 0xfd, 0x2b, 0xbd, 0x4d, 0x10, 0xa2, 0xd1, 0x2c, 0x4f,
	0x11, 0x7d, 0x98, 0x3e, 0x22, 0xb8, 0x21, 0x8a, 0x33, 0xa9, 0xe3, 0x0d, 0x88, 0x1b, 0x2f, 0x32,
	0xf9, 0xcf, 0x8a, 0x5c, 0xe1, 0x29, 0xd4, 0xa0, 0xa2, 0x49, 0xad, 0x95, 0x0e, 0x2c, 0xc5, 0x51,
	0x58, 0x86, 0x25, 0x55, 0x37, 0xf6, 0x3a, 0xfb, 0xba, 0xd1, 0x69, 0x37, 0x9a, 0xb2, 0x84, 0x97,
	0x21, 0xab, 0xea, 0x46, 0xab, 0xa9, 0x8e, 0x5d, 0x88, 0x83, 0x9a, 0xcf, 0xd5, 0x86, 0xde, 0x3a,
	0x0a, 0x3c, 0x89, 0x62, 0xfa, 0xf5, 0xfb, 0x92, 0xf4, 0xf9, 0x43, 0x49, 0x8a, 0xb6, 0x37, 0xb2,
	0x36, 0x9e, 0x42, 0x26, 0x76, 0x77, 0xf8, 0x3f, 0xc8, 0x6b, 0xcd, 0xc7, 0xdd, 0x96, 0xaa, 0x19,
	0xdd, 0xf6, 0xb3, 0x76, 0xe7, 0xb0, 0x2d, 0x4b, 0x38, 0x0b, 0x8b, 0x63, 0xa7, 0x2a, 0xa3, 0xf8,
	0xcf, 0x9d, 0x38, 0xff, 0xc6, 0x3b, 0x04, 0x10, 0x9d, 0x10, 0xc6, 0x90, 0x6b, 0x74, 0xf7, 0xf5,
	0xce, 0x5e, 0x8c, 0x2a, 0x0f, 0x99, 0xd0, 0x77, 0xb0, 0x69, 0xdc, 0x96, 0x11, 0x6f, 0x21, 0xe6,
	0x30, 0x36, 0xe5, 0x44, 0xf1, 0xe8, 0xe2, 0x6c, 0xb5, 0x0b, 0xfb, 0x96, 0xed, 0x0d, 0xfc, 0x63,
	0xe5, 0xc4, 0x19, 0xd5, 0xf4, 0x01, 0xd1, 0x85, 0x22, 0x4f, 0x68, 0xcf, 0xe7, 0xbb, 0x46, 0x58,
	0x4d, 0xfc, 0xf7, 0x9c, 0x54, 0x2d, 0x42, 0xab, 0x96, 0x53, 0x15, 0x4a, 0xd5, 0xf8, 0x18, 0x6b,
	0x5c, 0xd1, 0xd8, 0x25, 0x87, 0x87, 0x2f, 0x24, 0xde, 0x51, 0xbf, 0x9e, 0x97, 0xa4, 0xef, 0xe7,
	0x25, 0xf4, 0xe2, 0xde, 0x1f, 0xd2, 0x1f, 0xa7, 0xc4, 0xcb, 0x9d, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x37, 0x02, 0xf1, 0xdf, 0x4b, 0x07, 0x00, 0x00,
}

func (x CustomEnum) String() string {
	s, ok := CustomEnum_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
