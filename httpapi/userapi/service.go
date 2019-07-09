package userapi

import (
	"context"

	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/documents/entity"
	"github.com/centrifuge/go-centrifuge/documents/entityrelationship"
	"github.com/centrifuge/go-centrifuge/documents/invoice"
	"github.com/centrifuge/go-centrifuge/documents/purchaseorder"
	"github.com/centrifuge/go-centrifuge/extensions/transferdetails"
	"github.com/centrifuge/go-centrifuge/httpapi/coreapi"
	"github.com/centrifuge/go-centrifuge/jobs"
)

// Service provides functionality for User APIs.
type Service struct {
	coreAPISrv             coreapi.Service
	transferDetailsService transferdetails.Service
	entityRelationshipSrv  entityrelationship.Service
}

// TODO: this can be refactored into a generic Service which handles all kinds of custom attributes

// CreateTransferDetail creates and anchors a Transfer Detail
func (s Service) CreateTransferDetail(ctx context.Context, req transferdetails.CreateTransferDetailRequest) (documents.Model, jobs.JobID, error) {
	return s.transferDetailsService.CreateTransferDetail(ctx, req)
}

// UpdateTransferDetail updates and anchors a Transfer Detail
func (s Service) UpdateTransferDetail(ctx context.Context, req transferdetails.UpdateTransferDetailRequest) (documents.Model, jobs.JobID, error) {
	return s.transferDetailsService.UpdateTransferDetail(ctx, req)
}

// GetCurrentTransferDetail returns the current version on a Transfer Detail
func (s Service) GetCurrentTransferDetail(ctx context.Context, docID, transferID []byte) (*transferdetails.TransferDetail, documents.Model, error) {
	model, err := s.coreAPISrv.GetDocument(ctx, docID)
	if err != nil {
		return nil, nil, err
	}
	data, model, err := s.transferDetailsService.DeriveTransferDetail(ctx, model, transferID)
	if err != nil {
		return nil, nil, err
	}

	return data, model, nil
}

// GetCurrentTransferDetailsList returns a list of Transfer Details on the current version of a document
func (s Service) GetCurrentTransferDetailsList(ctx context.Context, docID []byte) (*transferdetails.TransferDetailList, documents.Model, error) {
	model, err := s.coreAPISrv.GetDocument(ctx, docID)
	if err != nil {
		return nil, nil, err
	}

	data, model, err := s.transferDetailsService.DeriveTransferList(ctx, model)
	if err != nil {
		return nil, nil, err
	}

	return data, model, nil
}

// GetVersionTransferDetail returns a Transfer Detail on a particular version of a Document
func (s Service) GetVersionTransferDetail(ctx context.Context, docID, versionID, transferID []byte) (*transferdetails.TransferDetail, documents.Model, error) {
	model, err := s.coreAPISrv.GetDocumentVersion(ctx, docID, versionID)
	if err != nil {
		return nil, nil, err
	}

	data, model, err := s.transferDetailsService.DeriveTransferDetail(ctx, model, transferID)
	if err != nil {
		return nil, nil, err
	}

	return data, model, nil
}

// GetVersionTransferDetailsList returns a list of Transfer Details on a particular version of a Document
func (s Service) GetVersionTransferDetailsList(ctx context.Context, docID, versionID []byte) (*transferdetails.TransferDetailList, documents.Model, error) {
	model, err := s.coreAPISrv.GetDocumentVersion(ctx, docID, versionID)
	if err != nil {
		return nil, nil, err
	}

	data, model, err := s.transferDetailsService.DeriveTransferList(ctx, model)
	if err != nil {
		return nil, nil, err
	}

	return data, model, nil
}

func convertInvRequest(req CreateInvoiceRequest) (documents.CreatePayload, error) {
	coreAPIReq := coreapi.CreateDocumentRequest{
		Scheme:      invoice.Scheme,
		WriteAccess: req.WriteAccess,
		ReadAccess:  req.ReadAccess,
		Data:        req.Data,
		Attributes:  req.Attributes,
	}

	return coreapi.ToDocumentsCreatePayload(coreAPIReq)
}

// CreateInvoice creates an invoice.
func (s Service) CreateInvoice(ctx context.Context, req CreateInvoiceRequest) (documents.Model, jobs.JobID, error) {
	docReq, err := convertInvRequest(req)
	if err != nil {
		return nil, jobs.NilJobID(), err
	}

	return s.coreAPISrv.CreateDocument(ctx, docReq)
}

func convertPORequest(req CreatePurchaseOrderRequest) (documents.CreatePayload, error) {
	coreAPIReq := coreapi.CreateDocumentRequest{
		Scheme:      purchaseorder.Scheme,
		WriteAccess: req.WriteAccess,
		ReadAccess:  req.ReadAccess,
		Data:        req.Data,
		Attributes:  req.Attributes,
	}

	return coreapi.ToDocumentsCreatePayload(coreAPIReq)
}

// CreatePurchaseOrder creates a purchase Order.
func (s Service) CreatePurchaseOrder(ctx context.Context, req CreatePurchaseOrderRequest) (documents.Model, jobs.JobID, error) {
	docReq, err := convertPORequest(req)
	if err != nil {
		return nil, jobs.NilJobID(), err
	}

	return s.coreAPISrv.CreateDocument(ctx, docReq)
}

// GetPurchaseOrder returns the latest version of the PurchaseOrder associated with Document ID.
func (s Service) GetPurchaseOrder(ctx context.Context, docID []byte) (documents.Model, error) {
	return s.coreAPISrv.GetDocument(ctx, docID)
}

// UpdatePurchaseOrder updates a purchase Order.
func (s Service) UpdatePurchaseOrder(ctx context.Context, docID []byte, req CreatePurchaseOrderRequest) (documents.Model, jobs.JobID, error) {
	docReq, err := convertPORequest(req)
	if err != nil {
		return nil, jobs.NilJobID(), err
	}

	return s.coreAPISrv.UpdateDocument(ctx, documents.UpdatePayload{CreatePayload: docReq, DocumentID: docID})
}

// GetPurchaseOrderVersion returns specific version of the PurchaseOrder associated with Document ID.
func (s Service) GetPurchaseOrderVersion(ctx context.Context, docID, versionID []byte) (documents.Model, error) {
	return s.coreAPISrv.GetDocumentVersion(ctx, docID, versionID)
}

func convertEntityRequest(req CreateEntityRequest) (documents.CreatePayload, error) {
	coreAPIReq := coreapi.CreateDocumentRequest{
		Scheme:      entity.Scheme,
		WriteAccess: req.WriteAccess,
		ReadAccess:  req.ReadAccess,
		Data:        req.Data,
		Attributes:  req.Attributes,
	}

	return coreapi.ToDocumentsCreatePayload(coreAPIReq)
}

// CreateEntity creates Entity document and anchors it.
func (s Service) CreateEntity(ctx context.Context, req CreateEntityRequest) (documents.Model, jobs.JobID, error) {
	docReq, err := convertEntityRequest(req)
	if err != nil {
		return nil, jobs.NilJobID(), err
	}

	return s.coreAPISrv.CreateDocument(ctx, docReq)
}

// UpdateEntity updates existing entity associated with docID  with provided data and anchors it.
func (s Service) UpdateEntity(ctx context.Context, docID []byte, req CreateEntityRequest) (documents.Model, jobs.JobID, error) {
	docReq, err := convertEntityRequest(req)
	if err != nil {
		return nil, jobs.NilJobID(), err
	}

	return s.coreAPISrv.UpdateDocument(ctx, documents.UpdatePayload{
		DocumentID:    docID,
		CreatePayload: docReq,
	})
}

// GetEntity returns the Entity associated with docID.
func (s Service) GetEntity(ctx context.Context, docID []byte) (documents.Model, error) {
	return s.coreAPISrv.GetDocument(ctx, docID)
}
