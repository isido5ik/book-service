package repository

import (
	"RedisLesson/dtos"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	CreateBook(book dtos.Book) (int, error)
	GetAllBooks(pagination dtos.PaginationParams) ([]dtos.Book, error)
	GetBookById(id int) (dtos.Book, error)
	DeleteBook(id int) error
	UpdateBook(id int, input dtos.BookUpdateInput) error
	GetBooksByFilter(filter dtos.BookFilter, pagination dtos.PaginationParams) ([]dtos.Book, error)
}

type repository struct {
	redis redis.Client
	db    *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db:    db,
		redis: *NewRedisClient(),
	}
}
