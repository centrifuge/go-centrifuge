package documents

import (
	"github.com/centrifuge/go-centrifuge/storage"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	// BootstrappedRegistry is the key to ServiceRegistry in Bootstrap context
	BootstrappedRegistry = "BootstrappedRegistry"
	// BootstrappedDocumentRepository is the key to the database repository of documents
	BootstrappedDocumentRepository = "BootstrappedDocumentRepository"
)

// Bootstrapper implements bootstrap.Bootstrapper.
type Bootstrapper struct{}

// Bootstrap sets the required storage and registers
func (Bootstrapper) Bootstrap(ctx map[string]interface{}) error {
	ctx[BootstrappedRegistry] = NewServiceRegistry()
	ldb, ok := ctx[storage.BootstrappedLevelDB].(*leveldb.DB)
	if !ok {
		return ErrDocumentBootstrap
	}
	repo := NewLevelDBRepository(ldb)
	ctx[BootstrappedDocumentRepository] = repo
	return nil
}
