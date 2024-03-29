// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: halo/valsync/types/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ValidatorSetRequest struct {
	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Latest bool   `protobuf:"varint,2,opt,name=latest,proto3" json:"latest,omitempty"`
}

func (m *ValidatorSetRequest) Reset()         { *m = ValidatorSetRequest{} }
func (m *ValidatorSetRequest) String() string { return proto.CompactTextString(m) }
func (*ValidatorSetRequest) ProtoMessage()    {}
func (*ValidatorSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7457158ec95097a5, []int{0}
}
func (m *ValidatorSetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSetRequest.Merge(m, src)
}
func (m *ValidatorSetRequest) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSetRequest proto.InternalMessageInfo

func (m *ValidatorSetRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ValidatorSetRequest) GetLatest() bool {
	if m != nil {
		return m.Latest
	}
	return false
}

type ValidatorSetResponse struct {
	Id              uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedHeight   uint64       `protobuf:"varint,2,opt,name=created_height,json=createdHeight,proto3" json:"created_height,omitempty"`
	ActivatedHeight uint64       `protobuf:"varint,3,opt,name=activated_height,json=activatedHeight,proto3" json:"activated_height,omitempty"`
	Validators      []*Validator `protobuf:"bytes,4,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (m *ValidatorSetResponse) Reset()         { *m = ValidatorSetResponse{} }
func (m *ValidatorSetResponse) String() string { return proto.CompactTextString(m) }
func (*ValidatorSetResponse) ProtoMessage()    {}
func (*ValidatorSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7457158ec95097a5, []int{1}
}
func (m *ValidatorSetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSetResponse.Merge(m, src)
}
func (m *ValidatorSetResponse) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSetResponse proto.InternalMessageInfo

func (m *ValidatorSetResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ValidatorSetResponse) GetCreatedHeight() uint64 {
	if m != nil {
		return m.CreatedHeight
	}
	return 0
}

func (m *ValidatorSetResponse) GetActivatedHeight() uint64 {
	if m != nil {
		return m.ActivatedHeight
	}
	return 0
}

func (m *ValidatorSetResponse) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

type Validator struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Power   int64  `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_7457158ec95097a5, []int{2}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Validator) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func init() {
	proto.RegisterType((*ValidatorSetRequest)(nil), "halo.valsync.types.ValidatorSetRequest")
	proto.RegisterType((*ValidatorSetResponse)(nil), "halo.valsync.types.ValidatorSetResponse")
	proto.RegisterType((*Validator)(nil), "halo.valsync.types.Validator")
}

func init() { proto.RegisterFile("halo/valsync/types/query.proto", fileDescriptor_7457158ec95097a5) }

var fileDescriptor_7457158ec95097a5 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xcd, 0xf4, 0xef, 0xfb, 0xbc, 0xd6, 0x2a, 0x63, 0x91, 0x20, 0x38, 0x94, 0x80, 0x18, 0x41,
	0x52, 0xa8, 0x4b, 0xe9, 0xc6, 0x95, 0x5b, 0x47, 0x70, 0xe1, 0x46, 0xc6, 0xcc, 0xc5, 0x06, 0x42,
	0x27, 0x9d, 0x99, 0x56, 0xfa, 0x16, 0xbe, 0x8d, 0xaf, 0xe0, 0xb2, 0x4b, 0x97, 0xd2, 0xbe, 0x88,
	0x74, 0x9a, 0xc6, 0x48, 0x40, 0x97, 0xe7, 0x67, 0xee, 0x9c, 0x7b, 0x2e, 0xb0, 0x91, 0x48, 0x55,
	0x7f, 0x26, 0x52, 0x33, 0x1f, 0xc7, 0x7d, 0x3b, 0xcf, 0xd0, 0xf4, 0x27, 0x53, 0xd4, 0xf3, 0x28,
	0xd3, 0xca, 0x2a, 0x4a, 0xd7, 0x7a, 0x94, 0xeb, 0x91, 0xd3, 0x83, 0x21, 0x1c, 0xde, 0x8b, 0x34,
	0x91, 0xc2, 0x2a, 0x7d, 0x87, 0x96, 0xe3, 0x64, 0x8a, 0xc6, 0xd2, 0x0e, 0xd4, 0x12, 0xe9, 0x93,
	0x1e, 0x09, 0x1b, 0xbc, 0x96, 0x48, 0x7a, 0x04, 0xad, 0x54, 0x58, 0x34, 0xd6, 0xaf, 0xf5, 0x48,
	0xf8, 0x9f, 0xe7, 0x28, 0x78, 0x23, 0xd0, 0xfd, 0xf9, 0xde, 0x64, 0x6a, 0x6c, 0xb0, 0x32, 0xe0,
	0x14, 0x3a, 0xb1, 0x46, 0x61, 0x51, 0x3e, 0x8e, 0x30, 0x79, 0x1e, 0x6d, 0x06, 0x35, 0xf8, 0x5e,
	0xce, 0xde, 0x38, 0x92, 0x9e, 0xc3, 0x81, 0x88, 0x6d, 0x32, 0x2b, 0x1b, 0xeb, 0xce, 0xb8, 0x5f,
	0xf0, 0xb9, 0x75, 0x08, 0x30, 0xdb, 0xfe, 0x6c, 0xfc, 0x46, 0xaf, 0x1e, 0xee, 0x0e, 0x4e, 0xa2,
	0xea, 0x8a, 0x51, 0x91, 0x8f, 0x97, 0x1e, 0x04, 0x57, 0xb0, 0x53, 0x08, 0xd4, 0x87, 0x7f, 0x42,
	0x4a, 0x8d, 0xc6, 0xb8, 0xc8, 0x6d, 0xbe, 0x85, 0xb4, 0x0b, 0xcd, 0x4c, 0xbd, 0xa0, 0x76, 0x71,
	0xeb, 0x7c, 0x03, 0x06, 0x29, 0x34, 0x6f, 0xd7, 0xc5, 0xd2, 0x18, 0xda, 0xe5, 0xf5, 0xe9, 0xd9,
	0xaf, 0x01, 0xbe, 0x0b, 0x3e, 0x0e, 0xff, 0x36, 0x6e, 0x9a, 0x0c, 0xbc, 0xeb, 0x8b, 0xf7, 0x25,
	0x23, 0x8b, 0x25, 0x23, 0x9f, 0x4b, 0x46, 0x5e, 0x57, 0xcc, 0x5b, 0xac, 0x98, 0xf7, 0xb1, 0x62,
	0xde, 0x03, 0xad, 0x5e, 0xfc, 0xa9, 0xe5, 0x8e, 0x7d, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0xa7,
	0xcb, 0x03, 0x08, 0x0e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	ValidatorSet(ctx context.Context, in *ValidatorSetRequest, opts ...grpc.CallOption) (*ValidatorSetResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ValidatorSet(ctx context.Context, in *ValidatorSetRequest, opts ...grpc.CallOption) (*ValidatorSetResponse, error) {
	out := new(ValidatorSetResponse)
	err := c.cc.Invoke(ctx, "/halo.valsync.types.Query/ValidatorSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	ValidatorSet(context.Context, *ValidatorSetRequest) (*ValidatorSetResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ValidatorSet(ctx context.Context, req *ValidatorSetRequest) (*ValidatorSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorSet not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ValidatorSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatorSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halo.valsync.types.Query/ValidatorSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorSet(ctx, req.(*ValidatorSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "halo.valsync.types.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidatorSet",
			Handler:    _Query_ValidatorSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "halo/valsync/types/query.proto",
}

func (m *ValidatorSetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Latest {
		i--
		if m.Latest {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.ActivatedHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ActivatedHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.CreatedHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CreatedHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorSetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	if m.Latest {
		n += 2
	}
	return n
}

func (m *ValidatorSetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	if m.CreatedHeight != 0 {
		n += 1 + sovQuery(uint64(m.CreatedHeight))
	}
	if m.ActivatedHeight != 0 {
		n += 1 + sovQuery(uint64(m.ActivatedHeight))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovQuery(uint64(m.Power))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorSetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ValidatorSetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latest", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Latest = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ValidatorSetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ValidatorSetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedHeight", wireType)
			}
			m.CreatedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivatedHeight", wireType)
			}
			m.ActivatedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActivatedHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, &Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
