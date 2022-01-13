// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services.proto

package server

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

type Nil struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nil) Reset()         { *m = Nil{} }
func (m *Nil) String() string { return proto.CompactTextString(m) }
func (*Nil) ProtoMessage()    {}
func (*Nil) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{0}
}

func (m *Nil) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nil.Unmarshal(m, b)
}
func (m *Nil) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nil.Marshal(b, m, deterministic)
}
func (m *Nil) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nil.Merge(m, src)
}
func (m *Nil) XXX_Size() int {
	return xxx_messageInfo_Nil.Size(m)
}
func (m *Nil) XXX_DiscardUnknown() {
	xxx_messageInfo_Nil.DiscardUnknown(m)
}

var xxx_messageInfo_Nil proto.InternalMessageInfo

type Key struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{1}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type KeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{2}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type OprationResult struct {
	ActiveIp             string   `protobuf:"bytes,4,opt,name=activeIp,proto3" json:"activeIp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OprationResult) Reset()         { *m = OprationResult{} }
func (m *OprationResult) String() string { return proto.CompactTextString(m) }
func (*OprationResult) ProtoMessage()    {}
func (*OprationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{3}
}

func (m *OprationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OprationResult.Unmarshal(m, b)
}
func (m *OprationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OprationResult.Marshal(b, m, deterministic)
}
func (m *OprationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OprationResult.Merge(m, src)
}
func (m *OprationResult) XXX_Size() int {
	return xxx_messageInfo_OprationResult.Size(m)
}
func (m *OprationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_OprationResult.DiscardUnknown(m)
}

var xxx_messageInfo_OprationResult proto.InternalMessageInfo

func (m *OprationResult) GetActiveIp() string {
	if m != nil {
		return m.ActiveIp
	}
	return ""
}

type Result struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	MissCache            bool     `protobuf:"varint,2,opt,name=missCache,proto3" json:"missCache,omitempty"`
	ActiveIp             string   `protobuf:"bytes,4,opt,name=activeIp,proto3" json:"activeIp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{4}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Result) GetMissCache() bool {
	if m != nil {
		return m.MissCache
	}
	return false
}

func (m *Result) GetActiveIp() string {
	if m != nil {
		return m.ActiveIp
	}
	return ""
}

type Data struct {
	FromIp               string   `protobuf:"bytes,1,opt,name=fromIp,proto3" json:"fromIp,omitempty"`
	FromPort             int32    `protobuf:"varint,2,opt,name=fromPort,proto3" json:"fromPort,omitempty"`
	Distribute           bool     `protobuf:"varint,3,opt,name=distribute,proto3" json:"distribute,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{5}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetFromIp() string {
	if m != nil {
		return m.FromIp
	}
	return ""
}

func (m *Data) GetFromPort() int32 {
	if m != nil {
		return m.FromPort
	}
	return 0
}

func (m *Data) GetDistribute() bool {
	if m != nil {
		return m.Distribute
	}
	return false
}

type DistKeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Data                 *Data    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistKeyValue) Reset()         { *m = DistKeyValue{} }
func (m *DistKeyValue) String() string { return proto.CompactTextString(m) }
func (*DistKeyValue) ProtoMessage()    {}
func (*DistKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e16ccb8c5307b32, []int{6}
}

func (m *DistKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistKeyValue.Unmarshal(m, b)
}
func (m *DistKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistKeyValue.Marshal(b, m, deterministic)
}
func (m *DistKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistKeyValue.Merge(m, src)
}
func (m *DistKeyValue) XXX_Size() int {
	return xxx_messageInfo_DistKeyValue.Size(m)
}
func (m *DistKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DistKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_DistKeyValue proto.InternalMessageInfo

func (m *DistKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DistKeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *DistKeyValue) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Nil)(nil), "server.Nil")
	proto.RegisterType((*Key)(nil), "server.Key")
	proto.RegisterType((*KeyValue)(nil), "server.KeyValue")
	proto.RegisterType((*OprationResult)(nil), "server.OprationResult")
	proto.RegisterType((*Result)(nil), "server.Result")
	proto.RegisterType((*Data)(nil), "server.Data")
	proto.RegisterType((*DistKeyValue)(nil), "server.DistKeyValue")
}

func init() { proto.RegisterFile("services.proto", fileDescriptor_8e16ccb8c5307b32) }

var fileDescriptor_8e16ccb8c5307b32 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xe1, 0xaa, 0xda, 0x40,
	0x10, 0x85, 0xd5, 0x98, 0x54, 0x47, 0x11, 0x19, 0xc4, 0x86, 0x50, 0x8a, 0x6c, 0x4b, 0x91, 0x52,
	0x02, 0xb5, 0x4f, 0x50, 0x14, 0x8a, 0x08, 0x56, 0x52, 0x5a, 0xa4, 0xff, 0xd6, 0x38, 0xc5, 0xc5,
	0xd8, 0x84, 0xdd, 0x4d, 0x20, 0xaf, 0x76, 0x9f, 0xee, 0x92, 0x8d, 0xc6, 0xe8, 0x05, 0xe1, 0xfe,
	0xdb, 0x33, 0x73, 0xf2, 0xed, 0xd9, 0x99, 0xc0, 0x40, 0x91, 0xcc, 0x44, 0x48, 0xca, 0x4f, 0x64,
	0xac, 0x63, 0x74, 0x0a, 0x4d, 0x92, 0xd9, 0x60, 0xad, 0x45, 0xc4, 0xde, 0x82, 0xb5, 0xa2, 0x1c,
	0x87, 0x60, 0x1d, 0x29, 0x77, 0x9b, 0x93, 0xe6, 0xb4, 0x1b, 0x14, 0x47, 0x36, 0x83, 0xce, 0x8a,
	0xf2, 0x3f, 0x3c, 0x4a, 0xe9, 0x65, 0x17, 0x47, 0x60, 0x67, 0x45, 0xcb, 0x6d, 0x99, 0x5a, 0x29,
	0xd8, 0x17, 0x18, 0xfc, 0x4c, 0x24, 0xd7, 0x22, 0xfe, 0x1f, 0x90, 0x4a, 0x23, 0x8d, 0x1e, 0x74,
	0x78, 0xa8, 0x45, 0x46, 0xcb, 0xc4, 0x6d, 0x1b, 0x6b, 0xa5, 0xd9, 0x16, 0x9c, 0xb3, 0xab, 0xa2,
	0x35, 0x6b, 0x34, 0x7c, 0x07, 0xdd, 0x93, 0x50, 0x6a, 0xce, 0xc3, 0x43, 0x79, 0x4f, 0x27, 0xb8,
	0x16, 0x1e, 0x92, 0xff, 0x42, 0x7b, 0xc1, 0x35, 0xc7, 0x31, 0x38, 0xff, 0x64, 0x7c, 0x5a, 0x26,
	0x67, 0xf0, 0x59, 0x15, 0xdf, 0x16, 0xa7, 0x4d, 0x2c, 0xb5, 0x01, 0xdb, 0x41, 0xa5, 0xf1, 0x3d,
	0xc0, 0x5e, 0x28, 0x2d, 0xc5, 0x2e, 0xd5, 0xe4, 0x5a, 0xe6, 0xda, 0x5a, 0x85, 0x6d, 0xa1, 0xbf,
	0x10, 0x4a, 0xbf, 0x76, 0x36, 0x38, 0x81, 0xf6, 0x9e, 0x6b, 0x6e, 0x88, 0xbd, 0x59, 0xdf, 0x2f,
	0xd7, 0xe0, 0x17, 0x39, 0x03, 0xd3, 0x99, 0x3d, 0xb5, 0xa0, 0x6f, 0xde, 0xf6, 0xab, 0xdc, 0x18,
	0x7e, 0x04, 0xeb, 0x07, 0x69, 0xec, 0x5d, 0xbc, 0x2b, 0xca, 0xbd, 0xc1, 0x45, 0x94, 0xa3, 0x63,
	0x0d, 0xfc, 0x0a, 0xd6, 0x26, 0xd5, 0x38, 0xac, 0xb9, 0x4c, 0x32, 0x6f, 0x7c, 0xa9, 0xdc, 0xee,
	0x84, 0x35, 0xd0, 0x07, 0x7b, 0x1e, 0x11, 0x97, 0x57, 0xf4, 0x5a, 0x44, 0x0f, 0xfc, 0x9f, 0xc1,
	0xfa, 0xad, 0x08, 0x47, 0x55, 0xe8, 0xda, 0x00, 0xbc, 0x3a, 0x83, 0x35, 0x70, 0x0a, 0x5d, 0xc3,
	0x2e, 0x3c, 0x78, 0xf3, 0xcc, 0x7b, 0xe7, 0x07, 0xb0, 0xe7, 0x07, 0x0a, 0x8f, 0xb7, 0x29, 0xee,
	0x4c, 0x9f, 0xe0, 0xcd, 0xf2, 0xf4, 0x3d, 0x12, 0x19, 0x3d, 0x84, 0xed, 0x1c, 0xf3, 0x77, 0x7f,
	0x7b, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xd9, 0x60, 0x1f, 0xef, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheServiceClient interface {
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Result, error)
	Put(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*OprationResult, error)
	Clear(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*OprationResult, error)
	Use(ctx context.Context, in *DistKeyValue, opts ...grpc.CallOption) (*Nil, error)
	ClearDist(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Nil, error)
	Check(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*Nil, error)
	ImAlive(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Nil, error)
}

type cacheServiceClient struct {
	cc *grpc.ClientConn
}

func NewCacheServiceClient(cc *grpc.ClientConn) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/server.CacheService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Put(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*OprationResult, error) {
	out := new(OprationResult)
	err := c.cc.Invoke(ctx, "/server.CacheService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Clear(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*OprationResult, error) {
	out := new(OprationResult)
	err := c.cc.Invoke(ctx, "/server.CacheService/Clear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Use(ctx context.Context, in *DistKeyValue, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := c.cc.Invoke(ctx, "/server.CacheService/Use", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) ClearDist(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := c.cc.Invoke(ctx, "/server.CacheService/ClearDist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Check(ctx context.Context, in *Nil, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := c.cc.Invoke(ctx, "/server.CacheService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) ImAlive(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Nil, error) {
	out := new(Nil)
	err := c.cc.Invoke(ctx, "/server.CacheService/ImAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
type CacheServiceServer interface {
	Get(context.Context, *Key) (*Result, error)
	Put(context.Context, *KeyValue) (*OprationResult, error)
	Clear(context.Context, *Nil) (*OprationResult, error)
	Use(context.Context, *DistKeyValue) (*Nil, error)
	ClearDist(context.Context, *Data) (*Nil, error)
	Check(context.Context, *Nil) (*Nil, error)
	ImAlive(context.Context, *Data) (*Nil, error)
}

// UnimplementedCacheServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (*UnimplementedCacheServiceServer) Get(ctx context.Context, req *Key) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCacheServiceServer) Put(ctx context.Context, req *KeyValue) (*OprationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedCacheServiceServer) Clear(ctx context.Context, req *Nil) (*OprationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (*UnimplementedCacheServiceServer) Use(ctx context.Context, req *DistKeyValue) (*Nil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Use not implemented")
}
func (*UnimplementedCacheServiceServer) ClearDist(ctx context.Context, req *Data) (*Nil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearDist not implemented")
}
func (*UnimplementedCacheServiceServer) Check(ctx context.Context, req *Nil) (*Nil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedCacheServiceServer) ImAlive(ctx context.Context, req *Data) (*Nil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImAlive not implemented")
}

func RegisterCacheServiceServer(s *grpc.Server, srv CacheServiceServer) {
	s.RegisterService(&_CacheService_serviceDesc, srv)
}

func _CacheService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.CacheService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.CacheService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Put(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.CacheService/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Clear(ctx, req.(*Nil))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Use_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistKeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Use(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.CacheService/Use",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Use(ctx, req.(*DistKeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_ClearDist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).ClearDist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.CacheService/ClearDist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).ClearDist(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.CacheService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Check(ctx, req.(*Nil))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_ImAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).ImAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.CacheService/ImAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).ImAlive(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

var _CacheService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CacheService_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _CacheService_Put_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _CacheService_Clear_Handler,
		},
		{
			MethodName: "Use",
			Handler:    _CacheService_Use_Handler,
		},
		{
			MethodName: "ClearDist",
			Handler:    _CacheService_ClearDist_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _CacheService_Check_Handler,
		},
		{
			MethodName: "ImAlive",
			Handler:    _CacheService_ImAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
