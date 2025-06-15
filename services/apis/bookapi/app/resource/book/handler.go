package book

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"

	e "workspace.dev/shared/go/errors"
	"workspace.dev/shared/go/logger"
	mb "workspace.dev/shared/go/models/book"
	rb "workspace.dev/shared/go/repositories/book"
	v "workspace.dev/shared/go/validator"
)

type API struct {
	logger       *logger.Logger
	validator    *validator.Validate
	repositories *repositories
}

type repositories struct {
	book *rb.Repository
}

func New(logger *logger.Logger, validator *validator.Validate, db *gorm.DB) *API {
	return &API{
		logger:    logger,
		validator: validator,
		repositories: &repositories{
			book: rb.New(db),
		},
	}
}

// List godoc
//
//	@summary		List books
//	@description	List books
//	@tags			books
//	@accept			json
//	@produce		json
//	@success		200	{array}		mb.DTO
//	@failure		500	{object}	e.Error
//	@router			/books [get]
func (a *API) List(w http.ResponseWriter, r *http.Request) {
	books, err := a.repositories.book.List()
	if err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	if len(books) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(books.ToDto()); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

// Create godoc
//
//	@summary		Create book
//	@description	Create book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			body	body	mb.Form	true	"Book form"
//	@success		201
//	@failure		400	{object}	e.Error
//	@failure		422	{object}	e.Errors
//	@failure		500	{object}	e.Error
//	@router			/books [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	form := &mb.Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	if err := a.validator.Struct(form); err != nil {
		respBody, err := json.Marshal(v.ToErrResponse(err))
		if err != nil {
			a.logger.Error().Err(err).Msg("")
			e.ServerError(w, e.RespJSONEncodeFailure)
			return
		}

		e.ValidationErrors(w, respBody)
		return
	}

	newBook := form.ToModel()
	newBook.ID = uuid.New()

	book, err := a.repositories.book.Create(newBook)
	if err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespDBDataInsertFailure)
		return
	}

	a.logger.Info().Str("id", book.ID.String()).Msg("new book created")
	w.WriteHeader(http.StatusCreated)
}

// Read godoc
//
//	@summary		Read book
//	@description	Read book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			id	path		string	true	"Book ID"
//	@success		200	{object}	mb.DTO
//	@failure		400	{object}	e.Error
//	@failure		404
//	@failure		500	{object}	e.Error
//	@router			/books/{id} [get]
func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	book, err := a.repositories.book.Read(id)
	if err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespDBDataAccessFailure)
		return
	}

	if book == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	dto := book.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespJSONEncodeFailure)
		return
	}
}

// Update godoc
//
//	@summary		Update book
//	@description	Update book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			id		path	string	true	"Book ID"
//	@param			body	body	mb.Form	true	"Book form"
//	@success		200
//	@failure		400	{object}	e.Error
//	@failure		404
//	@failure		422	{object}	e.Errors
//	@failure		500	{object}	e.Error
//	@router			/books/{id} [put]
func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	form := &mb.Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.BadRequest(w, e.RespJSONDecodeFailure)
		return
	}

	if err := a.validator.Struct(form); err != nil {
		respBody, err := json.Marshal(v.ToErrResponse(err))
		if err != nil {
			a.logger.Error().Err(err).Msg("")
			e.ServerError(w, e.RespJSONEncodeFailure)
			return
		}

		e.ValidationErrors(w, respBody)
		return
	}

	book := form.ToModel()
	book.ID = id

	rows, err := a.repositories.book.Update(book)
	if err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespDBDataUpdateFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.logger.Info().Str("id", id.String()).Msg("book updated")
}

// Delete godoc
//
//	@summary		Delete book
//	@description	Delete book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			id	path	string	true	"Book ID"
//	@success		200
//	@failure		400	{object}	e.Error
//	@failure		404
//	@failure		500	{object}	e.Error
//	@router			/books/{id} [delete]
func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.RespInvalidURLParamID)
		return
	}

	rows, err := a.repositories.book.Delete(id)
	if err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.RespDBDataRemoveFailure)
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.logger.Info().Str("id", id.String()).Msg("book deleted")
}
