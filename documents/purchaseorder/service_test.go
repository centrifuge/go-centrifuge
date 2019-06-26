// +build unit

package purchaseorder

import (
	"context"
	"testing"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/jobs"
	clientpurchaseorderpb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/purchaseorder"
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
	accountID = did[:]
)

func getServiceWithMockedLayers() (*testingcommons.MockIdentityService, Service) {
	idService := &testingcommons.MockIdentityService{}
	idService.On("IsSignedWithPurpose", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(true, nil).Once()
	queueSrv := new(testingutils.MockQueue)
	queueSrv.On("EnqueueJob", mock.Anything, mock.Anything).Return(&gocelery.AsyncResult{}, nil)
	jobManager := ctx[jobs.BootstrappedService].(jobs.Manager)
	repo := testRepo()
	anchorRepo := &testinganchors.MockAnchorRepo{}
	anchorRepo.On("GetAnchorData", mock.Anything).Return(nil, errors.New("missing"))
	docSrv := documents.DefaultService(cfg, repo, anchorRepo, documents.NewServiceRegistry(), idService)
	return idService, DefaultService(docSrv, repo, queueSrv, jobManager, func() documents.TokenRegistry {
		return nil
	}, anchorRepo)
}

func TestService_Update(t *testing.T) {
	_, poSrv := getServiceWithMockedLayers()
	ctxh := testingconfig.CreateAccountContext(t, cfg)

	// missing last version
	po, _ := createCDWithEmbeddedPO(t)
	_, _, _, err := poSrv.Update(ctxh, po)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNotFound, err))

	assert.NoError(t, testRepo().Create(accountID, po.CurrentVersion(), po))
	// success
	data, err := poSrv.DerivePurchaseOrderData(po)
	assert.Nil(t, err)
	data.TotalAmount = "100"
	collab := testingidentity.GenerateRandomDID().String()
	newPO, err := poSrv.DeriveFromUpdatePayload(ctxh, &clientpurchaseorderpb.PurchaseOrderUpdatePayload{
		DocumentId:  hexutil.Encode(po.ID()),
		WriteAccess: []string{collab},
		Data:        data,
	})
	assert.Nil(t, err)
	newData, err := poSrv.DerivePurchaseOrderData(newPO)
	assert.Nil(t, err)
	assert.Equal(t, data, newData)
	po, _, _, err = poSrv.Update(ctxh, newPO)
	assert.Nil(t, err)
	assert.NotNil(t, po)
	assert.True(t, testRepo().Exists(accountID, po.ID()))
	assert.True(t, testRepo().Exists(accountID, po.CurrentVersion()))
	assert.True(t, testRepo().Exists(accountID, po.PreviousVersion()))

	newData, err = poSrv.DerivePurchaseOrderData(po)
	assert.Nil(t, err)
	assert.Equal(t, data, newData)
}

func TestService_DeriveFromUpdatePayload(t *testing.T) {
	_, poSrv := getServiceWithMockedLayers()
	ctxh := testingconfig.CreateAccountContext(t, cfg)

	// nil payload
	doc, err := poSrv.DeriveFromUpdatePayload(ctxh, nil)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))
	assert.Nil(t, doc)

	// nil payload data
	doc, err = poSrv.DeriveFromUpdatePayload(ctxh, &clientpurchaseorderpb.PurchaseOrderUpdatePayload{})
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))
	assert.Nil(t, doc)

	// messed up identifier
	contextHeader := testingconfig.CreateAccountContext(t, cfg)
	payload := &clientpurchaseorderpb.PurchaseOrderUpdatePayload{DocumentId: "some identifier", Data: &clientpurchaseorderpb.PurchaseOrderData{}}
	doc, err = poSrv.DeriveFromUpdatePayload(contextHeader, payload)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode identifier")
	assert.Nil(t, doc)

	// missing last version
	id := utils.RandomSlice(32)
	payload.DocumentId = hexutil.Encode(id)
	doc, err = poSrv.DeriveFromUpdatePayload(contextHeader, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNotFound, err))
	assert.Nil(t, doc)

	// failed to load from data
	old, _ := createCDWithEmbeddedPO(t)
	err = testRepo().Create(accountID, old.CurrentVersion(), old)
	assert.Nil(t, err)
	payload.Data = &clientpurchaseorderpb.PurchaseOrderData{
		Recipient: "some recipient",
		Currency:  "EUR",
	}

	payload.DocumentId = hexutil.Encode(old.ID())
	doc, err = poSrv.DeriveFromUpdatePayload(contextHeader, payload)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load purchase order from data")
	assert.Nil(t, doc)

	// failed core document new version
	payload.Data.Recipient = "0xEA939D5C0494b072c51565b191eE59B5D34fbf79"
	payload.WriteAccess = []string{"some wrong ID"}
	doc, err = poSrv.DeriveFromUpdatePayload(contextHeader, payload)
	assert.Error(t, err)
	assert.Nil(t, doc)

	// success
	wantCollab := testingidentity.GenerateRandomDID()
	payload.WriteAccess = []string{wantCollab.String()}
	doc, err = poSrv.DeriveFromUpdatePayload(contextHeader, payload)
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
	data, err := doc.(*PurchaseOrder).getClientData()
	assert.NoError(t, err)
	assert.Equal(t, payload.Data, data)
}

func TestService_DeriveFromCreatePayload(t *testing.T) {
	poSrv := service{}
	ctxh := testingconfig.CreateAccountContext(t, cfg)

	// nil payload
	m, err := poSrv.DeriveFromCreatePayload(ctxh, nil)
	assert.Nil(t, m)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))

	// nil data payload
	m, err = poSrv.DeriveFromCreatePayload(ctxh, &clientpurchaseorderpb.PurchaseOrderCreatePayload{})
	assert.Nil(t, m)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))

	// Init fails
	payload := &clientpurchaseorderpb.PurchaseOrderCreatePayload{
		Data: &clientpurchaseorderpb.PurchaseOrderData{
			Recipient: "some recipient",
		},
	}

	m, err = poSrv.DeriveFromCreatePayload(ctxh, payload)
	assert.Nil(t, m)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalid, err))

	// success
	payload.Data.Recipient = "0xEA939D5C0494b072c51565b191eE59B5D34fbf79"
	m, err = poSrv.DeriveFromCreatePayload(ctxh, payload)
	assert.Nil(t, err)
	assert.NotNil(t, m)
}

func TestService_DeriveFromCoreDocument(t *testing.T) {
	poSrv := service{repo: testRepo()}
	_, cd := createCDWithEmbeddedPO(t)
	m, err := poSrv.DeriveFromCoreDocument(cd)
	assert.Nil(t, err, "must return model")
	assert.NotNil(t, m, "model must be non-nil")
	po, ok := m.(*PurchaseOrder)
	assert.True(t, ok, "must be true")
	assert.Equal(t, po.Data.Recipient.String(), "0xEA939D5C0494b072c51565b191eE59B5D34fbf79")
	assert.Equal(t, po.Data.TotalAmount.String(), "42")
}

func TestService_Create(t *testing.T) {
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	_, poSrv := getServiceWithMockedLayers()

	// calculate data root fails
	m, _, _, err := poSrv.Create(ctxh, &testingdocuments.MockModel{})
	assert.Nil(t, m)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown document type")

	// success
	po, err := poSrv.DeriveFromCreatePayload(ctxh, testingdocuments.CreatePOPayload())
	assert.Nil(t, err)
	m, _, _, err = poSrv.Create(ctxh, po)
	assert.Nil(t, err)
	assert.NotNil(t, m)

	assert.Nil(t, err)
	assert.True(t, testRepo().Exists(accountID, po.ID()))
	assert.True(t, testRepo().Exists(accountID, po.CurrentVersion()))
}

func TestService_DerivePurchaseOrderData(t *testing.T) {
	var m documents.Model
	_, poSrv := getServiceWithMockedLayers()

	// unknown type
	m = &testingdocuments.MockModel{}
	d, err := poSrv.DerivePurchaseOrderData(m)
	assert.Nil(t, d)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalidType, err))

	// success
	payload := testingdocuments.CreatePOPayload()
	m, err = poSrv.DeriveFromCreatePayload(testingconfig.CreateAccountContext(t, cfg), payload)
	assert.Nil(t, err)
	d, err = poSrv.DerivePurchaseOrderData(m)
	assert.Nil(t, err)
	assert.Equal(t, d.Currency, payload.Data.Currency)
}

func TestService_DerivePurchaseOrderResponse(t *testing.T) {
	poSrv := service{tokenRegFinder: func() documents.TokenRegistry {
		return nil
	}}

	// derive data failed
	m := &testingdocuments.MockModel{}
	r, err := poSrv.DerivePurchaseOrderResponse(m)
	m.AssertExpectations(t)
	assert.Nil(t, r)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalidType, err))

	// success
	payload := testingdocuments.CreatePOPayload()
	po, err := poSrv.DeriveFromCreatePayload(testingconfig.CreateAccountContext(t, cfg), payload)
	assert.Nil(t, err)
	r, err = poSrv.DerivePurchaseOrderResponse(po)
	assert.Nil(t, err)
	assert.Equal(t, payload.Data, r.Data)
	assert.Contains(t, r.Header.WriteAccess, did.String())
}

func TestService_GetCurrentVersion(t *testing.T) {
	_, poSrv := getServiceWithMockedLayers()
	doc, _ := createCDWithEmbeddedPO(t)
	ctxh := testingconfig.CreateAccountContext(t, cfg)

	err := testRepo().Create(accountID, doc.CurrentVersion(), doc)
	assert.Nil(t, err)

	data, err := doc.(*PurchaseOrder).getClientData()
	assert.NoError(t, err)
	data.Currency = "INR"
	doc2 := new(PurchaseOrder)
	assert.NoError(t, doc2.PrepareNewVersion(doc, data, documents.CollaboratorsAccess{}, doc.(*PurchaseOrder).Attributes))
	assert.NoError(t, testRepo().Create(accountID, doc2.CurrentVersion(), doc2))

	doc3, err := poSrv.GetCurrentVersion(ctxh, doc.ID())
	assert.Nil(t, err)
	assert.Equal(t, doc2, doc3)
}

func TestService_GetVersion(t *testing.T) {
	_, poSrv := getServiceWithMockedLayers()
	po, _ := createCDWithEmbeddedPO(t)
	err := testRepo().Create(accountID, po.CurrentVersion(), po)
	assert.Nil(t, err)

	ctxh := testingconfig.CreateAccountContext(t, cfg)
	mod, err := poSrv.GetVersion(ctxh, po.ID(), po.CurrentVersion())
	assert.Nil(t, err)

	mod, err = poSrv.GetVersion(ctxh, mod.ID(), []byte{})
	assert.Error(t, err)
}

func TestService_Exists(t *testing.T) {
	_, poSrv := getServiceWithMockedLayers()
	po, _ := createCDWithEmbeddedPO(t)
	err := testRepo().Create(accountID, po.CurrentVersion(), po)
	assert.Nil(t, err)

	ctxh := testingconfig.CreateAccountContext(t, cfg)
	exists := poSrv.Exists(ctxh, po.CurrentVersion())
	assert.True(t, exists, "purchase order should exist")

	exists = poSrv.Exists(ctxh, utils.RandomSlice(32))
	assert.False(t, exists, "purchase order should not exist")
}

func TestService_calculateDataRoot(t *testing.T) {
	poSrv := service{repo: testRepo()}
	ctxh := testingconfig.CreateAccountContext(t, cfg)

	// type mismatch
	po, err := poSrv.validateAndPersist(ctxh, nil, &testingdocuments.MockModel{}, nil)
	assert.Nil(t, po)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown document type")

	// failed validator
	po, err = poSrv.DeriveFromCreatePayload(ctxh, testingdocuments.CreatePOPayload())
	assert.Nil(t, err)
	v := documents.ValidatorFunc(func(_, _ documents.Model) error {
		return errors.New("validations fail")
	})
	po, err = poSrv.validateAndPersist(ctxh, nil, po, v)
	assert.Nil(t, po)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validations fail")

	// create failed
	po, err = poSrv.DeriveFromCreatePayload(ctxh, testingdocuments.CreatePOPayload())
	assert.Nil(t, err)
	err = poSrv.repo.Create(accountID, po.CurrentVersion(), po)
	assert.Nil(t, err)
	po, err = poSrv.validateAndPersist(ctxh, nil, po, CreateValidator())
	assert.Nil(t, po)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), storage.ErrRepositoryModelCreateKeyExists)

	// success
	po, err = poSrv.DeriveFromCreatePayload(ctxh, testingdocuments.CreatePOPayload())
	assert.Nil(t, err)
	po, err = poSrv.validateAndPersist(ctxh, nil, po, CreateValidator())
	assert.Nil(t, err)
	assert.NotNil(t, po)
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
	testRepoGlobal.Register(&PurchaseOrder{})
	return testRepoGlobal
}

func createCDWithEmbeddedPO(t *testing.T) (documents.Model, coredocumentpb.CoreDocument) {
	po := new(PurchaseOrder)
	err := po.InitPurchaseOrderInput(testingdocuments.CreatePOPayload(), did)
	assert.NoError(t, err)
	po.GetTestCoreDocWithReset()
	_, err = po.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = po.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = po.CalculateDocumentRoot()
	assert.NoError(t, err)
	cd, err := po.PackCoreDocument()
	assert.NoError(t, err)
	return po, cd
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
	po := createPurchaseOrder(t)
	err = testRepo().Create(did[:], po.ID(), po)
	assert.NoError(t, err)
	payload.DocumentID = po.ID()
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
	assert.Equal(t, m.ID(), po.ID())
	assert.Equal(t, m.CurrentVersion(), po.NextVersion())
	jm.AssertExpectations(t)
}
