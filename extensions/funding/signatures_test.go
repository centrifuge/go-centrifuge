// +build unit

package funding

import (
	"context"
	"fmt"
	"testing"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/contextutil"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/documents/invoice"
	clientfunpb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/funding"
	"github.com/centrifuge/go-centrifuge/testingutils/documents"
	"github.com/centrifuge/go-centrifuge/testingutils/identity"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAccount struct {
	config.Account
	mock.Mock
}

func (m *mockAccount) SignMsg(msg []byte) ([]*coredocumentpb.Signature, error) {
	args := m.Called(msg)
	sig, _ := args.Get(0).([]*coredocumentpb.Signature)
	return sig, args.Error(1)
}

func (m *mockAccount) GetIdentityID() ([]byte, error) {
	args := m.Called()
	sig, _ := args.Get(0).([]byte)
	return sig, args.Error(1)
}

func setupFundingForTesting(t *testing.T, fundingAmount int) (Service, *testingdocuments.MockService, documents.Model, string) {
	testingdocuments.CreateInvoicePayload()
	inv := new(invoice.Invoice)
	err := inv.InitInvoiceInput(testingdocuments.CreateInvoicePayload(), testingidentity.GenerateRandomDID())
	assert.NoError(t, err)

	docSrv := new(testingdocuments.MockService)
	docSrv.On("GetCurrentVersion", mock.Anything, mock.Anything).Return(inv, nil)

	srv := DefaultService(docSrv, nil)

	var model documents.Model
	var payloads []*clientfunpb.FundingCreatePayload
	var lastFundingId string

	// create a list of fundings
	for i := 0; i < fundingAmount; i++ {
		p := createTestPayload()
		p.DocumentId = hexutil.Encode(inv.Document.DocumentIdentifier)
		payloads = append(payloads, p)
		model, err = srv.DeriveFromPayload(context.Background(), p)
		assert.NoError(t, err)
		lastFundingId = p.Data.AgreementId
	}

	return srv, docSrv, model, lastFundingId
}

func TestService_Sign(t *testing.T) {
	fundingAmount := 5
	srv, _, model, lastFundingId := setupFundingForTesting(t, fundingAmount)

	// add signature
	acc := new(mockAccount)
	acc.On("GetIdentityID").Return(utils.RandomSlice(20), nil)
	// success
	signature := utils.RandomSlice(32)
	acc.On("SignMsg", mock.Anything).Return([]*coredocumentpb.Signature{{Signature: signature}}, nil)
	ctx, err := contextutil.New(context.Background(), acc)
	assert.NoError(t, err)

	for i := 0; i < 5; i++ {
		model, err = srv.Sign(ctx, lastFundingId, utils.RandomSlice(32))
		assert.NoError(t, err)
		// signature should exist
		label := fmt.Sprintf("funding_agreement[%d].signatures[%d]", fundingAmount-1, i)
		key, err := documents.AttrKeyFromLabel(label)
		assert.NoError(t, err)
		attr, err := model.GetAttribute(key)
		assert.NoError(t, err)
		assert.Equal(t, documents.AttrSigned, attr.Value.Type)
		assert.Equal(t, signature, attr.Value.Signed.Signature)

		// array idx should exist
		label = fmt.Sprintf("funding_agreement[%d].signatures", fundingAmount-1)
		key, err = documents.AttrKeyFromLabel(label)
		assert.NoError(t, err)
		attr, err = model.GetAttribute(key)
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("%d", i), attr.Value.Int256.String())

	}

	// funding id not exists
	model, err = srv.Sign(ctx, hexutil.Encode(utils.RandomSlice(32)), utils.RandomSlice(32))
	assert.Error(t, err)
}

func TestService_SignVerify(t *testing.T) {
	fundingAmount := 5
	srv, docSrv, model, fundingID := setupFundingForTesting(t, fundingAmount)

	// add signature
	acc := new(mockAccount)
	acc.On("GetIdentityID").Return(utils.RandomSlice(20), nil)
	// success
	signature := utils.RandomSlice(32)
	acc.On("SignMsg", mock.Anything).Return([]*coredocumentpb.Signature{{Signature: signature, PublicKey: utils.RandomSlice(64)}}, nil)
	ctx, err := contextutil.New(context.Background(), acc)
	assert.NoError(t, err)

	// add signature
	model, err = srv.Sign(ctx, fundingID, utils.RandomSlice(32))
	assert.NoError(t, err)

	// funding current version: valid
	response, err := srv.DeriveFundingResponse(ctx, model, fundingID)
	assert.NoError(t, err)
	assert.Equal(t, "true", response.Data.Signatures[0].Valid)
	assert.Equal(t, "false", response.Data.Signatures[0].OutdatedSignature)

	// update funding after signature
	oldCD, err := model.PackCoreDocument()
	assert.NoError(t, err)
	oldInv := new(invoice.Invoice)
	err = oldInv.UnpackCoreDocument(oldCD)
	assert.NoError(t, err)

	p2 := &clientfunpb.FundingUpdatePayload{Data: createTestClientData(), DocumentId: hexutil.Encode(utils.RandomSlice(32)), AgreementId: fundingID}
	p2.Data.Currency = ""
	p2.Data.Fee = "13.37"
	updatedModel, err := srv.DeriveFromUpdatePayload(context.Background(), p2)
	assert.NoError(t, err)

	// older funding version signed: valid
	docSrv.On("GetVersion", mock.Anything, mock.Anything).Return(oldInv, nil).Once()
	response, err = srv.DeriveFundingResponse(ctx, updatedModel, fundingID)
	assert.NoError(t, err)
	assert.Equal(t, "true", response.Data.Signatures[0].Valid)
	assert.Equal(t, "true", response.Data.Signatures[0].OutdatedSignature)

	// older funding version signed: invalid
	invalidValue, err := hexutil.Decode("0x1234")
	assert.NoError(t, err)
	attr, err := documents.NewSignedAttribute("funding_agreement[4].signatures[0]", testingidentity.GenerateRandomDID(), acc, model, invalidValue)
	assert.NoError(t, err)
	err = oldInv.AddAttributes(documents.CollaboratorsAccess{}, true, attr)
	assert.NoError(t, err)

	docSrv.On("GetVersion", mock.Anything, mock.Anything).Return(oldInv, nil)
	response, err = srv.DeriveFundingResponse(ctx, oldInv, fundingID)
	assert.NoError(t, err)
	assert.Equal(t, "false", response.Data.Signatures[0].Valid)
	assert.Equal(t, "true", response.Data.Signatures[0].OutdatedSignature)
}
