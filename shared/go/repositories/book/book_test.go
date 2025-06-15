package book_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"

	m "workspace.dev/shared/go/models/book"
	r "workspace.dev/shared/go/repositories/book"
	rTest "workspace.dev/shared/go/repositories/test"
	uTest "workspace.dev/shared/go/utils/test"
)

func TestRepository_List(t *testing.T) {
	t.Parallel()

	db, mock, err := rTest.NewMockDB()
	uTest.NoError(t, err)

	repo := r.NewRepository(db)

	mockRows := sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(uuid.New(), "Book1", "Author1").
		AddRow(uuid.New(), "Book2", "Author2")

	mock.ExpectQuery("^SELECT (.+) FROM \"books\"").WillReturnRows(mockRows)

	books, err := repo.List()
	uTest.NoError(t, err)
	uTest.Equal(t, len(books), 2)
}

func TestRepository_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := rTest.NewMockDB()
	uTest.NoError(t, err)

	repo := r.NewRepository(db)

	id := uuid.New()
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO \"books\" ").
		WithArgs(id, "Title", "Author", rTest.AnyTime{}, "", "", rTest.AnyTime{}, rTest.AnyTime{}, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	book := &m.Book{ID: id, Title: "Title", Author: "Author", PublishedDate: time.Now()}
	_, err = repo.Create(book)
	uTest.NoError(t, err)
}

func TestRepository_Read(t *testing.T) {
	t.Parallel()

	db, mock, err := rTest.NewMockDB()
	uTest.NoError(t, err)

	repo := r.NewRepository(db)

	id := uuid.New()
	mockRows := sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(id, "Book1", "Author1")

	mock.ExpectQuery("^SELECT (.+) FROM \"books\" WHERE (.+)").
		WithArgs(id, 1).
		WillReturnRows(mockRows)

	book, err := repo.Read(id)
	uTest.NoError(t, err)
	uTest.Equal(t, "Book1", book.Title)
}

func TestRepository_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := rTest.NewMockDB()
	uTest.NoError(t, err)

	repo := r.NewRepository(db)

	id := uuid.New()
	_ = sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(id, "Book1", "Author1")

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE \"books\" SET").
		WithArgs("Title", "Author", rTest.AnyTime{}, "", "", rTest.AnyTime{}, id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	book := &m.Book{ID: id, Title: "Title", Author: "Author"}
	rows, err := repo.Update(book)
	uTest.NoError(t, err)
	uTest.Equal(t, 1, rows)
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := rTest.NewMockDB()
	uTest.NoError(t, err)

	repo := r.NewRepository(db)

	id := uuid.New()
	_ = sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(id, "Book1", "Author1")

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE \"books\" SET \"deleted_at\"").
		WithArgs(rTest.AnyTime{}, id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	rows, err := repo.Delete(id)
	uTest.NoError(t, err)
	uTest.Equal(t, 1, rows)
}
