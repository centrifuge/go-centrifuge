// +build unit

package invoiceservice_test

import (
	"context"
	"crypto/sha256"
	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/coredocument"
	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/invoice"
	cc "github.com/CentrifugeInc/go-centrifuge/centrifuge/context/testing"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/invoice"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/invoice/service"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/testingutils"
	"github.com/centrifuge/precise-proofs/proofs"
	"github.com/go-errors/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	cc.TestIntegrationBootstrap()
	result := m.Run()
	cc.TestIntegrationTearDown()
	os.Exit(result)
}

// ----- MOCKS -----
type MockInvoiceRepository struct {
	mock.Mock
}

func (m *MockInvoiceRepository) GetKey(id []byte) []byte {
	args := m.Called(id)
	return args.Get(0).([]byte)
}
func (m *MockInvoiceRepository) FindById(id []byte) (inv *invoicepb.InvoiceDocument, err error) {
	args := m.Called(id)
	return args.Get(0).(*invoicepb.InvoiceDocument), args.Error(1)
}
func (m *MockInvoiceRepository) Store(inv *invoicepb.InvoiceDocument) (err error) {
	args := m.Called(inv)
	return args.Error(0)
}
func (m *MockInvoiceRepository) StoreOnce(inv *invoicepb.InvoiceDocument) (err error) {
	args := m.Called(inv)
	return args.Error(0)
}

// ----- END MOCKS -----

// ----- HELPER FUNCTIONS -----
func generateSendableInvoice() *invoice.Invoice {
	doc := invoice.NewEmptyInvoice()
	doc.Document.CoreDocument = testingutils.GenerateCoreDocument()
	return doc
}

func generateMockedOutInvoiceService() (srv *invoiceservice.InvoiceDocumentService, repo *MockInvoiceRepository, coreDocumentProcessor *testingutils.MockCoreDocumentProcessor) {
	repo = new(MockInvoiceRepository)
	coreDocumentProcessor = new(testingutils.MockCoreDocumentProcessor)
	srv = &invoiceservice.InvoiceDocumentService{
		InvoiceRepository:     repo,
		CoreDocumentProcessor: coreDocumentProcessor,
	}
	return
}
func getTestSetupData() (doc *invoice.Invoice, srv *invoiceservice.InvoiceDocumentService, repo *MockInvoiceRepository, coreDocumentProcessor *testingutils.MockCoreDocumentProcessor) {
	doc = generateSendableInvoice()
	srv, repo, coreDocumentProcessor = generateMockedOutInvoiceService()
	return
}

// ----- END HELPER FUNCTIONS -----

// ----- TESTS -----
func TestInvoiceDocumentService_Anchor(t *testing.T) {
	doc, s, mockRepo, mockCDP := getTestSetupData()

	mockRepo.On("StoreOnce", doc.Document).Return(nil).Once()
	mockCDP.On("Anchor", mock.Anything).Return(nil).Once()

	anchoredDoc, err := s.HandleAnchorInvoiceDocument(context.Background(), &invoicepb.AnchorInvoiceEnvelope{Document: doc.Document})

	mockRepo.AssertExpectations(t)
	mockCDP.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, doc.Document.CoreDocument.DocumentIdentifier, anchoredDoc.CoreDocument.DocumentIdentifier)
}

func TestInvoiceDocumentService_AnchorFails(t *testing.T) {
	doc, s, mockRepo, mockCDP := getTestSetupData()

	mockRepo.On("StoreOnce", doc.Document).Return(nil).Once()
	mockCDP.On("Anchor", mock.Anything).Return(errors.New("error anchoring")).Once()

	anchoredDoc, err := s.HandleAnchorInvoiceDocument(context.Background(), &invoicepb.AnchorInvoiceEnvelope{Document: doc.Document})

	mockRepo.AssertExpectations(t)
	mockCDP.AssertExpectations(t)
	assert.Error(t, err)
	assert.Nil(t, anchoredDoc)
}

func TestInvoiceDocumentService_AnchorFailsWithNilDocument(t *testing.T) {
	_, s, _, _ := getTestSetupData()

	anchoredDoc, err := s.HandleAnchorInvoiceDocument(context.Background(), &invoicepb.AnchorInvoiceEnvelope{})

	assert.Error(t, err)
	assert.Nil(t, anchoredDoc)
}

func TestInvoiceDocumentService_Send(t *testing.T) {
	doc, s, mockRepo, mockCDP := getTestSetupData()

	recipients := testingutils.GenerateP2PRecipients(1)

	mockRepo.On("StoreOnce", doc.Document).Return(nil).Once()
	mockCDP.On("Anchor", mock.Anything).Return(nil).Once()
	mockCDP.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	_, err := s.HandleSendInvoiceDocument(context.Background(), &invoicepb.SendInvoiceEnvelope{Recipients: recipients, Document: doc.Document})

	mockRepo.AssertExpectations(t)
	mockCDP.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestInvoiceDocumentService_SendFails(t *testing.T) {
	doc, s, mockRepo, mockCDP := getTestSetupData()
	recipients := testingutils.GenerateP2PRecipients(2)

	mockRepo.On("StoreOnce", doc.Document).Return(nil).Once()
	mockCDP.On("Anchor", mock.Anything).Return(nil).Once()
	mockCDP.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error sending")).Twice()

	_, err := s.HandleSendInvoiceDocument(context.Background(), &invoicepb.SendInvoiceEnvelope{Recipients: recipients, Document: doc.Document})

	mockCDP.AssertExpectations(t)
	//the error handling in the send handler simply prints out the list of errors without much formatting
	//OK for now but could be done nicer in the future
	assert.Equal(t, "[error sending error sending]", err.Error())
}

func TestInvoiceDocumentService_Send_StoreFails(t *testing.T) {
	doc, s, mockRepo, _ := getTestSetupData()
	recipients := testingutils.GenerateP2PRecipients(2)

	mockRepo.On("StoreOnce", doc.Document).Return(errors.New("error storing")).Once()

	_, err := s.HandleSendInvoiceDocument(context.Background(), &invoicepb.SendInvoiceEnvelope{Recipients: recipients, Document: doc.Document})

	mockRepo.AssertExpectations(t)
	assert.Equal(t, "error storing", err.Error())
}

func TestInvoiceDocumentService_Send_AnchorFails(t *testing.T) {
	doc, s, mockRepo, mockCDP := getTestSetupData()
	recipients := testingutils.GenerateP2PRecipients(2)

	mockRepo.On("StoreOnce", doc.Document).Return(nil).Once()
	mockCDP.On("Anchor", mock.Anything).Return(errors.New("error anchoring")).Once()

	_, err := s.HandleSendInvoiceDocument(context.Background(), &invoicepb.SendInvoiceEnvelope{Recipients: recipients, Document: doc.Document})

	mockRepo.AssertExpectations(t)
	mockCDP.AssertExpectations(t)
	assert.Equal(t, "error anchoring", err.Error())
}

func TestInvoiceDocumentService_HandleCreateInvoiceProof(t *testing.T) {
	identifier := testingutils.Rand32Bytes()
	inv := invoice.NewEmptyInvoice()
	inv.Document.CoreDocument = &coredocumentpb.CoreDocument{
		DocumentIdentifier: identifier,
		CurrentIdentifier:  identifier,
		NextIdentifier:     testingutils.Rand32Bytes(),
		// TODO: below should be actual merkle root
		DataRoot: testingutils.Rand32Bytes(),
	}
	inv.CalculateMerkleRoot()

	s, mockRepo, _ := generateMockedOutInvoiceService()

	proofRequest := &invoicepb.CreateInvoiceProofEnvelope{
		DocumentIdentifier: identifier,
		Fields:             []string{"currency", "sender_country", "gross_amount"},
	}
	mockRepo.On("FindById", proofRequest.DocumentIdentifier).Return(inv.Document, nil).Once()

	invoiceProof, err := s.HandleCreateInvoiceProof(context.Background(), proofRequest)
	mockRepo.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, identifier, invoiceProof.DocumentIdentifier)
	assert.Equal(t, len(proofRequest.Fields), len(invoiceProof.FieldProofs))
	assert.Equal(t, proofRequest.Fields[0], invoiceProof.FieldProofs[0].Property)
	sha256Hash := sha256.New()
	valid, err := proofs.ValidateProof(invoiceProof.FieldProofs[0], inv.Document.CoreDocument.DocumentRoot, sha256Hash)
	assert.True(t, valid)
	assert.Nil(t, err)
}

func TestInvoiceDocumentService_HandleCreateInvoiceProof_NotExistingInvoice(t *testing.T) {
	identifier := testingutils.Rand32Bytes()
	inv := invoice.NewEmptyInvoice()
	inv.Document.CoreDocument = &coredocumentpb.CoreDocument{
		DocumentIdentifier: identifier,
		CurrentIdentifier:  identifier,
		NextIdentifier:     testingutils.Rand32Bytes(),
		// TODO: below should be actual merkle root
		DataRoot: testingutils.Rand32Bytes(),
	}
	inv.CalculateMerkleRoot()

	s, mockRepo, _ := generateMockedOutInvoiceService()

	proofRequest := &invoicepb.CreateInvoiceProofEnvelope{
		DocumentIdentifier: identifier,
		Fields:             []string{"currency", "sender_country", "gross_amount"},
	}
	//return an "empty" invoice as mock can't return nil
	mockRepo.On("FindById", proofRequest.DocumentIdentifier).Return(invoice.NewEmptyInvoice().Document, errors.New("couldn't find invoice")).Once()

	invoiceProof, err := s.HandleCreateInvoiceProof(context.Background(), proofRequest)
	mockRepo.AssertExpectations(t)

	assert.Nil(t, invoiceProof)
	assert.Error(t, err)
}
