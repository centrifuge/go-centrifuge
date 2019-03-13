package identity

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/satori/go.uuid"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/crypto/ed25519"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common"
)

const (

	// ErrMalformedAddress standard error for malformed address
	ErrMalformedAddress = errors.Error("malformed address provided")

	// BootstrappedDIDFactory stores the id of the factory
	BootstrappedDIDFactory string = "BootstrappedDIDFactory"

	// BootstrappedDIDService stores the id of the service
	BootstrappedDIDService string = "BootstrappedDIDService"

	// KeyTypeECDSA has the value one in the ERC725 identity contract
	KeyTypeECDSA = 1

	keyPurposeMgmt         = "MANAGEMENT"
	keyPurposeAction       = "ACTION"
	keyPurposeP2PDiscovery = "P2P_DISCOVERY"
	keyPurposeSigning      = "SIGNING"
)

var (
	// KeyPurposeManagement purpose stores the management key to interact with the ERC725 identity contract
	KeyPurposeManagement Purpose
	// KeyPurposeAction purpose stores the action key to interact with the ERC725 identity contract
	KeyPurposeAction Purpose
	// KeyPurposeP2PDiscovery purpose stores the action key to interact with the ERC725 identity contract
	KeyPurposeP2PDiscovery Purpose
	// KeyPurposeSigning purpose stores the action key to interact with the ERC725 identity contract
	KeyPurposeSigning Purpose
)

func init() {
	KeyPurposeManagement = getKeyPurposeManagement()
	KeyPurposeAction = getKeyPurposeAction()
	KeyPurposeP2PDiscovery = getKeyPurposeP2PDiscovery()
	KeyPurposeSigning = getKeyPurposeSigning()
}

// getKeyPurposeManagement is calculated out of Hex(leftPadding(1,32))
func getKeyPurposeManagement() Purpose {
	enc := "0000000000000000000000000000000000000000000000000000000000000001"
	v, _ := new(big.Int).SetString(enc, 16)
	return Purpose{Name: keyPurposeMgmt, HexValue: enc, Value: *v}
}

// getKeyPurposeAction is calculated out of Hex(leftPadding(2,32))
func getKeyPurposeAction() Purpose {
	enc := "0000000000000000000000000000000000000000000000000000000000000002"
	v, _ := new(big.Int).SetString(enc, 16)
	return Purpose{Name: keyPurposeAction, HexValue: enc, Value: *v}
}

// getKeyPurposeP2PDiscovery is calculated out of Hex(sha256("CENTRIFUGE@P2P_DISCOVERY"))
func getKeyPurposeP2PDiscovery() Purpose {
	hashed := "88dbd1f0b244e515ab5aee93b5dee6a2d8e326576a583822635a27e52e5b591e"
	v, _ := new(big.Int).SetString(hashed, 16)
	return Purpose{Name: keyPurposeP2PDiscovery, HexValue: hashed, Value: *v}
}

// getKeyPurposeSigning is calculated out of Hex(sha256("CENTRIFUGE@SIGNING"))
func getKeyPurposeSigning() Purpose {
	hashed := "774a43710604e3ce8db630136980a6ba5a65b5e6686ee51009ed5f3fded6ea7e"
	v, _ := new(big.Int).SetString(hashed, 16)
	return Purpose{Name: keyPurposeSigning, HexValue: hashed, Value: *v}
}

// Purpose contains the different representation of purpose along the code
type Purpose struct {
	Name     string
	HexValue string
	Value    big.Int
}

// GetPurposeByName retrieves the Purpose by name
func GetPurposeByName(name string) Purpose {
	switch name {
	case keyPurposeMgmt:
		return getKeyPurposeManagement()
	case keyPurposeAction:
		return getKeyPurposeAction()
	case keyPurposeP2PDiscovery:
		return getKeyPurposeP2PDiscovery()
	case keyPurposeSigning:
		return getKeyPurposeSigning()
	default:
		return Purpose{}
	}
}

// DID stores the identity address of the user
type DID common.Address

// DIDLength contains the length of a DID
const DIDLength = common.AddressLength

// ToAddress returns the DID as common.Address
func (d DID) ToAddress() common.Address {
	return common.Address(d)
}

// String returns the DID as HEX String
func (d DID) String() string {
	return d.ToAddress().String()
}

// BigInt returns DID in bigInt
func (d DID) BigInt() *big.Int {
	return utils.ByteSliceToBigInt(d[:])
}

// Equal checks if d == other
func (d DID) Equal(other DID) bool {
	for i := range d {
		if d[i] != other[i] {
			return false
		}
	}
	return true
}

// NewDID returns a DID based on a common.Address
func NewDID(address common.Address) DID {
	return DID(address)
}

// NewDIDFromString returns a DID based on a hex string
func NewDIDFromString(address string) (DID, error) {
	if !common.IsHexAddress(address) {
		return DID{}, ErrMalformedAddress
	}
	return DID(common.HexToAddress(address)), nil
}

// NewDIDsFromStrings converts hex ids to DIDs
func NewDIDsFromStrings(ids []string) ([]DID, error) {
	var cids []DID
	for _, id := range ids {
		cid, err := NewDIDFromString(id)
		if err != nil {
			return nil, err
		}

		cids = append(cids, cid)
	}

	return cids, nil
}

// NewDIDFromBytes returns a DID based on a bytes input
func NewDIDFromBytes(bAddr []byte) DID {
	return DID(common.BytesToAddress(bAddr))
}

//// NewDIDFromContext returns DID from context.Account
//func NewDIDFromContext(ctx context.Context) (DID, error) {
//	tc, err := contextutil.Account(ctx)
//	if err != nil {
//		return DID{}, err
//	}
//
//	addressByte, err := tc.GetIdentityID()
//	if err != nil {
//		return DID{}, err
//	}
//	return NewDID(common.BytesToAddress(addressByte)), nil
//}

// Factory is the interface for factory related interactions
type Factory interface {
	CreateIdentity(ctx context.Context) (id *DID, err error)
	CalculateIdentityAddress(ctx context.Context) (*common.Address, error)
}

// ServiceDID interface contains the methods to interact with the identity contract
type ServiceDID interface {
	// AddKey adds a key to identity contract
	AddKey(ctx context.Context, key KeyDID) error

	// AddKeysForAccount adds key from configuration
	AddKeysForAccount(acc config.Account) error

	// GetKey return a key from the identity contract
	GetKey(did DID, key [32]byte) (*KeyResponse, error)

	// RawExecute calls the execute method on the identity contract
	RawExecute(ctx context.Context, to common.Address, data []byte) (utxID uuid.UUID, done chan bool, err error)

	// Execute creates the abi encoding an calls the execute method on the identity contract
	Execute(ctx context.Context, to common.Address, contractAbi, methodName string, args ...interface{}) (utxID uuid.UUID, done chan bool, err error)

	// IsSignedWithPurpose verifies if a message is signed with one of the identities specific purpose keys
	IsSignedWithPurpose(did DID, message [32]byte, signature []byte, purpose *big.Int) (bool, error)

	// AddMultiPurposeKey adds a key with multiple purposes
	AddMultiPurposeKey(context context.Context, key [32]byte, purposes []*big.Int, keyType *big.Int) error

	// RevokeKey revokes an existing key in the smart contract
	RevokeKey(ctx context.Context, key [32]byte) error

	// GetClientP2PURL returns the p2p url associated with the did
	GetClientP2PURL(did DID) (string, error)

	//Exists checks if an identity contract exists
	Exists(ctx context.Context, did DID) error

	// ValidateKey checks if a given key is valid for the given centrifugeID.
	ValidateKey(ctx context.Context, did DID, key []byte, purpose *big.Int, at *time.Time) error

	// ValidateSignature checks if signature is valid for given identity
	ValidateSignature(signature *coredocumentpb.Signature, message []byte) error

	// CurrentP2PKey retrieves the last P2P key stored in the identity
	CurrentP2PKey(did DID) (ret string, err error)

	// GetClientsP2PURLs returns p2p urls associated with each centIDs
	// will error out at first failure
	GetClientsP2PURLs(dids []*DID) ([]string, error)

	// GetKeysByPurpose returns keys grouped by purpose from the identity contract.
	GetKeysByPurpose(did DID, purpose *big.Int) ([][32]byte, error)
}

// KeyDID defines a single ERC725 identity key
type KeyDID interface {
	GetKey() [32]byte
	GetPurpose() *big.Int
	GetRevokedAt() *big.Int
	GetType() *big.Int
}

// KeyResponse contains the needed fields of the GetKey response
type KeyResponse struct {
	Key       [32]byte
	Purposes  []*big.Int
	RevokedAt *big.Int
}

// Key holds the identity related details
type key struct {
	Key       [32]byte
	Purpose   *big.Int
	RevokedAt *big.Int
	Type      *big.Int
}

//NewKey returns a new key struct
func NewKey(pk [32]byte, purpose *big.Int, keyType *big.Int) KeyDID {
	return &key{pk, purpose, big.NewInt(0), keyType}
}

// GetKey returns the public key
func (idk *key) GetKey() [32]byte {
	return idk.Key
}

// GetPurposes returns the purposes intended for the key
func (idk *key) GetPurpose() *big.Int {
	return idk.Purpose
}

// GetRevokedAt returns the block at which the identity is revoked
func (idk *key) GetRevokedAt() *big.Int {
	return idk.RevokedAt
}

// GetType returns the type of the key
func (idk *key) GetType() *big.Int {
	return idk.Type
}

// String prints the peerID extracted from the key
func (idk *key) String() string {
	peerID, _ := ed25519.PublicKeyToP2PKey(idk.Key)
	return fmt.Sprintf("%s", peerID.Pretty())
}

// IDKey represents a key pair
type IDKey struct {
	PublicKey  []byte
	PrivateKey []byte
}

// IDKeys holds key of an identity
type IDKeys struct {
	ID   []byte
	Keys map[int]IDKey
}

// Config defines methods required for the package identity.
type Config interface {
	GetEthereumDefaultAccountName() string
	GetIdentityID() ([]byte, error)
	GetP2PKeyPair() (pub, priv string)
	GetSigningKeyPair() (pub, priv string)
	GetEthereumContextWaitTimeout() time.Duration
}

// ValidateDIDBytes validates a centrifuge ID given as bytes
func ValidateDIDBytes(givenDID []byte, did DID) error {
	calcdid := NewDIDFromBytes(givenDID)
	if !did.Equal(calcdid) {
		return errors.New("provided bytes doesn't match centID")
	}

	return nil
}
