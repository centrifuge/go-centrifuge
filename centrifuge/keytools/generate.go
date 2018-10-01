package keytools

import (
	"strings"

	"github.com/centrifuge/go-centrifuge/centrifuge/keytools/ed25519keys"
	"github.com/centrifuge/go-centrifuge/centrifuge/keytools/secp256k1"
	"github.com/centrifuge/go-centrifuge/centrifuge/utils"
)

func GenerateSigningKeyPair(publicFileName, privateFileName, curveType string) {
	var publicKey, privateKey []byte
	switch strings.ToLower(curveType) {
	case CurveSecp256K1:
		publicKey, privateKey = secp256k1.GenerateSigningKeyPair()
	case CurveEd25519:
		publicKey, privateKey = ed25519keys.GenerateSigningKeyPair()
	default:
		publicKey, privateKey = ed25519keys.GenerateSigningKeyPair()
	}

	utils.WriteKeyToPemFile(privateFileName, utils.PrivateKey, privateKey)
	utils.WriteKeyToPemFile(publicFileName, utils.PublicKey, publicKey)
}
