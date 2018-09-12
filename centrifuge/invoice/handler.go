package invoice

import (
	"fmt"

	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/coredocument"
	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/invoice"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/centerrors"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/code"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/coredocument"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/coredocument/processor"
	clientinvoicepb "github.com/CentrifugeInc/go-centrifuge/centrifuge/protobufs/gen/go/invoice"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
)

// Handler handles all the invoice document related actions
// anchoring, sending, proof generation, finding stored invoice document
type Handler struct {
	InvoiceRepository     storage.Repository
	CoreDocumentProcessor coredocumentprocessor.Processor
}

// anchorInvoiceDocument anchors the given invoice document and returns the anchored document
func (s *Handler) anchorInvoiceDocument(doc *invoicepb.InvoiceDocument) (*invoicepb.InvoiceDocument, error) {
	inv, err := New(doc)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	coreDoc, err := inv.ConvertToCoreDocument()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = s.CoreDocumentProcessor.Anchor(coreDoc)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// we do not need this conversion again
	newInvoice, err := NewFromCoreDocument(coreDoc)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return newInvoice.Document, nil
}

// CreateInvoiceProof creates proofs for a list of fields
func (s *Handler) CreateInvoiceProof(ctx context.Context, createInvoiceProofEnvelope *clientinvoicepb.CreateInvoiceProofEnvelope) (*clientinvoicepb.InvoiceProof, error) {
	invDoc := new(invoicepb.InvoiceDocument)
	err := s.InvoiceRepository.GetByID(createInvoiceProofEnvelope.DocumentIdentifier, invDoc)
	if err != nil {
		return nil, err
	}

	inv, err := Wrap(invDoc)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	proofs, err := inv.CreateProofs(createInvoiceProofEnvelope.Fields)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &clientinvoicepb.InvoiceProof{FieldProofs: proofs, DocumentIdentifier: inv.Document.CoreDocument.DocumentIdentifier}, nil

}

// AnchorInvoiceDocument anchors the given invoice document and returns the anchor details
func (s *Handler) AnchorInvoiceDocument(ctx context.Context, anchorInvoiceEnvelope *clientinvoicepb.AnchorInvoiceEnvelope) (*invoicepb.InvoiceDocument, error) {
	inv, err := New(anchorInvoiceEnvelope.Document)
	if err != nil {
		log.Error(err)
		return nil, centerrors.New(code.DocumentInvalid, err.Error())
	}

	err = s.InvoiceRepository.Create(inv.Document.CoreDocument.DocumentIdentifier, inv.Document)
	if err != nil {
		log.Error(err)
		return nil, centerrors.New(code.Unknown, fmt.Sprintf("error saving invoice: %v", err))
	}

	anchoredInvDoc, err := s.anchorInvoiceDocument(inv.Document)
	if err != nil {
		log.Error(err)
		return nil, centerrors.New(code.Unknown, fmt.Sprintf("failed to anchor: %v", err))
	}

	// Updating invoice with autogenerated fields after anchoring
	err = s.InvoiceRepository.Update(anchoredInvDoc.CoreDocument.DocumentIdentifier, anchoredInvDoc)
	if err != nil {
		log.Error(err)
		return nil, centerrors.New(code.Unknown, fmt.Sprintf("error saving document: %v", err))
	}

	return anchoredInvDoc, nil

}

// SendInvoiceDocument anchors and sends an invoice to the recipient
func (s *Handler) SendInvoiceDocument(ctx context.Context, sendInvoiceEnvelope *clientinvoicepb.SendInvoiceEnvelope) (*invoicepb.InvoiceDocument, error) {
	doc, err := s.AnchorInvoiceDocument(ctx, &clientinvoicepb.AnchorInvoiceEnvelope{Document: sendInvoiceEnvelope.Document})
	if err != nil {
		return nil, err
	}

	var errs []error
	for _, element := range sendInvoiceEnvelope.Recipients {
		err = s.CoreDocumentProcessor.Send(doc.CoreDocument, ctx, element[:])
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		log.Errorf("%v", errs)
		return nil, centerrors.New(code.Unknown, fmt.Sprintf("%v", errs))
	}

	return doc, nil
}

// GetInvoiceDocument returns already stored invoice document
func (s *Handler) GetInvoiceDocument(ctx context.Context, getInvoiceDocumentEnvelope *clientinvoicepb.GetInvoiceDocumentEnvelope) (*invoicepb.InvoiceDocument, error) {
	doc := new(invoicepb.InvoiceDocument)
	err := s.InvoiceRepository.GetByID(getInvoiceDocumentEnvelope.DocumentIdentifier, doc)
	if err == nil {
		return doc, nil
	}

	coreDoc := new(coredocumentpb.CoreDocument)
	err = coredocument.GetRepository().GetByID(getInvoiceDocumentEnvelope.DocumentIdentifier, coreDoc)
	if err != nil {
		return nil, centerrors.New(code.DocumentNotFound, err.Error())
	}

	inv, err := NewFromCoreDocument(coreDoc)
	if err != nil {
		return nil, centerrors.New(code.Unknown, err.Error())
	}

	return inv.Document, nil
}

// GetReceivedInvoiceDocuments returns all the received invoice documents
func (s *Handler) GetReceivedInvoiceDocuments(ctx context.Context, empty *empty.Empty) (*clientinvoicepb.ReceivedInvoices, error) {
	return nil, nil
}
