package nft

import (
	"errors"

	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/ethereum"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/ethereum/go-ethereum/common"
)

type Bootstrapper struct {
}

// Bootstrap initializes the payment obligation contract
func (*Bootstrapper) Bootstrap(context map[string]interface{}) error {
	if _, ok := context[config.BootstrappedConfig]; !ok {
		return errors.New("config hasn't been initialized")
	}
	cfg := context[config.BootstrappedConfig].(*config.Configuration)

	if _, ok := context[ethereum.BootstrappedEthereumClient]; !ok {
		return errors.New("ethereum client hasn't been initialized")
	}

	contract, err := getPaymentObligationContract(cfg.GetContractAddress("paymentObligation"))
	if err != nil {
		return err
	}

	setPaymentObligation(NewEthereumPaymentObligation(contract, identity.IDService, ethereum.GetClient(), cfg, setupMintListener))
	return queue.InstallQueuedTask(context,
		newMintingConfirmationTask(contract, ethereum.DefaultWaitForTransactionMiningContext))
}

func getPaymentObligationContract(obligationAddress common.Address) (*EthereumPaymentObligationContract, error) {
	client := ethereum.GetClient()
	return NewEthereumPaymentObligationContract(obligationAddress, client.GetEthClient())
}
