package router

import (
	"github.com/go-chi/chi/v5"

	"workspace.dev/shared/go/logger"

	"workspace.dev/services/apis/bookapi/app/resource/book"
	"workspace.dev/services/apis/bookapi/app/resource/health"
)

func New(l *logger.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/livez", health.Read)

	r.Route("/v1", func(r chi.Router) {
		bookAPI := book.New(l)
		r.Get("/books", bookAPI.List)
		r.Post("/books", bookAPI.Create)
		r.Get("/books/{id}", bookAPI.Read)
		r.Put("/books/{id}", bookAPI.Update)
		r.Delete("/books/{id}", bookAPI.Delete)
	})

	return r
}
