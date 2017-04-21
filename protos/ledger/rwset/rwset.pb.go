// Code generated by protoc-gen-gogo.
// source: ledger/rwset/rwset.proto
// DO NOT EDIT!

/*
	Package rwset is a generated protocol buffer package.

	It is generated from these files:
		ledger/rwset/rwset.proto

	It has these top-level messages:
		TxReadWriteSet
		NsReadWriteSet
*/
package rwset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TxReadWriteSet_DataModel int32

const (
	TxReadWriteSet_KV TxReadWriteSet_DataModel = 0
)

var TxReadWriteSet_DataModel_name = map[int32]string{
	0: "KV",
}
var TxReadWriteSet_DataModel_value = map[string]int32{
	"KV": 0,
}

func (x TxReadWriteSet_DataModel) String() string {
	return proto.EnumName(TxReadWriteSet_DataModel_name, int32(x))
}
func (TxReadWriteSet_DataModel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorRwset, []int{0, 0}
}

// TxReadWriteSet encapsulates a read-write set for a transaction
// DataModel specifies the enum value of the data model
// ns_rwset field specifies a list of chaincode specific read-write set (one for each chaincode)
type TxReadWriteSet struct {
	DataModel TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,proto3,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsRwset   []*NsReadWriteSet        `protobuf:"bytes,2,rep,name=ns_rwset,json=nsRwset" json:"ns_rwset,omitempty"`
}

func (m *TxReadWriteSet) Reset()                    { *m = TxReadWriteSet{} }
func (m *TxReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*TxReadWriteSet) ProtoMessage()               {}
func (*TxReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptorRwset, []int{0} }

func (m *TxReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if m != nil {
		return m.DataModel
	}
	return TxReadWriteSet_KV
}

func (m *TxReadWriteSet) GetNsRwset() []*NsReadWriteSet {
	if m != nil {
		return m.NsRwset
	}
	return nil
}

// NsReadWriteSet encapsulates the read-write set for a chaincode
type NsReadWriteSet struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Rwset     []byte `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"`
}

func (m *NsReadWriteSet) Reset()                    { *m = NsReadWriteSet{} }
func (m *NsReadWriteSet) String() string            { return proto.CompactTextString(m) }
func (*NsReadWriteSet) ProtoMessage()               {}
func (*NsReadWriteSet) Descriptor() ([]byte, []int) { return fileDescriptorRwset, []int{1} }

func (m *NsReadWriteSet) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NsReadWriteSet) GetRwset() []byte {
	if m != nil {
		return m.Rwset
	}
	return nil
}

func init() {
	proto.RegisterType((*TxReadWriteSet)(nil), "rwset.TxReadWriteSet")
	proto.RegisterType((*NsReadWriteSet)(nil), "rwset.NsReadWriteSet")
	proto.RegisterEnum("rwset.TxReadWriteSet_DataModel", TxReadWriteSet_DataModel_name, TxReadWriteSet_DataModel_value)
}
func (m *TxReadWriteSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxReadWriteSet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DataModel != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRwset(dAtA, i, uint64(m.DataModel))
	}
	if len(m.NsRwset) > 0 {
		for _, msg := range m.NsRwset {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRwset(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *NsReadWriteSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NsReadWriteSet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRwset(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Rwset) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRwset(dAtA, i, uint64(len(m.Rwset)))
		i += copy(dAtA[i:], m.Rwset)
	}
	return i, nil
}

func encodeFixed64Rwset(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Rwset(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRwset(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TxReadWriteSet) Size() (n int) {
	var l int
	_ = l
	if m.DataModel != 0 {
		n += 1 + sovRwset(uint64(m.DataModel))
	}
	if len(m.NsRwset) > 0 {
		for _, e := range m.NsRwset {
			l = e.Size()
			n += 1 + l + sovRwset(uint64(l))
		}
	}
	return n
}

func (m *NsReadWriteSet) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovRwset(uint64(l))
	}
	l = len(m.Rwset)
	if l > 0 {
		n += 1 + l + sovRwset(uint64(l))
	}
	return n
}

func sovRwset(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRwset(x uint64) (n int) {
	return sovRwset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxReadWriteSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRwset
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
			return fmt.Errorf("proto: TxReadWriteSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxReadWriteSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataModel", wireType)
			}
			m.DataModel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRwset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataModel |= (TxReadWriteSet_DataModel(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NsRwset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRwset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRwset
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NsRwset = append(m.NsRwset, &NsReadWriteSet{})
			if err := m.NsRwset[len(m.NsRwset)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRwset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRwset
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
func (m *NsReadWriteSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRwset
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
			return fmt.Errorf("proto: NsReadWriteSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NsReadWriteSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRwset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRwset
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rwset", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRwset
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
				return ErrInvalidLengthRwset
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rwset = append(m.Rwset[:0], dAtA[iNdEx:postIndex]...)
			if m.Rwset == nil {
				m.Rwset = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRwset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRwset
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
func skipRwset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRwset
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
					return 0, ErrIntOverflowRwset
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
					return 0, ErrIntOverflowRwset
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
				return 0, ErrInvalidLengthRwset
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRwset
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
				next, err := skipRwset(dAtA[start:])
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
	ErrInvalidLengthRwset = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRwset   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ledger/rwset/rwset.proto", fileDescriptorRwset) }

var fileDescriptorRwset = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0x49, 0x4d, 0x49,
	0x4f, 0x2d, 0xd2, 0x2f, 0x2a, 0x2f, 0x4e, 0x2d, 0x81, 0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0xac, 0x60, 0x8e, 0xd2, 0x74, 0x46, 0x2e, 0xbe, 0x90, 0x8a, 0xa0, 0xd4, 0xc4, 0x94, 0xf0,
	0xa2, 0xcc, 0x92, 0xd4, 0xe0, 0xd4, 0x12, 0x21, 0x3b, 0x2e, 0xae, 0x94, 0xc4, 0x92, 0xc4, 0xf8,
	0xdc, 0xfc, 0x94, 0xd4, 0x1c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x79, 0x3d, 0x88, 0x5e,
	0x54, 0xa5, 0x7a, 0x2e, 0x89, 0x25, 0x89, 0xbe, 0x20, 0x65, 0x41, 0x9c, 0x29, 0x30, 0xa6, 0x90,
	0x01, 0x17, 0x47, 0x5e, 0x71, 0x3c, 0x58, 0xbd, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x28,
	0x54, 0xb7, 0x5f, 0x31, 0xb2, 0xee, 0x20, 0xf6, 0xbc, 0xe2, 0x20, 0xb0, 0x23, 0x84, 0xb9, 0x38,
	0xe1, 0x26, 0x09, 0xb1, 0x71, 0x31, 0x79, 0x87, 0x09, 0x30, 0x28, 0xb9, 0x70, 0xf1, 0xa1, 0xaa,
	0x17, 0x92, 0xe1, 0xe2, 0xcc, 0x4b, 0xcc, 0x4d, 0x2d, 0x2e, 0x48, 0x4c, 0x4e, 0x05, 0xbb, 0x8b,
	0x33, 0x08, 0x21, 0x20, 0x24, 0xc2, 0xc5, 0x0a, 0xb3, 0x93, 0x51, 0x83, 0x27, 0x08, 0xc2, 0x71,
	0xca, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c,
	0x96, 0x63, 0xe0, 0xd2, 0xca, 0x2f, 0x4a, 0xd7, 0xcb, 0xa8, 0x2c, 0x48, 0x2d, 0x82, 0x84, 0x8d,
	0x5e, 0x5a, 0x62, 0x52, 0x51, 0x66, 0x32, 0x24, 0x58, 0x8a, 0xf5, 0xa0, 0x82, 0x60, 0xdd, 0x51,
	0x86, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x48, 0x5a, 0xf4, 0x21,
	0x5a, 0xf4, 0x21, 0x5a, 0xf4, 0x91, 0xc3, 0x38, 0x89, 0x0d, 0x2c, 0x68, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x63, 0x2c, 0xfc, 0x76, 0x7a, 0x01, 0x00, 0x00,
}
