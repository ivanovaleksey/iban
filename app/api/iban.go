package api

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/ivanovaleksey/iban/app/responses"
	"net/http"
)

type ibanSrv interface {
	Validate(string) error
}

var errEmptyParam = errors.New("empty parameter")

// swagger:parameters validateIBAN
type ValidateIBANParams struct {
	// IBAN to be validated
	// required: true
	// in: query
	IBAN string `json:"iban"`
}

// swagger:route GET /iban/validate validateIBAN
//
// Validate given IBAN.
//
//     Responses:
//       default: errorResponse
//       200: validateIBANResponse
func (a *API) ValidateIBAN() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ibanParam := r.URL.Query().Get("iban")

		if ibanParam == "" {
			render.Render(w, r, responses.ErrBadRequest(errEmptyParam))
			return
		}

		err := a.ibanSrv.Validate(ibanParam)
		if err != nil {
			render.Render(w, r, responses.ErrBadRequest(err))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
