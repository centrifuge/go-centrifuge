package coredocument

import (
	"crypto/sha256"
	"fmt"

	"github.com/CentrifugeInc/go-centrifuge/centrifuge/code"

	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/coredocument"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/errors"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/tools"
	"github.com/centrifuge/precise-proofs/proofs"
)

// GetDataProofHashes returns the hashes needed to create a proof from DataRoot to SigningRoot. This method is used
// to create field proofs
func GetDataProofHashes(document *coredocumentpb.CoreDocument) (hashes [][]byte, err error) {
	tree, err := GetDocumentSigningTree(document)
	if err != nil {
		return
	}

	signingProof, err := tree.CreateProof("data_root")
	if err != nil {
		return
	}

	tree, err = GetDocumentRootTree(document)
	if err != nil {
		return
	}
	rootProof, err := tree.CreateProof("signing_root")
	if err != nil {
		return
	}
	return append(signingProof.SortedHashes, rootProof.SortedHashes...), err
}

func CalculateSigningRoot(document *coredocumentpb.CoreDocument) error {
	valid, errMsg, errs := Validate(document) // TODO: Validation
	if !valid {
		return errors.NewWithErrors(code.DocumentInvalid, errMsg, errs)
	}

	tree, err := GetDocumentSigningTree(document)
	if err != nil {
		return err
	}
	document.SigningRoot = tree.RootHash()
	return nil
}

func CalculateDocumentRoot(document *coredocumentpb.CoreDocument) error {
	if len(document.SigningRoot) != 32 {
		return errors.New(code.DocumentInvalid, "signing root invalid")
	}
	tree, err := GetDocumentRootTree(document)
	if err != nil {
		return err
	}
	document.DocumentRoot = tree.RootHash()
	return nil
}

// GetDocumentRootTree returns the merkle tree for the document root
func GetDocumentRootTree(document *coredocumentpb.CoreDocument) (tree *proofs.DocumentTree, err error) {
	h := sha256.New()
	t := proofs.NewDocumentTree(proofs.TreeOptions{EnableHashSorting: true, Hash: h})
	tree = &t

	// The first leave added is the signing_root
	err = tree.AddLeaf(proofs.LeafNode{Hash: document.SigningRoot, Hashed: true, Property: "signing_root"})
	if err != nil {
		return nil, err
	}
	// For every signature we create a LeafNode
	// TODO: we should modify this to use the proper message flattener once precise proofs is modified to support it
	sigLeafList := make([]proofs.LeafNode, len(document.Signatures)+1)
	sigLengthNode := proofs.LeafNode{
		Property: "signatures.length",
		Salt:     make([]byte, 32),
		Value:    fmt.Sprintf("%d", len(document.Signatures)),
	}
	sigLengthNode.HashNode(h)
	sigLeafList[0] = sigLengthNode
	for i, sig := range document.Signatures {
		payload := sha256.Sum256(append(sig.EntityId, append(sig.PublicKey, sig.Signature...)...))
		leaf := proofs.LeafNode{
			Hash:     payload[:],
			Hashed:   true,
			Property: fmt.Sprintf("signatures[%d]", i),
		}
		leaf.HashNode(h)
		sigLeafList[i+1] = leaf
	}
	err = tree.AddLeaves(sigLeafList)
	if err != nil {
		return nil, err
	}
	err = tree.Generate()
	if err != nil {
		return nil, err
	}
	return tree, nil
}

// GetDocumentSigningTree returns the merkle tree for the signing root
func GetDocumentSigningTree(document *coredocumentpb.CoreDocument) (tree *proofs.DocumentTree, err error) {
	t := proofs.NewDocumentTree(proofs.TreeOptions{EnableHashSorting: true, Hash: sha256.New()})
	tree = &t
	err = tree.AddLeavesFromDocument(document, document.CoredocumentSalts)
	if err != nil {
		return nil, err
	}
	err = tree.Generate()
	if err != nil {
		return nil, err
	}
	return tree, nil
}

// Validate checks that all required fields are set before doing any processing with core document
func Validate(document *coredocumentpb.CoreDocument) (valid bool, errMsg string, errs map[string]string) {
	if document == nil {
		return false, errors.NilDocument, nil
	}

	errs = make(map[string]string)

	if tools.IsEmptyByteSlice(document.DocumentIdentifier) {
		errs["cd_identifier"] = errors.RequiredField
	}

	// TODO(ved): where do we fill these
	//if tools.IsEmptyByteSlice(document.DocumentRoot) {
	//	errs["cd_root"] = errors.RequiredField
	//}

	if tools.IsEmptyByteSlice(document.CurrentIdentifier) {
		errs["cd_current_identifier"] = errors.RequiredField
	}

	if tools.IsEmptyByteSlice(document.NextIdentifier) {
		errs["cd_next_identifier"] = errors.RequiredField
	}

	if tools.IsEmptyByteSlice(document.DataRoot) {
		errs["cd_data_root"] = errors.RequiredField
	}

	// double check the identifiers
	isSameBytes := tools.IsSameByteSlice

	// Problem (re-using an old identifier for NextIdentifier): CurrentIdentifier or DocumentIdentifier same as NextIdentifier
	if isSameBytes(document.NextIdentifier, document.DocumentIdentifier) ||
		isSameBytes(document.NextIdentifier, document.CurrentIdentifier) {
		errs["cd_overall"] = errors.IdentifierReUsed
	}

	// lets not do verbose check like earlier since these will be
	// generated by us mostly
	salts := document.CoredocumentSalts
	if salts == nil ||
		!tools.CheckMultiple32BytesFilled(
			salts.CurrentIdentifier,
			salts.NextIdentifier,
			salts.DocumentIdentifier,
			salts.PreviousRoot) {
		errs["cd_salts"] = errors.RequiredField
	}

	if len(errs) < 1 {
		return true, "", nil
	}

	return false, "Invalid CoreDocument", errs
}

// FillIdentifiers fills in missing identifiers for the given CoreDocument.
// It does checks on document consistency (e.g. re-using an old identifier).
// In the case of an error, it returns the error and an empty CoreDocument.
func FillIdentifiers(document coredocumentpb.CoreDocument) (coredocumentpb.CoreDocument, error) {
	isEmptyId := tools.IsEmptyByteSlice

	// check if the document identifier is empty
	if !isEmptyId(document.DocumentIdentifier) {
		// check and fill current and next identifiers
		if isEmptyId(document.CurrentIdentifier) {
			document.CurrentIdentifier = document.DocumentIdentifier
		}

		if isEmptyId(document.NextIdentifier) {
			document.NextIdentifier = tools.RandomSlice(32)
		}

		return document, nil
	}

	// check if current and next identifier are empty
	if !isEmptyId(document.CurrentIdentifier) {
		return document, fmt.Errorf("no DocumentIdentifier but has CurrentIdentifier")
	}

	// check if the next identifier is empty
	if !isEmptyId(document.NextIdentifier) {
		return document, fmt.Errorf("no CurrentIdentifier but has NextIdentifier")
	}

	// fill the identifiers
	document.DocumentIdentifier = tools.RandomSlice(32)
	document.CurrentIdentifier = document.DocumentIdentifier
	document.NextIdentifier = tools.RandomSlice(32)
	return document, nil
}

// New returns a new core document from the proto message
func New() *coredocumentpb.CoreDocument {
	doc, _ := FillIdentifiers(coredocumentpb.CoreDocument{})
	salts := &coredocumentpb.CoreDocumentSalts{}
	proofs.FillSalts(&doc, salts)
	doc.CoredocumentSalts = salts
	return &doc
}
