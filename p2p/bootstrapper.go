package p2p

import (
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/config/configstore"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/p2p/receiver"
)

// Bootstrapper implements Bootstrapper with p2p details
type Bootstrapper struct{}

// Bootstrap initiates p2p server and client into context
func (b Bootstrapper) Bootstrap(ctx map[string]interface{}) error {
	cfg, err := configstore.RetrieveConfig(true, ctx)
	if err != nil {
		return err
	}

	registry, ok := ctx[documents.BootstrappedRegistry].(*documents.ServiceRegistry)
	if !ok {
		return errors.New("registry not initialised")
	}

	srv := &peer{config: cfg, handlerCreator: func() *receiver.Handler {
		return receiver.New(cfg, registry)
	}}
	ctx[bootstrap.BootstrappedP2PServer] = srv
	ctx[bootstrap.BootstrappedP2PClient] = srv
	return nil
}
