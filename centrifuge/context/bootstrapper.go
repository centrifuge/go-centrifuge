package context

import (
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/anchoring"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/bootstrapper"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/config"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/coredocument"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/ethereum"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/identity"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/invoice"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/purchaseorder"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/queue"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/storage"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/version"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("context")

type MainBootstrapper struct {
	Bootstrappers []bootstrapper.Bootstrapper
}

func (m *MainBootstrapper) PopulateDefaultBootstrappers() {
	m.Bootstrappers = []bootstrapper.Bootstrapper{
		&version.Bootstrapper{},
		&config.Bootstrapper{},
		&storage.Bootstrapper{},
		&coredocument.Bootstrapper{},
		&invoice.Bootstrapper{},
		&purchaseorder.Bootstrapper{},
		&ethereum.Bootstrapper{},
		&anchoring.Bootstrapper{},
		&identity.Bootstrapper{},
		&queue.Bootstrapper{},
	}
}

func (m *MainBootstrapper) Bootstrap(context map[string]interface{}) error {
	for _, b := range m.Bootstrappers {
		err := b.Bootstrap(context)
		if err != nil {
			log.Error("Error encountered while bootstrapping", err)
			return err
		}
	}
	return nil
}
