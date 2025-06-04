package book

import (
	"net/http"

	"logger"
)

type API struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *API {
	return &API{
		logger: logger,
	}
}

func (a *API) List(w http.ResponseWriter, r *http.Request) {}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {}

func (a *API) Read(w http.ResponseWriter, r *http.Request) {}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {}
