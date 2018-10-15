// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fib.proto

package l2

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FIBEntry_Action int32

const (
	FIBEntry_FORWARD FIBEntry_Action = 0
	FIBEntry_DROP    FIBEntry_Action = 1
)

var FIBEntry_Action_name = map[int32]string{
	0: "FORWARD",
	1: "DROP",
}
var FIBEntry_Action_value = map[string]int32{
	"FORWARD": 0,
	"DROP":    1,
}

func (x FIBEntry_Action) String() string {
	return proto.EnumName(FIBEntry_Action_name, int32(x))
}
func (FIBEntry_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fib_b631b45f310d5ae3, []int{0, 0}
}

type FIBEntry struct {
	PhysAddress             string          `protobuf:"bytes,1,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	BridgeDomain            string          `protobuf:"bytes,2,opt,name=bridge_domain,json=bridgeDomain,proto3" json:"bridge_domain,omitempty"`
	Action                  FIBEntry_Action `protobuf:"varint,3,opt,name=action,proto3,enum=l2.FIBEntry_Action" json:"action,omitempty"`
	OutgoingInterface       string          `protobuf:"bytes,4,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	StaticConfig            bool            `protobuf:"varint,5,opt,name=static_config,json=staticConfig,proto3" json:"static_config,omitempty"`
	BridgedVirtualInterface bool            `protobuf:"varint,6,opt,name=bridged_virtual_interface,json=bridgedVirtualInterface,proto3" json:"bridged_virtual_interface,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}        `json:"-"`
	XXX_unrecognized        []byte          `json:"-"`
	XXX_sizecache           int32           `json:"-"`
}

func (m *FIBEntry) Reset()         { *m = FIBEntry{} }
func (m *FIBEntry) String() string { return proto.CompactTextString(m) }
func (*FIBEntry) ProtoMessage()    {}
func (*FIBEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_fib_b631b45f310d5ae3, []int{0}
}
func (m *FIBEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FIBEntry.Unmarshal(m, b)
}
func (m *FIBEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FIBEntry.Marshal(b, m, deterministic)
}
func (dst *FIBEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FIBEntry.Merge(dst, src)
}
func (m *FIBEntry) XXX_Size() int {
	return xxx_messageInfo_FIBEntry.Size(m)
}
func (m *FIBEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FIBEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FIBEntry proto.InternalMessageInfo

func (m *FIBEntry) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *FIBEntry) GetBridgeDomain() string {
	if m != nil {
		return m.BridgeDomain
	}
	return ""
}

func (m *FIBEntry) GetAction() FIBEntry_Action {
	if m != nil {
		return m.Action
	}
	return FIBEntry_FORWARD
}

func (m *FIBEntry) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *FIBEntry) GetStaticConfig() bool {
	if m != nil {
		return m.StaticConfig
	}
	return false
}

func (m *FIBEntry) GetBridgedVirtualInterface() bool {
	if m != nil {
		return m.BridgedVirtualInterface
	}
	return false
}

func init() {
	proto.RegisterType((*FIBEntry)(nil), "l2.FIBEntry")
	proto.RegisterEnum("l2.FIBEntry_Action", FIBEntry_Action_name, FIBEntry_Action_value)
}

func init() { proto.RegisterFile("fib.proto", fileDescriptor_fib_b631b45f310d5ae3) }

var fileDescriptor_fib_b631b45f310d5ae3 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x71, 0x28, 0x21, 0x7d, 0x0d, 0xa8, 0x98, 0x01, 0x33, 0x11, 0xca, 0x12, 0x09, 0x91,
	0x21, 0x6c, 0x6c, 0x81, 0x50, 0xa9, 0x53, 0x91, 0x07, 0x18, 0x2d, 0x27, 0x4e, 0x82, 0xa5, 0x60,
	0x57, 0x8e, 0x8b, 0xd4, 0xbf, 0xe2, 0x13, 0x51, 0xed, 0x06, 0xb1, 0x9e, 0x73, 0xfd, 0xee, 0x95,
	0x61, 0xda, 0xca, 0x2a, 0xdb, 0x18, 0x6d, 0x35, 0x0e, 0xfa, 0x7c, 0xf1, 0x13, 0x40, 0xb4, 0x5c,
	0x3d, 0xbf, 0x2a, 0x6b, 0x76, 0xf8, 0x16, 0xe2, 0xcd, 0xe7, 0x6e, 0x60, 0x5c, 0x08, 0xd3, 0x0c,
	0x03, 0x41, 0x09, 0x4a, 0xa7, 0x74, 0xb6, 0x67, 0x85, 0x47, 0xf8, 0x0e, 0xce, 0x2a, 0x23, 0x45,
	0xd7, 0x30, 0xa1, 0xbf, 0xb8, 0x54, 0x24, 0x70, 0x99, 0xd8, 0xc3, 0xd2, 0x31, 0x7c, 0x0f, 0x21,
	0xaf, 0xad, 0xd4, 0x8a, 0x1c, 0x27, 0x28, 0x3d, 0xcf, 0x2f, 0xb3, 0x3e, 0xcf, 0xc6, 0x96, 0xac,
	0x70, 0x8a, 0x1e, 0x22, 0xf8, 0x01, 0xb0, 0xde, 0xda, 0x4e, 0x4b, 0xd5, 0x31, 0xa9, 0x6c, 0x63,
	0x5a, 0x5e, 0x37, 0x64, 0xe2, 0xce, 0x5e, 0x8c, 0x66, 0x35, 0x8a, 0xfd, 0x80, 0xc1, 0x72, 0x2b,
	0x6b, 0x56, 0x6b, 0xd5, 0xca, 0x8e, 0x9c, 0x24, 0x28, 0x8d, 0x68, 0xec, 0xe1, 0x8b, 0x63, 0xf8,
	0x09, 0xae, 0xfd, 0x20, 0xc1, 0xbe, 0xa5, 0xb1, 0x5b, 0xde, 0xff, 0x3b, 0x1d, 0xba, 0x07, 0x57,
	0x87, 0xc0, 0xbb, 0xf7, 0x7f, 0x05, 0x8b, 0x1b, 0x08, 0xfd, 0x42, 0x3c, 0x83, 0xd3, 0xe5, 0x9a,
	0x7e, 0x14, 0xb4, 0x9c, 0x1f, 0xe1, 0x08, 0x26, 0x25, 0x5d, 0xbf, 0xcd, 0x51, 0x15, 0xba, 0xdf,
	0x7b, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xae, 0xf9, 0xa8, 0xfa, 0x4a, 0x01, 0x00, 0x00,
}
