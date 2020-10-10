// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mixlunch.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// TargetDate is represented as date format string. i.e. '2019-05-01'
type TargetDate struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetDate) Reset()         { *m = TargetDate{} }
func (m *TargetDate) String() string { return proto.CompactTextString(m) }
func (*TargetDate) ProtoMessage()    {}
func (*TargetDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f6016d559ef54e, []int{0}
}

func (m *TargetDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetDate.Unmarshal(m, b)
}
func (m *TargetDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetDate.Marshal(b, m, deterministic)
}
func (m *TargetDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetDate.Merge(m, src)
}
func (m *TargetDate) XXX_Size() int {
	return xxx_messageInfo_TargetDate.Size(m)
}
func (m *TargetDate) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetDate.DiscardUnknown(m)
}

var xxx_messageInfo_TargetDate proto.InternalMessageInfo

func (m *TargetDate) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

// UserModelForMatching is represented as user model for MixLunch matching program
type UserModelForMatching struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FreeFrom             string   `protobuf:"bytes,2,opt,name=free_from,json=freeFrom,proto3" json:"free_from,omitempty"`
	FreeTo               string   `protobuf:"bytes,3,opt,name=free_to,json=freeTo,proto3" json:"free_to,omitempty"`
	UserName             string   `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserModelForMatching) Reset()         { *m = UserModelForMatching{} }
func (m *UserModelForMatching) String() string { return proto.CompactTextString(m) }
func (*UserModelForMatching) ProtoMessage()    {}
func (*UserModelForMatching) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f6016d559ef54e, []int{1}
}

func (m *UserModelForMatching) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserModelForMatching.Unmarshal(m, b)
}
func (m *UserModelForMatching) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserModelForMatching.Marshal(b, m, deterministic)
}
func (m *UserModelForMatching) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserModelForMatching.Merge(m, src)
}
func (m *UserModelForMatching) XXX_Size() int {
	return xxx_messageInfo_UserModelForMatching.Size(m)
}
func (m *UserModelForMatching) XXX_DiscardUnknown() {
	xxx_messageInfo_UserModelForMatching.DiscardUnknown(m)
}

var xxx_messageInfo_UserModelForMatching proto.InternalMessageInfo

func (m *UserModelForMatching) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserModelForMatching) GetFreeFrom() string {
	if m != nil {
		return m.FreeFrom
	}
	return ""
}

func (m *UserModelForMatching) GetFreeTo() string {
	if m != nil {
		return m.FreeTo
	}
	return ""
}

func (m *UserModelForMatching) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserModelForMatching) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// Party is represented as party model MixLunch matching program created
type Party struct {
	StartFrom            string                  `protobuf:"bytes,1,opt,name=start_from,json=startFrom,proto3" json:"start_from,omitempty"`
	EndTo                string                  `protobuf:"bytes,2,opt,name=end_to,json=endTo,proto3" json:"end_to,omitempty"`
	Members              []*UserModelForMatching `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	ChatRoomId           string                  `protobuf:"bytes,4,opt,name=chat_room_id,json=chatRoomId,proto3" json:"chat_room_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Party) Reset()         { *m = Party{} }
func (m *Party) String() string { return proto.CompactTextString(m) }
func (*Party) ProtoMessage()    {}
func (*Party) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f6016d559ef54e, []int{2}
}

func (m *Party) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Party.Unmarshal(m, b)
}
func (m *Party) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Party.Marshal(b, m, deterministic)
}
func (m *Party) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Party.Merge(m, src)
}
func (m *Party) XXX_Size() int {
	return xxx_messageInfo_Party.Size(m)
}
func (m *Party) XXX_DiscardUnknown() {
	xxx_messageInfo_Party.DiscardUnknown(m)
}

var xxx_messageInfo_Party proto.InternalMessageInfo

func (m *Party) GetStartFrom() string {
	if m != nil {
		return m.StartFrom
	}
	return ""
}

func (m *Party) GetEndTo() string {
	if m != nil {
		return m.EndTo
	}
	return ""
}

func (m *Party) GetMembers() []*UserModelForMatching {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Party) GetChatRoomId() string {
	if m != nil {
		return m.ChatRoomId
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0f6016d559ef54e, []int{3}
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

func init() {
	proto.RegisterType((*TargetDate)(nil), "pb.TargetDate")
	proto.RegisterType((*UserModelForMatching)(nil), "pb.UserModelForMatching")
	proto.RegisterType((*Party)(nil), "pb.Party")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

func init() { proto.RegisterFile("mixlunch.proto", fileDescriptor_c0f6016d559ef54e) }

var fileDescriptor_c0f6016d559ef54e = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xd1, 0x6a, 0xea, 0x40,
	0x10, 0x35, 0x6a, 0xd4, 0xcc, 0xbd, 0xd7, 0x87, 0xbd, 0x96, 0x06, 0x4b, 0x21, 0xe4, 0xc9, 0x52,
	0x90, 0x62, 0x3f, 0xc1, 0x56, 0x11, 0x6a, 0x29, 0x62, 0x9f, 0x65, 0xe3, 0x8e, 0x1a, 0x70, 0xb3,
	0x61, 0x33, 0x82, 0x7e, 0x47, 0x5f, 0xfa, 0x07, 0xfd, 0xcd, 0x32, 0x1b, 0xc5, 0x52, 0xda, 0xb7,
	0x99, 0x33, 0xe7, 0xcc, 0x99, 0x99, 0x5d, 0x68, 0xeb, 0x74, 0xbf, 0xdd, 0x65, 0xcb, 0x4d, 0x3f,
	0xb7, 0x86, 0x8c, 0xa8, 0xe6, 0x49, 0x1c, 0x01, 0xcc, 0xa5, 0x5d, 0x23, 0x3d, 0x48, 0x42, 0x21,
	0xa0, 0xae, 0x24, 0x61, 0xe8, 0x45, 0x5e, 0x2f, 0x98, 0xb9, 0x38, 0x7e, 0xf7, 0xa0, 0xf3, 0x5a,
	0xa0, 0x9d, 0x1a, 0x85, 0xdb, 0x91, 0xb1, 0x53, 0x49, 0xcb, 0x4d, 0x9a, 0xad, 0xc5, 0x25, 0x34,
	0x77, 0x05, 0xda, 0x45, 0xaa, 0x8e, 0xfc, 0x06, 0xa7, 0x13, 0x25, 0xae, 0x20, 0x58, 0x59, 0xc4,
	0xc5, 0xca, 0x1a, 0x1d, 0x56, 0x5d, 0xa9, 0xc5, 0xc0, 0xc8, 0x1a, 0xcd, 0x2a, 0x57, 0x24, 0x13,
	0xd6, 0x4a, 0x15, 0xa7, 0x73, 0xc3, 0x2a, 0xd7, 0x2e, 0x93, 0x1a, 0xc3, 0x7a, 0xa9, 0x62, 0xe0,
	0x59, 0x6a, 0x14, 0x1d, 0xf0, 0x51, 0xcb, 0x74, 0x1b, 0xfa, 0xae, 0x50, 0x26, 0xf1, 0x9b, 0x07,
	0xfe, 0x8b, 0xb4, 0x74, 0x10, 0xd7, 0x00, 0x05, 0x49, 0x4b, 0xa5, 0x67, 0x39, 0x4e, 0xe0, 0x10,
	0x67, 0x7a, 0x01, 0x0d, 0xcc, 0x14, 0x7b, 0x56, 0x8f, 0xfa, 0x4c, 0xcd, 0x8d, 0x18, 0x40, 0x53,
	0xa3, 0x4e, 0xd0, 0x16, 0x61, 0x2d, 0xaa, 0xf5, 0xfe, 0x0c, 0xc2, 0x7e, 0x9e, 0xf4, 0x7f, 0x5a,
	0x76, 0x76, 0x22, 0x8a, 0x08, 0xfe, 0x2e, 0x37, 0x92, 0x16, 0xd6, 0x18, 0xcd, 0xab, 0x97, 0x93,
	0x02, 0x63, 0x33, 0x63, 0xf4, 0x44, 0xc5, 0x4d, 0xf0, 0x1f, 0x75, 0x4e, 0x87, 0xc1, 0x87, 0x07,
	0xad, 0x69, 0xba, 0x7f, 0xe2, 0x93, 0x8b, 0x21, 0xfc, 0x1f, 0x23, 0x71, 0xef, 0xe2, 0xeb, 0x11,
	0xdb, 0xec, 0x78, 0x7e, 0x81, 0xee, 0xaf, 0x13, 0xc4, 0x95, 0x3b, 0x4f, 0xdc, 0xc0, 0xbf, 0xa1,
	0x45, 0x49, 0xc8, 0x5b, 0xa7, 0x58, 0x88, 0x80, 0xe9, 0xee, 0x04, 0x5d, 0x17, 0x3a, 0xe3, 0xb8,
	0xd2, 0xf3, 0xc4, 0x2d, 0xc0, 0x18, 0xe9, 0xc4, 0xfb, 0x6e, 0x73, 0xd6, 0x71, 0xdf, 0xa4, 0xe1,
	0x3e, 0xc4, 0xfd, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xbc, 0x7e, 0xe7, 0x22, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MixLunchClient is the client API for MixLunch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MixLunchClient interface {
	// A server-to-client streaming RPC.
	//
	// Obtains the users information by passing target date
	GetUsersForMatching(ctx context.Context, in *TargetDate, opts ...grpc.CallOption) (MixLunch_GetUsersForMatchingClient, error)
	// A client-to-server streaming RPC.
	//
	// Uploads the parties information
	CreateParties(ctx context.Context, opts ...grpc.CallOption) (MixLunch_CreatePartiesClient, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the parties information by passing target date
	GetParties(ctx context.Context, in *TargetDate, opts ...grpc.CallOption) (MixLunch_GetPartiesClient, error)
}

type mixLunchClient struct {
	cc *grpc.ClientConn
}

func NewMixLunchClient(cc *grpc.ClientConn) MixLunchClient {
	return &mixLunchClient{cc}
}

func (c *mixLunchClient) GetUsersForMatching(ctx context.Context, in *TargetDate, opts ...grpc.CallOption) (MixLunch_GetUsersForMatchingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MixLunch_serviceDesc.Streams[0], "/pb.MixLunch/GetUsersForMatching", opts...)
	if err != nil {
		return nil, err
	}
	x := &mixLunchGetUsersForMatchingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MixLunch_GetUsersForMatchingClient interface {
	Recv() (*UserModelForMatching, error)
	grpc.ClientStream
}

type mixLunchGetUsersForMatchingClient struct {
	grpc.ClientStream
}

func (x *mixLunchGetUsersForMatchingClient) Recv() (*UserModelForMatching, error) {
	m := new(UserModelForMatching)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mixLunchClient) CreateParties(ctx context.Context, opts ...grpc.CallOption) (MixLunch_CreatePartiesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MixLunch_serviceDesc.Streams[1], "/pb.MixLunch/CreateParties", opts...)
	if err != nil {
		return nil, err
	}
	x := &mixLunchCreatePartiesClient{stream}
	return x, nil
}

type MixLunch_CreatePartiesClient interface {
	Send(*Party) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type mixLunchCreatePartiesClient struct {
	grpc.ClientStream
}

func (x *mixLunchCreatePartiesClient) Send(m *Party) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mixLunchCreatePartiesClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mixLunchClient) GetParties(ctx context.Context, in *TargetDate, opts ...grpc.CallOption) (MixLunch_GetPartiesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MixLunch_serviceDesc.Streams[2], "/pb.MixLunch/GetParties", opts...)
	if err != nil {
		return nil, err
	}
	x := &mixLunchGetPartiesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MixLunch_GetPartiesClient interface {
	Recv() (*Party, error)
	grpc.ClientStream
}

type mixLunchGetPartiesClient struct {
	grpc.ClientStream
}

func (x *mixLunchGetPartiesClient) Recv() (*Party, error) {
	m := new(Party)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MixLunchServer is the server API for MixLunch service.
type MixLunchServer interface {
	// A server-to-client streaming RPC.
	//
	// Obtains the users information by passing target date
	GetUsersForMatching(*TargetDate, MixLunch_GetUsersForMatchingServer) error
	// A client-to-server streaming RPC.
	//
	// Uploads the parties information
	CreateParties(MixLunch_CreatePartiesServer) error
	// A server-to-client streaming RPC.
	//
	// Obtains the parties information by passing target date
	GetParties(*TargetDate, MixLunch_GetPartiesServer) error
}

func RegisterMixLunchServer(s *grpc.Server, srv MixLunchServer) {
	s.RegisterService(&_MixLunch_serviceDesc, srv)
}

func _MixLunch_GetUsersForMatching_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TargetDate)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MixLunchServer).GetUsersForMatching(m, &mixLunchGetUsersForMatchingServer{stream})
}

type MixLunch_GetUsersForMatchingServer interface {
	Send(*UserModelForMatching) error
	grpc.ServerStream
}

type mixLunchGetUsersForMatchingServer struct {
	grpc.ServerStream
}

func (x *mixLunchGetUsersForMatchingServer) Send(m *UserModelForMatching) error {
	return x.ServerStream.SendMsg(m)
}

func _MixLunch_CreateParties_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MixLunchServer).CreateParties(&mixLunchCreatePartiesServer{stream})
}

type MixLunch_CreatePartiesServer interface {
	SendAndClose(*Empty) error
	Recv() (*Party, error)
	grpc.ServerStream
}

type mixLunchCreatePartiesServer struct {
	grpc.ServerStream
}

func (x *mixLunchCreatePartiesServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mixLunchCreatePartiesServer) Recv() (*Party, error) {
	m := new(Party)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MixLunch_GetParties_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TargetDate)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MixLunchServer).GetParties(m, &mixLunchGetPartiesServer{stream})
}

type MixLunch_GetPartiesServer interface {
	Send(*Party) error
	grpc.ServerStream
}

type mixLunchGetPartiesServer struct {
	grpc.ServerStream
}

func (x *mixLunchGetPartiesServer) Send(m *Party) error {
	return x.ServerStream.SendMsg(m)
}

var _MixLunch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MixLunch",
	HandlerType: (*MixLunchServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUsersForMatching",
			Handler:       _MixLunch_GetUsersForMatching_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateParties",
			Handler:       _MixLunch_CreateParties_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetParties",
			Handler:       _MixLunch_GetParties_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mixlunch.proto",
}