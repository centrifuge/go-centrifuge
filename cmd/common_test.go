// +build integration

package cmd

import (
	"context"
	"math/big"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/centrifuge/go-centrifuge/identity/ideth"

	"github.com/centrifuge/go-centrifuge/crypto/ed25519"
	"github.com/centrifuge/go-centrifuge/crypto/secp256k1"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/testingutils"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common"

	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/bootstrap/bootstrappers/testlogging"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/config/configstore"
	"github.com/centrifuge/go-centrifuge/ethereum"

	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/centrifuge/go-centrifuge/storage/leveldb"
	"github.com/centrifuge/go-centrifuge/transactions"
	"github.com/stretchr/testify/assert"
)

var cfg config.Configuration
var ctx = map[string]interface{}{}

func TestMain(m *testing.M) {
	var bootstrappers = []bootstrap.TestBootstrapper{
		&testlogging.TestLoggingBootstrapper{},
		&config.Bootstrapper{},
		&leveldb.Bootstrapper{},
		transactions.Bootstrapper{},
		&queue.Bootstrapper{},
		ethereum.Bootstrapper{},
		&ideth.Bootstrapper{},
		&configstore.Bootstrapper{},
		&queue.Starter{},
	}

	bootstrap.RunTestBootstrappers(bootstrappers, ctx)
	cfg = ctx[bootstrap.BootstrappedConfig].(config.Configuration)
	result := m.Run()
	bootstrap.RunTestTeardown(bootstrappers)
	os.Exit(result)
}

func TestCreateConfig(t *testing.T) {
	// create config
	dataDir := "testconfig"
	keyPath := path.Join(testingutils.GetProjectDir(), "build/scripts/test-dependencies/test-ethereum/migrateAccount.json")
	scAddrs := testingutils.GetSmartContractAddresses()
	err := CreateConfig(dataDir, "http://127.0.0.1:9545", keyPath, "", "russianhill", 8028, 38202, nil, true, "", scAddrs)
	assert.Nil(t, err, "Create Config should be successful")

	// config exists
	cfg := config.LoadConfiguration(path.Join(dataDir, "config.yaml"))
	client := ctx[ethereum.BootstrappedEthereumClient].(ethereum.Client)

	// contract exists
	id, err := cfg.GetIdentityID()
	accountId := identity.NewDID(common.BytesToAddress(id))

	assert.Nil(t, err, "did should exists")
	contractCode, err := client.GetEthClient().CodeAt(context.Background(), common.BytesToAddress(id), nil)
	assert.Nil(t, err, "should be successful to get the contract code")
	assert.Equal(t, true, len(contractCode) > 3000, "current contract code should be arround 3378 bytes")

	// Keys exists
	// type KeyPurposeEthMsgAuth
	idSrv := ctx[ideth.BootstrappedDIDService].(identity.ServiceDID)
	pk, _, err := secp256k1.GetEthAuthKey(cfg.GetEthAuthKeyPair())
	assert.Nil(t, err)
	address32Bytes := utils.AddressTo32Bytes(common.HexToAddress(secp256k1.GetAddress(pk)))
	assert.Nil(t, err)
	response, err := idSrv.GetKey(accountId, address32Bytes)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, big.NewInt(identity.KeyPurposeEthMsgAuth), response.Purposes[0], "purpose should be ETHMsgAuth")

	// type KeyPurposeP2P
	pk, _, err = ed25519.GetSigningKeyPair(cfg.GetP2PKeyPair())
	assert.Nil(t, err)
	pk32, err := utils.SliceToByte32(pk)
	assert.Nil(t, err)
	response, _ = idSrv.GetKey(accountId, pk32)
	assert.NotNil(t, response)
	assert.Equal(t, big.NewInt(identity.KeyPurposeP2P), response.Purposes[0], "purpose should be P2P")

	// type KeyPurposeSigning
	pk, _, err = ed25519.GetSigningKeyPair(cfg.GetSigningKeyPair())
	assert.Nil(t, err)
	pk32, err = utils.SliceToByte32(pk)
	assert.Nil(t, err)
	response, _ = idSrv.GetKey(accountId, pk32)
	assert.NotNil(t, response)
	assert.Equal(t, big.NewInt(identity.KeyPurposeSigning), response.Purposes[0], "purpose should be Signing")

	err = exec.Command("rm", "-rf", dataDir).Run()
	assert.Nil(t, err, "removing testconfig folder should be successful")

}
