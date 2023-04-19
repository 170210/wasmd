// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/wasm/v1/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"

	types "github.com/Finschia/wasmd/x/wasm/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

var (
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState - genesis state of x/wasm
type GenesisState struct {
	Params    types.Params                 `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Codes     []types.Code                 `protobuf:"bytes,2,rep,name=codes,proto3" json:"codes,omitempty"`
	Contracts []types.Contract             `protobuf:"bytes,3,rep,name=contracts,proto3" json:"contracts,omitempty"`
	Sequences []types.Sequence             `protobuf:"bytes,4,rep,name=sequences,proto3" json:"sequences,omitempty"`
	GenMsgs   []types.GenesisState_GenMsgs `protobuf:"bytes,5,rep,name=gen_msgs,json=genMsgs,proto3" json:"gen_msgs,omitempty"`
	// InactiveContractAddresses is a list of contract address that set inactive
	InactiveContractAddresses []string `protobuf:"bytes,6,rep,name=inactive_contract_addresses,json=inactiveContractAddresses,proto3" json:"inactive_contract_address,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3308f670fed712dc, []int{0}
}

func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}

func (m *GenesisState) XXX_Size() int {
	return m.Size()
}

func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() types.Params {
	if m != nil {
		return m.Params
	}
	return types.Params{}
}

func (m *GenesisState) GetCodes() []types.Code {
	if m != nil {
		return m.Codes
	}
	return nil
}

func (m *GenesisState) GetContracts() []types.Contract {
	if m != nil {
		return m.Contracts
	}
	return nil
}

func (m *GenesisState) GetSequences() []types.Sequence {
	if m != nil {
		return m.Sequences
	}
	return nil
}

func (m *GenesisState) GetGenMsgs() []types.GenesisState_GenMsgs {
	if m != nil {
		return m.GenMsgs
	}
	return nil
}

func (m *GenesisState) GetInactiveContractAddresses() []string {
	if m != nil {
		return m.InactiveContractAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lbm.wasm.v1.GenesisState")
}

func init() { proto.RegisterFile("lbm/wasm/v1/genesis.proto", fileDescriptor_3308f670fed712dc) }

var fileDescriptor_3308f670fed712dc = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x8b, 0xda, 0x40,
	0x14, 0xc7, 0x93, 0xfa, 0xa3, 0x35, 0x16, 0x0a, 0x69, 0x69, 0x63, 0x2c, 0x51, 0x5a, 0x68, 0xa5,
	0x94, 0x04, 0x2d, 0xf4, 0xde, 0xf4, 0xd7, 0xa9, 0x50, 0x94, 0xbd, 0x2c, 0x2c, 0x32, 0x99, 0x0c,
	0x63, 0xc0, 0xc9, 0x64, 0xf3, 0x46, 0x77, 0x3d, 0xef, 0x3f, 0xb0, 0x7f, 0x96, 0x47, 0x8f, 0x7b,
	0x92, 0x45, 0x6f, 0xfe, 0x15, 0x4b, 0x26, 0x19, 0xcd, 0xe2, 0x7a, 0x4a, 0xf2, 0xde, 0xe7, 0xfb,
	0x79, 0x99, 0xe1, 0x19, 0xad, 0x69, 0xc0, 0xbc, 0x2b, 0x04, 0xcc, 0x9b, 0xf7, 0x3d, 0x4a, 0x62,
	0x02, 0x11, 0xb8, 0x49, 0xca, 0x05, 0x37, 0x9b, 0xd3, 0x80, 0xb9, 0x59, 0xcb, 0x9d, 0xf7, 0xed,
	0x37, 0x94, 0x53, 0x2e, 0xeb, 0x5e, 0xf6, 0x96, 0x23, 0xf6, 0x7b, 0xcc, 0x81, 0xc9, 0xb4, 0x52,
	0x88, 0x45, 0x42, 0x0a, 0x81, 0xed, 0x1c, 0x75, 0x1f, 0x0d, 0xf8, 0x70, 0x53, 0x35, 0x5e, 0xfe,
	0xcd, 0x2b, 0x23, 0x81, 0x04, 0x31, 0xbf, 0x1b, 0xf5, 0x04, 0xa5, 0x88, 0x81, 0xa5, 0x77, 0xf5,
	0x5e, 0x73, 0x60, 0xb9, 0xca, 0xa0, 0xfe, 0xc3, 0xfd, 0x2f, 0xfb, 0x7e, 0x75, 0xb9, 0xee, 0x68,
	0xc3, 0x82, 0x36, 0x7f, 0x1b, 0x35, 0xcc, 0x43, 0x02, 0xd6, 0xb3, 0x6e, 0xa5, 0xd7, 0x1c, 0xbc,
	0x3d, 0x8e, 0xfd, 0xe4, 0x21, 0xf1, 0xdf, 0x65, 0xa1, 0xdd, 0xba, 0xf3, 0x4a, 0xc2, 0x5f, 0x39,
	0x8b, 0x04, 0x61, 0x89, 0x58, 0x0c, 0xf3, 0xb4, 0x79, 0x66, 0x34, 0x30, 0x8f, 0x45, 0x8a, 0xb0,
	0x00, 0xab, 0x22, 0x55, 0xf6, 0x53, 0xaa, 0x1c, 0xf1, 0xdb, 0x85, 0xee, 0xf5, 0x3e, 0x54, 0x52,
	0x1e, 0x4c, 0x99, 0x16, 0xc8, 0xe5, 0x8c, 0xc4, 0x98, 0x80, 0x55, 0x3d, 0xa5, 0x1d, 0x15, 0xc8,
	0x41, 0xbb, 0x0f, 0x95, 0xb5, 0xfb, 0xa2, 0x79, 0x61, 0xbc, 0xa0, 0x24, 0x1e, 0x33, 0xa0, 0x60,
	0xd5, 0xa4, 0xf5, 0xd3, 0xb1, 0xb5, 0x7c, 0xbd, 0xd9, 0xc7, 0x3f, 0xa0, 0xe0, 0xdb, 0xc5, 0x04,
	0x53, 0xe5, 0x4b, 0x03, 0x9e, 0xd3, 0x1c, 0x32, 0xa9, 0xd1, 0x8e, 0x62, 0x84, 0x45, 0x34, 0x27,
	0x63, 0x75, 0x96, 0x31, 0x0a, 0xc3, 0x94, 0x00, 0x10, 0xb0, 0xea, 0xdd, 0x4a, 0xaf, 0xe1, 0x7f,
	0xde, 0xad, 0x3b, 0x1f, 0x4f, 0x62, 0x25, 0x6d, 0x4b, 0x41, 0xea, 0xf6, 0x7e, 0x28, 0x93, 0xff,
	0x6b, 0xb9, 0x71, 0xf4, 0xd5, 0xc6, 0xd1, 0xef, 0x37, 0x8e, 0x7e, 0xbb, 0x75, 0xb4, 0xd5, 0xd6,
	0xd1, 0xee, 0xb6, 0x8e, 0x76, 0xfe, 0x85, 0x46, 0x62, 0x32, 0x0b, 0x5c, 0xcc, 0x99, 0xf7, 0x27,
	0x8a, 0x01, 0x4f, 0x22, 0x24, 0x57, 0x29, 0xf4, 0xae, 0xe5, 0x33, 0x99, 0xce, 0x20, 0xdf, 0xb8,
	0xa0, 0x2e, 0x57, 0xea, 0xdb, 0x43, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x26, 0x39, 0xc4, 0xd0,
	0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InactiveContractAddresses) > 0 {
		for iNdEx := len(m.InactiveContractAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.InactiveContractAddresses[iNdEx])
			copy(dAtA[i:], m.InactiveContractAddresses[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.InactiveContractAddresses[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.GenMsgs) > 0 {
		for iNdEx := len(m.GenMsgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenMsgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Sequences) > 0 {
		for iNdEx := len(m.Sequences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Sequences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Contracts) > 0 {
		for iNdEx := len(m.Contracts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Contracts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Codes) > 0 {
		for iNdEx := len(m.Codes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Codes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Codes) > 0 {
		for _, e := range m.Codes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Contracts) > 0 {
		for _, e := range m.Contracts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Sequences) > 0 {
		for _, e := range m.Sequences {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GenMsgs) > 0 {
		for _, e := range m.GenMsgs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.InactiveContractAddresses) > 0 {
		for _, s := range m.InactiveContractAddresses {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codes = append(m.Codes, types.Code{})
			if err := m.Codes[len(m.Codes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contracts = append(m.Contracts, types.Contract{})
			if err := m.Contracts[len(m.Contracts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sequences = append(m.Sequences, types.Sequence{})
			if err := m.Sequences[len(m.Sequences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenMsgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenMsgs = append(m.GenMsgs, types.GenesisState_GenMsgs{})
			if err := m.GenMsgs[len(m.GenMsgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveContractAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InactiveContractAddresses = append(m.InactiveContractAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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

func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)