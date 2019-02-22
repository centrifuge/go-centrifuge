package crypto

import (
	"github.com/centrifuge/go-centrifuge/crypto/ed25519"
	"github.com/centrifuge/go-centrifuge/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// VerifyMessage verifies message using the public key as per the curve type.
// if ethereumVerify is true, ethereum specific verification is done
func VerifyMessage(publicKey, message []byte, signature []byte, curveType string, ethereumVerify bool) bool {
	signatureBytes := make([]byte, len(signature))
	copy(signatureBytes, signature)

	switch curveType {
	case CurveSecp256K1:
		if ethereumVerify {
			address := secp256k1.GetAddress(publicKey)
			return secp256k1.VerifySignatureWithAddress(address, hexutil.Encode(signatureBytes), message)
		}

		return secp256k1.VerifySignature(publicKey, message, signatureBytes)
	case CurveEd25519:
		return ed25519.VerifySignature(publicKey, message, signature)
	default:
		return false
	}
}
