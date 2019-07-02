package coreapi

import (
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/go-chi/chi"
)

const (
	documentIDParam      = "document_id"
	versionIDParam       = "version_id"
	jobIDParam           = "job_id"
	tokenIDParam         = "token_id"
	registryAddressParam = "registry_address"
	accountIDParam       = "account_id"
)

// Register registers the core apis to the router.
func Register(ctx map[string]interface{}, r chi.Router) {
	coreAPISrv := ctx[BootstrappedCoreAPIService].(Service)
	tokenRegistry := ctx[bootstrap.BootstrappedInvoiceUnpaid].(documents.TokenRegistry)
	h := handler{
		srv:           coreAPISrv,
		tokenRegistry: tokenRegistry,
	}

	r.Post("/documents", h.CreateDocument)
	r.Put("/documents/{"+documentIDParam+"}", h.UpdateDocument)
	r.Get("/documents/{"+documentIDParam+"}", h.GetDocument)
	r.Get("/documents/{"+documentIDParam+"}/versions/{"+versionIDParam+"}", h.GetDocumentVersion)
	r.Post("/documents/{"+documentIDParam+"}/proofs", h.GenerateProofs)
	r.Post("/documents/{"+documentIDParam+"}/versions/{"+versionIDParam+"}/proofs", h.GenerateProofsForVersion)
	r.Get("/jobs/{"+jobIDParam+"}", h.GetJobStatus)
	r.Post("/nfts/registries/{"+registryAddressParam+"}/mint", h.MintNFT)
	r.Post("/nfts/registries/{"+registryAddressParam+"}/tokens/{"+tokenIDParam+"}/transfer", h.TransferNFT)
	r.Get("/nfts/registries/{"+registryAddressParam+"}/tokens/{"+tokenIDParam+"}/owner", h.OwnerOfNFT)
	r.Post("/accounts/{"+accountIDParam+"}/sign", h.SignPayload)
	r.Get("/accounts/{"+accountIDParam+"}", h.GetAccount)
}
