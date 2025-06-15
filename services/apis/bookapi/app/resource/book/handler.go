package book

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"workspace.dev/shared/go/logger"
	model "workspace.dev/shared/go/models/book"
	repo "workspace.dev/shared/go/repositories/book"
)

type API struct {
	logger       *logger.Logger
	repositories *repositories
}

type repositories struct {
	book *repo.Repository
}

func New(logger *logger.Logger, db *gorm.DB) *API {
	return &API{
		logger: logger,
		repositories: &repositories{
			book: repo.New(db),
		},
	}
}

func (a *API) List(w http.ResponseWriter, r *http.Request) {
	books, err := a.repositories.book.List()
	if err != nil {
		// handle later
		return
	}

	if len(books) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(books.ToDto()); err != nil {
		// handle later
		return
	}
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	form := &model.Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		// handle later
		return
	}

	newBook := form.ToModel()
	newBook.ID = uuid.New()

	_, err := a.repositories.book.Create(newBook)
	if err != nil {
		// handle later
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// handle later
		return
	}

	book, err := a.repositories.book.Read(id)
	if err != nil {
		// handle later
		return
	}

	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dto := book.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		// handle later
		return
	}
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// handle later
		return
	}

	form := &model.Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		// handle later
		return
	}

	book := form.ToModel()
	book.ID = id

	rows, err := a.repositories.book.Update(book)
	if err != nil {
		// handle later
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// handle later
		return
	}

	rows, err := a.repositories.book.Delete(id)
	if err != nil {
		// handle later
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
