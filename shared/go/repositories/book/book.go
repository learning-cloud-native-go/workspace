package book

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	m "workspace.dev/shared/go/models/book"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List() (m.Books, error) {
	books := make([]*m.Book, 0)
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *Repository) Create(book *m.Book) (*m.Book, error) {
	if err := r.db.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *Repository) Read(id uuid.UUID) (*m.Book, error) {
	book := &m.Book{}
	if err := r.db.Where("id = ?", id).First(book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return book, nil
}

func (r *Repository) Update(book *m.Book) (int64, error) {
	result := r.db.Model(&m.Book{}).
		Select("Title", "Author", "PublishedDate", "ImageURL", "Description", "UpdatedAt").
		Where("id = ?", book.ID).
		Updates(book)

	return result.RowsAffected, result.Error
}

func (r *Repository) Delete(id uuid.UUID) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&m.Book{})

	return result.RowsAffected, result.Error
}
