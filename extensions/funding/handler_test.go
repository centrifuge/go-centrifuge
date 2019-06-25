// +build unit

package funding

import (
	"context"
	"testing"

	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/contextutil"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/jobs"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/document"
	clientfunpb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/funding"
	"github.com/centrifuge/go-centrifuge/testingutils/config"
	"github.com/centrifuge/go-centrifuge/testingutils/documents"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	Service
	mock.Mock
}

var configService config.Service

func (m *mockService) DeriveFromPayload(ctx context.Context, req *clientfunpb.FundingCreatePayload) (documents.Model, error) {
	args := m.Called(ctx, req)
	model, _ := args.Get(0).(documents.Model)
	return model, args.Error(1)
}

func (m *mockService) DeriveFromUpdatePayload(ctx context.Context, req *clientfunpb.FundingUpdatePayload) (documents.Model, error) {
	args := m.Called(ctx, req)
	model, _ := args.Get(0).(documents.Model)
	return model, args.Error(1)
}

func (m *mockService) DeriveFundingResponse(ctx context.Context, doc documents.Model, fundingId string) (*clientfunpb.FundingResponse, error) {
	args := m.Called(doc)
	data, _ := args.Get(0).(*clientfunpb.FundingResponse)
	return data, args.Error(1)
}

func (m *mockService) DeriveFundingListResponse(ctx context.Context, doc documents.Model) (*clientfunpb.FundingListResponse, error) {
	args := m.Called(doc)
	data, _ := args.Get(0).(*clientfunpb.FundingListResponse)
	return data, args.Error(1)
}

func (m *mockService) Update(ctx context.Context, model documents.Model) (documents.Model, jobs.JobID, chan bool, error) {
	args := m.Called(ctx, model)
	doc1, _ := args.Get(0).(documents.Model)
	return doc1, contextutil.Job(ctx), nil, args.Error(2)
}

func (m *mockService) GetCurrentVersion(ctx context.Context, identifier []byte) (documents.Model, error) {
	args := m.Called(ctx, identifier)
	model, _ := args.Get(0).(documents.Model)
	return model, args.Error(1)
}

func (m *mockService) GetVersion(ctx context.Context, identifier, version []byte) (documents.Model, error) {
	args := m.Called(ctx, identifier)
	model, _ := args.Get(0).(documents.Model)
	return model, args.Error(1)
}

func (m *mockService) Sign(ctx context.Context, fundingID string, identifier []byte) (documents.Model, error) {
	args := m.Called(ctx, fundingID, identifier)
	model, _ := args.Get(0).(documents.Model)
	return model, args.Error(1)
}

func TestGRPCHandler_Create(t *testing.T) {
	srv := &mockService{}

	h := &grpcHandler{service: srv, config: configService}
	jobID := jobs.NewJobID()

	// successful
	srv.On("DeriveFromPayload", mock.Anything, mock.Anything).Return(documents.ErrDocumentIdentifier, nil)
	srv.On("Update", mock.Anything, mock.Anything).Return(nil, jobID, nil).Once()
	srv.On("DeriveFundingResponse", mock.Anything, mock.Anything, mock.Anything).Return(&clientfunpb.FundingResponse{Header: new(documentpb.ResponseHeader)}, nil).Once()
	response, err := h.Create(testingconfig.HandlerContext(configService), &clientfunpb.FundingCreatePayload{DocumentId: hexutil.Encode(utils.RandomSlice(32)), Data: &clientfunpb.FundingData{Currency: "eur"}})
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestGRPCHandler_Update(t *testing.T) {
	srv := &mockService{}

	h := &grpcHandler{service: srv, config: configService}
	jobID := jobs.NewJobID()

	// successful
	srv.On("DeriveFromUpdatePayload", mock.Anything, mock.Anything).Return(&testingdocuments.MockModel{}, nil)
	srv.On("Update", mock.Anything, mock.Anything).Return(nil, jobID, nil).Once()
	srv.On("DeriveFundingResponse", mock.Anything, mock.Anything, mock.Anything).Return(&clientfunpb.FundingResponse{Header: new(documentpb.ResponseHeader)}, nil).Once()

	response, err := h.Update(testingconfig.HandlerContext(configService), &clientfunpb.FundingUpdatePayload{DocumentId: hexutil.Encode(utils.RandomSlice(32)), Data: &clientfunpb.FundingData{Currency: "eur"}})
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestGRPCHandler_Sign(t *testing.T) {
	srv := &mockService{}
	h := &grpcHandler{service: srv, config: configService}
	jobID := jobs.NewJobID()

	// successful
	srv.On("GetCurrentVersion", mock.Anything, mock.Anything).Return(&testingdocuments.MockModel{}, nil)
	srv.On("DeriveFundingResponse", mock.Anything, mock.Anything, mock.Anything).Return(&clientfunpb.FundingResponse{Header: new(documentpb.ResponseHeader)}, nil).Once()
	srv.On("Sign", mock.Anything, mock.Anything, mock.Anything).Return(&testingdocuments.MockModel{}, nil).Once()
	srv.On("Update", mock.Anything, mock.Anything).Return(nil, jobID, nil).Once()

	response, err := h.Sign(testingconfig.HandlerContext(configService), &clientfunpb.Request{DocumentId: hexutil.Encode(utils.RandomSlice(32)), AgreementId: hexutil.Encode(utils.RandomSlice(32))})
	assert.NoError(t, err)
	assert.NotNil(t, response)

	// fail
	response, err = h.Sign(testingconfig.HandlerContext(configService), &clientfunpb.Request{AgreementId: hexutil.Encode(utils.RandomSlice(32))})
	assert.Error(t, err)
}

func TestGRPCHandler_Get(t *testing.T) {
	srv := &mockService{}
	h := &grpcHandler{service: srv, config: configService}

	srv.On("GetCurrentVersion", mock.Anything, mock.Anything).Return(&testingdocuments.MockModel{}, nil)
	srv.On("DeriveFundingResponse", mock.Anything, mock.Anything, mock.Anything).Return(&clientfunpb.FundingResponse{Header: new(documentpb.ResponseHeader)}, nil).Once()

	response, err := h.Get(testingconfig.HandlerContext(configService), &clientfunpb.Request{DocumentId: hexutil.Encode(utils.RandomSlice(32)), AgreementId: hexutil.Encode(utils.RandomSlice(32))})
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestGRPCHandler_GetVersion(t *testing.T) {
	srv := &mockService{}
	h := &grpcHandler{service: srv, config: configService}

	srv.On("GetVersion", mock.Anything, mock.Anything, mock.Anything).Return(&testingdocuments.MockModel{}, nil)
	srv.On("DeriveFundingResponse", mock.Anything, mock.Anything, mock.Anything).Return(&clientfunpb.FundingResponse{Header: new(documentpb.ResponseHeader)}, nil).Once()

	response, err := h.GetVersion(testingconfig.HandlerContext(configService), &clientfunpb.GetVersionRequest{DocumentId: hexutil.Encode(utils.RandomSlice(32)), VersionId: hexutil.Encode(utils.RandomSlice(32)), AgreementId: hexutil.Encode(utils.RandomSlice(32))})
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestGRPCHandler_GetList(t *testing.T) {
	srv := &mockService{}
	h := &grpcHandler{service: srv, config: configService}

	srv.On("GetVersion", mock.Anything, mock.Anything, mock.Anything).Return(&testingdocuments.MockModel{}, nil)
	srv.On("DeriveFundingListResponse", mock.Anything, mock.Anything).Return(&clientfunpb.FundingListResponse{Header: new(documentpb.ResponseHeader)}, nil).Once()

	response, err := h.GetListVersion(testingconfig.HandlerContext(configService), &clientfunpb.GetListVersionRequest{DocumentId: hexutil.Encode(utils.RandomSlice(32)), VersionId: hexutil.Encode(utils.RandomSlice(32))})
	assert.NoError(t, err)
	assert.NotNil(t, response)
}
