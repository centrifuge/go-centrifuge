package testingutils

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/coredocument"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/config"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/tools"
	"github.com/centrifuge/precise-proofs/proofs"
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
		recipients[i] = []byte(fmt.Sprintf("RecipientNo[%d]", i))
	}
	return recipients
}

func GenerateCoreDocument() *coredocumentpb.CoreDocument {
	identifier := Rand32Bytes()
	salts := &coredocumentpb.CoreDocumentSalts{}
	doc := &coredocumentpb.CoreDocument{
		DataRoot:           tools.RandomSlice(32),
		DocumentIdentifier: identifier,
		CurrentIdentifier:  identifier,
		NextIdentifier:     Rand32Bytes(),
		CoredocumentSalts:  salts,
	}
	proofs.FillSalts(doc, salts)
	return doc
}

type MockCoreDocumentProcessor struct {
	mock.Mock
}

func (m *MockCoreDocumentProcessor) Send(coreDocument *coredocumentpb.CoreDocument, ctx context.Context, recipient []byte) (err error) {
	args := m.Called(coreDocument, ctx, recipient)
	return args.Error(0)
}

func (m *MockCoreDocumentProcessor) Anchor(coreDocument *coredocumentpb.CoreDocument) (err error) {
	args := m.Called(coreDocument)
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
