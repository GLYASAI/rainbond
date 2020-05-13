// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/listener/v3/udp_listener_config.proto

package envoy_config_listener_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type UdpListenerConfig struct {
	UdpListenerName string `protobuf:"bytes,1,opt,name=udp_listener_name,json=udpListenerName,proto3" json:"udp_listener_name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*UdpListenerConfig_HiddenEnvoyDeprecatedConfig
	//	*UdpListenerConfig_TypedConfig
	ConfigType           isUdpListenerConfig_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *UdpListenerConfig) Reset()         { *m = UdpListenerConfig{} }
func (m *UdpListenerConfig) String() string { return proto.CompactTextString(m) }
func (*UdpListenerConfig) ProtoMessage()    {}
func (*UdpListenerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ac914914a155255, []int{0}
}

func (m *UdpListenerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpListenerConfig.Unmarshal(m, b)
}
func (m *UdpListenerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpListenerConfig.Marshal(b, m, deterministic)
}
func (m *UdpListenerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpListenerConfig.Merge(m, src)
}
func (m *UdpListenerConfig) XXX_Size() int {
	return xxx_messageInfo_UdpListenerConfig.Size(m)
}
func (m *UdpListenerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpListenerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_UdpListenerConfig proto.InternalMessageInfo

func (m *UdpListenerConfig) GetUdpListenerName() string {
	if m != nil {
		return m.UdpListenerName
	}
	return ""
}

type isUdpListenerConfig_ConfigType interface {
	isUdpListenerConfig_ConfigType()
}

type UdpListenerConfig_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type UdpListenerConfig_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*UdpListenerConfig_HiddenEnvoyDeprecatedConfig) isUdpListenerConfig_ConfigType() {}

func (*UdpListenerConfig_TypedConfig) isUdpListenerConfig_ConfigType() {}

func (m *UdpListenerConfig) GetConfigType() isUdpListenerConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *UdpListenerConfig) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*UdpListenerConfig_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *UdpListenerConfig) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*UdpListenerConfig_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UdpListenerConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UdpListenerConfig_HiddenEnvoyDeprecatedConfig)(nil),
		(*UdpListenerConfig_TypedConfig)(nil),
	}
}

type ActiveRawUdpListenerConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActiveRawUdpListenerConfig) Reset()         { *m = ActiveRawUdpListenerConfig{} }
func (m *ActiveRawUdpListenerConfig) String() string { return proto.CompactTextString(m) }
func (*ActiveRawUdpListenerConfig) ProtoMessage()    {}
func (*ActiveRawUdpListenerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ac914914a155255, []int{1}
}

func (m *ActiveRawUdpListenerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActiveRawUdpListenerConfig.Unmarshal(m, b)
}
func (m *ActiveRawUdpListenerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActiveRawUdpListenerConfig.Marshal(b, m, deterministic)
}
func (m *ActiveRawUdpListenerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActiveRawUdpListenerConfig.Merge(m, src)
}
func (m *ActiveRawUdpListenerConfig) XXX_Size() int {
	return xxx_messageInfo_ActiveRawUdpListenerConfig.Size(m)
}
func (m *ActiveRawUdpListenerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ActiveRawUdpListenerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ActiveRawUdpListenerConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UdpListenerConfig)(nil), "envoy.config.listener.v3.UdpListenerConfig")
	proto.RegisterType((*ActiveRawUdpListenerConfig)(nil), "envoy.config.listener.v3.ActiveRawUdpListenerConfig")
}

func init() {
	proto.RegisterFile("envoy/config/listener/v3/udp_listener_config.proto", fileDescriptor_4ac914914a155255)
}

var fileDescriptor_4ac914914a155255 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x18, 0x85, 0xd7, 0x09, 0x82, 0x99, 0x22, 0x1b, 0xa2, 0x75, 0xca, 0x98, 0xbb, 0xd0, 0xe9, 0x45,
	0x22, 0xdd, 0x85, 0xb8, 0xbb, 0x55, 0x85, 0x5d, 0x88, 0x8c, 0xca, 0xae, 0x4b, 0xd6, 0x64, 0x33,
	0xb0, 0x25, 0xa1, 0x4d, 0xab, 0x7d, 0x03, 0x9f, 0x41, 0xf0, 0x71, 0x7c, 0x2f, 0x69, 0x92, 0x59,
	0xb4, 0xee, 0xb2, 0xfd, 0xcf, 0xf9, 0xfe, 0x73, 0xfe, 0x00, 0x8f, 0xf2, 0x4c, 0xe4, 0x28, 0x12,
	0x7c, 0xce, 0x16, 0x68, 0xc9, 0x12, 0x45, 0x39, 0x8d, 0x51, 0x36, 0x40, 0x29, 0x91, 0xe1, 0xfa,
	0x3b, 0x34, 0x73, 0x28, 0x63, 0xa1, 0x44, 0xcb, 0xd5, 0x1e, 0x68, 0xff, 0xad, 0x35, 0x30, 0x1b,
	0xb4, 0x8f, 0x17, 0x42, 0x2c, 0x96, 0x14, 0x69, 0xdd, 0x2c, 0x9d, 0x23, 0xcc, 0x73, 0x63, 0x6a,
	0x9f, 0xfe, 0x1d, 0x25, 0x2a, 0x4e, 0x23, 0x65, 0xa7, 0x67, 0x29, 0x91, 0x18, 0x61, 0xce, 0x85,
	0xc2, 0x8a, 0x09, 0x9e, 0xa0, 0x8c, 0xc6, 0x09, 0x13, 0x9c, 0x71, 0xbb, 0xb5, 0xf7, 0x59, 0x07,
	0xcd, 0x29, 0x91, 0x8f, 0x76, 0xdd, 0x9d, 0xde, 0xde, 0xba, 0x02, 0xcd, 0x5f, 0x41, 0x39, 0x5e,
	0x51, 0xd7, 0xe9, 0x3a, 0xfd, 0x9d, 0x60, 0x3f, 0x2d, 0xd5, 0x4f, 0x78, 0x45, 0x5b, 0x33, 0xd0,
	0x79, 0x61, 0x84, 0x50, 0x1e, 0xea, 0x02, 0x21, 0xa1, 0x32, 0xa6, 0x11, 0x56, 0x94, 0xd8, 0x7e,
	0x6e, 0xbd, 0xeb, 0xf4, 0x1b, 0xde, 0x11, 0x34, 0x59, 0xe1, 0x3a, 0x2b, 0x7c, 0xd6, 0x59, 0xfd,
	0xba, 0xeb, 0x8c, 0x6b, 0xc1, 0x89, 0x81, 0x3c, 0x14, 0x8c, 0xfb, 0x1f, 0x84, 0xcd, 0x73, 0x0b,
	0x76, 0x55, 0x2e, 0x4b, 0xe2, 0x96, 0x26, 0x1e, 0x54, 0x88, 0x23, 0x9e, 0x8f, 0x6b, 0x41, 0x43,
	0x6b, 0x8d, 0x75, 0x08, 0x3f, 0xbe, 0xde, 0x3b, 0x97, 0xe0, 0xc2, 0x5c, 0x17, 0x4b, 0x06, 0x33,
	0xaf, 0xbc, 0x6e, 0xa5, 0xba, 0xbf, 0x07, 0x1a, 0x66, 0x49, 0x58, 0x50, 0x7a, 0x53, 0xd0, 0x1e,
	0x45, 0x8a, 0x65, 0x34, 0xc0, 0xaf, 0x15, 0xf1, 0xf0, 0xa6, 0x80, 0x7b, 0xe0, 0xfa, 0x7f, 0xf8,
	0x66, 0xa3, 0xef, 0x83, 0x73, 0x26, 0xa0, 0xb6, 0xc9, 0x58, 0xbc, 0xe5, 0x70, 0xd3, 0xe3, 0xfb,
	0x87, 0x15, 0xf3, 0xa4, 0x68, 0x3b, 0x71, 0x66, 0xdb, 0xba, 0xf6, 0xe0, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x33, 0x7f, 0xab, 0xd7, 0x6d, 0x02, 0x00, 0x00,
}