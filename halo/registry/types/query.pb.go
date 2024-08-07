// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: halo/registry/types/query.proto

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

type NetworkRequest struct {
	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Latest bool   `protobuf:"varint,2,opt,name=latest,proto3" json:"latest,omitempty"`
}

func (m *NetworkRequest) Reset()         { *m = NetworkRequest{} }
func (m *NetworkRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkRequest) ProtoMessage()    {}
func (*NetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a7e0fe95c853e3, []int{0}
}
func (m *NetworkRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkRequest.Merge(m, src)
}
func (m *NetworkRequest) XXX_Size() int {
	return m.Size()
}
func (m *NetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkRequest proto.InternalMessageInfo

func (m *NetworkRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NetworkRequest) GetLatest() bool {
	if m != nil {
		return m.Latest
	}
	return false
}

type NetworkResponse struct {
	Id            uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedHeight uint64    `protobuf:"varint,2,opt,name=created_height,json=createdHeight,proto3" json:"created_height,omitempty"`
	Portals       []*Portal `protobuf:"bytes,3,rep,name=portals,proto3" json:"portals,omitempty"`
}

func (m *NetworkResponse) Reset()         { *m = NetworkResponse{} }
func (m *NetworkResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkResponse) ProtoMessage()    {}
func (*NetworkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a7e0fe95c853e3, []int{1}
}
func (m *NetworkResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkResponse.Merge(m, src)
}
func (m *NetworkResponse) XXX_Size() int {
	return m.Size()
}
func (m *NetworkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkResponse proto.InternalMessageInfo

func (m *NetworkResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NetworkResponse) GetCreatedHeight() uint64 {
	if m != nil {
		return m.CreatedHeight
	}
	return 0
}

func (m *NetworkResponse) GetPortals() []*Portal {
	if m != nil {
		return m.Portals
	}
	return nil
}

type Portal struct {
	ChainId        uint64   `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Address        []byte   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	DeployHeight   uint64   `protobuf:"varint,3,opt,name=deploy_height,json=deployHeight,proto3" json:"deploy_height,omitempty"`
	ShardIds       []uint64 `protobuf:"varint,4,rep,packed,name=shard_ids,json=shardIds,proto3" json:"shard_ids,omitempty"`
	AttestInterval uint64   `protobuf:"varint,5,opt,name=attest_interval,json=attestInterval,proto3" json:"attest_interval,omitempty"`
	BlockPeriodMs  uint64   `protobuf:"varint,6,opt,name=block_period_ms,json=blockPeriodMs,proto3" json:"block_period_ms,omitempty"`
	Name           string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Portal) Reset()         { *m = Portal{} }
func (m *Portal) String() string { return proto.CompactTextString(m) }
func (*Portal) ProtoMessage()    {}
func (*Portal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a7e0fe95c853e3, []int{2}
}
func (m *Portal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Portal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Portal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Portal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Portal.Merge(m, src)
}
func (m *Portal) XXX_Size() int {
	return m.Size()
}
func (m *Portal) XXX_DiscardUnknown() {
	xxx_messageInfo_Portal.DiscardUnknown(m)
}

var xxx_messageInfo_Portal proto.InternalMessageInfo

func (m *Portal) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Portal) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Portal) GetDeployHeight() uint64 {
	if m != nil {
		return m.DeployHeight
	}
	return 0
}

func (m *Portal) GetShardIds() []uint64 {
	if m != nil {
		return m.ShardIds
	}
	return nil
}

func (m *Portal) GetAttestInterval() uint64 {
	if m != nil {
		return m.AttestInterval
	}
	return 0
}

func (m *Portal) GetBlockPeriodMs() uint64 {
	if m != nil {
		return m.BlockPeriodMs
	}
	return 0
}

func (m *Portal) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkRequest)(nil), "halo.registry.types.NetworkRequest")
	proto.RegisterType((*NetworkResponse)(nil), "halo.registry.types.NetworkResponse")
	proto.RegisterType((*Portal)(nil), "halo.registry.types.Portal")
}

func init() { proto.RegisterFile("halo/registry/types/query.proto", fileDescriptor_a5a7e0fe95c853e3) }

var fileDescriptor_a5a7e0fe95c853e3 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x89, 0x1b, 0xa7, 0x43, 0xeb, 0x48, 0x5b, 0x09, 0x2d, 0x54, 0x32, 0x96, 0xcb,
	0x1f, 0x5f, 0x70, 0xa4, 0x22, 0x24, 0xce, 0x9c, 0xc8, 0x01, 0x54, 0xf6, 0xc0, 0x81, 0x8b, 0xb5,
	0xcd, 0x8e, 0xea, 0x55, 0x5d, 0xaf, 0xbb, 0xbb, 0x05, 0xf9, 0xc4, 0x2b, 0xf0, 0x58, 0x1c, 0x7b,
	0xe4, 0x88, 0x12, 0x1e, 0x04, 0x65, 0x6d, 0x47, 0x42, 0x44, 0xdc, 0x76, 0x7e, 0xfa, 0x66, 0x67,
	0xf4, 0x7d, 0x03, 0x4f, 0x4a, 0x51, 0xe9, 0x85, 0xc1, 0x2b, 0x65, 0x9d, 0x69, 0x17, 0xae, 0x6d,
	0xd0, 0x2e, 0x6e, 0xef, 0xd0, 0xb4, 0x79, 0x63, 0xb4, 0xd3, 0xf4, 0x64, 0x2b, 0xc8, 0x07, 0x41,
	0xee, 0x05, 0xe9, 0x1b, 0x88, 0x3e, 0xa0, 0xfb, 0xaa, 0xcd, 0x35, 0xc7, 0xdb, 0x3b, 0xb4, 0x8e,
	0x46, 0x30, 0x56, 0x92, 0x91, 0x84, 0x64, 0x01, 0x1f, 0x2b, 0x49, 0x1f, 0xc2, 0xb4, 0x12, 0x0e,
	0xad, 0x63, 0xe3, 0x84, 0x64, 0x33, 0xde, 0x57, 0xe9, 0x37, 0x98, 0xef, 0x3a, 0x6d, 0xa3, 0x6b,
	0x8b, 0xff, 0xb4, 0x3e, 0x83, 0x68, 0x65, 0x50, 0x38, 0x94, 0x45, 0x89, 0xea, 0xaa, 0xec, 0xbe,
	0x08, 0xf8, 0x71, 0x4f, 0xdf, 0x79, 0x48, 0x5f, 0x43, 0xd8, 0x68, 0xe3, 0x44, 0x65, 0xd9, 0x24,
	0x99, 0x64, 0x0f, 0xce, 0x4f, 0xf3, 0x3d, 0xab, 0xe6, 0x17, 0x5e, 0xc3, 0x07, 0x6d, 0xfa, 0x9b,
	0xc0, 0xb4, 0x63, 0xf4, 0x11, 0xcc, 0x56, 0xa5, 0x50, 0x75, 0xb1, 0x1b, 0x1f, 0xfa, 0x7a, 0x29,
	0x29, 0x83, 0x50, 0x48, 0x69, 0xd0, 0x5a, 0x3f, 0xfc, 0x88, 0x0f, 0x25, 0x3d, 0x83, 0x63, 0x89,
	0x4d, 0xa5, 0xdb, 0x61, 0xb9, 0x89, 0xef, 0x3c, 0xea, 0x60, 0xbf, 0xdb, 0x29, 0x1c, 0xda, 0x52,
	0x18, 0x59, 0x28, 0x69, 0x59, 0x90, 0x4c, 0xb2, 0x80, 0xcf, 0x3c, 0x58, 0x4a, 0x4b, 0x5f, 0xc0,
	0x5c, 0xb8, 0xad, 0x19, 0x85, 0xaa, 0x1d, 0x9a, 0x2f, 0xa2, 0x62, 0x07, 0xfe, 0x8f, 0xa8, 0xc3,
	0xcb, 0x9e, 0xd2, 0xe7, 0x30, 0xbf, 0xac, 0xf4, 0xea, 0xba, 0x68, 0xd0, 0x28, 0x2d, 0x8b, 0x1b,
	0xcb, 0xa6, 0x9d, 0x13, 0x1e, 0x5f, 0x78, 0xfa, 0xde, 0x52, 0x0a, 0x41, 0x2d, 0x6e, 0x90, 0x85,
	0x09, 0xc9, 0x0e, 0xb9, 0x7f, 0x9f, 0x17, 0x70, 0xf0, 0x71, 0x9b, 0x22, 0xfd, 0x04, 0x61, 0x6f,
	0x38, 0x3d, 0xdb, 0x6b, 0xd0, 0xdf, 0x41, 0x3e, 0x7e, 0xfa, 0x7f, 0x51, 0x97, 0x59, 0x3a, 0x7a,
	0xfb, 0xf2, 0xc7, 0x3a, 0x26, 0xf7, 0xeb, 0x98, 0xfc, 0x5a, 0xc7, 0xe4, 0xfb, 0x26, 0x1e, 0xdd,
	0x6f, 0xe2, 0xd1, 0xcf, 0x4d, 0x3c, 0xfa, 0x7c, 0xb2, 0xe7, 0xa4, 0x2e, 0xa7, 0xfe, 0x9a, 0x5e,
	0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x16, 0xa2, 0x22, 0x73, 0x70, 0x02, 0x00, 0x00,
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
	Network(ctx context.Context, in *NetworkRequest, opts ...grpc.CallOption) (*NetworkResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Network(ctx context.Context, in *NetworkRequest, opts ...grpc.CallOption) (*NetworkResponse, error) {
	out := new(NetworkResponse)
	err := c.cc.Invoke(ctx, "/halo.registry.types.Query/Network", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Network(context.Context, *NetworkRequest) (*NetworkResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Network(ctx context.Context, req *NetworkRequest) (*NetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Network not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Network_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Network(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halo.registry.types.Query/Network",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Network(ctx, req.(*NetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "halo.registry.types.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Network",
			Handler:    _Query_Network_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "halo/registry/types/query.proto",
}

func (m *NetworkRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *NetworkResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetworkResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Portals) > 0 {
		for iNdEx := len(m.Portals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Portals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
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

func (m *Portal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Portal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Portal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x3a
	}
	if m.BlockPeriodMs != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BlockPeriodMs))
		i--
		dAtA[i] = 0x30
	}
	if m.AttestInterval != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.AttestInterval))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ShardIds) > 0 {
		dAtA2 := make([]byte, len(m.ShardIds)*10)
		var j1 int
		for _, num := range m.ShardIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintQuery(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if m.DeployHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.DeployHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChainId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x8
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
func (m *NetworkRequest) Size() (n int) {
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

func (m *NetworkResponse) Size() (n int) {
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
	if len(m.Portals) > 0 {
		for _, e := range m.Portals {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *Portal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainId != 0 {
		n += 1 + sovQuery(uint64(m.ChainId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.DeployHeight != 0 {
		n += 1 + sovQuery(uint64(m.DeployHeight))
	}
	if len(m.ShardIds) > 0 {
		l = 0
		for _, e := range m.ShardIds {
			l += sovQuery(uint64(e))
		}
		n += 1 + sovQuery(uint64(l)) + l
	}
	if m.AttestInterval != 0 {
		n += 1 + sovQuery(uint64(m.AttestInterval))
	}
	if m.BlockPeriodMs != 0 {
		n += 1 + sovQuery(uint64(m.BlockPeriodMs))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NetworkRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *NetworkResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NetworkResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Portals", wireType)
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
			m.Portals = append(m.Portals, &Portal{})
			if err := m.Portals[len(m.Portals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Portal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Portal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Portal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeployHeight", wireType)
			}
			m.DeployHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeployHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ShardIds = append(m.ShardIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthQuery
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthQuery
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ShardIds) == 0 {
					m.ShardIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ShardIds = append(m.ShardIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ShardIds", wireType)
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestInterval", wireType)
			}
			m.AttestInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AttestInterval |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockPeriodMs", wireType)
			}
			m.BlockPeriodMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockPeriodMs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
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
