package entity

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/centrifuge/centrifuge-protobufs/documenttypes"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/centrifuge-protobufs/gen/go/entity"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/identity"
	cliententitypb "github.com/centrifuge/go-centrifuge/protobufs/gen/go/entity"
	"github.com/centrifuge/go-centrifuge/utils/byteutils"
	"github.com/centrifuge/precise-proofs/proofs"
	"github.com/centrifuge/precise-proofs/proofs/proto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
)

const (
	prefix string = "entity"
	scheme        = prefix

	// ErrMultiplePaymentMethodsSet is a sentinel error when multiple payment methods are set in a single payment detail.
	ErrMultiplePaymentMethodsSet = errors.Error("multiple payment methods are set")

	// ErrNoPaymentMethodSet is a sentinel error when no payment method is set in a single payment detail.
	ErrNoPaymentMethodSet = errors.Error("no payment method is set")

	// ErrEntityInvalidData sentinel error when data unmarshal is failed.
	ErrEntityInvalidData = errors.Error("invalid entity data")
)

// tree prefixes for specific to documents use the second byte of a 4 byte slice by convention
func compactPrefix() []byte { return []byte{0, 3, 0, 0} }

// Address holds the address details of the entity.
type Address struct {
	IsMain        bool   `json:"is_main"`
	IsRemitTo     bool   `json:"is_remit_to"`
	IsShipTo      bool   `json:"is_ship_to"`
	IsPayTo       bool   `json:"is_pay_to"`
	Label         string `json:"label"`
	Zip           string `json:"zip"`
	State         string `json:"state"`
	Country       string `json:"country"`
	AddressLine1  string `json:"address_line_1"`
	AddressLine2  string `json:"address_line_2"`
	ContactPerson string `json:"contact_person"`
}

// BankPaymentMethod holds the bank details of the entity.
type BankPaymentMethod struct {
	Identifier        byteutils.HexBytes `json:"identifier"`
	Address           Address            `json:"address"`
	HolderName        string             `json:"holder_name"`
	BankKey           string             `json:"bank_key"`
	BankAccountNumber string             `json:"bank_account_number"`
	SupportedCurrency string             `json:"supported_currency"`
}

// CryptoPaymentMethod holds the crypto details of the entity.
type CryptoPaymentMethod struct {
	Identifier        byteutils.HexBytes `json:"identifier"`
	To                string             `json:"to"`
	ChainURI          string             `json:"chain_uri"`
	SupportedCurrency string             `json:"supported_currency"`
}

// OtherPaymentMethod represents any other payment methods entity accepts.
type OtherPaymentMethod struct {
	Identifier        byteutils.HexBytes `json:"identifier"`
	Type              string             `json:"type"`
	PayTo             string             `json:"pay_to"`
	SupportedCurrency string             `json:"supported_currency"`
}

// PaymentDetail contains the payments receiving details of the entity.
// Note: only one of the payment methods has to be set for a given payment detail.
type PaymentDetail struct {
	Predefined          bool                 `json:"predefined"`
	BankPaymentMethod   *BankPaymentMethod   `json:"bank_payment_method,omitempty"`
	CryptoPaymentMethod *CryptoPaymentMethod `json:"crypto_payment_method,omitempty"`
	OtherPaymentMethod  *OtherPaymentMethod  `json:"other_payment_method,omitempty"`
}

// Contact holds the entity contact details.
type Contact struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Fax   string `json:"fax"`
}

// Data represents the entity data.
type Data struct {
	Identity       *identity.DID   `json:"identity"`
	LegalName      string          `json:"legal_name"`
	Addresses      []Address       `json:"addresses"`
	PaymentDetails []PaymentDetail `json:"payment_details"`
	Contacts       []Contact       `json:"contacts"`
}

// Entity implements the documents.Model keeps track of entity related fields and state
type Entity struct {
	*documents.CoreDocument

	Data Data
}

// getClientData returns the client data from the entity model
func (e *Entity) getClientData() *cliententitypb.EntityData {
	d := e.Data
	dids := identity.DIDsToStrings(d.Identity)
	return &cliententitypb.EntityData{
		Identity:       dids[0],
		LegalName:      d.LegalName,
		Addresses:      toProtoAddresses(d.Addresses),
		PaymentDetails: toProtoPaymentDetails(d.PaymentDetails),
		Contacts:       toProtoContacts(d.Contacts),
	}
}

// createP2PProtobuf returns centrifuge protobuf specific entityData
func (e *Entity) createP2PProtobuf() *entitypb.Entity {
	d := e.Data
	dids := identity.DIDsToBytes(d.Identity)
	return &entitypb.Entity{
		Identity:       dids[0],
		LegalName:      d.LegalName,
		Addresses:      toProtoAddresses(d.Addresses),
		PaymentDetails: toProtoPaymentDetails(d.PaymentDetails),
		Contacts:       toProtoContacts(d.Contacts),
	}
}

// InitEntityInput initialize the model based on the received parameters from the rest api call
func (e *Entity) InitEntityInput(payload *cliententitypb.EntityCreatePayload, self identity.DID) error {
	err := e.initEntityFromData(payload.Data)
	if err != nil {
		return err
	}

	ca, err := documents.FromClientCollaboratorAccess(payload.ReadAccess, payload.WriteAccess)
	if err != nil {
		return errors.New("failed to decode collaborator: %v", err)
	}

	ca.ReadWriteCollaborators = append(ca.ReadWriteCollaborators, self)
	attrs, err := documents.FromClientAttributes(payload.Attributes)
	if err != nil {
		return err
	}

	cd, err := documents.NewCoreDocument(compactPrefix(), ca, attrs)
	if err != nil {
		return errors.New("failed to init core document: %v", err)
	}

	e.CoreDocument = cd
	return nil
}

// initEntityFromData initialises entity from entityData
func (e *Entity) initEntityFromData(data *cliententitypb.EntityData) error {
	data.Identity = strings.TrimSpace(data.Identity)
	if data.Identity == "" {
		return identity.ErrMalformedAddress
	}

	dids, err := identity.StringsToDIDs(data.Identity)
	if err != nil {
		return errors.NewTypedError(identity.ErrMalformedAddress, err)
	}

	var d Data
	d.Identity = dids[0]
	d.LegalName = data.LegalName
	d.Addresses = fromProtoAddresses(data.Addresses)
	d.PaymentDetails = fromProtoPaymentDetails(data.PaymentDetails)
	d.Contacts = fromProtoContacts(data.Contacts)
	e.Data = d
	return nil
}

// loadFromP2PProtobuf  loads the entity from centrifuge protobuf entity data
func (e *Entity) loadFromP2PProtobuf(data *entitypb.Entity) error {
	dids, err := identity.BytesToDIDs(data.Identity)
	if err != nil {
		return err
	}

	var d Data
	d.Identity = dids[0]
	d.LegalName = data.LegalName
	d.Addresses = fromProtoAddresses(data.Addresses)
	d.PaymentDetails = fromProtoPaymentDetails(data.PaymentDetails)
	d.Contacts = fromProtoContacts(data.Contacts)
	e.Data = d
	return nil
}

// PackCoreDocument packs the Entity into a CoreDocument.
func (e *Entity) PackCoreDocument() (cd coredocumentpb.CoreDocument, err error) {
	entityData := e.createP2PProtobuf()
	data, err := proto.Marshal(entityData)
	if err != nil {
		return cd, errors.New("couldn't serialise EntityData: %v", err)
	}

	embedData := &any.Any{
		TypeUrl: e.DocumentType(),
		Value:   data,
	}

	return e.CoreDocument.PackCoreDocument(embedData), nil
}

// UnpackCoreDocument unpacks the core document into Entity.
func (e *Entity) UnpackCoreDocument(cd coredocumentpb.CoreDocument) error {
	if cd.EmbeddedData == nil ||
		cd.EmbeddedData.TypeUrl != e.DocumentType() {
		return errors.New("trying to convert document with incorrect schema")
	}

	entityData := new(entitypb.Entity)
	err := proto.Unmarshal(cd.EmbeddedData.Value, entityData)
	if err != nil {
		return err
	}

	err = e.loadFromP2PProtobuf(entityData)
	if err != nil {
		return err
	}

	e.CoreDocument, err = documents.NewCoreDocumentFromProtobuf(cd)
	return err
}

// JSON marshals Entity into a json bytes
func (e *Entity) JSON() ([]byte, error) {
	return e.CoreDocument.MarshalJSON(e)
}

// FromJSON unmarshals the json bytes into Entity
func (e *Entity) FromJSON(jsonData []byte) error {
	if e.CoreDocument == nil {
		e.CoreDocument = new(documents.CoreDocument)
	}

	return e.CoreDocument.UnmarshalJSON(jsonData, e)
}

// Type gives the Entity type
func (e *Entity) Type() reflect.Type {
	return reflect.TypeOf(e)
}

func (e *Entity) getDataLeaves() ([]proofs.LeafNode, error) {
	t, err := e.getRawDataTree()
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDataTree, err)
	}
	return t.GetLeaves(), nil
}

func (e *Entity) getRawDataTree() (*proofs.DocumentTree, error) {
	entityProto := e.createP2PProtobuf()
	if e.CoreDocument == nil {
		return nil, errors.New("getDataTree error CoreDocument not set")
	}
	t, err := e.CoreDocument.DefaultTreeWithPrefix(prefix, compactPrefix())
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDataTree, err)
	}
	err = t.AddLeavesFromDocument(entityProto)
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDataTree, err)
	}
	return t, nil
}

// getDataTree creates precise-proofs data tree for the model
func (e *Entity) getDataTree() (*proofs.DocumentTree, error) {
	tree, err := e.getRawDataTree()
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDataTree, err)
	}
	err = tree.Generate()
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDataTree, err)
	}

	return tree, nil
}

// CreateNFTProofs creates proofs specific to NFT minting.
func (e *Entity) CreateNFTProofs(
	account identity.DID,
	registry common.Address,
	tokenID []byte,
	nftUniqueProof, readAccessProof bool) (proofs []*proofspb.Proof, err error) {

	dataLeaves, err := e.getDataLeaves()
	if err != nil {
		return nil, err
	}

	return e.CoreDocument.CreateNFTProofs(
		e.DocumentType(),
		dataLeaves,
		account, registry, tokenID, nftUniqueProof, readAccessProof)
}

// CreateProofs generates proofs for given fields.
func (e *Entity) CreateProofs(fields []string) (proofs []*proofspb.Proof, err error) {
	dataLeaves, err := e.getDataLeaves()
	if err != nil {
		return nil, errors.New("createProofs error %v", err)
	}

	return e.CoreDocument.CreateProofs(e.DocumentType(), dataLeaves, fields)
}

// DocumentType returns the entity document type.
func (*Entity) DocumentType() string {
	return documenttypes.EntityDataTypeUrl
}

// PrepareNewVersion prepares new version from the old entity.
func (e *Entity) PrepareNewVersion(old documents.Model, data *cliententitypb.EntityData, collaborators documents.CollaboratorsAccess, attrs map[documents.AttrKey]documents.Attribute) error {
	err := e.initEntityFromData(data)
	if err != nil {
		return err
	}

	oldCD := old.(*Entity).CoreDocument
	e.CoreDocument, err = oldCD.PrepareNewVersion(compactPrefix(), collaborators, attrs)
	if err != nil {
		return err
	}

	return nil
}

// AddNFT adds NFT to the Entity.
func (e *Entity) AddNFT(grantReadAccess bool, registry common.Address, tokenID []byte) error {
	cd, err := e.CoreDocument.AddNFT(grantReadAccess, registry, tokenID)
	if err != nil {
		return err
	}

	e.CoreDocument = cd
	return nil
}

// CalculateDataRoot calculates the document data root of the document.
func (e *Entity) CalculateDataRoot() ([]byte, error) {
	dataLeaves, err := e.getDataLeaves()
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDataTree, err)
	}
	return e.CoreDocument.CalculateDataRoot(e.DocumentType(), dataLeaves)
}

// CalculateDocumentRoot calculates the document root
func (e *Entity) CalculateDocumentRoot() ([]byte, error) {
	dataLeaves, err := e.getDataLeaves()
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDataTree, err)
	}
	return e.CoreDocument.CalculateDocumentRoot(e.DocumentType(), dataLeaves)
}

// DocumentRootTree creates and returns the document root tree
func (e *Entity) DocumentRootTree() (tree *proofs.DocumentTree, err error) {
	dataLeaves, err := e.getDataLeaves()
	if err != nil {
		return nil, errors.NewTypedError(documents.ErrDataTree, err)
	}
	return e.CoreDocument.DocumentRootTree(e.DocumentType(), dataLeaves)
}

// CollaboratorCanUpdate checks if the collaborator can update the document.
func (e *Entity) CollaboratorCanUpdate(updated documents.Model, collaborator identity.DID) error {
	newEntity, ok := updated.(*Entity)
	if !ok {
		return errors.NewTypedError(documents.ErrDocumentInvalidType, errors.New("expecting an entity but got %T", updated))
	}

	// check the core document changes
	err := e.CoreDocument.CollaboratorCanUpdate(newEntity.CoreDocument, collaborator, e.DocumentType())
	if err != nil {
		return err
	}

	// check entity specific changes
	oldTree, err := e.getDataTree()
	if err != nil {
		return err
	}

	newTree, err := newEntity.getDataTree()
	if err != nil {
		return err
	}

	rules := e.CoreDocument.TransitionRulesFor(collaborator)
	cf := documents.GetChangedFields(oldTree, newTree)
	return documents.ValidateTransitions(rules, cf)
}

// AddAttributes adds attributes to the Entity model.
func (e *Entity) AddAttributes(ca documents.CollaboratorsAccess, prepareNewVersion bool, attrs ...documents.Attribute) error {
	ncd, err := e.CoreDocument.AddAttributes(ca, prepareNewVersion, compactPrefix(), attrs...)
	if err != nil {
		return errors.NewTypedError(documents.ErrCDAttribute, err)
	}

	e.CoreDocument = ncd
	return nil
}

// DeleteAttribute deletes the attribute from the model.
func (e *Entity) DeleteAttribute(key documents.AttrKey, prepareNewVersion bool) error {
	ncd, err := e.CoreDocument.DeleteAttribute(key, prepareNewVersion, compactPrefix())
	if err != nil {
		return errors.NewTypedError(documents.ErrCDAttribute, err)
	}

	e.CoreDocument = ncd
	return nil
}

// GetData returns entity data
func (e *Entity) GetData() interface{} {
	return e.Data
}

func isOnlyOneSet(methods ...interface{}) error {
	var isSet bool
	for _, method := range methods {
		mv := reflect.ValueOf(method)
		if mv.IsNil() {
			continue
		}

		if isSet {
			return ErrMultiplePaymentMethodsSet
		}

		isSet = true
	}

	if !isSet {
		return ErrNoPaymentMethodSet
	}

	return nil
}

// loadData unmarshals json blob to Data.
// Only one of the payment method has to be set.
// errors out if multiple payment methods are set or none is set.
func (e *Entity) loadData(data []byte) error {
	var d Data
	err := json.Unmarshal(data, &d)
	if err != nil {
		return err
	}

	pds := d.PaymentDetails
	for _, pd := range pds {
		err = isOnlyOneSet(pd.BankPaymentMethod, pd.CryptoPaymentMethod, pd.OtherPaymentMethod)
		if err != nil {
			return err
		}
	}

	e.Data = d
	return nil
}

// unpackFromCreatePayload unpacks the entity data from the Payload.
func (e *Entity) unpackFromCreatePayload(did identity.DID, payload documents.CreatePayload) error {
	if err := e.loadData(payload.Data); err != nil {
		return errors.NewTypedError(ErrEntityInvalidData, err)
	}

	payload.Collaborators.ReadWriteCollaborators = append(payload.Collaborators.ReadWriteCollaborators, did)
	cd, err := documents.NewCoreDocument(compactPrefix(), payload.Collaborators, payload.Attributes)
	if err != nil {
		return errors.NewTypedError(documents.ErrCDCreate, err)
	}

	e.CoreDocument = cd
	return nil
}

// unpackFromUpdatePayload unpacks the update payload and prepares a new version.
func (e *Entity) unpackFromUpdatePayload(old *Entity, payload documents.UpdatePayload) error {
	if err := e.loadData(payload.Data); err != nil {
		return errors.NewTypedError(ErrEntityInvalidData, err)
	}

	ncd, err := old.CoreDocument.PrepareNewVersion(compactPrefix(), payload.Collaborators, payload.Attributes)
	if err != nil {
		return err
	}

	e.CoreDocument = ncd
	return nil
}

// Scheme returns the entity scheme.
func (e *Entity) Scheme() string {
	return scheme
}
