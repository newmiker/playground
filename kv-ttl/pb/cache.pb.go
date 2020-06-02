// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache.proto

package pb

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type T struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *T) Reset()         { *m = T{} }
func (m *T) String() string { return proto.CompactTextString(m) }
func (*T) ProtoMessage()    {}
func (*T) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{1}
}

func (m *T) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_T.Unmarshal(m, b)
}
func (m *T) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_T.Marshal(b, m, deterministic)
}
func (m *T) XXX_Merge(src proto.Message) {
	xxx_messageInfo_T.Merge(m, src)
}
func (m *T) XXX_Size() int {
	return xxx_messageInfo_T.Size(m)
}
func (m *T) XXX_DiscardUnknown() {
	xxx_messageInfo_T.DiscardUnknown(m)
}

var xxx_messageInfo_T proto.InternalMessageInfo

func (m *T) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AddRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *T       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{2}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AddRequest) GetValue() *T {
	if m != nil {
		return m.Value
	}
	return nil
}

type Ack struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{3}
}

func (m *Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack.Unmarshal(m, b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
}
func (m *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(m, src)
}
func (m *Ack) XXX_Size() int {
	return xxx_messageInfo_Ack.Size(m)
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func (m *Ack) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResponse struct {
	Value                *T       `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fca3b110c9bbf3a, []int{5}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetValue() *T {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *GetResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*T)(nil), "pb.T")
	proto.RegisterType((*AddRequest)(nil), "pb.AddRequest")
	proto.RegisterType((*Ack)(nil), "pb.Ack")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "pb.GetResponse")
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor_5fca3b110c9bbf3a) }

var fileDescriptor_5fca3b110c9bbf3a = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xcf, 0x4b, 0x84, 0x40,
	0x14, 0x76, 0x46, 0x5c, 0xdb, 0x27, 0x6c, 0xf1, 0x28, 0x28, 0x83, 0x45, 0xde, 0xc9, 0x93, 0xc4,
	0x76, 0xab, 0x93, 0x87, 0xf0, 0x6e, 0xfb, 0x0f, 0xf8, 0xe3, 0x51, 0x30, 0xd6, 0x4c, 0xeb, 0x18,
	0xec, 0x7f, 0x1f, 0x8e, 0x9a, 0x87, 0xe8, 0x36, 0x1f, 0xef, 0xfb, 0x35, 0x1f, 0x44, 0x4d, 0xd5,
	0xbc, 0x73, 0x66, 0x4e, 0xda, 0x6a, 0x94, 0xa6, 0xa6, 0x10, 0x82, 0x97, 0x0f, 0x63, 0xcf, 0x74,
	0x07, 0xe2, 0x88, 0xd7, 0x10, 0x7c, 0x57, 0xdd, 0xc0, 0xb7, 0x22, 0x11, 0xe9, 0xb6, 0x9c, 0x00,
	0x3d, 0x03, 0xe4, 0x6d, 0x5b, 0xf2, 0xd7, 0xc0, 0xbd, 0xc5, 0x2b, 0xf0, 0x15, 0x9f, 0x67, 0xc6,
	0xf8, 0xc4, 0xfb, 0x45, 0x25, 0x13, 0x91, 0x46, 0x87, 0x20, 0x33, 0x75, 0x76, 0x5c, 0xc4, 0x37,
	0xe0, 0xe7, 0x8d, 0xc2, 0x1d, 0x48, 0xad, 0x9c, 0xe8, 0xa2, 0x94, 0x5a, 0xd1, 0x1e, 0xa0, 0x60,
	0xfb, 0xaf, 0x27, 0x3d, 0x41, 0xe4, 0xee, 0xbd, 0xd1, 0x9f, 0x3d, 0xaf, 0x11, 0xe2, 0x6f, 0xc4,
	0xec, 0x2d, 0x17, 0xef, 0xc3, 0x00, 0xe1, 0xab, 0xd5, 0xa7, 0xea, 0x8d, 0x31, 0x01, 0x3f, 0x6f,
	0x5b, 0xdc, 0x8d, 0xfc, 0xf5, 0x0f, 0x71, 0xe8, 0x70, 0xa3, 0xc8, 0xc3, 0x14, 0xfc, 0x82, 0xed,
	0xc4, 0x58, 0x1b, 0xc5, 0x97, 0xbf, 0x78, 0x6a, 0x40, 0x1e, 0xee, 0x61, 0x53, 0xb0, 0xcd, 0xbb,
	0x0e, 0xb7, 0xe3, 0xd1, 0xcd, 0x16, 0x4f, 0x4d, 0xc8, 0x7b, 0x10, 0xf5, 0xc6, 0xad, 0xfa, 0xf8,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xce, 0x27, 0x01, 0x0a, 0x64, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*Ack, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Storage_GetAllClient, error)
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/pb.Storage/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/pb.Storage/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Storage_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Storage_serviceDesc.Streams[0], "/pb.Storage/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_GetAllClient interface {
	Recv() (*T, error)
	grpc.ClientStream
}

type storageGetAllClient struct {
	grpc.ClientStream
}

func (x *storageGetAllClient) Recv() (*T, error) {
	m := new(T)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServer is the server API for Storage service.
type StorageServer interface {
	Add(context.Context, *AddRequest) (*Ack, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(*Empty, Storage_GetAllServer) error
}

// UnimplementedStorageServer can be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (*UnimplementedStorageServer) Add(ctx context.Context, req *AddRequest) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedStorageServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedStorageServer) GetAll(req *Empty, srv Storage_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storage/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Storage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).GetAll(m, &storageGetAllServer{stream})
}

type Storage_GetAllServer interface {
	Send(*T) error
	grpc.ServerStream
}

type storageGetAllServer struct {
	grpc.ServerStream
}

func (x *storageGetAllServer) Send(m *T) error {
	return x.ServerStream.SendMsg(m)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Storage_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Storage_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _Storage_GetAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cache.proto",
}