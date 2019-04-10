// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity/service.proto

package entitypb

import (
	context "context"
	fmt "fmt"
	math "math"

	entity "github.com/centrifuge/centrifuge-protobufs/gen/go/entity"
	document "github.com/centrifuge/go-centrifuge/protobufs/gen/go/document"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetRequestRelationship struct {
	Identifier             string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	RelationshipIdentifier string   `protobuf:"bytes,2,opt,name=relationship_identifier,json=relationshipIdentifier,proto3" json:"relationship_identifier,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *GetRequestRelationship) Reset()         { *m = GetRequestRelationship{} }
func (m *GetRequestRelationship) String() string { return proto.CompactTextString(m) }
func (*GetRequestRelationship) ProtoMessage()    {}
func (*GetRequestRelationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{1}
}

func (m *GetRequestRelationship) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequestRelationship.Unmarshal(m, b)
}
func (m *GetRequestRelationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequestRelationship.Marshal(b, m, deterministic)
}
func (m *GetRequestRelationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequestRelationship.Merge(m, src)
}
func (m *GetRequestRelationship) XXX_Size() int {
	return xxx_messageInfo_GetRequestRelationship.Size(m)
}
func (m *GetRequestRelationship) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequestRelationship.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequestRelationship proto.InternalMessageInfo

func (m *GetRequestRelationship) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *GetRequestRelationship) GetRelationshipIdentifier() string {
	if m != nil {
		return m.RelationshipIdentifier
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
	return fileDescriptor_c1b437217b9e14a2, []int{2}
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
	return fileDescriptor_c1b437217b9e14a2, []int{3}
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
	return fileDescriptor_c1b437217b9e14a2, []int{4}
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
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *EntityDataResponse      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *EntityResponse) Reset()         { *m = EntityResponse{} }
func (m *EntityResponse) String() string { return proto.CompactTextString(m) }
func (*EntityResponse) ProtoMessage()    {}
func (*EntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{5}
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

func (m *EntityResponse) GetData() *EntityDataResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

type Relationship struct {
	Identity             string   `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Active               bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Relationship) Reset()         { *m = Relationship{} }
func (m *Relationship) String() string { return proto.CompactTextString(m) }
func (*Relationship) ProtoMessage()    {}
func (*Relationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{6}
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
	return fileDescriptor_c1b437217b9e14a2, []int{7}
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
type EntityDataResponse struct {
	Entity               *EntityData     `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Relationships        []*Relationship `protobuf:"bytes,2,rep,name=relationships,proto3" json:"relationships,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EntityDataResponse) Reset()         { *m = EntityDataResponse{} }
func (m *EntityDataResponse) String() string { return proto.CompactTextString(m) }
func (*EntityDataResponse) ProtoMessage()    {}
func (*EntityDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{8}
}

func (m *EntityDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityDataResponse.Unmarshal(m, b)
}
func (m *EntityDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityDataResponse.Marshal(b, m, deterministic)
}
func (m *EntityDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityDataResponse.Merge(m, src)
}
func (m *EntityDataResponse) XXX_Size() int {
	return xxx_messageInfo_EntityDataResponse.Size(m)
}
func (m *EntityDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntityDataResponse proto.InternalMessageInfo

func (m *EntityDataResponse) GetEntity() *EntityData {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *EntityDataResponse) GetRelationships() []*Relationship {
	if m != nil {
		return m.Relationships
	}
	return nil
}

type RelationshipPayload struct {
	// entity identifier
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	TargetIdentity       string   `protobuf:"bytes,2,opt,name=target_identity,json=targetIdentity,proto3" json:"target_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationshipPayload) Reset()         { *m = RelationshipPayload{} }
func (m *RelationshipPayload) String() string { return proto.CompactTextString(m) }
func (*RelationshipPayload) ProtoMessage()    {}
func (*RelationshipPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{9}
}

func (m *RelationshipPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationshipPayload.Unmarshal(m, b)
}
func (m *RelationshipPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationshipPayload.Marshal(b, m, deterministic)
}
func (m *RelationshipPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationshipPayload.Merge(m, src)
}
func (m *RelationshipPayload) XXX_Size() int {
	return xxx_messageInfo_RelationshipPayload.Size(m)
}
func (m *RelationshipPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationshipPayload.DiscardUnknown(m)
}

var xxx_messageInfo_RelationshipPayload proto.InternalMessageInfo

func (m *RelationshipPayload) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *RelationshipPayload) GetTargetIdentity() string {
	if m != nil {
		return m.TargetIdentity
	}
	return ""
}

type RelationshipData struct {
	// DID of relationship owner
	OwnerIdentity string `protobuf:"bytes,1,opt,name=owner_identity,json=ownerIdentity,proto3" json:"owner_identity,omitempty"`
	// DID of target identity
	TargetIdentity string `protobuf:"bytes,2,opt,name=target_identity,json=targetIdentity,proto3" json:"target_identity,omitempty"`
	// identifier of Entity whose data can be accessed via this relationship
	EntityIdentifier     string   `protobuf:"bytes,3,opt,name=entity_identifier,json=entityIdentifier,proto3" json:"entity_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationshipData) Reset()         { *m = RelationshipData{} }
func (m *RelationshipData) String() string { return proto.CompactTextString(m) }
func (*RelationshipData) ProtoMessage()    {}
func (*RelationshipData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{10}
}

func (m *RelationshipData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationshipData.Unmarshal(m, b)
}
func (m *RelationshipData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationshipData.Marshal(b, m, deterministic)
}
func (m *RelationshipData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationshipData.Merge(m, src)
}
func (m *RelationshipData) XXX_Size() int {
	return xxx_messageInfo_RelationshipData.Size(m)
}
func (m *RelationshipData) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationshipData.DiscardUnknown(m)
}

var xxx_messageInfo_RelationshipData proto.InternalMessageInfo

func (m *RelationshipData) GetOwnerIdentity() string {
	if m != nil {
		return m.OwnerIdentity
	}
	return ""
}

func (m *RelationshipData) GetTargetIdentity() string {
	if m != nil {
		return m.TargetIdentity
	}
	return ""
}

func (m *RelationshipData) GetEntityIdentifier() string {
	if m != nil {
		return m.EntityIdentifier
	}
	return ""
}

type RelationshipResponse struct {
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Relationship         *RelationshipData        `protobuf:"bytes,2,opt,name=relationship,proto3" json:"relationship,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RelationshipResponse) Reset()         { *m = RelationshipResponse{} }
func (m *RelationshipResponse) String() string { return proto.CompactTextString(m) }
func (*RelationshipResponse) ProtoMessage()    {}
func (*RelationshipResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{11}
}

func (m *RelationshipResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationshipResponse.Unmarshal(m, b)
}
func (m *RelationshipResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationshipResponse.Marshal(b, m, deterministic)
}
func (m *RelationshipResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationshipResponse.Merge(m, src)
}
func (m *RelationshipResponse) XXX_Size() int {
	return xxx_messageInfo_RelationshipResponse.Size(m)
}
func (m *RelationshipResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationshipResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RelationshipResponse proto.InternalMessageInfo

func (m *RelationshipResponse) GetHeader() *document.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RelationshipResponse) GetRelationship() *RelationshipData {
	if m != nil {
		return m.Relationship
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "entity.GetRequest")
	proto.RegisterType((*GetRequestRelationship)(nil), "entity.GetRequestRelationship")
	proto.RegisterType((*GetVersionRequest)(nil), "entity.GetVersionRequest")
	proto.RegisterType((*EntityCreatePayload)(nil), "entity.EntityCreatePayload")
	proto.RegisterType((*EntityUpdatePayload)(nil), "entity.EntityUpdatePayload")
	proto.RegisterType((*EntityResponse)(nil), "entity.EntityResponse")
	proto.RegisterType((*Relationship)(nil), "entity.Relationship")
	proto.RegisterType((*EntityData)(nil), "entity.EntityData")
	proto.RegisterType((*EntityDataResponse)(nil), "entity.EntityDataResponse")
	proto.RegisterType((*RelationshipPayload)(nil), "entity.RelationshipPayload")
	proto.RegisterType((*RelationshipData)(nil), "entity.RelationshipData")
	proto.RegisterType((*RelationshipResponse)(nil), "entity.RelationshipResponse")
}

func init() { proto.RegisterFile("entity/service.proto", fileDescriptor_c1b437217b9e14a2) }

var fileDescriptor_c1b437217b9e14a2 = []byte{
	// 982 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x51, 0x6e, 0xdb, 0x46,
	0x10, 0x05, 0x2d, 0x47, 0xb6, 0x47, 0x8a, 0x1d, 0xaf, 0x1d, 0x45, 0x61, 0xd2, 0x84, 0x61, 0xe0,
	0xc4, 0x70, 0x62, 0xa9, 0x70, 0x11, 0xa4, 0x30, 0x8a, 0x02, 0x72, 0x5c, 0xa8, 0x29, 0xd0, 0xd4,
	0x60, 0xd0, 0x16, 0xe8, 0x47, 0x85, 0x35, 0x39, 0x96, 0xd8, 0x4a, 0x24, 0xb3, 0xbb, 0xb6, 0x20,
	0xa4, 0xf9, 0xe9, 0x47, 0x7f, 0x03, 0x28, 0x27, 0xe8, 0x05, 0x7a, 0x80, 0x1e, 0xa2, 0x3f, 0x3d,
	0x42, 0x7b, 0x90, 0x82, 0xbb, 0x4b, 0x69, 0x29, 0xd3, 0xb1, 0xeb, 0x7c, 0x49, 0x3b, 0xf3, 0x66,
	0xde, 0xdb, 0x99, 0xd9, 0x5d, 0xc2, 0x3a, 0x46, 0x22, 0x14, 0xa3, 0x26, 0x47, 0x76, 0x12, 0xfa,
	0xd8, 0x48, 0x58, 0x2c, 0x62, 0x52, 0x56, 0x56, 0xbb, 0x16, 0xc4, 0xfe, 0xf1, 0x00, 0x23, 0x91,
	0xf7, 0xdb, 0x6b, 0x3a, 0x4a, 0xfd, 0x68, 0xe3, 0xed, 0x6e, 0x1c, 0x77, 0xfb, 0xd8, 0xa4, 0x49,
	0xd8, 0xa4, 0x51, 0x14, 0x0b, 0x2a, 0xc2, 0x38, 0xe2, 0xda, 0xfb, 0x58, 0xfe, 0xf8, 0xdb, 0x5d,
	0x8c, 0xb6, 0xf9, 0x90, 0x76, 0xbb, 0xc8, 0x9a, 0x71, 0x22, 0x11, 0xa7, 0xd1, 0xee, 0x63, 0x80,
	0x36, 0x0a, 0x0f, 0x5f, 0x1d, 0x23, 0x17, 0xe4, 0x0e, 0x40, 0x18, 0xa4, 0x5c, 0x47, 0x21, 0xb2,
	0xba, 0xe5, 0x58, 0x9b, 0x4b, 0x9e, 0x61, 0x71, 0x5f, 0x41, 0x6d, 0x8a, 0xf6, 0xb0, 0xaf, 0x52,
	0xf5, 0xc2, 0xe4, 0xbc, 0x48, 0xf2, 0x14, 0x6e, 0x30, 0x03, 0xdf, 0x31, 0xc0, 0x73, 0x12, 0x5c,
	0x33, 0xdd, 0xcf, 0xa7, 0x94, 0x5f, 0xc3, 0x6a, 0x1b, 0xc5, 0x77, 0xc8, 0x78, 0x18, 0x47, 0x17,
	0xd4, 0x49, 0xea, 0xb0, 0x70, 0xa2, 0x22, 0x74, 0xf6, 0x6c, 0xe9, 0xfe, 0x61, 0xc1, 0xda, 0x17,
	0xb2, 0x98, 0xcf, 0x18, 0x52, 0x81, 0x07, 0x74, 0xd4, 0x8f, 0x69, 0x40, 0x9e, 0x40, 0x85, 0x21,
	0x0d, 0x3a, 0xd4, 0xf7, 0x91, 0x73, 0x99, 0xb2, 0xb2, 0xb3, 0xde, 0xc8, 0xda, 0xd2, 0xf0, 0x90,
	0x06, 0x2d, 0xe9, 0xf3, 0x80, 0x4d, 0xfe, 0x93, 0x4f, 0xa1, 0x3a, 0x64, 0xa1, 0xc0, 0x2c, 0x6e,
	0x4e, 0xc6, 0x5d, 0x9f, 0xc6, 0x7d, 0x9f, 0x7a, 0x75, 0x60, 0x65, 0x38, 0x5d, 0x90, 0x07, 0x30,
	0x1f, 0x50, 0x41, 0xeb, 0x25, 0x19, 0x41, 0x1a, 0xba, 0xc3, 0x4a, 0xdb, 0x3e, 0x15, 0xd4, 0x93,
	0x7e, 0xf7, 0xaf, 0x89, 0xe0, 0x6f, 0x93, 0xc0, 0x10, 0x7c, 0x5e, 0x09, 0x66, 0x36, 0x34, 0x77,
	0xc9, 0x0d, 0x95, 0xfe, 0xf7, 0x86, 0xe6, 0xcf, 0xd9, 0x10, 0x83, 0x65, 0x65, 0xf3, 0x90, 0x27,
	0x71, 0xc4, 0x91, 0x7c, 0x0c, 0xe5, 0x1e, 0xd2, 0x40, 0x6f, 0xa3, 0xb2, 0x53, 0x37, 0x55, 0x2a,
	0xcc, 0x97, 0xd2, 0xef, 0x69, 0x1c, 0x69, 0x68, 0x2e, 0xb5, 0x2b, 0xbb, 0x80, 0x4b, 0xc7, 0x69,
	0xce, 0x3d, 0xa8, 0xe6, 0xa6, 0xd5, 0x86, 0x45, 0x55, 0x2a, 0x31, 0xd2, 0xa5, 0x9b, 0xac, 0x49,
	0x0d, 0xca, 0xd4, 0x17, 0xe1, 0x09, 0xca, 0xec, 0x8b, 0x9e, 0x5e, 0xb9, 0xff, 0x58, 0x00, 0x53,
	0x82, 0xf7, 0xa6, 0xf8, 0x08, 0xa0, 0x8f, 0x5d, 0xda, 0xef, 0x44, 0x74, 0x80, 0x7a, 0x02, 0x97,
	0xa4, 0xe5, 0x05, 0x1d, 0x20, 0xd9, 0x86, 0x25, 0x1a, 0x04, 0x0c, 0x39, 0xc7, 0xb4, 0xc0, 0xa5,
	0xcd, 0xca, 0xce, 0x4a, 0xb6, 0x85, 0x96, 0x72, 0x78, 0x53, 0x04, 0xf9, 0x1c, 0x56, 0x12, 0x3a,
	0x4a, 0xcb, 0xd1, 0x09, 0x50, 0xd0, 0xb0, 0xcf, 0xeb, 0xf3, 0x32, 0xe8, 0x7a, 0x16, 0x74, 0xa0,
	0xdc, 0xfb, 0xd2, 0xeb, 0x2d, 0x27, 0xe6, 0x92, 0x93, 0x47, 0xb0, 0xe8, 0xc7, 0x91, 0xa0, 0xbe,
	0xe0, 0xf5, 0x2b, 0x79, 0xb6, 0x67, 0xca, 0xee, 0x4d, 0x00, 0xee, 0x2f, 0x40, 0x4e, 0x57, 0x91,
	0x6c, 0x41, 0xd9, 0xd8, 0x6a, 0x71, 0x77, 0x35, 0x82, 0xec, 0xc2, 0x55, 0xf3, 0x28, 0xa7, 0xa3,
	0x57, 0x92, 0xa3, 0xa7, 0x43, 0xcc, 0x46, 0x78, 0x79, 0xa8, 0xfb, 0x23, 0xac, 0x99, 0xee, 0x8b,
	0xce, 0xfa, 0x43, 0x58, 0x11, 0x94, 0x75, 0x51, 0x74, 0x26, 0x2d, 0x51, 0x45, 0x5f, 0x56, 0xe6,
	0xe7, 0xda, 0xea, 0xbe, 0xb5, 0xe0, 0x9a, 0x49, 0x20, 0x3b, 0xb9, 0x01, 0xcb, 0xf1, 0x30, 0x42,
	0xd6, 0x99, 0xe9, 0xe7, 0x55, 0x69, 0xcd, 0x62, 0x2f, 0x4c, 0x42, 0x1e, 0xc1, 0xaa, 0xfa, 0x67,
	0x5e, 0x72, 0x25, 0x09, 0xbd, 0xa6, 0x1c, 0xc6, 0xf5, 0xf6, 0x9b, 0x05, 0xeb, 0xb9, 0x8a, 0x5c,
	0xfe, 0x50, 0x7c, 0x06, 0x55, 0xb3, 0x9a, 0xfa, 0x70, 0xd4, 0x8b, 0xea, 0x2e, 0x1b, 0x96, 0x43,
	0xef, 0xfc, 0xb9, 0x00, 0x2b, 0xfb, 0x9a, 0xe1, 0xa5, 0x7a, 0x83, 0x48, 0x17, 0xca, 0xea, 0x96,
	0x24, 0xb7, 0xf2, 0x0d, 0xcf, 0xdd, 0x9d, 0x76, 0x2d, 0xef, 0xcc, 0xe4, 0xb9, 0x9b, 0xe3, 0xd6,
	0x9a, 0xbd, 0xaa, 0xb0, 0xdc, 0xa1, 0x91, 0xa3, 0x60, 0xbf, 0xfe, 0xfd, 0xef, 0xbb, 0xb9, 0xaa,
	0xbb, 0xa0, 0x1f, 0xb5, 0x5d, 0x6b, 0x8b, 0x08, 0x28, 0xab, 0xdb, 0x6d, 0x96, 0x28, 0x77, 0xe7,
	0x9d, 0x49, 0xf4, 0x44, 0x12, 0x29, 0xec, 0x2c, 0xd1, 0x4d, 0x7b, 0x5d, 0x13, 0x35, 0x5f, 0x4f,
	0xfb, 0xf1, 0x26, 0x65, 0x7d, 0x6b, 0xc9, 0xc7, 0x4f, 0xbf, 0x2d, 0xe4, 0x66, 0x96, 0xfd, 0xd4,
	0x7b, 0x73, 0x26, 0xf1, 0x8b, 0x71, 0x6b, 0xc3, 0xbe, 0xdf, 0x46, 0xe1, 0x50, 0x87, 0x27, 0xe8,
	0x87, 0x47, 0xa1, 0xef, 0xe8, 0xa7, 0xc6, 0x89, 0x8f, 0x66, 0xa4, 0x38, 0xe4, 0x4e, 0x91, 0x94,
	0xe6, 0x6b, 0x1d, 0xf1, 0x86, 0xfc, 0x04, 0xa5, 0x36, 0x0a, 0x42, 0x0c, 0x25, 0xe7, 0x49, 0x78,
	0x3a, 0x6e, 0xd5, 0xed, 0xf4, 0x55, 0x76, 0x44, 0x0f, 0x1d, 0xff, 0x98, 0x31, 0x8c, 0x84, 0xc9,
	0x5a, 0x23, 0x85, 0x05, 0x20, 0xbf, 0x5b, 0x70, 0xa3, 0x8d, 0x42, 0xa5, 0xdb, 0x1b, 0xe5, 0x5f,
	0xf3, 0xd3, 0x02, 0x4c, 0xff, 0xfb, 0xea, 0xe1, 0xda, 0x4e, 0x2a, 0x46, 0xf9, 0x9d, 0x23, 0x16,
	0x0f, 0x9c, 0xc3, 0x63, 0x1e, 0x46, 0xc8, 0xb9, 0x93, 0x50, 0x26, 0x22, 0x64, 0x52, 0xd6, 0x7d,
	0x72, 0xaf, 0xb0, 0x18, 0xe6, 0x50, 0x92, 0x77, 0x16, 0x5c, 0x79, 0xd9, 0xa3, 0xcc, 0x98, 0x8b,
	0x82, 0xfb, 0xc1, 0xbe, 0x5d, 0x78, 0xb7, 0x64, 0xa2, 0xbe, 0x19, 0xb7, 0x1e, 0xda, 0x1b, 0x32,
	0x8d, 0xac, 0x91, 0x96, 0x96, 0x1d, 0x27, 0x67, 0x18, 0x8a, 0x9e, 0x13, 0x8b, 0x1e, 0x32, 0x2e,
	0x95, 0xdd, 0x75, 0xed, 0x42, 0x65, 0x3c, 0xcd, 0xa0, 0xe7, 0xa6, 0xec, 0xe1, 0x49, 0xfc, 0xf3,
	0x07, 0xc9, 0xfa, 0x6a, 0xdc, 0xba, 0x67, 0xdf, 0x65, 0x32, 0xcf, 0x74, 0x50, 0xa6, 0xb2, 0x24,
	0x9d, 0x9a, 0x1b, 0xf7, 0xd6, 0x19, 0xa5, 0x4a, 0x63, 0x77, 0xad, 0xad, 0xbd, 0x07, 0x00, 0x7e,
	0x3c, 0xd0, 0x74, 0x7b, 0x55, 0x7d, 0x7e, 0x0f, 0xd2, 0x2f, 0xbc, 0x03, 0xeb, 0x87, 0x45, 0x65,
	0x4f, 0x0e, 0x0f, 0xcb, 0xf2, 0xa3, 0xef, 0x93, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xd2,
	0xee, 0x3a, 0x8d, 0x0a, 0x00, 0x00,
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
	// Entity Relation Get
	GetEntityByRelationship(ctx context.Context, in *GetRequestRelationship, opts ...grpc.CallOption) (*EntityResponse, error)
	// Entity Relation Share
	Share(ctx context.Context, in *RelationshipPayload, opts ...grpc.CallOption) (*RelationshipResponse, error)
	// Entity Relation Revoke
	Revoke(ctx context.Context, in *RelationshipPayload, opts ...grpc.CallOption) (*RelationshipResponse, error)
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

func (c *documentServiceClient) GetEntityByRelationship(ctx context.Context, in *GetRequestRelationship, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.DocumentService/GetEntityByRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Share(ctx context.Context, in *RelationshipPayload, opts ...grpc.CallOption) (*RelationshipResponse, error) {
	out := new(RelationshipResponse)
	err := c.cc.Invoke(ctx, "/entity.DocumentService/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Revoke(ctx context.Context, in *RelationshipPayload, opts ...grpc.CallOption) (*RelationshipResponse, error) {
	out := new(RelationshipResponse)
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
	// Entity Relation Get
	GetEntityByRelationship(context.Context, *GetRequestRelationship) (*EntityResponse, error)
	// Entity Relation Share
	Share(context.Context, *RelationshipPayload) (*RelationshipResponse, error)
	// Entity Relation Revoke
	Revoke(context.Context, *RelationshipPayload) (*RelationshipResponse, error)
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

func _DocumentService_GetEntityByRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestRelationship)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).GetEntityByRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.DocumentService/GetEntityByRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).GetEntityByRelationship(ctx, req.(*GetRequestRelationship))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationshipPayload)
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
		return srv.(DocumentServiceServer).Share(ctx, req.(*RelationshipPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelationshipPayload)
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
		return srv.(DocumentServiceServer).Revoke(ctx, req.(*RelationshipPayload))
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
			MethodName: "GetEntityByRelationship",
			Handler:    _DocumentService_GetEntityByRelationship_Handler,
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
