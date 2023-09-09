// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eywa/eywa/handshake.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Handshake struct {
	Sender        string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver      string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	HashId        string `protobuf:"bytes,3,opt,name=hashId,proto3" json:"hashId,omitempty"`
	RoomId        string `protobuf:"bytes,4,opt,name=roomId,proto3" json:"roomId,omitempty"`
	ServerAddress string `protobuf:"bytes,5,opt,name=serverAddress,proto3" json:"serverAddress,omitempty"`
}

func (m *Handshake) Reset()         { *m = Handshake{} }
func (m *Handshake) String() string { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()    {}
func (*Handshake) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea6da9d4d8cdf1ab, []int{0}
}
func (m *Handshake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Handshake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Handshake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Handshake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handshake.Merge(m, src)
}
func (m *Handshake) XXX_Size() int {
	return m.Size()
}
func (m *Handshake) XXX_DiscardUnknown() {
	xxx_messageInfo_Handshake.DiscardUnknown(m)
}

var xxx_messageInfo_Handshake proto.InternalMessageInfo

func (m *Handshake) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Handshake) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *Handshake) GetHashId() string {
	if m != nil {
		return m.HashId
	}
	return ""
}

func (m *Handshake) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *Handshake) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Handshake)(nil), "github.com/eywa-foundation/eywa-go-client.eywa.Handshake")
}

func init() { proto.RegisterFile("github.com/eywa-foundation/eywa-go-client/eywa/handshake.proto", fileDescriptor_ea6da9d4d8cdf1ab) }

var fileDescriptor_ea6da9d4d8cdf1ab = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xad, 0x2c, 0x4f,
	0xd4, 0x07, 0x13, 0x19, 0x89, 0x79, 0x29, 0xc5, 0x19, 0x89, 0xd9, 0xa9, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x9c, 0x20, 0x51, 0x3d, 0x10, 0xa1, 0x34, 0x95, 0x91, 0x8b, 0xd3, 0x03, 0x26,
	0x2d, 0x24, 0xc6, 0xc5, 0x56, 0x9c, 0x9a, 0x97, 0x92, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0xe5, 0x09, 0x49, 0x71, 0x71, 0x14, 0xa5, 0x26, 0xa7, 0x66, 0x96, 0xa5, 0x16, 0x49,
	0x30, 0x81, 0x65, 0xe0, 0x7c, 0x90, 0x9e, 0x8c, 0xc4, 0xe2, 0x0c, 0xcf, 0x14, 0x09, 0x66, 0x88,
	0x1e, 0x08, 0x0f, 0x24, 0x5e, 0x94, 0x9f, 0x9f, 0xeb, 0x99, 0x22, 0xc1, 0x02, 0x11, 0x87, 0xf0,
	0x84, 0x54, 0xb8, 0x78, 0x8b, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0x1c, 0x53, 0x52, 0x8a, 0x52, 0x8b,
	0x8b, 0x25, 0x58, 0xc1, 0xd2, 0xa8, 0x82, 0x4e, 0xda, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0x25, 0x08, 0xf6, 0x52, 0x05, 0xc4, 0x67, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49,
	0x6c, 0x60, 0x6f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x55, 0xd2, 0xe1, 0xf3, 0x00,
	0x00, 0x00,
}

func (m *Handshake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Handshake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Handshake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServerAddress) > 0 {
		i -= len(m.ServerAddress)
		copy(dAtA[i:], m.ServerAddress)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.ServerAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RoomId) > 0 {
		i -= len(m.RoomId)
		copy(dAtA[i:], m.RoomId)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.RoomId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.HashId) > 0 {
		i -= len(m.HashId)
		copy(dAtA[i:], m.HashId)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.HashId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintHandshake(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHandshake(dAtA []byte, offset int, v uint64) int {
	offset -= sovHandshake(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Handshake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	l = len(m.HashId)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	l = len(m.RoomId)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	l = len(m.ServerAddress)
	if l > 0 {
		n += 1 + l + sovHandshake(uint64(l))
	}
	return n
}

func sovHandshake(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHandshake(x uint64) (n int) {
	return sovHandshake(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Handshake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHandshake
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
			return fmt.Errorf("proto: Handshake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Handshake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoomId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHandshake
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
				return ErrInvalidLengthHandshake
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHandshake
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHandshake(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHandshake
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
func skipHandshake(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHandshake
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
					return 0, ErrIntOverflowHandshake
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
					return 0, ErrIntOverflowHandshake
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
				return 0, ErrInvalidLengthHandshake
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHandshake
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHandshake
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHandshake        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHandshake          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHandshake = fmt.Errorf("proto: unexpected end of group")
)
