package entity

import (
	"github.com/centrifuge/centrifuge-protobufs/documenttypes"
	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/documents/entityrelationship"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/jobs"
	"github.com/centrifuge/go-centrifuge/queue"
)

const (
	// BootstrappedEntityHandler maps to grpc handler for entities
	BootstrappedEntityHandler string = "BootstrappedEntityHandler"

	// BootstrappedEntityService maps to the service for entities
	BootstrappedEntityService string = "BootstrappedEntityService"
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

	repo, ok := ctx[documents.BootstrappedDocumentRepository].(documents.Repository)
	if !ok {
		return errors.New("document db repository not initialised")
	}
	repo.Register(&Entity{})

	queueSrv, ok := ctx[bootstrap.BootstrappedQueueServer].(*queue.Server)
	if !ok {
		return errors.New("queue server not initialised")
	}

	jobManager, ok := ctx[jobs.BootstrappedService].(jobs.Manager)
	if !ok {
		return errors.New("transaction service not initialised")
	}

	cfgSrv, ok := ctx[config.BootstrappedConfigStorage].(config.Service)
	if !ok {
		return errors.New("config service not initialised")
	}

	factory, ok := ctx[identity.BootstrappedDIDFactory].(identity.Factory)
	if !ok {
		return errors.New("identity factory not initialised")
	}

	erService, ok := ctx[entityrelationship.BootstrappedEntityRelationshipService].(entityrelationship.Service)
	if !ok {
		return errors.New("entity relation service not initialised")
	}

	processor, ok := ctx[documents.BootstrappedAnchorProcessor].(documents.DocumentRequestProcessor)
	if !ok {
		return errors.New("processor not initialised")
	}

	anchorRepo, ok := ctx[anchors.BootstrappedAnchorRepo].(anchors.AnchorRepository)
	if !ok {
		return errors.New("anchor repository not initialised")
	}

	didService, ok := ctx[identity.BootstrappedDIDService].(identity.Service)
	if !ok {
		return errors.New("identity service not initialized")
	}

	// register service
	srv := DefaultService(
		docSrv,
		repo,
		queueSrv, jobManager, factory, erService, didService, anchorRepo, processor, func() documents.ValidatorGroup {
			return documents.PostAnchoredValidator(didService, anchorRepo)
		})

	err := registry.Register(documenttypes.EntityDataTypeUrl, srv)
	if err != nil {
		return errors.New("failed to register entity service: %v", err)
	}

	err = registry.Register(scheme, srv)
	if err != nil {
		return errors.New("failed to register entity service: %v", err)
	}

	ctx[BootstrappedEntityService] = srv
	ctx[BootstrappedEntityHandler] = GRPCHandler(cfgSrv, srv)

	return nil
}
