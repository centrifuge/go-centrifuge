package identity

import (
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/CentrifugeInc/go-centrifuge/centrifuge/config"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/errors"
	centrifugeED25519 "github.com/CentrifugeInc/go-centrifuge/centrifuge/keytools/ed25519"
)

const (
	CentIdByteLength     = 6
	ActionCreate         = "create"
	ActionAddKey         = "addkey"
	KeyPurposeP2p        = 1
	KeyPurposeSigning    = 2
	KeyPurposeEthMsgAuth = 3
)

type Identity interface {
	String() string
	GetCentrifugeID() []byte
	CentrifugeIDString() string
	CentrifugeIDBytes() [CentIdByteLength]byte
	CentrifugeIDBigInt() *big.Int
	SetCentrifugeID(b []byte) error
	GetCurrentP2PKey() (ret string, err error)
	GetLastKeyForPurpose(keyPurpose int) (key []byte, err error)
	AddKeyToIdentity(keyPurpose int, key []byte) (confirmations chan *WatchIdentity, err error)
	CheckIdentityExists() (exists bool, err error)
	FetchKey(key []byte) (IdentityKey, error)
}

type IdentityKey interface {
	GetKey() [32]byte
	GetPurposes() []*big.Int
	GetRevokedAt() *big.Int
}

type WatchIdentity struct {
	Identity Identity
	Error    error
}

// Service is used to fetch identities
type Service interface {
	LookupIdentityForID(centrifugeId []byte) (id Identity, err error)
	CreateIdentity(centrifugeId []byte) (id Identity, confirmations chan *WatchIdentity, err error)
	CheckIdentityExists(centrifugeId []byte) (exists bool, err error)
}

// CentrifugeIdStringToSlice takes a string and decodes it using base64 to convert it into a slice
// of length 32.
func CentrifugeIdStringToSlice(s string) (id []byte, err error) {
	id, err = base64.StdEncoding.DecodeString(s)
	if err != nil {
		return []byte{}, err
	}
	if len(id) != CentIdByteLength {
		return []byte{}, fmt.Errorf("CentrifugeId has invalid length [%d]", len(id))
	}
	return id, nil
}

// FromConfig reads the keys and ID from the config and returns a the Identity config
func FromConfig() (identityConfig *config.IdentityConfig, err error) {
	pk, pvk := centrifugeED25519.GetSigningKeyPairFromConfig()
	decodedId, err := base64.StdEncoding.DecodeString(string(config.Config.GetIdentityId()))
	if err != nil {
		return nil, err
	}

	identityConfig = &config.IdentityConfig{
		IdentityId: decodedId,
		PublicKey:  pk,
		PrivateKey: pvk,
	}
	return
}

// GetClientP2PURL returns the p2p url associated with the centID
func GetClientP2PURL(idService Service, centID []byte) (url string, err error) {
	target, err := idService.LookupIdentityForID(centID)
	if err != nil {
		return url, errors.Wrap(err, "error fetching receiver identity")
	}

	p2pKey, err := target.GetCurrentP2PKey()
	if err != nil {
		return url, errors.Wrap(err, "error fetching p2p key")
	}

	return fmt.Sprintf("/ipfs/%s", p2pKey), nil
}

// GetClientsP2PURLs returns p2p urls associated with each centIDs
// will error out at first failure
func GetClientsP2PURLs(idService Service, centIDs [][]byte) ([]string, error) {
	var p2pURLs []string
	for _, id := range centIDs {
		url, err := GetClientP2PURL(idService, id)
		if err != nil {
			return nil, err
		}

		p2pURLs = append(p2pURLs, url)
	}

	return p2pURLs, nil
}
