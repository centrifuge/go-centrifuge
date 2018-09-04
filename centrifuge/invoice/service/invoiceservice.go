package invoiceservice

import (
	"fmt"

	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/coredocument"
	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/invoice"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/code"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/coredocument/processor"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/coredocument/repository"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/errors"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/invoice"
	clientinvoicepb "github.com/CentrifugeInc/go-centrifuge/centrifuge/protobufs/gen/go/invoice"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/storage"
	"github.com/golang/protobuf/ptypes/empty"
	logging "github.com/ipfs/go-log"
	"golang.org/x/net/context"
)

var log = logging.Logger("rest-api")

// InvoiceDocumentService handles all the invoice document related actions
// anchoring, sending, proof generation, finding stored invoice document
type InvoiceDocumentService struct {
	InvoiceRepository     storage.Repository
	CoreDocumentProcessor coredocumentprocessor.Processor
}

// anchorInvoiceDocument anchors the given invoice document and returns the anchored document
func (s *InvoiceDocumentService) anchorInvoiceDocument(doc *invoicepb.InvoiceDocument) (*invoicepb.InvoiceDocument, error) {
	inv, err := invoice.New(doc)
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
	newInvoice, err := invoice.NewFromCoreDocument(coreDoc)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return newInvoice.Document, nil
}

// HandleCreateInvoiceProof creates proofs for a list of fields
func (s *InvoiceDocumentService) HandleCreateInvoiceProof(ctx context.Context, createInvoiceProofEnvelope *clientinvoicepb.CreateInvoiceProofEnvelope) (*clientinvoicepb.InvoiceProof, error) {
	invDoc := new(invoicepb.InvoiceDocument)
	err := s.InvoiceRepository.GetByID(createInvoiceProofEnvelope.DocumentIdentifier, invDoc)
	if err != nil {
		return nil, err
	}

	inv, err := invoice.Wrap(invDoc)
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

// HandleAnchorInvoiceDocument anchors the given invoice document and returns the anchor details
func (s *InvoiceDocumentService) HandleAnchorInvoiceDocument(ctx context.Context, anchorInvoiceEnvelope *clientinvoicepb.AnchorInvoiceEnvelope) (*invoicepb.InvoiceDocument, error) {
	inv, err := invoice.New(anchorInvoiceEnvelope.Document)
	if err != nil {
		log.Error(err)
		return nil, errors.New(code.DocumentInvalid, err.Error())
	}

	err = s.InvoiceRepository.Create(inv.Document.CoreDocument.DocumentIdentifier, inv.Document)
	if err != nil {
		log.Error(err)
		return nil, errors.New(code.Unknown, fmt.Sprintf("error saving invoice: %v", err))
	}

	anchoredInvDoc, err := s.anchorInvoiceDocument(inv.Document)
	if err != nil {
		log.Error(err)
		return nil, errors.New(code.Unknown, fmt.Sprintf("failed to anchor: %v", err))
	}

	// Updating invoice with autogenerated fields after anchoring
	err = s.InvoiceRepository.Update(anchoredInvDoc.CoreDocument.DocumentIdentifier, anchoredInvDoc)
	if err != nil {
		log.Error(err)
		return nil, errors.New(code.Unknown, fmt.Sprintf("error saving document: %v", err))
	}

	return anchoredInvDoc, nil

}

// HandleSendInvoiceDocument anchors and sends an invoice to the recipient
func (s *InvoiceDocumentService) HandleSendInvoiceDocument(ctx context.Context, sendInvoiceEnvelope *clientinvoicepb.SendInvoiceEnvelope) (*invoicepb.InvoiceDocument, error) {
	doc, err := s.HandleAnchorInvoiceDocument(ctx, &clientinvoicepb.AnchorInvoiceEnvelope{Document: sendInvoiceEnvelope.Document})
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
		return nil, errors.New(code.Unknown, fmt.Sprintf("%v", errs))
	}

	return doc, nil
}

// HandleGetInvoiceDocument returns already stored invoice document
func (s *InvoiceDocumentService) HandleGetInvoiceDocument(ctx context.Context, getInvoiceDocumentEnvelope *clientinvoicepb.GetInvoiceDocumentEnvelope) (*invoicepb.InvoiceDocument, error) {
	doc := new(invoicepb.InvoiceDocument)
	err := s.InvoiceRepository.GetByID(getInvoiceDocumentEnvelope.DocumentIdentifier, doc)
	if err == nil {
		return doc, nil
	}

	coreDoc := new(coredocumentpb.CoreDocument)
	err = coredocumentrepository.GetRepository().GetByID(getInvoiceDocumentEnvelope.DocumentIdentifier, coreDoc)
	if err != nil {
		return nil, errors.New(code.DocumentNotFound, err.Error())
	}

	inv, err := invoice.NewFromCoreDocument(coreDoc)
	if err != nil {
		return nil, errors.New(code.Unknown, err.Error())
	}

	return inv.Document, nil
}

// HandleGetReceivedInvoiceDocuments returns all the received invoice documents
func (s *InvoiceDocumentService) HandleGetReceivedInvoiceDocuments(ctx context.Context, empty *empty.Empty) (*clientinvoicepb.ReceivedInvoices, error) {
	return nil, nil
}
