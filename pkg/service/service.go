package service

import (
	"RedisLesson/dtos"
	"RedisLesson/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mock/mock.go
type Service interface {
	CreateBook(book dtos.Book) (int, error)
	GetAllBooks(pagination dtos.PaginationParams) ([]dtos.Book, error)
	GetBookById(id int) (dtos.Book, error)
	DeleteBook(id int) error
	UpdateBook(id int, input dtos.BookUpdateInput) error
	GetBooksByFilter(filter dtos.BookFilter, pagination dtos.PaginationParams) ([]dtos.Book, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}
