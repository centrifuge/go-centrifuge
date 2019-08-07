package v2

import (
	"net/http"

	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/httpapi/coreapi"
	"github.com/centrifuge/go-centrifuge/jobs"
	"github.com/centrifuge/go-centrifuge/utils/byteutils"
	"github.com/centrifuge/go-centrifuge/utils/httputils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// DocumentRequest is an alias to coreapi Document request.
// Aliased here to fix the swagger generation issues.
type DocumentRequest = coreapi.CreateDocumentRequest

// CreateDocumentRequest defines the payload for creating documents.
type CreateDocumentRequest struct {
	DocumentRequest
	DocumentID byteutils.OptionalHex `json:"document_id" swaggertype:"primitive,string"` // if provided, creates the next version of the document.
}

// UpdateDocumentRequest defines the payload to patch an existing document.
type UpdateDocumentRequest struct {
	DocumentRequest
}

// CreateDocument creates a document.
// @summary Creates a new document.
// @description Creates a new document.
// @id create_document_v2
// @tags Documents
// @accept json
// @param authorization header string true "Hex encoded centrifuge ID of the account for the intended API action"
// @param body body v2.CreateDocumentRequest true "Document Create request"
// @produce json
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Failure 403 {object} httputils.HTTPError
// @success 201 {object} coreapi.DocumentResponse
// @router /v2/documents [post]
func (h handler) CreateDocument(w http.ResponseWriter, r *http.Request) {
	var err error
	var code int
	defer httputils.RespondIfError(&code, &err, w, r)

	ctx := r.Context()
	var req CreateDocumentRequest
	err = unmarshalBody(r, &req)
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		return
	}

	payload, err := toDocumentsPayload(req.DocumentRequest, req.DocumentID.Bytes())
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		return
	}

	doc, err := h.srv.CreateDocument(ctx, payload)
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		return
	}

	resp, err := toDocumentResponse(doc, h.srv.tokenRegistry, jobs.NilJobID())
	if err != nil {
		code = http.StatusInternalServerError
		log.Error(err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

// Update updates a pending document.
// @summary Updates a pending document.
// @description Updates a pending document.
// @id update_document_v2
// @tags Documents
// @accept json
// @param authorization header string true "Hex encoded centrifuge ID of the account for the intended API action"
// @param body body v2.UpdateDocumentRequest true "Document Update request"
// @param document_id path string true "Document Identifier"
// @produce json
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Failure 403 {object} httputils.HTTPError
// @success 200 {object} coreapi.DocumentResponse
// @router /v2/documents/{document_id} [patch]
func (h handler) UpdateDocument(w http.ResponseWriter, r *http.Request) {
	var err error
	var code int
	defer httputils.RespondIfError(&code, &err, w, r)

	docID, err := hexutil.Decode(chi.URLParam(r, coreapi.DocumentIDParam))
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		err = coreapi.ErrInvalidDocumentID
		return
	}

	ctx := r.Context()
	var req UpdateDocumentRequest
	err = unmarshalBody(r, &req)
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		return
	}

	payload, err := toDocumentsPayload(req.DocumentRequest, docID)
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		return
	}

	doc, err := h.srv.UpdateDocument(ctx, payload)
	if err != nil {
		code = http.StatusNotFound
		err = coreapi.ErrDocumentNotFound
		return
	}

	resp, err := toDocumentResponse(doc, h.srv.tokenRegistry, jobs.NilJobID())
	if err != nil {
		code = http.StatusInternalServerError
		log.Error(err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

// Commit creates a document.
// @summary Commits a pending document.
// @description Commits a pending document.
// @id commit_document_v2
// @tags Documents
// @accept json
// @param authorization header string true "Hex encoded centrifuge ID of the account for the intended API action"
// @param document_id path string true "Document Identifier"
// @produce json
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @Failure 403 {object} httputils.HTTPError
// @success 202 {object} coreapi.DocumentResponse
// @router /v2/documents/{document_id}/commit [post]
func (h handler) Commit(w http.ResponseWriter, r *http.Request) {
	var err error
	var code int
	defer httputils.RespondIfError(&code, &err, w, r)

	docID, err := hexutil.Decode(chi.URLParam(r, coreapi.DocumentIDParam))
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		err = coreapi.ErrInvalidDocumentID
		return
	}

	ctx := r.Context()
	doc, jobID, err := h.srv.Commit(ctx, docID)
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		return
	}

	resp, err := toDocumentResponse(doc, h.srv.tokenRegistry, jobID)
	if err != nil {
		code = http.StatusInternalServerError
		log.Error(err)
		return
	}

	render.Status(r, http.StatusAccepted)
	render.JSON(w, r, resp)
}

func (h handler) getDocumentWithStatus(w http.ResponseWriter, r *http.Request, st documents.Status) {
	var err error
	var code int
	defer httputils.RespondIfError(&code, &err, w, r)

	docID, err := hexutil.Decode(chi.URLParam(r, coreapi.DocumentIDParam))
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		err = coreapi.ErrInvalidDocumentID
		return
	}

	ctx := r.Context()
	doc, err := h.srv.GetDocument(ctx, docID, st)
	if err != nil {
		code = http.StatusNotFound
		log.Error(err)
		err = coreapi.ErrDocumentNotFound
		return
	}

	resp, err := toDocumentResponse(doc, h.srv.tokenRegistry, jobs.NilJobID())
	if err != nil {
		code = http.StatusInternalServerError
		log.Error(err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

// GetPendingDocument returns the pending document associated with docID.
// @summary Returns the pending document associated with docID.
// @description Returns the pending document associated with docID.
// @id get_pending_document
// @tags Documents
// @param authorization header string true "Hex encoded centrifuge ID of the account for the intended API action"
// @param document_id path string true "Document Identifier"
// @produce json
// @Failure 403 {object} httputils.HTTPError
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @success 200 {object} coreapi.DocumentResponse
// @router /v2/documents/{document_id}/pending [get]
func (h handler) GetPendingDocument(w http.ResponseWriter, r *http.Request) {
	h.getDocumentWithStatus(w, r, documents.Pending)
}

// GetCommittedDocument returns the latest committed document associated with docID.
// @summary Returns the latest committed document associated with docID.
// @description Returns the latest committed document associated with docID.
// @id get_committed_document
// @tags Documents
// @param authorization header string true "Hex encoded centrifuge ID of the account for the intended API action"
// @param document_id path string true "Document Identifier"
// @produce json
// @Failure 403 {object} httputils.HTTPError
// @Failure 400 {object} httputils.HTTPError
// @Failure 404 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @success 200 {object} coreapi.DocumentResponse
// @router /v2/documents/{document_id}/committed [get]
func (h handler) GetCommittedDocument(w http.ResponseWriter, r *http.Request) {
	h.getDocumentWithStatus(w, r, documents.Committed)
}
