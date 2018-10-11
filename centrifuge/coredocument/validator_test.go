// +build unit

package coredocument

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/centrifuge/go-centrifuge/centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/centrifuge/testingutils/commons"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/centrifuge/config"
	cc "github.com/centrifuge/go-centrifuge/centrifuge/context/testingbootstrap"
	"github.com/centrifuge/go-centrifuge/centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/centrifuge/keytools/ed25519keys"
	"github.com/centrifuge/go-centrifuge/centrifuge/signatures"
	"github.com/centrifuge/go-centrifuge/centrifuge/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	cc.TestIntegrationBootstrap()
	flag.Parse()
	config.Config.V.Set("keys.signing.publicKey", "../../example/resources/signature1.pub.pem")
	config.Config.V.Set("keys.signing.privateKey", "../../example/resources/signature1.key.pem")
	result := m.Run()
	cc.TestIntegrationTearDown()
	os.Exit(result)
}

type mockModel struct {
	mock.Mock
	documents.Model
}

func (m mockModel) PackCoreDocument() (*coredocumentpb.CoreDocument, error) {
	args := m.Called()
	cd, _ := args.Get(0).(*coredocumentpb.CoreDocument)
	return cd, args.Error(1)
}

func Test_getCoreDocument(t *testing.T) {
	// nil document
	cd, err := getCoreDocument(nil)
	assert.Error(t, err)
	assert.Nil(t, cd)

	// pack core document fail
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("err")).Once()
	cd, err = getCoreDocument(model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Nil(t, cd)

	// success
	model = mockModel{}
	cd = New()
	model.On("PackCoreDocument").Return(cd, nil).Once()
	got, err := getCoreDocument(model)
	model.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, cd, got)
}

func TestValidator_fieldValidator(t *testing.T) {
	fv := fieldValidator()

	// fail getCoreDocument
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("err")).Once()
	err := fv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)

	// failed validator
	model = mockModel{}
	cd := New()
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = fv.Validate(nil, model)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cd_salts : Required field")

	// success
	model = mockModel{}
	cd.DataRoot = tools.RandomSlice(32)
	FillSalts(cd)
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = fv.Validate(nil, model)
	assert.Nil(t, err)
}

func TestValidator_signingRootValidator(t *testing.T) {
	sv := signingRootValidator()

	// fail getCoreDoc
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("err")).Once()
	err := sv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)

	// missing signing_root
	cd := New()
	FillSalts(cd)
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = sv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "signing root missing")

	// mismatch signing roots
	cd.SigningRoot = tools.RandomSlice(32)
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = sv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "signing root mismatch")

	// success
	tree, err := GetDocumentSigningTree(cd)
	assert.Nil(t, err)
	cd.SigningRoot = tree.RootHash()
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = sv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestValidator_documentRootValidator(t *testing.T) {
	dv := documentRootValidator()

	// fail getCoreDoc
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("err")).Once()
	err := dv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)

	// missing document root
	cd := New()
	FillSalts(cd)
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = dv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "document root missing")

	// mismatch signing roots
	cd.DocumentRoot = tools.RandomSlice(32)
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = dv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "document root mismatch")

	// success
	tree, err := GetDocumentRootTree(cd)
	assert.Nil(t, err)
	cd.DocumentRoot = tree.RootHash()
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = dv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestValidator_selfSignatureValidator(t *testing.T) {
	ssv := selfSignatureValidator()

	// fail getCoreDoc
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("err")).Once()
	err := ssv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)

	// signature length mismatch
	cd := New()
	FillSalts(cd)
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = ssv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "expecting only one signature")

	// mismatch
	cd.SigningRoot = tools.RandomSlice(32)
	s := &coredocumentpb.Signature{
		Signature: tools.RandomSlice(32),
		EntityId:  tools.RandomSlice(6),
		PublicKey: tools.RandomSlice(32),
	}
	cd.Signatures = append(cd.Signatures, s)
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = ssv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Len(t, documents.ConvertToMap(err), 3)

	// success
	cd.SigningRoot = tools.RandomSlice(32)
	c, err := ed25519keys.GetIDConfig()
	assert.Nil(t, err)
	s = signatures.Sign(c, cd.SigningRoot)
	cd.Signatures = []*coredocumentpb.Signature{s}
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = ssv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestValidator_senderSignatureValidator(t *testing.T) {
	ssv := senderSignatureValidator()

	// fail getCoreDoc
	model := mockModel{}
	model.On("PackCoreDocument").Return(nil, fmt.Errorf("err")).Once()
	err := ssv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)

	// signature length mismatch
	cd := New()
	FillSalts(cd)
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	err = ssv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "expecting only one signature")

	// failed validation
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	s := &coredocumentpb.Signature{EntityId: tools.RandomSlice(7)}
	cd.Signatures = append(cd.Signatures, s)
	err = ssv.Validate(nil, model)
	model.AssertExpectations(t)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to verify signature")

	// success
	model = mockModel{}
	model.On("PackCoreDocument").Return(cd, nil).Once()
	cd.SigningRoot = tools.RandomSlice(32)
	c, err := ed25519keys.GetIDConfig()
	assert.Nil(t, err)
	s = signatures.Sign(c, cd.SigningRoot)
	cd.Signatures = []*coredocumentpb.Signature{s}
	pubkey, err := tools.SliceToByte32(c.PublicKey)
	assert.Nil(t, err)
	idkey := &identity.EthereumIdentityKey{
		Key:       pubkey,
		Purposes:  []*big.Int{big.NewInt(identity.KeyPurposeSigning)},
		RevokedAt: big.NewInt(0),
	}
	id := &testingcommons.MockID{}
	srv := &testingcommons.MockIDService{}
	centID, err := identity.ToCentID(c.ID)
	assert.Nil(t, err)
	srv.On("LookupIdentityForID", centID).Return(id, nil).Once()
	id.On("FetchKey", pubkey[:]).Return(idkey, nil).Once()
	identity.IDService = srv
	err = ssv.Validate(nil, model)
	model.AssertExpectations(t)
	id.AssertExpectations(t)
	srv.AssertExpectations(t)
	assert.Nil(t, err)
}
