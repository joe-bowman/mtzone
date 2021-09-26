// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: microtick/v1beta1/msg/QueryQuote.proto

package msg

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type QueryQuoteRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (m *QueryQuoteRequest) Reset()         { *m = QueryQuoteRequest{} }
func (m *QueryQuoteRequest) String() string { return proto.CompactTextString(m) }
func (*QueryQuoteRequest) ProtoMessage()    {}
func (*QueryQuoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0852982b6936a37f, []int{0}
}
func (m *QueryQuoteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryQuoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryQuoteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryQuoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryQuoteRequest.Merge(m, src)
}
func (m *QueryQuoteRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryQuoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryQuoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryQuoteRequest proto.InternalMessageInfo

func (m *QueryQuoteRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryQuoteResponse struct {
	Id        uint32                                        `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Provider  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=provider,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"provider"`
	Market    string                                        `protobuf:"bytes,3,opt,name=market,proto3" json:"market"`
	Duration  string                                        `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration"`
	Backing   types.DecCoin                                 `protobuf:"bytes,5,opt,name=backing,proto3" json:"backing"`
	Ask       types.DecCoin                                 `protobuf:"bytes,6,opt,name=ask,proto3" json:"ask"`
	Bid       types.DecCoin                                 `protobuf:"bytes,7,opt,name=bid,proto3" json:"bid"`
	Quantity  types.DecCoin                                 `protobuf:"bytes,8,opt,name=quantity,proto3" json:"quantity"`
	Consensus types.DecCoin                                 `protobuf:"bytes,9,opt,name=consensus,proto3" json:"consensus"`
	Spot      types.DecCoin                                 `protobuf:"bytes,10,opt,name=spot,proto3" json:"spot"`
	CallAsk   types.DecCoin                                 `protobuf:"bytes,11,opt,name=call_ask,json=callAsk,proto3" json:"call_ask"`
	CallBid   types.DecCoin                                 `protobuf:"bytes,12,opt,name=call_bid,json=callBid,proto3" json:"call_bid"`
	PutAsk    types.DecCoin                                 `protobuf:"bytes,13,opt,name=put_ask,json=putAsk,proto3" json:"put_ask"`
	PutBid    types.DecCoin                                 `protobuf:"bytes,14,opt,name=put_bid,json=putBid,proto3" json:"put_bid"`
	Modified  int64                                         `protobuf:"varint,15,opt,name=modified,proto3" json:"modified"`
	CanModify int64                                         `protobuf:"varint,16,opt,name=can_modify,json=canModify,proto3" json:"can_modify"`
}

func (m *QueryQuoteResponse) Reset()         { *m = QueryQuoteResponse{} }
func (m *QueryQuoteResponse) String() string { return proto.CompactTextString(m) }
func (*QueryQuoteResponse) ProtoMessage()    {}
func (*QueryQuoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0852982b6936a37f, []int{1}
}
func (m *QueryQuoteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryQuoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryQuoteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryQuoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryQuoteResponse.Merge(m, src)
}
func (m *QueryQuoteResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryQuoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryQuoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryQuoteResponse proto.InternalMessageInfo

func (m *QueryQuoteResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueryQuoteResponse) GetProvider() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *QueryQuoteResponse) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *QueryQuoteResponse) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *QueryQuoteResponse) GetBacking() types.DecCoin {
	if m != nil {
		return m.Backing
	}
	return types.DecCoin{}
}

func (m *QueryQuoteResponse) GetAsk() types.DecCoin {
	if m != nil {
		return m.Ask
	}
	return types.DecCoin{}
}

func (m *QueryQuoteResponse) GetBid() types.DecCoin {
	if m != nil {
		return m.Bid
	}
	return types.DecCoin{}
}

func (m *QueryQuoteResponse) GetQuantity() types.DecCoin {
	if m != nil {
		return m.Quantity
	}
	return types.DecCoin{}
}

func (m *QueryQuoteResponse) GetConsensus() types.DecCoin {
	if m != nil {
		return m.Consensus
	}
	return types.DecCoin{}
}

func (m *QueryQuoteResponse) GetSpot() types.DecCoin {
	if m != nil {
		return m.Spot
	}
	return types.DecCoin{}
}

func (m *QueryQuoteResponse) GetCallAsk() types.DecCoin {
	if m != nil {
		return m.CallAsk
	}
	return types.DecCoin{}
}

func (m *QueryQuoteResponse) GetCallBid() types.DecCoin {
	if m != nil {
		return m.CallBid
	}
	return types.DecCoin{}
}

func (m *QueryQuoteResponse) GetPutAsk() types.DecCoin {
	if m != nil {
		return m.PutAsk
	}
	return types.DecCoin{}
}

func (m *QueryQuoteResponse) GetPutBid() types.DecCoin {
	if m != nil {
		return m.PutBid
	}
	return types.DecCoin{}
}

func (m *QueryQuoteResponse) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

func (m *QueryQuoteResponse) GetCanModify() int64 {
	if m != nil {
		return m.CanModify
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryQuoteRequest)(nil), "microtick.v1beta1.msg.QueryQuoteRequest")
	proto.RegisterType((*QueryQuoteResponse)(nil), "microtick.v1beta1.msg.QueryQuoteResponse")
}

func init() {
	proto.RegisterFile("microtick/v1beta1/msg/QueryQuote.proto", fileDescriptor_0852982b6936a37f)
}

var fileDescriptor_0852982b6936a37f = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xbf, 0x8f, 0xda, 0x3e,
	0x18, 0xc6, 0x09, 0xdc, 0x17, 0x82, 0xe1, 0x7e, 0x59, 0xdf, 0x56, 0xd6, 0xa9, 0x4a, 0x10, 0x43,
	0x15, 0xe9, 0x44, 0xa2, 0x6b, 0xc7, 0xaa, 0x95, 0x48, 0x7f, 0xa9, 0x95, 0x6e, 0x38, 0x6f, 0xed,
	0x72, 0x4a, 0x6c, 0x37, 0xb5, 0x72, 0x89, 0x73, 0xb1, 0x73, 0x2a, 0xfd, 0x2b, 0xfa, 0x67, 0xdd,
	0x78, 0x63, 0xa7, 0xa8, 0x82, 0x8d, 0xb1, 0x63, 0xa7, 0xca, 0x21, 0x04, 0x96, 0x4a, 0xb0, 0x38,
	0x0f, 0x2f, 0xcf, 0xf3, 0xf1, 0xcb, 0x8b, 0x63, 0xf0, 0x34, 0xe1, 0x24, 0x17, 0x8a, 0x93, 0xd8,
	0xbb, 0xbb, 0x08, 0x99, 0x0a, 0x2e, 0xbc, 0x44, 0x46, 0xde, 0x55, 0xc1, 0xf2, 0xd9, 0x55, 0x21,
	0x14, 0x73, 0xb3, 0x5c, 0x28, 0x01, 0x1f, 0x35, 0x3e, 0xb7, 0xf6, 0xb9, 0x89, 0x8c, 0xce, 0xfe,
	0x8f, 0x44, 0x24, 0x2a, 0x87, 0xa7, 0xd5, 0xca, 0x7c, 0x66, 0x11, 0x21, 0x13, 0x21, 0xbd, 0x30,
	0x90, 0xac, 0xc1, 0x12, 0xc1, 0xd3, 0xd5, 0xf7, 0xe3, 0x73, 0x70, 0xba, 0xd9, 0x00, 0xb3, 0xdb,
	0x82, 0x49, 0x05, 0x1f, 0x83, 0x36, 0xa7, 0xc8, 0x18, 0x19, 0xce, 0xa1, 0xdf, 0x5d, 0x96, 0x76,
	0x9b, 0x53, 0xdc, 0xe6, 0x74, 0xfc, 0xbb, 0x07, 0xe0, 0xb6, 0x5b, 0x66, 0x22, 0x95, 0xec, 0x5f,
	0x76, 0xf8, 0x09, 0x98, 0x59, 0x2e, 0xee, 0x38, 0x65, 0x39, 0x6a, 0x8f, 0x0c, 0x67, 0xe8, 0xbf,
	0x5c, 0x96, 0x76, 0x53, 0xfb, 0x53, 0xda, 0x93, 0x88, 0xab, 0xaf, 0x45, 0xe8, 0x12, 0x91, 0x78,
	0x75, 0xa3, 0xab, 0xc7, 0x44, 0xd2, 0xd8, 0x53, 0xb3, 0x8c, 0x49, 0x77, 0x4a, 0xc8, 0x94, 0xd2,
	0x9c, 0x49, 0x89, 0x9b, 0x28, 0x1c, 0x83, 0x6e, 0x12, 0xe4, 0x31, 0x53, 0xa8, 0x33, 0x32, 0x9c,
	0xbe, 0x0f, 0x96, 0xa5, 0x5d, 0x57, 0x70, 0xfd, 0x84, 0x0e, 0x30, 0x69, 0x91, 0x07, 0x8a, 0x8b,
	0x14, 0x1d, 0x54, 0xae, 0xa1, 0xde, 0x7e, 0x5d, 0xc3, 0x8d, 0x82, 0xef, 0x41, 0x2f, 0x0c, 0x48,
	0xcc, 0xd3, 0x08, 0xfd, 0x37, 0x32, 0x9c, 0xc1, 0xb3, 0x27, 0xee, 0xaa, 0x0d, 0x57, 0x8f, 0xad,
	0x99, 0xf2, 0x1b, 0x46, 0x5e, 0x0b, 0x9e, 0xfa, 0xc7, 0xf7, 0xa5, 0xdd, 0x5a, 0x96, 0xf6, 0x3a,
	0x84, 0xd7, 0x02, 0xbe, 0x00, 0x9d, 0x40, 0xc6, 0xa8, 0xbb, 0x03, 0x64, 0x50, 0x43, 0x74, 0x00,
	0xeb, 0x45, 0x87, 0x43, 0x4e, 0x51, 0x6f, 0x9f, 0x70, 0xc8, 0x29, 0xd6, 0x0b, 0xfc, 0x08, 0xcc,
	0xdb, 0x22, 0x48, 0x15, 0x57, 0x33, 0x64, 0xee, 0x40, 0x38, 0xa9, 0x09, 0x4d, 0x0a, 0x37, 0x0a,
	0x5e, 0x82, 0x3e, 0xd1, 0x7f, 0x6c, 0x2a, 0x0b, 0x89, 0xfa, 0x3b, 0xc0, 0x4e, 0x6b, 0xd8, 0x26,
	0x86, 0x37, 0x12, 0xbe, 0x02, 0x07, 0x32, 0x13, 0x0a, 0x81, 0x1d, 0x48, 0xc3, 0x9a, 0x54, 0x25,
	0x70, 0xb5, 0xc2, 0x0f, 0xc0, 0x24, 0xc1, 0xcd, 0xcd, 0xb5, 0x9e, 0xec, 0x60, 0x9f, 0x9f, 0xb6,
	0x4e, 0xe1, 0x9e, 0x56, 0x53, 0x19, 0x37, 0x28, 0x3d, 0xe7, 0xe1, 0xde, 0x28, 0x3d, 0xec, 0x0a,
	0xe5, 0x73, 0x0a, 0xdf, 0x81, 0x5e, 0x56, 0xa8, 0xaa, 0xa9, 0xc3, 0x7d, 0xce, 0x4c, 0x1d, 0xc2,
	0xdd, 0xac, 0x50, 0xba, 0xa5, 0x9a, 0xa3, 0x3b, 0x3a, 0xda, 0x97, 0xa3, 0x1b, 0xd2, 0x1c, 0xdd,
	0x8f, 0x03, 0xcc, 0x44, 0x50, 0xfe, 0x85, 0x33, 0x8a, 0x8e, 0x47, 0x86, 0xd3, 0x59, 0x9d, 0xf6,
	0x75, 0x0d, 0x37, 0x0a, 0x4e, 0x00, 0x20, 0x41, 0x7a, 0x5d, 0x7d, 0x9e, 0xa1, 0x93, 0xca, 0x7b,
	0xb4, 0x2c, 0xed, 0xad, 0x2a, 0xee, 0x93, 0x20, 0xbd, 0xac, 0xa4, 0xff, 0xf6, 0x7e, 0x6e, 0x19,
	0x0f, 0x73, 0xcb, 0xf8, 0x35, 0xb7, 0x8c, 0x1f, 0x0b, 0xab, 0xf5, 0xb0, 0xb0, 0x5a, 0x3f, 0x17,
	0x56, 0xeb, 0xf3, 0xf9, 0xd6, 0xdb, 0xbb, 0xb9, 0xbb, 0x12, 0xf5, 0x5d, 0xa4, 0xcc, 0xfb, 0xb6,
	0x5d, 0x92, 0x51, 0xd8, 0xad, 0xee, 0x9b, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x11, 0x12,
	0xbc, 0x2d, 0xe6, 0x04, 0x00, 0x00,
}

func (m *QueryQuoteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryQuoteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryQuoteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQueryQuote(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryQuoteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryQuoteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryQuoteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CanModify != 0 {
		i = encodeVarintQueryQuote(dAtA, i, uint64(m.CanModify))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.Modified != 0 {
		i = encodeVarintQueryQuote(dAtA, i, uint64(m.Modified))
		i--
		dAtA[i] = 0x78
	}
	{
		size, err := m.PutBid.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryQuote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size, err := m.PutAsk.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryQuote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size, err := m.CallBid.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryQuote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size, err := m.CallAsk.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryQuote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size, err := m.Spot.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryQuote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size, err := m.Consensus.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryQuote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size, err := m.Quantity.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryQuote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size, err := m.Bid.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryQuote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.Ask.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryQuote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Backing.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryQuote(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Duration) > 0 {
		i -= len(m.Duration)
		copy(dAtA[i:], m.Duration)
		i = encodeVarintQueryQuote(dAtA, i, uint64(len(m.Duration)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Market) > 0 {
		i -= len(m.Market)
		copy(dAtA[i:], m.Market)
		i = encodeVarintQueryQuote(dAtA, i, uint64(len(m.Market)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintQueryQuote(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintQueryQuote(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryQuote(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryQuote(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryQuoteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQueryQuote(uint64(m.Id))
	}
	return n
}

func (m *QueryQuoteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQueryQuote(uint64(m.Id))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovQueryQuote(uint64(l))
	}
	l = len(m.Market)
	if l > 0 {
		n += 1 + l + sovQueryQuote(uint64(l))
	}
	l = len(m.Duration)
	if l > 0 {
		n += 1 + l + sovQueryQuote(uint64(l))
	}
	l = m.Backing.Size()
	n += 1 + l + sovQueryQuote(uint64(l))
	l = m.Ask.Size()
	n += 1 + l + sovQueryQuote(uint64(l))
	l = m.Bid.Size()
	n += 1 + l + sovQueryQuote(uint64(l))
	l = m.Quantity.Size()
	n += 1 + l + sovQueryQuote(uint64(l))
	l = m.Consensus.Size()
	n += 1 + l + sovQueryQuote(uint64(l))
	l = m.Spot.Size()
	n += 1 + l + sovQueryQuote(uint64(l))
	l = m.CallAsk.Size()
	n += 1 + l + sovQueryQuote(uint64(l))
	l = m.CallBid.Size()
	n += 1 + l + sovQueryQuote(uint64(l))
	l = m.PutAsk.Size()
	n += 1 + l + sovQueryQuote(uint64(l))
	l = m.PutBid.Size()
	n += 1 + l + sovQueryQuote(uint64(l))
	if m.Modified != 0 {
		n += 1 + sovQueryQuote(uint64(m.Modified))
	}
	if m.CanModify != 0 {
		n += 2 + sovQueryQuote(uint64(m.CanModify))
	}
	return n
}

func sovQueryQuote(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryQuote(x uint64) (n int) {
	return sovQueryQuote(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryQuoteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryQuote
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
			return fmt.Errorf("proto: QueryQuoteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryQuoteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQueryQuote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryQuote
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
func (m *QueryQuoteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryQuote
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
			return fmt.Errorf("proto: QueryQuoteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryQuoteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = append(m.Provider[:0], dAtA[iNdEx:postIndex]...)
			if m.Provider == nil {
				m.Provider = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Market", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Market = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Duration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Backing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Backing.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ask", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ask.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consensus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Consensus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallAsk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CallAsk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallBid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CallBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PutAsk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PutAsk.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PutBid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
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
				return ErrInvalidLengthQueryQuote
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryQuote
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PutBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Modified", wireType)
			}
			m.Modified = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Modified |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanModify", wireType)
			}
			m.CanModify = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryQuote
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CanModify |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQueryQuote(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryQuote
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
func skipQueryQuote(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryQuote
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
					return 0, ErrIntOverflowQueryQuote
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
					return 0, ErrIntOverflowQueryQuote
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
				return 0, ErrInvalidLengthQueryQuote
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryQuote
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryQuote
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryQuote        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryQuote          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryQuote = fmt.Errorf("proto: unexpected end of group")
)
