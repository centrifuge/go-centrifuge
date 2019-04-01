// +build cmd

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/centrifuge/go-centrifuge/testingutils"
	"github.com/centrifuge/go-centrifuge/version"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	o, err := exec.Command(testingutils.GetBinaryPath(), "version").Output()
	assert.NoError(t, err)

	assert.Contains(t, string(o), version.CentrifugeNodeVersion)
}

func TestCreateConfigCmd(t *testing.T) {
	dataDir := path.Join(os.Getenv("HOME"), "datadir")
	scAddrs := testingutils.GetSmartContractAddresses()
	keyPath := path.Join(testingutils.GetProjectDir(), "build/scripts/test-dependencies/test-ethereum/migrateAccount.json")
	cmd := exec.Command(testingutils.GetBinaryPath(), "createconfig", "-n", "testing", "-t", dataDir, "-z", keyPath)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("CENT_NETWORKS_TESTING_CONTRACTADDRESSES_IDENTITYFACTORY=%s", scAddrs.IdentityFactoryAddr))
	cmd.Env = append(cmd.Env, fmt.Sprintf("CENT_NETWORKS_TESTING_CONTRACTADDRESSES_ANCHORREPOSITORY=%s", scAddrs.AnchorRepositoryAddr))
	cmd.Env = append(cmd.Env, fmt.Sprintf("CENT_NETWORKS_TESTING_CONTRACTADDRESSES_INVOICEUNPAID=%s", scAddrs.InvoiceUnpaidAddr))
	o, err := cmd.Output()
	assert.NoError(t, err)

	fmt.Printf("Output: %s\n", o)
}
