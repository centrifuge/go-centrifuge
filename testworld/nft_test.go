// +build testworld

package testworld

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/nft"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/assert"
)

func TestInvoiceUnpaidMint_invoice_successful(t *testing.T) {
	t.Parallel()
	invoiceUnpaidMint(t, typeInvoice, true, true, true, false, "invoice")
}

func TestPaymentObligationWrapperMint_invoice_successful(t *testing.T) {
	t.Parallel()
	invoiceUnpaidMint(t, typeInvoice, false, false, false, true, "invoice")
}

/* TODO: testcase not stable
func TestInvoiceUnpaidMint_po_successful(t *testing.T) {
	t.Parallel()
	invoiceUnpaidMint(t, typePO)
}
*/

func invoiceUnpaidMint(t *testing.T, documentType string, grantNFTAccess, tokenProof, nftReadAccessProof bool, poWrapper bool, proofPrefix string) nft.TokenID {
	alice := doctorFord.getHostTestSuite(t, "Alice")
	bob := doctorFord.getHostTestSuite(t, "Bob")
	registry := alice.host.config.GetContractAddress(config.InvoiceUnpaidNFT)

	// Alice shares document with Bob
	res := createDocument(alice.httpExpect, alice.id.String(), documentType, http.StatusOK, defaultNFTPayload(documentType, []string{bob.id.String()}, alice.id.String()))
	txID := getTransactionID(t, res)
	status, message := getTransactionStatusAndMessage(alice.httpExpect, alice.id.String(), txID)
	if status != "success" {
		t.Error(message)
	}

	docIdentifier := getDocumentIdentifier(t, res)
	if docIdentifier == "" {
		t.Error("docIdentifier empty")
	}

	params := map[string]interface{}{
		"document_id": docIdentifier,
		"currency":    "USD",
	}
	getDocumentAndCheck(t, alice.httpExpect, alice.id.String(), documentType, params, false)
	getDocumentAndCheck(t, bob.httpExpect, bob.id.String(), documentType, params, false)

	var response *httpexpect.Object
	var err error

	depositAddress := alice.id.String()

	if !poWrapper {
		acc, err := alice.host.configService.GetAccount(alice.id[:])
		if err != nil {
			t.Error(err)
		}
		keys, err := acc.GetKeys()
		if err != nil {
			t.Error(err)
		}
		signerId := hexutil.Encode(append(alice.id[:], keys[identity.KeyPurposeSigning.Name].PublicKey...))
		dataRoot := fmt.Sprintf("%s.%s", documents.DRTreePrefix, documents.DataRootField)
		signatureSender := fmt.Sprintf("%s.signatures[%s]", documents.SignaturesTreePrefix, signerId)

		// mint an NFT
		payload := map[string]interface{}{
			"document_id":                   docIdentifier,
			"registry_address":              registry.String(),
			"deposit_address":               depositAddress, // Centrifuge address
			"proof_fields":                  []string{proofPrefix + ".gross_amount", proofPrefix + ".currency", proofPrefix + ".date_due", proofPrefix + ".sender", proofPrefix + ".status", dataRoot, signatureSender, documents.CDTreePrefix + ".next_version"},
			"submit_token_proof":            tokenProof,
			"submit_nft_owner_access_proof": nftReadAccessProof,
			"grant_nft_access":              grantNFTAccess,
		}
		response, err = alice.host.mintNFT(alice.httpExpect, alice.id.String(), http.StatusCreated, payload)

	} else {
		// mint a PO NFT
		payload := map[string]interface{}{
			"document_id":     docIdentifier,
			"deposit_address": depositAddress, // Centrifuge address
		}
		response, err = alice.host.mintUnpaidInvoiceNFT(alice.httpExpect, alice.id.String(), http.StatusOK, docIdentifier, payload)
	}

	assert.NoError(t, err, "mintNFT should be successful")
	txID = getTransactionID(t, response)
	status, message = getTransactionStatusAndMessage(alice.httpExpect, alice.id.String(), txID)
	if status != "success" {
		t.Error(message)
	}

	docVal := getDocumentAndCheck(t, alice.httpExpect, alice.id.String(), documentType, params, false)
	assert.True(t, len(docVal.Path("$.header.nfts[0].token_id").String().Raw()) > 0, "successful tokenId should have length 77")
	assert.True(t, len(docVal.Path("$.header.nfts[0].token_index").String().Raw()) > 0, "successful tokenIndex should have a value")

	tokenID, err := nft.TokenIDFromString(docVal.Path("$.header.nfts[0].token_id").String().Raw())
	assert.NoError(t, err, "token ID should be correct")
	respOwner := docVal.Path("$.header.nfts[0].owner").String().Raw()
	assert.NoError(t, err, "token ID should be correct")
	owner, err := alice.host.tokenRegistry.OwnerOf(registry, tokenID.BigInt().Bytes())
	assert.NoError(t, err)
	assert.Equal(t, strings.ToLower(depositAddress), strings.ToLower(owner.Hex()))
	assert.Equal(t, strings.ToLower(respOwner), strings.ToLower(owner.Hex()))
	return tokenID
}

func TestInvoiceUnpaidMint_errors(t *testing.T) {
	t.Parallel()
	alice := doctorFord.getHostTestSuite(t, "Alice")
	tests := []struct {
		errorMsg   string
		httpStatus int
		payload    map[string]interface{}
	}{
		{

			"RegistryAddress is not a valid Ethereum address",
			http.StatusBadRequest,
			map[string]interface{}{

				"registry_address": "0x123",
			},
		},
		{
			"cannot unmarshal hex string without 0x",
			http.StatusBadRequest,
			map[string]interface{}{

				"registry_address": "0xf72855759a39fb75fc7341139f5d7a3974d4da08", //dummy address
				"deposit_address":  "abc",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.errorMsg, func(t *testing.T) {
			t.Parallel()
			response, err := alice.host.mintNFT(alice.httpExpect, alice.id.String(), test.httpStatus, test.payload)
			assert.Nil(t, err, "it should be possible to call the API endpoint")
			response.Value("message").String().Contains(test.errorMsg)
		})
	}
}

func TestTransferNFT_successful(t *testing.T) {
	t.Parallel()
	tokenID := invoiceUnpaidMint(t, typeInvoice, false, false, false, true, "invoice")
	alice := doctorFord.getHostTestSuite(t, "Alice")
	bob := doctorFord.getHostTestSuite(t, "Bob")
	registry := alice.host.config.GetContractAddress(config.InvoiceUnpaidNFT)

	ownerOfPayload := map[string]interface{}{
		"token_id":         tokenID.String(),
		"registry_address": registry.String(),
	}

	transferPayload := map[string]interface{}{
		"token_id":         tokenID.String(),
		"registry_address": registry.String(),
		"to":               bob.id.String(),
	}

	// nft owner should be alice
	resp, err := alice.host.ownerOfNFT(alice.httpExpect, alice.id.String(), http.StatusOK, ownerOfPayload)
	assert.NoError(t, err)
	resp.Path("$.owner").String().Equal(strings.ToLower(alice.id.String()))

	// transfer nft from alice to bob
	response, err := alice.host.transferNFT(alice.httpExpect, alice.id.String(), http.StatusOK, transferPayload)
	assert.NoError(t, err)
	txID := getTransactionID(t, response)
	status, message := getTransactionStatusAndMessage(alice.httpExpect, alice.id.String(), txID)
	if status != "success" {
		t.Error(message)
	}

	// nft owner should be bob
	resp, err = alice.host.ownerOfNFT(alice.httpExpect, alice.id.String(), http.StatusOK, ownerOfPayload)
	assert.NoError(t, err)
	resp.Path("$.owner").String().Equal(strings.ToLower(bob.id.String()))
}
