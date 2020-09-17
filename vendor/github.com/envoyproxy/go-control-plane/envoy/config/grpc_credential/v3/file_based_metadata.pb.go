// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/grpc_credential/v3/file_based_metadata.proto

package envoy_config_grpc_credential_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
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

type FileBasedMetadataConfig struct {
	SecretData           *v3.DataSource `protobuf:"bytes,1,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	HeaderKey            string         `protobuf:"bytes,2,opt,name=header_key,json=headerKey,proto3" json:"header_key,omitempty"`
	HeaderPrefix         string         `protobuf:"bytes,3,opt,name=header_prefix,json=headerPrefix,proto3" json:"header_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FileBasedMetadataConfig) Reset()         { *m = FileBasedMetadataConfig{} }
func (m *FileBasedMetadataConfig) String() string { return proto.CompactTextString(m) }
func (*FileBasedMetadataConfig) ProtoMessage()    {}
func (*FileBasedMetadataConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_413c1287d6760a42, []int{0}
}

func (m *FileBasedMetadataConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileBasedMetadataConfig.Unmarshal(m, b)
}
func (m *FileBasedMetadataConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileBasedMetadataConfig.Marshal(b, m, deterministic)
}
func (m *FileBasedMetadataConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileBasedMetadataConfig.Merge(m, src)
}
func (m *FileBasedMetadataConfig) XXX_Size() int {
	return xxx_messageInfo_FileBasedMetadataConfig.Size(m)
}
func (m *FileBasedMetadataConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FileBasedMetadataConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FileBasedMetadataConfig proto.InternalMessageInfo

func (m *FileBasedMetadataConfig) GetSecretData() *v3.DataSource {
	if m != nil {
		return m.SecretData
	}
	return nil
}

func (m *FileBasedMetadataConfig) GetHeaderKey() string {
	if m != nil {
		return m.HeaderKey
	}
	return ""
}

func (m *FileBasedMetadataConfig) GetHeaderPrefix() string {
	if m != nil {
		return m.HeaderPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*FileBasedMetadataConfig)(nil), "envoy.config.grpc_credential.v3.FileBasedMetadataConfig")
}

func init() {
	proto.RegisterFile("envoy/config/grpc_credential/v3/file_based_metadata.proto", fileDescriptor_413c1287d6760a42)
}

var fileDescriptor_413c1287d6760a42 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x99, 0x0a, 0x85, 0xa6, 0xba, 0x99, 0x85, 0x96, 0x82, 0x76, 0xd4, 0x4d, 0x37, 0x66,
	0xa0, 0xb3, 0x52, 0xc4, 0xc5, 0x54, 0x84, 0x22, 0x42, 0xa9, 0x07, 0x18, 0x5e, 0x67, 0x5e, 0xdb,
	0xe0, 0x98, 0x84, 0x24, 0x0d, 0x9d, 0x1b, 0x08, 0xde, 0xc0, 0xb5, 0x77, 0xd0, 0x0b, 0x78, 0x1c,
	0xef, 0x20, 0x99, 0xd4, 0x45, 0x2d, 0xea, 0xf6, 0x7f, 0xff, 0xff, 0xe5, 0x7f, 0x2f, 0xe4, 0x1c,
	0xb9, 0x15, 0x55, 0x9c, 0x0b, 0x3e, 0x63, 0xf3, 0x78, 0xae, 0x64, 0x9e, 0xe5, 0x0a, 0x0b, 0xe4,
	0x86, 0x41, 0x19, 0xdb, 0x24, 0x9e, 0xb1, 0x12, 0xb3, 0x29, 0x68, 0x2c, 0xb2, 0x47, 0x34, 0x50,
	0x80, 0x01, 0x2a, 0x95, 0x30, 0x22, 0xec, 0xd5, 0x51, 0xea, 0xa3, 0xf4, 0x47, 0x94, 0xda, 0xa4,
	0xdb, 0xdb, 0x60, 0xe7, 0x42, 0xa1, 0x03, 0x3a, 0x96, 0x27, 0x74, 0xa3, 0x65, 0x21, 0x21, 0x06,
	0xce, 0x85, 0x01, 0xc3, 0x04, 0xd7, 0xb1, 0x46, 0xae, 0x99, 0x61, 0xf6, 0xdb, 0x71, 0xbc, 0xe5,
	0xb0, 0xa8, 0x34, 0x13, 0x9c, 0xf1, 0xb9, 0xb7, 0x9c, 0x7c, 0x06, 0xe4, 0xe0, 0x86, 0x95, 0x98,
	0xba, 0x8e, 0x77, 0xeb, 0x8a, 0xc3, 0xfa, 0xcd, 0x70, 0x44, 0xda, 0x1a, 0x73, 0x85, 0x26, 0x73,
	0x62, 0x27, 0x88, 0x82, 0x7e, 0x7b, 0x10, 0xd1, 0x8d, 0xe2, 0xae, 0x17, 0xb5, 0x09, 0xbd, 0x06,
	0x03, 0xf7, 0x62, 0xa9, 0x72, 0x4c, 0x9b, 0xef, 0x6f, 0xcf, 0xaf, 0x8d, 0x60, 0x42, 0x7c, 0xd8,
	0x4d, 0xc2, 0x43, 0x42, 0x16, 0x08, 0x05, 0xaa, 0xec, 0x01, 0xab, 0x4e, 0x23, 0x0a, 0xfa, 0xad,
	0x49, 0xcb, 0x2b, 0xb7, 0x58, 0x85, 0xa7, 0x64, 0x6f, 0x3d, 0x96, 0x0a, 0x67, 0x6c, 0xd5, 0xd9,
	0xa9, 0x1d, 0xbb, 0x5e, 0x1c, 0xd7, 0xda, 0xc5, 0xf0, 0xe5, 0xe3, 0xe9, 0xe8, 0x8a, 0x5c, 0xfe,
	0x7d, 0xb8, 0x01, 0x94, 0x72, 0x01, 0xf4, 0x97, 0x9d, 0xd2, 0x11, 0x39, 0x63, 0xc2, 0xaf, 0x20,
	0x95, 0x58, 0x55, 0xf4, 0x9f, 0x6f, 0x48, 0xf7, 0xb7, 0x48, 0x63, 0x77, 0xb8, 0x71, 0x30, 0x6d,
	0xd6, 0x17, 0x4c, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x5c, 0xa2, 0xfd, 0x05, 0x02, 0x00,
	0x00,
}