// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/xiaofeng0303/containerd/protobuf/plugin/fieldpath.proto

package plugin

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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

var E_FieldpathAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63300,
	Name:          "containerd.plugin.fieldpath_all",
	Tag:           "varint,63300,opt,name=fieldpath_all",
	Filename:      "github.com/xiaofeng0303/containerd/protobuf/plugin/fieldpath.proto",
}

var E_Fieldpath = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64400,
	Name:          "containerd.plugin.fieldpath",
	Tag:           "varint,64400,opt,name=fieldpath",
	Filename:      "github.com/xiaofeng0303/containerd/protobuf/plugin/fieldpath.proto",
}

func init() {
	proto.RegisterExtension(E_FieldpathAll)
	proto.RegisterExtension(E_Fieldpath)
}

func init() {
	proto.RegisterFile("github.com/xiaofeng0303/containerd/protobuf/plugin/fieldpath.proto", fileDescriptor_bf62733bfe2faf41)
}

var fileDescriptor_bf62733bfe2faf41 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xaf, 0xc8, 0x4c, 0xcc, 0x4f, 0x4b, 0xcd, 0x4b, 0x37,
	0x30, 0x36, 0x30, 0xd6, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0x4a, 0xd1, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0xd3, 0x4f,
	0xcb, 0x4c, 0xcd, 0x49, 0x29, 0x48, 0x2c, 0xc9, 0xd0, 0x03, 0xcb, 0x08, 0x09, 0x22, 0xd4, 0xea,
	0x41, 0x94, 0x48, 0x29, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0x22, 0xb4, 0xa6, 0xa4, 0x16, 0x27,
	0x17, 0x65, 0x16, 0x94, 0xe4, 0x17, 0x41, 0x34, 0x59, 0x39, 0x73, 0xf1, 0xc2, 0xcd, 0x89, 0x4f,
	0xcc, 0xc9, 0x11, 0x92, 0xd1, 0x83, 0xe8, 0xd1, 0x83, 0xe9, 0xd1, 0x73, 0xcb, 0xcc, 0x49, 0xf5,
	0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b, 0x96, 0x38, 0xf2, 0x8e, 0x59, 0x81, 0x51, 0x83, 0x23, 0x88,
	0x07, 0xae, 0xc9, 0x31, 0x27, 0xc7, 0xca, 0x9e, 0x8b, 0x13, 0xce, 0x17, 0x92, 0xc7, 0x30, 0xc0,
	0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0x6e, 0xc6, 0x84, 0xef, 0x10, 0x33, 0x10, 0x7a, 0x9c, 0x24,
	0x4e, 0x3c, 0x94, 0x63, 0xb8, 0xf1, 0x50, 0x8e, 0xa1, 0xe1, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x08, 0x08, 0x00, 0x00, 0xff, 0xff, 0xbe,
	0x71, 0xab, 0x64, 0x19, 0x01, 0x00, 0x00,
}
