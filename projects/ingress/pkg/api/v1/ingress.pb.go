// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

//
//A simple wrapper for a Kubernetes Ingress Object.
type Ingress struct {
	// a raw byte representation of the kubernetes ingress this resource wraps
	KubeIngressSpec *types.Any `protobuf:"bytes,1,opt,name=kube_ingress_spec,json=kubeIngressSpec,proto3" json:"kube_ingress_spec,omitempty"`
	// a raw byte representation of the ingress status of the kubernetes ingress object
	KubeIngressStatus *types.Any `protobuf:"bytes,2,opt,name=kube_ingress_status,json=kubeIngressStatus,proto3" json:"kube_ingress_status,omitempty"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Ingress) Reset()         { *m = Ingress{} }
func (m *Ingress) String() string { return proto.CompactTextString(m) }
func (*Ingress) ProtoMessage()    {}
func (*Ingress) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e3857ca3a6e6b32, []int{0}
}
func (m *Ingress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ingress.Unmarshal(m, b)
}
func (m *Ingress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ingress.Marshal(b, m, deterministic)
}
func (m *Ingress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ingress.Merge(m, src)
}
func (m *Ingress) XXX_Size() int {
	return xxx_messageInfo_Ingress.Size(m)
}
func (m *Ingress) XXX_DiscardUnknown() {
	xxx_messageInfo_Ingress.DiscardUnknown(m)
}

var xxx_messageInfo_Ingress proto.InternalMessageInfo

func (m *Ingress) GetKubeIngressSpec() *types.Any {
	if m != nil {
		return m.KubeIngressSpec
	}
	return nil
}

func (m *Ingress) GetKubeIngressStatus() *types.Any {
	if m != nil {
		return m.KubeIngressStatus
	}
	return nil
}

func (m *Ingress) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*Ingress)(nil), "ingress.solo.io.Ingress")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/ingress/api/v1/ingress.proto", fileDescriptor_7e3857ca3a6e6b32)
}

var fileDescriptor_7e3857ca3a6e6b32 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0xfd, 0x52, 0x59, 0x5f, 0xc1, 0x1d, 0xaa, 0x9a, 0x0a, 0x4a, 0x87, 0x82, 0x98, 0x58, 0xb0,
	0x05, 0x5d, 0x10, 0x52, 0x25, 0xe8, 0x06, 0x12, 0x4b, 0xd9, 0x58, 0x2a, 0x27, 0x18, 0x63, 0x92,
	0xe6, 0x5a, 0xb1, 0x83, 0xda, 0xb5, 0x4f, 0xc3, 0x23, 0xf0, 0x08, 0x3c, 0x05, 0x03, 0x23, 0x5b,
	0x07, 0x76, 0x14, 0xc7, 0x46, 0x30, 0x54, 0x62, 0xf3, 0xf1, 0xb9, 0xe7, 0xc7, 0xd7, 0x78, 0x24,
	0x95, 0x7d, 0x28, 0x63, 0x9a, 0xc0, 0x8c, 0x19, 0xc8, 0xe0, 0x48, 0x01, 0x93, 0x19, 0x00, 0xd3,
	0x05, 0x3c, 0x8a, 0xc4, 0x1a, 0xa6, 0x72, 0x59, 0x08, 0x63, 0x18, 0xd7, 0x8a, 0x3d, 0x1d, 0x07,
	0x48, 0x75, 0x01, 0x16, 0x48, 0x3b, 0xc0, 0x4a, 0x4b, 0x15, 0xf4, 0xbb, 0x12, 0x24, 0x38, 0x8e,
	0x55, 0xa7, 0x7a, 0xac, 0xbf, 0x2b, 0x01, 0x64, 0x26, 0x98, 0x43, 0x71, 0x79, 0xcf, 0x78, 0xbe,
	0xf0, 0xd4, 0xc0, 0xa5, 0xa6, 0xca, 0x86, 0x80, 0x99, 0xb0, 0xfc, 0x8e, 0x5b, 0xbe, 0x8e, 0x0f,
	0xd8, 0xf3, 0x44, 0xcc, 0x6d, 0x9d, 0x27, 0xe6, 0xfe, 0xee, 0xe0, 0x23, 0xc2, 0xcd, 0xcb, 0xba,
	0x18, 0x39, 0xc7, 0x9d, 0xb4, 0x8c, 0xc5, 0xd4, 0x17, 0x9d, 0x1a, 0x2d, 0x92, 0x5e, 0xb4, 0x1f,
	0x1d, 0xb6, 0x4e, 0xba, 0xb4, 0xae, 0x45, 0x43, 0x2d, 0x7a, 0x91, 0x2f, 0x26, 0xed, 0x6a, 0xdc,
	0xab, 0x6f, 0xb4, 0x48, 0xc8, 0x15, 0xde, 0xfa, 0xed, 0x60, 0xb9, 0x2d, 0x4d, 0xaf, 0xb1, 0xde,
	0x63, 0x8c, 0x5e, 0x3e, 0x51, 0x34, 0xe9, 0xfc, 0x74, 0x72, 0x22, 0x72, 0x8a, 0x37, 0xc2, 0xfb,
	0x7a, 0x4d, 0x67, 0xb0, 0x4d, 0x13, 0x28, 0x44, 0xd8, 0x1f, 0xbd, 0xf6, 0xec, 0x18, 0xbd, 0xbe,
	0xed, 0xfd, 0x9b, 0x7c, 0x4f, 0x9f, 0xed, 0x2c, 0x57, 0x08, 0xe1, 0x86, 0x92, 0xcb, 0x15, 0x6a,
	0x91, 0x4d, 0x5f, 0x46, 0x98, 0xf1, 0xa8, 0xca, 0x7a, 0x7e, 0x1f, 0x44, 0xb7, 0xc3, 0x3f, 0xff,
	0xa5, 0x4e, 0xa5, 0x5f, 0x67, 0xfc, 0xdf, 0x15, 0x1f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x3c,
	0x98, 0x86, 0x6c, 0x09, 0x02, 0x00, 0x00,
}

func (this *Ingress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ingress)
	if !ok {
		that2, ok := that.(Ingress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.KubeIngressSpec.Equal(that1.KubeIngressSpec) {
		return false
	}
	if !this.KubeIngressStatus.Equal(that1.KubeIngressStatus) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
