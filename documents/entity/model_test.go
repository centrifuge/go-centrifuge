// +build unit

package entity

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/centrifuge/centrifuge-protobufs/documenttypes"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/entity"
	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/bootstrap/bootstrappers/testlogging"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/config/configstore"
	"github.com/centrifuge/go-centrifuge/contextutil"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/ethereum"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/identity/ideth"
	"github.com/centrifuge/go-centrifuge/jobs"
	"github.com/centrifuge/go-centrifuge/p2p"
	cliententitypb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/entity"
	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/centrifuge/go-centrifuge/storage/leveldb"
	"github.com/centrifuge/go-centrifuge/testingutils/config"
	"github.com/centrifuge/go-centrifuge/testingutils/documents"
	"github.com/centrifuge/go-centrifuge/testingutils/identity"
	"github.com/centrifuge/go-centrifuge/testingutils/testingjobs"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var ctx = map[string]interface{}{}
var cfg config.Configuration
var configService config.Service

var (
	did       = testingidentity.GenerateRandomDID()
	dIDBytes  = did[:]
	accountID = did[:]
)

type mockAnchorRepo struct {
	mock.Mock
	anchors.AnchorRepository
}

func (m *mockAnchorRepo) GetDocumentRootOf(anchorID anchors.AnchorID) (anchors.DocumentRoot, error) {
	args := m.Called(anchorID)
	docRoot, _ := args.Get(0).(anchors.DocumentRoot)
	return docRoot, args.Error(1)
}

func (m *mockAnchorRepo) GetAnchorData(anchorID anchors.AnchorID) (docRoot anchors.DocumentRoot, anchoredTime time.Time, err error) {
	args := m.Called(anchorID)
	docRoot, _ = args.Get(0).(anchors.DocumentRoot)
	anchoredTime, _ = args.Get(1).(time.Time)
	return docRoot, anchoredTime, args.Error(2)
}

func TestMain(m *testing.M) {
	ethClient := &ethereum.MockEthClient{}
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

func TestEntity_PackCoreDocument(t *testing.T) {
	ctx := testingconfig.CreateAccountContext(t, cfg)
	did, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)

	entity := new(Entity)
	assert.NoError(t, entity.InitEntityInput(testingdocuments.CreateEntityPayload(), did))

	cd, err := entity.PackCoreDocument()
	assert.NoError(t, err)
	assert.NotNil(t, cd.EmbeddedData)
}

func TestEntity_JSON(t *testing.T) {
	entity := new(Entity)
	ctx := testingconfig.CreateAccountContext(t, cfg)
	did, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)
	assert.NoError(t, entity.InitEntityInput(testingdocuments.CreateEntityPayload(), did))

	cd, err := entity.PackCoreDocument()
	assert.NoError(t, err)
	jsonBytes, err := entity.JSON()
	assert.Nil(t, err, "marshal to json didn't work correctly")
	assert.True(t, json.Valid(jsonBytes), "json format not correct")

	entity = new(Entity)
	err = entity.FromJSON(jsonBytes)
	assert.Nil(t, err, "unmarshal JSON didn't work correctly")

	ncd, err := entity.PackCoreDocument()
	assert.Nil(t, err, "JSON unmarshal damaged entity variables")
	assert.Equal(t, cd, ncd)
}

func TestEntityModel_UnpackCoreDocument(t *testing.T) {
	var model = new(Entity)
	var err error

	// embed data missing
	err = model.UnpackCoreDocument(coredocumentpb.CoreDocument{})
	assert.Error(t, err)

	// embed data type is wrong
	err = model.UnpackCoreDocument(coredocumentpb.CoreDocument{EmbeddedData: new(any.Any)})
	assert.Error(t, err, "unpack must fail due to missing embed data")

	// embed data is wrong
	err = model.UnpackCoreDocument(coredocumentpb.CoreDocument{
		EmbeddedData: &any.Any{
			Value:   utils.RandomSlice(32),
			TypeUrl: documenttypes.EntityDataTypeUrl,
		},
	})
	assert.Error(t, err)

	// successful
	entity, cd := createCDWithEmbeddedEntity(t)
	err = model.UnpackCoreDocument(cd)
	assert.NoError(t, err)

	d := model.getClientData()
	d1 := entity.(*Entity).getClientData()
	assert.Equal(t, d.Addresses[0], d1.Addresses[0])
	assert.Equal(t, model.ID(), entity.ID())
	assert.Equal(t, model.CurrentVersion(), entity.CurrentVersion())
	assert.Equal(t, model.PreviousVersion(), entity.PreviousVersion())
}

func TestEntityModel_getClientData(t *testing.T) {
	entityData := testingdocuments.CreateEntityData()
	entity := new(Entity)
	entity.CoreDocument = new(documents.CoreDocument)
	err := entity.loadFromP2PProtobuf(&entityData)
	assert.NoError(t, err)

	data := entity.getClientData()
	assert.NotNil(t, data, "entity data should not be nil")
	assert.Equal(t, data.Addresses, entityData.Addresses, "addresses should match")
	assert.Equal(t, data.Contacts, entityData.Contacts, "contacts should match")
	assert.Equal(t, data.LegalName, entityData.LegalName, "legal name should match")
}

func TestEntityModel_InitEntityInput(t *testing.T) {
	ctx := testingconfig.CreateAccountContext(t, cfg)
	did, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)

	// fail recipient
	data := &cliententitypb.EntityData{
		Identity:  testingidentity.GenerateRandomDID().ToAddress().String(),
		LegalName: "Company Test",
		Contacts:  []*entitypb.Contact{{Name: "Satoshi Nakamoto"}},
		Addresses: []*entitypb.Address{{IsMain: true,
			AddressLine1: "Sample Street 1",
			Zip:          "12345",
			State:        "Germany",
		}, {IsMain: false, State: "US"}},
	}
	e := new(Entity)
	err = e.InitEntityInput(&cliententitypb.EntityCreatePayload{Data: data}, did)
	assert.Nil(t, err, "should be successful")

	e = new(Entity)
	collabs := []string{"0x010102040506", "some id"}
	err = e.InitEntityInput(&cliententitypb.EntityCreatePayload{Data: data, WriteAccess: collabs}, did)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode collaborator")

	collab1, err := identity.NewDIDFromString("0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7")
	assert.NoError(t, err)
	collab2, err := identity.NewDIDFromString("0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF3")
	assert.NoError(t, err)
	collabs = []string{collab1.String(), collab2.String()}
	err = e.InitEntityInput(&cliententitypb.EntityCreatePayload{Data: data, WriteAccess: collabs}, did)
	assert.Nil(t, err, "must be nil")

}

func TestEntityModel_calculateDataRoot(t *testing.T) {
	ctx := testingconfig.CreateAccountContext(t, cfg)
	did, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)
	m := new(Entity)
	err = m.InitEntityInput(testingdocuments.CreateEntityPayload(), did)
	assert.Nil(t, err, "Init must pass")
	m.GetTestCoreDocWithReset()

	dr, err := m.CalculateDataRoot()
	assert.Nil(t, err, "calculate must pass")
	assert.False(t, utils.IsEmptyByteSlice(dr))
}

func TestEntity_CreateProofs(t *testing.T) {
	e := createEntity(t)
	rk := e.Document.Roles[0].RoleKey
	pf := fmt.Sprintf(documents.CDTreePrefix+".roles[%s].collaborators[0]", hexutil.Encode(rk))
	proof, err := e.CreateProofs([]string{"entity.legal_name", pf, documents.CDTreePrefix + ".document_type"})
	assert.NoError(t, err)
	assert.NotNil(t, proof)
	tree, err := e.DocumentRootTree()
	assert.NoError(t, err)
	h := sha256.New()
	dataLeaves, err := e.getDataLeaves()
	assert.NoError(t, err)

	// Validate entity_number
	err = e.CoreDocument.ValidateDataProof("entity.legal_name", e.DocumentType(), tree.RootHash(), proof[0], dataLeaves, h)
	assert.Nil(t, err)

	// Validate roles
	err = e.CoreDocument.ValidateDataProof(pf, e.DocumentType(), tree.RootHash(), proof[1], dataLeaves, h)
	assert.Nil(t, err)

	// Validate []byte value
	acc, err := identity.NewDIDFromBytes(proof[1].Value)
	assert.NoError(t, err)
	assert.True(t, e.AccountCanRead(acc))

	// Validate document_type
	err = e.CoreDocument.ValidateDataProof(documents.CDTreePrefix+".document_type", e.DocumentType(), tree.RootHash(), proof[2], dataLeaves, h)
	assert.Nil(t, err)
}

func createEntity(t *testing.T) *Entity {
	e := new(Entity)
	err := e.InitEntityInput(testingdocuments.CreateEntityPayload(), did)
	assert.NoError(t, err)
	e.GetTestCoreDocWithReset()
	_, err = e.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = e.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = e.CalculateDocumentRoot()
	assert.NoError(t, err)
	return e
}

func TestEntityModel_createProofsFieldDoesNotExist(t *testing.T) {
	e := createEntity(t)
	_, err := e.CreateProofs([]string{"nonexisting"})
	assert.NotNil(t, err)
}

func TestEntityModel_GetDocumentID(t *testing.T) {
	e := createEntity(t)
	assert.Equal(t, e.CoreDocument.ID(), e.ID())
}

func TestEntityModel_getDataTree(t *testing.T) {
	e := createEntity(t)
	tree, err := e.getDataTree()
	assert.Nil(t, err, "tree should be generated without error")
	_, leaf := tree.GetLeafByProperty("entity.legal_name")
	assert.NotNil(t, leaf)
	assert.Equal(t, "entity.legal_name", leaf.Property.ReadableName())
}

func TestEntity_CollaboratorCanUpdate(t *testing.T) {
	entity := createEntity(t)
	id1 := did
	id2 := testingidentity.GenerateRandomDID()
	id3 := testingidentity.GenerateRandomDID()

	// wrong type
	err := entity.CollaboratorCanUpdate(new(mockModel), id1)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalidType, err))
	assert.NoError(t, testRepo().Create(id1[:], entity.CurrentVersion(), entity))

	// update the document
	model, err := testRepo().Get(id1[:], entity.CurrentVersion())
	assert.NoError(t, err)
	oldEntity := model.(*Entity)
	data := oldEntity.getClientData()
	data.LegalName = "new legal name"
	err = entity.PrepareNewVersion(entity, data, documents.CollaboratorsAccess{ReadWriteCollaborators: []identity.DID{id3}}, oldEntity.Attributes)
	assert.NoError(t, err)

	// id1 should have permission
	assert.NoError(t, oldEntity.CollaboratorCanUpdate(entity, id1))

	// id2 should fail since it doesn't have the permission to update
	assert.Error(t, oldEntity.CollaboratorCanUpdate(entity, id2))

	// update the id3 rules to update only legal fields
	entity.CoreDocument.Document.TransitionRules[3].MatchType = coredocumentpb.FieldMatchType_FIELD_MATCH_TYPE_EXACT
	entity.CoreDocument.Document.TransitionRules[3].Field = append(compactPrefix(), 0, 0, 0, 2)
	assert.NoError(t, testRepo().Create(id1[:], entity.CurrentVersion(), entity))

	// fetch the document
	model, err = testRepo().Get(id1[:], entity.CurrentVersion())
	assert.NoError(t, err)
	oldEntity = model.(*Entity)
	data = oldEntity.getClientData()
	data.LegalName = "second new legal name"
	data.Contacts = nil
	err = entity.PrepareNewVersion(entity, data, documents.CollaboratorsAccess{}, oldEntity.Attributes)
	assert.NoError(t, err)

	// id1 should have permission
	assert.NoError(t, oldEntity.CollaboratorCanUpdate(entity, id1))

	// id2 should fail since it doesn't have the permission to update
	assert.Error(t, oldEntity.CollaboratorCanUpdate(entity, id2))

	// id3 should pass with just one error since changing contacts is not allowed
	err = oldEntity.CollaboratorCanUpdate(entity, id3)
	assert.Error(t, err)
	assert.Equal(t, 5, errors.Len(err)) //five contact fields have been changed
	assert.Contains(t, err.Error(), "entity.contacts")

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
	testRepoGlobal.Register(&Entity{})
	return testRepoGlobal
}

func createCDWithEmbeddedEntity(t *testing.T) (documents.Model, coredocumentpb.CoreDocument) {
	e := new(Entity)
	err := e.InitEntityInput(testingdocuments.CreateEntityPayload(), did)
	assert.NoError(t, err)
	_, err = e.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = e.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = e.CalculateDocumentRoot()
	assert.NoError(t, err)
	cd, err := e.PackCoreDocument()
	assert.NoError(t, err)
	return e, cd
}

func TestEntity_AddAttributes(t *testing.T) {
	e, _ := createCDWithEmbeddedEntity(t)
	label := "some key"
	value := "some value"
	attr, err := documents.NewAttribute(label, documents.AttrString, value)
	assert.NoError(t, err)

	// success
	err = e.AddAttributes(documents.CollaboratorsAccess{}, true, attr)
	assert.NoError(t, err)
	assert.True(t, e.AttributeExists(attr.Key))
	gattr, err := e.GetAttribute(attr.Key)
	assert.NoError(t, err)
	assert.Equal(t, attr, gattr)

	// fail
	attr.Value.Type = documents.AttributeType("some attr")
	err = e.AddAttributes(documents.CollaboratorsAccess{}, true, attr)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrCDAttribute, err))
}

func TestEntity_DeleteAttribute(t *testing.T) {
	e, _ := createCDWithEmbeddedEntity(t)
	label := "some key"
	value := "some value"
	attr, err := documents.NewAttribute(label, documents.AttrString, value)
	assert.NoError(t, err)

	// failed
	err = e.DeleteAttribute(attr.Key, true)
	assert.Error(t, err)

	// success
	assert.NoError(t, e.AddAttributes(documents.CollaboratorsAccess{}, true, attr))
	assert.True(t, e.AttributeExists(attr.Key))
	assert.NoError(t, e.DeleteAttribute(attr.Key, true))
	assert.False(t, e.AttributeExists(attr.Key))
}

func TestEntity_GetData(t *testing.T) {
	e := createEntity(t)
	data := e.GetData()
	assert.Equal(t, e.Data, data)
}

func marshallData(t *testing.T, m map[string]interface{}) []byte {
	data, err := json.Marshal(m)
	assert.NoError(t, err)
	return data
}

func emptyDIDData(t *testing.T) []byte {
	d := map[string]interface{}{
		"identity": "",
	}

	return marshallData(t, d)
}

func invalidDIDData(t *testing.T) []byte {
	d := map[string]interface{}{
		"identity": "1acdew123asdefres",
	}

	return marshallData(t, d)
}

func emptyPaymentDetail(t *testing.T) []byte {
	d := map[string]interface{}{
		"identity": "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
		"payment_details": []map[string]interface{}{
			{},
			{"predefined": true},
		},
	}

	return marshallData(t, d)
}

func multiPaymentDetail(t *testing.T) []byte {
	d := map[string]interface{}{
		"identity": "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
		"payment_details": []map[string]interface{}{
			{
				"predefined": true,
				"bank_payment_method": map[string]interface{}{
					"identifier": "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
				},
				"crypto_payment_method": map[string]interface{}{
					"identifier": "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
				},
			},
		},
	}

	return marshallData(t, d)
}

func validData(t *testing.T) []byte {
	d := map[string]interface{}{
		"legal_name": "Hello, World!",
		"payment_details": []map[string]interface{}{
			{
				"predefined": true,
				"bank_payment_method": map[string]interface{}{
					"identifier": "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
				},
			},
		},
	}

	return marshallData(t, d)
}

func validDataWithIdentity(t *testing.T) []byte {
	d := map[string]interface{}{
		"legal_name": "Hello, World!",
		"identity":   "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
		"payment_details": []map[string]interface{}{
			{
				"predefined": true,
				"bank_payment_method": map[string]interface{}{
					"identifier": "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
				},
			},
		},
	}

	return marshallData(t, d)
}

func checkEntityPayloadDataError(t *testing.T, e *Entity, payload documents.CreatePayload) {
	err := e.loadData(payload.Data)
	assert.Error(t, err)
}

func TestEntity_loadData(t *testing.T) {
	e := new(Entity)
	payload := documents.CreatePayload{}

	// empty did data
	payload.Data = emptyDIDData(t)
	checkEntityPayloadDataError(t, e, payload)

	// invalid did data
	payload.Data = invalidDIDData(t)
	checkEntityPayloadDataError(t, e, payload)

	// empty payment detail
	payload.Data = emptyPaymentDetail(t)
	checkEntityPayloadDataError(t, e, payload)

	// multiple payment detail
	payload.Data = multiPaymentDetail(t)
	checkEntityPayloadDataError(t, e, payload)

	// valid data
	payload.Data = validData(t)
	err := e.loadData(payload.Data)
	assert.NoError(t, err)
	data := e.GetData().(Data)
	assert.Equal(t, data.LegalName, "Hello, World!")
	assert.Len(t, data.PaymentDetails, 1)
	assert.NotNil(t, data.PaymentDetails[0].BankPaymentMethod)
	assert.Nil(t, data.PaymentDetails[0].CryptoPaymentMethod)
	assert.Nil(t, data.PaymentDetails[0].OtherPaymentMethod)
	assert.True(t, data.PaymentDetails[0].Predefined)
	assert.Equal(t, data.PaymentDetails[0].BankPaymentMethod.Identifier.String(), "0xbaeb33a61f05e6f269f1c4b4cff91a901b54daf7")
}

func TestEntity_unpackFromCreatePayload(t *testing.T) {
	payload := documents.CreatePayload{}
	e := new(Entity)

	// invalid data
	payload.Data = invalidDIDData(t)
	err := e.unpackFromCreatePayload(did, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(ErrEntityInvalidData, err))

	// invalid attributes
	attr, err := documents.NewAttribute("test", documents.AttrString, "value")
	assert.NoError(t, err)
	val := attr.Value
	val.Type = documents.AttributeType("some type")
	attr.Value = val
	payload.Attributes = map[documents.AttrKey]documents.Attribute{
		attr.Key: attr,
	}
	payload.Data = validData(t)
	err = e.unpackFromCreatePayload(did, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrCDCreate, err))

	// valid
	val.Type = documents.AttrString
	attr.Value = val
	payload.Attributes = map[documents.AttrKey]documents.Attribute{
		attr.Key: attr,
	}
	err = e.unpackFromCreatePayload(did, payload)
	assert.NoError(t, err)
}

func TestInvoice_unpackFromUpdatePayload(t *testing.T) {
	payload := documents.UpdatePayload{}
	old := createEntity(t)
	e := new(Entity)

	// invalid data
	payload.Data = invalidDIDData(t)
	err := e.unpackFromUpdatePayload(old, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(ErrEntityInvalidData, err))

	// invalid attributes
	attr, err := documents.NewAttribute("test", documents.AttrString, "value")
	assert.NoError(t, err)
	val := attr.Value
	val.Type = documents.AttributeType("some type")
	attr.Value = val
	payload.Attributes = map[documents.AttrKey]documents.Attribute{
		attr.Key: attr,
	}
	payload.Data = validData(t)
	err = e.unpackFromUpdatePayload(old, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrCDNewVersion, err))

	// valid
	val.Type = documents.AttrString
	attr.Value = val
	payload.Attributes = map[documents.AttrKey]documents.Attribute{
		attr.Key: attr,
	}
	err = e.unpackFromUpdatePayload(old, payload)
	assert.NoError(t, err)
}
