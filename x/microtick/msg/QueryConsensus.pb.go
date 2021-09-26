// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: microtick/v1beta1/msg/QueryConsensus.proto

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

type QueryConsensusRequest struct {
	Market string `protobuf:"bytes,1,opt,name=market,proto3" json:"market"`
}

func (m *QueryConsensusRequest) Reset()         { *m = QueryConsensusRequest{} }
func (m *QueryConsensusRequest) String() string { return proto.CompactTextString(m) }
func (*QueryConsensusRequest) ProtoMessage()    {}
func (*QueryConsensusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b478dc6a497c9e12, []int{0}
}
func (m *QueryConsensusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryConsensusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryConsensusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryConsensusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryConsensusRequest.Merge(m, src)
}
func (m *QueryConsensusRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryConsensusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryConsensusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryConsensusRequest proto.InternalMessageInfo

func (m *QueryConsensusRequest) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

type QueryConsensusResponse struct {
	Market       string        `protobuf:"bytes,1,opt,name=market,proto3" json:"market"`
	Consensus    types.DecCoin `protobuf:"bytes,2,opt,name=consensus,proto3" json:"consensus"`
	TotalBacking types.DecCoin `protobuf:"bytes,3,opt,name=total_backing,json=totalBacking,proto3" json:"total_backing"`
	TotalWeight  types.DecCoin `protobuf:"bytes,4,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight"`
}

func (m *QueryConsensusResponse) Reset()         { *m = QueryConsensusResponse{} }
func (m *QueryConsensusResponse) String() string { return proto.CompactTextString(m) }
func (*QueryConsensusResponse) ProtoMessage()    {}
func (*QueryConsensusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b478dc6a497c9e12, []int{1}
}
func (m *QueryConsensusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryConsensusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryConsensusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryConsensusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryConsensusResponse.Merge(m, src)
}
func (m *QueryConsensusResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryConsensusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryConsensusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryConsensusResponse proto.InternalMessageInfo

func (m *QueryConsensusResponse) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *QueryConsensusResponse) GetConsensus() types.DecCoin {
	if m != nil {
		return m.Consensus
	}
	return types.DecCoin{}
}

func (m *QueryConsensusResponse) GetTotalBacking() types.DecCoin {
	if m != nil {
		return m.TotalBacking
	}
	return types.DecCoin{}
}

func (m *QueryConsensusResponse) GetTotalWeight() types.DecCoin {
	if m != nil {
		return m.TotalWeight
	}
	return types.DecCoin{}
}

func init() {
	proto.RegisterType((*QueryConsensusRequest)(nil), "microtick.v1beta1.msg.QueryConsensusRequest")
	proto.RegisterType((*QueryConsensusResponse)(nil), "microtick.v1beta1.msg.QueryConsensusResponse")
}

func init() {
	proto.RegisterFile("microtick/v1beta1/msg/QueryConsensus.proto", fileDescriptor_b478dc6a497c9e12)
}

var fileDescriptor_b478dc6a497c9e12 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0x67, 0xf8, 0xbe, 0x90, 0x50, 0x74, 0xe1, 0x04, 0xcc, 0x84, 0x98, 0x42, 0x58, 0x11,
	0x4d, 0xda, 0xa0, 0x4b, 0x77, 0x83, 0x2e, 0x5d, 0x38, 0x89, 0xd1, 0xb8, 0x21, 0x33, 0x4d, 0x53,
	0x1a, 0x6c, 0x2f, 0x4e, 0x3b, 0xfe, 0x7b, 0x0a, 0x1f, 0xc8, 0x07, 0x60, 0xc9, 0xd2, 0x15, 0x31,
	0xb0, 0xe3, 0x29, 0x0c, 0x33, 0x23, 0x88, 0xab, 0x59, 0xf5, 0x9e, 0xdb, 0xf3, 0x3b, 0xbd, 0x4d,
	0x8b, 0x8e, 0x95, 0x64, 0x09, 0x58, 0xc9, 0xc6, 0xf4, 0xa9, 0x1f, 0x73, 0x1b, 0xf5, 0xa9, 0x32,
	0x82, 0x5e, 0xa7, 0x3c, 0x79, 0x1d, 0x80, 0x36, 0x5c, 0x9b, 0xd4, 0x90, 0x49, 0x02, 0x16, 0xbc,
	0xe6, 0xc6, 0x4b, 0x0a, 0x2f, 0x51, 0x46, 0xb4, 0x1a, 0x02, 0x04, 0x64, 0x0e, 0xba, 0xae, 0x72,
	0x73, 0x0b, 0x33, 0x30, 0x0a, 0x0c, 0x8d, 0x23, 0xc3, 0x37, 0xd1, 0x0c, 0xa4, 0xce, 0xf7, 0xbb,
	0xe7, 0xa8, 0xb9, 0x7b, 0x48, 0xc8, 0x1f, 0x53, 0x6e, 0xac, 0xd7, 0x45, 0x55, 0x15, 0x25, 0x63,
	0x6e, 0x7d, 0xb7, 0xe3, 0xf6, 0x6a, 0x01, 0x5a, 0xcd, 0xdb, 0x45, 0x27, 0x2c, 0xd6, 0xee, 0x47,
	0x05, 0x1d, 0xfe, 0xa5, 0xcd, 0x64, 0x5d, 0x97, 0xc1, 0xbd, 0x2b, 0x54, 0x63, 0x3f, 0xa0, 0x5f,
	0xe9, 0xb8, 0xbd, 0xfa, 0xe9, 0x11, 0xc9, 0xe7, 0x25, 0xeb, 0x79, 0x37, 0xd7, 0xbb, 0xe0, 0x6c,
	0x00, 0x52, 0x07, 0x07, 0xd3, 0x79, 0xdb, 0x59, 0xcd, 0xdb, 0x5b, 0x2c, 0xdc, 0x96, 0xde, 0x1d,
	0xda, 0xb7, 0x60, 0xa3, 0x87, 0x61, 0x1c, 0xb1, 0xb1, 0xd4, 0xc2, 0xff, 0x57, 0x22, 0xb2, 0x59,
	0x44, 0xee, 0xa2, 0xe1, 0x5e, 0x26, 0x83, 0x5c, 0x79, 0x37, 0x28, 0xd7, 0xc3, 0x67, 0x2e, 0xc5,
	0xc8, 0xfa, 0xff, 0x4b, 0x04, 0x37, 0x8a, 0xe0, 0x1d, 0x32, 0xac, 0x67, 0xea, 0x36, 0x13, 0xc1,
	0xe5, 0x74, 0x81, 0xdd, 0xd9, 0x02, 0xbb, 0x5f, 0x0b, 0xec, 0xbe, 0x2f, 0xb1, 0x33, 0x5b, 0x62,
	0xe7, 0x73, 0x89, 0x9d, 0xfb, 0x13, 0x21, 0xed, 0x28, 0x8d, 0x09, 0x03, 0x45, 0xb7, 0x3f, 0x43,
	0xd9, 0x37, 0xd0, 0x9c, 0xbe, 0xfc, 0x6e, 0x19, 0x11, 0x57, 0xb3, 0x97, 0x3c, 0xfb, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0xa9, 0xe1, 0x6c, 0x5e, 0x44, 0x02, 0x00, 0x00,
}

func (m *QueryConsensusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryConsensusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryConsensusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Market) > 0 {
		i -= len(m.Market)
		copy(dAtA[i:], m.Market)
		i = encodeVarintQueryConsensus(dAtA, i, uint64(len(m.Market)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryConsensusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryConsensusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryConsensusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TotalWeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryConsensus(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.TotalBacking.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryConsensus(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Consensus.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryConsensus(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Market) > 0 {
		i -= len(m.Market)
		copy(dAtA[i:], m.Market)
		i = encodeVarintQueryConsensus(dAtA, i, uint64(len(m.Market)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryConsensus(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryConsensus(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryConsensusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Market)
	if l > 0 {
		n += 1 + l + sovQueryConsensus(uint64(l))
	}
	return n
}

func (m *QueryConsensusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Market)
	if l > 0 {
		n += 1 + l + sovQueryConsensus(uint64(l))
	}
	l = m.Consensus.Size()
	n += 1 + l + sovQueryConsensus(uint64(l))
	l = m.TotalBacking.Size()
	n += 1 + l + sovQueryConsensus(uint64(l))
	l = m.TotalWeight.Size()
	n += 1 + l + sovQueryConsensus(uint64(l))
	return n
}

func sovQueryConsensus(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryConsensus(x uint64) (n int) {
	return sovQueryConsensus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryConsensusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryConsensus
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
			return fmt.Errorf("proto: QueryConsensusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryConsensusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Market", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryConsensus
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
				return ErrInvalidLengthQueryConsensus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryConsensus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Market = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryConsensus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryConsensus
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
func (m *QueryConsensusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryConsensus
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
			return fmt.Errorf("proto: QueryConsensusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryConsensusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Market", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryConsensus
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
				return ErrInvalidLengthQueryConsensus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryConsensus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Market = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consensus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryConsensus
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
				return ErrInvalidLengthQueryConsensus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryConsensus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Consensus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBacking", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryConsensus
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
				return ErrInvalidLengthQueryConsensus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryConsensus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalBacking.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryConsensus
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
				return ErrInvalidLengthQueryConsensus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryConsensus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryConsensus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryConsensus
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
func skipQueryConsensus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryConsensus
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
					return 0, ErrIntOverflowQueryConsensus
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
					return 0, ErrIntOverflowQueryConsensus
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
				return 0, ErrInvalidLengthQueryConsensus
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryConsensus
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryConsensus
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryConsensus        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryConsensus          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryConsensus = fmt.Errorf("proto: unexpected end of group")
)
