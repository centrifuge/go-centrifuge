// +build integration

package nft_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/centrifuge/centrifuge-protobufs/documenttypes"
	"github.com/centrifuge/go-centrifuge/bootstrap"
	cc "github.com/centrifuge/go-centrifuge/bootstrap/bootstrappers/testingbootstrap"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/coredocument"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/documents/invoice"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/nft"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/invoice"
	"github.com/centrifuge/go-centrifuge/testingutils/config"
	"github.com/centrifuge/go-centrifuge/testingutils/identity"
	"github.com/centrifuge/go-centrifuge/transactions"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

var registry *documents.ServiceRegistry
var cfg config.Configuration
var idService identity.Service
var payOb nft.PaymentObligation
var txManager transactions.Manager
var tokenRegistry coredocument.TokenRegistry

func TestMain(m *testing.M) {
	log.Debug("Test PreSetup for NFT")
	ctx := cc.TestFunctionalEthereumBootstrap()
	registry = ctx[documents.BootstrappedRegistry].(*documents.ServiceRegistry)
	idService = ctx[identity.BootstrappedIDService].(identity.Service)
	cfg = ctx[bootstrap.BootstrappedConfig].(config.Configuration)
	payOb = ctx[nft.BootstrappedPayObService].(nft.PaymentObligation)
	txManager = ctx[transactions.BootstrappedService].(transactions.Manager)
	tokenRegistry = ctx[nft.BootstrappedPayObService].(coredocument.TokenRegistry)
	result := m.Run()
	cc.TestFunctionalEthereumTearDown()
	os.Exit(result)
}

func prepareForNFTMinting(t *testing.T) (context.Context, []byte, common.Address, string, documents.Service, identity.CentID) {
	// create identity
	log.Debug("Create Identity for Testing")
	cid := testingidentity.CreateIdentityWithKeys(cfg, idService)

	// create invoice (anchor)
	service, err := registry.LocateService(documenttypes.InvoiceDataTypeUrl)
	assert.Nil(t, err, "should not error out when getting invoice service")
	ctx := testingconfig.CreateAccountContext(t, cfg)
	invSrv := service.(invoice.Service)
	dueDate := time.Now().Add(4 * 24 * time.Hour)
	model, err := invSrv.DeriveFromCreatePayload(ctx, &invoicepb.InvoiceCreatePayload{
		Collaborators: []string{},
		Data: &invoicepb.InvoiceData{
			InvoiceNumber: "2132131",
			GrossAmount:   123,
			NetAmount:     123,
			Currency:      "EUR",
			DueDate:       &timestamp.Timestamp{Seconds: dueDate.Unix()},
		},
	})
	assert.Nil(t, err, "should not error out when creating invoice model")
	modelUpdated, txID, _, err := invSrv.Create(ctx, model)
	err = txManager.WaitForTransaction(cid, txID)
	assert.Nil(t, err)

	// get ID
	id, err := modelUpdated.ID()
	assert.Nil(t, err, "should not error out when getting invoice ID")
	// call mint
	// assert no error
	depositAddr := "0xf72855759a39fb75fc7341139f5d7a3974d4da08"
	registry := cfg.GetContractAddress(config.PaymentObligation)

	return ctx, id, registry, depositAddr, invSrv, cid
}

func mintNFT(t *testing.T, ctx context.Context, req nft.MintNFTRequest, cid identity.CentID, registry common.Address) nft.TokenID {
	resp, done, err := payOb.MintNFT(ctx, req)
	assert.Nil(t, err, "should not error out when minting an invoice")
	assert.NotNil(t, resp.TokenID, "token id should be present")
	tokenID, err := nft.TokenIDFromString(resp.TokenID)
	assert.Nil(t, err, "should not error out when getting tokenID hex")
	<-done
	assert.NoError(t, txManager.WaitForTransaction(cid, uuid.Must(uuid.FromString(resp.TransactionID))))
	owner, err := tokenRegistry.OwnerOf(registry, tokenID.BigInt().Bytes())
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress(req.DepositAddress), owner)
	return tokenID
}

func TestPaymentObligationService_mint_grant_read_access(t *testing.T) {
	ctx, id, registry, depositAddr, invSrv, cid := prepareForNFTMinting(t)
	req := nft.MintNFTRequest{
		DocumentID:         id,
		RegistryAddress:    registry.String(),
		DepositAddress:     depositAddr,
		ProofFields:        []string{"invoice.gross_amount", "invoice.currency", "invoice.due_date", "collaborators[0]"},
		GrantNFTReadAccess: true,
	}
	tokenID := mintNFT(t, ctx, req, cid, registry)
	doc, err := invSrv.GetCurrentVersion(ctx, id)
	assert.NoError(t, err)
	cd, err := doc.PackCoreDocument()
	assert.NoError(t, err)
	assert.Len(t, cd.Roles, 2)
	assert.Len(t, cd.Roles[1].Nfts, 1)
	newNFT := cd.Roles[1].Nfts[0]
	enft, err := coredocument.ConstructNFT(registry, tokenID.BigInt().Bytes())
	assert.NoError(t, err)
	assert.Equal(t, enft, newNFT)

	// try to mint the NFT again
	_, _, err = payOb.MintNFT(ctx, req)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(nft.ErrNFTMinted, err))
}

func failMintNFT(t *testing.T, grantNFT, nftReadAccess bool, roleProof string) {
	ctx, id, registry, depositAddr, _, _ := prepareForNFTMinting(t)
	req := nft.MintNFTRequest{
		DocumentID:               id,
		RegistryAddress:          registry.String(),
		DepositAddress:           depositAddr,
		ProofFields:              []string{"invoice.gross_amount", "invoice.currency", "invoice.due_date", "collaborators[0]"},
		GrantNFTReadAccess:       grantNFT,
		SubmitRoleProof:          roleProof,
		SubmitNFTReadAccessProof: nftReadAccess,
	}

	_, _, err := payOb.MintNFT(ctx, req)
	assert.Error(t, err)
	if !nftReadAccess {
		assert.True(t, errors.IsOfType(nft.ErrNFTRoleMissing, err))
	}
}

func TestEthereumPaymentObligation_MintNFT_role_not_exists(t *testing.T) {
	failMintNFT(t, true, false, "Supplier")
}

func TestEthereumPaymentObligation_MintNFT_no_grant_access(t *testing.T) {
	failMintNFT(t, false, true, "")
}

func mintNFTWithProofs(t *testing.T, grantAccess, tokenProof, readAccessProof bool) {
	ctx, id, registry, depositAddr, invSrv, cid := prepareForNFTMinting(t)
	req := nft.MintNFTRequest{
		DocumentID:               id,
		RegistryAddress:          registry.String(),
		DepositAddress:           depositAddr,
		ProofFields:              []string{"invoice.gross_amount", "invoice.currency", "invoice.due_date", "collaborators[0]"},
		GrantNFTReadAccess:       grantAccess,
		SubmitTokenProof:         tokenProof,
		SubmitNFTReadAccessProof: readAccessProof,
	}
	mintNFT(t, ctx, req, cid, registry)
	doc, err := invSrv.GetCurrentVersion(ctx, id)
	assert.NoError(t, err)
	cd, err := doc.PackCoreDocument()
	assert.NoError(t, err)
	roleCount := 1
	if grantAccess {
		roleCount++
	}
	assert.Len(t, cd.Roles, roleCount)
}

func TestEthereumPaymentObligation_MintNFT(t *testing.T) {
	tests := []struct {
		grantAccess, tokenProof, readAccessProof bool
	}{
		{
			grantAccess: true,
		},

		{
			tokenProof: true,
		},

		{
			grantAccess: true,
			tokenProof:  true,
		},

		{
			grantAccess:     true,
			readAccessProof: true,
		},

		{
			grantAccess:     true,
			tokenProof:      true,
			readAccessProof: true,
		},
	}

	for _, c := range tests {
		mintNFTWithProofs(t, c.grantAccess, c.tokenProof, c.readAccessProof)
	}
}
