// Code generated by protoc-gen-go. DO NOT EDIT.
// source: waves/node/grpc/blocks_api.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	waves "github.com/wavesplatform/gowaves/pkg/grpc/generated/waves"
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

type BlockRequest struct {
	// Types that are valid to be assigned to Request:
	//	*BlockRequest_BlockId
	//	*BlockRequest_Height
	//	*BlockRequest_Reference
	Request              isBlockRequest_Request `protobuf_oneof:"request"`
	IncludeTransactions  bool                   `protobuf:"varint,100,opt,name=include_transactions,json=includeTransactions,proto3" json:"include_transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BlockRequest) Reset()         { *m = BlockRequest{} }
func (m *BlockRequest) String() string { return proto.CompactTextString(m) }
func (*BlockRequest) ProtoMessage()    {}
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef9a5b9a94f607c7, []int{0}
}

func (m *BlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRequest.Unmarshal(m, b)
}
func (m *BlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRequest.Marshal(b, m, deterministic)
}
func (m *BlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRequest.Merge(m, src)
}
func (m *BlockRequest) XXX_Size() int {
	return xxx_messageInfo_BlockRequest.Size(m)
}
func (m *BlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRequest proto.InternalMessageInfo

type isBlockRequest_Request interface {
	isBlockRequest_Request()
}

type BlockRequest_BlockId struct {
	BlockId []byte `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3,oneof"`
}

type BlockRequest_Height struct {
	Height int32 `protobuf:"varint,2,opt,name=height,proto3,oneof"`
}

type BlockRequest_Reference struct {
	Reference []byte `protobuf:"bytes,3,opt,name=reference,proto3,oneof"`
}

func (*BlockRequest_BlockId) isBlockRequest_Request() {}

func (*BlockRequest_Height) isBlockRequest_Request() {}

func (*BlockRequest_Reference) isBlockRequest_Request() {}

func (m *BlockRequest) GetRequest() isBlockRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *BlockRequest) GetBlockId() []byte {
	if x, ok := m.GetRequest().(*BlockRequest_BlockId); ok {
		return x.BlockId
	}
	return nil
}

func (m *BlockRequest) GetHeight() int32 {
	if x, ok := m.GetRequest().(*BlockRequest_Height); ok {
		return x.Height
	}
	return 0
}

func (m *BlockRequest) GetReference() []byte {
	if x, ok := m.GetRequest().(*BlockRequest_Reference); ok {
		return x.Reference
	}
	return nil
}

func (m *BlockRequest) GetIncludeTransactions() bool {
	if m != nil {
		return m.IncludeTransactions
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BlockRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BlockRequest_BlockId)(nil),
		(*BlockRequest_Height)(nil),
		(*BlockRequest_Reference)(nil),
	}
}

type BlockRangeRequest struct {
	FromHeight uint32 `protobuf:"varint,1,opt,name=from_height,json=fromHeight,proto3" json:"from_height,omitempty"`
	ToHeight   uint32 `protobuf:"varint,2,opt,name=to_height,json=toHeight,proto3" json:"to_height,omitempty"`
	// Types that are valid to be assigned to Filter:
	//	*BlockRangeRequest_Generator
	Filter               isBlockRangeRequest_Filter `protobuf_oneof:"filter"`
	IncludeTransactions  bool                       `protobuf:"varint,100,opt,name=include_transactions,json=includeTransactions,proto3" json:"include_transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *BlockRangeRequest) Reset()         { *m = BlockRangeRequest{} }
func (m *BlockRangeRequest) String() string { return proto.CompactTextString(m) }
func (*BlockRangeRequest) ProtoMessage()    {}
func (*BlockRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef9a5b9a94f607c7, []int{1}
}

func (m *BlockRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRangeRequest.Unmarshal(m, b)
}
func (m *BlockRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRangeRequest.Marshal(b, m, deterministic)
}
func (m *BlockRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRangeRequest.Merge(m, src)
}
func (m *BlockRangeRequest) XXX_Size() int {
	return xxx_messageInfo_BlockRangeRequest.Size(m)
}
func (m *BlockRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRangeRequest proto.InternalMessageInfo

func (m *BlockRangeRequest) GetFromHeight() uint32 {
	if m != nil {
		return m.FromHeight
	}
	return 0
}

func (m *BlockRangeRequest) GetToHeight() uint32 {
	if m != nil {
		return m.ToHeight
	}
	return 0
}

type isBlockRangeRequest_Filter interface {
	isBlockRangeRequest_Filter()
}

type BlockRangeRequest_Generator struct {
	Generator []byte `protobuf:"bytes,3,opt,name=generator,proto3,oneof"`
}

func (*BlockRangeRequest_Generator) isBlockRangeRequest_Filter() {}

func (m *BlockRangeRequest) GetFilter() isBlockRangeRequest_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *BlockRangeRequest) GetGenerator() []byte {
	if x, ok := m.GetFilter().(*BlockRangeRequest_Generator); ok {
		return x.Generator
	}
	return nil
}

func (m *BlockRangeRequest) GetIncludeTransactions() bool {
	if m != nil {
		return m.IncludeTransactions
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BlockRangeRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BlockRangeRequest_Generator)(nil),
	}
}

type BlockWithHeight struct {
	Block                *waves.Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Height               uint32       `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BlockWithHeight) Reset()         { *m = BlockWithHeight{} }
func (m *BlockWithHeight) String() string { return proto.CompactTextString(m) }
func (*BlockWithHeight) ProtoMessage()    {}
func (*BlockWithHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef9a5b9a94f607c7, []int{2}
}

func (m *BlockWithHeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockWithHeight.Unmarshal(m, b)
}
func (m *BlockWithHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockWithHeight.Marshal(b, m, deterministic)
}
func (m *BlockWithHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockWithHeight.Merge(m, src)
}
func (m *BlockWithHeight) XXX_Size() int {
	return xxx_messageInfo_BlockWithHeight.Size(m)
}
func (m *BlockWithHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockWithHeight.DiscardUnknown(m)
}

var xxx_messageInfo_BlockWithHeight proto.InternalMessageInfo

func (m *BlockWithHeight) GetBlock() *waves.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BlockWithHeight) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockRequest)(nil), "waves.node.grpc.BlockRequest")
	proto.RegisterType((*BlockRangeRequest)(nil), "waves.node.grpc.BlockRangeRequest")
	proto.RegisterType((*BlockWithHeight)(nil), "waves.node.grpc.BlockWithHeight")
}

func init() { proto.RegisterFile("waves/node/grpc/blocks_api.proto", fileDescriptor_ef9a5b9a94f607c7) }

var fileDescriptor_ef9a5b9a94f607c7 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xeb, 0xa2, 0xa6, 0xc9, 0x34, 0x51, 0xe8, 0x82, 0x2a, 0xcb, 0x81, 0x62, 0xf9, 0x94,
	0xd3, 0x1a, 0xd2, 0x27, 0x20, 0x15, 0x8a, 0x2b, 0x04, 0x07, 0x0b, 0x88, 0xc4, 0xc5, 0x72, 0xec,
	0xb1, 0x63, 0xd5, 0xf6, 0x2e, 0xeb, 0x35, 0x15, 0xcf, 0xc3, 0x8d, 0x03, 0x2f, 0xc8, 0x05, 0xed,
	0xae, 0xd3, 0xba, 0x01, 0x84, 0xd4, 0xe3, 0xce, 0xff, 0xef, 0xcc, 0xb7, 0xb3, 0x33, 0xe0, 0xde,
	0xc4, 0x5f, 0xb1, 0xf1, 0x6b, 0x96, 0xa2, 0x9f, 0x0b, 0x9e, 0xf8, 0x9b, 0x92, 0x25, 0xd7, 0x4d,
	0x14, 0xf3, 0x82, 0x72, 0xc1, 0x24, 0x23, 0x53, 0xed, 0xa0, 0xca, 0x41, 0x95, 0xc3, 0x39, 0x35,
	0x57, 0xb4, 0xd1, 0x78, 0x9c, 0x59, 0xce, 0x58, 0x5e, 0xa2, 0xaf, 0x4f, 0x9b, 0x36, 0xf3, 0xb1,
	0xe2, 0xf2, 0x5b, 0x27, 0x9e, 0xef, 0x8b, 0x37, 0x22, 0xe6, 0x1c, 0x45, 0x63, 0x74, 0xef, 0xbb,
	0x05, 0xe3, 0xa5, 0x4a, 0x16, 0xe2, 0x97, 0x16, 0x1b, 0x49, 0x66, 0x30, 0xd4, 0xc9, 0xa3, 0x22,
	0xb5, 0x2d, 0xd7, 0x9a, 0x8f, 0x83, 0x83, 0xf0, 0x58, 0x47, 0xae, 0x52, 0x62, 0xc3, 0x60, 0x8b,
	0x45, 0xbe, 0x95, 0xf6, 0xa1, 0x6b, 0xcd, 0x8f, 0x82, 0x83, 0xb0, 0x3b, 0x93, 0x73, 0x18, 0x09,
	0xcc, 0x50, 0x60, 0x9d, 0xa0, 0xfd, 0xa8, 0xbb, 0x77, 0x17, 0x22, 0xaf, 0xe0, 0x69, 0x51, 0x27,
	0x65, 0x9b, 0x62, 0x24, 0x45, 0x5c, 0x37, 0x71, 0x22, 0x0b, 0x56, 0x37, 0x76, 0xea, 0x5a, 0xf3,
	0x61, 0xf8, 0xa4, 0xd3, 0x3e, 0xf4, 0xa4, 0xe5, 0x08, 0x8e, 0x85, 0x81, 0xf2, 0x7e, 0x5a, 0x70,
	0x6a, 0x28, 0xe3, 0x3a, 0xc7, 0x1d, 0xea, 0x0b, 0x38, 0xc9, 0x04, 0xab, 0xa2, 0x0e, 0x49, 0xd1,
	0x4e, 0x42, 0x50, 0xa1, 0xc0, 0x40, 0xcd, 0x60, 0x24, 0x59, 0xd4, 0x23, 0x9e, 0x84, 0x43, 0xc9,
	0x82, 0x5b, 0xe2, 0x1c, 0x6b, 0x14, 0xb1, 0x64, 0xe2, 0x8e, 0xf8, 0x36, 0xf4, 0x10, 0xe2, 0x21,
	0x0c, 0xb2, 0xa2, 0x94, 0x28, 0xbc, 0x77, 0x30, 0xd5, 0xbc, 0xeb, 0x42, 0x6e, 0xbb, 0x7a, 0x1e,
	0x1c, 0xe9, 0x36, 0x6a, 0xce, 0x93, 0xc5, 0x98, 0x9a, 0xaf, 0x35, 0xcf, 0x32, 0x12, 0x39, 0xbb,
	0xd7, 0xdf, 0xc9, 0xae, 0xbb, 0x8b, 0x5f, 0x16, 0x8c, 0xb4, 0xb1, 0x79, 0xcd, 0x0b, 0xf2, 0x16,
	0x86, 0x2b, 0x94, 0xfa, 0x4c, 0x9e, 0xd3, 0xbd, 0x09, 0xa1, 0xfd, 0xdf, 0x74, 0xdc, 0xbf, 0xcb,
	0x3d, 0xac, 0x35, 0x4c, 0x76, 0xc9, 0x74, 0x73, 0x89, 0xf7, 0x8f, 0x8c, 0xbd, 0xce, 0xff, 0x3f,
	0xed, 0x4b, 0x8b, 0x04, 0xf0, 0x78, 0x85, 0xf2, 0xb2, 0x15, 0x02, 0x6b, 0xd9, 0x15, 0x3b, 0xa3,
	0x66, 0x1c, 0xe9, 0x6e, 0x1c, 0xe9, 0x1b, 0x35, 0xab, 0xce, 0xb3, 0x3f, 0xe2, 0x1f, 0xaf, 0x6a,
	0x79, 0xb1, 0xf8, 0x14, 0x97, 0x2d, 0x2e, 0x1b, 0x70, 0x12, 0x56, 0x99, 0x92, 0xbc, 0x8c, 0x65,
	0xc6, 0x44, 0x45, 0xd5, 0x8e, 0xa8, 0xca, 0x9f, 0x2f, 0xf3, 0x42, 0x6e, 0xdb, 0x0d, 0x4d, 0x58,
	0xe5, 0xdf, 0xb3, 0xf8, 0x39, 0x33, 0xcb, 0xc2, 0xaf, 0x73, 0xb3, 0x5e, 0xdd, 0xe7, 0x62, 0xea,
	0xef, 0x2d, 0xde, 0x8f, 0xc3, 0xe9, 0x5a, 0x3f, 0xe8, 0xbd, 0x7a, 0xd0, 0x4a, 0xf0, 0x64, 0x33,
	0xd0, 0x28, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x35, 0xfb, 0x88, 0xfa, 0xa4, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlocksApiClient is the client API for BlocksApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlocksApiClient interface {
	GetBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockWithHeight, error)
	GetBlockRange(ctx context.Context, in *BlockRangeRequest, opts ...grpc.CallOption) (BlocksApi_GetBlockRangeClient, error)
	GetCurrentHeight(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.UInt32Value, error)
}

type blocksApiClient struct {
	cc *grpc.ClientConn
}

func NewBlocksApiClient(cc *grpc.ClientConn) BlocksApiClient {
	return &blocksApiClient{cc}
}

func (c *blocksApiClient) GetBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockWithHeight, error) {
	out := new(BlockWithHeight)
	err := c.cc.Invoke(ctx, "/waves.node.grpc.BlocksApi/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blocksApiClient) GetBlockRange(ctx context.Context, in *BlockRangeRequest, opts ...grpc.CallOption) (BlocksApi_GetBlockRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlocksApi_serviceDesc.Streams[0], "/waves.node.grpc.BlocksApi/GetBlockRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &blocksApiGetBlockRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlocksApi_GetBlockRangeClient interface {
	Recv() (*BlockWithHeight, error)
	grpc.ClientStream
}

type blocksApiGetBlockRangeClient struct {
	grpc.ClientStream
}

func (x *blocksApiGetBlockRangeClient) Recv() (*BlockWithHeight, error) {
	m := new(BlockWithHeight)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blocksApiClient) GetCurrentHeight(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*wrappers.UInt32Value, error) {
	out := new(wrappers.UInt32Value)
	err := c.cc.Invoke(ctx, "/waves.node.grpc.BlocksApi/GetCurrentHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlocksApiServer is the server API for BlocksApi service.
type BlocksApiServer interface {
	GetBlock(context.Context, *BlockRequest) (*BlockWithHeight, error)
	GetBlockRange(*BlockRangeRequest, BlocksApi_GetBlockRangeServer) error
	GetCurrentHeight(context.Context, *empty.Empty) (*wrappers.UInt32Value, error)
}

// UnimplementedBlocksApiServer can be embedded to have forward compatible implementations.
type UnimplementedBlocksApiServer struct {
}

func (*UnimplementedBlocksApiServer) GetBlock(ctx context.Context, req *BlockRequest) (*BlockWithHeight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (*UnimplementedBlocksApiServer) GetBlockRange(req *BlockRangeRequest, srv BlocksApi_GetBlockRangeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlockRange not implemented")
}
func (*UnimplementedBlocksApiServer) GetCurrentHeight(ctx context.Context, req *empty.Empty) (*wrappers.UInt32Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentHeight not implemented")
}

func RegisterBlocksApiServer(s *grpc.Server, srv BlocksApiServer) {
	s.RegisterService(&_BlocksApi_serviceDesc, srv)
}

func _BlocksApi_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksApiServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/waves.node.grpc.BlocksApi/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksApiServer).GetBlock(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlocksApi_GetBlockRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockRangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlocksApiServer).GetBlockRange(m, &blocksApiGetBlockRangeServer{stream})
}

type BlocksApi_GetBlockRangeServer interface {
	Send(*BlockWithHeight) error
	grpc.ServerStream
}

type blocksApiGetBlockRangeServer struct {
	grpc.ServerStream
}

func (x *blocksApiGetBlockRangeServer) Send(m *BlockWithHeight) error {
	return x.ServerStream.SendMsg(m)
}

func _BlocksApi_GetCurrentHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlocksApiServer).GetCurrentHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/waves.node.grpc.BlocksApi/GetCurrentHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlocksApiServer).GetCurrentHeight(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlocksApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "waves.node.grpc.BlocksApi",
	HandlerType: (*BlocksApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlock",
			Handler:    _BlocksApi_GetBlock_Handler,
		},
		{
			MethodName: "GetCurrentHeight",
			Handler:    _BlocksApi_GetCurrentHeight_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlockRange",
			Handler:       _BlocksApi_GetBlockRange_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "waves/node/grpc/blocks_api.proto",
}