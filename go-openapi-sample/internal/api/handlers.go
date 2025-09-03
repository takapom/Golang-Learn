package api

import (
	"github.com/go-chi/chi/v5"
)

type StrictServerOptions struct{}

func NewStrictHandler(server ServerInterface, options *StrictServerOptions) ServerInterface {
	return server
}

func RegisterHandlers(router chi.Router, si ServerInterface) {
	HandlerFromMux(si, router)
}