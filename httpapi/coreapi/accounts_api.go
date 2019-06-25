package coreapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/utils/httputils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// ErrAccountIDInvalid is a sentinel error for invalid account IDs.
const ErrAccountIDInvalid = errors.Error("account ID is invalid")

// SignPayload signs the payload and returns the signature.
// @summary Signs and returns the signature of the Payload.
// @description Signs and returns the signature of the Payload.
// @id account_sign
// @tags Accounts
// @param account_id path string true "Account ID"
// @param body body coreapi.SignRequest true "Sign request"
// @produce json
// @Failure 400 {object} httputils.HTTPError
// @Failure 500 {object} httputils.HTTPError
// @success 200 {object} coreapi.SignResponse
// @router /v1/accounts/{account_id}/sign [post]
func (h handler) SignPayload(w http.ResponseWriter, r *http.Request) {
	var err error
	var code int
	defer httputils.RespondIfError(&code, &err, w, r)

	accID, err := hexutil.Decode(chi.URLParam(r, accountIDParam))
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		err = ErrAccountIDInvalid
		return
	}

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		code = http.StatusInternalServerError
		log.Error(err)
		return
	}

	var payload SignRequest
	err = json.Unmarshal(d, &payload)
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		return
	}

	sig, err := h.srv.SignPayload(accID, payload.Payload)
	if err != nil {
		code = http.StatusBadRequest
		log.Error(err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, SignResponse{
		Payload:   payload.Payload,
		PublicKey: sig.PublicKey,
		Signature: sig.Signature,
		SignerID:  sig.SignerId,
	})
}
