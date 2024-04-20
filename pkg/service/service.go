package service

import (
	"RedisLesson/dtos"
	"RedisLesson/pkg/repository"
)

type Books interface {
	Create(book dtos.Book) (int, error)
	GetAll() ([]dtos.Book, error)
	GetById(id int) (dtos.Book, error)
	Delete(id int) error
	Update(id int, input dtos.BookUpdateInput) error
}

type Service struct {
	Books
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewBookService(repo.Books),
	}
}
