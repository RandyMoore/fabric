// Code generated by protoc-gen-gogo.
// source: orderer/configuration.proto
// DO NOT EDIT!

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptorConfiguration, []int{0} }

func (m *ConsensusType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=maxMessageCount,proto3" json:"maxMessageCount,omitempty"`
	// The byte count of the serialized messages in a batch cannot
	// exceed this value.
	AbsoluteMaxBytes uint32 `protobuf:"varint,2,opt,name=absoluteMaxBytes,proto3" json:"absoluteMaxBytes,omitempty"`
	// The byte count of the serialized messages in a batch should not
	// exceed this value.
	PreferredMaxBytes uint32 `protobuf:"varint,3,opt,name=preferredMaxBytes,proto3" json:"preferredMaxBytes,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptorConfiguration, []int{1} }

func (m *BatchSize) GetMaxMessageCount() uint32 {
	if m != nil {
		return m.MaxMessageCount
	}
	return 0
}

func (m *BatchSize) GetAbsoluteMaxBytes() uint32 {
	if m != nil {
		return m.AbsoluteMaxBytes
	}
	return 0
}

func (m *BatchSize) GetPreferredMaxBytes() uint32 {
	if m != nil {
		return m.PreferredMaxBytes
	}
	return 0
}

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout string `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (m *BatchTimeout) Reset()                    { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string            { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()               {}
func (*BatchTimeout) Descriptor() ([]byte, []int) { return fileDescriptorConfiguration, []int{2} }

func (m *BatchTimeout) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

// When submitting a new chain configuration transaction to create a new chain,
// the first configuration item must be of type Orderer with Key CreationPolicy
// and contents of a Marshaled CreationPolicy. The policy should be set to the
// policy which was supplied by the ordering service for the client's chain
// creation. The digest should be the hash of the concatenation of the remaining
// ConfigurationItem bytes. The signatures of the configuration item should
// satisfy the policy for chain creation.
type CreationPolicy struct {
	// The name of the policy which should be used to validate the creation of
	// this chain
	Policy string `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (m *CreationPolicy) Reset()                    { *m = CreationPolicy{} }
func (m *CreationPolicy) String() string            { return proto.CompactTextString(m) }
func (*CreationPolicy) ProtoMessage()               {}
func (*CreationPolicy) Descriptor() ([]byte, []int) { return fileDescriptorConfiguration, []int{3} }

func (m *CreationPolicy) GetPolicy() string {
	if m != nil {
		return m.Policy
	}
	return ""
}

// ChainCreationPolicyNames is the set of policies which may be invoked for chain creation
type ChainCreationPolicyNames struct {
	// A list of policies, in evaluation these are 'or'-ed, note this is not a proper policy
	// because implementing referential policies in a general way is difficult, and dangerous
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *ChainCreationPolicyNames) Reset()         { *m = ChainCreationPolicyNames{} }
func (m *ChainCreationPolicyNames) String() string { return proto.CompactTextString(m) }
func (*ChainCreationPolicyNames) ProtoMessage()    {}
func (*ChainCreationPolicyNames) Descriptor() ([]byte, []int) {
	return fileDescriptorConfiguration, []int{4}
}

func (m *ChainCreationPolicyNames) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers []string `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
}

func (m *KafkaBrokers) Reset()                    { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string            { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()               {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) { return fileDescriptorConfiguration, []int{5} }

func (m *KafkaBrokers) GetBrokers() []string {
	if m != nil {
		return m.Brokers
	}
	return nil
}

// ChannelRestrictions is the mssage which conveys restrictions on channel creation for an orderer
type ChannelRestrictions struct {
	MaxCount uint64 `protobuf:"varint,1,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"`
}

func (m *ChannelRestrictions) Reset()                    { *m = ChannelRestrictions{} }
func (m *ChannelRestrictions) String() string            { return proto.CompactTextString(m) }
func (*ChannelRestrictions) ProtoMessage()               {}
func (*ChannelRestrictions) Descriptor() ([]byte, []int) { return fileDescriptorConfiguration, []int{6} }

func (m *ChannelRestrictions) GetMaxCount() uint64 {
	if m != nil {
		return m.MaxCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*CreationPolicy)(nil), "orderer.CreationPolicy")
	proto.RegisterType((*ChainCreationPolicyNames)(nil), "orderer.ChainCreationPolicyNames")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
	proto.RegisterType((*ChannelRestrictions)(nil), "orderer.ChannelRestrictions")
}
func (m *ConsensusType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfiguration(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	return i, nil
}

func (m *BatchSize) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchSize) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MaxMessageCount != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConfiguration(dAtA, i, uint64(m.MaxMessageCount))
	}
	if m.AbsoluteMaxBytes != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConfiguration(dAtA, i, uint64(m.AbsoluteMaxBytes))
	}
	if m.PreferredMaxBytes != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintConfiguration(dAtA, i, uint64(m.PreferredMaxBytes))
	}
	return i, nil
}

func (m *BatchTimeout) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchTimeout) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Timeout) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfiguration(dAtA, i, uint64(len(m.Timeout)))
		i += copy(dAtA[i:], m.Timeout)
	}
	return i, nil
}

func (m *CreationPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreationPolicy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Policy) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfiguration(dAtA, i, uint64(len(m.Policy)))
		i += copy(dAtA[i:], m.Policy)
	}
	return i, nil
}

func (m *ChainCreationPolicyNames) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainCreationPolicyNames) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Names) > 0 {
		for _, s := range m.Names {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *KafkaBrokers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KafkaBrokers) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Brokers) > 0 {
		for _, s := range m.Brokers {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *ChannelRestrictions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChannelRestrictions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MaxCount != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConfiguration(dAtA, i, uint64(m.MaxCount))
	}
	return i, nil
}

func encodeFixed64Configuration(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Configuration(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConfiguration(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ConsensusType) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovConfiguration(uint64(l))
	}
	return n
}

func (m *BatchSize) Size() (n int) {
	var l int
	_ = l
	if m.MaxMessageCount != 0 {
		n += 1 + sovConfiguration(uint64(m.MaxMessageCount))
	}
	if m.AbsoluteMaxBytes != 0 {
		n += 1 + sovConfiguration(uint64(m.AbsoluteMaxBytes))
	}
	if m.PreferredMaxBytes != 0 {
		n += 1 + sovConfiguration(uint64(m.PreferredMaxBytes))
	}
	return n
}

func (m *BatchTimeout) Size() (n int) {
	var l int
	_ = l
	l = len(m.Timeout)
	if l > 0 {
		n += 1 + l + sovConfiguration(uint64(l))
	}
	return n
}

func (m *CreationPolicy) Size() (n int) {
	var l int
	_ = l
	l = len(m.Policy)
	if l > 0 {
		n += 1 + l + sovConfiguration(uint64(l))
	}
	return n
}

func (m *ChainCreationPolicyNames) Size() (n int) {
	var l int
	_ = l
	if len(m.Names) > 0 {
		for _, s := range m.Names {
			l = len(s)
			n += 1 + l + sovConfiguration(uint64(l))
		}
	}
	return n
}

func (m *KafkaBrokers) Size() (n int) {
	var l int
	_ = l
	if len(m.Brokers) > 0 {
		for _, s := range m.Brokers {
			l = len(s)
			n += 1 + l + sovConfiguration(uint64(l))
		}
	}
	return n
}

func (m *ChannelRestrictions) Size() (n int) {
	var l int
	_ = l
	if m.MaxCount != 0 {
		n += 1 + sovConfiguration(uint64(m.MaxCount))
	}
	return n
}

func sovConfiguration(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfiguration(x uint64) (n int) {
	return sovConfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConsensusType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfiguration
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
			return fmt.Errorf("proto: ConsensusType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
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
				return ErrInvalidLengthConfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfiguration
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
func (m *BatchSize) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfiguration
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
			return fmt.Errorf("proto: BatchSize: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchSize: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMessageCount", wireType)
			}
			m.MaxMessageCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMessageCount |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbsoluteMaxBytes", wireType)
			}
			m.AbsoluteMaxBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AbsoluteMaxBytes |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreferredMaxBytes", wireType)
			}
			m.PreferredMaxBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PreferredMaxBytes |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfiguration
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
func (m *BatchTimeout) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfiguration
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
			return fmt.Errorf("proto: BatchTimeout: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchTimeout: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
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
				return ErrInvalidLengthConfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timeout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfiguration
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
func (m *CreationPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfiguration
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
			return fmt.Errorf("proto: CreationPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreationPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
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
				return ErrInvalidLengthConfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Policy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfiguration
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
func (m *ChainCreationPolicyNames) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfiguration
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
			return fmt.Errorf("proto: ChainCreationPolicyNames: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainCreationPolicyNames: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Names", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
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
				return ErrInvalidLengthConfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Names = append(m.Names, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfiguration
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
func (m *KafkaBrokers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfiguration
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
			return fmt.Errorf("proto: KafkaBrokers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KafkaBrokers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Brokers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
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
				return ErrInvalidLengthConfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Brokers = append(m.Brokers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfiguration
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
func (m *ChannelRestrictions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfiguration
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
			return fmt.Errorf("proto: ChannelRestrictions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChannelRestrictions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCount", wireType)
			}
			m.MaxCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCount |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfiguration
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
func skipConfiguration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfiguration
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
					return 0, ErrIntOverflowConfiguration
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
					return 0, ErrIntOverflowConfiguration
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
				return 0, ErrInvalidLengthConfiguration
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfiguration
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
				next, err := skipConfiguration(dAtA[start:])
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
	ErrInvalidLengthConfiguration = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfiguration   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptorConfiguration) }

var fileDescriptorConfiguration = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x8a, 0xd4, 0x40,
	0x10, 0xc6, 0x8d, 0xbb, 0xee, 0x9a, 0xc2, 0xf5, 0x4f, 0x2b, 0x12, 0x58, 0x08, 0x43, 0x44, 0x08,
	0x32, 0x24, 0xa2, 0x6f, 0x90, 0x1c, 0x65, 0x44, 0xe2, 0x9c, 0xbc, 0x48, 0x27, 0x53, 0x49, 0x9a,
	0x49, 0xba, 0x43, 0x75, 0x07, 0x12, 0x5f, 0xc2, 0xab, 0x8f, 0xe4, 0xd1, 0x47, 0x90, 0xf1, 0x45,
	0x24, 0x9d, 0xcc, 0xa0, 0x3b, 0xb7, 0xef, 0xab, 0xfa, 0x41, 0x7f, 0x5d, 0x1f, 0xdc, 0x2a, 0xda,
	0x21, 0x21, 0xc5, 0x85, 0x92, 0xa5, 0xa8, 0x7a, 0xe2, 0x46, 0x28, 0x19, 0x75, 0xa4, 0x8c, 0x62,
	0xd7, 0xcb, 0x32, 0x78, 0x05, 0x37, 0xa9, 0x92, 0x1a, 0xa5, 0xee, 0xf5, 0x76, 0xec, 0x90, 0x31,
	0xb8, 0x34, 0x63, 0x87, 0x9e, 0xb3, 0x72, 0x42, 0x37, 0xb3, 0x3a, 0xf8, 0xee, 0x80, 0x9b, 0x70,
	0x53, 0xd4, 0x9f, 0xc5, 0x37, 0x64, 0x21, 0x3c, 0x69, 0xf9, 0xb0, 0x41, 0xad, 0x79, 0x85, 0xa9,
	0xea, 0xa5, 0xb1, 0xf0, 0x4d, 0x76, 0x77, 0xcc, 0xde, 0xc0, 0x53, 0x9e, 0x6b, 0xd5, 0xf4, 0x06,
	0x37, 0x7c, 0x48, 0x46, 0x83, 0xda, 0xbb, 0x6f, 0xd1, 0xb3, 0x39, 0x5b, 0xc3, 0xb3, 0x8e, 0xb0,
	0x44, 0x22, 0xdc, 0x9d, 0xe0, 0x0b, 0x0b, 0x9f, 0x2f, 0x82, 0x10, 0x1e, 0xd9, 0x40, 0x5b, 0xd1,
	0xa2, 0xea, 0x0d, 0xf3, 0xe0, 0xda, 0xcc, 0x72, 0x09, 0x7e, 0xb4, 0x41, 0x08, 0x8f, 0x53, 0x42,
	0xfb, 0xf7, 0x4f, 0xaa, 0x11, 0xc5, 0xc8, 0x5e, 0xc2, 0x55, 0x67, 0xd5, 0x82, 0x2e, 0x2e, 0x78,
	0x0b, 0x5e, 0x5a, 0x73, 0x21, 0xff, 0xc7, 0x3f, 0xf2, 0x16, 0x35, 0x7b, 0x01, 0x0f, 0xe4, 0x24,
	0x3c, 0x67, 0x75, 0x11, 0xba, 0xd9, 0x6c, 0xa6, 0x14, 0x1f, 0x78, 0xb9, 0xe7, 0x09, 0xa9, 0x3d,
	0x92, 0x9e, 0x52, 0xe4, 0xb3, 0x5c, 0xb8, 0xa3, 0x0d, 0xde, 0xc1, 0xf3, 0xb4, 0xe6, 0x52, 0x62,
	0x93, 0xa1, 0x36, 0x24, 0x8a, 0xe9, 0x01, 0xcd, 0x6e, 0xc1, 0x6d, 0xf9, 0xf0, 0xb5, 0x38, 0x1d,
	0xf1, 0x32, 0x7b, 0xd8, 0xf2, 0xc1, 0x5e, 0x2f, 0xc1, 0x9f, 0x07, 0xdf, 0xf9, 0x75, 0xf0, 0x9d,
	0xdf, 0x07, 0xdf, 0xf9, 0xf1, 0xc7, 0xbf, 0x07, 0xaf, 0x15, 0x55, 0x51, 0x3d, 0x76, 0x48, 0x0d,
	0xee, 0x2a, 0xa4, 0xa8, 0xe4, 0x39, 0x89, 0x62, 0xee, 0x54, 0x47, 0x4b, 0xa7, 0x5f, 0xd6, 0x95,
	0x30, 0x75, 0x9f, 0x47, 0x85, 0x6a, 0xe3, 0x7f, 0xe8, 0x78, 0xa6, 0xe3, 0x99, 0x8e, 0x17, 0x3a,
	0xbf, 0xb2, 0xfe, 0xfd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xe1, 0x41, 0x35, 0x30, 0x02,
	0x00, 0x00,
}
