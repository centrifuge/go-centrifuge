package entityrelationship

import (
	"github.com/centrifuge/centrifuge-protobufs/documenttypes"
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/centrifuge/go-centrifuge/storage"
	"github.com/centrifuge/go-centrifuge/transactions"
)


const (
	// BootstrappedEntityHandler maps to grpc handler for entities
	BootstrappedEntityrelationService string = "BootstrappedEntityrelationService"
)

// Bootstrapper implements bootstrap.Bootstrapper.
type Bootstrapper struct{}

// Bootstrap sets the required storage and registers
func (Bootstrapper) Bootstrap(ctx map[string]interface{}) error {
	registry, ok := ctx[documents.BootstrappedRegistry].(*documents.ServiceRegistry)
	if !ok {
		return errors.New("service registry not initialised")
	}

	docSrv, ok := ctx[documents.BootstrappedDocumentService].(documents.Service)
	if !ok {
		return errors.New("document service not initialised")
	}

	repo, ok := ctx[documents.BootstrappedDocumentRepository].(storage.Repository)
	if !ok {
		return errors.New("document db repository not initialised")
	}


	entityRepo := newDBRepository(repo)

	repo.Register(&EntityRelationship{})

	queueSrv, ok := ctx[bootstrap.BootstrappedQueueServer].(*queue.Server)
	if !ok {
		return errors.New("queue server not initialised")
	}

	txManager, ok := ctx[transactions.BootstrappedService].(transactions.Manager)
	if !ok {
		return errors.New("transaction service not initialised")
	}

	factory, ok := ctx[identity.BootstrappedDIDFactory].(identity.Factory)
	if !ok {
		return errors.New("identity factory not initialised")
	}

	// register service
	srv := DefaultService(
		docSrv,
		entityRepo,
		queueSrv, txManager, factory)

	err := registry.Register(documenttypes.EntityDataTypeUrl, srv)
	if err != nil {
		return errors.New("failed to register entity service: %v", err)
	}

	ctx[BootstrappedEntityrelationService] = srv

	return nil
}
