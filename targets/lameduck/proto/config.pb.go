// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/targets/lameduck/proto/config.proto

package proto

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

type Options struct {
	// How often to check for lame-ducked targets
	ReEvalSec *int32 `protobuf:"varint,1,opt,name=re_eval_sec,json=reEvalSec,def=10" json:"re_eval_sec,omitempty"`
	// Runtime config project. If running on GCE, this defaults to the project
	// containing the VM.
	RuntimeconfigProject *string `protobuf:"bytes,2,opt,name=runtimeconfig_project,json=runtimeconfigProject" json:"runtimeconfig_project,omitempty"`
	// Lame duck targets runtime config name. An operator will create a variable
	// here to mark a target as lame-ducked.
	RuntimeconfigName *string `protobuf:"bytes,3,opt,name=runtimeconfig_name,json=runtimeconfigName,def=lame-duck-targets" json:"runtimeconfig_name,omitempty"`
	// Lame duck targets pubsub topic name. An operator will create a message
	// here to mark a target as lame-ducked.
	PubsubTopic *string `protobuf:"bytes,7,opt,name=pubsub_topic,json=pubsubTopic" json:"pubsub_topic,omitempty"`
	// Lame duck expiration time. We ignore variables (targets) that have been
	// updated more than these many seconds ago. This is a safety mechanism for
	// failing to cleanup. Also, the idea is that if a target has actually
	// disappeared, automatic targets expansion will take care of that some time
	// during this expiration period.
	ExpirationSec *int32 `protobuf:"varint,4,opt,name=expiration_sec,json=expirationSec,def=300" json:"expiration_sec,omitempty"`
	// Use an RDS client to get lame-duck-targets.
	// This option is always true now and will be removed after v0.10.5.
	UseRds *bool `protobuf:"varint,5,opt,name=use_rds,json=useRds" json:"use_rds,omitempty"`
	// RDS server address. If not specified, RDS server address from the global
	// options is used, which in turn defaults to localhost:9314.
	RdsServerAddress     *string  `protobuf:"bytes,6,opt,name=rds_server_address,json=rdsServerAddress" json:"rds_server_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_24d4cbeb973fddd0, []int{0}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

const Default_Options_ReEvalSec int32 = 10
const Default_Options_RuntimeconfigName string = "lame-duck-targets"
const Default_Options_ExpirationSec int32 = 300

func (m *Options) GetReEvalSec() int32 {
	if m != nil && m.ReEvalSec != nil {
		return *m.ReEvalSec
	}
	return Default_Options_ReEvalSec
}

func (m *Options) GetRuntimeconfigProject() string {
	if m != nil && m.RuntimeconfigProject != nil {
		return *m.RuntimeconfigProject
	}
	return ""
}

func (m *Options) GetRuntimeconfigName() string {
	if m != nil && m.RuntimeconfigName != nil {
		return *m.RuntimeconfigName
	}
	return Default_Options_RuntimeconfigName
}

func (m *Options) GetPubsubTopic() string {
	if m != nil && m.PubsubTopic != nil {
		return *m.PubsubTopic
	}
	return ""
}

func (m *Options) GetExpirationSec() int32 {
	if m != nil && m.ExpirationSec != nil {
		return *m.ExpirationSec
	}
	return Default_Options_ExpirationSec
}

func (m *Options) GetUseRds() bool {
	if m != nil && m.UseRds != nil {
		return *m.UseRds
	}
	return false
}

func (m *Options) GetRdsServerAddress() string {
	if m != nil && m.RdsServerAddress != nil {
		return *m.RdsServerAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Options)(nil), "cloudprober.targets.lameduck.Options")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/targets/lameduck/proto/config.proto", fileDescriptor_24d4cbeb973fddd0)
}

var fileDescriptor_24d4cbeb973fddd0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x53, 0x10, 0x90, 0x45, 0x8d, 0x6c, 0x34, 0xf6, 0xe0, 0x01, 0x39, 0x11, 0x23, 0x2d,
	0x86, 0x1b, 0x27, 0x39, 0x78, 0x55, 0x53, 0xbc, 0x6f, 0xb6, 0xbb, 0x63, 0xad, 0xb6, 0xdd, 0xcd,
	0xcc, 0x2e, 0xf1, 0xb7, 0xf9, 0xeb, 0x4c, 0x5b, 0x8c, 0xf6, 0x38, 0xef, 0xcd, 0x97, 0xf7, 0x1e,
	0xdb, 0x66, 0xb9, 0x7b, 0xf7, 0x69, 0xa4, 0x4c, 0x19, 0x67, 0xc6, 0x64, 0x05, 0xc4, 0xaa, 0x30,
	0x5e, 0x5b, 0x34, 0x29, 0x60, 0xec, 0x24, 0x66, 0xe0, 0x28, 0x2e, 0x64, 0x09, 0xda, 0xab, 0xcf,
	0xd8, 0xa2, 0x71, 0x26, 0x56, 0xa6, 0x7a, 0xcb, 0xb3, 0xa8, 0x39, 0xf8, 0xf5, 0x3f, 0x20, 0x3a,
	0x00, 0xd1, 0x2f, 0x30, 0xff, 0xee, 0xb1, 0xd1, 0xb3, 0x75, 0xb9, 0xa9, 0x88, 0xcf, 0xd9, 0x04,
	0x41, 0xc0, 0x5e, 0x16, 0x82, 0x40, 0x85, 0xc1, 0x2c, 0x58, 0x0c, 0x36, 0xbd, 0xfb, 0x55, 0x32,
	0x46, 0x78, 0xdc, 0xcb, 0x62, 0x07, 0x8a, 0xaf, 0xd9, 0x25, 0xfa, 0xca, 0xe5, 0x25, 0xb4, 0x21,
	0xc2, 0xa2, 0xf9, 0x00, 0xe5, 0xc2, 0xde, 0x2c, 0x58, 0x8c, 0x93, 0x8b, 0x8e, 0xf9, 0xd2, 0x7a,
	0xfc, 0x81, 0xf1, 0x2e, 0x54, 0xc9, 0x12, 0xc2, 0x7e, 0x4d, 0x6c, 0xa6, 0x75, 0x95, 0x65, 0xdd,
	0x65, 0x79, 0x28, 0x97, 0x4c, 0x3b, 0xcf, 0x4f, 0xb2, 0x04, 0x7e, 0xc3, 0x4e, 0xac, 0x4f, 0xc9,
	0xa7, 0xc2, 0x19, 0x9b, 0xab, 0x70, 0xd4, 0xa4, 0x4d, 0x5a, 0xed, 0xb5, 0x96, 0xf8, 0x2d, 0x3b,
	0x83, 0x2f, 0x9b, 0xa3, 0xac, 0xc7, 0x34, 0x03, 0x8e, 0x9a, 0x01, 0xfd, 0xf5, 0x6a, 0x95, 0x9c,
	0xfe, 0x59, 0xf5, 0x8a, 0x2b, 0x36, 0xf2, 0x04, 0x02, 0x35, 0x85, 0x83, 0x59, 0xb0, 0x38, 0x4e,
	0x86, 0x9e, 0x20, 0xd1, 0xc4, 0xef, 0x18, 0x47, 0x4d, 0x82, 0x00, 0xf7, 0x80, 0x42, 0x6a, 0x8d,
	0x40, 0x14, 0x0e, 0x9b, 0xb4, 0x73, 0xd4, 0xb4, 0x6b, 0x8c, 0x6d, 0xab, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x9f, 0xcd, 0xa8, 0x50, 0x9e, 0x01, 0x00, 0x00,
}
