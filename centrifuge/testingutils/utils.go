package testingutils

import (
	"context"
	"crypto/rand"

	"github.com/centrifuge/go-centrifuge/centrifuge/documents"

	"github.com/centrifuge/centrifuge-protobufs/documenttypes"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/centrifuge/config"
	"github.com/centrifuge/go-centrifuge/centrifuge/ethereum"
	"github.com/centrifuge/go-centrifuge/centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/centrifuge/keytools/secp256k1"
	"github.com/centrifuge/go-centrifuge/centrifuge/tools"
	"github.com/centrifuge/precise-proofs/proofs"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/stretchr/testify/mock"
)

func MockConfigOption(key string, value interface{}) func() {
	mockedValue := config.Config.V.Get(key)
	config.Config.V.Set(key, value)
	return func() {
		config.Config.V.Set(key, mockedValue)
	}
}

func Rand32Bytes() []byte {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return bytes
}

func GenerateP2PRecipients(quantity int) [][]byte {
	recipients := make([][]byte, quantity)

	for i := 0; i < quantity; i++ {
		ID := identity.NewRandomCentID()
		recipients[i] = ID[:]
	}
	return recipients
}

func GenerateCoreDocument() *coredocumentpb.CoreDocument {
	identifier := Rand32Bytes()
	salts := &coredocumentpb.CoreDocumentSalts{}
	doc := &coredocumentpb.CoreDocument{
		DataRoot:           tools.RandomSlice(32),
		DocumentIdentifier: identifier,
		CurrentVersion:     identifier,
		NextVersion:        Rand32Bytes(),
		CoredocumentSalts:  salts,
		EmbeddedData: &any.Any{
			TypeUrl: documenttypes.InvoiceDataTypeUrl,
		},
		EmbeddedDataSalts: &any.Any{
			TypeUrl: documenttypes.InvoiceSaltsTypeUrl,
		},
	}
	proofs.FillSalts(doc, salts)
	return doc
}

type MockCoreDocumentProcessor struct {
	mock.Mock
}

func (m *MockCoreDocumentProcessor) Send(ctx context.Context, coreDocument *coredocumentpb.CoreDocument, recipient identity.CentID) (err error) {
	args := m.Called(coreDocument, ctx, recipient)
	return args.Error(0)
}

func (m *MockCoreDocumentProcessor) Anchor(
	ctx context.Context,
	coreDocument *coredocumentpb.CoreDocument,
	saveState func(*coredocumentpb.CoreDocument) error) (err error) {
	args := m.Called(coreDocument)
	if saveState != nil {
		err := saveState(coreDocument)
		if err != nil {
			return err
		}
	}
	return args.Error(0)
}

func (m *MockCoreDocumentProcessor) PrepareForSignatureRequests(model documents.Model) error {
	args := m.Called(model)
	return args.Error(0)
}

func (m *MockCoreDocumentProcessor) RequestSignatures(ctx context.Context, model documents.Model) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockCoreDocumentProcessor) PrepareForAnchoring(model documents.Model) error {
	args := m.Called(model)
	return args.Error(0)
}

func (m *MockCoreDocumentProcessor) AnchorDocument(model documents.Model) error {
	args := m.Called(model)
	return args.Error(0)
}

func (m *MockCoreDocumentProcessor) SendDocument(ctx context.Context, model documents.Model) error {
	args := m.Called(ctx, model)
	return args.Error(0)
}

func (m *MockCoreDocumentProcessor) GetDataProofHashes(coreDocument *coredocumentpb.CoreDocument) (hashes [][]byte, err error) {
	args := m.Called(coreDocument)
	return args.Get(0).([][]byte), args.Error(1)
}

type MockSubscription struct {
	ErrChan chan error
}

func (m *MockSubscription) Err() <-chan error {
	return m.ErrChan
}

func (*MockSubscription) Unsubscribe() {}

func CreateIdentityWithKeys() identity.CentID {
	idConfig, _ := secp256k1.GetIDConfig()
	centIdTyped, _ := identity.ToCentID(idConfig.ID)
	// only create identity if it doesn't exist
	id, err := identity.IDService.LookupIdentityForID(centIdTyped)
	if err != nil {
		_, confirmations, _ := identity.IDService.CreateIdentity(centIdTyped)
		<-confirmations
		// LookupIdentityForId
		id, _ = identity.IDService.LookupIdentityForID(centIdTyped)
	}

	// only add key if it doesn't exist
	_, err = id.GetLastKeyForPurpose(identity.KeyPurposeEthMsgAuth)
	ctx, cancel := ethereum.DefaultWaitForTransactionMiningContext()
	defer cancel()
	if err != nil {
		confirmations, _ := id.AddKeyToIdentity(ctx, identity.KeyPurposeEthMsgAuth, idConfig.PublicKey)
		<-confirmations
	}
	return centIdTyped
}
