// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: faucet/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryWhenBrrRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryWhenBrrRequest) Reset()         { *m = QueryWhenBrrRequest{} }
func (m *QueryWhenBrrRequest) String() string { return proto.CompactTextString(m) }
func (*QueryWhenBrrRequest) ProtoMessage()    {}
func (*QueryWhenBrrRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32e01ab1e3e8ff22, []int{0}
}
func (m *QueryWhenBrrRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWhenBrrRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWhenBrrRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWhenBrrRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhenBrrRequest.Merge(m, src)
}
func (m *QueryWhenBrrRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryWhenBrrRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhenBrrRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhenBrrRequest proto.InternalMessageInfo

func (m *QueryWhenBrrRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryWhenBrrResponse struct {
	TimeLeft int64 `protobuf:"varint,1,opt,name=timeLeft,proto3" json:"timeLeft,omitempty"`
}

func (m *QueryWhenBrrResponse) Reset()         { *m = QueryWhenBrrResponse{} }
func (m *QueryWhenBrrResponse) String() string { return proto.CompactTextString(m) }
func (*QueryWhenBrrResponse) ProtoMessage()    {}
func (*QueryWhenBrrResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32e01ab1e3e8ff22, []int{1}
}
func (m *QueryWhenBrrResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryWhenBrrResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryWhenBrrResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryWhenBrrResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWhenBrrResponse.Merge(m, src)
}
func (m *QueryWhenBrrResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryWhenBrrResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWhenBrrResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWhenBrrResponse proto.InternalMessageInfo

func (m *QueryWhenBrrResponse) GetTimeLeft() int64 {
	if m != nil {
		return m.TimeLeft
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryWhenBrrRequest)(nil), "faucet.QueryWhenBrrRequest")
	proto.RegisterType((*QueryWhenBrrResponse)(nil), "faucet.QueryWhenBrrResponse")
}

func init() { proto.RegisterFile("faucet/query.proto", fileDescriptor_32e01ab1e3e8ff22) }

var fileDescriptor_32e01ab1e3e8ff22 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x1b, 0x10, 0x05, 0x2c, 0x26, 0xd3, 0xa1, 0x0d, 0x95, 0x85, 0x3a, 0x31, 0xe5, 0xa4,
	0xf2, 0x06, 0x9d, 0xbb, 0xd0, 0x05, 0x89, 0x05, 0x39, 0xe9, 0x35, 0xb5, 0xd4, 0xf8, 0x52, 0xdb,
	0x11, 0x44, 0x88, 0x85, 0x27, 0x40, 0xe2, 0xa5, 0x18, 0x2b, 0xb1, 0x30, 0xa2, 0x84, 0x07, 0x41,
	0x8d, 0x53, 0x04, 0x52, 0x37, 0xff, 0xff, 0xef, 0x4f, 0xf7, 0xdb, 0xc7, 0xf8, 0x42, 0x16, 0x09,
	0x3a, 0x58, 0x17, 0x68, 0xca, 0x28, 0x37, 0xe4, 0x88, 0x77, 0xbd, 0x17, 0x0e, 0x12, 0xb2, 0x19,
	0xd9, 0xfb, 0xc6, 0x05, 0x2f, 0xfc, 0x95, 0x70, 0x98, 0x12, 0xa5, 0x2b, 0x04, 0x99, 0x2b, 0x90,
	0x5a, 0x93, 0x93, 0x4e, 0x91, 0xde, 0xa5, 0xbd, 0x94, 0x52, 0xf2, 0xd4, 0xf6, 0xd4, 0xba, 0x83,
	0x96, 0x69, 0x54, 0x5c, 0x2c, 0x40, 0xea, 0x76, 0xe2, 0x08, 0xd8, 0xf9, 0xcd, 0xb6, 0xc0, 0xed,
	0x12, 0xf5, 0xc4, 0x98, 0x19, 0xae, 0x0b, 0xb4, 0x8e, 0xf7, 0xd9, 0xb1, 0x9c, 0xcf, 0x0d, 0x5a,
	0xdb, 0x0f, 0x2e, 0x83, 0xab, 0xd3, 0xd9, 0x4e, 0x8e, 0xc6, 0xac, 0xf7, 0x1f, 0xb0, 0x39, 0x69,
	0x8b, 0x3c, 0x64, 0x27, 0x4e, 0x65, 0x38, 0xc5, 0x85, 0x6b, 0x90, 0xc3, 0xd9, 0xaf, 0x1e, 0x6b,
	0x76, 0xd4, 0x30, 0x1c, 0xd9, 0xd9, 0x5f, 0x98, 0x5f, 0x44, 0xfe, 0xc1, 0xd1, 0x9e, 0x0e, 0xe1,
	0x70, 0x7f, 0xe8, 0xe7, 0x8d, 0xc2, 0x97, 0x8f, 0xef, 0xb7, 0x83, 0x1e, 0xe7, 0xf0, 0xe0, 0x13,
	0x78, 0x6a, 0x2b, 0x3e, 0x4f, 0xa6, 0xef, 0x95, 0x08, 0x36, 0x95, 0x08, 0xbe, 0x2a, 0x11, 0xbc,
	0xd6, 0xa2, 0xb3, 0xa9, 0x45, 0xe7, 0xb3, 0x16, 0x9d, 0xbb, 0x71, 0xaa, 0xdc, 0xb2, 0x88, 0xa3,
	0x84, 0x32, 0x50, 0xda, 0xa1, 0x49, 0x96, 0x52, 0xe9, 0x18, 0xcd, 0x4a, 0x69, 0xc8, 0x89, 0x56,
	0x8e, 0x4a, 0x78, 0x84, 0x76, 0x35, 0xae, 0xcc, 0xd1, 0xc6, 0xdd, 0xe6, 0xa7, 0xae, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x75, 0x5c, 0xbb, 0xa2, 0xb1, 0x01, 0x00, 0x00,
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
	// WhenBrr queries the last time an address minted
	QueryWhenBrr(ctx context.Context, in *QueryWhenBrrRequest, opts ...grpc.CallOption) (*QueryWhenBrrResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) QueryWhenBrr(ctx context.Context, in *QueryWhenBrrRequest, opts ...grpc.CallOption) (*QueryWhenBrrResponse, error) {
	out := new(QueryWhenBrrResponse)
	err := c.cc.Invoke(ctx, "/faucet.Query/QueryWhenBrr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// WhenBrr queries the last time an address minted
	QueryWhenBrr(context.Context, *QueryWhenBrrRequest) (*QueryWhenBrrResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) QueryWhenBrr(ctx context.Context, req *QueryWhenBrrRequest) (*QueryWhenBrrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryWhenBrr not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_QueryWhenBrr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWhenBrrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryWhenBrr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faucet.Query/QueryWhenBrr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryWhenBrr(ctx, req.(*QueryWhenBrrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "faucet.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryWhenBrr",
			Handler:    _Query_QueryWhenBrr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "faucet/query.proto",
}

func (m *QueryWhenBrrRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWhenBrrRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWhenBrrRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryWhenBrrResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryWhenBrrResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryWhenBrrResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeLeft != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.TimeLeft))
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
func (m *QueryWhenBrrRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryWhenBrrResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimeLeft != 0 {
		n += 1 + sovQuery(uint64(m.TimeLeft))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryWhenBrrRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryWhenBrrRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWhenBrrRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *QueryWhenBrrResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryWhenBrrResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryWhenBrrResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeLeft", wireType)
			}
			m.TimeLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeLeft |= int64(b&0x7F) << shift
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
