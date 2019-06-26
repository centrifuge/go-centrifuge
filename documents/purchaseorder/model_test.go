// +build unit

package purchaseorder

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/centrifuge/centrifuge-protobufs/documenttypes"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
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
	clientpurchaseorderpb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/purchaseorder"
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
var defaultDID = testingidentity.GenerateRandomDID()

func TestMain(m *testing.M) {
	ethClient := &ethereum.MockEthClient{}
	ethClient.On("GetEthClient").Return(nil)
	ctx[ethereum.BootstrappedEthereumClient] = ethClient
	jobManager := &testingjobs.MockJobManager{}
	ctx[jobs.BootstrappedService] = jobManager
	done := make(chan bool)
	jobManager.On("ExecuteWithinJob", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(jobs.NilJobID(), done, nil)
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
		&Bootstrapper{},
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

func TestPurchaseOrder_PackCoreDocument(t *testing.T) {
	ctx := testingconfig.CreateAccountContext(t, cfg)
	did, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)

	po := new(PurchaseOrder)
	assert.NoError(t, po.InitPurchaseOrderInput(testingdocuments.CreatePOPayload(), did))

	cd, err := po.PackCoreDocument()
	assert.NoError(t, err)
	assert.NotNil(t, cd.EmbeddedData)
}

func TestPurchaseOrder_JSON(t *testing.T) {
	po := new(PurchaseOrder)
	ctx := testingconfig.CreateAccountContext(t, cfg)
	did, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)
	assert.NoError(t, po.InitPurchaseOrderInput(testingdocuments.CreatePOPayload(), did))

	cd, err := po.PackCoreDocument()
	assert.NoError(t, err)
	jsonBytes, err := po.JSON()
	assert.NoError(t, err, "marshal to json didn't work correctly")
	assert.True(t, json.Valid(jsonBytes), "json format not correct")

	po = new(PurchaseOrder)
	err = po.FromJSON(jsonBytes)
	assert.NoError(t, err, "unmarshal JSON didn't work correctly")

	ncd, err := po.PackCoreDocument()
	assert.NoError(t, err, "JSON unmarshal damaged invoice variables")
	assert.Equal(t, cd, ncd)
}

func TestPO_UnpackCoreDocument(t *testing.T) {
	var model = new(PurchaseOrder)
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
			TypeUrl: documenttypes.PurchaseOrderDataTypeUrl,
		},
	})
	assert.Error(t, err)

	// successful
	po, cd := createCDWithEmbeddedPO(t)
	err = model.UnpackCoreDocument(cd)
	assert.NoError(t, err)
	data, err := model.getClientData()
	assert.NoError(t, err)
	data1, err := po.(*PurchaseOrder).getClientData()
	assert.NoError(t, err)
	assert.Equal(t, data, data1)
	assert.Equal(t, model.ID(), po.ID())
	assert.Equal(t, model.CurrentVersion(), po.CurrentVersion())
	assert.Equal(t, model.PreviousVersion(), po.PreviousVersion())
}

func TestPOModel_getClientData(t *testing.T) {
	poData := testingdocuments.CreatePOData()
	poModel := new(PurchaseOrder)
	poModel.CoreDocument = &documents.CoreDocument{}
	err := poModel.loadFromP2PProtobuf(&poData)
	assert.NoError(t, err)

	data, err := poModel.getClientData()
	assert.NoError(t, err)
	assert.NotNil(t, data, "purchase order data should not be nil")
	assert.Equal(t, data.TotalAmount, data.TotalAmount, "gross amount must match")
	assert.Equal(t, data.Recipient, poModel.Data.Recipient.String(), "recipient should match")
}

func TestPOOrderModel_InitPOInput(t *testing.T) {
	ctx := testingconfig.CreateAccountContext(t, cfg)
	did, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)

	// fail recipient
	data := &clientpurchaseorderpb.PurchaseOrderData{
		Recipient: "some recipient",
	}
	poModel := new(PurchaseOrder)
	err = poModel.InitPurchaseOrderInput(&clientpurchaseorderpb.PurchaseOrderCreatePayload{Data: data}, did)
	assert.Error(t, err, "must return err")
	assert.Contains(t, err.Error(), "malformed address provided")
	assert.Nil(t, poModel.Data.Recipient)

	data.Recipient = "0xed03fa80291ff5ddc284de6b51e716b130b05e20"
	err = poModel.InitPurchaseOrderInput(&clientpurchaseorderpb.PurchaseOrderCreatePayload{Data: data}, did)
	assert.Nil(t, err)
	assert.NotNil(t, poModel.Data.Recipient)

	collabs := []string{"0x010102040506", "some id"}
	err = poModel.InitPurchaseOrderInput(&clientpurchaseorderpb.PurchaseOrderCreatePayload{Data: data, WriteAccess: collabs}, did)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "malformed address provided")

	collab1, err := identity.NewDIDFromString("0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7")
	assert.NoError(t, err)
	collab2, err := identity.NewDIDFromString("0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF3")
	assert.NoError(t, err)
	collabs = []string{collab1.String(), collab2.String()}
	err = poModel.InitPurchaseOrderInput(&clientpurchaseorderpb.PurchaseOrderCreatePayload{Data: data, WriteAccess: collabs}, did)
	assert.Nil(t, err, "must be nil")

	did, err = identity.NewDIDFromString("0xed03fa80291ff5ddc284de6b51e716b130b05e20")
	assert.NoError(t, err)
	assert.Equal(t, poModel.Data.Recipient[:], did[:])
}

func TestPOModel_calculateDataRoot(t *testing.T) {
	ctx := testingconfig.CreateAccountContext(t, cfg)
	did, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)
	poModel := new(PurchaseOrder)
	err = poModel.InitPurchaseOrderInput(testingdocuments.CreatePOPayload(), did)
	assert.Nil(t, err, "Init must pass")

	dr, err := poModel.CalculateDataRoot()
	assert.Nil(t, err, "calculate must pass")
	assert.False(t, utils.IsEmptyByteSlice(dr))
}

func TestPOModel_CreateProofs(t *testing.T) {
	po := createPurchaseOrder(t)
	assert.NotNil(t, po)
	rk := po.CoreDocument.GetTestCoreDocWithReset().Roles[0].RoleKey
	pf := fmt.Sprintf(documents.CDTreePrefix+".roles[%s].collaborators[0]", hexutil.Encode(rk))
	proof, err := po.CreateProofs([]string{"po.number", pf, documents.CDTreePrefix + ".document_type", "po.line_items[0].status"})
	assert.Nil(t, err)
	assert.NotNil(t, proof)
	tree, err := po.DocumentRootTree()
	assert.NoError(t, err)
	h := sha256.New()
	dataLeaves, err := po.getDataLeaves()
	assert.NoError(t, err)

	// Validate po_number
	err = po.CoreDocument.ValidateDataProof("po.number", po.DocumentType(), tree.RootHash(), proof[0], dataLeaves, h)
	assert.Nil(t, err)

	// Validate roles collaborators
	err = po.CoreDocument.ValidateDataProof(pf, po.DocumentType(), tree.RootHash(), proof[1], dataLeaves, h)
	assert.Nil(t, err)

	// Validate []byte value
	acc, err := identity.NewDIDFromBytes(proof[1].Value)
	assert.NoError(t, err)
	assert.True(t, po.AccountCanRead(acc))

	// Validate document_type
	err = po.CoreDocument.ValidateDataProof(documents.CDTreePrefix+".document_type", po.DocumentType(), tree.RootHash(), proof[2], dataLeaves, h)
	assert.Nil(t, err)

	// validate line items
	err = po.CoreDocument.ValidateDataProof("po.line_items[0].status", po.DocumentType(), tree.RootHash(), proof[3], dataLeaves, h)
	assert.Nil(t, err)
}

func TestPOModel_createProofsFieldDoesNotExist(t *testing.T) {
	poModel := createPurchaseOrder(t)
	_, err := poModel.CreateProofs([]string{"nonexisting"})
	assert.NotNil(t, err)
}

func TestPOModel_getDataTree(t *testing.T) {
	na := new(documents.Decimal)
	assert.NoError(t, na.SetString("2"))
	poModel := createPurchaseOrder(t)
	poModel.Data.Number = "123"
	poModel.Data.TotalAmount = na
	tree, err := poModel.getDataTree()
	assert.Nil(t, err, "tree should be generated without error")
	_, leaf := tree.GetLeafByProperty("po.number")
	assert.NotNil(t, leaf)
	assert.Equal(t, "po.number", leaf.Property.ReadableName())
	assert.Equal(t, []byte(poModel.Data.Number), leaf.Value)
}

func createPurchaseOrder(t *testing.T) *PurchaseOrder {
	po := new(PurchaseOrder)
	payload := testingdocuments.CreatePOPayload()
	payload.Data.LineItems = []*clientpurchaseorderpb.LineItem{
		{
			Status:      "pending",
			AmountTotal: "1.1",
			Activities: []*clientpurchaseorderpb.LineItemActivity{
				{
					ItemNumber: "12345",
					Status:     "pending",
					Amount:     "1.1",
				},
			},
			TaxItems: []*clientpurchaseorderpb.TaxItem{
				{
					ItemNumber: "12345",
					TaxAmount:  "1.1",
				},
			},
		},
	}
	err := po.InitPurchaseOrderInput(payload, defaultDID)
	assert.NoError(t, err)
	po.GetTestCoreDocWithReset()
	_, err = po.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = po.CalculateDataRoot()
	assert.NoError(t, err)
	_, err = po.CalculateDocumentRoot()
	assert.NoError(t, err)
	return po
}

func TestPurchaseOrder_CollaboratorCanUpdate(t *testing.T) {
	po := createPurchaseOrder(t)
	id1 := defaultDID
	id2 := testingidentity.GenerateRandomDID()
	id3 := testingidentity.GenerateRandomDID()

	// wrong type
	err := po.CollaboratorCanUpdate(new(mockModel), id1)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalidType, err))
	assert.NoError(t, testRepo().Create(id1[:], po.CurrentVersion(), po))

	// update the document
	model, err := testRepo().Get(id1[:], po.CurrentVersion())
	assert.NoError(t, err)
	oldPO := model.(*PurchaseOrder)
	data, err := oldPO.getClientData()
	assert.NoError(t, err)
	data.TotalAmount = "50"
	err = po.PrepareNewVersion(po, data, documents.CollaboratorsAccess{
		ReadWriteCollaborators: []identity.DID{id3},
	}, oldPO.Attributes)
	assert.NoError(t, err)

	// id1 should have permission
	assert.NoError(t, oldPO.CollaboratorCanUpdate(po, id1))

	// id2 should fail since it doesn't have the permission to update
	assert.Error(t, oldPO.CollaboratorCanUpdate(po, id2))

	// update the id3 rules to update only total amount
	po.CoreDocument.Document.TransitionRules[3].MatchType = coredocumentpb.FieldMatchType_FIELD_MATCH_TYPE_EXACT
	po.CoreDocument.Document.TransitionRules[3].Field = append(compactPrefix(), 0, 0, 0, 18)
	assert.NoError(t, testRepo().Create(id1[:], po.CurrentVersion(), po))

	// fetch the document
	model, err = testRepo().Get(id1[:], po.CurrentVersion())
	assert.NoError(t, err)
	oldPO = model.(*PurchaseOrder)
	data, err = oldPO.getClientData()
	assert.NoError(t, err)
	data.TotalAmount = "55"
	data.Currency = "INR"
	err = po.PrepareNewVersion(po, data, documents.CollaboratorsAccess{}, oldPO.Attributes)
	assert.NoError(t, err)

	// id1 should have permission
	assert.NoError(t, oldPO.CollaboratorCanUpdate(po, id1))

	// id2 should fail since it doesn't have the permission to update
	assert.Error(t, oldPO.CollaboratorCanUpdate(po, id2))

	// id3 should fail with just one error since changing Currency is not allowed
	err = oldPO.CollaboratorCanUpdate(po, id3)
	assert.Error(t, err)
	assert.Equal(t, 1, errors.Len(err))
	assert.Contains(t, err.Error(), "po.currency")
}

func TestPurchaseOrder_AddAttributes(t *testing.T) {
	po, _ := createCDWithEmbeddedPO(t)
	label := "some key"
	value := "some value"
	attr, err := documents.NewAttribute(label, documents.AttrString, value)
	assert.NoError(t, err)

	// success
	err = po.AddAttributes(documents.CollaboratorsAccess{}, true, attr)
	assert.NoError(t, err)
	assert.True(t, po.AttributeExists(attr.Key))
	gattr, err := po.GetAttribute(attr.Key)
	assert.NoError(t, err)
	assert.Equal(t, attr, gattr)

	// fail
	attr.Value.Type = documents.AttributeType("some attr")
	err = po.AddAttributes(documents.CollaboratorsAccess{}, true, attr)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrCDAttribute, err))
}

func TestPurchaseOrder_DeleteAttribute(t *testing.T) {
	po, _ := createCDWithEmbeddedPO(t)
	label := "some key"
	value := "some value"
	attr, err := documents.NewAttribute(label, documents.AttrString, value)
	assert.NoError(t, err)

	// failed
	err = po.DeleteAttribute(attr.Key, true)
	assert.Error(t, err)

	// success
	assert.NoError(t, po.AddAttributes(documents.CollaboratorsAccess{}, true, attr))
	assert.True(t, po.AttributeExists(attr.Key))
	assert.NoError(t, po.DeleteAttribute(attr.Key, true))
	assert.False(t, po.AttributeExists(attr.Key))
}

func TestPurchaseOrder_GetData(t *testing.T) {
	po := createPurchaseOrder(t)
	data := po.GetData()
	assert.Equal(t, po.Data, data)
}

func marshallData(t *testing.T, m map[string]interface{}) []byte {
	data, err := json.Marshal(m)
	assert.NoError(t, err)
	return data
}

func emptyDecimalData(t *testing.T) []byte {
	d := map[string]interface{}{
		"total_amount": "",
	}

	return marshallData(t, d)
}

func invalidDecimalData(t *testing.T) []byte {
	d := map[string]interface{}{
		"total_amount": "10.10.",
	}

	return marshallData(t, d)
}

func emptyDIDData(t *testing.T) []byte {
	d := map[string]interface{}{
		"recipient": "",
	}

	return marshallData(t, d)
}

func invalidDIDData(t *testing.T) []byte {
	d := map[string]interface{}{
		"recipient": "1acdew123asdefres",
	}

	return marshallData(t, d)
}

func emptyTimeData(t *testing.T) []byte {
	d := map[string]interface{}{
		"date_sent": "",
	}

	return marshallData(t, d)
}

func invalidTimeData(t *testing.T) []byte {
	d := map[string]interface{}{
		"date_sent": "1920-12-10",
	}

	return marshallData(t, d)
}

func validData(t *testing.T) []byte {
	d := map[string]interface{}{
		"number":         "12345",
		"status":         "unpaid",
		"total_amount":   "12.345",
		"recipient":      "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
		"date_sent":      "2019-05-24T14:48:44.308854Z", // rfc3339nano
		"date_confirmed": "2019-05-24T14:48:44Z",        // rfc3339
		"attachments": []map[string]interface{}{
			{
				"name":      "test",
				"file_type": "pdf",
				"size":      1000202,
				"data":      "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
				"checksum":  "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF3",
			},
		},
	}

	return marshallData(t, d)
}

func validDataWithCurrency(t *testing.T) []byte {
	d := map[string]interface{}{
		"number":         "12345",
		"status":         "unpaid",
		"total_amount":   "12.345",
		"recipient":      "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
		"date_sent":      "2019-05-24T14:48:44.308854Z", // rfc3339nano
		"date_confirmed": "2019-05-24T14:48:44Z",        // rfc3339
		"currency":       "EUR",
		"attachments": []map[string]interface{}{
			{
				"name":      "test",
				"file_type": "pdf",
				"size":      1000202,
				"data":      "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7",
				"checksum":  "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF3",
			},
		},
	}

	return marshallData(t, d)
}

func checkPOPayloadDataError(t *testing.T, po *PurchaseOrder, payload documents.CreatePayload) {
	err := po.loadData(payload.Data)
	assert.Error(t, err)
}

func TestPurchaseOrder_loadData(t *testing.T) {
	po := new(PurchaseOrder)
	payload := documents.CreatePayload{}

	// empty decimal data
	payload.Data = emptyDecimalData(t)
	checkPOPayloadDataError(t, po, payload)

	// invalid decimal data
	payload.Data = invalidDecimalData(t)
	checkPOPayloadDataError(t, po, payload)

	// empty did data
	payload.Data = emptyDIDData(t)
	checkPOPayloadDataError(t, po, payload)

	// invalid did data
	payload.Data = invalidDIDData(t)
	checkPOPayloadDataError(t, po, payload)

	// empty time data
	payload.Data = emptyTimeData(t)
	checkPOPayloadDataError(t, po, payload)

	// invalid time data
	payload.Data = invalidTimeData(t)
	checkPOPayloadDataError(t, po, payload)

	// valid data
	payload.Data = validData(t)
	err := po.loadData(payload.Data)
	assert.NoError(t, err)
	data := po.GetData().(Data)
	assert.Equal(t, data.Number, "12345")
	assert.Equal(t, data.Status, "unpaid")
	assert.Equal(t, data.TotalAmount.String(), "12.345")
	assert.Equal(t, data.Recipient.String(), "0xBAEb33a61f05e6F269f1c4b4CFF91A901B54DaF7")
	assert.Equal(t, data.DateSent.UTC().Format(time.RFC3339Nano), "2019-05-24T14:48:44.308854Z")
	assert.Equal(t, data.DateConfirmed.UTC().Format(time.RFC3339), "2019-05-24T14:48:44Z")
	assert.Len(t, data.Attachments, 1)
	assert.Equal(t, data.Attachments[0].Name, "test")
	assert.Equal(t, data.Attachments[0].FileType, "pdf")
	assert.Equal(t, data.Attachments[0].Size, 1000202)
	assert.Equal(t, hexutil.Encode(data.Attachments[0].Checksum), "0xbaeb33a61f05e6f269f1c4b4cff91a901b54daf3")
	assert.Equal(t, hexutil.Encode(data.Attachments[0].Data), "0xbaeb33a61f05e6f269f1c4b4cff91a901b54daf7")
}

func TestPurchaseOrder_unpackFromCreatePayload(t *testing.T) {
	payload := documents.CreatePayload{}
	po := new(PurchaseOrder)

	// invalid data
	payload.Data = invalidDecimalData(t)
	err := po.unpackFromCreatePayload(did, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(ErrPOInvalidData, err))

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
	err = po.unpackFromCreatePayload(did, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrCDCreate, err))

	// valid
	val.Type = documents.AttrString
	attr.Value = val
	payload.Attributes = map[documents.AttrKey]documents.Attribute{
		attr.Key: attr,
	}
	err = po.unpackFromCreatePayload(did, payload)
	assert.NoError(t, err)
}

func TestPurchaseOrder_unpackFromUpdatePayload(t *testing.T) {
	payload := documents.UpdatePayload{}
	old := createPurchaseOrder(t)
	po := new(PurchaseOrder)

	// invalid data
	payload.Data = invalidDecimalData(t)
	err := po.unpackFromUpdatePayload(old, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(ErrPOInvalidData, err))

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
	err = po.unpackFromUpdatePayload(old, payload)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrCDNewVersion, err))

	// valid
	val.Type = documents.AttrString
	attr.Value = val
	payload.Attributes = map[documents.AttrKey]documents.Attribute{
		attr.Key: attr,
	}
	err = po.unpackFromUpdatePayload(old, payload)
	assert.NoError(t, err)
}
