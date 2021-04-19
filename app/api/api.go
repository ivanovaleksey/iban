// Package api API
//
// IBAN service methods description
//
// BasePath: /api
// Version: 0.1
// Produces:
//  - application/json
// Schemes: http, https
// swagger:meta
package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/ivanovaleksey/iban/app/services/iban"
	"net/http"
)

type API struct {
	router chi.Router

	ibanSrv ibanSrv
}

func New(opts ...Option) *API {
	a := &API{
		ibanSrv: iban.Service{},
	}
	for _, opt := range opts {
		opt(a)
	}
	a.initRouter()
	return a
}

func (a *API) initRouter() {
	r := chi.NewRouter()

	r.Route("/docs", func(r chi.Router) {
		r.Handle("/*", http.StripPrefix("/docs", http.FileServer(http.Dir("docs"))))
	})
	r.Route("/api", func(r chi.Router) {
		r.Route("/iban", func(r chi.Router) {
			r.Get("/validate", a.ValidateIBAN())
		})
	})

	a.router = r
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
