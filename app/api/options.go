package api

type Option func(*API)

func WithIBANSrv(srv ibanSrv) Option {
	return func(api *API) {
		api.ibanSrv = srv
	}
}
