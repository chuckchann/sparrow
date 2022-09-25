// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/protobuf_spec/user/user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	common "sparrow/api/protobuf_spec/common"
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

type GetUserInfoRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoRequest) Reset()         { *m = GetUserInfoRequest{} }
func (m *GetUserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoRequest) ProtoMessage()    {}
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_506d464f8c3de92c, []int{0}
}

func (m *GetUserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoRequest.Unmarshal(m, b)
}
func (m *GetUserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetUserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoRequest.Merge(m, src)
}
func (m *GetUserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoRequest.Size(m)
}
func (m *GetUserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoRequest proto.InternalMessageInfo

func (m *GetUserInfoRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender               int64    `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Mobile               string   `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Job                  string   `protobuf:"bytes,7,opt,name=job,proto3" json:"job,omitempty"`
	Test                 string   `protobuf:"bytes,8,opt,name=test,proto3" json:"test,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_506d464f8c3de92c, []int{1}
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

func (m *UserInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetGender() int64 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UserInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserInfo) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *UserInfo) GetTest() string {
	if m != nil {
		return m.Test
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserInfoRequest)(nil), "user.GetUserInfoRequest")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
}

func init() { proto.RegisterFile("api/protobuf_spec/user/user.proto", fileDescriptor_506d464f8c3de92c) }

var fileDescriptor_506d464f8c3de92c = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xcd, 0x9f, 0xa6, 0x75, 0x04, 0x29, 0x83, 0xc8, 0xd0, 0x83, 0xc4, 0xa0, 0xd0, 0x53,
	0x02, 0x7a, 0xf4, 0xa6, 0x07, 0xf1, 0x1a, 0xf0, 0xe2, 0x45, 0x92, 0x66, 0x2a, 0x91, 0x26, 0x1b,
	0x77, 0x12, 0xfc, 0x62, 0x7e, 0x40, 0xd9, 0x49, 0x0a, 0x62, 0x7b, 0xd9, 0x7d, 0xef, 0xed, 0x8f,
	0x65, 0xe6, 0xc1, 0x75, 0xd1, 0xd5, 0x59, 0x67, 0x4d, 0x6f, 0xca, 0x61, 0xfb, 0x2e, 0x1d, 0x6f,
	0xb2, 0x41, 0xd8, 0xea, 0x91, 0x6a, 0x8e, 0xa1, 0xd3, 0xab, 0xdb, 0x43, 0x70, 0x63, 0x9a, 0xc6,
	0xb4, 0xd3, 0x35, 0xc2, 0xc9, 0x0d, 0xe0, 0x33, 0xf7, 0xaf, 0xc2, 0xf6, 0xa5, 0xdd, 0x9a, 0x9c,
	0xbf, 0x06, 0x96, 0x1e, 0xcf, 0xc1, 0xaf, 0x2b, 0xf2, 0x62, 0x6f, 0x1d, 0xe4, 0x7e, 0x5d, 0x25,
	0x3f, 0x1e, 0x2c, 0xf6, 0xcc, 0xff, 0x47, 0x44, 0x08, 0xdb, 0xa2, 0x61, 0xf2, 0x63, 0x6f, 0x7d,
	0x9a, 0xab, 0xc6, 0x4b, 0x88, 0x3e, 0xb8, 0xad, 0xd8, 0x52, 0xa0, 0xdc, 0xe4, 0x90, 0x60, 0x5e,
	0x54, 0x95, 0x65, 0x11, 0x0a, 0x15, 0xdf, 0x5b, 0xbc, 0x80, 0x19, 0x37, 0x45, 0xbd, 0xa3, 0x99,
	0xe6, 0xa3, 0x71, 0xff, 0x34, 0xa6, 0xac, 0x77, 0x4c, 0x91, 0xc6, 0x93, 0xc3, 0x25, 0x04, 0x9f,
	0xa6, 0xa4, 0xb9, 0x86, 0x4e, 0xba, 0x29, 0x7a, 0x96, 0x9e, 0x16, 0xe3, 0x14, 0x4e, 0xdf, 0x3d,
	0x41, 0xe8, 0xa6, 0xc6, 0x07, 0x38, 0xfb, 0xb3, 0x24, 0x52, 0xaa, 0x6d, 0x1d, 0xee, 0xbd, 0x5a,
	0xa6, 0x53, 0x39, 0x39, 0x4b, 0x67, 0x5a, 0xe1, 0xe4, 0xe4, 0x31, 0x7e, 0xbb, 0x92, 0xae, 0xb0,
	0xd6, 0x7c, 0x67, 0xc7, 0xbb, 0x2f, 0x23, 0xcd, 0xee, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x53,
	0x73, 0x4b, 0x0a, 0x9c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around sctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*common.Response, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*common.Response, error) {
	out := new(common.Response)
	err := c.cc.Invoke(ctx, "/user.User/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetUserInfo(context.Context, *GetUserInfoRequest) (*common.Response, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUserInfo(ctx context.Context, req *GetUserInfoRequest) (*common.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _User_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobuf_spec/user/user.proto",
}
