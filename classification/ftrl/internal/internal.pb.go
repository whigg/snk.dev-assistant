// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: classification/ftrl/internal/internal.proto

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	classification/ftrl/internal/internal.proto

It has these top-level messages:
	Optimizer
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import blacksquaremedia_reason_core "github.com/bsm/reason/core"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Optimizer wraps the optimizer data.
type Optimizer struct {
	// The underlying model.
	Model *blacksquaremedia_reason_core.Model `protobuf:"bytes,1,opt,name=model" json:"model,omitempty"`
	// The target feature.
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// The gradient sums.
	Sums []float64 `protobuf:"fixed64,3,rep,packed,name=sums" json:"sums,omitempty"`
	// The weights.
	Weights []float64 `protobuf:"fixed64,4,rep,packed,name=weights" json:"weights,omitempty"`
}

func (m *Optimizer) Reset()                    { *m = Optimizer{} }
func (m *Optimizer) String() string            { return proto.CompactTextString(m) }
func (*Optimizer) ProtoMessage()               {}
func (*Optimizer) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

func init() {
	proto.RegisterType((*Optimizer)(nil), "blacksquaremedia.reason.classification.ftrl.Optimizer")
}

func init() { proto.RegisterFile("classification/ftrl/internal/internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3f, 0x4b, 0x03, 0x31,
	0x18, 0xc6, 0x1b, 0x5b, 0xab, 0x8d, 0x5b, 0x06, 0x09, 0x1d, 0x8e, 0x43, 0x11, 0x0e, 0x8b, 0x09,
	0xe8, 0xe4, 0xda, 0xc1, 0xc9, 0x7f, 0x1c, 0x4e, 0x6e, 0x49, 0x9a, 0x4b, 0x83, 0x97, 0x7b, 0x6b,
	0x92, 0x43, 0xf0, 0x33, 0x38, 0xf8, 0xb1, 0xfc, 0x08, 0xae, 0xc5, 0x2f, 0x22, 0x97, 0xb3, 0xa8,
	0x83, 0xe0, 0x12, 0x9e, 0xe7, 0xe5, 0x97, 0x27, 0x79, 0x1f, 0x3c, 0x53, 0xb5, 0x08, 0xc1, 0x56,
	0x56, 0x89, 0x68, 0xa1, 0xe1, 0x55, 0xf4, 0x35, 0xb7, 0x4d, 0xd4, 0xbe, 0x11, 0xdf, 0x82, 0xad,
	0x3c, 0x44, 0x20, 0x33, 0x59, 0x0b, 0xf5, 0x10, 0x1e, 0x5b, 0xe1, 0xb5, 0xd3, 0x0b, 0x2b, 0x98,
	0xd7, 0x22, 0x40, 0xc3, 0x7e, 0x87, 0xb0, 0x2e, 0x64, 0x7a, 0x64, 0x6c, 0x5c, 0xb6, 0x92, 0x29,
	0x70, 0x5c, 0x06, 0xc7, 0x7b, 0x94, 0x2b, 0xf0, 0x3a, 0x1d, 0x7d, 0xe6, 0xf4, 0xe4, 0x07, 0x66,
	0xc0, 0x00, 0x4f, 0x63, 0xd9, 0x56, 0xc9, 0x25, 0x93, 0x54, 0x8f, 0x1f, 0xbc, 0x20, 0x3c, 0xb9,
	0x59, 0x45, 0xeb, 0xec, 0xb3, 0xf6, 0xe4, 0x1c, 0x6f, 0x3b, 0x58, 0xe8, 0x9a, 0xa2, 0x1c, 0x15,
	0x7b, 0xa7, 0x87, 0xec, 0xcf, 0x0f, 0x76, 0x0f, 0x5e, 0x75, 0x68, 0xd9, 0xdf, 0x20, 0xfb, 0x78,
	0x1c, 0x85, 0x37, 0x3a, 0xd2, 0xad, 0x1c, 0x15, 0x93, 0xf2, 0xcb, 0x11, 0x82, 0x47, 0xa1, 0x75,
	0x81, 0x0e, 0xf3, 0x61, 0x81, 0xca, 0xa4, 0x09, 0xc5, 0x3b, 0x4f, 0xda, 0x9a, 0x65, 0x0c, 0x74,
	0x94, 0xc6, 0x1b, 0x3b, 0xbf, 0x7e, 0x5b, 0x67, 0x83, 0xf7, 0x75, 0x86, 0x5e, 0x3f, 0xb2, 0x01,
	0x3e, 0x56, 0xe0, 0xd8, 0xff, 0x3a, 0x9a, 0xe3, 0x8b, 0xbb, 0xf2, 0xf2, 0xb6, 0xdb, 0x29, 0xdc,
	0xef, 0x6e, 0x7a, 0x96, 0xe3, 0xb4, 0xe5, 0xd9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xbd,
	0x36, 0xef, 0x97, 0x01, 0x00, 0x00,
}
