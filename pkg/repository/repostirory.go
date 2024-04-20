package repository

import (
	"RedisLesson/dtos"

	"github.com/jmoiron/sqlx"
)

type Books interface {
	Create(book dtos.Book) (int, error)
	GetAll() ([]dtos.Book, error)
	GetById(id int) (dtos.Book, error)
	Delete(id int) error
	Update(id int, input dtos.BookUpdateInput) error
}

type Repository struct {
	// Redis *redis.Client
	Books
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Books: NewBooksPostgres(db),
		// Redis: NewRedisClient(),
	}
}
