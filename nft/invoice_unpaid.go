package nft

import (
	"context"
	"math/big"

	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	// TokenIDLength is the length of an NFT token ID
	TokenIDLength = 32

	// LowEntropyTokenIDLength is the length of a low entropy NFT token ID. Used only for special cases.
	LowEntropyTokenIDLength = 7
)

// TokenID is uint256 in Solidity (256 bits | max value is 2^256-1)
// tokenID should be random 32 bytes (32 byte = 256 bits)
type TokenID [TokenIDLength]byte

// NewTokenID returns a new random TokenID
func NewTokenID() TokenID {
	var tid [TokenIDLength]byte
	copy(tid[:], utils.RandomSlice(TokenIDLength))
	return tid
}

// NewLowEntropyTokenID returns a new random TokenID
func NewLowEntropyTokenID() TokenID {
	var tid [TokenIDLength]byte
	// prefix with zeroes to match the bigendian big integer bytes for smart contract
	copy(tid[:], append(make([]byte, TokenIDLength-LowEntropyTokenIDLength), utils.RandomSlice(LowEntropyTokenIDLength)...))
	return tid
}

// TokenIDFromString converts given hex string to a TokenID
func TokenIDFromString(hexStr string) (TokenID, error) {
	tokenIDBytes, err := hexutil.Decode(hexStr)
	if err != nil {
		return NewTokenID(), err
	}
	if len(tokenIDBytes) != TokenIDLength {
		return NewTokenID(), errors.New("the provided hex string doesn't match the TokenID representation length")
	}
	var tid [TokenIDLength]byte
	copy(tid[:], tokenIDBytes)
	return tid, nil
}

// BigInt converts tokenID to big int
func (t TokenID) BigInt() *big.Int {
	return utils.ByteSliceToBigInt(t[:])
}

func (t TokenID) String() string {
	return hexutil.Encode(t[:])
}

// MintNFTRequest holds required fields for minting NFT
type MintNFTRequest struct {
	DocumentID               []byte
	RegistryAddress          common.Address
	DepositAddress           common.Address
	ProofFields              []string
	GrantNFTReadAccess       bool
	SubmitTokenProof         bool
	SubmitNFTReadAccessProof bool
}

// InvoiceUnpaid handles transactions related to minting of NFTs for unpaid invoices
type InvoiceUnpaid interface {
	// MintNFT mints an NFT
	MintNFT(ctx context.Context, request MintNFTRequest) (*MintNFTResponse, chan bool, error)
	// GetRequiredInvoiceUnpaidProofFields returns the required proof field properties
	GetRequiredInvoiceUnpaidProofFields(ctx context.Context) ([]string, error)
}

// MintNFTResponse holds tokenID and transaction ID.
type MintNFTResponse struct {
	TokenID       string
	TransactionID string
}
