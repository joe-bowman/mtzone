// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: microtick/v1beta1/msg/QueryOrderBook.proto

package msg

import (
	fmt "fmt"
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

type QueryOrderBookRequest struct {
	Market   string `protobuf:"bytes,1,opt,name=market,proto3" json:"market"`
	Duration string `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration"`
	Offset   uint32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset"`
	Limit    uint32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit"`
}

func (m *QueryOrderBookRequest) Reset()         { *m = QueryOrderBookRequest{} }
func (m *QueryOrderBookRequest) String() string { return proto.CompactTextString(m) }
func (*QueryOrderBookRequest) ProtoMessage()    {}
func (*QueryOrderBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e15ffeece9131b, []int{0}
}
func (m *QueryOrderBookRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOrderBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOrderBookRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOrderBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOrderBookRequest.Merge(m, src)
}
func (m *QueryOrderBookRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryOrderBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOrderBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOrderBookRequest proto.InternalMessageInfo

func (m *QueryOrderBookRequest) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *QueryOrderBookRequest) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *QueryOrderBookRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *QueryOrderBookRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type OrderBookQuote struct {
	Id       uint32        `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Premium  types.DecCoin `protobuf:"bytes,2,opt,name=premium,proto3" json:"premium"`
	Quantity types.DecCoin `protobuf:"bytes,3,opt,name=quantity,proto3" json:"quantity"`
}

func (m *OrderBookQuote) Reset()         { *m = OrderBookQuote{} }
func (m *OrderBookQuote) String() string { return proto.CompactTextString(m) }
func (*OrderBookQuote) ProtoMessage()    {}
func (*OrderBookQuote) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e15ffeece9131b, []int{1}
}
func (m *OrderBookQuote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderBookQuote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderBookQuote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderBookQuote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBookQuote.Merge(m, src)
}
func (m *OrderBookQuote) XXX_Size() int {
	return m.Size()
}
func (m *OrderBookQuote) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBookQuote.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBookQuote proto.InternalMessageInfo

func (m *OrderBookQuote) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderBookQuote) GetPremium() types.DecCoin {
	if m != nil {
		return m.Premium
	}
	return types.DecCoin{}
}

func (m *OrderBookQuote) GetQuantity() types.DecCoin {
	if m != nil {
		return m.Quantity
	}
	return types.DecCoin{}
}

type QueryOrderBookResponse struct {
	Consensus  types.DecCoin     `protobuf:"bytes,1,opt,name=consensus,proto3" json:"consensus"`
	SumBacking types.DecCoin     `protobuf:"bytes,2,opt,name=sum_backing,json=sumBacking,proto3" json:"sum_backing"`
	SumWeight  types.DecCoin     `protobuf:"bytes,3,opt,name=sum_weight,json=sumWeight,proto3" json:"sum_weight"`
	Offset     uint32            `protobuf:"varint,4,opt,name=offset,proto3" json:"offset"`
	Limit      uint32            `protobuf:"varint,5,opt,name=limit,proto3" json:"limit"`
	Total      uint32            `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
	CallAsks   []*OrderBookQuote `protobuf:"bytes,7,rep,name=call_asks,json=callAsks,proto3" json:"call_asks"`
	CallBids   []*OrderBookQuote `protobuf:"bytes,8,rep,name=call_bids,json=callBids,proto3" json:"call_bids"`
	PutAsks    []*OrderBookQuote `protobuf:"bytes,9,rep,name=put_asks,json=putAsks,proto3" json:"put_asks"`
	PutBids    []*OrderBookQuote `protobuf:"bytes,10,rep,name=put_bids,json=putBids,proto3" json:"put_bids"`
}

func (m *QueryOrderBookResponse) Reset()         { *m = QueryOrderBookResponse{} }
func (m *QueryOrderBookResponse) String() string { return proto.CompactTextString(m) }
func (*QueryOrderBookResponse) ProtoMessage()    {}
func (*QueryOrderBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30e15ffeece9131b, []int{2}
}
func (m *QueryOrderBookResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOrderBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOrderBookResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOrderBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOrderBookResponse.Merge(m, src)
}
func (m *QueryOrderBookResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryOrderBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOrderBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOrderBookResponse proto.InternalMessageInfo

func (m *QueryOrderBookResponse) GetConsensus() types.DecCoin {
	if m != nil {
		return m.Consensus
	}
	return types.DecCoin{}
}

func (m *QueryOrderBookResponse) GetSumBacking() types.DecCoin {
	if m != nil {
		return m.SumBacking
	}
	return types.DecCoin{}
}

func (m *QueryOrderBookResponse) GetSumWeight() types.DecCoin {
	if m != nil {
		return m.SumWeight
	}
	return types.DecCoin{}
}

func (m *QueryOrderBookResponse) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *QueryOrderBookResponse) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryOrderBookResponse) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *QueryOrderBookResponse) GetCallAsks() []*OrderBookQuote {
	if m != nil {
		return m.CallAsks
	}
	return nil
}

func (m *QueryOrderBookResponse) GetCallBids() []*OrderBookQuote {
	if m != nil {
		return m.CallBids
	}
	return nil
}

func (m *QueryOrderBookResponse) GetPutAsks() []*OrderBookQuote {
	if m != nil {
		return m.PutAsks
	}
	return nil
}

func (m *QueryOrderBookResponse) GetPutBids() []*OrderBookQuote {
	if m != nil {
		return m.PutBids
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryOrderBookRequest)(nil), "microtick.v1beta1.msg.QueryOrderBookRequest")
	proto.RegisterType((*OrderBookQuote)(nil), "microtick.v1beta1.msg.OrderBookQuote")
	proto.RegisterType((*QueryOrderBookResponse)(nil), "microtick.v1beta1.msg.QueryOrderBookResponse")
}

func init() {
	proto.RegisterFile("microtick/v1beta1/msg/QueryOrderBook.proto", fileDescriptor_30e15ffeece9131b)
}

var fileDescriptor_30e15ffeece9131b = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0xd3, 0xe6, 0xe1, 0x09, 0xe1, 0x31, 0xd0, 0xca, 0xaa, 0x90, 0x1d, 0x45, 0x42, 0x8a,
	0x40, 0xb2, 0xd5, 0xf2, 0x05, 0x18, 0x10, 0x12, 0x12, 0x54, 0x99, 0x0d, 0x12, 0x9b, 0xc8, 0x8f,
	0xa9, 0x3b, 0x72, 0xc6, 0x93, 0x7a, 0x66, 0x80, 0xf2, 0x15, 0x7c, 0x07, 0x5f, 0xc0, 0x86, 0x7d,
	0x97, 0x5d, 0xb2, 0xb2, 0x50, 0xb2, 0xcb, 0x57, 0x20, 0xcf, 0x38, 0x4e, 0x8a, 0x2a, 0xa5, 0xed,
	0xea, 0x9e, 0xdc, 0x9c, 0x73, 0xe6, 0xf8, 0xea, 0xce, 0x80, 0xe7, 0x94, 0x44, 0x39, 0x13, 0x24,
	0x4a, 0xbd, 0x2f, 0x87, 0x21, 0x16, 0xc1, 0xa1, 0x47, 0x79, 0xe2, 0x8d, 0x25, 0xce, 0xcf, 0x8f,
	0xf3, 0x18, 0xe7, 0x3e, 0x63, 0xa9, 0x3b, 0xcb, 0x99, 0x60, 0x70, 0xaf, 0xe6, 0xba, 0x15, 0xd7,
	0xa5, 0x3c, 0x39, 0x78, 0x92, 0xb0, 0x84, 0x29, 0x86, 0x57, 0x22, 0x4d, 0x3e, 0xb0, 0x23, 0xc6,
	0x29, 0xe3, 0x5e, 0x18, 0x70, 0x5c, 0x5b, 0x47, 0x8c, 0x64, 0xfa, 0xff, 0xe1, 0x4f, 0x03, 0xec,
	0x5d, 0x3d, 0x05, 0xe1, 0x33, 0x89, 0xb9, 0x80, 0x43, 0xd0, 0xa6, 0x41, 0x9e, 0x62, 0x61, 0x19,
	0x03, 0x63, 0x64, 0xfa, 0x60, 0x59, 0x38, 0x55, 0x07, 0x55, 0x15, 0x8e, 0x40, 0x37, 0x96, 0x79,
	0x20, 0x08, 0xcb, 0xac, 0xa6, 0x62, 0xdd, 0x5b, 0x16, 0x4e, 0xdd, 0x43, 0x35, 0x2a, 0xdd, 0xd8,
	0xc9, 0x09, 0xc7, 0xc2, 0xda, 0x19, 0x18, 0xa3, 0xbe, 0x76, 0xd3, 0x1d, 0x54, 0x55, 0xe8, 0x80,
	0xd6, 0x94, 0x50, 0x22, 0xac, 0x5d, 0x45, 0x31, 0x97, 0x85, 0xa3, 0x1b, 0x48, 0x97, 0xe1, 0x6f,
	0x03, 0xdc, 0xaf, 0x73, 0x8e, 0x25, 0x13, 0x18, 0xee, 0x83, 0x26, 0x89, 0x55, 0xc2, 0xbe, 0xdf,
	0x5e, 0x16, 0x4e, 0x93, 0xc4, 0xa8, 0x49, 0x62, 0xf8, 0x0e, 0x74, 0x66, 0x39, 0xa6, 0x44, 0x52,
	0x15, 0xac, 0x77, 0xf4, 0xd4, 0xd5, 0x93, 0x70, 0xcb, 0x49, 0xd4, 0x83, 0x7b, 0x83, 0xa3, 0xd7,
	0x8c, 0x64, 0xfe, 0x83, 0x8b, 0xc2, 0x69, 0x2c, 0x0b, 0x67, 0x25, 0x42, 0x2b, 0x00, 0xdf, 0x83,
	0xee, 0x99, 0x0c, 0x32, 0x41, 0xc4, 0xb9, 0x8a, 0xbe, 0xcd, 0xe9, 0x61, 0xe5, 0x54, 0xab, 0x50,
	0x8d, 0x86, 0xbf, 0x5a, 0x60, 0xff, 0xff, 0x61, 0xf3, 0x19, 0xcb, 0x38, 0x86, 0x1f, 0x80, 0x19,
	0x95, 0x20, 0xe3, 0x92, 0xab, 0xcf, 0xd9, 0x76, 0xce, 0xa3, 0xea, 0x9c, 0xb5, 0x0c, 0xad, 0x21,
	0x44, 0xa0, 0xc7, 0x25, 0x9d, 0x84, 0x41, 0x94, 0x92, 0x2c, 0xb9, 0xd1, 0x08, 0x1e, 0x57, 0x86,
	0x9b, 0x42, 0x04, 0xb8, 0xa4, 0xbe, 0xc6, 0xf0, 0x18, 0x94, 0xbf, 0x26, 0x5f, 0x31, 0x49, 0x4e,
	0xc5, 0x8d, 0x66, 0x01, 0x2b, 0xcb, 0x0d, 0x1d, 0x32, 0xb9, 0xa4, 0x9f, 0x14, 0xdc, 0xd8, 0x89,
	0xdd, 0xed, 0x3b, 0xd1, 0xba, 0x7e, 0x27, 0x4a, 0x82, 0x60, 0x22, 0x98, 0x5a, 0xed, 0x35, 0x41,
	0x35, 0x90, 0x2e, 0x70, 0x0c, 0xcc, 0x28, 0x98, 0x4e, 0x27, 0x01, 0x4f, 0xb9, 0xd5, 0x19, 0xec,
	0x8c, 0x7a, 0x47, 0xcf, 0xdc, 0x6b, 0xaf, 0x90, 0x7b, 0x75, 0xb7, 0xfc, 0xbe, 0x1a, 0xef, 0x4a,
	0x8b, 0xba, 0x25, 0x7c, 0xc5, 0x53, 0x5e, 0x5b, 0x86, 0x24, 0xe6, 0x56, 0xf7, 0x6e, 0x96, 0xa5,
	0x56, 0x5b, 0xfa, 0x24, 0xe6, 0xf0, 0x23, 0xe8, 0xce, 0xa4, 0xd0, 0x21, 0xcd, 0xdb, 0x38, 0xaa,
	0x0b, 0xb7, 0x92, 0xa2, 0xce, 0x4c, 0x0a, 0x15, 0xb1, 0xf2, 0x53, 0x09, 0xc1, 0x9d, 0xfc, 0x54,
	0xc0, 0xd2, 0xaf, 0xcc, 0xe7, 0xbf, 0xbd, 0x98, 0xdb, 0xc6, 0xe5, 0xdc, 0x36, 0xfe, 0xce, 0x6d,
	0xe3, 0xc7, 0xc2, 0x6e, 0x5c, 0x2e, 0xec, 0xc6, 0x9f, 0x85, 0xdd, 0xf8, 0xfc, 0x22, 0x21, 0xe2,
	0x54, 0x86, 0x6e, 0xc4, 0xa8, 0xb7, 0x7e, 0xc5, 0xa8, 0xf8, 0xce, 0x32, 0xec, 0x7d, 0xdb, 0x6c,
	0xf1, 0x24, 0x6c, 0xab, 0x57, 0xe7, 0xe5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xc7, 0xc4,
	0xd5, 0xf0, 0x04, 0x00, 0x00,
}

func (m *QueryOrderBookRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOrderBookRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOrderBookRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Limit != 0 {
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x20
	}
	if m.Offset != 0 {
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Duration) > 0 {
		i -= len(m.Duration)
		copy(dAtA[i:], m.Duration)
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(len(m.Duration)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Market) > 0 {
		i -= len(m.Market)
		copy(dAtA[i:], m.Market)
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(len(m.Market)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OrderBookQuote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderBookQuote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderBookQuote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Quantity.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Premium.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryOrderBookResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOrderBookResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOrderBookResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PutBids) > 0 {
		for iNdEx := len(m.PutBids) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PutBids[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryOrderBook(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.PutAsks) > 0 {
		for iNdEx := len(m.PutAsks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PutAsks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryOrderBook(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.CallBids) > 0 {
		for iNdEx := len(m.CallBids) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CallBids[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryOrderBook(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CallAsks) > 0 {
		for iNdEx := len(m.CallAsks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CallAsks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryOrderBook(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Total != 0 {
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x30
	}
	if m.Limit != 0 {
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x28
	}
	if m.Offset != 0 {
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(m.Offset))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.SumWeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.SumBacking.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Consensus.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryOrderBook(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQueryOrderBook(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryOrderBook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryOrderBookRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Market)
	if l > 0 {
		n += 1 + l + sovQueryOrderBook(uint64(l))
	}
	l = len(m.Duration)
	if l > 0 {
		n += 1 + l + sovQueryOrderBook(uint64(l))
	}
	if m.Offset != 0 {
		n += 1 + sovQueryOrderBook(uint64(m.Offset))
	}
	if m.Limit != 0 {
		n += 1 + sovQueryOrderBook(uint64(m.Limit))
	}
	return n
}

func (m *OrderBookQuote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQueryOrderBook(uint64(m.Id))
	}
	l = m.Premium.Size()
	n += 1 + l + sovQueryOrderBook(uint64(l))
	l = m.Quantity.Size()
	n += 1 + l + sovQueryOrderBook(uint64(l))
	return n
}

func (m *QueryOrderBookResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Consensus.Size()
	n += 1 + l + sovQueryOrderBook(uint64(l))
	l = m.SumBacking.Size()
	n += 1 + l + sovQueryOrderBook(uint64(l))
	l = m.SumWeight.Size()
	n += 1 + l + sovQueryOrderBook(uint64(l))
	if m.Offset != 0 {
		n += 1 + sovQueryOrderBook(uint64(m.Offset))
	}
	if m.Limit != 0 {
		n += 1 + sovQueryOrderBook(uint64(m.Limit))
	}
	if m.Total != 0 {
		n += 1 + sovQueryOrderBook(uint64(m.Total))
	}
	if len(m.CallAsks) > 0 {
		for _, e := range m.CallAsks {
			l = e.Size()
			n += 1 + l + sovQueryOrderBook(uint64(l))
		}
	}
	if len(m.CallBids) > 0 {
		for _, e := range m.CallBids {
			l = e.Size()
			n += 1 + l + sovQueryOrderBook(uint64(l))
		}
	}
	if len(m.PutAsks) > 0 {
		for _, e := range m.PutAsks {
			l = e.Size()
			n += 1 + l + sovQueryOrderBook(uint64(l))
		}
	}
	if len(m.PutBids) > 0 {
		for _, e := range m.PutBids {
			l = e.Size()
			n += 1 + l + sovQueryOrderBook(uint64(l))
		}
	}
	return n
}

func sovQueryOrderBook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryOrderBook(x uint64) (n int) {
	return sovQueryOrderBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryOrderBookRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryOrderBook
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
			return fmt.Errorf("proto: QueryOrderBookRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOrderBookRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Market", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Market = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Duration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQueryOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryOrderBook
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
func (m *OrderBookQuote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryOrderBook
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
			return fmt.Errorf("proto: OrderBookQuote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderBookQuote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return fmt.Errorf("proto: wrong wireType = %d for field Premium", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Premium.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryOrderBook
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
func (m *QueryOrderBookResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryOrderBook
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
			return fmt.Errorf("proto: QueryOrderBookResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOrderBookResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consensus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Consensus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SumBacking", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SumBacking.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SumWeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SumWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallAsks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallAsks = append(m.CallAsks, &OrderBookQuote{})
			if err := m.CallAsks[len(m.CallAsks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallBids", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallBids = append(m.CallBids, &OrderBookQuote{})
			if err := m.CallBids[len(m.CallBids)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PutAsks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PutAsks = append(m.PutAsks, &OrderBookQuote{})
			if err := m.PutAsks[len(m.PutAsks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PutBids", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryOrderBook
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
				return ErrInvalidLengthQueryOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PutBids = append(m.PutBids, &OrderBookQuote{})
			if err := m.PutBids[len(m.PutBids)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryOrderBook
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
func skipQueryOrderBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryOrderBook
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
					return 0, ErrIntOverflowQueryOrderBook
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
					return 0, ErrIntOverflowQueryOrderBook
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
				return 0, ErrInvalidLengthQueryOrderBook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryOrderBook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryOrderBook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryOrderBook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryOrderBook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryOrderBook = fmt.Errorf("proto: unexpected end of group")
)
