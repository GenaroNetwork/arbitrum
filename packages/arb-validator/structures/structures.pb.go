// Code generated by protoc-gen-go. DO NOT EDIT.
// source: structures.proto

package structures

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/offchainlabs/arbitrum/packages/arb-util/common"
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

type ExecutionAssertionBuf struct {
	AfterHash            *common.HashBuf   `protobuf:"bytes,1,opt,name=afterHash,proto3" json:"afterHash,omitempty"`
	DidInboxInsn         bool              `protobuf:"varint,2,opt,name=didInboxInsn,proto3" json:"didInboxInsn,omitempty"`
	NumGas               uint64            `protobuf:"varint,3,opt,name=numGas,proto3" json:"numGas,omitempty"`
	Messages             []*common.HashBuf `protobuf:"bytes,4,rep,name=messages,proto3" json:"messages,omitempty"`
	Logs                 []*common.HashBuf `protobuf:"bytes,5,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExecutionAssertionBuf) Reset()         { *m = ExecutionAssertionBuf{} }
func (m *ExecutionAssertionBuf) String() string { return proto.CompactTextString(m) }
func (*ExecutionAssertionBuf) ProtoMessage()    {}
func (*ExecutionAssertionBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{0}
}

func (m *ExecutionAssertionBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionAssertionBuf.Unmarshal(m, b)
}
func (m *ExecutionAssertionBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionAssertionBuf.Marshal(b, m, deterministic)
}
func (m *ExecutionAssertionBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionAssertionBuf.Merge(m, src)
}
func (m *ExecutionAssertionBuf) XXX_Size() int {
	return xxx_messageInfo_ExecutionAssertionBuf.Size(m)
}
func (m *ExecutionAssertionBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionAssertionBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionAssertionBuf proto.InternalMessageInfo

func (m *ExecutionAssertionBuf) GetAfterHash() *common.HashBuf {
	if m != nil {
		return m.AfterHash
	}
	return nil
}

func (m *ExecutionAssertionBuf) GetDidInboxInsn() bool {
	if m != nil {
		return m.DidInboxInsn
	}
	return false
}

func (m *ExecutionAssertionBuf) GetNumGas() uint64 {
	if m != nil {
		return m.NumGas
	}
	return 0
}

func (m *ExecutionAssertionBuf) GetMessages() []*common.HashBuf {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *ExecutionAssertionBuf) GetLogs() []*common.HashBuf {
	if m != nil {
		return m.Logs
	}
	return nil
}

type InboxItemBuf struct {
	ValType              uint32          `protobuf:"varint,1,opt,name=valType,proto3" json:"valType,omitempty"`
	ValHash              *common.HashBuf `protobuf:"bytes,2,opt,name=valHash,proto3" json:"valHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *InboxItemBuf) Reset()         { *m = InboxItemBuf{} }
func (m *InboxItemBuf) String() string { return proto.CompactTextString(m) }
func (*InboxItemBuf) ProtoMessage()    {}
func (*InboxItemBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{1}
}

func (m *InboxItemBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InboxItemBuf.Unmarshal(m, b)
}
func (m *InboxItemBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InboxItemBuf.Marshal(b, m, deterministic)
}
func (m *InboxItemBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxItemBuf.Merge(m, src)
}
func (m *InboxItemBuf) XXX_Size() int {
	return xxx_messageInfo_InboxItemBuf.Size(m)
}
func (m *InboxItemBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxItemBuf.DiscardUnknown(m)
}

var xxx_messageInfo_InboxItemBuf proto.InternalMessageInfo

func (m *InboxItemBuf) GetValType() uint32 {
	if m != nil {
		return m.ValType
	}
	return 0
}

func (m *InboxItemBuf) GetValHash() *common.HashBuf {
	if m != nil {
		return m.ValHash
	}
	return nil
}

type InboxBuf struct {
	TopCount             *common.BigIntegerBuf `protobuf:"bytes,1,opt,name=topCount,proto3" json:"topCount,omitempty"`
	Items                []*InboxItemBuf       `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	HashOfRest           *common.HashBuf       `protobuf:"bytes,3,opt,name=hashOfRest,proto3" json:"hashOfRest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *InboxBuf) Reset()         { *m = InboxBuf{} }
func (m *InboxBuf) String() string { return proto.CompactTextString(m) }
func (*InboxBuf) ProtoMessage()    {}
func (*InboxBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ea84bc81126bff, []int{2}
}

func (m *InboxBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InboxBuf.Unmarshal(m, b)
}
func (m *InboxBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InboxBuf.Marshal(b, m, deterministic)
}
func (m *InboxBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InboxBuf.Merge(m, src)
}
func (m *InboxBuf) XXX_Size() int {
	return xxx_messageInfo_InboxBuf.Size(m)
}
func (m *InboxBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_InboxBuf.DiscardUnknown(m)
}

var xxx_messageInfo_InboxBuf proto.InternalMessageInfo

func (m *InboxBuf) GetTopCount() *common.BigIntegerBuf {
	if m != nil {
		return m.TopCount
	}
	return nil
}

func (m *InboxBuf) GetItems() []*InboxItemBuf {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *InboxBuf) GetHashOfRest() *common.HashBuf {
	if m != nil {
		return m.HashOfRest
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecutionAssertionBuf)(nil), "structures.ExecutionAssertionBuf")
	proto.RegisterType((*InboxItemBuf)(nil), "structures.InboxItemBuf")
	proto.RegisterType((*InboxBuf)(nil), "structures.InboxBuf")
}

func init() { proto.RegisterFile("structures.proto", fileDescriptor_66ea84bc81126bff) }

var fileDescriptor_66ea84bc81126bff = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x4e, 0x83, 0x30,
	0x14, 0xc6, 0xc3, 0xfe, 0x89, 0x75, 0x46, 0x53, 0x33, 0xd3, 0x78, 0x45, 0xf0, 0x66, 0xc6, 0x0c,
	0xa2, 0x3e, 0x81, 0xa8, 0xd1, 0x5d, 0x99, 0xa0, 0x57, 0xde, 0x15, 0x56, 0xa0, 0x11, 0x5a, 0xd2,
	0xd3, 0x2e, 0xf3, 0x65, 0x7c, 0x2a, 0x1f, 0xc8, 0x50, 0x86, 0x9b, 0x09, 0x57, 0xd0, 0xef, 0xfc,
	0x7a, 0xce, 0xf7, 0x1d, 0x40, 0xa7, 0xa0, 0x95, 0x49, 0xb5, 0x51, 0x0c, 0x82, 0x5a, 0x49, 0x2d,
	0x31, 0xda, 0x29, 0x17, 0x67, 0xa9, 0xac, 0x2a, 0x29, 0xc2, 0xf6, 0xd1, 0x02, 0xfe, 0x8f, 0x83,
	0x66, 0x4f, 0x1b, 0x96, 0x1a, 0xcd, 0xa5, 0xb8, 0x07, 0x60, 0xaa, 0x79, 0x89, 0x4c, 0x86, 0x17,
	0xe8, 0x90, 0x66, 0x9a, 0xa9, 0x17, 0x0a, 0x05, 0x71, 0x3c, 0x67, 0x7e, 0x74, 0x7b, 0x12, 0x6c,
	0xef, 0x36, 0x5a, 0x64, 0xb2, 0x78, 0x47, 0x60, 0x1f, 0x4d, 0x57, 0x7c, 0xb5, 0x14, 0x89, 0xdc,
	0x2c, 0x05, 0x08, 0x32, 0xf0, 0x9c, 0xb9, 0x1b, 0xff, 0xd3, 0xf0, 0x39, 0x9a, 0x08, 0x53, 0x3d,
	0x53, 0x20, 0x43, 0xcf, 0x99, 0x8f, 0xe2, 0xed, 0x09, 0x5f, 0x23, 0xb7, 0x62, 0x00, 0x34, 0x67,
	0x40, 0x46, 0xde, 0xb0, 0x6f, 0xd2, 0x1f, 0x80, 0x2f, 0xd1, 0xa8, 0x94, 0x39, 0x90, 0x71, 0x3f,
	0x68, 0x8b, 0xfe, 0x1b, 0x9a, 0xb6, 0x63, 0x35, 0xab, 0x9a, 0x30, 0x04, 0x1d, 0xac, 0x69, 0xf9,
	0xfe, 0x55, 0x33, 0x1b, 0xe5, 0x38, 0xee, 0x8e, 0xf8, 0xca, 0x56, 0x6c, 0xc8, 0x41, 0x7f, 0xc8,
	0xae, 0xee, 0x7f, 0x3b, 0xc8, 0xb5, 0x5d, 0x9b, 0x8e, 0x37, 0xc8, 0xd5, 0xb2, 0x7e, 0x90, 0x46,
	0xe8, 0xed, 0x76, 0x66, 0xdd, 0xc5, 0x88, 0xe7, 0x4b, 0xa1, 0x59, 0xce, 0x94, 0x75, 0xde, 0x61,
	0x38, 0x40, 0x63, 0xae, 0x59, 0x05, 0x64, 0x60, 0xad, 0x93, 0x60, 0xef, 0x73, 0xed, 0xbb, 0x8d,
	0x5b, 0x0c, 0x87, 0x08, 0x15, 0x14, 0x8a, 0xd7, 0x2c, 0x66, 0xa0, 0xed, 0xca, 0x7a, 0xdc, 0xed,
	0x21, 0xd1, 0xe3, 0x47, 0x94, 0x73, 0x5d, 0x98, 0xa4, 0x81, 0x42, 0x99, 0x65, 0x69, 0x41, 0xb9,
	0x28, 0x69, 0x02, 0x21, 0x55, 0x09, 0xd7, 0xca, 0x54, 0x61, 0x4d, 0xd3, 0xcf, 0x66, 0x8f, 0x8d,
	0xb2, 0x58, 0xd3, 0x92, 0xaf, 0xa8, 0x96, 0x2a, 0xdc, 0x59, 0x49, 0x26, 0xf6, 0xcf, 0xb8, 0xfb,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x60, 0xcd, 0x35, 0x8b, 0x4e, 0x02, 0x00, 0x00,
}
