// Code generated by protoc-gen-go. DO NOT EDIT.
// source: document/service.proto

package documentpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/centrifuge/precise-proofs/proofs/proto"
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

type UpdateAccessTokenPayload struct {
	// The document which should contain the access token referenced below
	DelegatingDocumentIdentifier string `protobuf:"bytes,1,opt,name=delegating_document_identifier,json=delegatingDocumentIdentifier,proto3" json:"delegating_document_identifier,omitempty"`
	// The access token to be appended to the indicated document above
	AccessTokenParams    *AccessTokenParams `protobuf:"bytes,2,opt,name=access_token_params,json=accessTokenParams,proto3" json:"access_token_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateAccessTokenPayload) Reset()         { *m = UpdateAccessTokenPayload{} }
func (m *UpdateAccessTokenPayload) String() string { return proto.CompactTextString(m) }
func (*UpdateAccessTokenPayload) ProtoMessage()    {}
func (*UpdateAccessTokenPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_a33bfc9dfc97cd3f, []int{0}
}

func (m *UpdateAccessTokenPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccessTokenPayload.Unmarshal(m, b)
}
func (m *UpdateAccessTokenPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccessTokenPayload.Marshal(b, m, deterministic)
}
func (m *UpdateAccessTokenPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccessTokenPayload.Merge(m, src)
}
func (m *UpdateAccessTokenPayload) XXX_Size() int {
	return xxx_messageInfo_UpdateAccessTokenPayload.Size(m)
}
func (m *UpdateAccessTokenPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccessTokenPayload.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccessTokenPayload proto.InternalMessageInfo

func (m *UpdateAccessTokenPayload) GetDelegatingDocumentIdentifier() string {
	if m != nil {
		return m.DelegatingDocumentIdentifier
	}
	return ""
}

func (m *UpdateAccessTokenPayload) GetAccessTokenParams() *AccessTokenParams {
	if m != nil {
		return m.AccessTokenParams
	}
	return nil
}

type AccessTokenParams struct {
	// The identity being granted access to the document
	Grantee string `protobuf:"bytes,4,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Original identifier of the document
	DocumentIdentifier   string   `protobuf:"bytes,2,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessTokenParams) Reset()         { *m = AccessTokenParams{} }
func (m *AccessTokenParams) String() string { return proto.CompactTextString(m) }
func (*AccessTokenParams) ProtoMessage()    {}
func (*AccessTokenParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a33bfc9dfc97cd3f, []int{1}
}

func (m *AccessTokenParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessTokenParams.Unmarshal(m, b)
}
func (m *AccessTokenParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessTokenParams.Marshal(b, m, deterministic)
}
func (m *AccessTokenParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTokenParams.Merge(m, src)
}
func (m *AccessTokenParams) XXX_Size() int {
	return xxx_messageInfo_AccessTokenParams.Size(m)
}
func (m *AccessTokenParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTokenParams.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTokenParams proto.InternalMessageInfo

func (m *AccessTokenParams) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *AccessTokenParams) GetDocumentIdentifier() string {
	if m != nil {
		return m.DocumentIdentifier
	}
	return ""
}

type CreateDocumentProofRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Fields               []string `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDocumentProofRequest) Reset()         { *m = CreateDocumentProofRequest{} }
func (m *CreateDocumentProofRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDocumentProofRequest) ProtoMessage()    {}
func (*CreateDocumentProofRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a33bfc9dfc97cd3f, []int{2}
}

func (m *CreateDocumentProofRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDocumentProofRequest.Unmarshal(m, b)
}
func (m *CreateDocumentProofRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDocumentProofRequest.Marshal(b, m, deterministic)
}
func (m *CreateDocumentProofRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDocumentProofRequest.Merge(m, src)
}
func (m *CreateDocumentProofRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDocumentProofRequest.Size(m)
}
func (m *CreateDocumentProofRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDocumentProofRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDocumentProofRequest proto.InternalMessageInfo

func (m *CreateDocumentProofRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *CreateDocumentProofRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateDocumentProofRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

// ResponseHeader contains a set of common fields for most documents
type ResponseHeader struct {
	DocumentId           string   `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	VersionId            string   `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_a33bfc9dfc97cd3f, []int{3}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
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

type DocumentProof struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	FieldProofs          []*Proof        `protobuf:"bytes,2,rep,name=field_proofs,json=fieldProofs,proto3" json:"field_proofs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DocumentProof) Reset()         { *m = DocumentProof{} }
func (m *DocumentProof) String() string { return proto.CompactTextString(m) }
func (*DocumentProof) ProtoMessage()    {}
func (*DocumentProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_a33bfc9dfc97cd3f, []int{4}
}

func (m *DocumentProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentProof.Unmarshal(m, b)
}
func (m *DocumentProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentProof.Marshal(b, m, deterministic)
}
func (m *DocumentProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentProof.Merge(m, src)
}
func (m *DocumentProof) XXX_Size() int {
	return xxx_messageInfo_DocumentProof.Size(m)
}
func (m *DocumentProof) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentProof.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentProof proto.InternalMessageInfo

func (m *DocumentProof) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DocumentProof) GetFieldProofs() []*Proof {
	if m != nil {
		return m.FieldProofs
	}
	return nil
}

type Proof struct {
	Property string `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Salt     string `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty"`
	// hash is filled if value & salt are not available
	Hash                 string   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	SortedHashes         []string `protobuf:"bytes,5,rep,name=sorted_hashes,json=sortedHashes,proto3" json:"sorted_hashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_a33bfc9dfc97cd3f, []int{5}
}

func (m *Proof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proof.Unmarshal(m, b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return xxx_messageInfo_Proof.Size(m)
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *Proof) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Proof) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *Proof) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Proof) GetSortedHashes() []string {
	if m != nil {
		return m.SortedHashes
	}
	return nil
}

type CreateDocumentProofForVersionRequest struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Fields               []string `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDocumentProofForVersionRequest) Reset()         { *m = CreateDocumentProofForVersionRequest{} }
func (m *CreateDocumentProofForVersionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDocumentProofForVersionRequest) ProtoMessage()    {}
func (*CreateDocumentProofForVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a33bfc9dfc97cd3f, []int{6}
}

func (m *CreateDocumentProofForVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDocumentProofForVersionRequest.Unmarshal(m, b)
}
func (m *CreateDocumentProofForVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDocumentProofForVersionRequest.Marshal(b, m, deterministic)
}
func (m *CreateDocumentProofForVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDocumentProofForVersionRequest.Merge(m, src)
}
func (m *CreateDocumentProofForVersionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDocumentProofForVersionRequest.Size(m)
}
func (m *CreateDocumentProofForVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDocumentProofForVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDocumentProofForVersionRequest proto.InternalMessageInfo

func (m *CreateDocumentProofForVersionRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *CreateDocumentProofForVersionRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateDocumentProofForVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CreateDocumentProofForVersionRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateAccessTokenPayload)(nil), "document.UpdateAccessTokenPayload")
	proto.RegisterType((*AccessTokenParams)(nil), "document.AccessTokenParams")
	proto.RegisterType((*CreateDocumentProofRequest)(nil), "document.CreateDocumentProofRequest")
	proto.RegisterType((*ResponseHeader)(nil), "document.ResponseHeader")
	proto.RegisterType((*DocumentProof)(nil), "document.DocumentProof")
	proto.RegisterType((*Proof)(nil), "document.Proof")
	proto.RegisterType((*CreateDocumentProofForVersionRequest)(nil), "document.CreateDocumentProofForVersionRequest")
}

func init() { proto.RegisterFile("document/service.proto", fileDescriptor_a33bfc9dfc97cd3f) }

var fileDescriptor_a33bfc9dfc97cd3f = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0x13, 0x49,
	0x10, 0xd6, 0xd8, 0xf9, 0x2d, 0x3b, 0x1b, 0xa5, 0xb3, 0xca, 0x8e, 0xbc, 0x49, 0xb6, 0x77, 0x36,
	0xda, 0xb5, 0xb2, 0xc4, 0x46, 0xe6, 0xc6, 0x2d, 0x21, 0x42, 0x89, 0xb8, 0x58, 0x86, 0x80, 0xc4,
	0x01, 0xab, 0x33, 0x53, 0x1e, 0x0f, 0x4c, 0xa6, 0x87, 0xee, 0xb6, 0x91, 0x15, 0x71, 0x80, 0x03,
	0x17, 0x38, 0x85, 0x17, 0xe0, 0x05, 0x78, 0x1a, 0x2e, 0x3c, 0x00, 0x4f, 0xc1, 0x09, 0x4d, 0x77,
	0x8f, 0xc7, 0xc6, 0x49, 0x38, 0xe4, 0x34, 0x5d, 0xd5, 0x5f, 0x57, 0x7f, 0xdf, 0x57, 0x3d, 0x05,
	0x1b, 0x01, 0xf7, 0x07, 0x67, 0x98, 0xa8, 0xa6, 0x44, 0x31, 0x8c, 0x7c, 0x6c, 0xa4, 0x82, 0x2b,
	0x4e, 0x96, 0xf2, 0x7c, 0x6d, 0x33, 0xe4, 0x3c, 0x8c, 0xb1, 0xc9, 0xd2, 0xa8, 0xc9, 0x92, 0x84,
	0x2b, 0xa6, 0x22, 0x9e, 0x48, 0x83, 0xab, 0xfd, 0x97, 0x0a, 0xf4, 0x23, 0x89, 0x7b, 0xa9, 0xe0,
	0xbc, 0x27, 0x9b, 0xc5, 0x47, 0x71, 0x13, 0x58, 0xe0, 0x2d, 0xfd, 0xf1, 0xf7, 0x42, 0x4c, 0xf6,
	0xe4, 0x2b, 0x16, 0x86, 0x28, 0x9a, 0x3c, 0xd5, 0xa5, 0x66, 0xcb, 0x7a, 0x9f, 0x1d, 0x70, 0x4f,
	0xd2, 0x80, 0x29, 0xdc, 0xf7, 0x7d, 0x94, 0xf2, 0x11, 0x7f, 0x81, 0x49, 0x9b, 0x8d, 0x62, 0xce,
	0x02, 0x72, 0x08, 0xdb, 0x01, 0xc6, 0x18, 0x32, 0x15, 0x25, 0x61, 0x37, 0x27, 0xda, 0x8d, 0x02,
	0x4c, 0x54, 0xd4, 0x8b, 0x50, 0xb8, 0x0e, 0x75, 0xea, 0xcb, 0x9d, 0xcd, 0x02, 0x75, 0x68, 0x41,
	0xc7, 0x63, 0x0c, 0x79, 0x00, 0xeb, 0x4c, 0xd7, 0xee, 0xaa, 0xac, 0x78, 0x37, 0x65, 0x82, 0x9d,
	0x49, 0xb7, 0x44, 0x9d, 0x7a, 0xa5, 0xf5, 0x67, 0x23, 0x2f, 0xdb, 0x98, 0x22, 0x90, 0x41, 0x3a,
	0x6b, 0xec, 0xe7, 0x94, 0xf7, 0x0c, 0xd6, 0x66, 0x70, 0xc4, 0x85, 0xc5, 0x50, 0xb0, 0x44, 0x21,
	0xba, 0x73, 0x9a, 0x50, 0x1e, 0x92, 0x26, 0xac, 0x5f, 0x46, 0xbb, 0xa4, 0x51, 0x24, 0x98, 0x21,
	0xeb, 0xf5, 0xa1, 0x76, 0x4f, 0x20, 0x53, 0x98, 0x0b, 0x69, 0x67, 0xd6, 0x76, 0xf0, 0xe5, 0x00,
	0xa5, 0x22, 0xdb, 0x00, 0x33, 0xe2, 0x27, 0x32, 0x84, 0xc0, 0x9c, 0x1a, 0xa5, 0x68, 0xeb, 0xeb,
	0x35, 0xd9, 0x80, 0x85, 0x5e, 0x84, 0x71, 0x20, 0xdd, 0x32, 0x2d, 0xd7, 0x97, 0x3b, 0x36, 0xf2,
	0x7a, 0xf0, 0x5b, 0x07, 0x65, 0xca, 0x13, 0x89, 0x47, 0xc8, 0x02, 0x14, 0xe4, 0x2f, 0xa8, 0x4c,
	0x90, 0xcd, 0xcb, 0x17, 0x24, 0xc9, 0x16, 0xc0, 0x10, 0x85, 0x8c, 0x78, 0x92, 0xed, 0x9b, 0x4b,
	0x96, 0x6d, 0xe6, 0x38, 0x20, 0xbf, 0xc3, 0xbc, 0x54, 0x4c, 0xa1, 0x5b, 0xd6, 0x3b, 0x26, 0xf0,
	0x06, 0xb0, 0x32, 0xa5, 0x85, 0xdc, 0x86, 0x85, 0xbe, 0xbe, 0x50, 0xdf, 0x50, 0x69, 0xb9, 0x45,
	0x0b, 0xa6, 0x09, 0x75, 0x2c, 0x8e, 0xb4, 0xa0, 0xaa, 0x49, 0x77, 0xcd, 0xa3, 0x73, 0x4b, 0xb4,
	0x5c, 0xaf, 0xb4, 0x56, 0x8b, 0x73, 0xc6, 0xa4, 0x8a, 0x06, 0xe9, 0xb5, 0xf4, 0xde, 0x39, 0x30,
	0x6f, 0xee, 0xab, 0xc1, 0x52, 0x2a, 0x78, 0x8a, 0x42, 0x8d, 0xac, 0xa6, 0x71, 0x9c, 0x51, 0x1e,
	0xb2, 0x78, 0x90, 0x3b, 0x66, 0x82, 0xcc, 0x46, 0xc9, 0x62, 0x65, 0x75, 0xe8, 0x75, 0x96, 0xeb,
	0x33, 0xd9, 0xb7, 0x0d, 0xd6, 0x6b, 0xf2, 0x0f, 0xac, 0x48, 0x2e, 0x14, 0x06, 0xdd, 0x2c, 0x44,
	0xe9, 0xce, 0x6b, 0x87, 0xab, 0x26, 0x79, 0xa4, 0x73, 0xde, 0x07, 0x07, 0x76, 0x2e, 0x69, 0xe9,
	0x7d, 0x2e, 0x1e, 0x1b, 0xe7, 0x6e, 0xd2, 0x5c, 0x17, 0x16, 0xad, 0xff, 0x96, 0x6c, 0x1e, 0x4e,
	0xb4, 0x7d, 0x6e, 0xb2, 0xed, 0xad, 0xef, 0x65, 0x58, 0xcd, 0x89, 0x3c, 0x34, 0x93, 0x80, 0x7c,
	0x75, 0x60, 0xfd, 0x12, 0x8a, 0x64, 0xa7, 0x70, 0xf8, 0xea, 0x47, 0x59, 0xfb, 0xa3, 0x40, 0x4d,
	0xed, 0x7b, 0x6f, 0x9c, 0x8b, 0xfd, 0x27, 0xb5, 0x13, 0x73, 0x54, 0x52, 0x46, 0xe3, 0x48, 0x2a,
	0xca, 0x7b, 0xd4, 0x8e, 0x12, 0x6a, 0xda, 0x49, 0x7b, 0x5c, 0x50, 0xd5, 0x47, 0x2a, 0x53, 0xf4,
	0x33, 0xa9, 0x01, 0x35, 0x5c, 0x33, 0x68, 0x96, 0xcf, 0xcb, 0xd3, 0x30, 0x1a, 0x62, 0x42, 0x4f,
	0x47, 0xf4, 0xf8, 0xf0, 0xed, 0x97, 0x6f, 0x1f, 0x4b, 0x7f, 0x7b, 0x9b, 0xcd, 0xf1, 0x58, 0x3b,
	0x2f, 0xac, 0x7a, 0x6d, 0x06, 0xd2, 0x5d, 0x67, 0x97, 0xbc, 0x2f, 0xc1, 0xd6, 0xb5, 0xee, 0x93,
	0xc6, 0xb5, 0x22, 0x67, 0xda, 0x74, 0xb5, 0xdc, 0x4f, 0xce, 0xc5, 0x7e, 0x5c, 0x7b, 0x7e, 0x63,
	0xb9, 0x46, 0xa5, 0xed, 0xe3, 0x2f, 0x3d, 0xf8, 0xdf, 0xfb, 0xf7, 0x0a, 0x0f, 0xce, 0x6d, 0x89,
	0xc2, 0x8d, 0x83, 0x5d, 0xa8, 0xfa, 0xfc, 0x6c, 0x2c, 0xe0, 0xa0, 0x6a, 0x5f, 0x40, 0x3b, 0x9b,
	0xc5, 0x6d, 0xe7, 0xe9, 0xf8, 0x67, 0x4f, 0x4f, 0x4f, 0x17, 0xf4, 0x80, 0xbe, 0xf3, 0x23, 0x00,
	0x00, 0xff, 0xff, 0x17, 0xaa, 0x39, 0x7f, 0x39, 0x06, 0x00, 0x00,
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
	CreateDocumentProof(ctx context.Context, in *CreateDocumentProofRequest, opts ...grpc.CallOption) (*DocumentProof, error)
	CreateDocumentProofForVersion(ctx context.Context, in *CreateDocumentProofForVersionRequest, opts ...grpc.CallOption) (*DocumentProof, error)
}

type documentServiceClient struct {
	cc *grpc.ClientConn
}

func NewDocumentServiceClient(cc *grpc.ClientConn) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) CreateDocumentProof(ctx context.Context, in *CreateDocumentProofRequest, opts ...grpc.CallOption) (*DocumentProof, error) {
	out := new(DocumentProof)
	err := c.cc.Invoke(ctx, "/document.DocumentService/CreateDocumentProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) CreateDocumentProofForVersion(ctx context.Context, in *CreateDocumentProofForVersionRequest, opts ...grpc.CallOption) (*DocumentProof, error) {
	out := new(DocumentProof)
	err := c.cc.Invoke(ctx, "/document.DocumentService/CreateDocumentProofForVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServiceServer is the server API for DocumentService service.
type DocumentServiceServer interface {
	CreateDocumentProof(context.Context, *CreateDocumentProofRequest) (*DocumentProof, error)
	CreateDocumentProofForVersion(context.Context, *CreateDocumentProofForVersionRequest) (*DocumentProof, error)
}

func RegisterDocumentServiceServer(s *grpc.Server, srv DocumentServiceServer) {
	s.RegisterService(&_DocumentService_serviceDesc, srv)
}

func _DocumentService_CreateDocumentProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).CreateDocumentProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document.DocumentService/CreateDocumentProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).CreateDocumentProof(ctx, req.(*CreateDocumentProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_CreateDocumentProofForVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentProofForVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).CreateDocumentProofForVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document.DocumentService/CreateDocumentProofForVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).CreateDocumentProofForVersion(ctx, req.(*CreateDocumentProofForVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DocumentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "document.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDocumentProof",
			Handler:    _DocumentService_CreateDocumentProof_Handler,
		},
		{
			MethodName: "CreateDocumentProofForVersion",
			Handler:    _DocumentService_CreateDocumentProofForVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "document/service.proto",
}
