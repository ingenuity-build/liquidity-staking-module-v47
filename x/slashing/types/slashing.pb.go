// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/slashing/v1beta1/slashing.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-proto/rapidproto"
	proto "github.com/cosmos/cosmos-proto"
	github_com_gogo_protobuf_types "github.com/cosmos/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
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

// ValidatorSigningInfo defines a validator's signing info for monitoring their
// liveness activity.
type ValidatorSigningInfo struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Height at which validator was first a candidate OR was unjailed
	StartHeight int64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty" yaml:"start_height"`
	// Index which is incremented each time the validator was a bonded
	// in a block and may have signed a precommit or not. This in conjunction with the
	// `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`.
	IndexOffset int64 `protobuf:"varint,3,opt,name=index_offset,json=indexOffset,proto3" json:"index_offset,omitempty" yaml:"index_offset"`
	// Timestamp until which the validator is jailed due to liveness downtime.
	JailedUntil time.Time `protobuf:"bytes,4,opt,name=jailed_until,json=jailedUntil,proto3,stdtime" json:"jailed_until" yaml:"jailed_until"`
	// Whether or not a validator has been tombstoned (killed out of validator set). It is set
	// once the validator commits an equivocation or for any other configured misbehiavor.
	Tombstoned bool `protobuf:"varint,5,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	// A counter kept to avoid unnecessary array reads.
	// Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`.
	MissedBlocksCounter int64 `protobuf:"varint,6,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter,omitempty" yaml:"missed_blocks_counter"`
}

func (m *ValidatorSigningInfo) Reset()      { *m = ValidatorSigningInfo{} }
func (*ValidatorSigningInfo) ProtoMessage() {}
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1078e5d96a74cc52, []int{0}
}
func (m *ValidatorSigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfo.Merge(m, src)
}
func (m *ValidatorSigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfo proto.InternalMessageInfo

func (m *ValidatorSigningInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorSigningInfo) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *ValidatorSigningInfo) GetIndexOffset() int64 {
	if m != nil {
		return m.IndexOffset
	}
	return 0
}

func (m *ValidatorSigningInfo) GetJailedUntil() time.Time {
	if m != nil {
		return m.JailedUntil
	}
	return time.Time{}
}

func (m *ValidatorSigningInfo) GetTombstoned() bool {
	if m != nil {
		return m.Tombstoned
	}
	return false
}

func (m *ValidatorSigningInfo) GetMissedBlocksCounter() int64 {
	if m != nil {
		return m.MissedBlocksCounter
	}
	return 0
}

// Params represents the parameters used for by the slashing module.
type Params struct {
	SignedBlocksWindow      int64                                  `protobuf:"varint,1,opt,name=signed_blocks_window,json=signedBlocksWindow,proto3" json:"signed_blocks_window,omitempty" yaml:"signed_blocks_window"`
	MinSignedPerWindow      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=min_signed_per_window,json=minSignedPerWindow,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_signed_per_window" yaml:"min_signed_per_window"`
	DowntimeJailDuration    time.Duration                          `protobuf:"bytes,3,opt,name=downtime_jail_duration,json=downtimeJailDuration,proto3,stdduration" json:"downtime_jail_duration" yaml:"downtime_jail_duration"`
	SlashFractionDoubleSign github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=slash_fraction_double_sign,json=slashFractionDoubleSign,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_double_sign" yaml:"slash_fraction_double_sign"`
	SlashFractionDowntime   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=slash_fraction_downtime,json=slashFractionDowntime,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_downtime" yaml:"slash_fraction_downtime"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1078e5d96a74cc52, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetSignedBlocksWindow() int64 {
	if m != nil {
		return m.SignedBlocksWindow
	}
	return 0
}

func (m *Params) GetDowntimeJailDuration() time.Duration {
	if m != nil {
		return m.DowntimeJailDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*ValidatorSigningInfo)(nil), "liquidstaking.slashing.v1beta1.ValidatorSigningInfo")
	proto.RegisterType((*Params)(nil), "liquidstaking.slashing.v1beta1.Params")
}

func init() {
	proto.RegisterFile("cosmos/slashing/v1beta1/slashing.proto", fileDescriptor_1078e5d96a74cc52)
}

var fileDescriptor_1078e5d96a74cc52 = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xbd, 0x6f, 0xd3, 0x4e,
	0x18, 0xce, 0xfd, 0xd2, 0x5f, 0x29, 0x97, 0x4c, 0x6e, 0x4a, 0x4c, 0x00, 0x3b, 0x78, 0xa8, 0xc2,
	0x50, 0x5b, 0x2d, 0x5b, 0x47, 0x53, 0x21, 0x3e, 0x24, 0x28, 0x6e, 0x01, 0x89, 0x01, 0xeb, 0x9c,
	0xbb, 0x38, 0x47, 0xed, 0xbb, 0xc8, 0x77, 0xa1, 0x2d, 0x1b, 0x5b, 0xc7, 0x8e, 0x1d, 0x3b, 0xf2,
	0xa7, 0x74, 0xec, 0x88, 0x18, 0x02, 0x4a, 0x17, 0xe6, 0x6e, 0x6c, 0xc8, 0x77, 0x76, 0x1b, 0xa5,
	0x29, 0x52, 0xa7, 0xe4, 0x7d, 0xde, 0xe7, 0x7d, 0xee, 0x79, 0x3f, 0x12, 0xb8, 0xdc, 0xe5, 0x22,
	0xe5, 0xc2, 0x13, 0x09, 0x12, 0x7d, 0xca, 0x62, 0xef, 0xf3, 0x6a, 0x44, 0x24, 0x5a, 0xbd, 0x00,
	0xdc, 0x41, 0xc6, 0x25, 0x37, 0x9a, 0x9a, 0xe7, 0x5e, 0xc0, 0x05, 0xaf, 0xd5, 0x88, 0x79, 0xcc,
	0x15, 0xc7, 0xcb, 0xbf, 0x69, 0x7a, 0xcb, 0x8a, 0x39, 0x8f, 0x13, 0xe2, 0xa9, 0x28, 0x1a, 0xf6,
	0x3c, 0x3c, 0xcc, 0x90, 0xa4, 0x9c, 0x15, 0x79, 0x7b, 0x3a, 0x2f, 0x69, 0x4a, 0x84, 0x44, 0xe9,
	0x40, 0x13, 0x9c, 0x83, 0x2a, 0x6c, 0xbc, 0x43, 0x09, 0xc5, 0x48, 0xf2, 0x6c, 0x8b, 0xc6, 0x8c,
	0xb2, 0xf8, 0x39, 0xeb, 0x71, 0xc3, 0x84, 0xb7, 0x10, 0xc6, 0x19, 0x11, 0xc2, 0x04, 0x6d, 0xd0,
	0xb9, 0x1d, 0x94, 0xa1, 0xb1, 0x0e, 0xeb, 0x42, 0xa2, 0x4c, 0x86, 0x7d, 0x42, 0xe3, 0xbe, 0x34,
	0xff, 0x6b, 0x83, 0x4e, 0xd5, 0x6f, 0x9e, 0x8f, 0xec, 0xc5, 0x7d, 0x94, 0x26, 0xeb, 0xce, 0x64,
	0xd6, 0x09, 0x6a, 0x2a, 0x7c, 0xa6, 0xa2, 0xbc, 0x96, 0x32, 0x4c, 0xf6, 0x42, 0xde, 0xeb, 0x09,
	0x22, 0xcd, 0xea, 0x74, 0xed, 0x64, 0xd6, 0x09, 0x6a, 0x2a, 0x7c, 0xad, 0x22, 0xe3, 0x23, 0xac,
	0x7f, 0x42, 0x34, 0x21, 0x38, 0x1c, 0x32, 0x49, 0x13, 0x73, 0xae, 0x0d, 0x3a, 0xb5, 0xb5, 0x96,
	0xab, 0x5b, 0x74, 0xcb, 0x16, 0xdd, 0xed, 0xb2, 0x45, 0xdf, 0x3e, 0x19, 0xd9, 0x95, 0x4b, 0xed,
	0xc9, 0x6a, 0xe7, 0xf0, 0xa7, 0x0d, 0x82, 0x9a, 0x86, 0xde, 0xe6, 0x88, 0x61, 0x41, 0x28, 0x79,
	0x1a, 0x09, 0xc9, 0x19, 0xc1, 0xe6, 0xff, 0x6d, 0xd0, 0x59, 0x08, 0x26, 0x10, 0x63, 0x1b, 0x2e,
	0xa5, 0x54, 0x08, 0x82, 0xc3, 0x28, 0xe1, 0xdd, 0x1d, 0x11, 0x76, 0xf9, 0x90, 0x49, 0x92, 0x99,
	0xf3, 0xaa, 0x89, 0xf6, 0xf9, 0xc8, 0xbe, 0xaf, 0x1f, 0x9a, 0x49, 0x73, 0x82, 0x45, 0x8d, 0xfb,
	0x0a, 0x7e, 0xa2, 0xd1, 0xf5, 0x85, 0xa3, 0x63, 0xbb, 0xf2, 0xfb, 0xd8, 0x06, 0xce, 0x9f, 0x39,
	0x38, 0xbf, 0x89, 0x32, 0x94, 0x0a, 0xe3, 0x0d, 0x6c, 0x08, 0x1a, 0xb3, 0x4b, 0x8d, 0x5d, 0xca,
	0x30, 0xdf, 0x55, 0x9b, 0xa8, 0xfa, 0xf6, 0xf9, 0xc8, 0xbe, 0x57, 0x8c, 0x7a, 0x06, 0xcb, 0x09,
	0x0c, 0x0d, 0xeb, 0x87, 0xde, 0x2b, 0xd0, 0xf8, 0x0a, 0x72, 0xfb, 0x2c, 0x2c, 0x2a, 0x06, 0x24,
	0x2b, 0x45, 0xf3, 0xfd, 0xd5, 0xfd, 0x57, 0xf9, 0xac, 0x7e, 0x8c, 0xec, 0xe5, 0x98, 0xca, 0xfe,
	0x30, 0x72, 0xbb, 0x3c, 0xf5, 0x8a, 0x9b, 0xd5, 0x1f, 0x2b, 0x02, 0xef, 0x78, 0x72, 0x7f, 0x40,
	0x84, 0xbb, 0x41, 0xba, 0x93, 0xcd, 0xce, 0x10, 0x75, 0x02, 0x23, 0xa5, 0x6c, 0x4b, 0xc1, 0x9b,
	0x24, 0x2b, 0x3c, 0x7c, 0x81, 0x77, 0x30, 0xdf, 0x65, 0xf9, 0x0d, 0x86, 0xf9, 0xe4, 0xc3, 0xf2,
	0x5a, 0xd5, 0x1d, 0xd4, 0xd6, 0xee, 0x5e, 0xd9, 0xe5, 0x46, 0x41, 0xf0, 0x1f, 0x15, 0xab, 0x7c,
	0xa0, 0x1f, 0x9d, 0x2d, 0xe3, 0x1c, 0xe5, 0x4b, 0x6d, 0x94, 0xc9, 0x17, 0x88, 0x26, 0xa5, 0x80,
	0x71, 0x08, 0x60, 0x4b, 0xfd, 0xa8, 0xc2, 0x5e, 0x86, 0xba, 0x39, 0x14, 0x62, 0x3e, 0x8c, 0x12,
	0xa2, 0xcc, 0xab, 0x63, 0xaa, 0xfb, 0x5b, 0x37, 0x1e, 0xc2, 0xc3, 0x62, 0x0f, 0xd7, 0x2a, 0x3b,
	0x41, 0x53, 0x25, 0x9f, 0x16, 0xb9, 0x0d, 0x95, 0xca, 0x27, 0x63, 0x1c, 0x00, 0xd8, 0xbc, 0x52,
	0xa8, 0xad, 0xab, 0xf3, 0xab, 0xfb, 0x9b, 0x37, 0xf6, 0x63, 0x5d, 0xe3, 0x47, 0xcb, 0x3a, 0xc1,
	0xd2, 0x94, 0x19, 0x8d, 0xfb, 0x2f, 0xbf, 0x8d, 0x2d, 0x70, 0x32, 0xb6, 0xc0, 0xe9, 0xd8, 0x02,
	0xbf, 0xc6, 0x16, 0x38, 0x3c, 0xb3, 0x2a, 0xa7, 0x67, 0x56, 0xe5, 0xfb, 0x99, 0x55, 0xf9, 0xb0,
	0xf2, 0xcf, 0xe7, 0xf7, 0x2e, 0xff, 0xd4, 0x94, 0x93, 0x68, 0x5e, 0xad, 0xef, 0xf1, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7b, 0x78, 0xfd, 0x40, 0xf4, 0x04, 0x00, 0x00,
}

func (this *ValidatorSigningInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSigningInfo)
	if !ok {
		that2, ok := that.(ValidatorSigningInfo)
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
	if this.Address != that1.Address {
		return false
	}
	if this.StartHeight != that1.StartHeight {
		return false
	}
	if this.IndexOffset != that1.IndexOffset {
		return false
	}
	if !this.JailedUntil.Equal(that1.JailedUntil) {
		return false
	}
	if this.Tombstoned != that1.Tombstoned {
		return false
	}
	if this.MissedBlocksCounter != that1.MissedBlocksCounter {
		return false
	}
	return true
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.SignedBlocksWindow != that1.SignedBlocksWindow {
		return false
	}
	if !this.MinSignedPerWindow.Equal(that1.MinSignedPerWindow) {
		return false
	}
	if this.DowntimeJailDuration != that1.DowntimeJailDuration {
		return false
	}
	if !this.SlashFractionDoubleSign.Equal(that1.SlashFractionDoubleSign) {
		return false
	}
	if !this.SlashFractionDowntime.Equal(that1.SlashFractionDowntime) {
		return false
	}
	return true
}
func (m *ValidatorSigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissedBlocksCounter != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.MissedBlocksCounter))
		i--
		dAtA[i] = 0x30
	}
	if m.Tombstoned {
		i--
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.JailedUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSlashing(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.IndexOffset != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.IndexOffset))
		i--
		dAtA[i] = 0x18
	}
	if m.StartHeight != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSlashing(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SlashFractionDowntime.Size()
		i -= size
		if _, err := m.SlashFractionDowntime.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSlashing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SlashFractionDoubleSign.Size()
		i -= size
		if _, err := m.SlashFractionDoubleSign.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSlashing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DowntimeJailDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DowntimeJailDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSlashing(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinSignedPerWindow.Size()
		i -= size
		if _, err := m.MinSignedPerWindow.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSlashing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SignedBlocksWindow != 0 {
		i = encodeVarintSlashing(dAtA, i, uint64(m.SignedBlocksWindow))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSlashing(dAtA []byte, offset int, v uint64) int {
	offset -= sovSlashing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorSigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSlashing(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovSlashing(uint64(m.StartHeight))
	}
	if m.IndexOffset != 0 {
		n += 1 + sovSlashing(uint64(m.IndexOffset))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil)
	n += 1 + l + sovSlashing(uint64(l))
	if m.Tombstoned {
		n += 2
	}
	if m.MissedBlocksCounter != 0 {
		n += 1 + sovSlashing(uint64(m.MissedBlocksCounter))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedBlocksWindow != 0 {
		n += 1 + sovSlashing(uint64(m.SignedBlocksWindow))
	}
	l = m.MinSignedPerWindow.Size()
	n += 1 + l + sovSlashing(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DowntimeJailDuration)
	n += 1 + l + sovSlashing(uint64(l))
	l = m.SlashFractionDoubleSign.Size()
	n += 1 + l + sovSlashing(uint64(l))
	l = m.SlashFractionDowntime.Size()
	n += 1 + l + sovSlashing(uint64(l))
	return n
}

func sovSlashing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSlashing(x uint64) (n int) {
	return sovSlashing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorSigningInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
			return fmt.Errorf("proto: ValidatorSigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexOffset", wireType)
			}
			m.IndexOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexOffset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.JailedUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
			m.Tombstoned = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocksCounter", wireType)
			}
			m.MissedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSlashing
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashing
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedBlocksWindow", wireType)
			}
			m.SignedBlocksWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedBlocksWindow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSignedPerWindow", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSignedPerWindow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DowntimeJailDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
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
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DowntimeJailDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionDoubleSign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionDoubleSign.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionDowntime", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashing
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSlashing
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionDowntime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSlashing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSlashing
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSlashing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlashing
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
					return 0, ErrIntOverflowSlashing
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
					return 0, ErrIntOverflowSlashing
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
				return 0, ErrInvalidLengthSlashing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSlashing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSlashing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSlashing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlashing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSlashing = fmt.Errorf("proto: unexpected end of group")
)
