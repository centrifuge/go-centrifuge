// +build unit

package entity

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

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
	"github.com/centrifuge/go-centrifuge/nft"
	"github.com/centrifuge/go-centrifuge/p2p"
	cliententitypb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/entity"
	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/centrifuge/go-centrifuge/storage/leveldb"
	"github.com/centrifuge/go-centrifuge/testingutils/commons"
	"github.com/centrifuge/go-centrifuge/testingutils/config"
	"github.com/centrifuge/go-centrifuge/testingutils/documents"
	"github.com/centrifuge/go-centrifuge/testingutils/identity"
	"github.com/centrifuge/go-centrifuge/testingutils/testingtx"
	"github.com/centrifuge/go-centrifuge/transactions"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var ctx = map[string]interface{}{}
var cfg config.Configuration
var configService config.Service
var defaultDID = testingidentity.GenerateRandomDID()

var (
	did       = testingidentity.GenerateRandomDID()
	dIDBytes  = did[:]
	accountID = did[:]
)

type mockAnchorRepo struct {
	mock.Mock
	anchors.AnchorRepository
}

func (r *mockAnchorRepo) GetDocumentRootOf(anchorID anchors.AnchorID) (anchors.DocumentRoot, error) {
	args := r.Called(anchorID)
	docRoot, _ := args.Get(0).(anchors.DocumentRoot)
	return docRoot, args.Error(1)
}

func TestMain(m *testing.M) {
	ethClient := &testingcommons.MockEthClient{}
	ethClient.On("GetEthClient").Return(nil)
	ctx[ethereum.BootstrappedEthereumClient] = ethClient
	txMan := &testingtx.MockTxManager{}
	ctx[transactions.BootstrappedService] = txMan
	done := make(chan bool)
	txMan.On("ExecuteWithinTX", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(transactions.NilTxID(), done, nil)
	ctx[nft.BootstrappedPayObService] = new(testingdocuments.MockRegistry)
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
		//&Bootstrapper{}, TODO activate bootstrapper for entity
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
	assert.NoError(t, entity.InitEntityInput(testingdocuments.CreateEntityPayload(), did.String()))

	cd, err := entity.PackCoreDocument()
	assert.NoError(t, err)
	assert.NotNil(t, cd.EmbeddedData)
}

func TestEntity_JSON(t *testing.T) {
	entity := new(Entity)
	ctx := testingconfig.CreateAccountContext(t, cfg)
	did, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)
	assert.NoError(t, entity.InitEntityInput(testingdocuments.CreateEntityPayload(), did.String()))

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
	assert.Equal(t, model.getClientData(), model.getClientData(), entity.(*Entity).getClientData())
	assert.Equal(t, model.ID(), entity.ID())
	assert.Equal(t, model.CurrentVersion(), entity.CurrentVersion())
	assert.Equal(t, model.PreviousVersion(), entity.PreviousVersion())
}

func TestEntityModel_getClientData(t *testing.T) {
	entityData := testingdocuments.CreateEntityData()
	entity := new(Entity)
	entity.loadFromP2PProtobuf(&entityData)

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
	err = e.InitEntityInput(&cliententitypb.EntityCreatePayload{Data: data}, did.String())
	assert.Nil(t, err, "should be successful")

	e = new(Entity)
	collabs := []string{"0x010102040506", "some id"}
	err = e.InitEntityInput(&cliententitypb.EntityCreatePayload{Data: data, Collaborators: collabs}, did.String())
	assert.Contains(t, err.Error(), "failed to decode collaborator")

	collab1, err := identity.NewDIDFromString("0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7")
	assert.NoError(t, err)
	collab2, err := identity.NewDIDFromString("0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF3")
	assert.NoError(t, err)
	collabs = []string{collab1.String(), collab2.String()}
	err = e.InitEntityInput(&cliententitypb.EntityCreatePayload{Data: data, Collaborators: collabs}, did.String())
	assert.Nil(t, err, "must be nil")

}

func TestEntityModel_calculateDataRoot(t *testing.T) {
	ctx := testingconfig.CreateAccountContext(t, cfg)
	did, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)
	m := new(Entity)
	err = m.InitEntityInput(testingdocuments.CreateEntityPayload(), did.String())
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

	// Validate entity_number
	valid, err := tree.ValidateProof(proof[0])
	assert.Nil(t, err)
	assert.True(t, valid)

	// Validate roles
	valid, err = tree.ValidateProof(proof[1])
	assert.Nil(t, err)
	assert.True(t, valid)

	// Validate []byte value
	acc := identity.NewDIDFromBytes(proof[1].Value)
	assert.True(t, e.AccountCanRead(acc))

	// Validate document_type
	valid, err = tree.ValidateProof(proof[2])
	assert.Nil(t, err)
	assert.True(t, valid)
}

func createEntity(t *testing.T) *Entity {
	e := new(Entity)
	err := e.InitEntityInput(testingdocuments.CreateEntityPayload(), defaultDID.String())
	assert.NoError(t, err)
	e.GetTestCoreDocWithReset()
	_, err = e.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = e.CalculateSigningRoot()
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

func TestEntityModel_getDocumentDataTree(t *testing.T) {
	e := createEntity(t)
	tree, err := e.getDocumentDataTree()
	assert.Nil(t, err, "tree should be generated without error")
	_, leaf := tree.GetLeafByProperty("entity.legal_name")
	assert.NotNil(t, leaf)
	assert.Equal(t, "entity.legal_name", leaf.Property.ReadableName())
}

func TestEntity_CollaboratorCanUpdate(t *testing.T) {
	entity := createEntity(t)
	id1 := defaultDID
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
	err = entity.PrepareNewVersion(entity, data, []string{id3.String()})
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
	err = entity.PrepareNewVersion(entity, data, nil)
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
	if testRepoGlobal == nil {
		ldb, err := leveldb.NewLevelDBStorage(leveldb.GetRandomTestStoragePath())
		if err != nil {
			panic(err)
		}
		testRepoGlobal = documents.NewDBRepository(leveldb.NewLevelDBRepository(ldb))
		testRepoGlobal.Register(&Entity{})
	}
	return testRepoGlobal
}

func createCDWithEmbeddedEntity(t *testing.T) (documents.Model, coredocumentpb.CoreDocument) {
	e := new(Entity)
	err := e.InitEntityInput(testingdocuments.CreateEntityPayload(), did.String())
	assert.NoError(t, err)
	e.GetTestCoreDocWithReset()
	_, err = e.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = e.CalculateSigningRoot()
	assert.NoError(t, err)
	_, err = e.CalculateDocumentRoot()
	assert.NoError(t, err)
	cd, err := e.PackCoreDocument()
	assert.NoError(t, err)
	return e, cd
}
