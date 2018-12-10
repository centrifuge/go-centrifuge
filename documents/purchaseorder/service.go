package purchaseorder

import (
	"bytes"
	"fmt"
	"time"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/notification"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/p2p"
	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/coredocument"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/header"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/notification"
	clientpopb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/purchaseorder"
	"github.com/centrifuge/go-centrifuge/signatures"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/golang/protobuf/ptypes"
	logging "github.com/ipfs/go-log"
)

var srvLog = logging.Logger("po-service")

// Service defines specific functions for purchase order
type Service interface {
	documents.Service

	// DeriverFromPayload derives purchase order from clientPayload
	DeriveFromCreatePayload(payload *clientpopb.PurchaseOrderCreatePayload, hdr *header.ContextHeader) (documents.Model, error)

	// DeriveFromUpdatePayload derives purchase order from update payload
	DeriveFromUpdatePayload(payload *clientpopb.PurchaseOrderUpdatePayload, hdr *header.ContextHeader) (documents.Model, error)

	// Create validates and persists purchase order and returns a Updated model
	Create(ctx *header.ContextHeader, po documents.Model) (documents.Model, error)

	// Update validates and updates the purchase order and return the updated model
	Update(ctx *header.ContextHeader, po documents.Model) (documents.Model, error)

	// DerivePurchaseOrderData returns the purchase order data as client data
	DerivePurchaseOrderData(po documents.Model) (*clientpopb.PurchaseOrderData, error)

	// DerivePurchaseOrderResponse returns the purchase order in our standard client format
	DerivePurchaseOrderResponse(po documents.Model) (*clientpopb.PurchaseOrderResponse, error)
}

// service implements Service and handles all purchase order related persistence and validations
// service always returns errors of type `errors.Error` or `errors.TypedError`
type service struct {
	repo             documents.LegacyRepository
	coreDocProcessor coredocument.Processor
	notifier         notification.Sender
	anchorRepository anchors.AnchorRepository
	identityService  identity.Service
}

// DefaultService returns the default implementation of the service
func DefaultService(config config.Configuration, repo documents.LegacyRepository, processor coredocument.Processor, anchorRepository anchors.AnchorRepository, identityService identity.Service) Service {
	return service{repo: repo, coreDocProcessor: processor, notifier: notification.NewWebhookSender(config), anchorRepository: anchorRepository, identityService: identityService}
}

// DeriveFromCoreDocument takes a core document and returns a purchase order
func (s service) DeriveFromCoreDocument(cd *coredocumentpb.CoreDocument) (documents.Model, error) {
	var model documents.Model = new(PurchaseOrder)
	err := model.UnpackCoreDocument(cd)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentUnPackingCoreDocument, err)
	}

	return model, nil
}

// calculateDataRoot validates the document, calculates the data root, and persists to DB
func (s service) calculateDataRoot(old, new documents.Model, validator documents.Validator) (documents.Model, error) {
	po, ok := new.(*PurchaseOrder)
	if !ok {
		return nil, errors.NewTypeError(documents.ErrDocumentInvalidType, fmt.Errorf("unknown document type: %T", new))
	}

	// create data root, has to be done at the model level to access fields
	err := po.calculateDataRoot()
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentInvalid, err)
	}

	// validate the invoice
	err = validator.Validate(old, po)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentInvalid, err)
	}

	// we use CurrentVersion as the id since that will be unique across multiple versions of the same document
	err = s.repo.Create(po.CoreDocument.CurrentVersion, po)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentPersistence, err)
	}

	return po, nil
}

// Create validates, persists, and anchors a purchase order
func (s service) Create(ctx *header.ContextHeader, po documents.Model) (documents.Model, error) {
	po, err := s.calculateDataRoot(nil, po, CreateValidator())
	if err != nil {
		return nil, err
	}

	po, err = documents.AnchorDocument(ctx, po, s.coreDocProcessor, s.repo.Update)
	if err != nil {
		return nil, err
	}

	return po, nil
}

// Update validates, persists, and anchors a new version of purchase order
func (s service) Update(ctx *header.ContextHeader, po documents.Model) (documents.Model, error) {
	cd, err := po.PackCoreDocument()
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentPackingCoreDocument, err)
	}

	old, err := s.GetCurrentVersion(cd.DocumentIdentifier)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentNotFound, err)
	}

	po, err = s.calculateDataRoot(old, po, UpdateValidator())
	if err != nil {
		return nil, err
	}

	po, err = documents.AnchorDocument(ctx, po, s.coreDocProcessor, s.repo.Update)
	if err != nil {
		return nil, err
	}

	return po, nil
}

// DeriveFromCreatePayload derives purchase order from create payload
func (s service) DeriveFromCreatePayload(payload *clientpopb.PurchaseOrderCreatePayload, ctxH *header.ContextHeader) (documents.Model, error) {
	if payload == nil || payload.Data == nil {
		return nil, documents.ErrDocumentProvidedIsNil
	}

	po := new(PurchaseOrder)
	err := po.InitPurchaseOrderInput(payload, ctxH)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentInvalid, err)
	}

	return po, nil
}

// DeriveFromUpdatePayload derives purchase order from update payload
func (s service) DeriveFromUpdatePayload(payload *clientpopb.PurchaseOrderUpdatePayload, ctxH *header.ContextHeader) (documents.Model, error) {
	if payload == nil || payload.Data == nil {
		return nil, documents.ErrDocumentProvidedIsNil
	}

	// get latest old version of the document
	id, err := hexutil.Decode(payload.Identifier)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentIdentifier, fmt.Errorf("failed to decode identifier: %v", err))
	}

	old, err := s.GetCurrentVersion(id)
	if err != nil {
		return nil, err
	}

	// load purchase order data
	po := new(PurchaseOrder)
	err = po.initPurchaseOrderFromData(payload.Data)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentInvalid, fmt.Errorf("failed to load purchase order from data: %v", err))
	}

	// update core document
	oldCD, err := old.PackCoreDocument()
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentPackingCoreDocument, err)
	}

	collaborators := append([]string{ctxH.Self().ID.String()}, payload.Collaborators...)
	po.CoreDocument, err = coredocument.PrepareNewVersion(*oldCD, collaborators)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentPrepareCoreDocument, err)
	}

	return po, nil
}

// DerivePurchaseOrderData returns po data from the model
func (s service) DerivePurchaseOrderData(doc documents.Model) (*clientpopb.PurchaseOrderData, error) {
	po, ok := doc.(*PurchaseOrder)
	if !ok {
		return nil, documents.ErrDocumentInvalidType
	}

	return po.getClientData(), nil
}

// DerivePurchaseOrderResponse returns po response from the model
func (s service) DerivePurchaseOrderResponse(doc documents.Model) (*clientpopb.PurchaseOrderResponse, error) {
	cd, err := doc.PackCoreDocument()
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentPackingCoreDocument, err)
	}

	collaborators := make([]string, len(cd.Collaborators))
	for i, c := range cd.Collaborators {
		cid, err := identity.ToCentID(c)
		if err != nil {
			return nil, errors.NewTypeError(documents.ErrDocumentCollaborator, err)
		}
		collaborators[i] = cid.String()
	}

	header := &clientpopb.ResponseHeader{
		DocumentId:    hexutil.Encode(cd.DocumentIdentifier),
		VersionId:     hexutil.Encode(cd.CurrentVersion),
		Collaborators: collaborators,
	}

	data, err := s.DerivePurchaseOrderData(doc)
	if err != nil {
		return nil, err
	}

	return &clientpopb.PurchaseOrderResponse{
		Header: header,
		Data:   data,
	}, nil
}

func (s service) getPurchaseOrderVersion(documentID, version []byte) (model *PurchaseOrder, err error) {
	var doc documents.Model = new(PurchaseOrder)
	err = s.repo.LoadByID(version, doc)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentVersionNotFound, err)
	}
	model, ok := doc.(*PurchaseOrder)
	if !ok {
		return nil, documents.ErrDocumentInvalidType
	}

	if !bytes.Equal(model.CoreDocument.DocumentIdentifier, documentID) {
		return nil, errors.NewTypeError(documents.ErrDocumentVersionNotFound, fmt.Errorf("version is not valid for this identifier"))
	}
	return model, nil
}

// GetLastVersion returns the latest version of the document
func (s service) GetCurrentVersion(documentID []byte) (documents.Model, error) {
	model, err := s.getPurchaseOrderVersion(documentID, documentID)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentNotFound, err)
	}
	nextVersion := model.CoreDocument.NextVersion
	for nextVersion != nil {
		temp, err := s.getPurchaseOrderVersion(documentID, nextVersion)
		if err != nil {
			// here the err is returned as nil because it is expected that the nextVersion is not available in the db at some stage of the iteration
			return model, nil
		}

		model = temp
		nextVersion = model.CoreDocument.NextVersion
	}
	return model, nil
}

// GetVersion returns the specific version of the document
func (s service) GetVersion(documentID []byte, version []byte) (documents.Model, error) {
	po, err := s.getPurchaseOrderVersion(documentID, version)
	if err != nil {
		return nil, err
	}
	return po, nil

}

// purchaseOrderProof creates proofs for purchaseOrder model fields
func (s service) purchaseOrderProof(model documents.Model, fields []string) (*documents.DocumentProof, error) {
	po, ok := model.(*PurchaseOrder)
	if !ok {
		return nil, documents.ErrDocumentInvalidType
	}
	if err := coredocument.PostAnchoredValidator(s.identityService, s.anchorRepository).Validate(nil, po); err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentInvalid, err)
	}
	coreDoc, proofs, err := po.createProofs(fields)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentProof, err)
	}
	return &documents.DocumentProof{
		DocumentID:  coreDoc.DocumentIdentifier,
		VersionID:   coreDoc.CurrentVersion,
		FieldProofs: proofs,
	}, nil
}

// CreateProofs generates proofs for given document
func (s service) CreateProofs(documentID []byte, fields []string) (*documents.DocumentProof, error) {
	model, err := s.GetCurrentVersion(documentID)
	if err != nil {
		return nil, err
	}
	return s.purchaseOrderProof(model, fields)
}

// CreateProofsForVersion generates proofs for specific version of the document
func (s service) CreateProofsForVersion(documentID, version []byte, fields []string) (*documents.DocumentProof, error) {
	model, err := s.GetVersion(documentID, version)
	if err != nil {
		return nil, err
	}
	return s.purchaseOrderProof(model, fields)
}

// RequestDocumentSignature validates the document and returns the signature
// Note: this is document agnostic. But since we do not have a common implementation, adding it here.
// will remove this once we have a common implementation for documents.Service
func (s service) RequestDocumentSignature(contextHeader *header.ContextHeader, model documents.Model) (*coredocumentpb.Signature, error) {
	if err := coredocument.SignatureRequestValidator(s.identityService).Validate(nil, model); err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentInvalid, err)
	}

	cd, err := model.PackCoreDocument()
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentPackingCoreDocument, err)
	}

	srvLog.Infof("coredoc received %x with signing root %x", cd.DocumentIdentifier, cd.SigningRoot)

	idKeys, ok := contextHeader.Self().Keys[identity.KeyPurposeSigning]
	if !ok {
		return nil, errors.NewTypeError(documents.ErrDocumentSigning, fmt.Errorf("missing signing key"))
	}
	sig := signatures.Sign(contextHeader.Self().ID[:], idKeys.PrivateKey, idKeys.PublicKey, cd.SigningRoot)
	cd.Signatures = append(cd.Signatures, sig)
	err = model.UnpackCoreDocument(cd)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentUnPackingCoreDocument, err)
	}

	// Logic for receiving version n (n > 1) of the document for the first time
	if !s.repo.Exists(cd.DocumentIdentifier) && !utils.IsSameByteSlice(cd.DocumentIdentifier, cd.CurrentVersion) {
		err = s.repo.Create(cd.DocumentIdentifier, model)
		if err != nil {
			return nil, errors.NewTypeError(documents.ErrDocumentPersistence, err)
		}
	}

	err = s.repo.Create(cd.CurrentVersion, model)
	if err != nil {
		return nil, errors.NewTypeError(documents.ErrDocumentPersistence, err)
	}

	srvLog.Infof("signed coredoc %x with version %x", cd.DocumentIdentifier, cd.CurrentVersion)
	return sig, nil
}

// ReceiveAnchoredDocument validates the anchored document and updates it on DB
// Note: this is document agnostic. But since we do not have a common implementation, adding it here.
// will remove this once we have a common implementation for documents.Service
func (s service) ReceiveAnchoredDocument(model documents.Model, headers *p2ppb.CentrifugeHeader) error {
	if err := coredocument.PostAnchoredValidator(s.identityService, s.anchorRepository).Validate(nil, model); err != nil {
		return errors.NewTypeError(documents.ErrDocumentInvalid, err)
	}

	doc, err := model.PackCoreDocument()
	if err != nil {
		return errors.NewTypeError(documents.ErrDocumentPackingCoreDocument, err)
	}

	err = s.repo.Update(doc.CurrentVersion, model)
	if err != nil {
		return errors.NewTypeError(documents.ErrDocumentPersistence, err)
	}

	ts, _ := ptypes.TimestampProto(time.Now().UTC())
	notificationMsg := &notificationpb.NotificationMessage{
		EventType:    uint32(notification.ReceivedPayload),
		CentrifugeId: hexutil.Encode(headers.SenderCentrifugeId),
		Recorded:     ts,
		DocumentType: doc.EmbeddedData.TypeUrl,
		DocumentId:   hexutil.Encode(doc.DocumentIdentifier),
	}

	// Async until we add queuing
	go s.notifier.Send(notificationMsg)

	return nil
}

// Exists checks if an purchase order exists
func (s service) Exists(documentID []byte) bool {
	return s.repo.Exists(documentID)
}
