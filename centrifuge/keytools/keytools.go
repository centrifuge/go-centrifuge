package keytools

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("keytools")

const (
	PUBLIC_KEY  = "PUBLIC KEY"
	PRIVATE_KEY = "PRIVATE KEY"
)

const (
	CURVE_ED25519   string = "ed25519"
	CURVE_SECP256K1 string = "secp256k1"
)

const MAX_MSG_LEN = 32

func writeKeyToPemFile(fileName string, keyType string, key []byte) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	block := &pem.Block{
		Type:  keyType,
		Bytes: key,
	}
	if err := pem.Encode(f, block); err != nil {
		return err
	}
	f.Close()
	return nil
}

func readKeyFromPemFile(fileName, keyType string) (key []byte, err error) {
	pemData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []byte{}, err
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		return []byte{}, fmt.Errorf("File [%s] is not a valid pem file", fileName)
	}
	if block.Type != keyType {
		return []byte{}, fmt.Errorf("Key type mismatch got [%s] but expected [%s]", block.Type, keyType)
	}

	return block.Bytes, nil
}


