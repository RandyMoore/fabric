// Code generated by protoc-gen-gogo.
// source: config.proto
// DO NOT EDIT!

/*
	Package sbft is a generated protocol buffer package.

	It is generated from these files:
		config.proto

	It has these top-level messages:
		ConsensusConfig
		JsonConfig
		Peer
*/
package sbft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import simplebft "github.com/hyperledger/fabric/orderer/sbft/simplebft"

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

type ConsensusConfig struct {
	Consensus *simplebft.Config `protobuf:"bytes,1,opt,name=consensus" json:"consensus,omitempty"`
	Peers     map[string][]byte `protobuf:"bytes,2,rep,name=peers" json:"peers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ConsensusConfig) Reset()                    { *m = ConsensusConfig{} }
func (m *ConsensusConfig) String() string            { return proto.CompactTextString(m) }
func (*ConsensusConfig) ProtoMessage()               {}
func (*ConsensusConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *ConsensusConfig) GetConsensus() *simplebft.Config {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *ConsensusConfig) GetPeers() map[string][]byte {
	if m != nil {
		return m.Peers
	}
	return nil
}

type JsonConfig struct {
	Consensus *simplebft.Config `protobuf:"bytes,1,opt,name=consensus" json:"consensus,omitempty"`
	Peers     []*Peer           `protobuf:"bytes,2,rep,name=peers" json:"peers,omitempty"`
}

func (m *JsonConfig) Reset()                    { *m = JsonConfig{} }
func (m *JsonConfig) String() string            { return proto.CompactTextString(m) }
func (*JsonConfig) ProtoMessage()               {}
func (*JsonConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *JsonConfig) GetConsensus() *simplebft.Config {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *JsonConfig) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type Peer struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Cert    string `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *Peer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Peer) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func init() {
	proto.RegisterType((*ConsensusConfig)(nil), "sbft.consensus_config")
	proto.RegisterType((*JsonConfig)(nil), "sbft.json_config")
	proto.RegisterType((*Peer)(nil), "sbft.peer")
}
func (m *ConsensusConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Consensus != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Consensus.Size()))
		n1, err := m.Consensus.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Peers) > 0 {
		for k, _ := range m.Peers {
			dAtA[i] = 0x12
			i++
			v := m.Peers[k]
			byteSize := 0
			if len(v) > 0 {
				byteSize = 1 + len(v) + sovConfig(uint64(len(v)))
			}
			mapSize := 1 + len(k) + sovConfig(uint64(len(k))) + byteSize
			i = encodeVarintConfig(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintConfig(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if len(v) > 0 {
				dAtA[i] = 0x12
				i++
				i = encodeVarintConfig(dAtA, i, uint64(len(v)))
				i += copy(dAtA[i:], v)
			}
		}
	}
	return i, nil
}

func (m *JsonConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JsonConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Consensus != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(m.Consensus.Size()))
		n2, err := m.Consensus.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Peers) > 0 {
		for _, msg := range m.Peers {
			dAtA[i] = 0x12
			i++
			i = encodeVarintConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Cert) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Cert)))
		i += copy(dAtA[i:], m.Cert)
	}
	return i, nil
}

func encodeFixed64Config(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Config(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ConsensusConfig) Size() (n int) {
	var l int
	_ = l
	if m.Consensus != nil {
		l = m.Consensus.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.Peers) > 0 {
		for k, v := range m.Peers {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovConfig(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovConfig(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovConfig(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *JsonConfig) Size() (n int) {
	var l int
	_ = l
	if m.Consensus != nil {
		l = m.Consensus.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if len(m.Peers) > 0 {
		for _, e := range m.Peers {
			l = e.Size()
			n += 1 + l + sovConfig(uint64(l))
		}
	}
	return n
}

func (m *Peer) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Cert)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConsensusConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: consensus_config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: consensus_config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consensus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Consensus == nil {
				m.Consensus = &simplebft.Config{}
			}
			if err := m.Consensus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthConfig
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Peers == nil {
				m.Peers = make(map[string][]byte)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapbyteLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapbyteLen |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intMapbyteLen := int(mapbyteLen)
				if intMapbyteLen < 0 {
					return ErrInvalidLengthConfig
				}
				postbytesIndex := iNdEx + intMapbyteLen
				if postbytesIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := make([]byte, mapbyteLen)
				copy(mapvalue, dAtA[iNdEx:postbytesIndex])
				iNdEx = postbytesIndex
				m.Peers[mapkey] = mapvalue
			} else {
				var mapvalue []byte
				m.Peers[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func (m *JsonConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: json_config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: json_config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Consensus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Consensus == nil {
				m.Consensus = &simplebft.Config{}
			}
			if err := m.Consensus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peers = append(m.Peers, &Peer{})
			if err := m.Peers[len(m.Peers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x91, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x71, 0x2f, 0xa0, 0x9e, 0x76, 0x28, 0x16, 0x43, 0xe8, 0x10, 0x85, 0x4c, 0x11, 0x83,
	0x2d, 0x15, 0x24, 0x2a, 0xc4, 0x04, 0x62, 0x47, 0x19, 0x59, 0x20, 0x97, 0x93, 0x34, 0x90, 0xc6,
	0x91, 0xed, 0x20, 0xe5, 0x4d, 0x78, 0x12, 0x9e, 0x81, 0x91, 0x47, 0x40, 0xe1, 0x45, 0x90, 0x93,
	0xd2, 0x08, 0x46, 0xb6, 0xff, 0x1c, 0x7d, 0xc7, 0xff, 0x27, 0x19, 0x66, 0x91, 0x28, 0x92, 0x2c,
	0x65, 0xa5, 0x14, 0x5a, 0xd0, 0x91, 0x0a, 0x13, 0xbd, 0x38, 0x56, 0xd9, 0xa6, 0xcc, 0x31, 0x4c,
	0x34, 0xdf, 0xa5, 0x0e, 0x70, 0xdf, 0x08, 0xcc, 0x23, 0x51, 0x28, 0x2c, 0x54, 0xa5, 0x1e, 0xba,
	0x5b, 0xca, 0x61, 0xb2, 0xdb, 0x59, 0xc4, 0x21, 0xde, 0x74, 0x79, 0xc8, 0xfa, 0xcb, 0x9b, 0x96,
	0xf2, 0x7b, 0x86, 0x5e, 0xc0, 0xb8, 0x44, 0x94, 0xca, 0x1a, 0x38, 0x43, 0x6f, 0xba, 0x3c, 0x61,
	0xa6, 0x96, 0xfd, 0x7d, 0x97, 0xdd, 0x19, 0xe6, 0xb6, 0xd0, 0xb2, 0xf6, 0x3b, 0x7e, 0xb1, 0x02,
	0xe8, 0x97, 0x74, 0x0e, 0xc3, 0x67, 0xac, 0xdb, 0xc6, 0x89, 0x6f, 0x22, 0x3d, 0x82, 0xf1, 0x4b,
	0x90, 0x57, 0x68, 0x0d, 0x1c, 0xe2, 0xcd, 0xfc, 0x6e, 0xb8, 0x1c, 0xac, 0x88, 0xfb, 0x08, 0xd3,
	0x27, 0x25, 0x8a, 0x7f, 0x2b, 0x3b, 0xbf, 0x95, 0xa1, 0x53, 0x36, 0xab, 0xad, 0x9b, 0x7b, 0x0e,
	0x23, 0x13, 0xa8, 0x05, 0x07, 0x41, 0x1c, 0x4b, 0x54, 0x6a, 0x6b, 0xf6, 0x33, 0x52, 0x0a, 0xa3,
	0x08, 0xa5, 0x6e, 0xe5, 0x26, 0x7e, 0x9b, 0xaf, 0xaf, 0xde, 0x1b, 0x9b, 0x7c, 0x34, 0x36, 0xf9,
	0x6c, 0x6c, 0xf2, 0xfa, 0x65, 0xef, 0xdd, 0x9f, 0xa6, 0x99, 0x5e, 0x57, 0x21, 0x8b, 0xc4, 0x86,
	0xaf, 0xeb, 0x12, 0x65, 0x8e, 0x71, 0x8a, 0x92, 0x27, 0x41, 0x28, 0xb3, 0x88, 0x0b, 0x19, 0xa3,
	0x44, 0xc9, 0x4d, 0x7f, 0xb8, 0xdf, 0xfe, 0xca, 0xd9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0xb3, 0xd9, 0x62, 0xc6, 0x01, 0x00, 0x00,
}
