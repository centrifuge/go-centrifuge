package node

import (
	"context"
	"errors"
	"os"
	"os/signal"

	"github.com/centrifuge/go-centrifuge/centrifuge/api"
	"github.com/centrifuge/go-centrifuge/centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/centrifuge/config"
	"github.com/centrifuge/go-centrifuge/centrifuge/keytools/ed25519keys"
	"github.com/centrifuge/go-centrifuge/centrifuge/p2p"
)

type Bootstrapper struct {
}

func (*Bootstrapper) Bootstrap(c map[string]interface{}) error {
	if _, ok := c[bootstrap.BootstrappedConfig]; ok {
		services := defaultServerList()
		n := NewNode(services)
		startUpErr := make(chan error)
		// may be we can pass a context that exists in c here
		ctx, canc := context.WithCancel(context.Background())
		go n.Start(ctx, startUpErr)
		controlC := make(chan os.Signal, 1)
		signal.Notify(controlC, os.Interrupt)
		for {
			select {
			case err := <-startUpErr:
				panic(err)
			case sig := <-controlC:
				log.Info("Node shutting down because of ", sig)
				canc()
				<-startUpErr
				return nil
			}
		}
		return nil
	}
	return errors.New("could not initialize node")
}

func defaultServerList() []Server {
	publicKey, privateKey := ed25519keys.GetSigningKeyPairFromConfig()
	services := []Server{
		api.NewCentAPIServer(
			config.Config.GetServerAddress(),
			config.Config.GetServerPort(),
			config.Config.GetNetworkString(),
		),
		p2p.NewCentP2PServer(
			config.Config.GetP2PPort(),
			config.Config.GetBootstrapPeers(),
			publicKey, privateKey,
		),
	}
	return services
}
