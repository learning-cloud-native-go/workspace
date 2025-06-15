package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"workspace.dev/services/apis/bookapi/app/resource/book"
	"workspace.dev/services/apis/bookapi/app/resource/health"
	"workspace.dev/services/apis/bookapi/app/router/middleware"
	mrl "workspace.dev/services/apis/bookapi/app/router/middleware/requestlog"
	"workspace.dev/shared/go/logger"
)

func New(l *logger.Logger, v *validator.Validate, db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/livez", health.Read)

	r.Route("/v1", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.ContentTypeJSON)

		bookAPI := book.New(l, v, db)
		r.Method(http.MethodGet, "/books", mrl.NewHandler(bookAPI.List, l))
		r.Method(http.MethodPost, "/books", mrl.NewHandler(bookAPI.Create, l))
		r.Method(http.MethodGet, "/books/{id}", mrl.NewHandler(bookAPI.Read, l))
		r.Method(http.MethodPut, "/books/{id}", mrl.NewHandler(bookAPI.Update, l))
		r.Method(http.MethodDelete, "/books/{id}", mrl.NewHandler(bookAPI.Delete, l))
	})

	return r
}
