// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_nodeagent_secret.proto

package security

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type SecretRequest struct {
	Secrets              []*Secret        `protobuf:"bytes,1,rep,name=Secrets,proto3" json:"Secrets,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SecretRequest) Reset()         { *m = SecretRequest{} }
func (m *SecretRequest) String() string { return proto.CompactTextString(m) }
func (*SecretRequest) ProtoMessage()    {}
func (*SecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9c3fe257a0b458, []int{0}
}

func (m *SecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretRequest.Unmarshal(m, b)
}
func (m *SecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretRequest.Marshal(b, m, deterministic)
}
func (m *SecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretRequest.Merge(m, src)
}
func (m *SecretRequest) XXX_Size() int {
	return xxx_messageInfo_SecretRequest.Size(m)
}
func (m *SecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SecretRequest proto.InternalMessageInfo

func (m *SecretRequest) GetSecrets() []*Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *SecretRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type SecretResponse struct {
	Secrets              []*Secret           `protobuf:"bytes,1,rep,name=Secrets,proto3" json:"Secrets,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SecretResponse) Reset()         { *m = SecretResponse{} }
func (m *SecretResponse) String() string { return proto.CompactTextString(m) }
func (*SecretResponse) ProtoMessage()    {}
func (*SecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9c3fe257a0b458, []int{1}
}

func (m *SecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretResponse.Unmarshal(m, b)
}
func (m *SecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretResponse.Marshal(b, m, deterministic)
}
func (m *SecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretResponse.Merge(m, src)
}
func (m *SecretResponse) XXX_Size() int {
	return xxx_messageInfo_SecretResponse.Size(m)
}
func (m *SecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SecretResponse proto.InternalMessageInfo

func (m *SecretResponse) GetSecrets() []*Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *SecretResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *SecretResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Secret struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Filename             string         `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Value                string         `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	VaultId              string         `protobuf:"bytes,5,opt,name=vaultId,proto3" json:"vaultId,omitempty"`
	VaultName            string         `protobuf:"bytes,6,opt,name=vaultName,proto3" json:"vaultName,omitempty"`
	GroupName            string         `protobuf:"bytes,7,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Status               *common.Status `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity `protobuf:"bytes,9,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags   `protobuf:"bytes,10,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9c3fe257a0b458, []int{2}
}

func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Secret) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Secret) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Secret) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Secret) GetVaultId() string {
	if m != nil {
		return m.VaultId
	}
	return ""
}

func (m *Secret) GetVaultName() string {
	if m != nil {
		return m.VaultName
	}
	return ""
}

func (m *Secret) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Secret) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Secret) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Secret) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*SecretRequest)(nil), "moc.nodeagent.security.SecretRequest")
	proto.RegisterType((*SecretResponse)(nil), "moc.nodeagent.security.SecretResponse")
	proto.RegisterType((*Secret)(nil), "moc.nodeagent.security.Secret")
}

func init() { proto.RegisterFile("moc_nodeagent_secret.proto", fileDescriptor_fa9c3fe257a0b458) }

var fileDescriptor_fa9c3fe257a0b458 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x69, 0xea, 0xd4, 0x13, 0x35, 0x87, 0x15, 0x82, 0x95, 0x05, 0x55, 0x14, 0x04, 0xca,
	0xc9, 0x16, 0x86, 0x03, 0x57, 0x2a, 0xf5, 0xd0, 0x0b, 0x48, 0x6e, 0x85, 0x04, 0x97, 0xca, 0x71,
	0x26, 0xae, 0x85, 0xed, 0x31, 0xfb, 0x11, 0x94, 0x13, 0x7f, 0x82, 0x2b, 0xff, 0x15, 0x79, 0xd6,
	0x36, 0x20, 0x21, 0xe5, 0xd0, 0xd3, 0xee, 0xbc, 0xf7, 0xe6, 0xbd, 0xd5, 0xcc, 0x42, 0x58, 0x53,
	0x7e, 0xd7, 0xd0, 0x16, 0xb3, 0x02, 0x1b, 0x73, 0xa7, 0x31, 0x57, 0x68, 0xa2, 0x56, 0x91, 0x21,
	0xf1, 0xa4, 0xa6, 0x3c, 0x1a, 0xb9, 0x48, 0x63, 0x6e, 0x55, 0x69, 0x0e, 0xe1, 0x45, 0x41, 0x54,
	0x54, 0x18, 0xb3, 0x6a, 0x63, 0x77, 0xf1, 0x77, 0x95, 0xb5, 0x2d, 0x2a, 0xed, 0xfa, 0xc2, 0xa7,
	0x9d, 0x67, 0x4e, 0x75, 0x4d, 0x4d, 0x7f, 0x38, 0x62, 0xf5, 0x03, 0xce, 0x6f, 0x38, 0x20, 0xc5,
	0x6f, 0x16, 0xb5, 0x11, 0xef, 0x60, 0xe6, 0x00, 0x2d, 0xbd, 0xe5, 0xc9, 0x7a, 0x9e, 0x5c, 0x44,
	0xff, 0xcf, 0x8c, 0xfa, 0xbe, 0x41, 0x2e, 0xde, 0xc2, 0xf9, 0xc7, 0x16, 0x55, 0x66, 0x4a, 0x6a,
	0x6e, 0x0f, 0x2d, 0xca, 0xc9, 0xd2, 0x5b, 0x2f, 0x92, 0x05, 0xf7, 0x8f, 0x4c, 0xfa, 0xaf, 0x68,
	0xf5, 0xd3, 0x83, 0xc5, 0xf0, 0x02, 0xdd, 0x52, 0xa3, 0xf1, 0x01, 0x4f, 0x48, 0xc0, 0x4f, 0x51,
	0xdb, 0xca, 0x70, 0xf6, 0x3c, 0x09, 0x23, 0x37, 0x97, 0x68, 0x98, 0x4b, 0x74, 0x49, 0x54, 0x7d,
	0xca, 0x2a, 0x8b, 0x69, 0xaf, 0x14, 0x8f, 0xe1, 0xf4, 0x4a, 0x29, 0x52, 0xf2, 0x64, 0xe9, 0xad,
	0x83, 0xd4, 0x15, 0xab, 0x5f, 0x13, 0xf0, 0x9d, 0xab, 0x10, 0x30, 0x6d, 0xb2, 0x1a, 0xa5, 0xc7,
	0x3c, 0xdf, 0xc5, 0x02, 0x26, 0xe5, 0x96, 0x43, 0x82, 0x74, 0x52, 0x6e, 0x45, 0x08, 0x67, 0xbb,
	0xb2, 0x42, 0xd6, 0x39, 0x9f, 0xb1, 0xee, 0x02, 0xf6, 0x5d, 0xa2, 0x9c, 0xba, 0x00, 0x2e, 0x84,
	0x84, 0xd9, 0x3e, 0xb3, 0x95, 0xb9, 0xde, 0xca, 0x53, 0xc6, 0x87, 0x52, 0x3c, 0x83, 0x80, 0xaf,
	0x1f, 0x3a, 0x33, 0x9f, 0xb9, 0x3f, 0x40, 0xc7, 0x16, 0x8a, 0x6c, 0xcb, 0xec, 0xcc, 0xb1, 0x23,
	0x20, 0x5e, 0x80, 0xaf, 0x4d, 0x66, 0xac, 0x96, 0x67, 0x3c, 0x80, 0x39, 0x4f, 0xee, 0x86, 0xa1,
	0xb4, 0xa7, 0x3a, 0x11, 0x36, 0xa6, 0x34, 0x07, 0x19, 0xfc, 0x25, 0xba, 0x62, 0x28, 0xed, 0x29,
	0xf1, 0x1c, 0xa6, 0x26, 0x2b, 0xb4, 0x04, 0x96, 0x04, 0x2c, 0xb9, 0xcd, 0x0a, 0x9d, 0x32, 0x9c,
	0xdc, 0xc3, 0xdc, 0x8d, 0xe7, 0x7d, 0xb7, 0x11, 0xf1, 0x19, 0xfc, 0xeb, 0x66, 0x4f, 0x5f, 0x51,
	0xbc, 0x3c, 0xb2, 0x2b, 0xf7, 0xcd, 0xc2, 0x57, 0xc7, 0x64, 0xee, 0x2f, 0xac, 0x1e, 0x5d, 0xbe,
	0xfe, 0x12, 0x17, 0xa5, 0xb9, 0xb7, 0x9b, 0x28, 0xa7, 0x3a, 0xae, 0xcb, 0x5c, 0x91, 0xa6, 0x9d,
	0x89, 0x6b, 0xca, 0x63, 0xd5, 0xe6, 0xf1, 0xe8, 0x11, 0x0f, 0x1e, 0x1b, 0x9f, 0xd7, 0xfd, 0xe6,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0xdc, 0x1f, 0xf8, 0x4a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretAgentClient is the client API for SecretAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretAgentClient interface {
	Invoke(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*SecretResponse, error)
}

type secretAgentClient struct {
	cc *grpc.ClientConn
}

func NewSecretAgentClient(cc *grpc.ClientConn) SecretAgentClient {
	return &secretAgentClient{cc}
}

func (c *secretAgentClient) Invoke(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*SecretResponse, error) {
	out := new(SecretResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.SecretAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretAgentServer is the server API for SecretAgent service.
type SecretAgentServer interface {
	Invoke(context.Context, *SecretRequest) (*SecretResponse, error)
}

// UnimplementedSecretAgentServer can be embedded to have forward compatible implementations.
type UnimplementedSecretAgentServer struct {
}

func (*UnimplementedSecretAgentServer) Invoke(ctx context.Context, req *SecretRequest) (*SecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterSecretAgentServer(s *grpc.Server, srv SecretAgentServer) {
	s.RegisterService(&_SecretAgent_serviceDesc, srv)
}

func _SecretAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.SecretAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretAgentServer).Invoke(ctx, req.(*SecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.security.SecretAgent",
	HandlerType: (*SecretAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _SecretAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_nodeagent_secret.proto",
}
