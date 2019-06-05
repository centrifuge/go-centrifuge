package documents

import (
	"context"
	"math/big"
	"time"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/storage"
	"github.com/centrifuge/precise-proofs/proofs"
	"github.com/centrifuge/precise-proofs/proofs/proto"
	"github.com/ethereum/go-ethereum/common"
)

// Model is an interface to abstract away model specificness like invoice or purchaseOrder
// The interface can cast into the type specified by the model if required
// It should only handle protocol-level document actions
type Model interface {
	storage.Model

	// ID returns the document identifier
	ID() []byte

	// CurrentVersion returns the current version identifier of the document
	CurrentVersion() []byte

	// CurrentVersionPreimage returns the current version pre-image of the document. This is intended to hide the next version of an updated version of the document.
	CurrentVersionPreimage() []byte

	// PreviousVersion returns the previous version identifier of the document
	PreviousVersion() []byte

	// NextVersion returns the next version identifier of the document.
	NextVersion() []byte

	// PackCoreDocument packs the implementing document into a core document
	// Should only be called when the document is about to be put on wire.
	PackCoreDocument() (coredocumentpb.CoreDocument, error)

	// UnpackCoreDocument takes a core document protobuf and loads the data into the model.
	UnpackCoreDocument(cd coredocumentpb.CoreDocument) error

	// DocumentType returns the type of the document
	DocumentType() string

	// CalculateDataRoot calculates the data root of the model.
	CalculateDataRoot() ([]byte, error)

	// CalculateDocumentDataRoot calculates the document data root of the model.
	CalculateDocumentDataRoot() ([]byte, error)

	// CalculateDocumentRoot returns the document root of the model.
	CalculateDocumentRoot() ([]byte, error)

	// CalculateSignaturesRoot returns signatures root of the model.
	CalculateSignaturesRoot() ([]byte, error)

	// DocumentRootTree returns the document root tree
	DocumentRootTree() (tree *proofs.DocumentTree, err error)

	// AppendSignatures appends the signatures to the model.
	AppendSignatures(signatures ...*coredocumentpb.Signature)

	// Signatures returns a copy of the signatures on the document
	Signatures() []coredocumentpb.Signature

	// CreateProofs creates precise-proofs for given fields
	CreateProofs(fields []string) (proofs []*proofspb.Proof, err error)

	// CreateNFTProofs creates NFT proofs for minting.
	CreateNFTProofs(
		account identity.DID,
		registry common.Address,
		tokenID []byte,
		nftUniqueProof, readAccessProof bool) (proofs []*proofspb.Proof, err error)

	// IsNFTMinted checks if there is any NFT minted for the registry given
	IsNFTMinted(tr TokenRegistry, registry common.Address) bool

	// AddNFT adds an NFT to the document.
	// Note: The document should be anchored after successfully adding the NFT.
	AddNFT(grantReadAccess bool, registry common.Address, tokenID []byte) error

	// NFTs returns the list of NFTs created for this model
	NFTs() []*coredocumentpb.NFT

	// GetCollaborators returns the collaborators of this document.
	// filter ids should not be returned
	// Note: returns all the collaborators with Read and Read_Sign permission
	GetCollaborators(filterIDs ...identity.DID) (CollaboratorsAccess, error)

	// GetSignerCollaborators works like GetCollaborators except it returns only those with Read_Sign permission.
	GetSignerCollaborators(filterIDs ...identity.DID) ([]identity.DID, error)

	// AccountCanRead returns true if the account can read the document
	AccountCanRead(account identity.DID) bool

	// NFTOwnerCanRead returns error if the NFT cannot read the document.
	NFTOwnerCanRead(tokenRegistry TokenRegistry, registry common.Address, tokenID []byte, account identity.DID) error

	// ATGranteeCanRead returns error if the access token grantee cannot read the document.
	ATGranteeCanRead(ctx context.Context, docSrv Service, idSrv identity.Service, tokenID, docID []byte, grantee identity.DID) (err error)

	// AddUpdateLog adds a log to the model to persist an update related meta data such as author
	AddUpdateLog(account identity.DID) error

	// Author is the author of the document version represented by the model
	Author() (identity.DID, error)

	// Timestamp is the time of update in UTC of the document version represented by the model
	Timestamp() (time.Time, error)

	// CollaboratorCanUpdate returns an error if indicated identity does not have the capacity to update the document.
	CollaboratorCanUpdate(updated Model, collaborator identity.DID) error

	// IsDIDCollaborator returns true if the did is a collaborator of the document
	IsDIDCollaborator(did identity.DID) (bool, error)

	// AddAttributes adds a custom attribute to the model with the given value. If an attribute with the given name already exists, it's updated.
	AddAttributes(ca CollaboratorsAccess, prepareNewVersion bool, attrs ...Attribute) error

	// GetAttribute gets the attribute with the given name from the model, it returns a non-nil error if the attribute doesn't exist or can't be retrieved.
	GetAttribute(key AttrKey) (Attribute, error)

	// GetAttributes returns all the attributes in the current document
	GetAttributes() []Attribute

	// DeleteAttribute deletes a custom attribute from the model
	DeleteAttribute(key AttrKey, prepareNewVersion bool) error

	// AttributeExists checks if the attribute with the key exists
	AttributeExists(key AttrKey) bool

	// GetAccessTokens returns the access tokens of a core document
	GetAccessTokens() ([]*coredocumentpb.AccessToken, error)

	// SetUsedAnchorRepoAddress sets the anchor repository address to which document is anchored to.
	SetUsedAnchorRepoAddress(addr common.Address)

	// AnchorRepoAddress returns the used anchor repo address to which document is/will be anchored to.
	AnchorRepoAddress() common.Address
}

// TokenRegistry defines NFT related functions.
type TokenRegistry interface {
	// OwnerOf to retrieve owner of the tokenID
	OwnerOf(registry common.Address, tokenID []byte) (common.Address, error)

	// CurrentIndexOfToken get the current index of the token
	CurrentIndexOfToken(registry common.Address, tokenID []byte) (*big.Int, error)
}

// CreatePayload holds the scheme, CollaboratorsAccess, Attributes, and Data of the document.
type CreatePayload struct {
	Scheme        string
	Collaborators CollaboratorsAccess
	Attributes    map[AttrKey]Attribute
	Data          []byte
}

// UpdatePayload holds the scheme, CollaboratorsAccess, Attributes, Data and document identifier.
type UpdatePayload struct {
	CreatePayload
	DocumentID []byte
}
