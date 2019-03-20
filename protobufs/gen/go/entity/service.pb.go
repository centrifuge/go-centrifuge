// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity/service.proto

package entitypb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import entity "github.com/centrifuge/centrifuge-protobufs/gen/go/entity"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_45431f28cb32fa51, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type GetVersionRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_45431f28cb32fa51, []int{1}
}
func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionRequest.Unmarshal(m, b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
}
func (dst *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(dst, src)
}
func (m *GetVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetVersionRequest.Size(m)
}
func (m *GetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionRequest proto.InternalMessageInfo

func (m *GetVersionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *GetVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type EntityCreatePayload struct {
	Collaborators        []string    `protobuf:"bytes,1,rep,name=collaborators" json:"collaborators,omitempty"`
	Data                 *EntityData `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EntityCreatePayload) Reset()         { *m = EntityCreatePayload{} }
func (m *EntityCreatePayload) String() string { return proto.CompactTextString(m) }
func (*EntityCreatePayload) ProtoMessage()    {}
func (*EntityCreatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_45431f28cb32fa51, []int{2}
}
func (m *EntityCreatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityCreatePayload.Unmarshal(m, b)
}
func (m *EntityCreatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityCreatePayload.Marshal(b, m, deterministic)
}
func (dst *EntityCreatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityCreatePayload.Merge(dst, src)
}
func (m *EntityCreatePayload) XXX_Size() int {
	return xxx_messageInfo_EntityCreatePayload.Size(m)
}
func (m *EntityCreatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityCreatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_EntityCreatePayload proto.InternalMessageInfo

func (m *EntityCreatePayload) GetCollaborators() []string {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *EntityCreatePayload) GetData() *EntityData {
	if m != nil {
		return m.Data
	}
	return nil
}

type EntityUpdatePayload struct {
	Identifier           string      `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Collaborators        []string    `protobuf:"bytes,2,rep,name=collaborators" json:"collaborators,omitempty"`
	Data                 *EntityData `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EntityUpdatePayload) Reset()         { *m = EntityUpdatePayload{} }
func (m *EntityUpdatePayload) String() string { return proto.CompactTextString(m) }
func (*EntityUpdatePayload) ProtoMessage()    {}
func (*EntityUpdatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_45431f28cb32fa51, []int{3}
}
func (m *EntityUpdatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityUpdatePayload.Unmarshal(m, b)
}
func (m *EntityUpdatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityUpdatePayload.Marshal(b, m, deterministic)
}
func (dst *EntityUpdatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityUpdatePayload.Merge(dst, src)
}
func (m *EntityUpdatePayload) XXX_Size() int {
	return xxx_messageInfo_EntityUpdatePayload.Size(m)
}
func (m *EntityUpdatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityUpdatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_EntityUpdatePayload proto.InternalMessageInfo

func (m *EntityUpdatePayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *EntityUpdatePayload) GetCollaborators() []string {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *EntityUpdatePayload) GetData() *EntityData {
	if m != nil {
		return m.Data
	}
	return nil
}

type EntityResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Data                 *EntityData     `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EntityResponse) Reset()         { *m = EntityResponse{} }
func (m *EntityResponse) String() string { return proto.CompactTextString(m) }
func (*EntityResponse) ProtoMessage()    {}
func (*EntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_45431f28cb32fa51, []int{4}
}
func (m *EntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityResponse.Unmarshal(m, b)
}
func (m *EntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityResponse.Marshal(b, m, deterministic)
}
func (dst *EntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityResponse.Merge(dst, src)
}
func (m *EntityResponse) XXX_Size() int {
	return xxx_messageInfo_EntityResponse.Size(m)
}
func (m *EntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntityResponse proto.InternalMessageInfo

func (m *EntityResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *EntityResponse) GetData() *EntityData {
	if m != nil {
		return m.Data
	}
	return nil
}

// ResponseHeader contains a set of common fields for most document
type ResponseHeader struct {
	DocumentId           string   `protobuf:"bytes,1,opt,name=document_id,json=documentId" json:"document_id,omitempty"`
	VersionId            string   `protobuf:"bytes,2,opt,name=version_id,json=versionId" json:"version_id,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Collaborators        []string `protobuf:"bytes,4,rep,name=collaborators" json:"collaborators,omitempty"`
	TransactionId        string   `protobuf:"bytes,5,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_45431f28cb32fa51, []int{5}
}
func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (dst *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(dst, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func (m *ResponseHeader) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *ResponseHeader) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ResponseHeader) GetCollaborators() []string {
	if m != nil {
		return m.Collaborators
	}
	return nil
}

func (m *ResponseHeader) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

// EntityData is the default entity schema
type EntityData struct {
	Identity  []byte `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	LegalName string `protobuf:"bytes,2,opt,name=legal_name,json=legalName" json:"legal_name,omitempty"`
	// address
	Addresses []*entity.Address `protobuf:"bytes,3,rep,name=addresses" json:"addresses,omitempty"`
	// tax information
	PaymentDetails []*entity.PaymentDetail `protobuf:"bytes,4,rep,name=payment_details,json=paymentDetails" json:"payment_details,omitempty"`
	// Entity contact list
	Contacts             []*entity.Contact `protobuf:"bytes,5,rep,name=contacts" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EntityData) Reset()         { *m = EntityData{} }
func (m *EntityData) String() string { return proto.CompactTextString(m) }
func (*EntityData) ProtoMessage()    {}
func (*EntityData) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_45431f28cb32fa51, []int{6}
}
func (m *EntityData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityData.Unmarshal(m, b)
}
func (m *EntityData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityData.Marshal(b, m, deterministic)
}
func (dst *EntityData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityData.Merge(dst, src)
}
func (m *EntityData) XXX_Size() int {
	return xxx_messageInfo_EntityData.Size(m)
}
func (m *EntityData) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityData.DiscardUnknown(m)
}

var xxx_messageInfo_EntityData proto.InternalMessageInfo

func (m *EntityData) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *EntityData) GetLegalName() string {
	if m != nil {
		return m.LegalName
	}
	return ""
}

func (m *EntityData) GetAddresses() []*entity.Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *EntityData) GetPaymentDetails() []*entity.PaymentDetail {
	if m != nil {
		return m.PaymentDetails
	}
	return nil
}

func (m *EntityData) GetContacts() []*entity.Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "entity.GetRequest")
	proto.RegisterType((*GetVersionRequest)(nil), "entity.GetVersionRequest")
	proto.RegisterType((*EntityCreatePayload)(nil), "entity.EntityCreatePayload")
	proto.RegisterType((*EntityUpdatePayload)(nil), "entity.EntityUpdatePayload")
	proto.RegisterType((*EntityResponse)(nil), "entity.EntityResponse")
	proto.RegisterType((*ResponseHeader)(nil), "entity.ResponseHeader")
	proto.RegisterType((*EntityData)(nil), "entity.EntityData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DocumentService service

type DocumentServiceClient interface {
	Create(ctx context.Context, in *EntityCreatePayload, opts ...grpc.CallOption) (*EntityResponse, error)
	Update(ctx context.Context, in *EntityUpdatePayload, opts ...grpc.CallOption) (*EntityResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*EntityResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*EntityResponse, error)
}

type documentServiceClient struct {
	cc *grpc.ClientConn
}

func NewDocumentServiceClient(cc *grpc.ClientConn) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) Create(ctx context.Context, in *EntityCreatePayload, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := grpc.Invoke(ctx, "/entity.DocumentService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Update(ctx context.Context, in *EntityUpdatePayload, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := grpc.Invoke(ctx, "/entity.DocumentService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := grpc.Invoke(ctx, "/entity.DocumentService/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := grpc.Invoke(ctx, "/entity.DocumentService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DocumentService service

type DocumentServiceServer interface {
	Create(context.Context, *EntityCreatePayload) (*EntityResponse, error)
	Update(context.Context, *EntityUpdatePayload) (*EntityResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*EntityResponse, error)
	Get(context.Context, *GetRequest) (*EntityResponse, error)
}

func RegisterDocumentServiceServer(s *grpc.Server, srv DocumentServiceServer) {
	s.RegisterService(&_DocumentService_serviceDesc, srv)
}

func _DocumentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityCreatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.DocumentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Create(ctx, req.(*EntityCreatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityUpdatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.DocumentService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Update(ctx, req.(*EntityUpdatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.DocumentService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.DocumentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DocumentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "entity.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DocumentService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DocumentService_Update_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _DocumentService_GetVersion_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DocumentService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entity/service.proto",
}

func init() { proto.RegisterFile("entity/service.proto", fileDescriptor_service_45431f28cb32fa51) }

var fileDescriptor_service_45431f28cb32fa51 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x4e, 0x13, 0x4d,
	0x14, 0xce, 0x52, 0x28, 0xf4, 0x14, 0x4a, 0x18, 0xf8, 0x9b, 0xb2, 0xbf, 0xe2, 0x66, 0x15, 0xd2,
	0x28, 0xd0, 0xa4, 0xc6, 0x98, 0x78, 0x61, 0x52, 0xc0, 0x54, 0x2e, 0x24, 0xcd, 0x1a, 0xbd, 0xf0,
	0x86, 0x0c, 0xbb, 0x87, 0x65, 0xcd, 0x76, 0x67, 0xdd, 0x39, 0x60, 0x1a, 0xc2, 0x8d, 0x3e, 0x80,
	0x09, 0x3e, 0x8a, 0x8f, 0xe2, 0x23, 0xe8, 0xb5, 0xcf, 0x60, 0x3a, 0x33, 0x0b, 0x5d, 0x0a, 0xc2,
	0xd5, 0x66, 0xbe, 0xef, 0xcc, 0xf9, 0xbe, 0x73, 0xce, 0xce, 0x81, 0x25, 0x4c, 0x28, 0xa2, 0x41,
	0x4b, 0x62, 0x76, 0x12, 0xf9, 0xb8, 0x99, 0x66, 0x82, 0x04, 0x2b, 0x6b, 0xd4, 0x5e, 0x34, 0xac,
	0xfe, 0x68, 0xd2, 0xbe, 0x17, 0x0a, 0x11, 0xc6, 0xd8, 0xe2, 0x69, 0xd4, 0xe2, 0x49, 0x22, 0x88,
	0x53, 0x24, 0x12, 0x69, 0xd8, 0x75, 0xf5, 0xf1, 0x37, 0x42, 0x4c, 0x36, 0xe4, 0x67, 0x1e, 0x86,
	0x98, 0xb5, 0x44, 0xaa, 0x22, 0xc6, 0xa3, 0xdd, 0x75, 0x80, 0x2e, 0x92, 0x87, 0x9f, 0x8e, 0x51,
	0x12, 0x5b, 0x01, 0x88, 0x82, 0xa1, 0xd6, 0x61, 0x84, 0x59, 0xc3, 0x72, 0xac, 0x66, 0xc5, 0x1b,
	0x41, 0xdc, 0x37, 0xb0, 0xd0, 0x45, 0x7a, 0x8f, 0x99, 0x8c, 0x44, 0x72, 0xc7, 0x4b, 0xac, 0x01,
	0xd3, 0x27, 0xfa, 0x46, 0x63, 0x42, 0x91, 0xf9, 0xd1, 0xf5, 0x61, 0xf1, 0x95, 0x2a, 0x6c, 0x3b,
	0x43, 0x4e, 0xd8, 0xe3, 0x83, 0x58, 0xf0, 0x80, 0x3d, 0x82, 0x39, 0x5f, 0xc4, 0x31, 0x3f, 0x10,
	0x19, 0x27, 0x91, 0xc9, 0x86, 0xe5, 0x94, 0x9a, 0x15, 0xaf, 0x08, 0xb2, 0x35, 0x98, 0x0c, 0x38,
	0x71, 0x95, 0xb3, 0xda, 0x66, 0x9b, 0xa6, 0x45, 0x3a, 0xe1, 0x0e, 0x27, 0xee, 0x29, 0xde, 0xfd,
	0x6a, 0xe5, 0x2a, 0xef, 0xd2, 0x60, 0x44, 0xe5, 0x36, 0xdb, 0x63, 0x2e, 0x26, 0xfe, 0xe5, 0xa2,
	0x74, 0x8b, 0x8b, 0x23, 0xa8, 0x69, 0xcc, 0x43, 0x99, 0x8a, 0x44, 0x22, 0xdb, 0x84, 0xf2, 0x11,
	0xf2, 0xc0, 0x68, 0x57, 0xdb, 0xf5, 0xfc, 0x6e, 0x1e, 0xf1, 0x5a, 0xb1, 0x9e, 0x89, 0xba, 0x73,
	0xbd, 0x3f, 0x2c, 0xa8, 0x15, 0x53, 0xb0, 0x07, 0x50, 0x0d, 0x84, 0x7f, 0xdc, 0xc7, 0x84, 0xf6,
	0xa3, 0x20, 0xaf, 0x35, 0x87, 0x76, 0x03, 0x76, 0x1f, 0xc0, 0xcc, 0x64, 0xc8, 0xeb, 0x29, 0x55,
	0x0c, 0xb2, 0x1b, 0xb0, 0x25, 0x98, 0x92, 0xc4, 0x09, 0x55, 0x95, 0x15, 0x4f, 0x1f, 0xc6, 0x1b,
	0x34, 0x79, 0x5d, 0x83, 0x56, 0xa1, 0x46, 0x19, 0x4f, 0x24, 0xf7, 0xc9, 0xa4, 0x9f, 0x52, 0x49,
	0xe6, 0x46, 0xd0, 0xdd, 0xc0, 0xfd, 0x65, 0x01, 0x5c, 0x96, 0xc2, 0x6c, 0x98, 0xd1, 0xa3, 0xa0,
	0x81, 0xb2, 0x3b, 0xeb, 0x5d, 0x9c, 0x87, 0x66, 0x63, 0x0c, 0x79, 0xbc, 0x9f, 0xf0, 0x3e, 0xe6,
	0x66, 0x15, 0xb2, 0xc7, 0xfb, 0xc8, 0x36, 0xa0, 0xc2, 0x83, 0x20, 0x43, 0x29, 0x51, 0x36, 0x4a,
	0x4e, 0xa9, 0x59, 0x6d, 0xcf, 0xe7, 0xcd, 0xea, 0x68, 0xc2, 0xbb, 0x8c, 0x60, 0x2f, 0x61, 0x3e,
	0xe5, 0x03, 0xd5, 0x9a, 0x00, 0x89, 0x47, 0xb1, 0xae, 0xa3, 0xda, 0xfe, 0x2f, 0xbf, 0xd4, 0xd3,
	0xf4, 0x8e, 0x62, 0xbd, 0x5a, 0x3a, 0x7a, 0x94, 0xec, 0x09, 0xcc, 0xf8, 0x22, 0x21, 0xee, 0x93,
	0x6c, 0x4c, 0x15, 0xd5, 0xb6, 0x35, 0xee, 0x5d, 0x04, 0xb4, 0xff, 0x94, 0x60, 0x7e, 0xc7, 0xb4,
	0xfd, 0xad, 0x7e, 0xf0, 0x2c, 0x84, 0xb2, 0xfe, 0xfd, 0xd9, 0xff, 0xc5, 0x99, 0x16, 0x1e, 0x85,
	0x5d, 0x2f, 0x92, 0xf9, 0x84, 0xdd, 0xe6, 0x79, 0x67, 0xd1, 0x5e, 0xd0, 0xb1, 0xd2, 0xe1, 0x89,
	0xa3, 0xc3, 0xbe, 0xfc, 0xfc, 0xfd, 0x7d, 0x62, 0xd6, 0x9d, 0x36, 0x9b, 0xe3, 0x85, 0xf5, 0x98,
	0x11, 0x94, 0xf5, 0x0b, 0xb8, 0x2a, 0x54, 0x78, 0x17, 0x37, 0x0a, 0x3d, 0x53, 0x42, 0x3a, 0xf6,
	0xaa, 0xd0, 0xb2, 0xbd, 0x64, 0x84, 0x5a, 0xa7, 0x97, 0x4f, 0xe8, 0x6c, 0xa8, 0xfa, 0xcd, 0x52,
	0x1b, 0xc6, 0xec, 0x0c, 0xb6, 0x9c, 0x67, 0x1f, 0xdb, 0x23, 0x37, 0x0a, 0xef, 0x9d, 0x77, 0x56,
	0xed, 0x87, 0x5d, 0x24, 0x87, 0x3b, 0x32, 0x45, 0x3f, 0x3a, 0x8c, 0x7c, 0xc7, 0xfc, 0x9c, 0x8e,
	0x38, 0xbc, 0x62, 0xc5, 0x61, 0x2b, 0xd7, 0x59, 0x69, 0x9d, 0x9a, 0x1b, 0x67, 0xec, 0x23, 0x94,
	0xba, 0x48, 0x8c, 0x8d, 0x38, 0xb9, 0xcd, 0xc2, 0xf3, 0xf3, 0x4e, 0xc3, 0xae, 0x0f, 0x2d, 0xd0,
	0x11, 0x3a, 0xfe, 0x71, 0x96, 0x61, 0x42, 0xa3, 0xaa, 0x75, 0x76, 0x6d, 0x03, 0xb6, 0xd6, 0x00,
	0x7c, 0xd1, 0x37, 0x59, 0xb7, 0x66, 0xcd, 0xcc, 0x7b, 0xc3, 0xd5, 0xdb, 0xb3, 0x3e, 0xcc, 0x68,
	0x3c, 0x3d, 0x38, 0x28, 0xab, 0x6d, 0xfc, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x8d,
	0xab, 0xf8, 0x0e, 0x06, 0x00, 0x00,
}
