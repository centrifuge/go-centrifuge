package invoice

import (
	"context"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/contextutil"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/jobs"
	clientinvoicepb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/invoice"
	"github.com/centrifuge/go-centrifuge/queue"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Service defines specific functions for invoice
type Service interface {
	documents.Service

	// DeriverFromPayload derives Invoice from clientPayload
	DeriveFromCreatePayload(ctx context.Context, payload *clientinvoicepb.InvoiceCreatePayload) (documents.Model, error)

	// DeriveFromUpdatePayload derives invoice model from update payload
	DeriveFromUpdatePayload(ctx context.Context, payload *clientinvoicepb.InvoiceUpdatePayload) (documents.Model, error)

	// DeriveInvoiceData returns the invoice data as client data
	DeriveInvoiceData(inv documents.Model) (*clientinvoicepb.InvoiceData, error)

	// DeriveInvoiceResponse returns the invoice model in our standard client format
	DeriveInvoiceResponse(inv documents.Model) (*clientinvoicepb.InvoiceResponse, error)
}

// service implements Service and handles all invoice related persistence and validations
// service always returns errors of type `errors.Error` or `errors.TypedError`
type service struct {
	documents.Service
	repo           documents.Repository
	queueSrv       queue.TaskQueuer
	jobManager     jobs.Manager
	tokenRegFinder func() documents.TokenRegistry
	anchorRepo     anchors.AnchorRepository
}

// DefaultService returns the default implementation of the service.
func DefaultService(
	srv documents.Service,
	repo documents.Repository,
	queueSrv queue.TaskQueuer,
	jobManager jobs.Manager,
	tokenRegFinder func() documents.TokenRegistry,
	anchorRepo anchors.AnchorRepository,
) Service {
	return service{
		repo:           repo,
		queueSrv:       queueSrv,
		jobManager:     jobManager,
		Service:        srv,
		tokenRegFinder: tokenRegFinder,
		anchorRepo:     anchorRepo,
	}
}

// DeriveFromCoreDocument takes a core document model and returns an invoice
func (s service) DeriveFromCoreDocument(cd coredocumentpb.CoreDocument) (documents.Model, error) {
	inv := new(Invoice)
	err := inv.UnpackCoreDocument(cd)
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDocumentUnPackingCoreDocument, err)
	}

	return inv, nil
}

// UnpackFromCreatePayload initializes the model with parameters provided from the rest-api call
func (s service) DeriveFromCreatePayload(ctx context.Context, payload *clientinvoicepb.InvoiceCreatePayload) (documents.Model, error) {
	if payload == nil || payload.Data == nil {
		return nil, documents.ErrDocumentNil
	}

	did, err := contextutil.AccountDID(ctx)
	if err != nil {
		return nil, documents.ErrDocumentConfigAccountID
	}

	invoiceModel := new(Invoice)
	err = invoiceModel.InitInvoiceInput(payload, did)
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDocumentInvalid, err)
	}

	return invoiceModel, nil
}

// validateAndPersist validates the document, calculates the data root, and persists to DB
func (s service) validateAndPersist(ctx context.Context, old, new documents.Model, validator documents.Validator) (documents.Model, error) {
	selfDID, err := contextutil.AccountDID(ctx)
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDocumentConfigAccountID, err)
	}

	inv, ok := new.(*Invoice)
	if !ok {
		return nil, errors.NewTypedError(documents.ErrDocumentInvalidType, errors.New("unknown document type: %T", new))
	}

	// validate the invoice
	err = validator.Validate(old, inv)
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDocumentInvalid, err)
	}

	// we use CurrentVersion as the id since that will be unique across multiple versions of the same document
	err = s.repo.Create(selfDID[:], inv.CurrentVersion(), inv)
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDocumentPersistence, err)
	}

	return inv, nil
}

// Create takes and invoice model and does required validation checks, tries to persist to DB
func (s service) Create(ctx context.Context, inv documents.Model) (documents.Model, jobs.JobID, chan bool, error) {
	selfDID, err := contextutil.AccountDID(ctx)
	if err != nil {
		return nil, jobs.NilJobID(), nil, errors.NewTypedError(documents.ErrDocumentConfigAccountID, err)
	}

	inv, err = s.validateAndPersist(ctx, nil, inv, CreateValidator())
	if err != nil {
		return nil, jobs.NilJobID(), nil, err
	}

	jobID := contextutil.Job(ctx)
	jobID, done, err := documents.CreateAnchorJob(ctx, s.jobManager, s.queueSrv, selfDID, jobID, inv.CurrentVersion())
	if err != nil {
		return nil, jobs.NilJobID(), nil, err
	}
	return inv, jobID, done, nil
}

// Update finds the old document, validates the new version and persists the updated document
func (s service) Update(ctx context.Context, new documents.Model) (documents.Model, jobs.JobID, chan bool, error) {
	selfDID, err := contextutil.AccountDID(ctx)
	if err != nil {
		return nil, jobs.NilJobID(), nil, errors.NewTypedError(documents.ErrDocumentConfigAccountID, err)
	}

	old, err := s.GetCurrentVersion(ctx, new.ID())
	if err != nil {
		return nil, jobs.NilJobID(), nil, errors.NewTypedError(documents.ErrDocumentNotFound, err)
	}

	new, err = s.validateAndPersist(ctx, old, new, UpdateValidator(s.anchorRepo))
	if err != nil {
		return nil, jobs.NilJobID(), nil, err
	}

	jobID := contextutil.Job(ctx)
	jobID, done, err := documents.CreateAnchorJob(ctx, s.jobManager, s.queueSrv, selfDID, jobID, new.CurrentVersion())
	if err != nil {
		return nil, jobs.NilJobID(), nil, err
	}
	return new, jobID, done, nil
}

// DeriveInvoiceResponse returns create response from invoice model
func (s service) DeriveInvoiceResponse(model documents.Model) (*clientinvoicepb.InvoiceResponse, error) {
	data, err := s.DeriveInvoiceData(model)
	if err != nil {
		return nil, err
	}

	h, err := documents.DeriveResponseHeader(s.tokenRegFinder(), model)
	if err != nil {
		return nil, errors.New("failed to derive response: %v", err)
	}

	attrs, err := documents.ToClientAttributes(model.GetAttributes())
	if err != nil {
		return nil, err
	}

	return &clientinvoicepb.InvoiceResponse{
		Header:     h,
		Data:       data,
		Attributes: attrs,
	}, nil

}

// DeriveInvoiceData returns create response from invoice model
func (s service) DeriveInvoiceData(doc documents.Model) (*clientinvoicepb.InvoiceData, error) {
	inv, ok := doc.(*Invoice)
	if !ok {
		return nil, documents.ErrDocumentInvalidType
	}

	return inv.getClientData()
}

// DeriveFromUpdatePayload returns a new version of the old invoice identified by identifier in payload
func (s service) DeriveFromUpdatePayload(ctx context.Context, payload *clientinvoicepb.InvoiceUpdatePayload) (documents.Model, error) {
	if payload == nil || payload.Data == nil {
		return nil, documents.ErrDocumentNil
	}

	// get latest old version of the document
	id, err := hexutil.Decode(payload.DocumentId)
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDocumentIdentifier, errors.New("failed to decode identifier: %v", err))
	}

	old, err := s.GetCurrentVersion(ctx, id)
	if err != nil {
		return nil, err
	}

	cs, err := documents.FromClientCollaboratorAccess(payload.ReadAccess, payload.WriteAccess)
	if err != nil {
		return nil, err
	}

	attrs, err := documents.FromClientAttributes(payload.Attributes)
	if err != nil {
		return nil, err
	}

	inv := new(Invoice)
	if err := inv.PrepareNewVersion(old, payload.Data, cs, attrs); err != nil {
		return nil, errors.NewTypedError(documents.ErrDocumentPrepareCoreDocument, errors.New("failed to load invoice from data: %v", err))
	}

	return inv, nil
}

// CreateModel creates invoice from the payload, validates, persists, and returns the invoice.
func (s service) CreateModel(ctx context.Context, payload documents.CreatePayload) (documents.Model, jobs.JobID, error) {
	if payload.Data == nil {
		return nil, jobs.NilJobID(), documents.ErrDocumentNil
	}

	did, err := contextutil.AccountDID(ctx)
	if err != nil {
		return nil, jobs.NilJobID(), documents.ErrDocumentConfigAccountID
	}

	inv := new(Invoice)
	if err := inv.unpackFromCreatePayload(did, payload); err != nil {
		return nil, jobs.NilJobID(), errors.NewTypedError(documents.ErrDocumentInvalid, err)
	}

	// validate invoice
	err = CreateValidator().Validate(nil, inv)
	if err != nil {
		return nil, jobs.NilJobID(), errors.NewTypedError(documents.ErrDocumentInvalid, err)
	}

	// we use CurrentVersion as the id since that will be unique across multiple versions of the same document
	err = s.repo.Create(did[:], inv.CurrentVersion(), inv)
	if err != nil {
		return nil, jobs.NilJobID(), errors.NewTypedError(documents.ErrDocumentPersistence, err)
	}

	jobID := contextutil.Job(ctx)
	jobID, _, err = documents.CreateAnchorJob(ctx, s.jobManager, s.queueSrv, did, jobID, inv.CurrentVersion())
	return inv, jobID, err
}

// UpdateModel updates the migrates the current invoice to next version with data from the update payload
func (s service) UpdateModel(ctx context.Context, payload documents.UpdatePayload) (documents.Model, jobs.JobID, error) {
	if payload.Data == nil {
		return nil, jobs.NilJobID(), documents.ErrDocumentNil
	}

	did, err := contextutil.AccountDID(ctx)
	if err != nil {
		return nil, jobs.NilJobID(), documents.ErrDocumentConfigAccountID
	}

	old, err := s.GetCurrentVersion(ctx, payload.DocumentID)
	if err != nil {
		return nil, jobs.NilJobID(), err
	}

	oldInv, ok := old.(*Invoice)
	if !ok {
		return nil, jobs.NilJobID(), errors.NewTypedError(documents.ErrDocumentInvalidType, errors.New("%v is not an invoice", hexutil.Encode(payload.DocumentID)))
	}

	inv := new(Invoice)
	err = inv.unpackFromUpdatePayload(oldInv, payload)
	if err != nil {
		return nil, jobs.NilJobID(), errors.NewTypedError(documents.ErrDocumentInvalid, err)
	}

	err = UpdateValidator(s.anchorRepo).Validate(old, inv)
	if err != nil {
		return nil, jobs.NilJobID(), errors.NewTypedError(documents.ErrDocumentInvalid, err)
	}

	err = s.repo.Create(did[:], inv.CurrentVersion(), inv)
	if err != nil {
		return nil, jobs.NilJobID(), errors.NewTypedError(documents.ErrDocumentPersistence, err)
	}

	jobID := contextutil.Job(ctx)
	jobID, _, err = documents.CreateAnchorJob(ctx, s.jobManager, s.queueSrv, did, jobID, inv.CurrentVersion())
	return inv, jobID, err
}
