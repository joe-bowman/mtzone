// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: microtick/v1beta1/msg/TxDeposit.proto

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

type TxDepositQuote struct {
	Id        uint32                                        `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Requester github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=requester,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"requester"`
	Deposit   string                                        `protobuf:"bytes,3,opt,name=deposit,proto3" json:"deposit"`
}

func (m *TxDepositQuote) Reset()         { *m = TxDepositQuote{} }
func (m *TxDepositQuote) String() string { return proto.CompactTextString(m) }
func (*TxDepositQuote) ProtoMessage()    {}
func (*TxDepositQuote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae54d804beae1656, []int{0}
}
func (m *TxDepositQuote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxDepositQuote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxDepositQuote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxDepositQuote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxDepositQuote.Merge(m, src)
}
func (m *TxDepositQuote) XXX_Size() int {
	return m.Size()
}
func (m *TxDepositQuote) XXX_DiscardUnknown() {
	xxx_messageInfo_TxDepositQuote.DiscardUnknown(m)
}

var xxx_messageInfo_TxDepositQuote proto.InternalMessageInfo

func (m *TxDepositQuote) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TxDepositQuote) GetRequester() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Requester
	}
	return nil
}

func (m *TxDepositQuote) GetDeposit() string {
	if m != nil {
		return m.Deposit
	}
	return ""
}

type DepositQuoteData struct {
	Time       int64         `protobuf:"varint,1,opt,name=time,proto3" json:"time"`
	Market     string        `protobuf:"bytes,2,opt,name=market,proto3" json:"market"`
	Duration   string        `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration"`
	Consensus  types.DecCoin `protobuf:"bytes,4,opt,name=consensus,proto3" json:"consensus"`
	Backing    types.DecCoin `protobuf:"bytes,5,opt,name=backing,proto3" json:"backing"`
	Commission types.DecCoin `protobuf:"bytes,6,opt,name=commission,proto3" json:"commission"`
	Reward     types.Coin    `protobuf:"bytes,7,opt,name=reward,proto3" json:"reward"`
	Adjustment string        `protobuf:"bytes,8,opt,name=adjustment,proto3" json:"adjustment"`
}

func (m *DepositQuoteData) Reset()         { *m = DepositQuoteData{} }
func (m *DepositQuoteData) String() string { return proto.CompactTextString(m) }
func (*DepositQuoteData) ProtoMessage()    {}
func (*DepositQuoteData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae54d804beae1656, []int{1}
}
func (m *DepositQuoteData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositQuoteData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositQuoteData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositQuoteData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositQuoteData.Merge(m, src)
}
func (m *DepositQuoteData) XXX_Size() int {
	return m.Size()
}
func (m *DepositQuoteData) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositQuoteData.DiscardUnknown(m)
}

var xxx_messageInfo_DepositQuoteData proto.InternalMessageInfo

func (m *DepositQuoteData) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *DepositQuoteData) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *DepositQuoteData) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *DepositQuoteData) GetConsensus() types.DecCoin {
	if m != nil {
		return m.Consensus
	}
	return types.DecCoin{}
}

func (m *DepositQuoteData) GetBacking() types.DecCoin {
	if m != nil {
		return m.Backing
	}
	return types.DecCoin{}
}

func (m *DepositQuoteData) GetCommission() types.DecCoin {
	if m != nil {
		return m.Commission
	}
	return types.DecCoin{}
}

func (m *DepositQuoteData) GetReward() types.Coin {
	if m != nil {
		return m.Reward
	}
	return types.Coin{}
}

func (m *DepositQuoteData) GetAdjustment() string {
	if m != nil {
		return m.Adjustment
	}
	return ""
}

func init() {
	proto.RegisterType((*TxDepositQuote)(nil), "microtick.v1beta1.msg.TxDepositQuote")
	proto.RegisterType((*DepositQuoteData)(nil), "microtick.v1beta1.msg.DepositQuoteData")
}

func init() {
	proto.RegisterFile("microtick/v1beta1/msg/TxDeposit.proto", fileDescriptor_ae54d804beae1656)
}

var fileDescriptor_ae54d804beae1656 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0x6e, 0xda, 0x35, 0x6d, 0x67, 0xd7, 0xaa, 0x83, 0x4a, 0x5c, 0x96, 0xa4, 0x14, 0x16, 0x02,
	0xb2, 0x09, 0xab, 0x77, 0xa1, 0xb1, 0xe2, 0x49, 0xd0, 0xc1, 0x93, 0x78, 0x49, 0x66, 0x86, 0x38,
	0x96, 0xc9, 0xd4, 0xcc, 0x44, 0x57, 0x7f, 0x85, 0x7f, 0xc4, 0xff, 0xb1, 0xc7, 0x3d, 0x0a, 0x42,
	0x90, 0xf6, 0x96, 0x9f, 0xe0, 0x49, 0x32, 0x99, 0x26, 0x39, 0xec, 0xa1, 0x97, 0xbc, 0x2f, 0x5f,
	0xde, 0xf7, 0xbd, 0x37, 0x2f, 0x6f, 0xc0, 0x39, 0x67, 0x38, 0x17, 0x8a, 0xe1, 0x75, 0xf8, 0xf5,
	0x32, 0xa1, 0x2a, 0xbe, 0x0c, 0xb9, 0x4c, 0xc3, 0xf7, 0x57, 0x2b, 0xba, 0x11, 0x92, 0xa9, 0x60,
	0x93, 0x0b, 0x25, 0xe0, 0xa3, 0x36, 0x2d, 0x30, 0x69, 0x01, 0x97, 0xe9, 0xe9, 0xc3, 0x54, 0xa4,
	0x42, 0x67, 0x84, 0x35, 0x6a, 0x92, 0x4f, 0x5d, 0x2c, 0x24, 0x17, 0x32, 0x4c, 0x62, 0x49, 0x5b,
	0x57, 0x2c, 0x58, 0xd6, 0x7c, 0x5f, 0xfc, 0xb2, 0xc0, 0xac, 0x2d, 0xf0, 0xae, 0x10, 0x8a, 0xc2,
	0xc7, 0x60, 0xc8, 0x88, 0x63, 0xcd, 0x2d, 0xff, 0x6e, 0x64, 0x57, 0xa5, 0x37, 0x64, 0x04, 0x0d,
	0x19, 0x81, 0x1f, 0xc1, 0x34, 0xa7, 0x5f, 0x0a, 0x2a, 0x15, 0xcd, 0x9d, 0xe1, 0xdc, 0xf2, 0x4f,
	0xa2, 0x17, 0x55, 0xe9, 0x75, 0xe4, 0xbf, 0xd2, 0xbb, 0x48, 0x99, 0xfa, 0x54, 0x24, 0x01, 0x16,
	0x3c, 0x34, 0x95, 0x9b, 0x70, 0x21, 0xc9, 0x3a, 0x54, 0xdf, 0x37, 0x54, 0x06, 0x4b, 0x8c, 0x97,
	0x84, 0xe4, 0x54, 0x4a, 0xd4, 0x69, 0xe1, 0x39, 0x18, 0x93, 0xa6, 0x0b, 0x67, 0x34, 0xb7, 0xfc,
	0x69, 0x74, 0x5c, 0x95, 0xde, 0x9e, 0x42, 0x7b, 0xb0, 0xf8, 0x33, 0x02, 0xf7, 0xfb, 0xdd, 0xae,
	0x62, 0x15, 0xc3, 0x33, 0x70, 0xa4, 0x18, 0xa7, 0xba, 0xe7, 0x51, 0x34, 0xa9, 0x4a, 0x4f, 0xbf,
	0x23, 0xfd, 0x84, 0x0b, 0x60, 0xf3, 0x38, 0x5f, 0x53, 0xa5, 0x9b, 0x9e, 0x46, 0xa0, 0x2a, 0x3d,
	0xc3, 0x20, 0x13, 0xa1, 0x0f, 0x26, 0xa4, 0xc8, 0x63, 0xc5, 0x44, 0x66, 0xca, 0x9f, 0x54, 0xa5,
	0xd7, 0x72, 0xa8, 0x45, 0xf0, 0x0d, 0x98, 0x62, 0x91, 0x49, 0x9a, 0xc9, 0x42, 0x3a, 0x47, 0x73,
	0xcb, 0x3f, 0x7e, 0x76, 0x16, 0x34, 0x67, 0x0c, 0xea, 0x21, 0xb7, 0xff, 0x64, 0x45, 0xf1, 0x4b,
	0xc1, 0xb2, 0xe8, 0xc1, 0x75, 0xe9, 0x0d, 0xea, 0x39, 0xb5, 0x32, 0xd4, 0x41, 0xf8, 0x1a, 0x8c,
	0x93, 0x18, 0xaf, 0x59, 0x96, 0x3a, 0x77, 0x0e, 0x30, 0xbb, 0x67, 0xcc, 0xf6, 0x22, 0xb4, 0x07,
	0xf0, 0x2d, 0x00, 0x58, 0x70, 0xce, 0xa4, 0xac, 0xcf, 0x60, 0x1f, 0xe0, 0x05, 0x8d, 0x57, 0x4f,
	0x87, 0x7a, 0x18, 0x2e, 0x81, 0x9d, 0xd3, 0x6f, 0x71, 0x4e, 0x9c, 0xb1, 0x76, 0x7b, 0x72, 0xab,
	0x9b, 0xb6, 0x9a, 0x19, 0x2b, 0x23, 0x40, 0x26, 0xc2, 0x00, 0x80, 0x98, 0x7c, 0x2e, 0xa4, 0xe2,
	0x34, 0x53, 0xce, 0x44, 0x0f, 0x76, 0x56, 0x97, 0xec, 0x58, 0xd4, 0xc3, 0xd1, 0xab, 0xeb, 0xad,
	0x6b, 0xdd, 0x6c, 0x5d, 0xeb, 0xef, 0xd6, 0xb5, 0x7e, 0xee, 0xdc, 0xc1, 0xcd, 0xce, 0x1d, 0xfc,
	0xde, 0xb9, 0x83, 0x0f, 0x4f, 0x7b, 0x8b, 0xd5, 0x5d, 0x13, 0xae, 0x7e, 0x88, 0x8c, 0x86, 0x57,
	0x7d, 0x4a, 0xa6, 0x89, 0xad, 0x77, 0xfb, 0xf9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x3b,
	0x69, 0xfd, 0x51, 0x03, 0x00, 0x00,
}

func (m *TxDepositQuote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxDepositQuote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxDepositQuote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		i -= len(m.Deposit)
		copy(dAtA[i:], m.Deposit)
		i = encodeVarintTxDeposit(dAtA, i, uint64(len(m.Deposit)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Requester) > 0 {
		i -= len(m.Requester)
		copy(dAtA[i:], m.Requester)
		i = encodeVarintTxDeposit(dAtA, i, uint64(len(m.Requester)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTxDeposit(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DepositQuoteData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositQuoteData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositQuoteData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Adjustment) > 0 {
		i -= len(m.Adjustment)
		copy(dAtA[i:], m.Adjustment)
		i = encodeVarintTxDeposit(dAtA, i, uint64(len(m.Adjustment)))
		i--
		dAtA[i] = 0x42
	}
	{
		size, err := m.Reward.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTxDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.Commission.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTxDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Backing.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTxDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Consensus.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTxDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Duration) > 0 {
		i -= len(m.Duration)
		copy(dAtA[i:], m.Duration)
		i = encodeVarintTxDeposit(dAtA, i, uint64(len(m.Duration)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Market) > 0 {
		i -= len(m.Market)
		copy(dAtA[i:], m.Market)
		i = encodeVarintTxDeposit(dAtA, i, uint64(len(m.Market)))
		i--
		dAtA[i] = 0x12
	}
	if m.Time != 0 {
		i = encodeVarintTxDeposit(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTxDeposit(dAtA []byte, offset int, v uint64) int {
	offset -= sovTxDeposit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxDepositQuote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTxDeposit(uint64(m.Id))
	}
	l = len(m.Requester)
	if l > 0 {
		n += 1 + l + sovTxDeposit(uint64(l))
	}
	l = len(m.Deposit)
	if l > 0 {
		n += 1 + l + sovTxDeposit(uint64(l))
	}
	return n
}

func (m *DepositQuoteData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Time != 0 {
		n += 1 + sovTxDeposit(uint64(m.Time))
	}
	l = len(m.Market)
	if l > 0 {
		n += 1 + l + sovTxDeposit(uint64(l))
	}
	l = len(m.Duration)
	if l > 0 {
		n += 1 + l + sovTxDeposit(uint64(l))
	}
	l = m.Consensus.Size()
	n += 1 + l + sovTxDeposit(uint64(l))
	l = m.Backing.Size()
	n += 1 + l + sovTxDeposit(uint64(l))
	l = m.Commission.Size()
	n += 1 + l + sovTxDeposit(uint64(l))
	l = m.Reward.Size()
	n += 1 + l + sovTxDeposit(uint64(l))
	l = len(m.Adjustment)
	if l > 0 {
		n += 1 + l + sovTxDeposit(uint64(l))
	}
	return n
}

func sovTxDeposit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTxDeposit(x uint64) (n int) {
	return sovTxDeposit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxDepositQuote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxDeposit
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
			return fmt.Errorf("proto: TxDepositQuote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxDepositQuote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
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
				return fmt.Errorf("proto: wrong wireType = %d for field Requester", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
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
				return ErrInvalidLengthTxDeposit
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requester = append(m.Requester[:0], dAtA[iNdEx:postIndex]...)
			if m.Requester == nil {
				m.Requester = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
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
				return ErrInvalidLengthTxDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxDeposit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxDeposit
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
func (m *DepositQuoteData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxDeposit
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
			return fmt.Errorf("proto: DepositQuoteData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositQuoteData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Market", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
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
				return ErrInvalidLengthTxDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Market = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
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
				return ErrInvalidLengthTxDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Duration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consensus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
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
				return ErrInvalidLengthTxDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Consensus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Backing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
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
				return ErrInvalidLengthTxDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxDeposit
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
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
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
				return ErrInvalidLengthTxDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Commission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
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
				return ErrInvalidLengthTxDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTxDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Adjustment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxDeposit
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
				return ErrInvalidLengthTxDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Adjustment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxDeposit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxDeposit
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
func skipTxDeposit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxDeposit
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
					return 0, ErrIntOverflowTxDeposit
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
					return 0, ErrIntOverflowTxDeposit
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
				return 0, ErrInvalidLengthTxDeposit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTxDeposit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTxDeposit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTxDeposit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxDeposit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTxDeposit = fmt.Errorf("proto: unexpected end of group")
)