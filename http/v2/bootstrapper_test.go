// +build unit

package v2

import (
	"testing"

	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/documents/entity"
	"github.com/centrifuge/go-centrifuge/documents/entityrelationship"
	"github.com/centrifuge/go-centrifuge/jobs"
	"github.com/centrifuge/go-centrifuge/oracle"
	"github.com/centrifuge/go-centrifuge/pending"
	testingnfts "github.com/centrifuge/go-centrifuge/testingutils/nfts"
	"github.com/stretchr/testify/assert"
)

func TestBootstrapper_Bootstrap(t *testing.T) {
	ctx := make(map[string]interface{})
	b := Bootstrapper{}

	ctx[pending.BootstrappedPendingDocumentService] = new(pending.MockService)
	ctx[bootstrap.BootstrappedNFTService] = new(testingnfts.MockNFTService)
	ctx[oracle.BootstrappedOracleService] = new(oracle.MockService)
	ctx[config.BootstrappedConfigStorage] = new(config.MockService)
	ctx[jobs.BootstrappedDispatcher] = new(jobs.MockDispatcher)
	ctx[entity.BootstrappedEntityService] = new(entity.MockService)
	ctx[entityrelationship.BootstrappedEntityRelationshipService] = new(entity.MockEntityRelationService)
	ctx[documents.BootstrappedDocumentService] = new(documents.MockService)
	assert.NoError(t, b.Bootstrap(ctx))
	assert.NotNil(t, ctx[BootstrappedService])
}
