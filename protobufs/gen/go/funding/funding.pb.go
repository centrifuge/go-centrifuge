// Code generated by protoc-gen-go. DO NOT EDIT.
// source: funding/funding.proto

package funpb

import (
	context "context"
	fmt "fmt"
	math "math"

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

type FundingCreatePayload struct {
	DocumentId           string       `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	Data                 *FundingData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FundingCreatePayload) Reset()         { *m = FundingCreatePayload{} }
func (m *FundingCreatePayload) String() string { return proto.CompactTextString(m) }
func (*FundingCreatePayload) ProtoMessage()    {}
func (*FundingCreatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{0}
}

func (m *FundingCreatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingCreatePayload.Unmarshal(m, b)
}
func (m *FundingCreatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingCreatePayload.Marshal(b, m, deterministic)
}
func (m *FundingCreatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingCreatePayload.Merge(m, src)
}
func (m *FundingCreatePayload) XXX_Size() int {
	return xxx_messageInfo_FundingCreatePayload.Size(m)
}
func (m *FundingCreatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingCreatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_FundingCreatePayload proto.InternalMessageInfo

func (m *FundingCreatePayload) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func (m *FundingCreatePayload) GetData() *FundingData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FundingUpdatePayload struct {
	DocumentId           string       `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	AgreementId          string       `protobuf:"bytes,2,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	Data                 *FundingData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FundingUpdatePayload) Reset()         { *m = FundingUpdatePayload{} }
func (m *FundingUpdatePayload) String() string { return proto.CompactTextString(m) }
func (*FundingUpdatePayload) ProtoMessage()    {}
func (*FundingUpdatePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{1}
}

func (m *FundingUpdatePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingUpdatePayload.Unmarshal(m, b)
}
func (m *FundingUpdatePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingUpdatePayload.Marshal(b, m, deterministic)
}
func (m *FundingUpdatePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingUpdatePayload.Merge(m, src)
}
func (m *FundingUpdatePayload) XXX_Size() int {
	return xxx_messageInfo_FundingUpdatePayload.Size(m)
}
func (m *FundingUpdatePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingUpdatePayload.DiscardUnknown(m)
}

var xxx_messageInfo_FundingUpdatePayload proto.InternalMessageInfo

func (m *FundingUpdatePayload) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func (m *FundingUpdatePayload) GetAgreementId() string {
	if m != nil {
		return m.AgreementId
	}
	return ""
}

func (m *FundingUpdatePayload) GetData() *FundingData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FundingResponse struct {
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 *FundingResponseData     `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FundingResponse) Reset()         { *m = FundingResponse{} }
func (m *FundingResponse) String() string { return proto.CompactTextString(m) }
func (*FundingResponse) ProtoMessage()    {}
func (*FundingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{2}
}

func (m *FundingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingResponse.Unmarshal(m, b)
}
func (m *FundingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingResponse.Marshal(b, m, deterministic)
}
func (m *FundingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingResponse.Merge(m, src)
}
func (m *FundingResponse) XXX_Size() int {
	return xxx_messageInfo_FundingResponse.Size(m)
}
func (m *FundingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FundingResponse proto.InternalMessageInfo

func (m *FundingResponse) GetHeader() *document.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FundingResponse) GetData() *FundingResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FundingListResponse struct {
	Header               *document.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data                 []*FundingResponseData   `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FundingListResponse) Reset()         { *m = FundingListResponse{} }
func (m *FundingListResponse) String() string { return proto.CompactTextString(m) }
func (*FundingListResponse) ProtoMessage()    {}
func (*FundingListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{3}
}

func (m *FundingListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingListResponse.Unmarshal(m, b)
}
func (m *FundingListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingListResponse.Marshal(b, m, deterministic)
}
func (m *FundingListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingListResponse.Merge(m, src)
}
func (m *FundingListResponse) XXX_Size() int {
	return xxx_messageInfo_FundingListResponse.Size(m)
}
func (m *FundingListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FundingListResponse proto.InternalMessageInfo

func (m *FundingListResponse) GetHeader() *document.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FundingListResponse) GetData() []*FundingResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type Request struct {
	DocumentId           string   `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	AgreementId          string   `protobuf:"bytes,2,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{4}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func (m *Request) GetAgreementId() string {
	if m != nil {
		return m.AgreementId
	}
	return ""
}

type GetVersionRequest struct {
	DocumentId           string   `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	VersionId            string   `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	AgreementId          string   `protobuf:"bytes,3,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{5}
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

func (m *GetVersionRequest) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func (m *GetVersionRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *GetVersionRequest) GetAgreementId() string {
	if m != nil {
		return m.AgreementId
	}
	return ""
}

type GetListRequest struct {
	DocumentId           string   `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListRequest) Reset()         { *m = GetListRequest{} }
func (m *GetListRequest) String() string { return proto.CompactTextString(m) }
func (*GetListRequest) ProtoMessage()    {}
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{6}
}

func (m *GetListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListRequest.Unmarshal(m, b)
}
func (m *GetListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListRequest.Marshal(b, m, deterministic)
}
func (m *GetListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListRequest.Merge(m, src)
}
func (m *GetListRequest) XXX_Size() int {
	return xxx_messageInfo_GetListRequest.Size(m)
}
func (m *GetListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListRequest proto.InternalMessageInfo

func (m *GetListRequest) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

type GetListVersionRequest struct {
	DocumentId           string   `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	VersionId            string   `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListVersionRequest) Reset()         { *m = GetListVersionRequest{} }
func (m *GetListVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetListVersionRequest) ProtoMessage()    {}
func (*GetListVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{7}
}

func (m *GetListVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListVersionRequest.Unmarshal(m, b)
}
func (m *GetListVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListVersionRequest.Marshal(b, m, deterministic)
}
func (m *GetListVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListVersionRequest.Merge(m, src)
}
func (m *GetListVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetListVersionRequest.Size(m)
}
func (m *GetListVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListVersionRequest proto.InternalMessageInfo

func (m *GetListVersionRequest) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

func (m *GetListVersionRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

// FundingData is the default funding extension schema
type FundingData struct {
	AgreementId           string   `protobuf:"bytes,1,opt,name=agreement_id,json=agreementId,proto3" json:"agreement_id,omitempty"`
	Amount                string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Apr                   string   `protobuf:"bytes,3,opt,name=apr,proto3" json:"apr,omitempty"`
	Days                  string   `protobuf:"bytes,4,opt,name=days,proto3" json:"days,omitempty"`
	Fee                   string   `protobuf:"bytes,5,opt,name=fee,proto3" json:"fee,omitempty"`
	RepaymentDueDate      string   `protobuf:"bytes,7,opt,name=repayment_due_date,json=repaymentDueDate,proto3" json:"repayment_due_date,omitempty"`
	RepaymentOccurredDate string   `protobuf:"bytes,8,opt,name=repayment_occurred_date,json=repaymentOccurredDate,proto3" json:"repayment_occurred_date,omitempty"`
	RepaymentAmount       string   `protobuf:"bytes,9,opt,name=repayment_amount,json=repaymentAmount,proto3" json:"repayment_amount,omitempty"`
	Currency              string   `protobuf:"bytes,10,opt,name=currency,proto3" json:"currency,omitempty"`
	NftAddress            string   `protobuf:"bytes,11,opt,name=nft_address,json=nftAddress,proto3" json:"nft_address,omitempty"`
	PaymentDetailsId      string   `protobuf:"bytes,12,opt,name=payment_details_id,json=paymentDetailsId,proto3" json:"payment_details_id,omitempty"`
	FunderId              string   `protobuf:"bytes,13,opt,name=funder_id,json=funderId,proto3" json:"funder_id,omitempty"`
	BorrowerId            string   `protobuf:"bytes,14,opt,name=borrower_id,json=borrowerId,proto3" json:"borrower_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *FundingData) Reset()         { *m = FundingData{} }
func (m *FundingData) String() string { return proto.CompactTextString(m) }
func (*FundingData) ProtoMessage()    {}
func (*FundingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{8}
}

func (m *FundingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingData.Unmarshal(m, b)
}
func (m *FundingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingData.Marshal(b, m, deterministic)
}
func (m *FundingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingData.Merge(m, src)
}
func (m *FundingData) XXX_Size() int {
	return xxx_messageInfo_FundingData.Size(m)
}
func (m *FundingData) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingData.DiscardUnknown(m)
}

var xxx_messageInfo_FundingData proto.InternalMessageInfo

func (m *FundingData) GetAgreementId() string {
	if m != nil {
		return m.AgreementId
	}
	return ""
}

func (m *FundingData) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *FundingData) GetApr() string {
	if m != nil {
		return m.Apr
	}
	return ""
}

func (m *FundingData) GetDays() string {
	if m != nil {
		return m.Days
	}
	return ""
}

func (m *FundingData) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *FundingData) GetRepaymentDueDate() string {
	if m != nil {
		return m.RepaymentDueDate
	}
	return ""
}

func (m *FundingData) GetRepaymentOccurredDate() string {
	if m != nil {
		return m.RepaymentOccurredDate
	}
	return ""
}

func (m *FundingData) GetRepaymentAmount() string {
	if m != nil {
		return m.RepaymentAmount
	}
	return ""
}

func (m *FundingData) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *FundingData) GetNftAddress() string {
	if m != nil {
		return m.NftAddress
	}
	return ""
}

func (m *FundingData) GetPaymentDetailsId() string {
	if m != nil {
		return m.PaymentDetailsId
	}
	return ""
}

func (m *FundingData) GetFunderId() string {
	if m != nil {
		return m.FunderId
	}
	return ""
}

func (m *FundingData) GetBorrowerId() string {
	if m != nil {
		return m.BorrowerId
	}
	return ""
}

type FundingResponseData struct {
	Funding              *FundingData        `protobuf:"bytes,1,opt,name=funding,proto3" json:"funding,omitempty"`
	Signatures           []*FundingSignature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FundingResponseData) Reset()         { *m = FundingResponseData{} }
func (m *FundingResponseData) String() string { return proto.CompactTextString(m) }
func (*FundingResponseData) ProtoMessage()    {}
func (*FundingResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{9}
}

func (m *FundingResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingResponseData.Unmarshal(m, b)
}
func (m *FundingResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingResponseData.Marshal(b, m, deterministic)
}
func (m *FundingResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingResponseData.Merge(m, src)
}
func (m *FundingResponseData) XXX_Size() int {
	return xxx_messageInfo_FundingResponseData.Size(m)
}
func (m *FundingResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_FundingResponseData proto.InternalMessageInfo

func (m *FundingResponseData) GetFunding() *FundingData {
	if m != nil {
		return m.Funding
	}
	return nil
}

func (m *FundingResponseData) GetSignatures() []*FundingSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type FundingSignature struct {
	Valid                string   `protobuf:"bytes,1,opt,name=valid,proto3" json:"valid,omitempty"`
	OutdatedSignature    string   `protobuf:"bytes,2,opt,name=outdated_signature,json=outdatedSignature,proto3" json:"outdated_signature,omitempty"`
	Identity             string   `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	SignedVersion        string   `protobuf:"bytes,4,opt,name=signed_version,json=signedVersion,proto3" json:"signed_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FundingSignature) Reset()         { *m = FundingSignature{} }
func (m *FundingSignature) String() string { return proto.CompactTextString(m) }
func (*FundingSignature) ProtoMessage()    {}
func (*FundingSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_e74e73a4c22d632c, []int{10}
}

func (m *FundingSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FundingSignature.Unmarshal(m, b)
}
func (m *FundingSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FundingSignature.Marshal(b, m, deterministic)
}
func (m *FundingSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingSignature.Merge(m, src)
}
func (m *FundingSignature) XXX_Size() int {
	return xxx_messageInfo_FundingSignature.Size(m)
}
func (m *FundingSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingSignature.DiscardUnknown(m)
}

var xxx_messageInfo_FundingSignature proto.InternalMessageInfo

func (m *FundingSignature) GetValid() string {
	if m != nil {
		return m.Valid
	}
	return ""
}

func (m *FundingSignature) GetOutdatedSignature() string {
	if m != nil {
		return m.OutdatedSignature
	}
	return ""
}

func (m *FundingSignature) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *FundingSignature) GetSignedVersion() string {
	if m != nil {
		return m.SignedVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*FundingCreatePayload)(nil), "fun.FundingCreatePayload")
	proto.RegisterType((*FundingUpdatePayload)(nil), "fun.FundingUpdatePayload")
	proto.RegisterType((*FundingResponse)(nil), "fun.FundingResponse")
	proto.RegisterType((*FundingListResponse)(nil), "fun.FundingListResponse")
	proto.RegisterType((*Request)(nil), "fun.Request")
	proto.RegisterType((*GetVersionRequest)(nil), "fun.GetVersionRequest")
	proto.RegisterType((*GetListRequest)(nil), "fun.GetListRequest")
	proto.RegisterType((*GetListVersionRequest)(nil), "fun.GetListVersionRequest")
	proto.RegisterType((*FundingData)(nil), "fun.FundingData")
	proto.RegisterType((*FundingResponseData)(nil), "fun.FundingResponseData")
	proto.RegisterType((*FundingSignature)(nil), "fun.FundingSignature")
}

func init() { proto.RegisterFile("funding/funding.proto", fileDescriptor_e74e73a4c22d632c) }

var fileDescriptor_e74e73a4c22d632c = []byte{
	// 986 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0xd6, 0x89, 0x1d, 0x3f, 0xa7, 0x6e, 0x3a, 0x4d, 0x82, 0x59, 0x0a, 0x2c, 0x2b, 0x2a,
	0xda, 0x92, 0x78, 0x93, 0x54, 0x70, 0x40, 0x42, 0x68, 0x43, 0x54, 0x13, 0x09, 0x44, 0xe4, 0xaa,
	0x20, 0x55, 0x42, 0xd6, 0xc4, 0x33, 0x6b, 0x16, 0x39, 0x33, 0xee, 0xce, 0xac, 0x83, 0x15, 0x55,
	0x02, 0xee, 0x5c, 0x8c, 0xb8, 0x70, 0x43, 0xf0, 0x05, 0xb8, 0xf2, 0x11, 0xb8, 0xf2, 0x15, 0x90,
	0xf8, 0x1a, 0x68, 0xfe, 0xec, 0x7a, 0x1d, 0xc7, 0xc1, 0x51, 0x73, 0xb2, 0xe7, 0xbd, 0xdf, 0xbc,
	0xf7, 0x7b, 0x6f, 0xde, 0x9f, 0x85, 0x8d, 0x28, 0x65, 0x24, 0x66, 0xbd, 0xc0, 0xfe, 0x36, 0x07,
	0x09, 0x97, 0x1c, 0x95, 0xa2, 0x94, 0xb9, 0x9b, 0x84, 0x77, 0xd3, 0x13, 0xca, 0x64, 0x20, 0x68,
	0x32, 0x8c, 0xbb, 0xd4, 0x28, 0xdd, 0xbb, 0x3d, 0xce, 0x7b, 0x7d, 0x1a, 0xe0, 0x41, 0x1c, 0x60,
	0xc6, 0xb8, 0xc4, 0x32, 0xe6, 0x4c, 0x58, 0xed, 0x96, 0xfe, 0xe9, 0x6e, 0xf7, 0x28, 0xdb, 0x16,
	0xa7, 0xb8, 0xd7, 0xa3, 0x49, 0xc0, 0x07, 0x1a, 0x31, 0x8b, 0xf6, 0xbf, 0x82, 0xf5, 0xc7, 0xc6,
	0xf3, 0xc7, 0x09, 0xc5, 0x92, 0x1e, 0xe1, 0x51, 0x9f, 0x63, 0x82, 0xde, 0x84, 0x5a, 0xe6, 0xbd,
	0x13, 0x93, 0x86, 0xe3, 0x39, 0xf7, 0xab, 0x6d, 0xc8, 0x44, 0x87, 0x04, 0xbd, 0x0d, 0x4b, 0x04,
	0x4b, 0xdc, 0xb8, 0xe1, 0x39, 0xf7, 0x6b, 0x7b, 0x6b, 0xcd, 0x28, 0x65, 0x4d, 0x6b, 0xe9, 0x00,
	0x4b, 0xdc, 0xd6, 0x5a, 0xff, 0x3b, 0x27, 0xb7, 0xff, 0x74, 0x40, 0xae, 0x62, 0xff, 0x2d, 0x58,
	0xc5, 0xbd, 0x84, 0xd2, 0x0c, 0x71, 0x43, 0x23, 0x6a, 0xb9, 0xac, 0x40, 0xa1, 0x74, 0x29, 0x85,
	0xe7, 0x70, 0xcb, 0x0a, 0xdb, 0x54, 0x0c, 0x38, 0x13, 0x14, 0xed, 0x40, 0xf9, 0x6b, 0x8a, 0x09,
	0x4d, 0xb4, 0xdf, 0xda, 0x5e, 0xa3, 0x99, 0x39, 0x6e, 0x66, 0x98, 0x4f, 0xb4, 0xbe, 0x6d, 0x71,
	0x68, 0x6b, 0x2a, 0xda, 0x46, 0xd1, 0x55, 0x76, 0xa3, 0xe0, 0x32, 0x85, 0x3b, 0x56, 0xf9, 0x69,
	0x2c, 0xe4, 0xb5, 0xb8, 0x2d, 0x2d, 0xe0, 0xf6, 0x33, 0xa8, 0xb4, 0xe9, 0xf3, 0x94, 0x0a, 0x79,
	0x1d, 0xe9, 0xf5, 0x87, 0x70, 0xbb, 0x45, 0xe5, 0x17, 0x34, 0x11, 0x31, 0x67, 0x0b, 0x1b, 0x7e,
	0x1d, 0x60, 0x68, 0xae, 0x4c, 0xcc, 0x56, 0xad, 0xe4, 0x02, 0xbf, 0xa5, 0x59, 0xbf, 0xbb, 0x50,
	0x6f, 0x51, 0x69, 0x32, 0xb7, 0x98, 0x53, 0xff, 0x4b, 0xd8, 0xb0, 0x57, 0xae, 0x97, 0xae, 0xff,
	0x67, 0x09, 0x6a, 0x85, 0x92, 0x9a, 0xa1, 0xef, 0xcc, 0x56, 0xe5, 0x26, 0x94, 0xf1, 0x09, 0x4f,
	0x99, 0xb4, 0xd6, 0xec, 0x09, 0xad, 0x41, 0x09, 0x0f, 0x12, 0x1b, 0xb0, 0xfa, 0x8b, 0x90, 0x7a,
	0xdd, 0x91, 0x68, 0x2c, 0x69, 0x91, 0xfe, 0xaf, 0x50, 0x11, 0xa5, 0x8d, 0x65, 0x83, 0x8a, 0x28,
	0x45, 0x5b, 0x80, 0x12, 0x3a, 0xc0, 0x23, 0xed, 0x92, 0xa4, 0xb4, 0xa3, 0xda, 0xa8, 0x51, 0xd1,
	0x80, 0xb5, 0x5c, 0x73, 0x90, 0xaa, 0x4a, 0xa0, 0xe8, 0x7d, 0x78, 0x65, 0x82, 0xe6, 0xdd, 0x6e,
	0x9a, 0x24, 0x94, 0x98, 0x2b, 0x2b, 0xfa, 0xca, 0x46, 0xae, 0xfe, 0xdc, 0x6a, 0xf5, 0xbd, 0x07,
	0x30, 0xb1, 0xd5, 0xb1, 0xfc, 0xab, 0xfa, 0xc2, 0xad, 0x5c, 0x1e, 0x9a, 0x40, 0x5c, 0x58, 0xd1,
	0x17, 0x59, 0x77, 0xd4, 0x00, 0x0d, 0xc9, 0xcf, 0x2a, 0xdf, 0x2c, 0x92, 0x1d, 0x4c, 0x48, 0x42,
	0x85, 0x68, 0xd4, 0x4c, 0xbe, 0x59, 0x24, 0x43, 0x23, 0x51, 0xd1, 0xe4, 0xb1, 0x50, 0x89, 0xe3,
	0xbe, 0x50, 0x69, 0x5c, 0x35, 0xd1, 0x64, 0xb1, 0x18, 0xc5, 0x21, 0x41, 0xaf, 0x41, 0x55, 0xcd,
	0x45, 0x9a, 0x28, 0xd0, 0x4d, 0xe3, 0xcb, 0x08, 0x0e, 0xf5, 0x08, 0x39, 0xe6, 0x49, 0xc2, 0x4f,
	0x8d, 0xba, 0x6e, 0x7c, 0x65, 0xa2, 0x43, 0xe2, 0x7f, 0x9b, 0xb7, 0x61, 0xb1, 0x59, 0xd0, 0x43,
	0xa8, 0xd8, 0x61, 0x6b, 0xfb, 0x70, 0x76, 0x72, 0x64, 0x00, 0xf4, 0x1e, 0x80, 0x88, 0x7b, 0x0c,
	0xcb, 0x34, 0xa1, 0xc2, 0xb6, 0xe1, 0x46, 0x11, 0xfe, 0x24, 0xd3, 0xb6, 0x0b, 0x40, 0xff, 0x17,
	0x07, 0xd6, 0xce, 0x03, 0xd0, 0x3a, 0x2c, 0x0f, 0x71, 0x3f, 0x2f, 0x1a, 0x73, 0x40, 0xdb, 0x80,
	0x78, 0x2a, 0xd5, 0x03, 0x91, 0x4e, 0x6e, 0xc1, 0x96, 0xce, 0xed, 0x4c, 0x33, 0x31, 0xe2, 0xc2,
	0x4a, 0x4c, 0x28, 0x93, 0xb1, 0x1c, 0xd9, 0x52, 0xca, 0xcf, 0xe8, 0x1e, 0xd4, 0x95, 0x05, 0x4a,
	0x3a, 0xb6, 0x80, 0x6d, 0x65, 0xdd, 0x34, 0x52, 0xdb, 0x1a, 0x7b, 0xff, 0x56, 0xa1, 0x9e, 0x91,
	0x33, 0x7b, 0x05, 0xfd, 0xec, 0x40, 0xd9, 0xcc, 0x7f, 0xf4, 0x6a, 0x31, 0xba, 0xa9, 0x9d, 0xe0,
	0xae, 0x5f, 0x34, 0x7f, 0xfc, 0x67, 0xe3, 0xf0, 0x0d, 0xf7, 0x6e, 0x48, 0x88, 0xf0, 0xb0, 0x67,
	0xf3, 0xe6, 0x49, 0xee, 0x61, 0x2f, 0x6b, 0xb3, 0x1f, 0xfe, 0xfe, 0xe7, 0xa7, 0x1b, 0x8f, 0xfc,
	0x66, 0x30, 0xdc, 0x0d, 0x32, 0x99, 0x08, 0xce, 0x0a, 0x8d, 0xf9, 0x22, 0x5b, 0x7f, 0x9d, 0xbc,
	0x93, 0xc4, 0x07, 0xce, 0x43, 0xf4, 0x87, 0x03, 0x65, 0xb3, 0x37, 0xa6, 0x79, 0x4d, 0xed, 0x92,
	0x39, 0xbc, 0x86, 0xe3, 0xf0, 0x5d, 0xf7, 0x81, 0x41, 0x16, 0xa9, 0xe5, 0x1e, 0xbc, 0x98, 0x9d,
	0x27, 0xb9, 0xef, 0x7e, 0x78, 0x35, 0x92, 0xc1, 0x59, 0x71, 0x1c, 0xbc, 0x50, 0x9c, 0x7f, 0x73,
	0x60, 0x49, 0xbd, 0x17, 0x5a, 0xd5, 0xb4, 0xec, 0x24, 0x9a, 0x43, 0xf2, 0x74, 0x1c, 0xbe, 0xe3,
	0xde, 0x53, 0x70, 0xb1, 0x10, 0xc1, 0x96, 0xbf, 0xff, 0x52, 0x04, 0x03, 0x55, 0x0a, 0x8a, 0xe5,
	0xaf, 0x0e, 0x94, 0x5a, 0x54, 0x2e, 0x44, 0x72, 0x38, 0x0e, 0x9b, 0xee, 0x56, 0x8b, 0xca, 0x0b,
	0xb3, 0xc8, 0x23, 0x0f, 0x7b, 0x7d, 0x95, 0x64, 0x39, 0xcd, 0xf5, 0x23, 0xf4, 0x72, 0xc9, 0x44,
	0x7f, 0x39, 0x00, 0x93, 0x0d, 0x84, 0x36, 0x35, 0xb9, 0x99, 0x95, 0x34, 0x87, 0xf4, 0x8f, 0xce,
	0x38, 0x0c, 0xdc, 0xed, 0x4b, 0x59, 0x67, 0x7c, 0x3c, 0xdb, 0x2f, 0x9a, 0xf6, 0x53, 0xf4, 0xe4,
	0x32, 0xda, 0x16, 0x2a, 0x82, 0xb3, 0xc9, 0xde, 0x58, 0x24, 0x98, 0xdf, 0x1d, 0xa8, 0xd8, 0x1d,
	0x85, 0xee, 0x64, 0x91, 0x14, 0x96, 0x9c, 0x3b, 0xb5, 0xdd, 0x8b, 0xdf, 0x0d, 0xfe, 0x37, 0xe3,
	0x70, 0xd7, 0x0d, 0x74, 0x24, 0xfd, 0xfe, 0x6c, 0x2c, 0x62, 0xfe, 0x13, 0xec, 0xa0, 0x2b, 0x36,
	0x9d, 0xca, 0x79, 0x7d, 0x7a, 0x95, 0x22, 0xb7, 0xc8, 0xf6, 0x5c, 0xee, 0xe7, 0x93, 0xfe, 0xde,
	0x19, 0x87, 0x7b, 0xee, 0xce, 0xff, 0xb1, 0xbe, 0xf0, 0x09, 0x1e, 0xa3, 0x83, 0xeb, 0x78, 0x82,
	0x7d, 0x0f, 0x2a, 0x5d, 0x7e, 0xa2, 0x28, 0xee, 0xaf, 0x5a, 0x8e, 0x47, 0xea, 0xab, 0xf7, 0xc8,
	0x79, 0xb6, 0x1c, 0xa5, 0x6c, 0x70, 0x7c, 0x5c, 0xd6, 0x5f, 0xc1, 0x8f, 0xfe, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xf2, 0xff, 0xc2, 0xce, 0x87, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FundingServiceClient is the client API for FundingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FundingServiceClient interface {
	Create(ctx context.Context, in *FundingCreatePayload, opts ...grpc.CallOption) (*FundingResponse, error)
	Update(ctx context.Context, in *FundingUpdatePayload, opts ...grpc.CallOption) (*FundingResponse, error)
	Sign(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error)
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*FundingResponse, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*FundingListResponse, error)
	GetListVersion(ctx context.Context, in *GetListVersionRequest, opts ...grpc.CallOption) (*FundingListResponse, error)
}

type fundingServiceClient struct {
	cc *grpc.ClientConn
}

func NewFundingServiceClient(cc *grpc.ClientConn) FundingServiceClient {
	return &fundingServiceClient{cc}
}

func (c *fundingServiceClient) Create(ctx context.Context, in *FundingCreatePayload, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) Update(ctx context.Context, in *FundingUpdatePayload, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) Sign(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*FundingListResponse, error) {
	out := new(FundingListResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundingServiceClient) GetListVersion(ctx context.Context, in *GetListVersionRequest, opts ...grpc.CallOption) (*FundingListResponse, error) {
	out := new(FundingListResponse)
	err := c.cc.Invoke(ctx, "/fun.FundingService/GetListVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FundingServiceServer is the server API for FundingService service.
type FundingServiceServer interface {
	Create(context.Context, *FundingCreatePayload) (*FundingResponse, error)
	Update(context.Context, *FundingUpdatePayload) (*FundingResponse, error)
	Sign(context.Context, *Request) (*FundingResponse, error)
	Get(context.Context, *Request) (*FundingResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*FundingResponse, error)
	GetList(context.Context, *GetListRequest) (*FundingListResponse, error)
	GetListVersion(context.Context, *GetListVersionRequest) (*FundingListResponse, error)
}

func RegisterFundingServiceServer(s *grpc.Server, srv FundingServiceServer) {
	s.RegisterService(&_FundingService_serviceDesc, srv)
}

func _FundingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundingCreatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Create(ctx, req.(*FundingCreatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundingUpdatePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Update(ctx, req.(*FundingUpdatePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Sign(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundingService_GetListVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundingServiceServer).GetListVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fun.FundingService/GetListVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundingServiceServer).GetListVersion(ctx, req.(*GetListVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FundingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fun.FundingService",
	HandlerType: (*FundingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FundingService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FundingService_Update_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _FundingService_Sign_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _FundingService_Get_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _FundingService_GetVersion_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _FundingService_GetList_Handler,
		},
		{
			MethodName: "GetListVersion",
			Handler:    _FundingService_GetListVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "funding/funding.proto",
}
