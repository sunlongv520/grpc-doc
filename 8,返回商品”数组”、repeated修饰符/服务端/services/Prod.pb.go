// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Prod.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ProdRequest struct {
	ProdId               int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdRequest) Reset()         { *m = ProdRequest{} }
func (m *ProdRequest) String() string { return proto.CompactTextString(m) }
func (*ProdRequest) ProtoMessage()    {}
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b02cd6816510a0e, []int{0}
}

func (m *ProdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdRequest.Unmarshal(m, b)
}
func (m *ProdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdRequest.Marshal(b, m, deterministic)
}
func (m *ProdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdRequest.Merge(m, src)
}
func (m *ProdRequest) XXX_Size() int {
	return xxx_messageInfo_ProdRequest.Size(m)
}
func (m *ProdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProdRequest proto.InternalMessageInfo

func (m *ProdRequest) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

type ProdResponse struct {
	ProdStock            int32    `protobuf:"varint,1,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdResponse) Reset()         { *m = ProdResponse{} }
func (m *ProdResponse) String() string { return proto.CompactTextString(m) }
func (*ProdResponse) ProtoMessage()    {}
func (*ProdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b02cd6816510a0e, []int{1}
}

func (m *ProdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdResponse.Unmarshal(m, b)
}
func (m *ProdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdResponse.Marshal(b, m, deterministic)
}
func (m *ProdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdResponse.Merge(m, src)
}
func (m *ProdResponse) XXX_Size() int {
	return xxx_messageInfo_ProdResponse.Size(m)
}
func (m *ProdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdResponse proto.InternalMessageInfo

func (m *ProdResponse) GetProdStock() int32 {
	if m != nil {
		return m.ProdStock
	}
	return 0
}

type QuerySize struct {
	Size                 int32    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuerySize) Reset()         { *m = QuerySize{} }
func (m *QuerySize) String() string { return proto.CompactTextString(m) }
func (*QuerySize) ProtoMessage()    {}
func (*QuerySize) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b02cd6816510a0e, []int{2}
}

func (m *QuerySize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuerySize.Unmarshal(m, b)
}
func (m *QuerySize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuerySize.Marshal(b, m, deterministic)
}
func (m *QuerySize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySize.Merge(m, src)
}
func (m *QuerySize) XXX_Size() int {
	return xxx_messageInfo_QuerySize.Size(m)
}
func (m *QuerySize) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySize.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySize proto.InternalMessageInfo

func (m *QuerySize) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ProdResponseList struct {
	Prodres              []*ProdResponse `protobuf:"bytes,1,rep,name=prodres,proto3" json:"prodres,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProdResponseList) Reset()         { *m = ProdResponseList{} }
func (m *ProdResponseList) String() string { return proto.CompactTextString(m) }
func (*ProdResponseList) ProtoMessage()    {}
func (*ProdResponseList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b02cd6816510a0e, []int{3}
}

func (m *ProdResponseList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdResponseList.Unmarshal(m, b)
}
func (m *ProdResponseList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdResponseList.Marshal(b, m, deterministic)
}
func (m *ProdResponseList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdResponseList.Merge(m, src)
}
func (m *ProdResponseList) XXX_Size() int {
	return xxx_messageInfo_ProdResponseList.Size(m)
}
func (m *ProdResponseList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdResponseList.DiscardUnknown(m)
}

var xxx_messageInfo_ProdResponseList proto.InternalMessageInfo

func (m *ProdResponseList) GetProdres() []*ProdResponse {
	if m != nil {
		return m.Prodres
	}
	return nil
}

func init() {
	proto.RegisterType((*ProdRequest)(nil), "services.ProdRequest")
	proto.RegisterType((*ProdResponse)(nil), "services.ProdResponse")
	proto.RegisterType((*QuerySize)(nil), "services.QuerySize")
	proto.RegisterType((*ProdResponseList)(nil), "services.ProdResponseList")
}

func init() { proto.RegisterFile("Prod.proto", fileDescriptor_8b02cd6816510a0e) }

var fileDescriptor_8b02cd6816510a0e = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0a, 0x28, 0xca, 0x4f,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d,
	0x96, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb,
	0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x53, 0x52, 0xe3, 0xe2, 0x06, 0xe9,
	0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe7, 0x62, 0x2f, 0x28, 0xca, 0x4f, 0x89,
	0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0x03, 0x71, 0x3d, 0x53, 0x94, 0x74,
	0xb9, 0x78, 0x20, 0xea, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x64, 0xb9, 0xb8, 0xc0, 0x0a,
	0x8b, 0x4b, 0xf2, 0x93, 0xb3, 0xa1, 0x6a, 0x39, 0x41, 0x22, 0xc1, 0x20, 0x01, 0x25, 0x79, 0x2e,
	0xce, 0xc0, 0xd2, 0xd4, 0xa2, 0xca, 0xe0, 0xcc, 0xaa, 0x54, 0x21, 0x21, 0x2e, 0x96, 0xe2, 0xcc,
	0xaa, 0x54, 0xa8, 0x2a, 0x30, 0x5b, 0xc9, 0x85, 0x4b, 0x00, 0xd9, 0x3c, 0x9f, 0xcc, 0xe2, 0x12,
	0x21, 0x03, 0x88, 0xe5, 0x45, 0xa9, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x62, 0x7a,
	0x30, 0x5f, 0xe8, 0x21, 0x2b, 0x0e, 0x82, 0x29, 0x33, 0x5a, 0xc3, 0x08, 0x71, 0x7e, 0x30, 0x44,
	0x99, 0x50, 0x24, 0x17, 0x8f, 0x7b, 0x6a, 0x49, 0x00, 0xcc, 0x19, 0x42, 0xa2, 0xe8, 0x06, 0x80,
	0x7d, 0x29, 0x85, 0xc3, 0x5c, 0x25, 0xa9, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x89, 0x08, 0x09, 0xe9,
	0x97, 0x19, 0xea, 0x83, 0xec, 0xd0, 0xaf, 0x86, 0x86, 0x46, 0xad, 0x90, 0x13, 0x17, 0x2f, 0xb2,
	0xd1, 0xc5, 0x42, 0xc2, 0x08, 0x43, 0xe0, 0x5e, 0x95, 0x92, 0xc2, 0x6e, 0x32, 0xc8, 0x7b, 0x4a,
	0x0c, 0x49, 0x6c, 0xe0, 0x30, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xb6, 0x1d, 0xa2,
	0xa9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProdServiceClient interface {
	GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error)
	GetProdStocks(ctx context.Context, in *QuerySize, opts ...grpc.CallOption) (*ProdResponseList, error)
}

type prodServiceClient struct {
	cc *grpc.ClientConn
}

func NewProdServiceClient(cc *grpc.ClientConn) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error) {
	out := new(ProdResponse)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) GetProdStocks(ctx context.Context, in *QuerySize, opts ...grpc.CallOption) (*ProdResponseList, error) {
	out := new(ProdResponseList)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdServiceServer is the server API for ProdService service.
type ProdServiceServer interface {
	GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error)
	GetProdStocks(context.Context, *QuerySize) (*ProdResponseList, error)
}

// UnimplementedProdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (*UnimplementedProdServiceServer) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStock not implemented")
}
func (*UnimplementedProdServiceServer) GetProdStocks(ctx context.Context, req *QuerySize) (*ProdResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStocks not implemented")
}

func RegisterProdServiceServer(s *grpc.Server, srv ProdServiceServer) {
	s.RegisterService(&_ProdService_serviceDesc, srv)
}

func _ProdService_GetProdStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStock(ctx, req.(*ProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_GetProdStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStocks(ctx, req.(*QuerySize))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProdStock",
			Handler:    _ProdService_GetProdStock_Handler,
		},
		{
			MethodName: "GetProdStocks",
			Handler:    _ProdService_GetProdStocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Prod.proto",
}
