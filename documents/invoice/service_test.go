// +build unit

package invoice

import (
	"context"
	"testing"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/jobs"
	clientinvoicepb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/invoice"
	"github.com/centrifuge/go-centrifuge/storage"
	"github.com/centrifuge/go-centrifuge/storage/leveldb"
	"github.com/centrifuge/go-centrifuge/testingutils"
	"github.com/centrifuge/go-centrifuge/testingutils/anchors"
	"github.com/centrifuge/go-centrifuge/testingutils/commons"
	"github.com/centrifuge/go-centrifuge/testingutils/config"
	"github.com/centrifuge/go-centrifuge/testingutils/documents"
	"github.com/centrifuge/go-centrifuge/testingutils/identity"
	"github.com/centrifuge/go-centrifuge/testingutils/testingjobs"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/centrifuge/gocelery"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	did       = testingidentity.GenerateRandomDID()
	didBytes  = did[:]
	accountID = did[:]
)

func getServiceWithMockedLayers() (testingcommons.MockIdentityService, Service) {
	c := &testingconfig.MockConfig{}
	c.On("GetIdentityID").Return(didBytes, nil)
	idService := testingcommons.MockIdentityService{}
	idService.On("IsSignedWithPurpose", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(true, nil).Once()
	queueSrv := new(testingutils.MockQueue)
	queueSrv.On("EnqueueJob", mock.Anything, mock.Anything).Return(&gocelery.AsyncResult{}, nil)

	repo := testRepo()
	anchorRepo := &testinganchors.MockAnchorRepo{}
	anchorRepo.On("GetAnchorData", mock.Anything).Return(nil, errors.New("missing"))
	docSrv := documents.DefaultService(cfg, repo, anchorRepo, documents.NewServiceRegistry(), &idService)
	return idService, DefaultService(
		docSrv,
		repo,
		queueSrv,
		ctx[jobs.BootstrappedService].(jobs.Manager),
		func() documents.TokenRegistry { return nil }, anchorRepo)
}

func TestService_Update(t *testing.T) {
	_, srv := getServiceWithMockedLayers()
	invSrv := srv.(service)
	ctxh := testingconfig.CreateAccountContext(t, cfg)

	// missing last version
	model, _ := createCDWithEmbeddedInvoice(t)
	_, _, _, err := invSrv.Update(ctxh, model)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNotFound, err))
	assert.NoError(t, testRepo().Create(accountID, model.CurrentVersion(), model))

	// calculate data root fails
	nm := new(mockModel)
	nm.On("ID").Return(model.ID(), nil).Once()
	_, _, _, err = invSrv.Update(ctxh, nm)
	nm.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown document type")

	// success
	data, err := invSrv.DeriveInvoiceData(model)
	assert.Nil(t, err)
	data.GrossAmount = "100"
	collab := testingidentity.GenerateRandomDID().String()
	newInv, err := invSrv.DeriveFromUpdatePayload(ctxh, &clientinvoicepb.InvoiceUpdatePayload{
		DocumentId:  hexutil.Encode(model.ID()),
		WriteAccess: []string{collab},
		Data:        data,
	})
	assert.Nil(t, err)
	newData, err := invSrv.DeriveInvoiceData(newInv)
	assert.Nil(t, err)
	assert.Equal(t, data, newData)

	model, _, _, err = invSrv.Update(ctxh, newInv)
	assert.Nil(t, err)
	assert.NotNil(t, model)
	assert.True(t, testRepo().Exists(accountID, model.ID()))
	assert.True(t, testRepo().Exists(accountID, model.CurrentVersion()))
	assert.True(t, testRepo().Exists(accountID, model.PreviousVersion()))

	newData, err = invSrv.DeriveInvoiceData(model)
	assert.Nil(t, err)
	assert.Equal(t, data, newData)
}

func TestService_DeriveFromUpdatePayload(t *testing.T) {
	_, invSrv := getServiceWithMockedLayers()
	// nil payload
	doc, err := invSrv.DeriveFromUpdatePayload(nil, nil)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))
	assert.Nil(t, doc)

	// nil payload data
	doc, err = invSrv.DeriveFromUpdatePayload(nil, &clientinvoicepb.InvoiceUpdatePayload{})
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))
	assert.Nil(t, doc)

	// messed up identifier
	contextHeader := testingconfig.CreateAccountContext(t, cfg)
	payload := &clientinvoicepb.InvoiceUpdatePayload{DocumentId: "some identifier", Data: &clientinvoicepb.InvoiceData{}}
	doc, err = invSrv.DeriveFromUpdatePayload(contextHeader, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentIdentifier, err))
	assert.Contains(t, err.Error(), "failed to decode identifier")
	assert.Nil(t, doc)

	// missing last version
	id := utils.RandomSlice(32)
	payload.DocumentId = hexutil.Encode(id)
	doc, err = invSrv.DeriveFromUpdatePayload(contextHeader, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNotFound, err))
	assert.Nil(t, doc)

	// failed to load from data
	old, _ := createCDWithEmbeddedInvoice(t)
	err = testRepo().Create(accountID, old.CurrentVersion(), old)
	assert.Nil(t, err)
	payload.Data = &clientinvoicepb.InvoiceData{
		Sender:      "0xed03Fa80291fF5DDC284DE6b51E716B130b05e20",
		Recipient:   "0xEA939D5C0494b072c51565b191eE59B5D34fbf79",
		Payee:       "some data",
		GrossAmount: "42",
		Currency:    "EUR",
	}

	payload.DocumentId = hexutil.Encode(old.ID())
	doc, err = invSrv.DeriveFromUpdatePayload(contextHeader, payload)
	assert.Error(t, err)
	assert.Nil(t, doc)

	// failed core document new version
	payload.Data.Payee = "0x087D8ca6A16E6ce8d9fF55672E551A2828Ab8e8C"
	payload.WriteAccess = []string{"some wrong ID"}
	doc, err = invSrv.DeriveFromUpdatePayload(contextHeader, payload)
	assert.Error(t, err)
	assert.Nil(t, doc)

	// success
	wantCollab := testingidentity.GenerateRandomDID()
	payload.WriteAccess = []string{wantCollab.String()}
	doc, err = invSrv.DeriveFromUpdatePayload(contextHeader, payload)
	assert.Nil(t, err)
	assert.NotNil(t, doc)
	cs, err := doc.GetCollaborators()
	assert.NoError(t, err)
	assert.Len(t, cs.ReadWriteCollaborators, 3)
	assert.Contains(t, cs.ReadWriteCollaborators, wantCollab)
	assert.Equal(t, old.ID(), doc.ID())
	assert.Equal(t, payload.DocumentId, hexutil.Encode(doc.ID()))
	assert.Equal(t, old.CurrentVersion(), doc.PreviousVersion())
	assert.Equal(t, old.NextVersion(), doc.CurrentVersion())
	assert.NotNil(t, doc.NextVersion())
	data, err := doc.(*Invoice).getClientData()
	assert.NoError(t, err)
	assert.Equal(t, payload.Data, data)
}

func TestService_DeriveFromCreatePayload(t *testing.T) {
	invSrv := service{}
	ctxh := testingconfig.CreateAccountContext(t, cfg)

	// nil payload
	m, err := invSrv.DeriveFromCreatePayload(ctxh, nil)
	assert.Nil(t, m)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))

	// nil data payload
	m, err = invSrv.DeriveFromCreatePayload(ctxh, &clientinvoicepb.InvoiceCreatePayload{})
	assert.Nil(t, m)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))

	// Init fails
	payload := &clientinvoicepb.InvoiceCreatePayload{
		Data: &clientinvoicepb.InvoiceData{
			Payee: "some payee",
		},
	}

	m, err = invSrv.DeriveFromCreatePayload(ctxh, payload)
	assert.Nil(t, m)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalid, err))

	// success
	payload.Data.Payee = testingidentity.GenerateRandomDID().String()
	m, err = invSrv.DeriveFromCreatePayload(ctxh, payload)
	assert.Nil(t, err)
	assert.NotNil(t, m)
}

func TestService_DeriveFromCoreDocument(t *testing.T) {
	invSrv := service{repo: testRepo()}
	_, cd := createCDWithEmbeddedInvoice(t)
	m, err := invSrv.DeriveFromCoreDocument(cd)
	assert.Nil(t, err, "must return model")
	assert.NotNil(t, m, "model must be non-nil")
	inv, ok := m.(*Invoice)
	assert.True(t, ok, "must be true")
	assert.Equal(t, inv.Data.Recipient.String(), "0xEA939D5C0494b072c51565b191eE59B5D34fbf79")
	assert.Equal(t, inv.Data.GrossAmount.String(), "42")
}

func TestService_Create(t *testing.T) {
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	_, srv := getServiceWithMockedLayers()
	invSrv := srv.(service)

	// calculate data root fails
	m, _, _, err := invSrv.Create(ctxh, &mockModel{})
	assert.Nil(t, m)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown document type")

	// success
	inv, err := invSrv.DeriveFromCreatePayload(ctxh, testingdocuments.CreateInvoicePayload())
	assert.Nil(t, err)
	m, _, _, err = invSrv.Create(ctxh, inv)
	assert.Nil(t, err)
	assert.True(t, testRepo().Exists(accountID, m.ID()))
	assert.True(t, testRepo().Exists(accountID, m.CurrentVersion()))
}

func TestService_DeriveInvoiceData(t *testing.T) {
	_, invSrv := getServiceWithMockedLayers()

	// some random model
	_, err := invSrv.DeriveInvoiceData(&mockModel{})
	assert.Error(t, err, "Derive must fail")

	// success
	payload := testingdocuments.CreateInvoicePayload()
	inv, err := invSrv.DeriveFromCreatePayload(testingconfig.CreateAccountContext(t, cfg), payload)
	assert.Nil(t, err, "must be non nil")
	data, err := invSrv.DeriveInvoiceData(inv)
	assert.Nil(t, err, "Derive must succeed")
	assert.NotNil(t, data, "data must be non nil")
}

func TestService_DeriveInvoiceResponse(t *testing.T) {
	// success
	invSrv := service{repo: testRepo(), tokenRegFinder: func() documents.TokenRegistry {
		return nil
	}}

	// derive data failed
	m := new(mockModel)
	r, err := invSrv.DeriveInvoiceResponse(m)
	m.AssertExpectations(t)
	assert.Nil(t, r)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalidType, err))

	// success
	inv, _ := createCDWithEmbeddedInvoice(t)
	r, err = invSrv.DeriveInvoiceResponse(inv)
	payload := testingdocuments.CreateInvoicePayload()
	assert.Nil(t, err)
	assert.Equal(t, payload.Data, r.Data)
	assert.Contains(t, r.Header.WriteAccess, did.String())
}

func TestService_GetCurrentVersion(t *testing.T) {
	_, invSrv := getServiceWithMockedLayers()
	doc, _ := createCDWithEmbeddedInvoice(t)
	ctxh := testingconfig.CreateAccountContext(t, cfg)

	err := testRepo().Create(accountID, doc.CurrentVersion(), doc)
	assert.Nil(t, err)

	data, err := doc.(*Invoice).getClientData()
	assert.NoError(t, err)
	data.Currency = "INR"
	doc2 := new(Invoice)
	assert.NoError(t, doc2.PrepareNewVersion(doc, data, documents.CollaboratorsAccess{}, doc.(*Invoice).Attributes))
	assert.NoError(t, testRepo().Create(accountID, doc2.CurrentVersion(), doc2))

	doc3, err := invSrv.GetCurrentVersion(ctxh, doc.ID())
	assert.Nil(t, err)
	assert.Equal(t, doc2, doc3)
}

func TestService_GetVersion(t *testing.T) {
	_, invSrv := getServiceWithMockedLayers()
	inv, _ := createCDWithEmbeddedInvoice(t)
	err := testRepo().Create(accountID, inv.CurrentVersion(), inv)
	assert.Nil(t, err)

	ctxh := testingconfig.CreateAccountContext(t, cfg)
	mod, err := invSrv.GetVersion(ctxh, inv.ID(), inv.CurrentVersion())
	assert.Nil(t, err)

	mod, err = invSrv.GetVersion(ctxh, mod.ID(), []byte{})
	assert.Error(t, err)
}

func TestService_Exists(t *testing.T) {
	_, invSrv := getServiceWithMockedLayers()
	inv, _ := createCDWithEmbeddedInvoice(t)
	err := testRepo().Create(accountID, inv.CurrentVersion(), inv)
	assert.Nil(t, err)

	ctxh := testingconfig.CreateAccountContext(t, cfg)
	exists := invSrv.Exists(ctxh, inv.CurrentVersion())
	assert.True(t, exists, "invoice should exist")

	exists = invSrv.Exists(ctxh, utils.RandomSlice(32))
	assert.False(t, exists, " invoice should not exist")
}

func TestService_calculateDataRoot(t *testing.T) {
	invSrv := service{repo: testRepo()}
	ctxh := testingconfig.CreateAccountContext(t, cfg)

	// type mismatch
	inv, err := invSrv.validateAndPersist(ctxh, nil, &mockModel{}, nil)
	assert.Nil(t, inv)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown document type")

	// failed validator
	inv, err = invSrv.DeriveFromCreatePayload(ctxh, testingdocuments.CreateInvoicePayload())
	assert.Nil(t, err)
	v := documents.ValidatorFunc(func(_, _ documents.Model) error {
		return errors.New("validations fail")
	})
	inv, err = invSrv.validateAndPersist(ctxh, nil, inv, v)
	assert.Nil(t, inv)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validations fail")

	// create failed
	inv, err = invSrv.DeriveFromCreatePayload(ctxh, testingdocuments.CreateInvoicePayload())
	assert.Nil(t, err)
	err = invSrv.repo.Create(accountID, inv.CurrentVersion(), inv)
	assert.Nil(t, err)
	inv, err = invSrv.validateAndPersist(ctxh, nil, inv, CreateValidator())
	assert.Nil(t, inv)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), storage.ErrRepositoryModelCreateKeyExists)

	// success
	inv, err = invSrv.DeriveFromCreatePayload(ctxh, testingdocuments.CreateInvoicePayload())
	assert.Nil(t, err)
	inv, err = invSrv.validateAndPersist(ctxh, nil, inv, CreateValidator())
	assert.Nil(t, err)
	assert.NotNil(t, inv)
}

var testRepoGlobal documents.Repository

func testRepo() documents.Repository {
	if testRepoGlobal != nil {
		return testRepoGlobal
	}

	ldb, err := leveldb.NewLevelDBStorage(leveldb.GetRandomTestStoragePath())
	if err != nil {
		panic(err)
	}
	testRepoGlobal = documents.NewDBRepository(leveldb.NewLevelDBRepository(ldb))
	testRepoGlobal.Register(&Invoice{})
	return testRepoGlobal
}

func createCDWithEmbeddedInvoice(t *testing.T) (documents.Model, coredocumentpb.CoreDocument) {
	i := new(Invoice)
	err := i.InitInvoiceInput(testingdocuments.CreateInvoicePayload(), did)
	assert.NoError(t, err)
	i.GetTestCoreDocWithReset()
	_, err = i.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = i.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = i.CalculateDocumentRoot()
	assert.NoError(t, err)
	cd, err := i.PackCoreDocument()
	assert.NoError(t, err)
	return i, cd
}

func TestService_CreateModel(t *testing.T) {
	payload := documents.CreatePayload{}
	srv := service{}

	// nil  model
	_, _, err := srv.CreateModel(context.Background(), payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))

	// empty context
	payload.Data = utils.RandomSlice(32)
	_, _, err = srv.CreateModel(context.Background(), payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentConfigAccountID, err))

	// invalid data
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	_, _, err = srv.CreateModel(ctxh, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalid, err))

	// validator failed
	payload.Data = validData(t)
	_, _, err = srv.CreateModel(ctxh, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalid, err))

	// success
	payload.Data = validDataWithCurrency(t)
	srv.repo = testRepo()
	jm := testingjobs.MockJobManager{}
	jm.On("ExecuteWithinJob", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(jobs.NilJobID(), make(chan bool), nil)
	srv.jobManager = jm
	m, _, err := srv.CreateModel(ctxh, payload)
	assert.NoError(t, err)
	assert.NotNil(t, m)
	jm.AssertExpectations(t)
}

func TestService_UpdateModel(t *testing.T) {
	payload := documents.UpdatePayload{}
	srv := service{}

	// nil  model
	_, _, err := srv.UpdateModel(context.Background(), payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))

	// empty context
	payload.Data = utils.RandomSlice(32)
	_, _, err = srv.UpdateModel(context.Background(), payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentConfigAccountID, err))

	// missing id
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	_, srvr := getServiceWithMockedLayers()
	srv = srvr.(service)
	payload.DocumentID = utils.RandomSlice(32)
	_, _, err = srv.UpdateModel(ctxh, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNotFound, err))

	// payload invalid
	inv := createInvoice(t)
	err = testRepo().Create(did[:], inv.ID(), inv)
	assert.NoError(t, err)
	payload.DocumentID = inv.ID()
	_, _, err = srv.UpdateModel(ctxh, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalid, err))

	// validator failed
	payload.Data = validData(t)
	_, _, err = srv.UpdateModel(ctxh, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalid, err))

	// Success
	payload.Data = validDataWithCurrency(t)
	jm := testingjobs.MockJobManager{}
	jm.On("ExecuteWithinJob", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(jobs.NilJobID(), make(chan bool), nil)
	srv.jobManager = jm
	m, _, err := srv.UpdateModel(ctxh, payload)
	assert.NoError(t, err)
	assert.Equal(t, m.ID(), inv.ID())
	assert.Equal(t, m.CurrentVersion(), inv.NextVersion())
	jm.AssertExpectations(t)
}
