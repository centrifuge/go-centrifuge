package identity

import (
	"github.com/spf13/viper"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/keytools"
	"log"
	"fmt"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"errors"
)

type IdentityKey struct {
	Key [32]byte
}

func (idk *IdentityKey) String() string {
	peerdId, _ := keytools.PublicKeyToP2PKey(idk.Key)
	return fmt.Sprintf("%s", peerdId.Pretty())
}

type Identity struct {
	CentrifugeId string
	Keys map[int][]IdentityKey
}

func (id *Identity) String() string {
	joinedKeys := ""
	for k, v := range id.Keys {
		for i, _ := range v {
			joinedKeys += fmt.Sprintf("[%v]%s ", k, v[i].String())
		}
	}
	return fmt.Sprintf("CentrifugeId [%s], Keys [%s]", id.CentrifugeId, joinedKeys)
}

func (id *Identity) GetLastB58Key(keyType int) (ret string, err error) {
	if len(id.Keys[keyType]) == 0 {
		return
	}
	switch keyType {
	case 0:
		log.Printf("Error not authorized type")
	case 1:
		p2pId, err1 := keytools.PublicKeyToP2PKey(id.Keys[keyType][len(id.Keys[keyType])-1].Key)
		if err1 != nil {
			err = err1
			return
		}
		ret = p2pId.Pretty()
	default:
		log.Printf("keyType Not found")
	}
	return
}

func CheckIdentityExists(centrifugeId string) (exists bool, err error) {
	if (viper.GetBool("identity.ethereum.enabled")) {
		idContract, err := doFindIdentity(centrifugeId)
		if err != nil {
			return false, err
		}
		if idContract != nil {
			opts := ethereum.GetGethCallOpts()
			centId, err := idContract.CentrifugeId(opts)
			if err == bind.ErrNoCode { //no contract in specified address, meaning Identity was not created
				log.Printf("Identity contract does not exist!")
				err = nil
			} else if len(centId) != 0 {
				log.Printf("Identity exists!")
				exists = true
			}
		}
	} else {
		err = errors.New("Ethereum Identity not enabled")
	}
	return
}

func ResolveP2PIdentityForId(centrifugeId string, keyType int) (id Identity, err error) {
	if (viper.GetBool("identity.ethereum.enabled")) {
		id, err = doResolveIdentityForKeyType(centrifugeId, keyType)
	} else {
		err = errors.New("Ethereum Identity not enabled")
	}
	return
}

func CreateIdentity(identity Identity, confirmations chan<- *Identity) (err error) {
	if (viper.GetBool("identity.ethereum.enabled")) {
		err = doCreateIdentity(identity, confirmations)
	} else {
		err = errors.New("Ethereum Identity not enabled")
	}
	return
}

func AddKeyToIdentity(identity Identity, keyType int, confirmations chan<- *Identity) (err error) {
	if (viper.GetBool("identity.ethereum.enabled")) {
		err = doAddKeyToIdentity(identity, keyType, confirmations)
	} else {
		err = errors.New("Ethereum Identity not enabled")
	}
	return
}