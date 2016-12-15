// Code generated by protoc-gen-go.
// source: op.proto
// DO NOT EDIT!

/*
Package statepb is a generated protocol buffer package.

It is generated from these files:
	op.proto

It has these top-level messages:
	Op
*/
package statepb

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

type Op_Type int32

const (
	Op_SET                    Op_Type = 0
	Op_INCREMENT_NEXT_NODE_ID Op_Type = 1
)

var Op_Type_name = map[int32]string{
	0: "SET",
	1: "INCREMENT_NEXT_NODE_ID",
}
var Op_Type_value = map[string]int32{
	"SET": 0,
	"INCREMENT_NEXT_NODE_ID": 1,
}

func (x Op_Type) String() string {
	return proto.EnumName(Op_Type_name, int32(x))
}
func (Op_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Op struct {
	Type       Op_Type `protobuf:"varint,1,opt,name=type,enum=statepb.Op_Type" json:"type,omitempty"`
	NextNodeId uint64  `protobuf:"varint,2,opt,name=next_node_id,json=nextNodeId" json:"next_node_id,omitempty"`
	Key        string  `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value      string  `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *Op) Reset()                    { *m = Op{} }
func (m *Op) String() string            { return proto.CompactTextString(m) }
func (*Op) ProtoMessage()               {}
func (*Op) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Op) GetType() Op_Type {
	if m != nil {
		return m.Type
	}
	return Op_SET
}

func (m *Op) GetNextNodeId() uint64 {
	if m != nil {
		return m.NextNodeId
	}
	return 0
}

func (m *Op) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Op) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Op)(nil), "statepb.Op")
	proto.RegisterEnum("statepb.Op_Type", Op_Type_name, Op_Type_value)
}

func init() { proto.RegisterFile("op.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x2f, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x2e, 0x49, 0x2c, 0x49, 0x2d, 0x48, 0x52, 0x5a, 0xc8,
	0xc8, 0xc5, 0xe4, 0x5f, 0x20, 0xa4, 0xc2, 0xc5, 0x52, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x67, 0x24, 0xa0, 0x07, 0x95, 0xd6, 0xf3, 0x2f, 0xd0, 0x0b, 0xa9, 0x2c, 0x48, 0x0d,
	0x02, 0xcb, 0x0a, 0x29, 0x70, 0xf1, 0xe4, 0xa5, 0x56, 0x94, 0xc4, 0xe7, 0xe5, 0xa7, 0xa4, 0xc6,
	0x67, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x04, 0x71, 0x81, 0xc4, 0xfc, 0xf2, 0x53, 0x52,
	0x3d, 0x53, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x98, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x16, 0xb0, 0x18, 0x84,
	0xa3, 0xa4, 0xcd, 0xc5, 0x02, 0x32, 0x57, 0x88, 0x9d, 0x8b, 0x39, 0xd8, 0x35, 0x44, 0x80, 0x41,
	0x48, 0x8a, 0x4b, 0xcc, 0xd3, 0xcf, 0x39, 0xc8, 0xd5, 0xd7, 0xd5, 0x2f, 0x24, 0xde, 0xcf, 0x35,
	0x22, 0x24, 0xde, 0xcf, 0xdf, 0xc5, 0x35, 0xde, 0xd3, 0x45, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x66,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x1f, 0x40, 0xcd, 0xbf, 0x00, 0x00, 0x00,
}