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

type GetRequestRelationship struct {
	RelationshipId       string   `protobuf:"bytes,1,opt,name=relationship_id,json=relationshipId,proto3" json:"relationship_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequestRelationship) Reset()         { *m = GetRequestRelationship{} }
func (m *GetRequestRelationship) String() string { return proto.CompactTextString(m) }
func (*GetRequestRelationship) ProtoMessage()    {}
func (*GetRequestRelationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{0}
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

func (m *GetRequestRelationship) GetRelationshipId() string {
	if m != nil {
		return m.RelationshipId
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
	ReadAccess  []string    `protobuf:"bytes,1,rep,name=read_access,json=readAccess,proto3" json:"read_access,omitempty"`
	WriteAccess []string    `protobuf:"bytes,2,rep,name=write_access,json=writeAccess,proto3" json:"write_access,omitempty"`
	Data        *EntityData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// custom attributes
	Attributes           map[string]*document.Attribute `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
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

func (m *EntityCreatePayload) GetReadAccess() []string {
	if m != nil {
		return m.ReadAccess
	}
	return nil
}

func (m *EntityCreatePayload) GetWriteAccess() []string {
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

func (m *EntityCreatePayload) GetAttributes() map[string]*document.Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type EntityUpdatePayload struct {
	DocumentId  string      `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	ReadAccess  []string    `protobuf:"bytes,2,rep,name=read_access,json=readAccess,proto3" json:"read_access,omitempty"`
	WriteAccess []string    `protobuf:"bytes,3,rep,name=write_access,json=writeAccess,proto3" json:"write_access,omitempty"`
	Data        *EntityData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// custom attributes
	Attributes           map[string]*document.Attribute `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
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

func (m *EntityUpdatePayload) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func (m *EntityUpdatePayload) GetReadAccess() []string {
	if m != nil {
		return m.ReadAccess
	}
	return nil
}

func (m *EntityUpdatePayload) GetWriteAccess() []string {
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

func (m *EntityUpdatePayload) GetAttributes() map[string]*document.Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type EntityResponse struct {
	Header *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   *EntityDataResponse      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// custom attributes
	Attributes           map[string]*document.Attribute `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
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

func (m *EntityResponse) GetData() *EntityDataResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *EntityResponse) GetAttributes() map[string]*document.Attribute {
	if m != nil {
		return m.Attributes
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
	return fileDescriptor_c1b437217b9e14a2, []int{7}
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
	DocumentId           string   `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	TargetIdentity       string   `protobuf:"bytes,2,opt,name=target_identity,json=targetIdentity,proto3" json:"target_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationshipPayload) Reset()         { *m = RelationshipPayload{} }
func (m *RelationshipPayload) String() string { return proto.CompactTextString(m) }
func (*RelationshipPayload) ProtoMessage()    {}
func (*RelationshipPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{8}
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

func (m *RelationshipPayload) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
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
	return fileDescriptor_c1b437217b9e14a2, []int{9}
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
	Relationship         []*RelationshipData      `protobuf:"bytes,2,rep,name=relationship,proto3" json:"relationship,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RelationshipResponse) Reset()         { *m = RelationshipResponse{} }
func (m *RelationshipResponse) String() string { return proto.CompactTextString(m) }
func (*RelationshipResponse) ProtoMessage()    {}
func (*RelationshipResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b437217b9e14a2, []int{10}
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

func (m *RelationshipResponse) GetRelationship() []*RelationshipData {
	if m != nil {
		return m.Relationship
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequestRelationship)(nil), "entity.GetRequestRelationship")
	proto.RegisterType((*GetVersionRequest)(nil), "entity.GetVersionRequest")
	proto.RegisterType((*EntityCreatePayload)(nil), "entity.EntityCreatePayload")
	proto.RegisterMapType((map[string]*document.Attribute)(nil), "entity.EntityCreatePayload.AttributesEntry")
	proto.RegisterType((*EntityUpdatePayload)(nil), "entity.EntityUpdatePayload")
	proto.RegisterMapType((map[string]*document.Attribute)(nil), "entity.EntityUpdatePayload.AttributesEntry")
	proto.RegisterType((*EntityResponse)(nil), "entity.EntityResponse")
	proto.RegisterMapType((map[string]*document.Attribute)(nil), "entity.EntityResponse.AttributesEntry")
	proto.RegisterType((*Relationship)(nil), "entity.Relationship")
	proto.RegisterType((*EntityData)(nil), "entity.EntityData")
	proto.RegisterType((*EntityDataResponse)(nil), "entity.EntityDataResponse")
	proto.RegisterType((*RelationshipPayload)(nil), "entity.RelationshipPayload")
	proto.RegisterType((*RelationshipData)(nil), "entity.RelationshipData")
	proto.RegisterType((*RelationshipResponse)(nil), "entity.RelationshipResponse")
}

func init() { proto.RegisterFile("entity/service.proto", fileDescriptor_c1b437217b9e14a2) }

var fileDescriptor_c1b437217b9e14a2 = []byte{
	// 826 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0x93, 0x36, 0xa4, 0x27, 0x69, 0xd2, 0x4e, 0x4a, 0xb0, 0x22, 0x28, 0xc6, 0x12, 0x6d,
	0xe9, 0x4f, 0x02, 0xe1, 0x06, 0x55, 0x08, 0x29, 0x69, 0x4b, 0x88, 0x10, 0x28, 0x32, 0x02, 0x24,
	0x6e, 0xa2, 0x89, 0x7d, 0x9a, 0x5a, 0x24, 0xb6, 0x99, 0x99, 0xa4, 0x8a, 0x80, 0x1b, 0x2e, 0xb8,
	0x04, 0x69, 0xf7, 0x0d, 0xf6, 0x72, 0x9f, 0x61, 0xef, 0xf6, 0x11, 0xf6, 0x11, 0x76, 0x1f, 0x64,
	0x95, 0x99, 0x71, 0x62, 0xb7, 0xd9, 0xdd, 0xee, 0x5e, 0xf4, 0xca, 0x9e, 0xef, 0xfb, 0xce, 0x99,
	0x99, 0xef, 0xcc, 0xf1, 0x18, 0x76, 0x30, 0x10, 0xbe, 0x98, 0x35, 0x38, 0xb2, 0xa9, 0xef, 0x62,
	0x3d, 0x62, 0xa1, 0x08, 0x49, 0x4e, 0xa1, 0xb5, 0xaa, 0x17, 0xba, 0x93, 0x31, 0x06, 0x22, 0xcd,
	0xd7, 0x2a, 0x3a, 0x4a, 0x3d, 0x34, 0xf8, 0xe1, 0x30, 0x0c, 0x87, 0x23, 0x6c, 0xd0, 0xc8, 0x6f,
	0xd0, 0x20, 0x08, 0x05, 0x15, 0x7e, 0x18, 0x70, 0xcd, 0x1e, 0xcb, 0x87, 0x7b, 0x32, 0xc4, 0xe0,
	0x84, 0x5f, 0xd3, 0xe1, 0x10, 0x59, 0x23, 0x8c, 0xa4, 0xe2, 0xb6, 0xda, 0x6e, 0x41, 0xb5, 0x83,
	0xc2, 0xc1, 0x3f, 0x26, 0xc8, 0x85, 0x83, 0x23, 0x45, 0x5e, 0xf9, 0x11, 0xd9, 0x87, 0x32, 0x4b,
	0x8c, 0xfb, 0xbe, 0x67, 0x1a, 0x96, 0x71, 0xb0, 0xe1, 0x94, 0x92, 0x70, 0xd7, 0xb3, 0x7f, 0x80,
	0xed, 0x0e, 0x8a, 0x5f, 0x90, 0x71, 0x3f, 0x0c, 0x74, 0x26, 0xb2, 0x0b, 0xe0, 0x7b, 0xf3, 0x55,
	0x5f, 0xfa, 0xc8, 0x74, 0x60, 0x02, 0x21, 0x26, 0xbc, 0x37, 0x55, 0x11, 0x66, 0x46, 0x92, 0xf1,
	0xd0, 0x7e, 0x94, 0x81, 0xca, 0x85, 0xdc, 0xee, 0x19, 0x43, 0x2a, 0xb0, 0x47, 0x67, 0xa3, 0x90,
	0x7a, 0xe4, 0x63, 0x28, 0x30, 0xa4, 0x5e, 0x9f, 0xba, 0x2e, 0x72, 0x6e, 0x1a, 0x56, 0x76, 0x9e,
	0x72, 0x0e, 0xb5, 0x24, 0x42, 0x3e, 0x81, 0xe2, 0x35, 0xf3, 0x05, 0xc6, 0x8a, 0x8c, 0x54, 0x14,
	0x24, 0xa6, 0x25, 0x7b, 0xb0, 0xe6, 0x51, 0x41, 0xcd, 0xac, 0x65, 0x1c, 0x14, 0x9a, 0xa4, 0xae,
	0x6d, 0x55, 0xd3, 0x9d, 0x53, 0x41, 0x1d, 0xc9, 0x93, 0xef, 0x01, 0xa8, 0x10, 0xcc, 0x1f, 0x4c,
	0x04, 0x72, 0x73, 0xcd, 0xca, 0x1e, 0x14, 0x9a, 0x47, 0x69, 0x75, 0x6a, 0x71, 0xf5, 0xd6, 0x42,
	0x7d, 0x11, 0x08, 0x36, 0x73, 0x12, 0xe1, 0x35, 0x07, 0xca, 0x37, 0x68, 0xb2, 0x05, 0xd9, 0xdf,
	0x71, 0xa6, 0x6d, 0x99, 0xbf, 0x92, 0xcf, 0x60, 0x7d, 0x4a, 0x47, 0x13, 0x94, 0x6e, 0x14, 0x9a,
	0x95, 0x7a, 0x7c, 0x20, 0x96, 0xa9, 0x1d, 0xa5, 0x38, 0xcd, 0x7c, 0x65, 0xd8, 0x4f, 0x17, 0x26,
	0xfd, 0x1c, 0x79, 0x69, 0x93, 0xe2, 0xc0, 0x65, 0xc1, 0x20, 0x86, 0xba, 0xb7, 0x5c, 0xcc, 0xbc,
	0xd1, 0xc5, 0xec, 0xab, 0x5d, 0x5c, 0x7b, 0x2b, 0x17, 0xd7, 0x57, 0xb9, 0x98, 0x5a, 0xfd, 0xbd,
	0xbb, 0xf8, 0x5f, 0x06, 0x4a, 0x6a, 0x1d, 0x0e, 0xf2, 0x28, 0x0c, 0x38, 0x92, 0xcf, 0x21, 0x77,
	0x85, 0xd4, 0xd3, 0x67, 0xb6, 0xd0, 0x34, 0x97, 0x29, 0x62, 0xcd, 0x77, 0x92, 0x77, 0xb4, 0x8e,
	0xd4, 0xb5, 0x1b, 0x6a, 0xca, 0xda, 0x0a, 0x37, 0x74, 0x9c, 0x76, 0xe5, 0xdb, 0x94, 0x2b, 0x59,
	0xe9, 0xca, 0x5e, 0x3a, 0x2a, 0x8e, 0xb8, 0x77, 0x43, 0xda, 0x50, 0x4c, 0x7d, 0x03, 0x6a, 0x90,
	0x57, 0x3d, 0x2b, 0xe2, 0xac, 0x8b, 0x31, 0xa9, 0x42, 0x8e, 0xba, 0xc2, 0x9f, 0xaa, 0xdc, 0x79,
	0x47, 0x8f, 0xec, 0xe7, 0x06, 0xc0, 0x72, 0xf3, 0xaf, 0x4d, 0xf1, 0x11, 0xc0, 0x08, 0x87, 0x74,
	0xd4, 0x0f, 0xe8, 0x18, 0xf5, 0x77, 0x60, 0x43, 0x22, 0x3f, 0xd2, 0x31, 0x92, 0x13, 0xd8, 0xa0,
	0x9e, 0xc7, 0x90, 0xf3, 0x85, 0x51, 0xe5, 0xd8, 0xa8, 0x96, 0x22, 0x9c, 0xa5, 0x82, 0x7c, 0x03,
	0xe5, 0x88, 0xce, 0xe4, 0xd1, 0xf7, 0x50, 0x50, 0x7f, 0x14, 0x77, 0xee, 0xfb, 0x71, 0x50, 0x4f,
	0xd1, 0xe7, 0x92, 0x75, 0x4a, 0x51, 0x72, 0xc8, 0xc9, 0x11, 0xe4, 0xdd, 0x30, 0x10, 0xd4, 0x15,
	0xf1, 0x61, 0x5d, 0xcc, 0x76, 0xa6, 0x70, 0x67, 0x21, 0xb0, 0xff, 0x02, 0x72, 0xbb, 0xc2, 0xe4,
	0x10, 0x72, 0x89, 0xad, 0xae, 0xee, 0x0d, 0xad, 0x20, 0xa7, 0xb0, 0x99, 0xfc, 0x90, 0xaa, 0x5e,
	0x2c, 0x34, 0x77, 0xe2, 0x90, 0x64, 0x21, 0x9c, 0xb4, 0xd4, 0xee, 0x43, 0x25, 0x49, 0xdf, 0xb9,
	0xfb, 0xf7, 0xa1, 0x2c, 0x28, 0x1b, 0xe2, 0x9c, 0xd6, 0x0b, 0x55, 0xae, 0x97, 0x14, 0xdc, 0xd5,
	0xa8, 0xfd, 0xbf, 0x01, 0x5b, 0xc9, 0x19, 0x64, 0x29, 0x3f, 0x85, 0x52, 0x78, 0x1d, 0x20, 0xeb,
	0xdf, 0x28, 0xe8, 0xa6, 0x44, 0xe3, 0xd8, 0x3b, 0x4f, 0x42, 0x8e, 0x60, 0x5b, 0xbd, 0xf5, 0x13,
	0x57, 0x45, 0x56, 0x4a, 0xb7, 0x14, 0xd1, 0x5d, 0xe0, 0xf6, 0xbf, 0x06, 0xec, 0xa4, 0x2c, 0x79,
	0xf7, 0x8e, 0xfd, 0x1a, 0x8a, 0x49, 0x3b, 0xb5, 0xf1, 0xe6, 0x2a, 0xe3, 0x65, 0xc5, 0x52, 0xea,
	0xe6, 0x13, 0x03, 0x36, 0x55, 0x39, 0x7f, 0x52, 0x57, 0x35, 0x79, 0x6c, 0xc0, 0x07, 0x1d, 0x14,
	0x0a, 0x6c, 0xcf, 0x52, 0x1d, 0xb4, 0x1b, 0x67, 0x5d, 0x7d, 0xcb, 0xd6, 0xaa, 0xab, 0x3b, 0xdf,
	0xfe, 0xf5, 0x41, 0xcb, 0xae, 0x59, 0x1d, 0x14, 0x96, 0xe2, 0xad, 0x4b, 0x16, 0x8e, 0xad, 0xc1,
	0x84, 0xfb, 0x01, 0x72, 0x6e, 0x45, 0x94, 0x89, 0x00, 0xd9, 0x3f, 0xcf, 0x5e, 0x3c, 0xcc, 0x1c,
	0x93, 0xc3, 0xc6, 0xf4, 0x8b, 0x46, 0xea, 0x84, 0x34, 0xfe, 0xbc, 0x71, 0x77, 0xff, 0xad, 0x7f,
	0x21, 0xda, 0x7b, 0x00, 0x6e, 0x38, 0xd6, 0xb3, 0xb6, 0x8b, 0x7a, 0x0f, 0xbd, 0xf9, 0xcf, 0x40,
	0xcf, 0xf8, 0x2d, 0xaf, 0xf0, 0x68, 0x30, 0xc8, 0xc9, 0xff, 0x83, 0x2f, 0x5f, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x46, 0xab, 0x4a, 0x18, 0xb8, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EntityServiceClient is the client API for EntityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EntityServiceClient interface {
	// Entity Relation Get
	GetEntityByRelationship(ctx context.Context, in *GetRequestRelationship, opts ...grpc.CallOption) (*EntityResponse, error)
}

type entityServiceClient struct {
	cc *grpc.ClientConn
}

func NewEntityServiceClient(cc *grpc.ClientConn) EntityServiceClient {
	return &entityServiceClient{cc}
}

func (c *entityServiceClient) GetEntityByRelationship(ctx context.Context, in *GetRequestRelationship, opts ...grpc.CallOption) (*EntityResponse, error) {
	out := new(EntityResponse)
	err := c.cc.Invoke(ctx, "/entity.EntityService/GetEntityByRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityServiceServer is the server API for EntityService service.
type EntityServiceServer interface {
	// Entity Relation Get
	GetEntityByRelationship(context.Context, *GetRequestRelationship) (*EntityResponse, error)
}

func RegisterEntityServiceServer(s *grpc.Server, srv EntityServiceServer) {
	s.RegisterService(&_EntityService_serviceDesc, srv)
}

func _EntityService_GetEntityByRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestRelationship)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityServiceServer).GetEntityByRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.EntityService/GetEntityByRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityServiceServer).GetEntityByRelationship(ctx, req.(*GetRequestRelationship))
	}
	return interceptor(ctx, in, info, handler)
}

var _EntityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "entity.EntityService",
	HandlerType: (*EntityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntityByRelationship",
			Handler:    _EntityService_GetEntityByRelationship_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entity/service.proto",
}
