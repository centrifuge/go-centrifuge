// +build integration

package receiver_test

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/centrifuge/go-centrifuge/errors"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/p2p"
	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/bootstrap/bootstrappers/testingbootstrap"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/config/configstore"
	"github.com/centrifuge/go-centrifuge/contextutil"
	"github.com/centrifuge/go-centrifuge/crypto"
	cented25519 "github.com/centrifuge/go-centrifuge/crypto/ed25519"
	"github.com/centrifuge/go-centrifuge/crypto/secp256k1"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/documents/purchaseorder"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/p2p/common"
	"github.com/centrifuge/go-centrifuge/p2p/receiver"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/protocol"
	"github.com/centrifuge/go-centrifuge/storage"
	"github.com/centrifuge/go-centrifuge/testingutils/config"
	"github.com/centrifuge/go-centrifuge/testingutils/documents"
	"github.com/centrifuge/go-centrifuge/testingutils/identity"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

var (
	handler    *receiver.Handler
	anchorRepo anchors.AnchorRepository
	cfg        config.Configuration
	idService  identity.Service
	idFactory  identity.Factory
	cfgService config.Service
	docSrv     documents.Service
	defaultDID identity.DID
)

func TestMain(m *testing.M) {
	flag.Parse()
	ctx := testingbootstrap.TestFunctionalEthereumBootstrap()
	cfg = ctx[bootstrap.BootstrappedConfig].(config.Configuration)
	cfgService = ctx[config.BootstrappedConfigStorage].(config.Service)
	docSrv = ctx[documents.BootstrappedDocumentService].(documents.Service)
	anchorRepo = ctx[anchors.BootstrappedAnchorRepo].(anchors.AnchorRepository)
	idService = ctx[identity.BootstrappedDIDService].(identity.Service)
	idFactory = ctx[identity.BootstrappedDIDFactory].(identity.Factory)
	handler = receiver.New(cfgService, receiver.HandshakeValidator(cfg.GetNetworkID(), idService), docSrv, new(testingdocuments.MockRegistry), idService)
	defaultDID = createIdentity(&testing.T{})
	errors.MaskErrs = false
	result := m.Run()
	testingbootstrap.TestFunctionalEthereumTearDown()
	os.Exit(result)
}

func TestHandler_GetDocument_nonexistentIdentifier(t *testing.T) {
	b := utils.RandomSlice(32)
	req := &p2ppb.GetDocumentRequest{DocumentIdentifier: b}
	resp, err := handler.GetDocument(context.Background(), req, defaultDID)
	assert.Error(t, err, "must return error")
	assert.Nil(t, resp, "must be nil")
}

func TestHandler_HandleInterceptorReqSignature(t *testing.T) {
	tc, err := configstore.NewAccount("main", cfg)
	assert.Nil(t, err)
	acc := tc.(*configstore.Account)
	acc.IdentityID = defaultDID[:]
	ctxh, err := contextutil.New(context.Background(), acc)
	assert.Nil(t, err)
	_, err = cfgService.CreateAccount(acc)
	assert.NoError(t, err)
	po, cd := prepareDocumentForP2PHandler(t, nil)
	p2pEnv, err := p2pcommon.PrepareP2PEnvelope(ctxh, cfg.GetNetworkID(), p2pcommon.MessageTypeRequestSignature, &p2ppb.SignatureRequest{Document: &cd})

	pub, _ := acc.GetP2PKeyPair()
	publicKey, err := cented25519.GetPublicSigningKey(pub)
	assert.NoError(t, err)
	var bPk [32]byte
	copy(bPk[:], publicKey)
	peerID, err := cented25519.PublicKeyToP2PKey(bPk)
	assert.NoError(t, err)

	p2pResp, err := handler.HandleInterceptor(ctxh, peerID, p2pcommon.ProtocolForDID(&defaultDID), p2pEnv)
	assert.Nil(t, err, "must be nil")
	assert.NotNil(t, p2pResp, "must be non nil")
	resp := resolveSignatureResponse(t, p2pResp)
	assert.NotNil(t, resp.Signatures[0].Signature, "must be non nil")
	sig := resp.Signatures[0]
	ddr, err := po.CalculateDataRoot()
	assert.NoError(t, err)
	assert.True(t,
		secp256k1.VerifySignatureWithAddress(
			common.BytesToAddress(sig.PublicKey).String(),
			hexutil.Encode(sig.Signature),
			documents.ConsensusSignaturePayload(ddr, byte(0)),
		), "signature must be valid")
}

func TestHandler_RequestDocumentSignature(t *testing.T) {
	tc, err := configstore.NewAccount("main", cfg)
	assert.Nil(t, err)
	acc := tc.(*configstore.Account)
	acc.IdentityID = defaultDID[:]

	ctxh, err := contextutil.New(context.Background(), acc)
	assert.Nil(t, err)

	po, cd := prepareDocumentForP2PHandler(t, nil)

	// nil sigRequest
	id2 := testingidentity.GenerateRandomDID()
	_, err = handler.RequestDocumentSignature(ctxh, nil, defaultDID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nil document provided")

	// requestDocumentSignature, no previous versions
	_, err = handler.RequestDocumentSignature(ctxh, &p2ppb.SignatureRequest{Document: &cd}, defaultDID)
	assert.NoError(t, err)

	// we can update the document so that there are two versions in the repo
	po, ncd := updateDocumentForP2Phandler(t, po)
	assert.NotEqual(t, cd.DocumentIdentifier, ncd.CurrentVersion)

	// invalid transition for non-collaborator id
	_, err = handler.RequestDocumentSignature(ctxh, &p2ppb.SignatureRequest{Document: &ncd}, id2)
	assert.Error(t, err)

	// valid transition for collaborator id
	resp, err := handler.RequestDocumentSignature(ctxh, &p2ppb.SignatureRequest{Document: &ncd}, defaultDID)
	fmt.Println(ncd.PreviousVersion, ncd.CurrentVersion)
	assert.NoError(t, err)
	assert.NotNil(t, resp, "must be non nil")
	assert.NotNil(t, resp.Signatures[0].Signature, "must be non nil")
	sig := resp.Signatures[0]
	ddr, err := po.CalculateDataRoot()
	assert.NoError(t, err)
	assert.True(t,
		secp256k1.VerifySignatureWithAddress(
			common.BytesToAddress(sig.PublicKey).String(),
			hexutil.Encode(sig.Signature),
			documents.ConsensusSignaturePayload(ddr, byte(1)),
		), "signature must be valid")

	// document already exists
	_, err = handler.RequestDocumentSignature(ctxh, &p2ppb.SignatureRequest{Document: &cd}, defaultDID)
	assert.NotNil(t, err, "must not be nil")
	assert.Contains(t, err.Error(), storage.ErrRepositoryModelCreateKeyExists.Error())
}

func TestHandler_SendAnchoredDocument_update_fail(t *testing.T) {
	po, cd := prepareDocumentForP2PHandler(t, nil)
	ctx := testingconfig.CreateAccountContext(t, cfg)

	// Anchor document
	accDID, err := contextutil.AccountDID(ctx)
	assert.NoError(t, err)
	anchorIDTyped, err := anchors.ToAnchorID(cd.CurrentPreimage)
	assert.NoError(t, err)
	docRoot, err := po.CalculateDocumentRoot()
	assert.NoError(t, err)
	docRootTyped, err := anchors.ToDocumentRoot(docRoot)
	assert.NoError(t, err)

	anchorConfirmations, err := anchorRepo.CommitAnchor(ctx, anchorIDTyped, docRootTyped, utils.RandomByte32())
	assert.Nil(t, err)

	watchCommittedAnchor := <-anchorConfirmations
	assert.True(t, watchCommittedAnchor, "No error should be thrown by context")

	anchorResp, err := handler.SendAnchoredDocument(ctx, &p2ppb.AnchorDocumentRequest{Document: &cd}, accDID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), storage.ErrRepositoryModelUpdateKeyNotFound.Error())
	assert.Nil(t, anchorResp)
}

func TestHandler_SendAnchoredDocument_EmptyDocument(t *testing.T) {
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	id, err := cfg.GetIdentityID()
	collaborator, err := identity.NewDIDFromBytes(id)
	assert.NoError(t, err)
	resp, err := handler.SendAnchoredDocument(ctxh, &p2ppb.AnchorDocumentRequest{}, collaborator)
	assert.NotNil(t, err)
	assert.Nil(t, resp, "must be nil")
}

func TestHandler_SendAnchoredDocument(t *testing.T) {
	tc, err := configstore.NewAccount("main", cfg)
	assert.Nil(t, err)
	acc := tc.(*configstore.Account)
	acc.IdentityID = defaultDID[:]

	ctxh, err := contextutil.New(context.Background(), acc)
	assert.Nil(t, err)

	po, cd := prepareDocumentForP2PHandler(t, nil)
	resp, err := handler.RequestDocumentSignature(ctxh, &p2ppb.SignatureRequest{Document: &cd}, defaultDID)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	// Add signature received
	po.AppendSignatures(resp.Signatures...)

	// Since we have changed the coredocument by adding signatures lets generate salts again
	tree, err := po.DocumentRootTree()

	// Anchor document
	anchorIDTyped, err := anchors.ToAnchorID(po.GetTestCoreDocWithReset().CurrentPreimage)
	assert.NoError(t, err)
	docRootTyped, err := anchors.ToDocumentRoot(tree.RootHash())
	assert.NoError(t, err)

	anchorConfirmations, err := anchorRepo.CommitAnchor(ctxh, anchorIDTyped, docRootTyped, utils.RandomByte32())
	assert.Nil(t, err)

	watchCommittedAnchor := <-anchorConfirmations
	assert.True(t, watchCommittedAnchor, "No error should be thrown by context")
	cd, err = po.PackCoreDocument()
	assert.NoError(t, err)

	// this should succeed since this is the first document version
	anchorResp, err := handler.SendAnchoredDocument(ctxh, &p2ppb.AnchorDocumentRequest{Document: &cd}, defaultDID)
	assert.Nil(t, err)
	assert.NotNil(t, anchorResp, "must be non nil")
	assert.True(t, anchorResp.Accepted)

	// update the document
	npo, ncd := updateDocumentForP2Phandler(t, po)
	resp, err = handler.RequestDocumentSignature(ctxh, &p2ppb.SignatureRequest{Document: &ncd}, defaultDID)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	// Add signature received
	npo.AppendSignatures(resp.Signatures...)
	tree, err = npo.DocumentRootTree()

	// Anchor document
	anchorIDTyped, err = anchors.ToAnchorID(npo.GetTestCoreDocWithReset().CurrentPreimage)
	assert.NoError(t, err)
	docRootTyped, err = anchors.ToDocumentRoot(tree.RootHash())
	assert.NoError(t, err)
	anchorConfirmations, err = anchorRepo.CommitAnchor(ctxh, anchorIDTyped, docRootTyped, utils.RandomByte32())
	assert.Nil(t, err)

	watchCommittedAnchor = <-anchorConfirmations
	assert.True(t, watchCommittedAnchor, "No error should be thrown by context")
	ncd, err = npo.PackCoreDocument()
	assert.NoError(t, err)

	// transition failure for random ID
	id := testingidentity.GenerateRandomDID()
	_, err = handler.SendAnchoredDocument(ctxh, &p2ppb.AnchorDocumentRequest{Document: &ncd}, id)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid document state transition")

	anchorResp, err = handler.SendAnchoredDocument(ctxh, &p2ppb.AnchorDocumentRequest{Document: &ncd}, defaultDID)
	assert.Nil(t, err)
	assert.NotNil(t, anchorResp, "must be non nil")
	assert.True(t, anchorResp.Accepted)
}

func createIdentity(t *testing.T) identity.DID {
	// Create Identity
	didAddr, err := idFactory.CalculateIdentityAddress(context.Background())
	assert.NoError(t, err)
	tc, err := configstore.NewAccount("main", cfg)
	assert.Nil(t, err)
	acc := tc.(*configstore.Account)
	acc.IdentityID = didAddr.Bytes()

	ctx, err := contextutil.New(context.Background(), tc)
	assert.Nil(t, err)
	did, err := idFactory.CreateIdentity(ctx)
	assert.Nil(t, err, "should not error out when creating identity")
	assert.Equal(t, did.String(), didAddr.String(), "Resulting Identity should have the same ID as the input")

	// Add Keys
	accKeys, err := tc.GetKeys()
	assert.NoError(t, err)
	pk, err := utils.SliceToByte32(accKeys[identity.KeyPurposeP2PDiscovery.Name].PublicKey)
	assert.NoError(t, err)
	keyDID := identity.NewKey(pk, &(identity.KeyPurposeP2PDiscovery.Value), big.NewInt(identity.KeyTypeECDSA), 0)
	err = idService.AddKey(ctx, keyDID)
	assert.Nil(t, err, "should not error out when adding key to identity")

	sPk, err := utils.SliceToByte32(accKeys[identity.KeyPurposeSigning.Name].PublicKey)
	assert.NoError(t, err)
	keyDID = identity.NewKey(sPk, &(identity.KeyPurposeSigning.Value), big.NewInt(identity.KeyTypeECDSA), 0)
	err = idService.AddKey(ctx, keyDID)
	assert.Nil(t, err, "should not error out when adding key to identity")

	return *did
}

func prepareDocumentForP2PHandler(t *testing.T, po *purchaseorder.PurchaseOrder) (*purchaseorder.PurchaseOrder, coredocumentpb.CoreDocument) {
	ctx := testingconfig.CreateAccountContext(t, cfg)
	accCfg, err := contextutil.Account(ctx)
	assert.NoError(t, err)
	acc := accCfg.(*configstore.Account)
	acc.IdentityID = defaultDID[:]
	accKeys, err := acc.GetKeys()
	assert.NoError(t, err)
	if po == nil {
		payload := testingdocuments.CreatePOPayload()
		po = new(purchaseorder.PurchaseOrder)
		err = po.InitPurchaseOrderInput(payload, defaultDID)
		assert.NoError(t, err)
	}
	po.SetUsedAnchorRepoAddress(cfg.GetContractAddress(config.AnchorRepo))
	err = po.AddUpdateLog(defaultDID)
	assert.NoError(t, err)
	_, err = po.CalculateDataRoot()
	assert.NoError(t, err)
	ddr, err := po.CalculateDataRoot()
	assert.NoError(t, err)
	s, err := crypto.SignMessage(accKeys[identity.KeyPurposeSigning.Name].PrivateKey, documents.ConsensusSignaturePayload(ddr, byte(0)), crypto.CurveSecp256K1)
	assert.NoError(t, err)
	sig := &coredocumentpb.Signature{
		SignatureId: append(defaultDID[:], accKeys[identity.KeyPurposeSigning.Name].PublicKey...),
		SignerId:    defaultDID[:],
		PublicKey:   accKeys[identity.KeyPurposeSigning.Name].PublicKey,
		Signature:   s,
	}
	po.AppendSignatures(sig)
	_, err = po.CalculateDocumentRoot()
	assert.NoError(t, err)
	cd, err := po.PackCoreDocument()
	assert.NoError(t, err)
	return po, cd
}

func updateDocumentForP2Phandler(t *testing.T, po *purchaseorder.PurchaseOrder) (*purchaseorder.PurchaseOrder, coredocumentpb.CoreDocument) {
	cd, err := po.CoreDocument.PrepareNewVersion(nil, documents.CollaboratorsAccess{}, nil)
	assert.NoError(t, err)
	po.CoreDocument = cd
	return prepareDocumentForP2PHandler(t, po)
}

func resolveSignatureResponse(t *testing.T, p2pEnv *protocolpb.P2PEnvelope) *p2ppb.SignatureResponse {
	signResp := new(p2ppb.SignatureResponse)
	dataEnv, err := p2pcommon.ResolveDataEnvelope(p2pEnv)
	assert.NoError(t, err)
	err = proto.Unmarshal(dataEnv.Body, signResp)
	assert.NoError(t, err)
	return signResp
}
