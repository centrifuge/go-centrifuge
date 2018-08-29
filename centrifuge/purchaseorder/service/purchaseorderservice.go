package purchaseorderservice

import (
	"fmt"

	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/purchaseorder"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/coredocument"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/coredocument/processor"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/coredocument/repository"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/errors"
	clientpurchaseorderpb "github.com/CentrifugeInc/go-centrifuge/centrifuge/protobufs/gen/go/purchaseorder"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/purchaseorder"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/storage"
	googleprotobuf2 "github.com/golang/protobuf/ptypes/empty"
	logging "github.com/ipfs/go-log"
	"golang.org/x/net/context"
)

var log = logging.Logger("rest-api")

// PurchaseOrderDocumentService needed as it is used to register the grpc services attached to the grpc server
type PurchaseOrderDocumentService struct {
	Repository            storage.Repository
	CoreDocumentProcessor coredocumentprocessor.Processor
}

func fillCoreDocIdentifiers(doc *purchaseorderpb.PurchaseOrderDocument) error {
	if doc == nil {
		return errors.NilError(doc)
	}
	err := coredocument.FillIdentifiers(doc.CoreDocument)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// HandleCreatePurchaseOrderProof creates proofs for a list of fields
func (s *PurchaseOrderDocumentService) HandleCreatePurchaseOrderProof(ctx context.Context, createPurchaseOrderProofEnvelope *clientpurchaseorderpb.CreatePurchaseOrderProofEnvelope) (*clientpurchaseorderpb.PurchaseOrderProof, error) {
	orderDoc := new(purchaseorderpb.PurchaseOrderDocument)
	err := s.Repository.GetByID(createPurchaseOrderProofEnvelope.DocumentIdentifier, orderDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to get document from DB: %v", err)
	}

	order, err := purchaseorder.NewPurchaseOrder(orderDoc)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to create new order: %v", err)
	}

	proofs, err := order.CreateProofs(createPurchaseOrderProofEnvelope.Fields)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to create proofs: %v", err)
	}

	return &clientpurchaseorderpb.PurchaseOrderProof{FieldProofs: proofs, DocumentIdentifier: order.Document.CoreDocument.DocumentIdentifier}, nil
}

// HandleAnchorPurchaseOrderDocument anchors the given purchaseorder document and returns the anchor details
func (s *PurchaseOrderDocumentService) HandleAnchorPurchaseOrderDocument(ctx context.Context, anchorPurchaseOrderEnvelope *clientpurchaseorderpb.AnchorPurchaseOrderEnvelope) (*purchaseorderpb.PurchaseOrderDocument, error) {
	doc := anchorPurchaseOrderEnvelope.Document

	err := fillCoreDocIdentifiers(doc)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to fill document IDs: %v", err)
	}

	err = s.Repository.Create(doc.CoreDocument.DocumentIdentifier, doc)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to save document: %v", err)
	}

	anchoredPurchaseOrder, err := s.anchorPurchaseOrderDocument(doc)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to anchor document: %v", err)
	}

	return anchoredPurchaseOrder, nil
}

// HandleSendPurchaseOrderDocument anchors and sends an purchaseorder to the recipient
func (s *PurchaseOrderDocumentService) HandleSendPurchaseOrderDocument(ctx context.Context, sendPurchaseOrderEnvelope *clientpurchaseorderpb.SendPurchaseOrderEnvelope) (*purchaseorderpb.PurchaseOrderDocument, error) {
	doc := sendPurchaseOrderEnvelope.Document

	err := fillCoreDocIdentifiers(doc)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = s.Repository.Create(doc.CoreDocument.DocumentIdentifier, doc)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	anchoredPurchaseOrder, err := s.anchorPurchaseOrderDocument(sendPurchaseOrderEnvelope.Document)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var errs []error
	for _, recipient := range sendPurchaseOrderEnvelope.Recipients {
		err = s.CoreDocumentProcessor.Send(anchoredPurchaseOrder.CoreDocument, ctx, recipient)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		log.Errorf("%v", errs)
		return nil, fmt.Errorf("%v", errs)
	}

	return anchoredPurchaseOrder, nil
}

func (s *PurchaseOrderDocumentService) HandleGetPurchaseOrderDocument(ctx context.Context, getPurchaseOrderDocumentEnvelope *clientpurchaseorderpb.GetPurchaseOrderDocumentEnvelope) (*purchaseorderpb.PurchaseOrderDocument, error) {
	doc := new(purchaseorderpb.PurchaseOrderDocument)
	err := s.Repository.GetByID(getPurchaseOrderDocumentEnvelope.DocumentIdentifier, doc)
	if err == nil {
		return doc, nil
	}

	docFound, err := coredocumentrepository.GetRepository().FindById(getPurchaseOrderDocumentEnvelope.DocumentIdentifier)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to get document: %v", err)
	}

	purchaseOrder, err := purchaseorder.NewPurchaseOrderFromCoreDocument(docFound)
	if err != nil {
		return nil, fmt.Errorf("failed convert coredoc to purchase order: %v", err)
	}

	return purchaseOrder.Document, nil
}

func (s *PurchaseOrderDocumentService) HandleGetReceivedPurchaseOrderDocuments(ctx context.Context, empty *googleprotobuf2.Empty) (*clientpurchaseorderpb.ReceivedPurchaseOrders, error) {
	return nil, nil
}

// anchorPurchaseOrderDocument anchors the given purchaseorder document and returns the anchor details
func (s *PurchaseOrderDocumentService) anchorPurchaseOrderDocument(doc *purchaseorderpb.PurchaseOrderDocument) (*purchaseorderpb.PurchaseOrderDocument, error) {
	// TODO: the calculated merkle root should be persisted locally as well.
	orderDoc, err := purchaseorder.NewPurchaseOrder(doc)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	orderDoc.CalculateMerkleRoot()
	coreDoc := orderDoc.ConvertToCoreDocument()

	err = s.CoreDocumentProcessor.Anchor(coreDoc)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	newPo, err := purchaseorder.NewPurchaseOrderFromCoreDocument(coreDoc)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return newPo.Document, nil
}
