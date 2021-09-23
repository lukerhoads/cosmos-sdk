// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/vesting/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgCreateVestingAccount defines a message that enables creating a vesting
// account.
type MsgCreateVestingAccount struct {
	FromAddress string                                   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress   string                                   `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	EndTime     int64                                    `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Delayed     bool                                     `protobuf:"varint,5,opt,name=delayed,proto3" json:"delayed,omitempty"`
}

func (m *MsgCreateVestingAccount) Reset()         { *m = MsgCreateVestingAccount{} }
func (m *MsgCreateVestingAccount) String() string { return proto.CompactTextString(m) }
func (*MsgCreateVestingAccount) ProtoMessage()    {}
func (*MsgCreateVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_5338ca97811f9792, []int{0}
}
func (m *MsgCreateVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateVestingAccount.Merge(m, src)
}
func (m *MsgCreateVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateVestingAccount proto.InternalMessageInfo

func (m *MsgCreateVestingAccount) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgCreateVestingAccount) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgCreateVestingAccount) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MsgCreateVestingAccount) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *MsgCreateVestingAccount) GetDelayed() bool {
	if m != nil {
		return m.Delayed
	}
	return false
}

// MsgCreateVestingAccountResponse defines the Msg/CreateVestingAccount response type.
type MsgCreateVestingAccountResponse struct {
}

func (m *MsgCreateVestingAccountResponse) Reset()         { *m = MsgCreateVestingAccountResponse{} }
func (m *MsgCreateVestingAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateVestingAccountResponse) ProtoMessage()    {}
func (*MsgCreateVestingAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5338ca97811f9792, []int{1}
}
func (m *MsgCreateVestingAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateVestingAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateVestingAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateVestingAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateVestingAccountResponse.Merge(m, src)
}
func (m *MsgCreateVestingAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateVestingAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateVestingAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateVestingAccountResponse proto.InternalMessageInfo

// MsgCreateVestingAccount defines a message that enables creating a vesting
// account.
type MsgCreatePeriodicVestingAccount struct {
	FromAddress    string   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress      string   `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	StartTime      int64    `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" yaml:"start_time"`
	VestingPeriods []Period `protobuf:"bytes,4,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods"`
}

func (m *MsgCreatePeriodicVestingAccount) Reset()         { *m = MsgCreatePeriodicVestingAccount{} }
func (m *MsgCreatePeriodicVestingAccount) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePeriodicVestingAccount) ProtoMessage()    {}
func (*MsgCreatePeriodicVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_5338ca97811f9792, []int{2}
}
func (m *MsgCreatePeriodicVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePeriodicVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePeriodicVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePeriodicVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePeriodicVestingAccount.Merge(m, src)
}
func (m *MsgCreatePeriodicVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePeriodicVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePeriodicVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePeriodicVestingAccount proto.InternalMessageInfo

func (m *MsgCreatePeriodicVestingAccount) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgCreatePeriodicVestingAccount) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgCreatePeriodicVestingAccount) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *MsgCreatePeriodicVestingAccount) GetVestingPeriods() []Period {
	if m != nil {
		return m.VestingPeriods
	}
	return nil
}

// MsgCreateVestingAccountResponse defines the Msg/CreatePeriodicVestingAccount
// response type.
type MsgCreatePeriodicVestingAccountResponse struct {
}

func (m *MsgCreatePeriodicVestingAccountResponse) Reset() {
	*m = MsgCreatePeriodicVestingAccountResponse{}
}
func (m *MsgCreatePeriodicVestingAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePeriodicVestingAccountResponse) ProtoMessage()    {}
func (*MsgCreatePeriodicVestingAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5338ca97811f9792, []int{3}
}
func (m *MsgCreatePeriodicVestingAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePeriodicVestingAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePeriodicVestingAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePeriodicVestingAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePeriodicVestingAccountResponse.Merge(m, src)
}
func (m *MsgCreatePeriodicVestingAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePeriodicVestingAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePeriodicVestingAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePeriodicVestingAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateVestingAccount)(nil), "cosmos.vesting.v1beta1.MsgCreateVestingAccount")
	proto.RegisterType((*MsgCreateVestingAccountResponse)(nil), "cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse")
	proto.RegisterType((*MsgCreatePeriodicVestingAccount)(nil), "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount")
	proto.RegisterType((*MsgCreatePeriodicVestingAccountResponse)(nil), "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse")
}

func init() { proto.RegisterFile("cosmos/vesting/v1beta1/tx.proto", fileDescriptor_5338ca97811f9792) }

var fileDescriptor_5338ca97811f9792 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x31, 0x6f, 0xd3, 0x4e,
	0x1c, 0xf5, 0xc5, 0xf9, 0xb7, 0xcd, 0xf5, 0x2f, 0x2a, 0xdc, 0x40, 0xdd, 0x08, 0xd9, 0xc1, 0x42,
	0xc2, 0x0c, 0xb1, 0x49, 0xa9, 0x84, 0x94, 0x05, 0xd5, 0x1d, 0x21, 0x12, 0x32, 0x88, 0x81, 0x25,
	0x72, 0xec, 0xc3, 0x3d, 0x51, 0xfb, 0x22, 0xdf, 0xa5, 0x6a, 0x36, 0x3e, 0x02, 0x23, 0x13, 0x42,
	0x62, 0x63, 0xe6, 0x43, 0x74, 0xac, 0x98, 0x98, 0x02, 0x4a, 0x16, 0x46, 0x94, 0x4f, 0x80, 0x7c,
	0x77, 0x4e, 0x03, 0xb2, 0x5b, 0xc1, 0x94, 0xf8, 0xde, 0xfb, 0x3d, 0x3f, 0xbf, 0x77, 0x77, 0xd0,
	0x0c, 0x09, 0x4d, 0x08, 0x75, 0x4f, 0x10, 0x65, 0x38, 0x8d, 0xdd, 0x93, 0xee, 0x10, 0xb1, 0xa0,
	0xeb, 0xb2, 0x53, 0x67, 0x94, 0x11, 0x46, 0xb4, 0x9b, 0x82, 0xe0, 0x48, 0x82, 0x23, 0x09, 0xad,
	0x66, 0x4c, 0x62, 0xc2, 0x29, 0x6e, 0xfe, 0x4f, 0xb0, 0x5b, 0x86, 0x94, 0x1b, 0x06, 0x14, 0x2d,
	0xb5, 0x42, 0x82, 0x53, 0x89, 0xef, 0x0a, 0x7c, 0x20, 0x06, 0xa5, 0xb4, 0x80, 0xee, 0x54, 0x38,
	0x29, 0x5e, 0xcc, 0x59, 0xd6, 0xcf, 0x1a, 0xdc, 0xe9, 0xd3, 0xf8, 0x30, 0x43, 0x01, 0x43, 0x2f,
	0x04, 0x74, 0x10, 0x86, 0x64, 0x9c, 0x32, 0xcd, 0x87, 0xff, 0xbf, 0xca, 0x48, 0x32, 0x08, 0xa2,
	0x28, 0x43, 0x94, 0xea, 0xa0, 0x0d, 0xec, 0x86, 0xe7, 0x2e, 0xa6, 0xe6, 0xf6, 0x24, 0x48, 0x8e,
	0x7b, 0xd6, 0x2a, 0x6a, 0x7d, 0xf9, 0xdc, 0x69, 0x4a, 0x03, 0x07, 0x62, 0xe9, 0x19, 0xcb, 0x70,
	0x1a, 0xfb, 0x9b, 0x39, 0x4d, 0x2e, 0x69, 0x4f, 0x20, 0x64, 0x64, 0xa9, 0x58, 0xe3, 0x8a, 0x9d,
	0xc5, 0xd4, 0xbc, 0x2e, 0x14, 0x2f, 0xb0, 0x6a, 0xbd, 0x06, 0x23, 0x85, 0x5a, 0x08, 0xd7, 0x82,
	0x24, 0xf7, 0xaa, 0xab, 0x6d, 0xd5, 0xde, 0xdc, 0xdb, 0x75, 0xe4, 0x44, 0x9e, 0x57, 0x11, 0xad,
	0x73, 0x48, 0x70, 0xea, 0xdd, 0x3f, 0x9b, 0x9a, 0xca, 0xa7, 0x6f, 0xa6, 0x1d, 0x63, 0x76, 0x34,
	0x1e, 0x3a, 0x21, 0x49, 0x64, 0x5e, 0xf2, 0xa7, 0x43, 0xa3, 0xd7, 0x2e, 0x9b, 0x8c, 0x10, 0xe5,
	0x03, 0xd4, 0x97, 0xd2, 0x9a, 0x03, 0x37, 0x50, 0x1a, 0x0d, 0x18, 0x4e, 0x90, 0x5e, 0x6f, 0x03,
	0x5b, 0xf5, 0xb6, 0x17, 0x53, 0x73, 0x4b, 0x18, 0x2e, 0x10, 0xcb, 0x5f, 0x47, 0x69, 0xf4, 0x1c,
	0x27, 0x48, 0xd3, 0xe1, 0x7a, 0x84, 0x8e, 0x83, 0x09, 0x8a, 0xf4, 0xff, 0xda, 0xc0, 0xde, 0xf0,
	0x8b, 0xc7, 0x5e, 0xfd, 0xc7, 0x07, 0x13, 0x58, 0xb7, 0xa1, 0x59, 0x91, 0xb8, 0x8f, 0xe8, 0x88,
	0xa4, 0x14, 0x59, 0xef, 0x6b, 0x2b, 0x9c, 0xa7, 0x28, 0xc3, 0x24, 0xc2, 0xe1, 0x1f, 0xed, 0xf4,
	0x4a, 0xdb, 0xd9, 0xa9, 0x68, 0xe7, 0xf7, 0x16, 0xf6, 0x4b, 0x5a, 0xb8, 0x51, 0xda, 0xc2, 0x6a,
	0xda, 0xfb, 0x10, 0x52, 0x16, 0x64, 0x4c, 0x44, 0xa1, 0xf2, 0x28, 0x56, 0xa6, 0x2e, 0x30, 0xcb,
	0x6f, 0xf0, 0x07, 0x1e, 0x47, 0x1f, 0x6e, 0xc9, 0x2d, 0x37, 0x18, 0xf1, 0x2f, 0xa1, 0x7a, 0x9d,
	0x97, 0x65, 0x38, 0xe5, 0x47, 0xc1, 0x11, 0x1f, 0xec, 0xd5, 0xf3, 0xc6, 0xfc, 0x6b, 0x12, 0x15,
	0x8b, 0x94, 0x67, 0xa8, 0x58, 0xf7, 0xe0, 0xdd, 0x2b, 0xf2, 0x29, 0xb2, 0xdc, 0xfb, 0x58, 0x83,
	0x6a, 0x9f, 0xc6, 0xda, 0x1b, 0x00, 0x9b, 0xa5, 0xdb, 0xdc, 0xad, 0xf2, 0x51, 0xd1, 0x52, 0xeb,
	0xe1, 0x5f, 0x0e, 0x14, 0x56, 0xb4, 0x77, 0x00, 0xde, 0xba, 0xb4, 0xd3, 0xab, 0x95, 0xcb, 0x07,
	0x5b, 0x8f, 0xfe, 0x71, 0xb0, 0xb0, 0xe6, 0x3d, 0x3e, 0x9b, 0x19, 0xe0, 0x7c, 0x66, 0x80, 0xef,
	0x33, 0x03, 0xbc, 0x9d, 0x1b, 0xca, 0xf9, 0xdc, 0x50, 0xbe, 0xce, 0x0d, 0xe5, 0x65, 0xf7, 0xd2,
	0x03, 0x73, 0xea, 0x06, 0x63, 0x76, 0xb4, 0xbc, 0x64, 0xf8, 0xf9, 0x19, 0xae, 0xf1, 0xbb, 0xe5,
	0xc1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x1b, 0x2d, 0xe5, 0x0d, 0x05, 0x00, 0x00,
}

func (this *MsgCreateVestingAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgCreateVestingAccount)
	if !ok {
		that2, ok := that.(MsgCreateVestingAccount)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FromAddress != that1.FromAddress {
		return false
	}
	if this.ToAddress != that1.ToAddress {
		return false
	}
	if len(this.Amount) != len(that1.Amount) {
		return false
	}
	for i := range this.Amount {
		if !this.Amount[i].Equal(&that1.Amount[i]) {
			return false
		}
	}
	if this.EndTime != that1.EndTime {
		return false
	}
	if this.Delayed != that1.Delayed {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateVestingAccount defines a method that enables creating a vesting
	// account.
	CreateVestingAccount(ctx context.Context, in *MsgCreateVestingAccount, opts ...grpc.CallOption) (*MsgCreateVestingAccountResponse, error)
	// CreatePeriodicVestingAccount defines a method that enables creating a
	// periodic vesting account.
	CreatePeriodicVestingAccount(ctx context.Context, in *MsgCreatePeriodicVestingAccount, opts ...grpc.CallOption) (*MsgCreatePeriodicVestingAccountResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateVestingAccount(ctx context.Context, in *MsgCreateVestingAccount, opts ...grpc.CallOption) (*MsgCreateVestingAccountResponse, error) {
	out := new(MsgCreateVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.vesting.v1beta1.Msg/CreateVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePeriodicVestingAccount(ctx context.Context, in *MsgCreatePeriodicVestingAccount, opts ...grpc.CallOption) (*MsgCreatePeriodicVestingAccountResponse, error) {
	out := new(MsgCreatePeriodicVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.vesting.v1beta1.Msg/CreatePeriodicVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateVestingAccount defines a method that enables creating a vesting
	// account.
	CreateVestingAccount(context.Context, *MsgCreateVestingAccount) (*MsgCreateVestingAccountResponse, error)
	// CreatePeriodicVestingAccount defines a method that enables creating a
	// periodic vesting account.
	CreatePeriodicVestingAccount(context.Context, *MsgCreatePeriodicVestingAccount) (*MsgCreatePeriodicVestingAccountResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateVestingAccount(ctx context.Context, req *MsgCreateVestingAccount) (*MsgCreateVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVestingAccount not implemented")
}
func (*UnimplementedMsgServer) CreatePeriodicVestingAccount(ctx context.Context, req *MsgCreatePeriodicVestingAccount) (*MsgCreatePeriodicVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePeriodicVestingAccount not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.vesting.v1beta1.Msg/CreateVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateVestingAccount(ctx, req.(*MsgCreateVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePeriodicVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePeriodicVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePeriodicVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.vesting.v1beta1.Msg/CreatePeriodicVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePeriodicVestingAccount(ctx, req.(*MsgCreatePeriodicVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.vesting.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVestingAccount",
			Handler:    _Msg_CreateVestingAccount_Handler,
		},
		{
			MethodName: "CreatePeriodicVestingAccount",
			Handler:    _Msg_CreatePeriodicVestingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/vesting/v1beta1/tx.proto",
}

func (m *MsgCreateVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Delayed {
		i--
		if m.Delayed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.EndTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateVestingAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateVestingAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateVestingAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCreatePeriodicVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePeriodicVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePeriodicVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.StartTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePeriodicVestingAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePeriodicVestingAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePeriodicVestingAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.EndTime != 0 {
		n += 1 + sovTx(uint64(m.EndTime))
	}
	if m.Delayed {
		n += 2
	}
	return n
}

func (m *MsgCreateVestingAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCreatePeriodicVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.StartTime != 0 {
		n += 1 + sovTx(uint64(m.StartTime))
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreatePeriodicVestingAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delayed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.Delayed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateVestingAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateVestingAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreatePeriodicVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePeriodicVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePeriodicVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreatePeriodicVestingAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePeriodicVestingAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePeriodicVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
