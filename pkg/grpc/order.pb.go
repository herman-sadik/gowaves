// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Order_Side int32

const (
	Order_BUY  Order_Side = 0
	Order_SELL Order_Side = 1
)

var Order_Side_name = map[int32]string{
	0: "BUY",
	1: "SELL",
}

var Order_Side_value = map[string]int32{
	"BUY":  0,
	"SELL": 1,
}

func (x Order_Side) String() string {
	return proto.EnumName(Order_Side_name, int32(x))
}

func (Order_Side) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1, 0}
}

type AssetPair struct {
	AmountAssetId        []byte   `protobuf:"bytes,1,opt,name=amount_asset_id,json=amountAssetId,proto3" json:"amount_asset_id,omitempty"`
	PriceAssetId         []byte   `protobuf:"bytes,2,opt,name=price_asset_id,json=priceAssetId,proto3" json:"price_asset_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetPair) Reset()         { *m = AssetPair{} }
func (m *AssetPair) String() string { return proto.CompactTextString(m) }
func (*AssetPair) ProtoMessage()    {}
func (*AssetPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *AssetPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetPair.Unmarshal(m, b)
}
func (m *AssetPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetPair.Marshal(b, m, deterministic)
}
func (m *AssetPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetPair.Merge(m, src)
}
func (m *AssetPair) XXX_Size() int {
	return xxx_messageInfo_AssetPair.Size(m)
}
func (m *AssetPair) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetPair.DiscardUnknown(m)
}

var xxx_messageInfo_AssetPair proto.InternalMessageInfo

func (m *AssetPair) GetAmountAssetId() []byte {
	if m != nil {
		return m.AmountAssetId
	}
	return nil
}

func (m *AssetPair) GetPriceAssetId() []byte {
	if m != nil {
		return m.PriceAssetId
	}
	return nil
}

type Order struct {
	ChainId              int32      `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	SenderPublicKey      []byte     `protobuf:"bytes,2,opt,name=sender_public_key,json=senderPublicKey,proto3" json:"sender_public_key,omitempty"`
	MatcherPublicKey     []byte     `protobuf:"bytes,3,opt,name=matcher_public_key,json=matcherPublicKey,proto3" json:"matcher_public_key,omitempty"`
	AssetPair            *AssetPair `protobuf:"bytes,4,opt,name=asset_pair,json=assetPair,proto3" json:"asset_pair,omitempty"`
	OrderSide            Order_Side `protobuf:"varint,5,opt,name=order_side,json=orderSide,proto3,enum=waves.Order_Side" json:"order_side,omitempty"`
	Amount               int64      `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Price                int64      `protobuf:"varint,7,opt,name=price,proto3" json:"price,omitempty"`
	Timestamp            int64      `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Expiration           int64      `protobuf:"varint,9,opt,name=expiration,proto3" json:"expiration,omitempty"`
	MatcherFee           *Amount    `protobuf:"bytes,10,opt,name=matcher_fee,json=matcherFee,proto3" json:"matcher_fee,omitempty"`
	Version              int32      `protobuf:"varint,11,opt,name=version,proto3" json:"version,omitempty"`
	Proofs               [][]byte   `protobuf:"bytes,12,rep,name=proofs,proto3" json:"proofs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetChainId() int32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Order) GetSenderPublicKey() []byte {
	if m != nil {
		return m.SenderPublicKey
	}
	return nil
}

func (m *Order) GetMatcherPublicKey() []byte {
	if m != nil {
		return m.MatcherPublicKey
	}
	return nil
}

func (m *Order) GetAssetPair() *AssetPair {
	if m != nil {
		return m.AssetPair
	}
	return nil
}

func (m *Order) GetOrderSide() Order_Side {
	if m != nil {
		return m.OrderSide
	}
	return Order_BUY
}

func (m *Order) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Order) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Order) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *Order) GetMatcherFee() *Amount {
	if m != nil {
		return m.MatcherFee
	}
	return nil
}

func (m *Order) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Order) GetProofs() [][]byte {
	if m != nil {
		return m.Proofs
	}
	return nil
}

func init() {
	proto.RegisterEnum("waves.Order_Side", Order_Side_name, Order_Side_value)
	proto.RegisterType((*AssetPair)(nil), "waves.AssetPair")
	proto.RegisterType((*Order)(nil), "waves.Order")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xe1, 0x6a, 0xd4, 0x40,
	0x14, 0x85, 0x4d, 0xb3, 0xd9, 0xdd, 0xdc, 0x4d, 0xdb, 0xed, 0x20, 0x32, 0x15, 0x91, 0xb0, 0x88,
	0x04, 0x91, 0x58, 0xea, 0x13, 0x58, 0x50, 0x28, 0x16, 0x2c, 0x29, 0x22, 0xf5, 0x4f, 0x98, 0x4d,
	0x6e, 0xec, 0x60, 0x93, 0x19, 0x66, 0x66, 0xab, 0x7d, 0x25, 0x9f, 0xc4, 0xc7, 0x92, 0xdc, 0x99,
	0xed, 0xda, 0x7f, 0x7b, 0xcf, 0xf9, 0xf6, 0xde, 0x39, 0x87, 0xc0, 0x42, 0x99, 0x16, 0x4d, 0xa9,
	0x8d, 0x72, 0x8a, 0x25, 0xbf, 0xc4, 0x1d, 0xda, 0xe7, 0x99, 0xe8, 0xd5, 0x66, 0x70, 0x5e, 0x5c,
	0x5d, 0x43, 0xfa, 0xc1, 0x5a, 0x74, 0x97, 0x42, 0x1a, 0xf6, 0x1a, 0x0e, 0xbd, 0x59, 0x8b, 0x51,
	0xab, 0x65, 0xcb, 0xa3, 0x3c, 0x2a, 0xb2, 0x6a, 0xdf, 0xcb, 0x44, 0x9e, 0xb7, 0xec, 0x15, 0x1c,
	0x68, 0x23, 0x1b, 0xdc, 0x61, 0x7b, 0x84, 0x65, 0xa4, 0x06, 0x6a, 0xf5, 0x37, 0x86, 0xe4, 0xcb,
	0x78, 0x9f, 0x1d, 0xc3, 0xbc, 0xb9, 0x11, 0x72, 0xd8, 0x2e, 0x4c, 0xaa, 0x19, 0xcd, 0xe7, 0x2d,
	0x7b, 0x03, 0x47, 0x16, 0x87, 0x16, 0x4d, 0xad, 0x37, 0xeb, 0x5b, 0xd9, 0xd4, 0x3f, 0xf1, 0x3e,
	0x6c, 0x3b, 0xf4, 0xc6, 0x25, 0xe9, 0x9f, 0xf1, 0x9e, 0xbd, 0x05, 0xd6, 0x0b, 0xd7, 0xdc, 0x3c,
	0x86, 0x63, 0x82, 0x97, 0xc1, 0xd9, 0xd1, 0xef, 0x00, 0xfc, 0xf3, 0xb4, 0x90, 0x86, 0x4f, 0xf2,
	0xa8, 0x58, 0x9c, 0x2e, 0x4b, 0xea, 0xa0, 0x7c, 0x88, 0x5c, 0xa5, 0xe2, 0x21, 0xfd, 0x09, 0x00,
	0xd5, 0x55, 0x5b, 0xd9, 0x22, 0x4f, 0xf2, 0xa8, 0x38, 0x38, 0x3d, 0x0a, 0x7f, 0xa0, 0x1c, 0xe5,
	0x95, 0x6c, 0xb1, 0x4a, 0x09, 0x1a, 0x7f, 0xb2, 0x67, 0x30, 0xf5, 0xc5, 0xf0, 0x69, 0x1e, 0x15,
	0x71, 0x15, 0x26, 0xf6, 0x14, 0x12, 0x6a, 0x82, 0xcf, 0x48, 0xf6, 0x03, 0x7b, 0x01, 0xa9, 0x93,
	0x3d, 0x5a, 0x27, 0x7a, 0xcd, 0xe7, 0xe4, 0xec, 0x04, 0xf6, 0x12, 0x00, 0x7f, 0x6b, 0x69, 0x84,
	0x93, 0x6a, 0xe0, 0x29, 0xd9, 0xff, 0x29, 0xac, 0x84, 0xc5, 0x36, 0x7c, 0x87, 0xc8, 0x81, 0xf2,
	0xec, 0x6f, 0xf3, 0xd0, 0xdd, 0x0a, 0x02, 0xf1, 0x09, 0x91, 0x71, 0x98, 0xdd, 0xa1, 0xb1, 0xe3,
	0xb2, 0x85, 0xaf, 0x3c, 0x8c, 0xe3, 0xab, 0xb5, 0x51, 0xaa, 0xb3, 0x3c, 0xcb, 0xe3, 0x22, 0xab,
	0xc2, 0xb4, 0x3a, 0x86, 0x09, 0xa5, 0x9a, 0x41, 0x7c, 0xf6, 0xf5, 0x7a, 0xf9, 0x84, 0xcd, 0x61,
	0x72, 0xf5, 0xf1, 0xe2, 0x62, 0x19, 0x9d, 0x9d, 0x40, 0xde, 0xa8, 0xde, 0x1f, 0xd3, 0xb7, 0xc2,
	0x75, 0xca, 0xf4, 0xfe, 0x03, 0x5a, 0x6f, 0xba, 0x92, 0x0a, 0xf9, 0x3e, 0xf9, 0x61, 0x74, 0xf3,
	0x67, 0x2f, 0xf9, 0x36, 0x42, 0xeb, 0x29, 0xb9, 0xef, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x1e, 0x08, 0xc8, 0x82, 0x02, 0x00, 0x00,
}