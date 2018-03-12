// +build ethereum

package ethereum_test

import (
	"testing"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/ethereum"
	"os"
	"github.com/spf13/viper"
	"github.com/magiconair/properties/assert"
)

func TestMain(m *testing.M) {
	//for now set up the env vars manually in integration test
	//TODO move to generalized config once it is available
	viper.BindEnv("ethereum.gethSocket", "CENT_ETHEREUM_GETH_SOCKET")

	result := m.Run()
	os.Exit(result)
}

func TestGetConnection_returnsSameConnection(t *testing.T) {
	//TODO this will currently fail if concurrency is at play - e.g. running with 3 go-routines the test will fail
	howMany := 5
	confChannel := make(chan ethereum.EthereumClient, howMany)
	for ix := 0; ix < howMany; ix++ {
		go func() {
			confChannel <- ethereum.GetConnection()
		}()
	}
	for ix := 0; ix < howMany; ix++ {
		multiThreadCreatedCon := <-confChannel
		assert.Equal(t, multiThreadCreatedCon , ethereum.GetConnection(), "Should only return a single ethereum client")
	}
}
