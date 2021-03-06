// Code generated by protoc-gen-gogo.
// source: common/ledger.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Contains information about the blockchain ledger such as height, current
// block hash, and previous block hash.
type BlockchainInfo struct {
	Height            uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	CurrentBlockHash  []byte `protobuf:"bytes,2,opt,name=currentBlockHash,proto3" json:"currentBlockHash,omitempty"`
	PreviousBlockHash []byte `protobuf:"bytes,3,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
}

func (m *BlockchainInfo) Reset()                    { *m = BlockchainInfo{} }
func (m *BlockchainInfo) String() string            { return proto.CompactTextString(m) }
func (*BlockchainInfo) ProtoMessage()               {}
func (*BlockchainInfo) Descriptor() ([]byte, []int) { return fileDescriptorLedger, []int{0} }

func (m *BlockchainInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockchainInfo) GetCurrentBlockHash() []byte {
	if m != nil {
		return m.CurrentBlockHash
	}
	return nil
}

func (m *BlockchainInfo) GetPreviousBlockHash() []byte {
	if m != nil {
		return m.PreviousBlockHash
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockchainInfo)(nil), "common.BlockchainInfo")
}
func (m *BlockchainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockchainInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLedger(dAtA, i, uint64(m.Height))
	}
	if len(m.CurrentBlockHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLedger(dAtA, i, uint64(len(m.CurrentBlockHash)))
		i += copy(dAtA[i:], m.CurrentBlockHash)
	}
	if len(m.PreviousBlockHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLedger(dAtA, i, uint64(len(m.PreviousBlockHash)))
		i += copy(dAtA[i:], m.PreviousBlockHash)
	}
	return i, nil
}

func encodeFixed64Ledger(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Ledger(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLedger(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BlockchainInfo) Size() (n int) {
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovLedger(uint64(m.Height))
	}
	l = len(m.CurrentBlockHash)
	if l > 0 {
		n += 1 + l + sovLedger(uint64(l))
	}
	l = len(m.PreviousBlockHash)
	if l > 0 {
		n += 1 + l + sovLedger(uint64(l))
	}
	return n
}

func sovLedger(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLedger(x uint64) (n int) {
	return sovLedger(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockchainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLedger
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockchainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockchainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLedger
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentBlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLedger
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLedger
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentBlockHash = append(m.CurrentBlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.CurrentBlockHash == nil {
				m.CurrentBlockHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousBlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLedger
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthLedger
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousBlockHash = append(m.PreviousBlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PreviousBlockHash == nil {
				m.PreviousBlockHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLedger(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLedger
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
func skipLedger(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLedger
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
					return 0, ErrIntOverflowLedger
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLedger
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLedger
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLedger
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLedger(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLedger = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLedger   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("common/ledger.proto", fileDescriptorLedger) }

var fileDescriptorLedger = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0xcf, 0x49, 0x4d, 0x49, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x83, 0x08, 0x2a, 0x35, 0x31, 0x72, 0xf1, 0x39, 0xe5, 0xe4, 0x27, 0x67, 0x27, 0x67, 0x24,
	0x66, 0xe6, 0x79, 0xe6, 0xa5, 0xe5, 0x0b, 0x89, 0x71, 0xb1, 0x65, 0xa4, 0x66, 0xa6, 0x67, 0x94,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x41, 0x79, 0x42, 0x5a, 0x5c, 0x02, 0xc9, 0xa5, 0x45,
	0x45, 0xa9, 0x79, 0x25, 0x60, 0x0d, 0x1e, 0x89, 0xc5, 0x19, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x41, 0x18, 0xe2, 0x42, 0x3a, 0x5c, 0x82, 0x05, 0x45, 0xa9, 0x65, 0x99, 0xf9, 0xa5, 0xc5, 0x08,
	0xc5, 0xcc, 0x60, 0xc5, 0x98, 0x12, 0x4e, 0xc9, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x5c, 0x2a, 0xf9, 0x45, 0xe9, 0x7a, 0x19,
	0x95, 0x05, 0xa9, 0x45, 0x50, 0x57, 0xa7, 0x25, 0x26, 0x15, 0x65, 0x26, 0x43, 0x1c, 0x5f, 0xac,
	0x07, 0x71, 0x7c, 0x94, 0x76, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x12, 0x88, 0xab, 0x8f, 0xa4, 0x58,
	0x1f, 0xa2, 0x58, 0x1f, 0xa2, 0x58, 0x1f, 0xa2, 0x38, 0x89, 0x0d, 0xcc, 0x35, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x9f, 0xab, 0x69, 0x54, 0x0f, 0x01, 0x00, 0x00,
}
