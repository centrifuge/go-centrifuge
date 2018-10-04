// +build unit

package invoice

import (
	"testing"

	"github.com/centrifuge/go-centrifuge/centrifuge/documents"
	clientinvoicepb "github.com/centrifuge/go-centrifuge/centrifuge/protobufs/gen/go/invoice"
	"github.com/centrifuge/go-centrifuge/centrifuge/testingutils/documents"
	"github.com/stretchr/testify/assert"
)

var invService Service

func createPayload() *clientinvoicepb.InvoiceCreatePayload {
	return &clientinvoicepb.InvoiceCreatePayload{
		Data: &clientinvoicepb.InvoiceData{
			GrossAmount: 42,
		},
	}
}

func TestDefaultService(t *testing.T) {
	srv := DefaultService(GetRepository())
	assert.NotNil(t, srv, "must be non-nil")
}

func TestService_DeriveFromCoreDocument(t *testing.T) {
	// nil doc
	_, err := invService.DeriveFromCoreDocument(nil)
	assert.Error(t, err, "must fail to derive")

	// successful
	data := testinginvoice.CreateInvoiceData()
	coreDoc := testinginvoice.CreateCDWithEmbeddedInvoice(t, data)
	model, err := invService.DeriveFromCoreDocument(coreDoc)
	assert.Nil(t, err, "must return model")
	assert.NotNil(t, model, "model must be non-nil")
	inv, ok := model.(*InvoiceModel)
	assert.True(t, ok, "must be true")
	assert.Equal(t, inv.Payee[:], data.Payee)
	assert.Equal(t, inv.Sender[:], data.Sender)
	assert.Equal(t, inv.Recipient[:], data.Recipient)
	assert.Equal(t, inv.GrossAmount, data.GrossAmount)
}

func TestService_DeriveFromPayload(t *testing.T) {
	payload := createPayload()
	var model documents.Model
	var err error

	// fail due to nil payload
	_, err = invService.DeriveFromCreatePayload(nil)
	assert.Error(t, err, "DeriveWithInvoiceInput should produce an error if invoiceInput equals nil")

	model, err = invService.DeriveFromCreatePayload(payload)
	assert.Nil(t, err, "valid invoiceData shouldn't produce an error")

	receivedCoreDocument, err := model.PackCoreDocument()
	assert.Nil(t, err, "model should be able to return the core document with embedded invoice")
	assert.NotNil(t, receivedCoreDocument.EmbeddedData, "embeddedData should be field")
}

func TestService_Create(t *testing.T) {
	payload := createPayload()
	inv, err := invService.DeriveFromCreatePayload(payload)
	assert.Nil(t, err, "must be non nil")

	// successful creation
	err = invService.Create(inv)
	assert.Nil(t, err, "create must pass")

	coredoc, err := inv.PackCoreDocument()
	assert.Nil(t, err, "must be converted to coredocument")

	loadInv := new(InvoiceModel)
	err = GetRepository().LoadByID(coredoc.CurrentIdentifier, loadInv)
	assert.Nil(t, err, "Load must pass")
	assert.NotNil(t, loadInv, "must be non nil")

	invType := inv.(*InvoiceModel)
	assert.Equal(t, loadInv.GrossAmount, invType.GrossAmount)
	assert.Equal(t, loadInv.CoreDocument, invType.CoreDocument)

	// failed creation
	err = invService.Create(inv)
	assert.Error(t, err, "must fail")
	assert.Contains(t, err.Error(), "document already exists")
}

func TestService_DeriveCreateResponse(t *testing.T) {
	// some random model
	_, err := invService.DeriveCreateResponse(&mockModel{})
	assert.Error(t, err, "Derive must fail")

	// success
	payload := createPayload()
	inv, err := invService.DeriveFromCreatePayload(payload)
	assert.Nil(t, err, "must be non nil")
	data, err := invService.DeriveCreateResponse(inv)
	assert.Nil(t, err, "Derive must succeed")
	assert.NotNil(t, data, "data must be non nil")
}
