// +build unit

package invoice

import (
	"context"
	"testing"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/contextutil"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/jobs"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/document"
	clientinvoicepb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/invoice"
	"github.com/centrifuge/go-centrifuge/testingutils/config"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	Service
	mock.Mock
}

func (m *mockService) DeriveFromCreatePayload(ctx context.Context, payload *clientinvoicepb.InvoiceCreatePayload) (documents.Model, error) {
	args := m.Called(ctx, payload)
	model, _ := args.Get(0).(documents.Model)
	return model, args.Error(1)
}

func (m *mockService) Create(ctx context.Context, model documents.Model) (documents.Model, jobs.JobID, chan bool, error) {
	args := m.Called(ctx, model)
	model, _ = args.Get(0).(documents.Model)
	return model, contextutil.Job(ctx), nil, args.Error(2)
}

func (m *mockService) GetCurrentVersion(ctx context.Context, documentID []byte) (documents.Model, error) {
	args := m.Called(ctx, documentID)
	data, _ := args.Get(0).(documents.Model)
	return data, args.Error(1)
}

func (m *mockService) GetVersion(ctx context.Context, documentID []byte, version []byte) (documents.Model, error) {
	args := m.Called(ctx, documentID, version)
	data, _ := args.Get(0).(documents.Model)
	return data, args.Error(1)
}

func (m *mockService) DeriveInvoiceData(doc documents.Model) (*clientinvoicepb.InvoiceData, error) {
	args := m.Called(doc)
	data, _ := args.Get(0).(*clientinvoicepb.InvoiceData)
	return data, args.Error(1)
}

func (m *mockService) DeriveInvoiceResponse(doc documents.Model) (*clientinvoicepb.InvoiceResponse, error) {
	args := m.Called(doc)
	data, _ := args.Get(0).(*clientinvoicepb.InvoiceResponse)
	return data, args.Error(1)
}

func (m *mockService) Update(ctx context.Context, model documents.Model) (documents.Model, jobs.JobID, chan bool, error) {
	args := m.Called(ctx, model)
	doc1, _ := args.Get(0).(documents.Model)
	return doc1, contextutil.Job(ctx), nil, args.Error(2)
}

func (m *mockService) DeriveFromUpdatePayload(ctx context.Context, payload *clientinvoicepb.InvoiceUpdatePayload) (documents.Model, error) {
	args := m.Called(ctx, payload)
	doc, _ := args.Get(0).(documents.Model)
	return doc, args.Error(1)
}

func getHandler() *grpcHandler {
	return &grpcHandler{service: &mockService{}, config: configService}
}

func TestGRPCHandler_Create_derive_fail(t *testing.T) {
	// DeriveFrom payload fails
	h := getHandler()
	srv := h.service.(*mockService)
	srv.On("DeriveFromCreatePayload", mock.Anything, mock.Anything).Return(nil, errors.New("derive failed")).Once()
	_, err := h.Create(testingconfig.HandlerContext(configService), nil)
	srv.AssertExpectations(t)
	assert.Error(t, err, "must be non nil")
	assert.Contains(t, err.Error(), "derive failed")
}

func TestGRPCHandler_Create_create_fail(t *testing.T) {
	h := getHandler()
	srv := h.service.(*mockService)
	srv.On("DeriveFromCreatePayload", mock.Anything, mock.Anything).Return(new(Invoice), nil).Once()
	srv.On("Create", mock.Anything, mock.Anything).Return(nil, jobs.NilJobID().String(), errors.New("create failed")).Once()
	payload := &clientinvoicepb.InvoiceCreatePayload{Data: &clientinvoicepb.InvoiceData{GrossAmount: "300"}}
	_, err := h.Create(testingconfig.HandlerContext(configService), payload)
	srv.AssertExpectations(t)
	assert.Error(t, err, "must be non nil")
	assert.Contains(t, err.Error(), "create failed")
}

type mockModel struct {
	documents.Model
	mock.Mock
	CoreDocument *coredocumentpb.CoreDocument
}

func (m *mockModel) ID() []byte {
	args := m.Called()
	id, _ := args.Get(0).([]byte)
	return id
}

func TestGRPCHandler_Create_DeriveInvoiceResponse_fail(t *testing.T) {
	h := getHandler()
	srv := h.service.(*mockService)
	model := new(Invoice)
	srv.On("DeriveFromCreatePayload", mock.Anything, mock.Anything).Return(model, nil).Once()
	srv.On("Create", mock.Anything, mock.Anything).Return(model, jobs.NilJobID().String(), nil).Once()
	srv.On("DeriveInvoiceResponse", mock.Anything).Return(nil, errors.New("derive response failed"))
	payload := &clientinvoicepb.InvoiceCreatePayload{Data: &clientinvoicepb.InvoiceData{Currency: "EUR"}}
	_, err := h.Create(testingconfig.HandlerContext(configService), payload)
	srv.AssertExpectations(t)
	assert.Error(t, err, "must be non nil")
	assert.Contains(t, err.Error(), "derive response failed")
}

func TestGrpcHandler_Create(t *testing.T) {
	h := getHandler()
	srv := h.service.(*mockService)
	model := new(Invoice)
	txID := jobs.NewJobID()
	payload := &clientinvoicepb.InvoiceCreatePayload{
		Data:        &clientinvoicepb.InvoiceData{GrossAmount: "300"},
		WriteAccess: []string{"0x010203040506"}}
	response := &clientinvoicepb.InvoiceResponse{Header: new(documentpb.ResponseHeader)}
	srv.On("DeriveFromCreatePayload", mock.Anything, mock.Anything).Return(model, nil).Once()
	srv.On("Create", mock.Anything, mock.Anything).Return(model, txID.String(), nil).Once()
	srv.On("DeriveInvoiceResponse", model).Return(response, nil)
	res, err := h.Create(testingconfig.HandlerContext(configService), payload)
	srv.AssertExpectations(t)
	assert.Nil(t, err, "must be nil")
	assert.NotNil(t, res, "must be non nil")
	assert.Equal(t, res, response)
}

func TestGrpcHandler_Get_invalid_input(t *testing.T) {
	identifier := "0x01010101"
	identifierBytes, _ := hexutil.Decode(identifier)
	h := getHandler()
	srv := h.service.(*mockService)
	payload := &clientinvoicepb.GetRequest{DocumentId: "invalid"}

	res, err := h.Get(testingconfig.HandlerContext(configService), payload)
	assert.Nil(t, res)
	assert.EqualError(t, err, "identifier is an invalid hex string: hex string without 0x prefix")

	payload.DocumentId = identifier
	srv.On("GetCurrentVersion", mock.Anything, identifierBytes).Return(nil, errors.New("not found"))
	res, err = h.Get(testingconfig.HandlerContext(configService), payload)
	srv.AssertExpectations(t)
	assert.Nil(t, res)
	assert.EqualError(t, err, "document not found: not found")
}

func TestGrpcHandler_Get(t *testing.T) {
	identifier := "0x01010101"
	identifierBytes, _ := hexutil.Decode(identifier)
	h := getHandler()
	srv := h.service.(*mockService)
	model := new(mockModel)
	payload := &clientinvoicepb.GetRequest{DocumentId: identifier}
	response := &clientinvoicepb.InvoiceResponse{}
	srv.On("GetCurrentVersion", mock.Anything, identifierBytes).Return(model, nil)
	srv.On("DeriveInvoiceResponse", model).Return(response, nil)
	res, err := h.Get(testingconfig.HandlerContext(configService), payload)
	model.AssertExpectations(t)
	srv.AssertExpectations(t)
	assert.Nil(t, err, "must be nil")
	assert.NotNil(t, res, "must be non nil")
	assert.Equal(t, res, response)
}

func TestGrpcHandler_GetVersion_invalid_input(t *testing.T) {
	h := getHandler()
	srv := h.service.(*mockService)
	payload := &clientinvoicepb.GetVersionRequest{DocumentId: "0x0x", VersionId: "0x00"}
	res, err := h.GetVersion(testingconfig.HandlerContext(configService), payload)
	assert.Error(t, err)
	assert.EqualError(t, err, "identifier is invalid: invalid hex string")
	payload.VersionId = "0x0x"
	payload.DocumentId = "0x01"

	res, err = h.GetVersion(testingconfig.HandlerContext(configService), payload)
	assert.Error(t, err)
	assert.EqualError(t, err, "version is invalid: invalid hex string")
	payload.VersionId = "0x00"
	payload.DocumentId = "0x01"

	mockErr := errors.New("not found")
	srv.On("GetVersion", mock.Anything, []byte{0x01}, []byte{0x00}).Return(nil, mockErr)
	res, err = h.GetVersion(testingconfig.HandlerContext(configService), payload)
	srv.AssertExpectations(t)
	assert.EqualError(t, err, "document not found: not found")
	assert.Nil(t, res)
}

func TestGrpcHandler_GetVersion(t *testing.T) {
	h := getHandler()
	srv := h.service.(*mockService)
	model := new(mockModel)
	payload := &clientinvoicepb.GetVersionRequest{DocumentId: "0x01", VersionId: "0x00"}

	response := &clientinvoicepb.InvoiceResponse{}
	srv.On("GetVersion", mock.Anything, []byte{0x01}, []byte{0x00}).Return(model, nil)
	srv.On("DeriveInvoiceResponse", model).Return(response, nil)
	res, err := h.GetVersion(testingconfig.HandlerContext(configService), payload)
	model.AssertExpectations(t)
	srv.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res, response)
}

func TestGrpcHandler_Update_derive_fail(t *testing.T) {
	h := getHandler()
	srv := h.service.(*mockService)
	payload := &clientinvoicepb.InvoiceUpdatePayload{DocumentId: "0x010201"}
	srv.On("DeriveFromUpdatePayload", mock.Anything, payload).Return(nil, errors.New("derive error")).Once()
	res, err := h.Update(testingconfig.HandlerContext(configService), payload)
	srv.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "derive error")
	assert.Nil(t, res)
}

func TestGrpcHandler_Update_update_fail(t *testing.T) {
	h := getHandler()
	srv := h.service.(*mockService)
	model := &mockModel{}
	ctx := testingconfig.HandlerContext(configService)
	payload := &clientinvoicepb.InvoiceUpdatePayload{DocumentId: "0x010201"}
	srv.On("DeriveFromUpdatePayload", mock.Anything, payload).Return(model, nil).Once()
	srv.On("Update", mock.Anything, model).Return(nil, jobs.NilJobID().String(), errors.New("update error")).Once()
	res, err := h.Update(ctx, payload)
	srv.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "update error")
	assert.Nil(t, res)
}

func TestGrpcHandler_Update_derive_response_fail(t *testing.T) {
	h := getHandler()
	srv := h.service.(*mockService)
	model := &mockModel{}
	ctx := testingconfig.HandlerContext(configService)
	payload := &clientinvoicepb.InvoiceUpdatePayload{DocumentId: "0x010201"}
	srv.On("DeriveFromUpdatePayload", mock.Anything, payload).Return(model, nil).Once()
	srv.On("Update", mock.Anything, model).Return(model, jobs.NilJobID().String(), nil).Once()
	srv.On("DeriveInvoiceResponse", model).Return(nil, errors.New("derive response error")).Once()
	res, err := h.Update(ctx, payload)
	srv.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "derive response error")
	assert.Nil(t, res)
}

func TestGrpcHandler_Update(t *testing.T) {
	h := getHandler()
	srv := h.service.(*mockService)
	model := &mockModel{}
	ctx := testingconfig.HandlerContext(configService)
	jobID := jobs.NewJobID()
	payload := &clientinvoicepb.InvoiceUpdatePayload{DocumentId: "0x010201"}
	resp := &clientinvoicepb.InvoiceResponse{Header: new(documentpb.ResponseHeader)}
	srv.On("DeriveFromUpdatePayload", mock.Anything, payload).Return(model, nil).Once()
	srv.On("Update", mock.Anything, model).Return(model, jobID.String(), nil).Once()
	srv.On("DeriveInvoiceResponse", model).Return(resp, nil).Once()
	res, err := h.Update(ctx, payload)
	srv.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, resp, res)
}
