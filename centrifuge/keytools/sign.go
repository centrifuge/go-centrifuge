package keytools

import (
	"fmt"
	"strings"

	"github.com/CentrifugeInc/go-centrifuge/centrifuge/keytools/io"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/keytools/secp256k1"
)

func SignMessage(privateKeyPath, message, curveType string, ethereumSign bool) []byte {

	privateKey, err := io.ReadKeyFromPemFile(privateKeyPath, PrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	curveType = strings.ToLower(curveType)

	var msg []byte
	if ethereumSign == true {
		msg = []byte(message)
	} else {
		if len(message) > MaxMsgLen {
			log.Fatal("max message len is 32 bytes current len:", len(message))
		}

		msg = make([]byte, MaxMsgLen)
		copy(msg, message)
	}

	switch curveType {

	case CurveSecp256K1:
		if ethereumSign {
			return secp256k1.SignEthereum(msg, privateKey)
		}
		return secp256k1.Sign(msg, privateKey)

	case CurveEd25519:
		fmt.Println("curve ed25519 not support yet")
		return []byte("")

	default:
		if ethereumSign {
			return secp256k1.SignEthereum(msg, privateKey)
		}
		return secp256k1.Sign(msg, privateKey)

	}

}
