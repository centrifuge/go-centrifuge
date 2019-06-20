// +build unit

package funding

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/bootstrap/bootstrappers/testlogging"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/config/configstore"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/documents/invoice"
	"github.com/centrifuge/go-centrifuge/ethereum"
	"github.com/centrifuge/go-centrifuge/extensions"
	"github.com/centrifuge/go-centrifuge/identity/ideth"
	"github.com/centrifuge/go-centrifuge/jobs"
	"github.com/centrifuge/go-centrifuge/p2p"
	clientfunpb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/funding"
	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/centrifuge/go-centrifuge/storage/leveldb"
	"github.com/centrifuge/go-centrifuge/testingutils/config"
	"github.com/centrifuge/go-centrifuge/testingutils/documents"
	"github.com/centrifuge/go-centrifuge/testingutils/identity"
	"github.com/centrifuge/go-centrifuge/testingutils/testingjobs"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var ctx = map[string]interface{}{}
var cfg config.Configuration

var (
	did = testingidentity.GenerateRandomDID()
)

func TestMain(m *testing.M) {
	ethClient := new(ethereum.MockEthClient)
	ethClient.On("GetEthClient").Return(nil)
	ctx[ethereum.BootstrappedEthereumClient] = ethClient
	jobMan := &testingjobs.MockJobManager{}
	ctx[jobs.BootstrappedService] = jobMan
	done := make(chan bool)
	jobMan.On("ExecuteWithinJob", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(jobs.NilJobID(), done, nil)
	ctx[bootstrap.BootstrappedInvoiceUnpaid] = new(testingdocuments.MockRegistry)
	ibootstrappers := []bootstrap.TestBootstrapper{
		&testlogging.TestLoggingBootstrapper{},
		&config.Bootstrapper{},
		&leveldb.Bootstrapper{},
		&queue.Bootstrapper{},
		&ideth.Bootstrapper{},
		&configstore.Bootstrapper{},
		anchors.Bootstrapper{},
		documents.Bootstrapper{},
		p2p.Bootstrapper{},
		documents.PostBootstrapper{},
		&queue.Starter{},
	}
	bootstrap.RunTestBootstrappers(ibootstrappers, ctx)
	cfg = ctx[bootstrap.BootstrappedConfig].(config.Configuration)
	cfg.Set("identityId", did.String())
	configService = ctx[config.BootstrappedConfigStorage].(config.Service)
	result := m.Run()
	bootstrap.RunTestTeardown(ibootstrappers)
	os.Exit(result)
}

func TestAttributesUtils(t *testing.T) {
	testingdocuments.CreateInvoicePayload()
	inv := new(invoice.Invoice)
	err := inv.InitInvoiceInput(testingdocuments.CreateInvoicePayload(), testingidentity.GenerateRandomDID())
	assert.NoError(t, err)
	docSrv := new(testingdocuments.MockService)
	docSrv.On("GetCurrentVersion", mock.Anything, mock.Anything).Return(inv, nil)
	srv := DefaultService(docSrv, nil)

	data := createTestData()

	// Fill attributes list
	a, err := extensions.FillAttributeList(data, "0", fundingFieldKey)
	assert.NoError(t, err)
	// fill attribute list does not add the idx of attribute set as an attribute
	assert.Len(t, a, 12)

	// Creating an attributes list generates the correct attributes and adds an idx as an attribute
	attributes, err := extensions.CreateAttributesList(inv, data, fundingFieldKey, fundingLabel)
	assert.NoError(t, err)
	assert.Len(t, attributes, 13)

	for _, attribute := range attributes {
		if attribute.KeyLabel == "funding_agreement[0].currency" {
			assert.Equal(t, "eur", attribute.Value.Str)
			break
		}

		// apr was not set
		assert.NotEqual(t, "funding_agreement[0].apr", attribute.KeyLabel)
	}

	// add attributes to Document
	err = inv.AddAttributes(documents.CollaboratorsAccess{}, true, attributes...)
	assert.NoError(t, err)

	var agreementID string
	for _, attribute := range attributes {
		if attribute.KeyLabel == "funding_agreement[0].agreement_id" {
			agreementID, err = attribute.Value.String()
			assert.NoError(t, err)
			break
		}
	}

	// wrong attributeSetID
	idx, err := extensions.FindAttributeSetIDX(inv, "randomID", fundingLabel, agreementIDLabel, fundingFieldKey)
	assert.Error(t, err)

	// correct
	idx, err = extensions.FindAttributeSetIDX(inv, agreementID, fundingLabel, agreementIDLabel, fundingFieldKey)
	assert.Equal(t, "0", idx)
	assert.NoError(t, err)

	// add second attributeSet
	data.AgreementId = extensions.NewAttributeSetID()
	a2, err := extensions.CreateAttributesList(inv, data, fundingFieldKey, fundingLabel)
	assert.NoError(t, err)

	var aID string
	for _, attribute := range a2 {
		if attribute.KeyLabel == "funding_agreement[1].agreement_id" {
			aID, err = attribute.Value.String()
			assert.NoError(t, err)
			//break
		}
	}

	err = inv.AddAttributes(documents.CollaboratorsAccess{}, true, a2...)
	assert.NoError(t, err)

	// latest idx
	model, err := srv.GetCurrentVersion(context.Background(), inv.Document.DocumentIdentifier)
	assert.NoError(t, err)

	lastIdx, err := extensions.GetArrayLatestIDX(model, fundingLabel)
	assert.NoError(t, err)

	n, err := documents.NewInt256("1")
	assert.NoError(t, err)
	assert.Equal(t, lastIdx, n)

	// index should be 1
	idx, err = extensions.FindAttributeSetIDX(inv, aID, fundingLabel, agreementIDLabel, fundingFieldKey)
	assert.Equal(t, "1", idx)
	assert.NoError(t, err)

	// delete the first attribute set
	idx, err = extensions.FindAttributeSetIDX(inv, agreementID, fundingLabel, agreementIDLabel, fundingFieldKey)
	assert.NoError(t, err)

	model, err = extensions.DeleteAttributesSet(model, Data{}, idx, fundingFieldKey)
	assert.NoError(t, err)
	assert.Len(t, model.GetAttributes(), 13)

	// error when trying to delete non existing attribute set
	idx, err = extensions.FindAttributeSetIDX(inv, agreementID, fundingLabel, agreementIDLabel, fundingFieldKey)
	assert.Error(t, err)

	// check that latest idx is still 1 even though the first set of attributes have been deleted ?
	latest, err := extensions.GetArrayLatestIDX(model, fundingLabel)
	assert.NoError(t, err)
	assert.Equal(t, latest, n)

	// non existent typeLabel for attribute set
	_, err = extensions.GetArrayLatestIDX(model, "randomLabel")
	assert.Error(t, err)

	// check that we can no longer find the attributes from the first set
	idx, err = extensions.FindAttributeSetIDX(inv, agreementID, fundingLabel, agreementIDLabel, fundingFieldKey)
	assert.Error(t, err)

	// test increment array attr idx
	n, err = documents.NewInt256("2")
	assert.NoError(t, err)

	newIdx, err := extensions.IncrementArrayAttrIDX(model, fundingLabel)
	assert.NoError(t, err)

	v, err := newIdx.Value.String()
	assert.NoError(t, err)
	assert.Equal(t, "2", v)
	assert.Equal(t, fundingLabel, newIdx.KeyLabel)
}

func TestDeriveFromPayload(t *testing.T) {
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	testingdocuments.CreateInvoicePayload()
	inv := new(invoice.Invoice)
	err := inv.InitInvoiceInput(testingdocuments.CreateInvoicePayload(), testingidentity.GenerateRandomDID())
	assert.NoError(t, err)

	docSrv := new(testingdocuments.MockService)
	docSrv.On("GetCurrentVersion", mock.Anything, mock.Anything).Return(inv, nil)
	srv := DefaultService(docSrv, nil)

	payload := createTestPayload()
	payload.DocumentId = hexutil.Encode(inv.Document.DocumentIdentifier)

	for i := 0; i < 10; i++ {
		model, err := srv.DeriveFromPayload(ctxh, payload)
		assert.NoError(t, err)
		label := fmt.Sprintf("funding_agreement[%d].currency", i)
		key, err := documents.AttrKeyFromLabel(label)
		assert.NoError(t, err)

		attr, err := model.GetAttribute(key)
		assert.NoError(t, err)
		assert.Equal(t, "eur", attr.Value.Str)
	}

	payload.DocumentId = ""
	_, err = srv.DeriveFromPayload(ctxh, payload)
	assert.Error(t, err)
}

func TestDeriveFundingResponse(t *testing.T) {
	testingdocuments.CreateInvoicePayload()
	inv := new(invoice.Invoice)
	err := inv.InitInvoiceInput(testingdocuments.CreateInvoicePayload(), testingidentity.GenerateRandomDID())
	assert.NoError(t, err)

	docSrv := new(testingdocuments.MockService)
	docSrv.On("GetCurrentVersion", mock.Anything, mock.Anything).Return(inv, nil)
	srv := DefaultService(docSrv, nil)

	ctxh := testingconfig.CreateAccountContext(t, cfg)

	for i := 0; i < 10; i++ {
		payload := createTestPayload()
		payload.DocumentId = hexutil.Encode(inv.Document.DocumentIdentifier)
		model, err := srv.DeriveFromPayload(context.Background(), payload)
		assert.NoError(t, err)

		response, err := srv.DeriveFundingResponse(ctxh, model, payload.Data.AgreementId)
		assert.NoError(t, err)
		checkResponse(t, payload, response.Data.Funding)
	}

}

func TestDeriveFundingListResponse(t *testing.T) {
	testingdocuments.CreateInvoicePayload()
	inv := new(invoice.Invoice)
	err := inv.InitInvoiceInput(testingdocuments.CreateInvoicePayload(), testingidentity.GenerateRandomDID())
	assert.NoError(t, err)

	docSrv := new(testingdocuments.MockService)
	docSrv.On("GetCurrentVersion", mock.Anything, mock.Anything).Return(inv, nil)
	srv := DefaultService(docSrv, nil)

	var model documents.Model
	var payloads []*clientfunpb.FundingCreatePayload
	for i := 0; i < 10; i++ {
		p := createTestPayload()
		p.DocumentId = hexutil.Encode(inv.Document.DocumentIdentifier)
		payloads = append(payloads, p)
		model, err = srv.DeriveFromPayload(context.Background(), p)
		assert.NoError(t, err)
	}

	response, err := srv.DeriveFundingListResponse(context.Background(), model)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(response.Data))

	for i := 0; i < 10; i++ {
		checkResponse(t, payloads[i], response.Data[i].Funding)
	}
}

func TestService_DeriveFromUpdatePayload(t *testing.T) {
	testingdocuments.CreateInvoicePayload()
	inv := new(invoice.Invoice)
	err := inv.InitInvoiceInput(testingdocuments.CreateInvoicePayload(), testingidentity.GenerateRandomDID())
	assert.NoError(t, err)

	docSrv := new(testingdocuments.MockService)
	docSrv.On("GetCurrentVersion", mock.Anything, mock.Anything).Return(inv, nil)
	srv := DefaultService(docSrv, nil)

	var model documents.Model
	p := createTestPayload()
	p.DocumentId = hexutil.Encode(inv.Document.DocumentIdentifier)
	model, err = srv.DeriveFromPayload(context.Background(), p)
	assert.NoError(t, err)

	// update
	docSrv.On("GetCurrentVersion", mock.Anything, mock.Anything).Return(model, nil)
	p2 := &clientfunpb.FundingUpdatePayload{Data: createTestClientData(), DocumentId: hexutil.Encode(utils.RandomSlice(32)), AgreementId: p.Data.AgreementId}
	p2.Data.Currency = ""
	p2.Data.Fee = "13.37"

	model, err = srv.DeriveFromUpdatePayload(context.Background(), p2)
	assert.NoError(t, err)

	response, err := srv.DeriveFundingListResponse(context.Background(), model)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, p2.Data.Fee, response.Data[0].Funding.Fee)

	// fee was not set in the update old fee field should not exist
	assert.NotEqual(t, p.Data.Fee, response.Data[0].Funding.Fee)

	// non existing funding id
	p3 := &clientfunpb.FundingUpdatePayload{Data: createTestClientData(), DocumentId: hexutil.Encode(utils.RandomSlice(32)), AgreementId: hexutil.Encode(utils.RandomSlice(32))}
	model, err = srv.DeriveFromUpdatePayload(context.Background(), p3)
	assert.Error(t, err)
	assert.Contains(t, err, extensions.ErrAttributeSetNotFound)

	p2.DocumentId = ""
	_, err = srv.DeriveFromUpdatePayload(context.Background(), p2)
	assert.Error(t, err)
}

func createTestClientData() *clientfunpb.FundingData {
	fundingId := extensions.NewAttributeSetID()
	return &clientfunpb.FundingData{
		AgreementId:           fundingId,
		Currency:              "eur",
		Days:                  "90",
		Amount:                "1000",
		RepaymentAmount:       "1200.12",
		Fee:                   "10",
		BorrowerId:            testingidentity.GenerateRandomDID().String(),
		FunderId:              testingidentity.GenerateRandomDID().String(),
		NftAddress:            hexutil.Encode(utils.RandomSlice(32)),
		RepaymentDueDate:      time.Now().UTC().Format(time.RFC3339),
		RepaymentOccurredDate: time.Now().UTC().Format(time.RFC3339),
		PaymentDetailsId:      hexutil.Encode(utils.RandomSlice(32)),
	}
}

func createTestData() Data {
	fundingId := extensions.NewAttributeSetID()
	return Data{
		AgreementId:           fundingId,
		Currency:              "eur",
		Days:                  "90",
		Amount:                "1000",
		RepaymentAmount:       "1200.12",
		Fee:                   "10",
		BorrowerId:            testingidentity.GenerateRandomDID().String(),
		FunderId:              testingidentity.GenerateRandomDID().String(),
		NftAddress:            hexutil.Encode(utils.RandomSlice(32)),
		RepaymentDueDate:      time.Now().UTC().Format(time.RFC3339),
		RepaymentOccurredDate: time.Now().UTC().Format(time.RFC3339),
		PaymentDetailsId:      hexutil.Encode(utils.RandomSlice(32)),
	}
}

func createTestPayload() *clientfunpb.FundingCreatePayload {
	return &clientfunpb.FundingCreatePayload{Data: createTestClientData()}
}

func checkResponse(t *testing.T, payload *clientfunpb.FundingCreatePayload, response *clientfunpb.FundingData) {
	assert.Equal(t, payload.Data.AgreementId, response.AgreementId)
	assert.Equal(t, payload.Data.Currency, response.Currency)
	assert.Equal(t, payload.Data.Days, response.Days)
	assert.Equal(t, payload.Data.Amount, response.Amount)
	assert.Equal(t, payload.Data.RepaymentDueDate, response.RepaymentDueDate)
}
