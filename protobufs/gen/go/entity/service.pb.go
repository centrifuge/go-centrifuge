// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity/service.proto

package entitypb

import (
	context "context"
	fmt "fmt"
	entity "github.com/centrifuge/centrifuge-protobufs/gen/go/entity"
	document "github.com/centrifuge/go-centrifuge/protobufs/gen/go/document"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{0}
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

func (m *GetRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type GetVersionRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{1}
}

func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionRequest.Unmarshal(m, b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
}
func (m *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(m, src)
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
	ReadAccess           *document.ReadAccess  `protobuf:"bytes,1,opt,name=read_access,json=readAccess,proto3" json:"read_access,omitempty"`
	WriteAccess          *document.WriteAccess `protobuf:"bytes,2,opt,name=write_access,json=writeAccess,proto3" json:"write_access,omitempty"`
	Data                 *EntityData           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *EntityCreatePayload) Reset()         { *m = EntityCreatePayload{} }
func (m *EntityCreatePayload) String() string { return proto.CompactTextString(m) }
func (*EntityCreatePayload) ProtoMessage()    {}
func (*EntityCreatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{2}
}

func (m *EntityCreatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityCreatePayload.Unmarshal(m, b)
}
func (m *EntityCreatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityCreatePayload.Marshal(b, m, deterministic)
}
func (m *EntityCreatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityCreatePayload.Merge(m, src)
}
func (m *EntityCreatePayload) XXX_Size() int {
	return xxx_messageInfo_EntityCreatePayload.Size(m)
}
func (m *EntityCreatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityCreatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_EntityCreatePayload proto.InternalMessageInfo

func (m *EntityCreatePayload) GetReadAccess() *document.ReadAccess {
	if m != nil {
		return m.ReadAccess
	}
	return nil
}

func (m *EntityCreatePayload) GetWriteAccess() *document.WriteAccess {
	if m != nil {
		return m.WriteAccess
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
	Identifier           string                `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	ReadAccess           *document.ReadAccess  `protobuf:"bytes,2,opt,name=read_access,json=readAccess,proto3" json:"read_access,omitempty"`
	WriteAccess          *document.WriteAccess `protobuf:"bytes,3,opt,name=write_access,json=writeAccess,proto3" json:"write_access,omitempty"`
	Data                 *EntityData           `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *EntityUpdatePayload) Reset()         { *m = EntityUpdatePayload{} }
func (m *EntityUpdatePayload) String() string { return proto.CompactTextString(m) }
func (*EntityUpdatePayload) ProtoMessage()    {}
func (*EntityUpdatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{3}
}

func (m *EntityUpdatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityUpdatePayload.Unmarshal(m, b)
}
func (m *EntityUpdatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityUpdatePayload.Marshal(b, m, deterministic)
}
func (m *EntityUpdatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityUpdatePayload.Merge(m, src)
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

func (m *EntityUpdatePayload) GetReadAccess() *document.ReadAccess {
	if m != nil {
		return m.ReadAccess
	}
	return nil
}

func (m *EntityUpdatePayload) GetWriteAccess() *document.WriteAccess {
	if m != nil {
		return m.WriteAccess
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
	Header               *document.ResponseHeader  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *EntityWithReleationships `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EntityResponse) Reset()         { *m = EntityResponse{} }
func (m *EntityResponse) String() string { return proto.CompactTextString(m) }
func (*EntityResponse) ProtoMessage()    {}
func (*EntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{4}
}

func (m *EntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityResponse.Unmarshal(m, b)
}
func (m *EntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityResponse.Marshal(b, m, deterministic)
}
func (m *EntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityResponse.Merge(m, src)
}
func (m *EntityResponse) XXX_Size() int {
	return xxx_messageInfo_EntityResponse.Size(m)
}
func (m *EntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntityResponse proto.InternalMessageInfo

func (m *EntityResponse) GetHeader() *document.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *EntityResponse) GetData() *EntityWithReleationships {
	if m != nil {
		return m.Data
	}
	return nil
}

type Relationship struct {
	Identity                  string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Active                    bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	LastUpdateAtEntityVersion string   `protobuf:"bytes,3,opt,name=last_update_at_entity_version,json=lastUpdateAtEntityVersion,proto3" json:"last_update_at_entity_version,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Relationship) Reset()         { *m = Relationship{} }
func (m *Relationship) String() string { return proto.CompactTextString(m) }
func (*Relationship) ProtoMessage()    {}
func (*Relationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{5}
}

func (m *Relationship) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Relationship.Unmarshal(m, b)
}
func (m *Relationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Relationship.Marshal(b, m, deterministic)
}
func (m *Relationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relationship.Merge(m, src)
}
func (m *Relationship) XXX_Size() int {
	return xxx_messageInfo_Relationship.Size(m)
}
func (m *Relationship) XXX_DiscardUnknown() {
	xxx_messageInfo_Relationship.DiscardUnknown(m)
}

var xxx_messageInfo_Relationship proto.InternalMessageInfo

func (m *Relationship) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *Relationship) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Relationship) GetLastUpdateAtEntityVersion() string {
	if m != nil {
		return m.LastUpdateAtEntityVersion
	}
	return ""
}

// EntityData is the default entity schema
type EntityData struct {
	Identity  string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	LegalName string `protobuf:"bytes,2,opt,name=legal_name,json=legalName,proto3" json:"legal_name,omitempty"`
	// address
	Addresses []*entity.Address `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// tax information
	PaymentDetails []*entity.PaymentDetail `protobuf:"bytes,4,rep,name=payment_details,json=paymentDetails,proto3" json:"payment_details,omitempty"`
	// Entity contact list
	Contacts             []*entity.Contact `protobuf:"bytes,5,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EntityData) Reset()         { *m = EntityData{} }
func (m *EntityData) String() string { return proto.CompactTextString(m) }
func (*EntityData) ProtoMessage()    {}
func (*EntityData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{6}
}

func (m *EntityData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityData.Unmarshal(m, b)
}
func (m *EntityData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityData.Marshal(b, m, deterministic)
}
func (m *EntityData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityData.Merge(m, src)
}
func (m *EntityData) XXX_Size() int {
	return xxx_messageInfo_EntityData.Size(m)
}
func (m *EntityData) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityData.DiscardUnknown(m)
}

var xxx_messageInfo_EntityData proto.InternalMessageInfo

func (m *EntityData) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
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

// Entity Relationships
type EntityWithReleationships struct {
	Entity               *EntityData     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Relationships        []*Relationship `protobuf:"bytes,6,rep,name=relationships,proto3" json:"relationships,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EntityWithReleationships) Reset()         { *m = EntityWithReleationships{} }
func (m *EntityWithReleationships) String() string { return proto.CompactTextString(m) }
func (*EntityWithReleationships) ProtoMessage()    {}
func (*EntityWithReleationships) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{7}
}

func (m *EntityWithReleationships) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityWithReleationships.Unmarshal(m, b)
}
func (m *EntityWithReleationships) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityWithReleationships.Marshal(b, m, deterministic)
}
func (m *EntityWithReleationships) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityWithReleationships.Merge(m, src)
}
func (m *EntityWithReleationships) XXX_Size() int {
	return xxx_messageInfo_EntityWithReleationships.Size(m)
}
func (m *EntityWithReleationships) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityWithReleationships.DiscardUnknown(m)
}

var xxx_messageInfo_EntityWithReleationships proto.InternalMessageInfo

func (m *EntityWithReleationships) GetEntity() *EntityData {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *EntityWithReleationships) GetRelationships() []*Relationship {
	if m != nil {
		return m.Relationships
	}
	return nil
}

type EntityRelationPayload struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	TargetIdentity       string   `protobuf:"bytes,2,opt,name=target_identity,json=targetIdentity,proto3" json:"target_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityRelationPayload) Reset()         { *m = EntityRelationPayload{} }
func (m *EntityRelationPayload) String() string { return proto.CompactTextString(m) }
func (*EntityRelationPayload) ProtoMessage()    {}
func (*EntityRelationPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{8}
}

func (m *EntityRelationPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityRelationPayload.Unmarshal(m, b)
}
func (m *EntityRelationPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityRelationPayload.Marshal(b, m, deterministic)
}
func (m *EntityRelationPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityRelationPayload.Merge(m, src)
}
func (m *EntityRelationPayload) XXX_Size() int {
	return xxx_messageInfo_EntityRelationPayload.Size(m)
}
func (m *EntityRelationPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityRelationPayload.DiscardUnknown(m)
}

var xxx_messageInfo_EntityRelationPayload proto.InternalMessageInfo

func (m *EntityRelationPayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *EntityRelationPayload) GetTargetIdentity() string {
	if m != nil {
		return m.TargetIdentity
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "entity.GetRequest")
	proto.RegisterType((*GetVersionRequest)(nil), "entity.GetVersionRequest")
	proto.RegisterType((*EntityCreatePayload)(nil), "entity.EntityCreatePayload")
	proto.RegisterType((*EntityUpdatePayload)(nil), "entity.EntityUpdatePayload")
	proto.RegisterType((*EntityResponse)(nil), "entity.EntityResponse")
	proto.RegisterType((*Relationship)(nil), "entity.Relationship")
	proto.RegisterType((*EntityData)(nil), "entity.EntityData")
	proto.RegisterType((*EntityWithReleationships)(nil), "entity.EntityWithReleationships")
	proto.RegisterType((*EntityRelationPayload)(nil), "entity.EntityRelationPayload")
}

func init() { proto.RegisterFile("entity/service.proto", fileDescriptor_c1b437217b9e14a2) }

var fileDescriptor_c1b437217b9e14a2 = []byte{
	// 877 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0x96, 0xb3, 0xa9, 0x9b, 0x9c, 0x0d, 0x89, 0x3a, 0x49, 0x56, 0x8e, 0x4b, 0x82, 0x65, 0xd4,
	0x36, 0x2a, 0x4d, 0x16, 0x05, 0x2a, 0x50, 0x2f, 0x10, 0x6e, 0x83, 0x42, 0x2f, 0x28, 0x91, 0x2b,
	0xa8, 0xc4, 0x8d, 0x99, 0xda, 0x27, 0x5e, 0x83, 0xe3, 0x31, 0x33, 0x93, 0x84, 0x55, 0xd5, 0x9b,
	0x8a, 0x2b, 0x24, 0x84, 0xb4, 0xbc, 0x07, 0x8f, 0xc2, 0x0d, 0x8f, 0x00, 0x0f, 0x82, 0x3c, 0x33,
	0xde, 0xb5, 0x77, 0xb3, 0xa4, 0xed, 0x95, 0x33, 0xe7, 0x9c, 0x6f, 0xbe, 0xef, 0xfc, 0xcc, 0xc9,
	0xc2, 0x06, 0x16, 0x32, 0x93, 0xc3, 0xbe, 0x40, 0x7e, 0x9e, 0xc5, 0xb8, 0x5f, 0x72, 0x26, 0x19,
	0xb1, 0xb5, 0xd5, 0xed, 0x25, 0x2c, 0x3e, 0x3b, 0xc5, 0x42, 0xb6, 0xfd, 0xee, 0xba, 0x41, 0xe9,
	0x8f, 0x31, 0xbe, 0x9b, 0x32, 0x96, 0xe6, 0xd8, 0xa7, 0x65, 0xd6, 0xa7, 0x45, 0xc1, 0x24, 0x95,
	0x19, 0x2b, 0x84, 0xf1, 0xde, 0x53, 0x9f, 0x78, 0x2f, 0xc5, 0x62, 0x4f, 0x5c, 0xd0, 0x34, 0x45,
	0xde, 0x67, 0xa5, 0x8a, 0x98, 0x8d, 0xf6, 0xef, 0x01, 0x1c, 0xa1, 0x0c, 0xf1, 0xa7, 0x33, 0x14,
	0x92, 0xec, 0x00, 0x64, 0x49, 0xc5, 0x75, 0x92, 0x21, 0x77, 0x2c, 0xcf, 0xda, 0x5d, 0x0e, 0x1b,
	0x16, 0xff, 0x2b, 0xb8, 0x71, 0x84, 0xf2, 0x5b, 0xe4, 0x22, 0x63, 0xc5, 0x6b, 0x82, 0x88, 0x03,
	0xd7, 0xcf, 0x35, 0xc2, 0x59, 0x50, 0xce, 0xfa, 0xe8, 0xff, 0x69, 0xc1, 0xfa, 0x17, 0x2a, 0xb3,
	0x47, 0x1c, 0xa9, 0xc4, 0x63, 0x3a, 0xcc, 0x19, 0x4d, 0xc8, 0x7d, 0xe8, 0x72, 0xa4, 0x49, 0x44,
	0xe3, 0x18, 0x85, 0x50, 0x57, 0x76, 0x0f, 0x36, 0xf6, 0xeb, 0x1a, 0xed, 0x87, 0x48, 0x93, 0x40,
	0xf9, 0x42, 0xe0, 0xe3, 0xbf, 0xc9, 0xa7, 0xb0, 0x72, 0xc1, 0x33, 0x89, 0x35, 0x6e, 0x41, 0xe1,
	0x36, 0x27, 0xb8, 0x67, 0x95, 0xd7, 0x00, 0xbb, 0x17, 0x93, 0x03, 0xb9, 0x0d, 0x8b, 0x09, 0x95,
	0xd4, 0xe9, 0x28, 0x04, 0xd9, 0x37, 0xe5, 0xd6, 0xda, 0x0e, 0xa9, 0xa4, 0xa1, 0xf2, 0xfb, 0x7f,
	0x8d, 0x05, 0x7f, 0x53, 0x26, 0x0d, 0xc1, 0x57, 0x95, 0x60, 0x2a, 0xa1, 0x85, 0xb7, 0x4c, 0xa8,
	0xf3, 0xc6, 0x09, 0x2d, 0x5e, 0x91, 0xd0, 0xcf, 0xb0, 0xaa, 0x6d, 0x21, 0x8a, 0x92, 0x15, 0x02,
	0xc9, 0x87, 0x60, 0x0f, 0x90, 0x26, 0x26, 0x8d, 0xee, 0x81, 0xd3, 0x54, 0xa9, 0x63, 0xbe, 0x54,
	0xfe, 0xd0, 0xc4, 0x91, 0x8f, 0x0d, 0x97, 0xce, 0xca, 0x6b, 0x73, 0x3d, 0xcb, 0xe4, 0x20, 0xc4,
	0x1c, 0xf5, 0xd0, 0x0d, 0xb2, 0x52, 0x18, 0xe6, 0x5f, 0x2c, 0x58, 0x09, 0x31, 0x1f, 0xdb, 0x89,
	0x0b, 0x4b, 0xba, 0x62, 0x72, 0x68, 0x2a, 0x38, 0x3e, 0x93, 0x1e, 0xd8, 0x34, 0x96, 0xd9, 0x39,
	0x2a, 0x92, 0xa5, 0xd0, 0x9c, 0xc8, 0xe7, 0xb0, 0x9d, 0x53, 0x21, 0xa3, 0x33, 0xd5, 0x8d, 0x88,
	0xca, 0x48, 0x03, 0xa2, 0x7a, 0xe0, 0x3a, 0xea, 0xa2, 0xad, 0x2a, 0x48, 0x77, 0x2c, 0x90, 0x5a,
	0x97, 0x99, 0x61, 0xff, 0x1f, 0x0b, 0x60, 0x52, 0x95, 0xff, 0x15, 0xb1, 0x0d, 0x90, 0x63, 0x4a,
	0xf3, 0xa8, 0xa0, 0xa7, 0x68, 0x46, 0x79, 0x59, 0x59, 0x9e, 0xd0, 0x53, 0x24, 0x7b, 0xb0, 0x4c,
	0x93, 0x84, 0xa3, 0x10, 0x58, 0x75, 0xaa, 0xb3, 0xdb, 0x3d, 0x58, 0xab, 0x6b, 0x11, 0x68, 0x47,
	0x38, 0x89, 0x20, 0x9f, 0xc1, 0x5a, 0x49, 0x87, 0x55, 0x5d, 0xa3, 0x04, 0x25, 0xcd, 0x72, 0xe1,
	0x2c, 0x2a, 0xd0, 0x66, 0x0d, 0x3a, 0xd6, 0xee, 0x43, 0xe5, 0x0d, 0x57, 0xcb, 0xe6, 0x51, 0x90,
	0x0f, 0x60, 0x29, 0x66, 0x85, 0xa4, 0xb1, 0x14, 0xce, 0xb5, 0x36, 0xdb, 0x23, 0x6d, 0x0f, 0xc7,
	0x01, 0xfe, 0x2b, 0x0b, 0x9c, 0x79, 0xfd, 0x20, 0x77, 0xc1, 0x6e, 0x64, 0x7c, 0xf9, 0xb4, 0x98,
	0x08, 0xf2, 0x00, 0xde, 0xe1, 0x8d, 0xa6, 0x09, 0xc7, 0x56, 0xd4, 0x1b, 0x35, 0xa4, 0xd9, 0xd1,
	0xb0, 0x1d, 0xea, 0x7f, 0x0f, 0x9b, 0xf5, 0xac, 0x69, 0xf3, 0xeb, 0xbe, 0x9e, 0x3b, 0xb0, 0x26,
	0x29, 0x4f, 0x51, 0x46, 0xe3, 0xde, 0xe8, 0xea, 0xaf, 0x6a, 0xf3, 0x63, 0x63, 0x3d, 0xf8, 0xcd,
	0x86, 0xb5, 0x43, 0x33, 0xad, 0x4f, 0xf5, 0x1e, 0x25, 0x29, 0xd8, 0x7a, 0xb9, 0x90, 0x9b, 0xed,
	0xbc, 0x5a, 0x2b, 0xc7, 0xed, 0xb5, 0x9d, 0xf5, 0xa8, 0xfb, 0xbb, 0xa3, 0x60, 0xdd, 0xbd, 0xa1,
	0x63, 0x85, 0x47, 0x0b, 0x4f, 0x87, 0xbd, 0xfa, 0xfb, 0xdf, 0x3f, 0x16, 0x56, 0xfc, 0xeb, 0x66,
	0x31, 0x3f, 0xb0, 0xee, 0x12, 0x09, 0xb6, 0x1e, 0xb1, 0x69, 0xa2, 0xd6, 0xaa, 0x98, 0x4b, 0x74,
	0x5f, 0x11, 0xe9, 0xd8, 0x69, 0xa2, 0x2d, 0x77, 0xc3, 0x10, 0xf5, 0x5f, 0x4c, 0xea, 0xf2, 0xb2,
	0x62, 0xfd, 0xdd, 0x52, 0x0b, 0xdc, 0x8c, 0x33, 0xd9, 0xaa, 0x6f, 0x9f, 0x59, 0xd3, 0x73, 0x89,
	0x9f, 0x8c, 0x82, 0x5b, 0xee, 0xfb, 0x47, 0x28, 0x3d, 0xea, 0x89, 0x12, 0xe3, 0xec, 0x24, 0x8b,
	0x3d, 0xf3, 0x7e, 0x3c, 0x76, 0x32, 0x25, 0xc5, 0x23, 0x3b, 0x97, 0x49, 0xe9, 0xbf, 0x30, 0x88,
	0x97, 0xe4, 0x07, 0xe8, 0x1c, 0xa1, 0x24, 0xa4, 0xa1, 0xe4, 0x2a, 0x09, 0x9f, 0x8c, 0x02, 0xc7,
	0xed, 0x55, 0x12, 0xe4, 0x00, 0xbd, 0xf8, 0x8c, 0x73, 0x2c, 0x64, 0x93, 0xb5, 0x47, 0x2e, 0x2d,
	0x40, 0x95, 0xfd, 0xb5, 0xa7, 0x03, 0xca, 0x91, 0x6c, 0x4f, 0x5f, 0xdd, 0x1a, 0xb1, 0xb9, 0xcc,
	0x5f, 0x8f, 0x82, 0x3b, 0xee, 0x2d, 0x75, 0x85, 0xe2, 0xd6, 0x51, 0x5e, 0xbd, 0xf2, 0xbc, 0x8b,
	0x4c, 0x0e, 0x3c, 0x26, 0x07, 0xc8, 0x85, 0x12, 0xf2, 0x9e, 0xef, 0xd6, 0x42, 0x44, 0x05, 0x9a,
	0xe9, 0xc7, 0xaf, 0x16, 0xd8, 0x21, 0x9e, 0xb3, 0x1f, 0xdf, 0x5a, 0xd2, 0xe3, 0x51, 0xe0, 0xb9,
	0x3b, 0x5c, 0xdd, 0xe1, 0xd1, 0x19, 0x45, 0x8a, 0x56, 0xb7, 0xc2, 0xbf, 0x59, 0x6b, 0xd1, 0xd1,
	0xd3, 0x62, 0x1e, 0xde, 0x06, 0x88, 0xd9, 0xa9, 0xe1, 0x79, 0xb8, 0x62, 0x9e, 0xc4, 0x71, 0xf5,
	0x8f, 0xff, 0xd8, 0xfa, 0x6e, 0x49, 0xdb, 0xcb, 0xe7, 0xcf, 0x6d, 0xf5, 0x5b, 0xe0, 0xa3, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x67, 0x1c, 0x3f, 0xd6, 0xa4, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DocumentServiceClient interface {
	Create(ctx context.Context, in *EntityCreatePayload, opts ...grpc.CallOption) (*EntityResponse, error)
	Update(ctx context.Context, in *EntityUpdatePayload, opts ...grpc.CallOption) (*EntityResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*EntityResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*EntityResponse, error)
	// Entity Relation Share
	Share(ctx context.Context, in *EntityRelationPayload, opts ...grpc.CallOption) (*EntityResponse, error)
	// Entity Relation Revoke
	Revoke(ctx context.Context, in *EntityRelationPayload, opts ...grpc.CallOption) (*EntityResponse, error)
}

type documentServiceClient struct {
	cc *grpc.ClientConn
}

func NewDocumentServiceClient(cc *grpc.ClientConn) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) Create(ctx context.Context, in *EntityCreatePayload, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.DocumentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Update(ctx context.Context, in *EntityUpdatePayload, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.DocumentService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.DocumentService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.DocumentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Share(ctx context.Context, in *EntityRelationPayload, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.DocumentService/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Revoke(ctx context.Context, in *EntityRelationPayload, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.DocumentService/Revoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServiceServer is the server API for DocumentService service.
type DocumentServiceServer interface {
	Create(context.Context, *EntityCreatePayload) (*EntityResponse, error)
	Update(context.Context, *EntityUpdatePayload) (*EntityResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*EntityResponse, error)
	Get(context.Context, *GetRequest) (*EntityResponse, error)
	// Entity Relation Share
	Share(context.Context, *EntityRelationPayload) (*EntityResponse, error)
	// Entity Relation Revoke
	Revoke(context.Context, *EntityRelationPayload) (*EntityResponse, error)
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

func _DocumentService_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityRelationPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.DocumentService/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Share(ctx, req.(*EntityRelationPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityRelationPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.DocumentService/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Revoke(ctx, req.(*EntityRelationPayload))
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
		{
			MethodName: "Share",
			Handler:    _DocumentService_Share_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _DocumentService_Revoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entity/service.proto",
}
