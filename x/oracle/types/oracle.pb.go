// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/oracle.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Params struct {
	VotePeriod        uint64                                 `protobuf:"varint,1,opt,name=vote_period,json=votePeriod,proto3" json:"vote_period,omitempty" yaml:"vote_period"`
	VoteThreshold     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=vote_threshold,json=voteThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"vote_threshold" yaml:"vote_threshold"`
	RewardBand        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=reward_band,json=rewardBand,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_band" yaml:"reward_band"`
	Whitelist         DenomList                              `protobuf:"bytes,4,rep,name=whitelist,proto3,castrepeated=DenomList" json:"whitelist" yaml:"whitelist"`
	SlashFraction     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=slash_fraction,json=slashFraction,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction" yaml:"slash_fraction"`
	SlashWindow       uint64                                 `protobuf:"varint,6,opt,name=slash_window,json=slashWindow,proto3" json:"slash_window,omitempty" yaml:"slash_window"`
	MinValidPerWindow github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=min_valid_per_window,json=minValidPerWindow,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_valid_per_window" yaml:"min_valid_per_window"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc470b50b143d488, []int{0}
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

func (m *Params) GetVotePeriod() uint64 {
	if m != nil {
		return m.VotePeriod
	}
	return 0
}

func (m *Params) GetWhitelist() DenomList {
	if m != nil {
		return m.Whitelist
	}
	return nil
}

func (m *Params) GetSlashWindow() uint64 {
	if m != nil {
		return m.SlashWindow
	}
	return 0
}

type Denom struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
}

func (m *Denom) Reset()      { *m = Denom{} }
func (*Denom) ProtoMessage() {}
func (*Denom) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc470b50b143d488, []int{1}
}
func (m *Denom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Denom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Denom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Denom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Denom.Merge(m, src)
}
func (m *Denom) XXX_Size() int {
	return m.Size()
}
func (m *Denom) XXX_DiscardUnknown() {
	xxx_messageInfo_Denom.DiscardUnknown(m)
}

var xxx_messageInfo_Denom proto.InternalMessageInfo

type AggregateExchangeRatePrevote struct {
	Hash        string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty" yaml:"hash"`
	Voter       string `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty" yaml:"voter"`
	SubmitBlock uint64 `protobuf:"varint,3,opt,name=submit_block,json=submitBlock,proto3" json:"submit_block,omitempty" yaml:"submit_block"`
}

func (m *AggregateExchangeRatePrevote) Reset()      { *m = AggregateExchangeRatePrevote{} }
func (*AggregateExchangeRatePrevote) ProtoMessage() {}
func (*AggregateExchangeRatePrevote) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc470b50b143d488, []int{2}
}
func (m *AggregateExchangeRatePrevote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AggregateExchangeRatePrevote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AggregateExchangeRatePrevote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AggregateExchangeRatePrevote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateExchangeRatePrevote.Merge(m, src)
}
func (m *AggregateExchangeRatePrevote) XXX_Size() int {
	return m.Size()
}
func (m *AggregateExchangeRatePrevote) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateExchangeRatePrevote.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateExchangeRatePrevote proto.InternalMessageInfo

type AggregateExchangeRateVote struct {
	ExchangeRateTuples ExchangeRateTuples `protobuf:"bytes,1,rep,name=exchange_rate_tuples,json=exchangeRateTuples,proto3,castrepeated=ExchangeRateTuples" json:"exchange_rate_tuples" yaml:"exchange_rate_tuples"`
	Voter              string             `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty" yaml:"voter"`
}

func (m *AggregateExchangeRateVote) Reset()      { *m = AggregateExchangeRateVote{} }
func (*AggregateExchangeRateVote) ProtoMessage() {}
func (*AggregateExchangeRateVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc470b50b143d488, []int{3}
}
func (m *AggregateExchangeRateVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AggregateExchangeRateVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AggregateExchangeRateVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AggregateExchangeRateVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateExchangeRateVote.Merge(m, src)
}
func (m *AggregateExchangeRateVote) XXX_Size() int {
	return m.Size()
}
func (m *AggregateExchangeRateVote) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateExchangeRateVote.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateExchangeRateVote proto.InternalMessageInfo

type ExchangeRateTuple struct {
	Denom        string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	ExchangeRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=exchange_rate,json=exchangeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"exchange_rate" yaml:"exchange_rate"`
}

func (m *ExchangeRateTuple) Reset()      { *m = ExchangeRateTuple{} }
func (*ExchangeRateTuple) ProtoMessage() {}
func (*ExchangeRateTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc470b50b143d488, []int{4}
}
func (m *ExchangeRateTuple) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExchangeRateTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExchangeRateTuple.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExchangeRateTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeRateTuple.Merge(m, src)
}
func (m *ExchangeRateTuple) XXX_Size() int {
	return m.Size()
}
func (m *ExchangeRateTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeRateTuple.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeRateTuple proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "seiprotocol.seichain.oracle.Params")
	proto.RegisterType((*Denom)(nil), "seiprotocol.seichain.oracle.Denom")
	proto.RegisterType((*AggregateExchangeRatePrevote)(nil), "seiprotocol.seichain.oracle.AggregateExchangeRatePrevote")
	proto.RegisterType((*AggregateExchangeRateVote)(nil), "seiprotocol.seichain.oracle.AggregateExchangeRateVote")
	proto.RegisterType((*ExchangeRateTuple)(nil), "seiprotocol.seichain.oracle.ExchangeRateTuple")
}

func init() { proto.RegisterFile("oracle/oracle.proto", fileDescriptor_dc470b50b143d488) }

var fileDescriptor_dc470b50b143d488 = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0x8d, 0x69, 0x53, 0xc8, 0x25, 0x85, 0xd6, 0x0d, 0x10, 0x5a, 0x14, 0x57, 0x87, 0xa8, 0xba,
	0xd4, 0xa6, 0x30, 0x20, 0xb2, 0x61, 0x95, 0x22, 0x21, 0x90, 0xa2, 0x53, 0x55, 0x24, 0x96, 0xe8,
	0x62, 0x1f, 0xf6, 0xa9, 0xb6, 0x2f, 0xba, 0x73, 0x9b, 0x76, 0x61, 0x66, 0x64, 0x44, 0x62, 0xe9,
	0xcc, 0x0e, 0xfc, 0x85, 0x8e, 0x1d, 0x11, 0x83, 0x41, 0xed, 0xc2, 0x86, 0x94, 0x5f, 0x80, 0xee,
	0xce, 0x6d, 0x9d, 0x26, 0xaa, 0xa8, 0x98, 0xec, 0xf7, 0xbd, 0xbb, 0xf7, 0x3d, 0x7f, 0xdf, 0x4b,
	0xc0, 0x1c, 0xe3, 0xd8, 0x8b, 0x88, 0xa3, 0x1f, 0x76, 0x8f, 0xb3, 0x94, 0x99, 0x0b, 0x82, 0x50,
	0xf5, 0xe6, 0xb1, 0xc8, 0x16, 0x84, 0x7a, 0x21, 0xa6, 0x89, 0xad, 0x8f, 0xcc, 0xd7, 0x03, 0x16,
	0x30, 0xc5, 0x3a, 0xf2, 0x4d, 0x5f, 0x99, 0x6f, 0x7a, 0x4c, 0xc4, 0x4c, 0x38, 0x5d, 0x2c, 0x88,
	0xb3, 0xb3, 0xda, 0x25, 0x29, 0x5e, 0x75, 0x3c, 0x46, 0x13, 0xcd, 0xc3, 0x6f, 0x65, 0x30, 0xd5,
	0xc6, 0x1c, 0xc7, 0xc2, 0x7c, 0x0c, 0xaa, 0x3b, 0x2c, 0x25, 0x9d, 0x1e, 0xe1, 0x94, 0xf9, 0x0d,
	0x63, 0xd1, 0x58, 0x9e, 0x74, 0x6f, 0x0d, 0x32, 0xcb, 0xdc, 0xc3, 0x71, 0xd4, 0x82, 0x05, 0x12,
	0x22, 0x20, 0x51, 0x5b, 0x01, 0x33, 0x01, 0xd7, 0x15, 0x97, 0x86, 0x9c, 0x88, 0x90, 0x45, 0x7e,
	0xe3, 0xca, 0xa2, 0xb1, 0x5c, 0x71, 0x9f, 0x1f, 0x64, 0x56, 0xe9, 0x47, 0x66, 0x2d, 0x05, 0x34,
	0x0d, 0xb7, 0xbb, 0xb6, 0xc7, 0x62, 0x27, 0xb7, 0xa3, 0x1f, 0x2b, 0xc2, 0xdf, 0x72, 0xd2, 0xbd,
	0x1e, 0x11, 0xf6, 0x1a, 0xf1, 0x06, 0x99, 0x75, 0xb3, 0xd0, 0xe9, 0x54, 0x0d, 0xa2, 0x69, 0x59,
	0xd8, 0x38, 0xc1, 0x26, 0x01, 0x55, 0x4e, 0xfa, 0x98, 0xfb, 0x9d, 0x2e, 0x4e, 0xfc, 0xc6, 0x84,
	0x6a, 0xb6, 0x76, 0xe9, 0x66, 0xf9, 0x67, 0x15, 0xa4, 0x20, 0x02, 0x1a, 0xb9, 0x38, 0xf1, 0xcd,
	0x00, 0x54, 0xfa, 0x21, 0x4d, 0x49, 0x44, 0x45, 0xda, 0x98, 0x5c, 0x9c, 0x58, 0xae, 0x3e, 0x84,
	0xf6, 0x05, 0x1b, 0xb0, 0xd7, 0x48, 0xc2, 0x62, 0xf7, 0xbe, 0x34, 0x32, 0xc8, 0xac, 0x19, 0x2d,
	0x7f, 0x2a, 0x01, 0x3f, 0xff, 0xb4, 0x2a, 0xea, 0xc8, 0x4b, 0x2a, 0x52, 0x74, 0xa6, 0x2d, 0xe7,
	0x27, 0x22, 0x2c, 0xc2, 0xce, 0x5b, 0x8e, 0xbd, 0x94, 0xb2, 0xa4, 0x51, 0xfe, 0xbf, 0xf9, 0x0d,
	0xab, 0x41, 0x34, 0xad, 0x0a, 0xeb, 0x39, 0x36, 0x5b, 0xa0, 0xa6, 0x4f, 0xf4, 0x69, 0xe2, 0xb3,
	0x7e, 0x63, 0x4a, 0x6d, 0xfa, 0xf6, 0x20, 0xb3, 0xe6, 0x8a, 0xf7, 0x35, 0x0b, 0x51, 0x55, 0xc1,
	0xd7, 0x0a, 0x99, 0xef, 0x40, 0x3d, 0xa6, 0x49, 0x67, 0x07, 0x47, 0xd4, 0x97, 0x61, 0x38, 0xd1,
	0xb8, 0xaa, 0x1c, 0xbf, 0xba, 0xb4, 0xe3, 0x05, 0xdd, 0x71, 0x9c, 0x26, 0x44, 0xb3, 0x31, 0x4d,
	0x36, 0x65, 0xb5, 0x4d, 0xb8, 0xee, 0xdf, 0xba, 0xf6, 0x71, 0xdf, 0x2a, 0xfd, 0xde, 0xb7, 0x0c,
	0xd8, 0x02, 0x65, 0x35, 0x4d, 0xf3, 0x1e, 0x98, 0x4c, 0x70, 0x4c, 0x54, 0x60, 0x2b, 0xee, 0x8d,
	0x41, 0x66, 0x55, 0xb5, 0xa8, 0xac, 0x42, 0xa4, 0xc8, 0x56, 0xed, 0xfd, 0xbe, 0x55, 0xca, 0xef,
	0x96, 0xe0, 0x17, 0x03, 0xdc, 0x7d, 0x1a, 0x04, 0x9c, 0x04, 0x38, 0x25, 0xcf, 0x76, 0xbd, 0x10,
	0x27, 0x01, 0x41, 0x38, 0x25, 0x6d, 0x4e, 0x64, 0xd6, 0xa4, 0x66, 0x88, 0x45, 0x38, 0xaa, 0x29,
	0xab, 0x10, 0x29, 0xd2, 0x5c, 0x02, 0x65, 0x79, 0x98, 0xe7, 0x71, 0x9f, 0x19, 0x64, 0x56, 0xed,
	0x2c, 0xc0, 0x1c, 0x22, 0x4d, 0xab, 0x79, 0x6f, 0x77, 0x63, 0x9a, 0x76, 0xba, 0x11, 0xf3, 0xb6,
	0x54, 0x60, 0x87, 0xe7, 0x5d, 0x60, 0xe5, 0xbc, 0x15, 0x74, 0x25, 0x3a, 0xe7, 0xfb, 0x8f, 0x01,
	0xee, 0x8c, 0xf5, 0xbd, 0x29, 0x4d, 0x7f, 0x32, 0x40, 0x9d, 0xe4, 0xc5, 0x0e, 0xc7, 0xf2, 0x37,
	0xb4, 0xdd, 0x8b, 0x88, 0x68, 0x18, 0x2a, 0xbc, 0xf6, 0x85, 0xe1, 0x2d, 0xaa, 0x6d, 0xc8, 0x6b,
	0xee, 0x93, 0x3c, 0xc8, 0xf9, 0x8a, 0xc6, 0x29, 0xcb, 0x4c, 0x9b, 0x23, 0x37, 0x05, 0x32, 0xc9,
	0x48, 0xed, 0x5f, 0xa7, 0x75, 0xee, 0x8b, 0xbf, 0x1a, 0x60, 0x76, 0xa4, 0x81, 0xd4, 0xf2, 0xe5,
	0xee, 0xf3, 0xfd, 0x14, 0xb4, 0x54, 0x19, 0x22, 0x4d, 0x9b, 0x5b, 0x60, 0x7a, 0xc8, 0x76, 0xde,
	0x7b, 0xfd, 0xd2, 0x31, 0xad, 0x8f, 0x99, 0x01, 0x44, 0xb5, 0xe2, 0x67, 0x0e, 0x1b, 0x77, 0x5f,
	0x1c, 0x1c, 0x35, 0x8d, 0xc3, 0xa3, 0xa6, 0xf1, 0xeb, 0xa8, 0x69, 0x7c, 0x38, 0x6e, 0x96, 0x0e,
	0x8f, 0x9b, 0xa5, 0xef, 0xc7, 0xcd, 0xd2, 0x9b, 0x07, 0x85, 0xae, 0x82, 0xd0, 0x95, 0x93, 0x95,
	0x28, 0xa0, 0x76, 0xe2, 0xec, 0xe6, 0xff, 0xfb, 0xda, 0x43, 0x77, 0x4a, 0x1d, 0x79, 0xf4, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xc9, 0x7e, 0x83, 0xc5, 0x15, 0x06, 0x00, 0x00,
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
	if this.VotePeriod != that1.VotePeriod {
		return false
	}
	if !this.VoteThreshold.Equal(that1.VoteThreshold) {
		return false
	}
	if !this.RewardBand.Equal(that1.RewardBand) {
		return false
	}
	if len(this.Whitelist) != len(that1.Whitelist) {
		return false
	}
	for i := range this.Whitelist {
		if !this.Whitelist[i].Equal(&that1.Whitelist[i]) {
			return false
		}
	}
	if !this.SlashFraction.Equal(that1.SlashFraction) {
		return false
	}
	if this.SlashWindow != that1.SlashWindow {
		return false
	}
	if !this.MinValidPerWindow.Equal(that1.MinValidPerWindow) {
		return false
	}
	return true
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
		size := m.MinValidPerWindow.Size()
		i -= size
		if _, err := m.MinValidPerWindow.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.SlashWindow != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.SlashWindow))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.SlashFraction.Size()
		i -= size
		if _, err := m.SlashFraction.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Whitelist) > 0 {
		for iNdEx := len(m.Whitelist) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Whitelist[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOracle(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.RewardBand.Size()
		i -= size
		if _, err := m.RewardBand.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.VoteThreshold.Size()
		i -= size
		if _, err := m.VoteThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.VotePeriod != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.VotePeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Denom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Denom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Denom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AggregateExchangeRatePrevote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregateExchangeRatePrevote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggregateExchangeRatePrevote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SubmitBlock != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.SubmitBlock))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AggregateExchangeRateVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregateExchangeRateVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggregateExchangeRateVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ExchangeRateTuples) > 0 {
		for iNdEx := len(m.ExchangeRateTuples) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExchangeRateTuples[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOracle(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ExchangeRateTuple) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExchangeRateTuple) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExchangeRateTuple) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ExchangeRate.Size()
		i -= size
		if _, err := m.ExchangeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracle(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VotePeriod != 0 {
		n += 1 + sovOracle(uint64(m.VotePeriod))
	}
	l = m.VoteThreshold.Size()
	n += 1 + l + sovOracle(uint64(l))
	l = m.RewardBand.Size()
	n += 1 + l + sovOracle(uint64(l))
	if len(m.Whitelist) > 0 {
		for _, e := range m.Whitelist {
			l = e.Size()
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	l = m.SlashFraction.Size()
	n += 1 + l + sovOracle(uint64(l))
	if m.SlashWindow != 0 {
		n += 1 + sovOracle(uint64(m.SlashWindow))
	}
	l = m.MinValidPerWindow.Size()
	n += 1 + l + sovOracle(uint64(l))
	return n
}

func (m *Denom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func (m *AggregateExchangeRatePrevote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.SubmitBlock != 0 {
		n += 1 + sovOracle(uint64(m.SubmitBlock))
	}
	return n
}

func (m *AggregateExchangeRateVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ExchangeRateTuples) > 0 {
		for _, e := range m.ExchangeRateTuples {
			l = e.Size()
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func (m *ExchangeRateTuple) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = m.ExchangeRate.Size()
	n += 1 + l + sovOracle(uint64(l))
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
				return fmt.Errorf("proto: wrong wireType = %d for field VotePeriod", wireType)
			}
			m.VotePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VoteThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardBand", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardBand.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Whitelist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Whitelist = append(m.Whitelist, Denom{})
			if err := m.Whitelist[len(m.Whitelist)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFraction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFraction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashWindow", wireType)
			}
			m.SlashWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SlashWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinValidPerWindow", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinValidPerWindow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *Denom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Denom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Denom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *AggregateExchangeRatePrevote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: AggregateExchangeRatePrevote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregateExchangeRatePrevote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitBlock", wireType)
			}
			m.SubmitBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubmitBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *AggregateExchangeRateVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: AggregateExchangeRateVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregateExchangeRateVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRateTuples", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExchangeRateTuples = append(m.ExchangeRateTuples, ExchangeRateTuple{})
			if err := m.ExchangeRateTuples[len(m.ExchangeRateTuples)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *ExchangeRateTuple) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: ExchangeRateTuple: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExchangeRateTuple: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExchangeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)
