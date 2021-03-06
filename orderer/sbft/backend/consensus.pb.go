// Code generated by protoc-gen-gogo.
// source: backend/consensus.proto
// DO NOT EDIT!

/*
	Package backend is a generated protocol buffer package.

	It is generated from these files:
		backend/consensus.proto

	It has these top-level messages:
		Handshake
*/
package backend

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import simplebft "github.com/hyperledger/fabric/orderer/sbft/simplebft"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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

type Handshake struct {
}

func (m *Handshake) Reset()                    { *m = Handshake{} }
func (m *Handshake) String() string            { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()               {}
func (*Handshake) Descriptor() ([]byte, []int) { return fileDescriptorConsensus, []int{0} }

func init() {
	proto.RegisterType((*Handshake)(nil), "backend.handshake")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Consensus service

type ConsensusClient interface {
	Consensus(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (Consensus_ConsensusClient, error)
}

type consensusClient struct {
	cc *grpc.ClientConn
}

func NewConsensusClient(cc *grpc.ClientConn) ConsensusClient {
	return &consensusClient{cc}
}

func (c *consensusClient) Consensus(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (Consensus_ConsensusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Consensus_serviceDesc.Streams[0], c.cc, "/backend.consensus/consensus", opts...)
	if err != nil {
		return nil, err
	}
	x := &consensusConsensusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Consensus_ConsensusClient interface {
	Recv() (*simplebft.MultiChainMsg, error)
	grpc.ClientStream
}

type consensusConsensusClient struct {
	grpc.ClientStream
}

func (x *consensusConsensusClient) Recv() (*simplebft.MultiChainMsg, error) {
	m := new(simplebft.MultiChainMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Consensus service

type ConsensusServer interface {
	Consensus(*Handshake, Consensus_ConsensusServer) error
}

func RegisterConsensusServer(s *grpc.Server, srv ConsensusServer) {
	s.RegisterService(&_Consensus_serviceDesc, srv)
}

func _Consensus_Consensus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Handshake)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsensusServer).Consensus(m, &consensusConsensusServer{stream})
}

type Consensus_ConsensusServer interface {
	Send(*simplebft.MultiChainMsg) error
	grpc.ServerStream
}

type consensusConsensusServer struct {
	grpc.ServerStream
}

func (x *consensusConsensusServer) Send(m *simplebft.MultiChainMsg) error {
	return x.ServerStream.SendMsg(m)
}

var _Consensus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backend.consensus",
	HandlerType: (*ConsensusServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "consensus",
			Handler:       _Consensus_Consensus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "backend/consensus.proto",
}

func (m *Handshake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Handshake) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64Consensus(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Consensus(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintConsensus(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Handshake) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovConsensus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConsensus(x uint64) (n int) {
	return sovConsensus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Handshake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensus
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
			return fmt.Errorf("proto: handshake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: handshake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipConsensus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConsensus
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
func skipConsensus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConsensus
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
					return 0, ErrIntOverflowConsensus
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
					return 0, ErrIntOverflowConsensus
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
				return 0, ErrInvalidLengthConsensus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConsensus
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
				next, err := skipConsensus(dAtA[start:])
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
	ErrInvalidLengthConsensus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConsensus   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("backend/consensus.proto", fileDescriptorConsensus) }

var fileDescriptorConsensus = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x4a, 0x4c, 0xce,
	0x4e, 0xcd, 0x4b, 0xd1, 0x4f, 0xce, 0xcf, 0x2b, 0x4e, 0xcd, 0x2b, 0x2e, 0x2d, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x4a, 0x48, 0x49, 0x16, 0x67, 0xe6, 0x16, 0xe4, 0xa4, 0x26,
	0xa5, 0x95, 0xe8, 0xc3, 0x59, 0x10, 0x35, 0x4a, 0xdc, 0x5c, 0x9c, 0x19, 0x89, 0x79, 0x29, 0xc5,
	0x19, 0x89, 0xd9, 0xa9, 0x46, 0x5e, 0x5c, 0x9c, 0x70, 0x33, 0x84, 0x6c, 0x91, 0x39, 0x42, 0x7a,
	0x50, 0xb3, 0xf4, 0xe0, 0xaa, 0xa5, 0x24, 0xf4, 0x10, 0x86, 0xf9, 0x96, 0xe6, 0x94, 0x64, 0x3a,
	0x67, 0x24, 0x66, 0xe6, 0xf9, 0x16, 0xa7, 0x2b, 0x31, 0x18, 0x30, 0x3a, 0xb9, 0x9c, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0x19,
	0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x67, 0x54, 0x16, 0xa4, 0x16,
	0xe5, 0xa4, 0xa6, 0xa4, 0xa7, 0x16, 0xe9, 0xa7, 0x25, 0x26, 0x15, 0x65, 0x26, 0xeb, 0xe7, 0x17,
	0xa5, 0xa4, 0x16, 0xa5, 0x16, 0xe9, 0x17, 0x83, 0x9c, 0x09, 0xb5, 0x2d, 0x89, 0x0d, 0xec, 0x4a,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0x3f, 0x06, 0x09, 0xe4, 0x00, 0x00, 0x00,
}
