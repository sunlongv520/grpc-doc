// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Models.proto

package services

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ProdModel struct {
	ProdId               int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdName             string   `protobuf:"bytes,2,opt,name=prod_name,json=prodName,proto3" json:"prod_name,omitempty"`
	ProdPrice            float32  `protobuf:"fixed32,3,opt,name=prod_price,json=prodPrice,proto3" json:"prod_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdModel) Reset()         { *m = ProdModel{} }
func (m *ProdModel) String() string { return proto.CompactTextString(m) }
func (*ProdModel) ProtoMessage()    {}
func (*ProdModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{0}
}

func (m *ProdModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdModel.Unmarshal(m, b)
}
func (m *ProdModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdModel.Marshal(b, m, deterministic)
}
func (m *ProdModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdModel.Merge(m, src)
}
func (m *ProdModel) XXX_Size() int {
	return xxx_messageInfo_ProdModel.Size(m)
}
func (m *ProdModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdModel.DiscardUnknown(m)
}

var xxx_messageInfo_ProdModel proto.InternalMessageInfo

func (m *ProdModel) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

func (m *ProdModel) GetProdName() string {
	if m != nil {
		return m.ProdName
	}
	return ""
}

func (m *ProdModel) GetProdPrice() float32 {
	if m != nil {
		return m.ProdPrice
	}
	return 0
}

type OrderMain struct {
	OrderId              int32                `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderNo              string               `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	UserId               int32                `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderMoney           float32              `protobuf:"fixed32,4,opt,name=order_money,json=orderMoney,proto3" json:"order_money,omitempty"`
	OrderTime            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=order_time,json=orderTime,proto3" json:"order_time,omitempty"`
	OrderDetails         []*OrderDetail       `protobuf:"bytes,6,rep,name=order_details,json=orderDetails,proto3" json:"order_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderMain) Reset()         { *m = OrderMain{} }
func (m *OrderMain) String() string { return proto.CompactTextString(m) }
func (*OrderMain) ProtoMessage()    {}
func (*OrderMain) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{1}
}

func (m *OrderMain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderMain.Unmarshal(m, b)
}
func (m *OrderMain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderMain.Marshal(b, m, deterministic)
}
func (m *OrderMain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderMain.Merge(m, src)
}
func (m *OrderMain) XXX_Size() int {
	return xxx_messageInfo_OrderMain.Size(m)
}
func (m *OrderMain) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderMain.DiscardUnknown(m)
}

var xxx_messageInfo_OrderMain proto.InternalMessageInfo

func (m *OrderMain) GetOrderId() int32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderMain) GetOrderNo() string {
	if m != nil {
		return m.OrderNo
	}
	return ""
}

func (m *OrderMain) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *OrderMain) GetOrderMoney() float32 {
	if m != nil {
		return m.OrderMoney
	}
	return 0
}

func (m *OrderMain) GetOrderTime() *timestamp.Timestamp {
	if m != nil {
		return m.OrderTime
	}
	return nil
}

func (m *OrderMain) GetOrderDetails() []*OrderDetail {
	if m != nil {
		return m.OrderDetails
	}
	return nil
}

//子订单模型
type OrderDetail struct {
	DetailId             int32    `protobuf:"varint,1,opt,name=detail_id,json=detailId,proto3" json:"detail_id,omitempty"`
	OrderNo              string   `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	ProdId               int32    `protobuf:"varint,3,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdPrice            float32  `protobuf:"fixed32,4,opt,name=prod_price,json=prodPrice,proto3" json:"prod_price,omitempty"`
	ProdNum              int32    `protobuf:"varint,5,opt,name=prod_num,json=prodNum,proto3" json:"prod_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderDetail) Reset()         { *m = OrderDetail{} }
func (m *OrderDetail) String() string { return proto.CompactTextString(m) }
func (*OrderDetail) ProtoMessage()    {}
func (*OrderDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{2}
}

func (m *OrderDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderDetail.Unmarshal(m, b)
}
func (m *OrderDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderDetail.Marshal(b, m, deterministic)
}
func (m *OrderDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderDetail.Merge(m, src)
}
func (m *OrderDetail) XXX_Size() int {
	return xxx_messageInfo_OrderDetail.Size(m)
}
func (m *OrderDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderDetail.DiscardUnknown(m)
}

var xxx_messageInfo_OrderDetail proto.InternalMessageInfo

func (m *OrderDetail) GetDetailId() int32 {
	if m != nil {
		return m.DetailId
	}
	return 0
}

func (m *OrderDetail) GetOrderNo() string {
	if m != nil {
		return m.OrderNo
	}
	return ""
}

func (m *OrderDetail) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

func (m *OrderDetail) GetProdPrice() float32 {
	if m != nil {
		return m.ProdPrice
	}
	return 0
}

func (m *OrderDetail) GetProdNum() int32 {
	if m != nil {
		return m.ProdNum
	}
	return 0
}

//用户模型
type UserInfo struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserScore            int32    `protobuf:"varint,2,opt,name=user_score,json=userScore,proto3" json:"user_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{3}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfo) GetUserScore() int32 {
	if m != nil {
		return m.UserScore
	}
	return 0
}

func init() {
	proto.RegisterType((*ProdModel)(nil), "services.ProdModel")
	proto.RegisterType((*OrderMain)(nil), "services.OrderMain")
	proto.RegisterType((*OrderDetail)(nil), "services.OrderDetail")
	proto.RegisterType((*UserInfo)(nil), "services.UserInfo")
}

func init() { proto.RegisterFile("Models.proto", fileDescriptor_96b05f67b8e9f86a) }

var fileDescriptor_96b05f67b8e9f86a = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x9b, 0xa6, 0x69, 0x93, 0x9b, 0xfe, 0xff, 0x60, 0x09, 0x11, 0x8a, 0x2a, 0xa2, 0x48,
	0x48, 0x59, 0x48, 0xa5, 0x32, 0xc1, 0x82, 0x54, 0xb1, 0x74, 0x68, 0xa9, 0x0c, 0xcc, 0x25, 0xad,
	0xdd, 0x2a, 0x52, 0x12, 0x47, 0x76, 0x52, 0x89, 0xad, 0x8f, 0xc1, 0xb3, 0x30, 0xf1, 0x3a, 0xbc,
	0x05, 0xb2, 0x4d, 0x14, 0xc2, 0xc2, 0x78, 0xce, 0xb9, 0xbe, 0xb6, 0xbf, 0x03, 0xc3, 0x05, 0x23,
	0x34, 0x15, 0x51, 0xc1, 0x59, 0xc9, 0x90, 0x2d, 0x28, 0x3f, 0x24, 0x5b, 0x2a, 0x46, 0x17, 0x7b,
	0xc6, 0xf6, 0x29, 0x9d, 0x28, 0x7f, 0x53, 0xed, 0x26, 0x65, 0x92, 0x51, 0x51, 0xc6, 0x59, 0xa1,
	0x47, 0x47, 0xff, 0x0f, 0x71, 0x9a, 0x90, 0xb8, 0xa4, 0x5a, 0x07, 0x2f, 0xe0, 0xac, 0x38, 0x23,
	0x6a, 0x1d, 0x3a, 0x85, 0x41, 0xc1, 0x19, 0x59, 0x27, 0xc4, 0x33, 0x7c, 0x23, 0xb4, 0x70, 0x5f,
	0xca, 0x39, 0x41, 0xe7, 0xe0, 0xa8, 0x20, 0x8f, 0x33, 0xea, 0x75, 0x7d, 0x23, 0x74, 0xb0, 0x2d,
	0x8d, 0x65, 0x9c, 0x51, 0x34, 0x06, 0x50, 0x61, 0xc1, 0x93, 0x2d, 0xf5, 0x4c, 0xdf, 0x08, 0xbb,
	0x58, 0x8d, 0xaf, 0xa4, 0x11, 0x1c, 0xbb, 0xe0, 0x3c, 0x70, 0x42, 0xf9, 0x22, 0x4e, 0x72, 0x74,
	0x06, 0x36, 0x93, 0xa2, 0xb9, 0x63, 0xa0, 0xf4, 0x9c, 0x34, 0x51, 0xce, 0xbe, 0xef, 0xd0, 0xd1,
	0x92, 0xc9, 0x87, 0x55, 0x42, 0x1f, 0x32, 0xf5, 0xc3, 0xa4, 0x9c, 0x13, 0x74, 0x05, 0xae, 0x3e,
	0x93, 0xb1, 0x9c, 0xbe, 0x7a, 0x3d, 0x79, 0xf9, 0x6c, 0xf8, 0xfe, 0xf9, 0x61, 0x0e, 0xc0, 0xba,
	0xec, 0x74, 0x8e, 0x77, 0x18, 0xd4, 0xc0, 0x42, 0xe6, 0xe8, 0x06, 0xb4, 0x5a, 0x4b, 0x2c, 0x9e,
	0xe5, 0x1b, 0xa1, 0x3b, 0x1d, 0x45, 0x9a, 0x59, 0x54, 0x33, 0x8b, 0x9e, 0x6a, 0x66, 0xd8, 0x51,
	0xd3, 0x52, 0xa3, 0x5b, 0xf8, 0xa7, 0x8f, 0x12, 0x5a, 0xc6, 0x49, 0x2a, 0xbc, 0xbe, 0x6f, 0x86,
	0xee, 0xf4, 0x24, 0xaa, 0xd9, 0x47, 0xea, 0x93, 0xf7, 0x2a, 0xc5, 0x43, 0xd6, 0x08, 0x11, 0xbc,
	0x19, 0xe0, 0xfe, 0x48, 0x25, 0x4e, 0xbd, 0xa5, 0xa1, 0x60, 0x6b, 0xe3, 0x4f, 0x0c, 0x75, 0x3f,
	0x66, 0xab, 0x9f, 0x76, 0x05, 0xbd, 0x5f, 0x15, 0xc8, 0x95, 0xba, 0xbe, 0x2a, 0x53, 0x9f, 0xb6,
	0xb0, 0xda, 0xb3, 0xac, 0xb2, 0x60, 0x06, 0xf6, 0xb3, 0x44, 0x99, 0xef, 0x5a, 0x94, 0x8d, 0x16,
	0xe5, 0x31, 0x80, 0x0a, 0xc4, 0x96, 0x71, 0xdd, 0xbf, 0x85, 0x1d, 0xe9, 0x3c, 0x4a, 0x63, 0xd3,
	0x57, 0xe4, 0xae, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x25, 0x97, 0xdf, 0x44, 0x95, 0x02, 0x00,
	0x00,
}
