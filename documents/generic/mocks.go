// +build integration unit testworld

package generic

import (
	"context"
	"testing"

	coredocumentpb "github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/contextutil"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/identity"
	testingidentity "github.com/centrifuge/go-centrifuge/testingutils/identity"
	"github.com/stretchr/testify/assert"
)

func InitGeneric(t *testing.T, did identity.DID, payload documents.CreatePayload) *Generic {
	gen := new(Generic)
	payload.Collaborators.ReadWriteCollaborators = append(payload.Collaborators.ReadWriteCollaborators, did)
	assert.NoError(t, gen.DeriveFromCreatePayload(context.Background(), payload))
	return gen
}

func (b Bootstrapper) TestBootstrap(context map[string]interface{}) error {
	return b.Bootstrap(context)
}

func (Bootstrapper) TestTearDown() error {
	return nil
}

func CreateGenericWithEmbedCDWithPayload(t *testing.T, ctx context.Context, did identity.DID, payload documents.CreatePayload) (*Generic, coredocumentpb.CoreDocument) {
	g := new(Generic)
	payload.Collaborators.ReadWriteCollaborators = append(payload.Collaborators.ReadWriteCollaborators, did)
	err := g.DeriveFromCreatePayload(ctx, payload)
	assert.NoError(t, err)
	g.GetTestCoreDocWithReset()
	sr, err := g.CalculateSigningRoot()
	assert.NoError(t, err)
	// if acc errors out, just skip it
	if ctx == nil {
		ctx = context.Background()
	}
	acc, err := contextutil.Account(ctx)
	if err == nil {
		sig, err := acc.SignMsg(sr)
		assert.NoError(t, err)
		g.AppendSignatures(sig)
	}
	_, err = g.CalculateDocumentRoot()
	assert.NoError(t, err)
	cd, err := g.PackCoreDocument()
	assert.NoError(t, err)

	return g, cd
}

func CreateGenericWithEmbedCD(t *testing.T, ctx context.Context, did identity.DID, collaborators []identity.DID) (*Generic, coredocumentpb.CoreDocument) {
	payload := CreateGenericPayload(t, collaborators)
	return CreateGenericWithEmbedCDWithPayload(t, ctx, did, payload)
}

func CreateGenericPayload(t *testing.T, collaborators []identity.DID) documents.CreatePayload {
	if collaborators == nil {
		collaborators = []identity.DID{testingidentity.GenerateRandomDID()}
	}
	return documents.CreatePayload{
		Scheme: Scheme,
		Collaborators: documents.CollaboratorsAccess{
			ReadWriteCollaborators: collaborators,
		},
	}
}
