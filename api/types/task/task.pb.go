// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/xiaofeng0303/containerd/api/types/task/task.proto

package task

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	StatusUnknown Status = 0
	StatusCreated Status = 1
	StatusRunning Status = 2
	StatusStopped Status = 3
	StatusPaused  Status = 4
	StatusPausing Status = 5
)

var Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "CREATED",
	2: "RUNNING",
	3: "STOPPED",
	4: "PAUSED",
	5: "PAUSING",
}

var Status_value = map[string]int32{
	"UNKNOWN": 0,
	"CREATED": 1,
	"RUNNING": 2,
	"STOPPED": 3,
	"PAUSED":  4,
	"PAUSING": 5,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5130e2b447e52321, []int{0}
}

type Process struct {
	ContainerID          string    `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ID                   string    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Pid                  uint32    `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Status               Status    `protobuf:"varint,4,opt,name=status,proto3,enum=containerd.v1.types.Status" json:"status,omitempty"`
	Stdin                string    `protobuf:"bytes,5,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout               string    `protobuf:"bytes,6,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr               string    `protobuf:"bytes,7,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Terminal             bool      `protobuf:"varint,8,opt,name=terminal,proto3" json:"terminal,omitempty"`
	ExitStatus           uint32    `protobuf:"varint,9,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	ExitedAt             time.Time `protobuf:"bytes,10,opt,name=exited_at,json=exitedAt,proto3,stdtime" json:"exited_at"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Process) Reset()      { *m = Process{} }
func (*Process) ProtoMessage() {}
func (*Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_5130e2b447e52321, []int{0}
}
func (m *Process) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Process.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process.Merge(m, src)
}
func (m *Process) XXX_Size() int {
	return m.Size()
}
func (m *Process) XXX_DiscardUnknown() {
	xxx_messageInfo_Process.DiscardUnknown(m)
}

var xxx_messageInfo_Process proto.InternalMessageInfo

type ProcessInfo struct {
	// PID is the process ID.
	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	// Info contains additional process information.
	//
	// Info varies by platform.
	Info                 *types.Any `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProcessInfo) Reset()      { *m = ProcessInfo{} }
func (*ProcessInfo) ProtoMessage() {}
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5130e2b447e52321, []int{1}
}
func (m *ProcessInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessInfo.Merge(m, src)
}
func (m *ProcessInfo) XXX_Size() int {
	return m.Size()
}
func (m *ProcessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessInfo proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("containerd.v1.types.Status", Status_name, Status_value)
	proto.RegisterType((*Process)(nil), "containerd.v1.types.Process")
	proto.RegisterType((*ProcessInfo)(nil), "containerd.v1.types.ProcessInfo")
}

func init() {
	proto.RegisterFile("github.com/xiaofeng0303/containerd/api/types/task/task.proto", fileDescriptor_5130e2b447e52321)
}

var fileDescriptor_5130e2b447e52321 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x6f, 0x9b, 0x4c,
	0x18, 0xc7, 0x39, 0x92, 0x10, 0xe7, 0x48, 0xf2, 0xf2, 0xd2, 0x28, 0xa2, 0xb4, 0x02, 0x94, 0x09,
	0x75, 0x80, 0xd4, 0x5e, 0xbb, 0x38, 0xb6, 0x55, 0xa1, 0x4a, 0xc4, 0xc2, 0xb6, 0xda, 0xcd, 0xc2,
	0xe6, 0x4c, 0x4f, 0x89, 0xef, 0x10, 0x1c, 0x6d, 0xbc, 0x75, 0xac, 0x32, 0xf5, 0x0b, 0x64, 0x6a,
	0x3f, 0x45, 0x3f, 0x81, 0xc7, 0x4e, 0x55, 0x27, 0xb7, 0xe1, 0x93, 0x54, 0x07, 0xd8, 0xb1, 0xda,
	0x2e, 0x88, 0xe7, 0xf9, 0xfd, 0xee, 0xb9, 0xe7, 0xfe, 0xf0, 0x45, 0x8c, 0xd9, 0xdb, 0x7c, 0xe2,
	0x4c, 0xe9, 0xdc, 0xbd, 0xc1, 0x21, 0x9d, 0x21, 0x12, 0x9f, 0xb7, 0xce, 0x5b, 0xee, 0x94, 0x12,
	0x16, 0x62, 0x82, 0xd2, 0xc8, 0x0d, 0x13, 0xec, 0xb2, 0x45, 0x82, 0x32, 0x97, 0x85, 0xd9, 0x55,
	0xf9, 0x71, 0x92, 0x94, 0x32, 0xaa, 0x3e, 0x7a, 0xb0, 0x9c, 0x77, 0xcf, 0x9d, 0x52, 0xd2, 0x4f,
	0x62, 0x1a, 0xd3, 0x92, 0xbb, 0xfc, 0xaf, 0x52, 0x75, 0x33, 0xa6, 0x34, 0xbe, 0x46, 0x6e, 0x59,
	0x4d, 0xf2, 0x99, 0xcb, 0xf0, 0x1c, 0x65, 0x2c, 0x9c, 0x27, 0xb5, 0xf0, 0xf8, 0x4f, 0x21, 0x24,
	0x8b, 0x0a, 0x9d, 0x15, 0x22, 0xdc, 0xef, 0xa7, 0x74, 0x8a, 0xb2, 0x4c, 0x6d, 0xc2, 0xc3, 0xcd,
	0xa5, 0x63, 0x1c, 0x69, 0xc0, 0x02, 0xf6, 0xc1, 0xc5, 0x7f, 0xc5, 0xca, 0x94, 0x3b, 0xeb, 0xbe,
	0xd7, 0x0d, 0xe4, 0x8d, 0xe4, 0x45, 0xea, 0x29, 0x14, 0x71, 0xa4, 0x89, 0xa5, 0x29, 0x15, 0x2b,
	0x53, 0xf4, 0xba, 0x81, 0x88, 0x23, 0x55, 0x81, 0x3b, 0x09, 0x8e, 0xb4, 0x1d, 0x0b, 0xd8, 0x47,
	0x01, 0xff, 0x55, 0x5b, 0x50, 0xca, 0x58, 0xc8, 0xf2, 0x4c, 0xdb, 0xb5, 0x80, 0x7d, 0xdc, 0x7c,
	0xe2, 0xfc, 0xe3, 0x85, 0xce, 0xa0, 0x54, 0x82, 0x5a, 0x55, 0x4f, 0xe0, 0x5e, 0xc6, 0x22, 0x4c,
	0xb4, 0x3d, 0x7e, 0x43, 0x50, 0x15, 0xea, 0x29, 0x1f, 0x15, 0xd1, 0x9c, 0x69, 0x52, 0xd9, 0xae,
	0xab, 0xba, 0x8f, 0xd2, 0x54, 0xdb, 0xdf, 0xf4, 0x51, 0x9a, 0xaa, 0x3a, 0x6c, 0x30, 0x94, 0xce,
	0x31, 0x09, 0xaf, 0xb5, 0x86, 0x05, 0xec, 0x46, 0xb0, 0xa9, 0x55, 0x13, 0xca, 0xe8, 0x06, 0xb3,
	0x71, 0xbd, 0xdb, 0x41, 0xb9, 0x30, 0xe4, 0xad, 0x6a, 0x15, 0xb5, 0x0d, 0x0f, 0x78, 0x85, 0xa2,
	0x71, 0xc8, 0x34, 0x68, 0x01, 0x5b, 0x6e, 0xea, 0x4e, 0x15, 0xa8, 0xb3, 0x0e, 0xd4, 0x19, 0xae,
	0x13, 0xbf, 0x68, 0x2c, 0x57, 0xa6, 0xf0, 0xe9, 0xa7, 0x09, 0x82, 0x46, 0x75, 0xac, 0xcd, 0xce,
	0x3c, 0x28, 0xd7, 0x19, 0x7b, 0x64, 0x46, 0xd7, 0xd9, 0x80, 0x87, 0x6c, 0x6c, 0xb8, 0x8b, 0xc9,
	0x8c, 0x96, 0x39, 0xca, 0xcd, 0x93, 0xbf, 0xc6, 0xb7, 0xc9, 0x22, 0x28, 0x8d, 0x67, 0xdf, 0x01,
	0x94, 0xea, 0xc5, 0x0c, 0xb8, 0x3f, 0xf2, 0x5f, 0xf9, 0x97, 0xaf, 0x7d, 0x45, 0xd0, 0xff, 0xbf,
	0xbd, 0xb3, 0x8e, 0x2a, 0x30, 0x22, 0x57, 0x84, 0xbe, 0x27, 0x9c, 0x77, 0x82, 0x5e, 0x7b, 0xd8,
	0xeb, 0x2a, 0x60, 0x9b, 0x77, 0x52, 0x14, 0x32, 0x14, 0x71, 0x1e, 0x8c, 0x7c, 0xdf, 0xf3, 0x5f,
	0x2a, 0xe2, 0x36, 0x0f, 0x72, 0x42, 0x30, 0x89, 0x39, 0x1f, 0x0c, 0x2f, 0xfb, 0xfd, 0x5e, 0x57,
	0xd9, 0xd9, 0xe6, 0x03, 0x46, 0x93, 0x04, 0x45, 0xea, 0x53, 0x28, 0xf5, 0xdb, 0xa3, 0x41, 0xaf,
	0xab, 0xec, 0xea, 0xca, 0xed, 0x9d, 0x75, 0x58, 0xe1, 0x7e, 0x98, 0x67, 0xd5, 0x74, 0x4e, 0xf9,
	0xf4, 0xbd, 0xed, 0xd3, 0x1c, 0x63, 0x12, 0xeb, 0xc7, 0x1f, 0x3f, 0x1b, 0xc2, 0xd7, 0x2f, 0x46,
	0xfd, 0x9a, 0x0b, 0x6d, 0x79, 0x6f, 0x08, 0x3f, 0xee, 0x0d, 0xe1, 0x43, 0x61, 0x80, 0x65, 0x61,
	0x80, 0x6f, 0x85, 0x01, 0x7e, 0x15, 0x06, 0x78, 0x23, 0x4c, 0xa4, 0x32, 0x88, 0xd6, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x62, 0x6d, 0xbd, 0x5a, 0x52, 0x03, 0x00, 0x00,
}

func (m *Process) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Process) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Process) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExitedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.ExitedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTask(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	if m.ExitStatus != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.ExitStatus))
		i--
		dAtA[i] = 0x48
	}
	if m.Terminal {
		i--
		if m.Terminal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.Stderr) > 0 {
		i -= len(m.Stderr)
		copy(dAtA[i:], m.Stderr)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stderr)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Stdout) > 0 {
		i -= len(m.Stdout)
		copy(dAtA[i:], m.Stdout)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stdout)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Stdin) > 0 {
		i -= len(m.Stdin)
		copy(dAtA[i:], m.Stdin)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Stdin)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Status != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.Pid != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTask(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContainerID) > 0 {
		i -= len(m.ContainerID)
		copy(dAtA[i:], m.ContainerID)
		i = encodeVarintTask(dAtA, i, uint64(len(m.ContainerID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProcessInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProcessInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTask(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Pid != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTask(dAtA []byte, offset int, v uint64) int {
	offset -= sovTask(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Process) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Pid != 0 {
		n += 1 + sovTask(uint64(m.Pid))
	}
	if m.Status != 0 {
		n += 1 + sovTask(uint64(m.Status))
	}
	l = len(m.Stdin)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stdout)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Stderr)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Terminal {
		n += 2
	}
	if m.ExitStatus != 0 {
		n += 1 + sovTask(uint64(m.ExitStatus))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExitedAt)
	n += 1 + l + sovTask(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ProcessInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pid != 0 {
		n += 1 + sovTask(uint64(m.Pid))
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovTask(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTask(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTask(x uint64) (n int) {
	return sovTask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Process) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Process{`,
		`ContainerID:` + fmt.Sprintf("%v", this.ContainerID) + `,`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Stdin:` + fmt.Sprintf("%v", this.Stdin) + `,`,
		`Stdout:` + fmt.Sprintf("%v", this.Stdout) + `,`,
		`Stderr:` + fmt.Sprintf("%v", this.Stderr) + `,`,
		`Terminal:` + fmt.Sprintf("%v", this.Terminal) + `,`,
		`ExitStatus:` + fmt.Sprintf("%v", this.ExitStatus) + `,`,
		`ExitedAt:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ExitedAt), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessInfo{`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`Info:` + strings.Replace(fmt.Sprintf("%v", this.Info), "Any", "types.Any", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTask(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Process) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Process: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Process: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stdout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stderr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stderr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Terminal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Terminal = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitStatus", wireType)
			}
			m.ExitStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExitStatus |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExitedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProcessInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProcessInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &types.Any{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTask
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTask
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTask
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTask
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTask        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTask          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTask = fmt.Errorf("proto: unexpected end of group")
)
