// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commitment_messages.proto

package messages

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CommitmentResponse struct {
	Error                string   `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitmentResponse) Reset()         { *m = CommitmentResponse{} }
func (m *CommitmentResponse) String() string { return proto.CompactTextString(m) }
func (*CommitmentResponse) ProtoMessage()    {}
func (*CommitmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02948687e8213eef, []int{0}
}

func (m *CommitmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitmentResponse.Unmarshal(m, b)
}
func (m *CommitmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitmentResponse.Marshal(b, m, deterministic)
}
func (m *CommitmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitmentResponse.Merge(m, src)
}
func (m *CommitmentResponse) XXX_Size() int {
	return xxx_messageInfo_CommitmentResponse.Size(m)
}
func (m *CommitmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommitmentResponse proto.InternalMessageInfo

func (m *CommitmentResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Commitment struct {
	IDF                  []byte   `protobuf:"bytes,1,opt,name=IDF,proto3" json:"IDF,omitempty"`
	Contract             []byte   `protobuf:"bytes,2,opt,name=Contract,proto3" json:"Contract,omitempty"`
	Wallet               string   `protobuf:"bytes,3,opt,name=Wallet,proto3" json:"Wallet,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Commitment) Reset()         { *m = Commitment{} }
func (m *Commitment) String() string { return proto.CompactTextString(m) }
func (*Commitment) ProtoMessage()    {}
func (*Commitment) Descriptor() ([]byte, []int) {
	return fileDescriptor_02948687e8213eef, []int{1}
}

func (m *Commitment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commitment.Unmarshal(m, b)
}
func (m *Commitment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commitment.Marshal(b, m, deterministic)
}
func (m *Commitment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commitment.Merge(m, src)
}
func (m *Commitment) XXX_Size() int {
	return xxx_messageInfo_Commitment.Size(m)
}
func (m *Commitment) XXX_DiscardUnknown() {
	xxx_messageInfo_Commitment.DiscardUnknown(m)
}

var xxx_messageInfo_Commitment proto.InternalMessageInfo

func (m *Commitment) GetIDF() []byte {
	if m != nil {
		return m.IDF
	}
	return nil
}

func (m *Commitment) GetContract() []byte {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *Commitment) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func (m *Commitment) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*CommitmentResponse)(nil), "messages.CommitmentResponse")
	proto.RegisterType((*Commitment)(nil), "messages.Commitment")
}

func init() {
	proto.RegisterFile("commitment_messages.proto", fileDescriptor_02948687e8213eef)
}

var fileDescriptor_02948687e8213eef = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xad, 0xd5, 0xa5, 0x1d, 0x3d, 0xc8, 0x50, 0x64, 0x2d, 0x05, 0x75, 0x4f, 0xe2, 0x21,
	0x05, 0xbd, 0x7b, 0xb0, 0x2a, 0x78, 0xf0, 0xb2, 0x82, 0x82, 0x17, 0x89, 0x65, 0x5c, 0x82, 0x9b,
	0x64, 0x99, 0x8c, 0x18, 0xff, 0xbd, 0x18, 0xdd, 0xcd, 0xa5, 0xb7, 0x7c, 0x8f, 0xef, 0x25, 0xbc,
	0xc0, 0xd1, 0xda, 0x5b, 0x6b, 0xc4, 0x92, 0x93, 0x57, 0x4b, 0x21, 0xe8, 0x86, 0x82, 0xea, 0xd8,
	0x8b, 0xc7, 0x49, 0xcf, 0xd5, 0x39, 0xe0, 0x6a, 0xd0, 0x6a, 0x0a, 0x9d, 0x77, 0x81, 0x70, 0x06,
	0xbb, 0xb7, 0xcc, 0x9e, 0xcb, 0xd1, 0xc9, 0xe8, 0x6c, 0x5a, 0xff, 0x41, 0xd5, 0x01, 0x64, 0x17,
	0x0f, 0x60, 0x7c, 0x7f, 0x73, 0x97, 0x8c, 0xfd, 0xfa, 0xf7, 0x88, 0x73, 0x98, 0xac, 0xbc, 0x13,
	0xd6, 0x6b, 0x29, 0xb7, 0x53, 0x3c, 0x30, 0x1e, 0x42, 0xf1, 0xac, 0xdb, 0x96, 0xa4, 0x1c, 0xa7,
	0x2b, 0xff, 0x09, 0x17, 0x30, 0x7d, 0x34, 0x8d, 0xd3, 0xf2, 0xc9, 0x54, 0xee, 0xa4, 0x52, 0x0e,
	0x2e, 0x1e, 0x60, 0x2f, 0xbf, 0x18, 0xf0, 0x0a, 0x8a, 0x27, 0x62, 0xf3, 0xfe, 0x8d, 0x33, 0x35,
	0x2c, 0xca, 0xc2, 0x7c, 0xb1, 0x29, 0xed, 0x47, 0x55, 0x5b, 0xd7, 0xa7, 0x2f, 0xc7, 0x8d, 0x11,
	0x15, 0xa3, 0x72, 0x24, 0x5f, 0x9e, 0x3f, 0x96, 0xd4, 0x9a, 0x18, 0x0d, 0x2f, 0xfb, 0xde, 0x5b,
	0x91, 0x3e, 0xe8, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xeb, 0x78, 0x79, 0x3d, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommitmentsClient is the client API for Commitments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommitmentsClient interface {
	Verify(ctx context.Context, in *Commitment, opts ...grpc.CallOption) (*CommitmentResponse, error)
}

type commitmentsClient struct {
	cc grpc.ClientConnInterface
}

func NewCommitmentsClient(cc grpc.ClientConnInterface) CommitmentsClient {
	return &commitmentsClient{cc}
}

func (c *commitmentsClient) Verify(ctx context.Context, in *Commitment, opts ...grpc.CallOption) (*CommitmentResponse, error) {
	out := new(CommitmentResponse)
	err := c.cc.Invoke(ctx, "/messages.Commitments/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommitmentsServer is the server API for Commitments service.
type CommitmentsServer interface {
	Verify(context.Context, *Commitment) (*CommitmentResponse, error)
}

// UnimplementedCommitmentsServer can be embedded to have forward compatible implementations.
type UnimplementedCommitmentsServer struct {
}

func (*UnimplementedCommitmentsServer) Verify(ctx context.Context, req *Commitment) (*CommitmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}

func RegisterCommitmentsServer(s *grpc.Server, srv CommitmentsServer) {
	s.RegisterService(&_Commitments_serviceDesc, srv)
}

func _Commitments_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Commitment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitmentsServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Commitments/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitmentsServer).Verify(ctx, req.(*Commitment))
	}
	return interceptor(ctx, in, info, handler)
}

var _Commitments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.Commitments",
	HandlerType: (*CommitmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _Commitments_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commitment_messages.proto",
}
